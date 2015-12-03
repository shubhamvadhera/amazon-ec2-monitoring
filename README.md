Amazon Ec2 instance monitoring using data fetched through CloudWatch API.

The dashboard indicates which instances are under high CPU load, memory load, network I/O, Disk I/O.

Technologies used:

GoLang
NodeJs
ElasticSearch
Kibana

Amazon AWS SDK of GoLang is used for fetching performance data of instances and save them as array of JSON objects to a file.

Node server is used to parse the file and save data on elasticsearch

The data is analyzed and visualized on the dashboard using Kibana.

The dashborad monitors various instances and shows comparable graphs of each metrice.

Following metrices are being monitored:
1. CPUUTilization
2. NetworkOut
3. NetworkIn
4. VolumeWriteOps
5. VolumeReadOps
6. RequestCount

To run the program, first run the script. You need to first update the file according to your GOPATH and GOROOT.

./pushdata.sh

The script first run metrics.go and then server.js.
metrics.go collects stats from AWS console and writes output in .json files.
server.js parses .json and saves data into elastic search.

Elasticsearch indexes the data. For this project, mapping described in Mapping.txt is used to define schema of data stored in elasticsearch.

For front end, Kibana is used. Change /config.kibana.yml to update the path of elastic search server.



 



