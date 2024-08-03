package modules

import (
	"os"
	"fmt"
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
)

func Lookupprice(army, unitname string){
	urlsuffix := army+"-"+unitname
	urlprefix := "https://www.warhammer.com/en-US/plp?search=lord-kroak"
	url := urlprefix+urlsuffix
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: jar,
	}
	resp, err := client.Get(url)
	if err != nil {
	   log.Fatalln(err)
	}
 //We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
	   log.Fatalln(err)
	}
 //Convert the body to type string
	sb := string(body)
	log.Printf(sb)
}

func Csviterate(){
	file, err := os.Open("units.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()


	reader := csv.NewReader(file)

	lineNumber := 0

	records := make ([][]string, 0)

	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("err while reading CSV file: %s", err)
		}
		if lineNumber == 0 {
			lineNumber++
			continue
		}
		lineNumber++

		records = append(records,record)
	}

	for _, v := range records {
		if len(v) != 2 {
			// malformation of file ignore record
			continue
		}
		Lookupprice(v[1], v[0])
	}
}