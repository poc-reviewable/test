package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
)

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })

    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hi")
    })

    log.Fatal(http.ListenAndServe(":8081", nil))

}
In the above code, we essentially define two different Handlers. These handlers are what respond to any HTTP request that matches the string pattern we define as the first parameter. So essentially whenever a request is made for the home page or http://localhost:8081/, we’ll see our first handler respond as the query matches that pattern.

Running Our Server
Ok so now that we’ve created our own very simplistic server we can try running it by typing go run server.go in to our console. Once this is done head over to your browser and head to http://localhost:8081/world. On this page, you should hopefully see your query string echoed back to you in true “hello world” fashion.

Adding a bit of Complexity
So now that we’ve got a basic web server set up, let’s try incrementing a counter every time a specific URL is hit. Due to the fact that the web server is asynchronous, we’ll have to guard our counter using a mutex in order to prevent us from being hit with race-condition bugs.

Note - If you are unsure as to what a mutex is, don’t worry, this is just being used to highlight that these servers aren’t guarded against race conditions. If you want to learn more on mutexes, you can check out my other tutorial here: Go Mutex Tutorial

package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"
    "sync"
)

var counter int
var mutex = &sync.Mutex{}

func echoString(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello")
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
    mutex.Lock()
    counter++
    fmt.Fprintf(w, strconv.Itoa(counter))
    mutex.Unlock()
}

func main() {
    http.HandleFunc("/", echoString)

    http.HandleFunc("/increment", incrementCounter)

    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hi")
    })

    log.Fatal(http.ListenAndServe(":8081", nil))

}
