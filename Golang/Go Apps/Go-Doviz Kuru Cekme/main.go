package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.tcmb.gov.tr/kurlar/today.xml")
	if err != nil {
		log.Fatal("Hata: %s", err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(data))
}

type stock struct {
	XMLName      xml.Name `xml:"STOCK"`
	CURRENCYNAME string   `xml:"CURRENCYNAME"`
	FOREXBUYING  string   `xml:"FOREXBUYING"`
	FOREXSELLING string   `xml:"FOREXSELLING"`
}

type icpiyasa struct {
}
