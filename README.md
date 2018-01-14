## Installation:
```
go get golang.org/x/net/html
go get github.com/bobesa/go-domain-util/domainutil
go get github.com/ccding/go-logging/logging
go get github.com/deckarep/golang-set
go build
```

## Usage:
```
./crawler --domain https://xkcd.com --output-file output.tsv --timeout 1m --level INFO 
```
