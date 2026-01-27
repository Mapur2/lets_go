package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"text/tabwriter"
)

var w *tabwriter.Writer

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	w = tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	fmt.Fprintln(w, "DOMAIN\tHAS_MX\tHAS_SPF\tHAS_DMARC\tSPF_RECORD\tDMARC_RECORD")
	w.Flush()

	for scanner.Scan() {
		checkDomain(strings.TrimSpace(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	w.Flush()
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	// MX record
	mxRecords, _ := net.LookupMX(domain)
	if len(mxRecords) > 0 {
		hasMX = true
	}

	// TXT records (for SPF)
	txtRecords, _ := net.LookupTXT(domain)
	for _, r := range txtRecords {
		if strings.HasPrefix(r, "v=spf1") {
			hasSPF = true
			spfRecord = r
			break
		}
	}

	// DMARC record
	dmarcRecords, _ := net.LookupTXT("_dmarc." + domain)
	for _, r := range dmarcRecords {
		if strings.HasPrefix(r, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = r
			break
		}
	}

	fmt.Fprintf(
		w,
		"%s\t%t\t%t\t%t\t%s\t%s\n",
		domain,
		hasMX,
		hasSPF,
		hasDMARC,
		spfRecord,
		dmarcRecord,
	)
	w.Flush()
}
