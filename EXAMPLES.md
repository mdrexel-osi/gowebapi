# Go PI Web API Examples

We have added a WebID2 wrapper to this library. Make sure you use the WebID functions in [webid.go](webid.go) to construct WebIDs from your AF and PI object paths.

I have updated the examples here to show you how WebIDs are constructed. Please also review [webid_tests.go](webid_tests.go) to see how each different AF and PI object constucts their WebIDs.

## Setting up an HTTP Client
```go
package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	pi "github.com/christoofar/gowebapi"
)

// The http client is best shared.
var cfg = pi.NewConfiguration()
var client *pi.APIClient
var auth context.Context

func Init() {
	cfg.BasePath = "https://myhostname.mycompany.com/piwebapi"
	//TODO Fix this
	auth = context.WithValue(context.Background(), pi.ContextBasicAuth, pi.BasicAuth{
		UserName: "COMPANY\MyID",
		Password: "p@$$w0rd",
	})
	client = pi.NewAPIClient(cfg)
}

func main() {
	Init()
}
```

Here I have stood up a global http client that is ready to send requests to the PI Web API endpoint. I have replaced the "gowebapi" package name with the faster-to-type mnemonic called `pi`.

If you try to compile this, Go will immediately complain that the variable `auth` is not being used. That's because it isn't.  It is the security context used to send stateless web requests. To make the error go away we need to use `auth` in something.

## Getting an AF Attribute
At my office I have a AF database called `Chris` where I have some of my computers being monitored by PI Perfmon. I would like to see the details about the Attribute I created called `Memory % In Use` hanging off a machine called CLSAF, which also happens to be my PI Server. 

**Keep in mind that my WebIDs will not work in your environment. You can use a web browser to manually find the WebIDs in your system by going to https://yourserver/piwebapi.**


![Here's my AF](./af1.png)

To get at that attribute's detail, let's call `AttributeGet`

```go
func GetAttribute() {

	// Construct the PathOnly WebID for \\CLSAF\Chris\CLSAF|Memory % In Use

	webid := pi.EncodeWebID(pi.NewAFAttributeWebID("CLSAF\\Chris\\CLSAF|Memory % In Use", pi.IS_AF_ELEMENT))

	options := make(map[string]interface{})
	options["webIdType"] = "PathOnly"

	value, resp, err := client.AttributeApi.AttributeGet(auth, webid, options)
	if err != nil {
		log.Println("The call to WebAPI failed [", resp.StatusCode, "]")
		body, _ := ioutil.ReadAll(resp.Body)
		log.Fatal(string(body))
	}

	fmt.Println("》CLSAF\\Chris\\Memory % In Use")
	fmt.Println("  ConfigString:\t", value.ConfigString)
	fmt.Println("   Path:\t", value.Path)
	fmt.Println("   WebId:\t", value.WebId)
}
```
Add a call to this `func` in your main() and then compile and run your program. I've enhanced my program a bit more than the example, but you get the idea:
![here's the output](./af2.png)

## Recorded Values (the PIPoint way)

Here's how to pull some recorded values out of a PI Point (aka "PI tag")

```go
func GetSinusoidTag() {

	webid := pi.EncodeWebID(pi.NewPIPointWebID("CLSAF\\SINUSOID"))

	optionals := make(map[string]interface{})

	optionals["startTime"] = "*-5d"
	optionals["endTime"] = "*"
	optionals["maxCount"] = int32(5000)

	value, resp, err := client.StreamApi.StreamGetRecorded(auth, webid, optionals)
	if err != nil {
		log.Fatal("Couldn't get PI tag CLSAF\\SINUSOID [", resp.StatusCode, "] ", err)
	}

	// Let us split the data that came back into 2 arrays then print in two columns in the console.
	col1 := value.Items[0 : len(value.Items)/2]
	col2 := value.Items[len(value.Items)/2:]

	for i := 0; i < len(col1); i++ {
		fmt.Println(col1[i].Timestamp, "\t", *col1[i].Value, "\t",
			col2[i].Timestamp, "\t", *col2[i].Value)
	}
}
```

Result:
![here's the output](./af3.png)

## Recorded Values (the AF Attribute way)

Here's how to pull some recorded values out of an AF Attribute that has a DataReference.

```go
func GetRecordedValues() {

	webid := pi.EncodeWebID(pi.NewAFAttributeWebID("CLSAF\\Chris\\CLSAF|Memory % In Use", pi.IS_AF_ELEMENT))

	optionals := make(map[string]interface{})

	optionals["endTime"] = "*"
	optionals["startTime"] = "*-100d"
	optionals["maxCount"] = int32(500)

	value, resp, err := client.StreamApi.StreamGetRecorded(auth, webid, optionals)
	if err != nil {
		log.Println("The call to WebAPI failed [", resp.StatusCode, "]")
		body, _ := ioutil.ReadAll(resp.Body)
		log.Fatal(string(body))
	}

	// Let us split the data that came back into 2 arrays to print in two columns in the console.
	col1 := value.Items[0 : len(value.Items)/2]
	col2 := value.Items[len(value.Items)/2:]

	for i := 0; i < len(col1); i++ {
		fmt.Println(col1[i].Timestamp, "\t", *col1[i].Value, "\t",
			col2[i].Timestamp, "\t", *col2[i].Value)
	}
}
```

Notice in the code that to retrieve the PIPoint values a pointer derefrence must occur. Some variables in golang will be memory pointers which you must dereference. Usually any variable that will accept a variety of structures will be by reference rather than by value.

A rule of thumb: If you print a variable to the console get back something like _0x00005244fff..._, you need to put a `*` in front of the variable to get at the array, map, or struct that is behind it.

Result:
![here's the output](./af4.png)

Notice how the Go `map` datatype is able to handle a mix of data value types coming back. A `map` is a name:value pair data type in golang.  Not everything will be a floating point number. This means you need to check what comes back.

## Batch Requests
Strangely enough these are the simplest things in the world to implement in Go.

There's only one method `Execute` which takes an array of stuff.

### Making a batch request
Let's send two requests at once to pull recorded values from the server. To implement the individual request batches we need to define a struct where we can put our batch request content, which I've called `BatchRequest`. If I don't fill something out in the `BatchRequest` struct I want to make sure the JSON for the empty fields do not get sent. You can attach hints to struct fields like this:

```go

type BatchRequest struct {
	Resource string `json:"Resource,omitempty"`
	Method   string `json:"Method,omitempty"`
	Content  string `json:"Content,omitempty"`
}

func BatchExample() {

	webid := pi.EncodeWebID(pi.NewPIPointWebID("CLSAF\\SINUSOID"))	
	webid2 := pi.EncodeWebID(pi.NewPIPointWebID("CLSAF\\CDT158"))	

	batches := make(map[string]BatchRequest)
	
	// Setup request 1
	var batch1 BatchRequest
	batch1.Method = "GET"
	batch1.Resource = "https://localhost/piwebapi/streams/" + webid + "/recorded"

	// Setup Request 2
	var batch2 BatchRequest
	batch2.Method = "GET"
	batch2.Resource = "https://localhost/piwebapi/streams/" + webid2 + "/recorded"

	// Stick 'em in the map
	batches["1"] = batch1
	batches["2"] = batch2

	// Ship it!!!
	value, resp, err := client.BatchApi.BatchExecute(auth, batches)
	if err != nil {
		log.Fatal("Batch request failed.  [", resp.StatusCode, "] ", err)
	}

	// Succeeded, now process what we got back
	fmt.Println(value)
	fmt.Println("Request completed.")

}
```

Result:

```
map[1:{200 map[Content-Type:application/json; charset=utf-8] 0xc420304090} 
2:{200 map[Content-Type:application/json; charset=utf-8] 0xc420304220}]
Request completed.
~/go/piwebapitest $ 
```

### Dealing with dynamic data returns

As you can see, we got back a memory address for the data (0xc420304090 and 0xc420304220), which means we were given pointers to some structure since those are memory addresses. Let's call `fmt.Println()` again but this time we will do a print of each dereferenced item inside the map.

```go
	// Succeeded

	for index, element := range value {
		fmt.Println("Index:", index, "Status:", element)
		fmt.Println("   Headers: ")
		for index, header := range element.Headers {
			fmt.Println("\t", index, ":\t", header)
		}
		fmt.Println("   Content: ")
		fmt.Println("\t", *element.Content)
	}
```

Now, let's look at the output:
![here's the output](./af5.png)

My that's a lot of output! Since Batch output can be literally anything, you can use the Go Delve debugger to figure out the data structure that was returned to you in the `gowebapi.Response` type.

Let's change this routine so we loop through the information the server sent back.

```go

	for index, element := range value {
		fmt.Println("Index:", index, "Status:", element)
		fmt.Println("   Headers: ")
		for index, header := range element.Headers {
			fmt.Println("\t", index, ":\t", header)
		}
		fmt.Println("   Content: ")
		fmt.Println("\t", *element.Content)

		fmt.Println(*element.Content)
	}
```

By running the [Go Delve](https://github.com/derekparker/delve) debugger and setting a breakpoint at the last `Println` statement we can probe what the `Content` is. It is this:

![debugger output](./af6.png)

So it looks like for one of the two batches I got back 232 items. We know they're `time:value` records from the PI Data Archive, but here we see yet more memory addresses - so we need to deference these as well. Let's do that, but to even get there we need to dereference the element list we have using a technique called [Type assertions](https://tour.golang.org/methods/15) which you can use to figure out what you have, like this:

![more output](./af7.png)

Aha, the data is structured! It looks like we get back several layers of nested maps, which the `json` library in Go converts to arrays of run-time interfaces.

We're seeing more pointers to more dictionary values. That's easy, we can loop through these to print them all out, dereferenced. But since they're hiding inside `interface{}` we'll have to use type assertions yet again to remind Go that we got back nested dictionaries of things.

![more output](./af8.png)

My, Batch can be complicated! It's all structured with references and it takes three levels of `for` loops to scan through all the results you got back - which is a perfect job for a goroutine to handle by itself.

Since Delve told me what type these arrays of things are, I used those hints from the debugger to write this code to exract out the data.

### Let's make the dynamic data strongly-typed again

Since the Batch controller literally returns "stuff" the shape of the data return is only discovered after you've formulated the data request.

We also now know that when we made two calls to get streams from two PI tags - using the debugger, we saw we got a stream of PI Point values back. However, they came in the form of a `map[string]interface{}`, which is literally a dictionary of unknown things.

You *could* use a `switch` statement to loop through all the indexes, compare the names of the data fields in the JSON and then do one final type assertion to the value itself. Personally, I think that's gross and creates ugly code.

Luckily there is this awesome go package called [`mapstructure`](github.com/mitchellh/mapstructure) which will cast your crazy dynamic maps into `structs`.

So, let's grab mapstructure:

```
~/go/piwebapitest $ go get -u github.com/mitchellh/mapstructure
```

Now, let's add a new struct to get the PI archive values...

```go
type BatchTimedValues struct {
	Timestamp string
	Good      bool
	Annotated bool
	Value     *interface{}
}
```

We have to make `Timestamp` a string here because the Batch controller doesn't send the time back with milliseconds. We will use the `time` package to do a nice casting of the time and put it in a local time zone.

Then, we'll transform the `map[string]interface{}` into a `BatchTimedValues` struct so we can refer to the PI data. Now, the completed `BatchRequest()` function looks like this

```go

type BatchRequest struct {
	Resource string `json:"Resource,omitempty"`
	Method   string `json:"Method,omitempty"`
	Content  string `json:"Content,omitempty"`
}

type BatchTimedValues struct {
	Timestamp string
	Good      bool
	Annotated bool
	Value     *interface{}
}


func BatchExample() {

	webid := pi.EncodeWebID(pi.NewPIPointWebID("CLSAF\\SINUSOID"))	
	webid2 := pi.EncodeWebID(pi.NewPIPointWebID("CLSAF\\CDT158"))	

	batches := make(map[string]BatchRequest)
	var batch1 BatchRequest
	batch1.Method = "GET"
	batch1.Resource = "https://localhost/piwebapi/streams/" + webid + "/recorded"

	var batch2 BatchRequest
	batch2.Method = "GET"
	batch2.Resource = "https://localhost/piwebapi/streams/" + webid2 + "/recorded"

	batches["1"] = batch1
	batches["2"] = batch2

	value, resp, err := client.BatchApi.BatchExecute(auth, batches)
	if err != nil {
		log.Fatal("Batch request failed.  [", resp.StatusCode, "] ", err)
	}

	// Succeeded

	for index, element := range value {
		fmt.Println("Index:", index, "Status:", element)
		fmt.Println("   Headers: ")
		for index, header := range element.Headers {
			fmt.Println("\t", index, ":\t", header)
		}
		fmt.Println("   Content: ")

		data := (*element.Content).(map[string]interface{})

		layout := "2006-01-02T15:04:05Z"
		location, _ := time.LoadLocation("America/New_York")

		for _, dataitem := range data["Items"].([]interface{}) {

			var pidata BatchTimedValues
			mapstructure.Decode(dataitem, &pidata)

			pitime, _ := time.Parse(layout, pidata.Timestamp)

			fmt.Println(pitime.In(location),
				*pidata.Value,
				"Annotated:", pidata.Annotated,
				"Good:", pidata.Good)

		}

		fmt.Println("----------", index, "--------------")

	}

	fmt.Println("Request completed.")

}
```

And here's the output:
![here's the output](./af10.png)
