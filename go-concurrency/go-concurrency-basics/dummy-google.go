package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

var sleep = 1

func main() {
	if len(os.Args) > 2 {
		sleep, _ = strconv.Atoi(os.Args[2])
	}
	http.HandleFunc("/", dummySearch)
	fmt.Printf("Listening on port %s \n", os.Args[1])
	http.ListenAndServe(":"+os.Args[1], nil)
}

func dummySearch(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func(start time.Time) { fmt.Printf("Query took: %v \n", time.Since(start)) }(start)

	str := r.URL.Query().Get("q")
	fmt.Printf("Request for string=%s\n", str)

	time.Sleep(time.Millisecond * time.Duration(rand.Intn(sleep)))
	fmt.Fprint(w, response)
}

var response = `{
results: [
{
GsearchResultClass: "GwebSearch",
unescapedUrl: "https://golang.org/",
url: "https://golang.org/",
visibleUrl: "golang.org",
cacheUrl: "http://www.google.com/search?q=cache:rie1WixWbVcJ:golang.org",
title: "The Go Programming Language",
titleNoFormatting: "The Go Programming Language",
content: "On August 21st the Go community gathered in London for the first edition of <b>Golang</b> UK. The conference featured two parallel tracks and nearly 400 gophers  ..."
},
{
GsearchResultClass: "GwebSearch",
unescapedUrl: "https://github.com/golang",
url: "https://github.com/golang",
visibleUrl: "github.com",
cacheUrl: "http://www.google.com/search?q=cache:It6zosxIs54J:github.com",
title: "<b>golang</b> - GitHub",
titleNoFormatting: "golang - GitHub",
content: "sys. [mirror] Go packages for low-level interaction with the operating system. Updated 2 minutes ago. Go 61 11 · review. [mirror] Tool for working with Gerrit code ..."
},
{
GsearchResultClass: "GwebSearch",
unescapedUrl: "https://twitter.com/golang",
url: "https://twitter.com/golang",
visibleUrl: "twitter.com",
cacheUrl: "http://www.google.com/search?q=cache:KnPnz5GJJTsJ:twitter.com",
title: "Go (@<b>golang</b>) | Twitter",
titleNoFormatting: "Go (@golang) | Twitter",
content: "The latest Tweets from Go (@<b>golang</b>). Go will make you love programming again. I promise."
},
{
GsearchResultClass: "GwebSearch",
unescapedUrl: "https://www.reddit.com/r/golang",
url: "https://www.reddit.com/r/golang",
visibleUrl: "www.reddit.com",
cacheUrl: "http://www.google.com/search?q=cache:WuYKHPcyrwcJ:www.reddit.com",
title: "r/<b>Golang</b> - Reddit",
titleNoFormatting: "r/Golang - Reddit",
content: "<b>Golang</b> image analysis (youtube.com). submitted 2 hours ago by tscottmcleod · comment; share. loading... 3. 7. 8. 9. Developing API backend in Go (self.<b>golang</b>)."
}
]`
