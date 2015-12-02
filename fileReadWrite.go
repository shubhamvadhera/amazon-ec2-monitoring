package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"
)

func doOps(dir string, interval int) {
	for {
		readFile := "/" + dir + "/foo.txt"
		writeFile := "/" + dir + "/bar.txt"
		path, _ := filepath.Abs(readFile)
		b, _ := ioutil.ReadFile(path)

		path, _ = filepath.Abs(writeFile)
		ioutil.WriteFile(path, b, 0644)
		fmt.Println("running operation on", dir)

		time.Sleep(time.Minute * time.Duration(interval))
	}
}

func main() {
	doOps("data_xvdb", 15)
}
