# Pattern for Stream of Events in Golang

1. Create Channel as Stream:  
```
   stream := make(chan ...)
```
2. Populate Channel with Events/Stream-Data in Goroutine: 
```
   go emit(chan ...)
```   
3. Consume Stream-Data by subscribing to stream/channel:
```
   go subscribe(stream)
