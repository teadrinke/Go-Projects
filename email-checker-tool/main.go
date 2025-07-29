package main

import (
	"fmt"
	"bufio"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord")

	for scanner.Scan() {
		checkDomain(scanner.Text())
		//whatever you type goes into checkDomain function

	}

	// Check for errors after scanning
	// This is important to ensure we handle any issues with reading input.
	if err:= scanner.Err(); err!=nil {
		log.Fatal("Error:could not read from input:", err) //for non-recoverable errors
		// This will log the error and exit the program.
	}
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool //declaration of variables
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain) // Check for MX records

	if err != nil {
		log.Printf("Error: %v\n", err) // for recoverable errors 
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain) // Check for TXT records

	if err != nil {
		log.Printf("Error : %v\n", err) 
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record // Store the SPF record
			break
		}
	}

	dmarcRecords , err := net.LookupTXT("_dmarc." + domain) // Check for DMARC records

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1"){
			hasDMARC = true
			dmarcRecord = record // Store the DMARC record
			break
		}
	}

	fmt.Printf("%v, %v, %v, %v, %v, %v", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}