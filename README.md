#Amazon EC2 instance monitoring
##Using CloudWatch API.

The dashboard monitors various EC2 instances and plots the metrics on a graph, thus allowing a comparison of all the instances so that we know which instances are under high CPU load, memory load, network I/O, Disk I/O, as well as resource utilization by number of HTTP/TCP requests

###Technology stack:

 * Golang
 * NodeJs
 * ElasticSearch
 * Kibana

Amazon AWS SDK for GoLang is used for fetching performance data of instances and save them as array of JSON objects to a file.

Node server is used to parse the file and push data to elasticsearch.

The data is analyzed and visualized on the dashboard using Kibana.

The dashborad monitors various instances and shows comparable graphs of each metric.

Following metrics are being monitored for this project:

1. CPUUtilization
2. NetworkOut
3. NetworkIn
4. VolumeWriteOps
5. VolumeReadOps
6. RequestCount

The program supports all the metrics which are supported by aws-sdk-go

###Instructions:
1. Start elastic search server

2. Start kibana server


3. Run this script. You need to first update the file according to your GOPATH and GOROOT.

```shell
./pushdata.sh
```

The script triggers metrics.go and then server.js.

metrics.go collects stats from AWS console and writes output in .json files.
server.js parses .json and saves data into elastic search.

Elasticsearch indexes the data. For this project, mapping described in Mapping.txt is used to define schema of data stored in elasticsearch.

For front end, Kibana is used. 
Change /config.kibana.yml to update the path of elastic search server.
