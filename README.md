faustnet is a cli program that can perform GET, POST, and DELETE requests as a Client
Can also start a simple HTTP server(still in progress) and can serve text files to clients(i.e. other faustnet instances)


To install faustnet as a CLI command use the following command from within the project directory
```
go install ./cmd/main/faustnet.go
```

From there you can invoke faustnet by calling 
```
faustntet
```

Upon the first time of calling faustnet it will create a directory(and sub directories) for files you might export after receiving with the GET requests or (text)files you may want to host with the HTTP Server

If you would like to perform a GET request you can use 
```
faustnet httpreq -u "https://httpbin.org/get"
```
and receive this as the output
```
{
  "args": {},
  "headers": {
    "Accept": "application/json, application/yaml, application/xml, text/csv, text/html",
    "Accept-Encoding": "gzip",
    "Host": "httpbin.org",
    "User-Agent": "Go-http-client/2.0",
    "X-Amzn-Trace-Id": "REDACTED"
  },
  "origin": "REDACTED",
  "url": "https://httpbin.org/get"
}
```

To create an http server you can run

```
faustnet httpserv run
```

This is currently hardcoded to port :8080 and will ping ipify.org to retrieve your public IP Address. This should make it relatively easy to share text files with users outside your network.
After that it will refresh server uptime every 10 seconds and can closed by pressing <Ctrl+^c>

For any updates that might happen, I think it would be nice to implement some sort of websocket support. I know that Gorilla is widely used but it would be interesting to learn how to handle that without a prebuilt solution.

I am still new to Go and would like feedback if you're able to provide it
