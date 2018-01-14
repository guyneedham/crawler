package main

import (
	"crawler/sub"
	"flag"
        "time"
)

func main() {
        level := flag.String("level", "INFO", "logging level. Valid logging levels: {DEBUG, INFO, WARN, ERROR, CRITICAL}")
        domain := flag.String("domain", "https://xkcd.com", "domain to crawl")
        outputFile := flag.String("output-file", "output.tsv", "output file path")
        timeout := flag.String("timeout", "1m", "duration to wait for completion. Valid time units: {ns, us (or µs), ms, s, m, h}.") 
	flag.Parse()
	
        d, err := time.ParseDuration(*timeout)

        if err != nil {
            panic(err)
        }
        crawl := &crawler.Crawler{*domain, *outputFile, d, *level}
        crawl.Crawl()
}
