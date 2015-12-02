package main

import ( //"http"
	"fmt"
	"net/http"

//"net/rpc"
//"net/rpc/jsonrpc"
)

func hitELB(url string) error {

	//url2 := "http://maps.google.com/maps/api/geocode/json?address=655+South+Fair+Oaks+Avenue,+Sunnyvale,+CA&sensor=false"
	fmt.Println()
	fmt.Println(url)

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	/*	body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		fmt.Println(body)*/

	return err

}

func main() {
	//url := "https://query.yahooapis.com/v1/public/yql?q=select%20LastTradePriceOnly%2Csymbol%20from%20yahoo.finance.quote%20where%20symbol%20in%20(%22YHOO%22%2C%22GOOG%22)&format=json&env=store%3A%2F%2Fdatatables.org%2Falltableswithkeys&callback="

	for {
		hitELB("http://squad-loadbalancer-590657324.us-west-2.elb.amazonaws.com/hello.html")
		hitELB("http://squad-loadbalancer-590657324.us-west-2.elb.amazonaws.com/helloinstance2.html")
	}

}
