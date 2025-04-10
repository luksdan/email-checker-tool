package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	file, err := os.Create("result.csv")
	if err != nil {
		log.Fatalf("Error by creating new file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	writer.Write([]string{"domain", "hasMX", "hasSPF", "sprRecord", "hasDMARC", "dmarcRecord"})
	defer writer.Flush()

	for scanner.Scan() {
		domain := strings.TrimSpace(scanner.Text())
		if domain == "" {
			continue
		}
		result := checkDomain(domain)
		writer.Write(result)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: Could not read from input: %v\n", err)
	}
}

func checkDomain(domain string) []string {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	hasMX = len(mxRecords) > 0

	txtRecord, err := net.LookupTXT(domain)
	if err != nil {
		log.Fatal("Error: %v\n", err)
	}

	for _, v := range txtRecord {
		if strings.HasPrefix(v, "v=spf1") {
			hasSPF = true
			spfRecord = v
			break
		}
	}

	dmarcRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for _, v := range dmarcRecords {
		if strings.HasPrefix(v, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = v
			break
		}
	}

	return []string{
		domain,
		fmt.Sprintf("%t", hasMX),
		fmt.Sprintf("%t", hasSPF),
		spfRecord,
		fmt.Sprintf("%t", hasDMARC),
		dmarcRecord,
	}
}
