# Domain Scanner (MX, SPF, DMARC)
This is a Go utility that automatically checks whether one or more domains have MX, SPF, and DMARC records and exports the results to a CSV file.

## Features
- Checks if the domain has an MX (Mail Exchange) record

- Detects if the domain has an SPF (Sender Policy Framework) record and captures its value

- Detects if the domain has a DMARC record and captures its value

- Exports the results to the result.csv file

## Requirements
- Go 1.16 or higher installed

## Installation
Clone the repository and build the project:
```bash
git clone https://github.com/luksdan/email-checker-tool.git
cd email-checker-tool
go build -o email-checker-tool
```
## How to use
Run the program and enter the domains (one per line). Finish by pressing Ctrl+D (Linux/Mac) or Ctrl+Z + Enter (Windows):
```bash
./email-checker-tool
google.com
example.com
your-domain.com
```
## Example output (result.csv)
```csv
domain,hasMX,hasSPF,spfRecord,hasDMARC,dmarcRecord
google.com,true,true,v=spf1 include:_spf.google.com ~all,true,v=DMARC1; p=none; rua=mailto:dmarc-reports@google.com
example.com,false,false,,false,
```