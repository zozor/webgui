/*
Javascript: Communicate(handlername, jsondata, successfunction(xml))
*/
package webgui

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
	"time"
)
//Resource System
var resources = map[string][]byte{}
func UseResource(files map[string][]byte) {
	resources = files
}

//Binding functions to strings
var bind = map[string]func([]byte) []byte{}


func SetHandler(key string, value func([]byte) []byte) {
	bind[key] = value
}

//Killing the server when not active
var pingChannel chan bool

func dieCounter() {
	for {
		select {
		case <-pingChannel:
			println("Pong")
			continue
		case <-time.After(15e9):
			println("Ping timeout")
			os.Exit(0)
		}
	}
}

var root string

func StartServer(addr string) { //"127.0.0.1:3939"
	http.Handle("/", http.HandlerFunc(requests))
	root, _ = os.Getwd()
	println("root:", root)

	cmd := exec.Command("x-www-browser", addr)
	cmd.Start()

	pingChannel = make(chan bool)
	go dieCounter()

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}
}

func requests(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/ping" {
		pingChannel <- true
		return
	}
	
	//Handling webgui's ajax calls
	req.ParseForm()
	ajax := req.URL.Query().Get("ajax")
	if ajax != "" {
		if _, ok := bind[ajax]; ok {
			w.Header().Set("Content-Type", "text/html; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			out := bind[ajax]([]byte(req.FormValue("data")))
			w.Write(out)
		}
		return
	}
	
	//Handling files to be displayed
	println(req.URL.Path)
	
	if req.URL.Path == "/" {
		req.URL.Path = "/index.html"
	}
	
	if req.URL.Path == "/webgui" { //<script type="text/javascript" src="/webgui"></script>
		w.Header().Set("Content-Type", "text/javascript")
		w.WriteHeader(http.StatusOK)
		w.Write(fileJQuery)
		w.Write([]byte("\n\n"))
		w.Write(fileWebguijs)
		return
	}

	setContentType(w, req.URL.Path)

	//Open file
	f, err := os.Open(root + req.URL.Path)
	if err == nil {
		defer f.Close()
		w.WriteHeader(http.StatusOK)
		io.Copy(w, f)
		return
	}
	
	//If file not found or other error - use try resource
	b, ok := resources[req.URL.Path]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Page Not Found - 404"))
		return
	}
	
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func setContentType(w http.ResponseWriter, url string) {
	switch path.Ext(url) { //Add content type på forespørgelserne
	case ".html":
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	case ".js":
		w.Header().Set("Content-Type", "text/javascript")
	case ".css":
		w.Header().Set("Content-Type", "text/css")
	}
}

func WriteJSON(w http.ResponseWriter, v interface{}) error {
	n := json.NewEncoder(w)
	err := n.Encode(v)
	if err != nil {
		return err
	}
	return nil
}

