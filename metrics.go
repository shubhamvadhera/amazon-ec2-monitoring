package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

/*-----Debugging variables-----*/

var out io.Writer
var debugModeActivated bool

/*-----Structs-----*/

// for storing instance details
type instance struct {
	name      string
	namespace string
	typ       string
	region    string
	id        string
	secret    string
}

//MetricWriteStruct struct to write in file
type MetricWriteStruct struct {
	InstanceName string  `json:"instanceName"`
	Region       string  `json:"region"`
	PostDate     string  `json:"post_date"`
	MetricType   string  `json:"metric_type"`
	MetricValue  float64 `json:"metric_value"`
	Unit         string  `json:"unit"`
}

//stores metric response alongwith the instance
type instanceMetricResponse struct {
	resp      *cloudwatch.GetMetricStatisticsOutput
	ins       instance
	statistic string
}

/*--- Constants ---*/
const accessKeyID string = "-------------------------"
const secretAccessKey string = "-------------------------------"

/*--- Global Variables ---*/
var instance1 = instance{name: "i-b56a2571", typ: "InstanceId", region: "us-west-2"}
var instance2 = instance{name: "i-45cf8281", typ: "InstanceId", region: "us-west-2"}
var instance3 = instance{name: "i-30aae5f4", typ: "InstanceId", region: "us-west-2"}
var instance4 = instance{name: "i-0bf021cf", typ: "InstanceId", region: "us-west-2"}
var volumeInstance1 = instance{name: "vol-f0765131", typ: "VolumeId", region: "us-west-2"}
var loadBalancer1 = instance{name: "SQUAD-LoadBalancer", typ: "LoadBalancerName", region: "us-west-2"}

/*--- Functions ---*/

//construct function acts as program constructor
func construct() {
	debugModeActivated = true //change to true to see all developer messages
	out = ioutil.Discard
	if debugModeActivated {
		out = os.Stdout
	}
	deleteOldFiles()
}

//deletes old files
func deleteOldFiles() {
	dir := "./"
	path, err := filepath.Abs(dir)
	check(err)
	dirAbs, err := os.Open(path)
	check(err)
	defer dirAbs.Close()

	files, err := dirAbs.Readdir(-1)
	check(err)

	for _, file := range files {
		if file.Mode().IsRegular() {
			if filepath.Ext(file.Name()) == ".json" {
				os.Remove(file.Name())
				fmt.Println("Deleted ", file.Name())
			}
		}
	}
}

//to check each error
func check(e error) {

	if e != nil {
		fmt.Println(e.Error())
		panic(e)
	}
}

//writes a .json file type in Kibana format
func writeMetricResp(imrA []instanceMetricResponse) {
	var mwsA []MetricWriteStruct
	for _, z := range imrA {
		for _, k := range z.resp.Datapoints {
			var mws MetricWriteStruct
			mws.InstanceName = z.ins.name
			mws.Region = z.ins.region
			temp := k.Timestamp.UnixNano() / int64(time.Millisecond)
			mws.PostDate = strconv.FormatInt(temp, 10)
			mws.MetricType = *z.resp.Label
			if z.statistic == "Maximum" {
				mws.MetricValue = *k.Maximum
			} else if z.statistic == "Average" {
				mws.MetricValue = *k.Average
			} else if z.statistic == "Sum" {
				mws.MetricValue = *k.Sum
			} else {
				panic("unknown statistic")
			}
			mws.Unit = *k.Unit
			mwsA = append(mwsA, mws)
		}
		fmt.Fprintln(out, "Instance: ", z.ins.name, "metric resp: ", z.resp)
	}
	jsonOut, err := json.Marshal(mwsA)
	check(err)

	writeToFile(jsonOut, *imrA[0].resp.Label)

	fmt.Println(*imrA[0].resp.Label, "written to file @ ", time.Now())
}

//writes any bytes to any filename in current folder
func writeToFile(b []byte, filename string) {
	file := "./" + filename + ".json"
	path, err := filepath.Abs(file)
	check(err)

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
	check(err)

	defer f.Close()

	fmt.Fprintln(out, "writing", string(b), "to file", file)
	_, err = f.WriteString(string(b))
	check(err)

	return
}

//writes given metric type for given instances to a file
func metricMonitor(metricName string, namespace string, statistic string, unit string, inss ...instance) {

	var imrA []instanceMetricResponse
	for _, ins := range inss {
		var imr instanceMetricResponse
		imr.resp = getMetrics(ins, metricName, namespace, statistic, unit)
		imr.ins = ins
		imr.statistic = statistic
		imrA = append(imrA, imr)
	}
	writeMetricResp(imrA)
	return
}

//returns metrics for one instance and one metric type
func getMetrics(ins instance, metricName, namespace, statistic, unit string) *cloudwatch.GetMetricStatisticsOutput {
	svc := cloudwatch.New(session.New(), &aws.Config{Region: aws.String(ins.region), Credentials: credentials.NewStaticCredentials(accessKeyID, secretAccessKey, "")})

	now := time.Now()
	then := now.Add(-24 * time.Hour) //duration
	params := &cloudwatch.GetMetricStatisticsInput{
		EndTime:    aws.Time(now),          // Required
		MetricName: aws.String(metricName), // Required
		Namespace:  aws.String(namespace),  // Required
		Period:     aws.Int64(3600),        // Required //interval
		StartTime:  aws.Time(then),         // Required
		Statistics: []*string{ // Required
			aws.String(statistic), // Required
			// More values...
		},
		Dimensions: []*cloudwatch.Dimension{
			{ // Required
				Name:  aws.String(ins.typ),  // Required
				Value: aws.String(ins.name), // Required
			},
			// More values...
		},
		Unit: aws.String(unit),
	}
	resp, err := svc.GetMetricStatistics(params)
	check(err)
	return resp
}

func main() {
	construct()
	metricMonitor("CPUUtilization", "AWS/EC2", "Maximum", "Percent", instance1, instance2, instance3, instance4)
	metricMonitor("NetworkOut", "AWS/EC2", "Maximum", "Bytes", instance1, instance2, instance3, instance4)
	metricMonitor("NetworkIn", "AWS/EC2", "Maximum", "Bytes", instance1, instance2, instance3, instance4)
	metricMonitor("VolumeWriteOps", "AWS/EBS", "Average", "Count", volumeInstance1)
	metricMonitor("VolumeReadOps", "AWS/EBS", "Average", "Count", volumeInstance1)
	metricMonitor("MemoryUtilization", "System/Linux", "Maximum", "Percent", instance1, instance2, instance3, instance4)
	metricMonitor("RequestCount", "AWS/ELB", "Sum", "Count", loadBalancer1)

	fmt.Println("Program finished")
}
