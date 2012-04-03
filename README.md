Simple GUI toolkit with go
------------------------------

Install
===========================

`go get github.com/zozor/webgui`


The idea
===========================

You import `webgui` and set "handlers" that can be called with javascript, after that you start the server.
When the server is started, your browser will be opened and will request index.html. If you close the browser, the server will die from no pings after 15 seconds.

in index.html you use 

`<script type="text/javascript" src="webgui"></script>`

this loads jquery and the webgui function Communicate into the browser, which lets you call your handlers in go. 
It also loads a ping function that makes sure the server is closed when the index.html page is not active anymore.

It is possible to use resource packing with

`webgui.UseResource(files map[string][]byte)`

if the client requests a file that does not exist, it looks in this map, before returning 404


API
===========================
##Go functions

`webgui.SetHandler(handlername string, handler func([]byte) []byte)`

>This sets a handler that can be called from javascript

>`handlername` is the name which identifies this handler

>`handler` is the handler function

`webgui.StartServer(addr string)`

>This function starts the server. `addr` should for normal use be `localhost:someport`

`webgui.UseResource(files map[string][]byte)`

>If you want to pack resources into your binary (to make a standalone application), use this function. If the server cannot find
"/index.html" as a local file, it will do a lookup `files["/index.html"]` in the map, before returning 404.

##Javascript functions

`Communicate(handlername, data, returntype, function(data))`

>`handlername` calls the handler with this name

>`data` is the data sent to the handler, it should either be a string or json obj

>`returntype` can be whatever is used by $.ajax() in jquery. "text" (returns plain text) or "json" (returns json object) are good choices for go.

>`function(data)` is the "success" function, data is the data returned from the go handler, its type is specified in returntype

If returntype and what the go handler returns does not match or server has exited, an `alert` with the error will be called.
This could be better in the future.

##Notes

Use the examples to understand how it works.

This library is only compatible with linux at the moment, But should be pretty easy to fix for windows. The server uses the command `x-www-browser` to open the browser in `webgui.go`, you only need to change the path to your own browser and then `go install`