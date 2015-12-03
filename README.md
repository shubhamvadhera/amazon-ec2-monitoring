Amazon Ec2 instance monitoring using data fetched through CloudWatch API.
The dashboard indicates which instances are under high CPU load, memory load, network I/O, Disk I/O.
Technologies used:
GoLang
NodeJs
ElasticSearch
Kibana

Amazon AWS SDK of GoLang is used for fetching performance data of instances and save them as array of JSON objects to a file.

Node server is used to parse the file and save data on elasticsearch

The data is analyzed and visualized on the dashboard using Kibana

