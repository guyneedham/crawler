package crawler

import (
    "bufio"
    "fmt"
    "github.com/bobesa/go-domain-util/domainutil"
    "github.com/ccding/go-logging/logging"
    "net/http"
    "os"
    "sync"
    "time"
)

var logger, _ = logging.SimpleLogger("crawler")

var encountered = make(map[string]bool)
var mapLock = &sync.Mutex{}
var writeLock = &sync.Mutex{}

type Page struct {
    url       string
    links     []string
    resources []string
}

type Crawler struct {
    Domain string 
    Path string
    Duration time.Duration
    LoggingLevel string
}

func (args *Crawler) Crawl() {
    logger.SetLevel(logging.GetLevelValue(args.LoggingLevel))

    f, err := os.Create(args.Path); if err != nil {
    	panic(err)
    }
    w := bufio.NewWriter(f)
    
    defer f.Close()
    defer w.Flush()

    lines := make(chan string)
    links := make(chan string)

    var wg sync.WaitGroup

    ticker := time.NewTicker(time.Second * 5)
    go func() {
        for {
            select {
                case t := <- ticker.C:
                    logger.Info(t.String() + " channel length: " +
                        fmt.Sprintf("%d", len(links)), ", channel capacity: " +
                        fmt.Sprintf("%d", cap(links)))
                case link := <- links:
                    logger.Debug("got link: "+link)
                    wg.Add(1)
                    go processLink(link, lines, links, &wg)
                case line := <- lines:
                    w.WriteString(line)
                 
            }
        }
    }()
  

    logger.Info("started processing...") 
 
    links <- args.Domain // kick off
    <- time.After(args.Duration)
    logger.Info("timeout reached, waiting for threads")
    wg.Wait()
    ticker.Stop()
}

func processLink(address string, toWrite chan string, toProcess chan string, wg *sync.WaitGroup) {
    if getEnc(address) {
    	return
    }

    setEnc(address)
    logger.Debug("working on "+address)
    p, err := pullPage(address)
    // where the Get caused an error, return
    if err != nil {
        wg.Done()
        return
    }

    toWrite <- fmt.Sprintf("%s\t%s\t%s\n", p.url, p.links, p.resources)
    for _, link := range p.links {
        if !getEnc(link) {
            logger.Debug("preparing to push "+link)
            toProcess <- link
            logger.Debug("pushed "+link+" into channel") 
        }  
    }
    wg.Done()
}

func extractDomain(address string) string {
    return domainutil.Domain(address)
}

func sameDomain(a1 string, a2 string) bool {
    return extractDomain(a1) == extractDomain(a2)
}

func pullPage(address string) (*Page, error) {
    resp, err := http.Get(address)
    // if an exception is thrown, return a blank Page
    if err != nil {
        logger.Error("error pulling page "+address)
        return nil, err
    }
    defer resp.Body.Close() // push closing resource onto stack
    linkSet, resourceSet := Find(resp.Body, extractDomain(address))
    return &Page{
        url:       address,
        links:     linkSet,
        resources: resourceSet,
    }, nil
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
