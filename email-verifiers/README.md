Domain Email DNS Checker (Go)

A fast and simple Go CLI tool to check whether a domain has:

✅ MX record (Mail server configured)

✅ SPF record (Sender Policy Framework)

✅ DMARC record (Anti-spoofing policy)

It prints the result in a nice table format.

🚀 Features

Reads domains from stdin (file or manual input)

Checks:

MX

SPF

DMARC

Displays:

Boolean status

Actual SPF & DMARC records

Outputs in a pretty aligned table

Streams results live

🛠️ Requirements

Go 1.18+

Internet connection (for DNS lookups)