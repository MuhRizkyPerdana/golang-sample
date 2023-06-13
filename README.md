
#testing

# golang-sample
Golang web server sample app
Port: 8000

## Docker Run
```
# docker run --name goapp -p 8000:8000 MuhRizkyPerdana/goapp:13

Application running on port :8000
```


## Running Binary

Download on [Release page](https://github.com/aditirvan/golang-sample/releases)


### Print logs to stdout/stderr
`./goapp`
### Arguments
```
--log-file : Log app to saved to file (ex: goapp.log)
--trace-url=http://tempo-host-url:port : enable tracing with grafana tempo
```
