package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/belovai/goipapi"

	"github.com/pborman/getopt/v2"
)

var help *bool
var format *string
var client *goipapi.Client

type IpInfo struct {
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	District    string  `json:"district"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	ISP         string  `json:"isp"`
	Org         string  `json:"org"`
	AS          string  `json:"as"`
	Query       string  `json:"query"`
	Reverse     string  `json:"reverse"`
	Mobile      string  `json:"mobile"`
	Proxy       string  `json:"proxy"`
}

func main() {
	setOptions()
	args := getopt.Args()

	if *help {
		getopt.Usage()
		os.Exit(0)
	}

	if len(args) == 0 {
		info, _ := os.Stdin.Stat()

		if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
			getopt.Usage()
			os.Exit(1)
		}

		reader := bufio.NewReader(os.Stdin)
		for {
			record, _, err := reader.ReadLine()
			if err == io.EOF {
				break
			}

			args = append(args, string(record))
		}
	}

	if *format == "json" || *format == "pretty" {
		client = goipapi.NewClient("json")
	} else {
		client = goipapi.NewClient(*format)
	}

	for _, ip := range args {
		resp, err := client.LookupIP(ip)
		if err != nil {
			log.Fatal(err)
		}

		if *format == "pretty" {
			print(resp)
			continue
		}

		fmt.Print(resp)
	}

}

func setOptions() {
	help = getopt.BoolLong("help", 'h', "Shows this help")
	format = getopt.EnumLong("format", 'f', []string{"pretty", "json", "xml", "csv", "line", "php"}, "pretty", "Format")
	getopt.Parse()
}

func print(data string) {
	var ipInfo IpInfo
	json.Unmarshal([]byte(data), &ipInfo)

	fmt.Printf("%s info\n\n", ipInfo.Query)
	fmt.Printf("Country: %s (%s)\n", ipInfo.Country, ipInfo.CountryCode)
	fmt.Printf("Region: %s (%s)\n", ipInfo.Region, ipInfo.RegionName)
	fmt.Printf("City: %s\n", ipInfo.City)
	fmt.Printf("Zip: %s\n", ipInfo.Zip)
	fmt.Printf("District: %s\n", ipInfo.District)
	fmt.Printf("GPS: %f %f\n", ipInfo.Lat, ipInfo.Lon)
	fmt.Printf("Timezone: %s\n", ipInfo.Timezone)
	fmt.Printf("\n")
	fmt.Printf("ISP: %s\n", ipInfo.ISP)
	fmt.Printf("AS: %s\n", ipInfo.AS)
	fmt.Printf("Org: %s\n", ipInfo.Org)
	fmt.Printf("Mobile: %s\n", ipInfo.Mobile)
	fmt.Printf("Reverse: %s\n", ipInfo.Reverse)
	fmt.Printf("Proxy: %s\n", ipInfo.Proxy)
	fmt.Printf("\n======================\n")
}
