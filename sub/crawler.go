package crawler

import (
	"bufio"
	"fmt"
	"github.com/bobesa/go-domain-util/domainutil"
	"net/http"
	"os"
	"sync"
)

var encountered = make(map[string]bool)
var mapLock = &sync.Mutex{}
var writeLock = &sync.Mutex{}

type Page struct {
	url       string
	links     []string
	resources []string
}

func extractDomain(address string) string {
	return domainutil.Domain(address)
}

func sameDomain(a1 string, a2 string) bool {
	return extractDomain(a1) == extractDomain(a2)
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func pullPage(address string) *Page {
	resp, err := http.Get(address)
	// if an exception is thrown, return a blank Page
	if err != nil {
		return new(Page)
	}
	defer resp.Body.Close() // push closing resource onto stack
	linkSet, resourceSet := Find(resp.Body, extractDomain(address))
	return &Page{
		url:       address,
		links:     linkSet,
		resources: resourceSet,
	}
}

func getEnc(address string) bool {
	mapLock.Lock()
	enc := encountered[address]
	mapLock.Unlock()
	return enc
}

func setEnc(address string) {
	mapLock.Lock()
	encountered[address] = true
	mapLock.Unlock()
}

func write(w *bufio.Writer, message string) {
	writeLock.Lock()
	w.WriteString(message)
	writeLock.Unlock()
}

func processLink(address string, w *bufio.Writer) {
	if !getEnc(address) {
		setEnc(address)
		p := pullPage(address)
		// where the Get caused an error, url will be empty
		if p.url == "" {
			return
		}
		write(w, fmt.Sprintf("%s\t%s\t%s\n", p.url, p.links, p.resources))
		// using WaitGroup to ensure program exit waits for goroutines
		// links are followed async
		var wg sync.WaitGroup
		for _, link := range p.links {
			if !getEnc(link) {
				wg.Add(1)
				go func(link string, w *bufio.Writer) {
					// once processLink returns, will be marked as done
					defer wg.Done()
					processLink(link, w)
				}(link, w)
			}
		}
		wg.Wait()
	}
}

func Run(startPage string, path string) {
	f, err := os.Create(path)
	checkError(err)
	w := bufio.NewWriter(f)
	defer f.Close()
	processLink(startPage, w)
	w.Flush()
}
