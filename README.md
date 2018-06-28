# Go Client for OSIsoft PI Web API

This is the Go programming language interface to PI Web API.

It was built against the **PI Web API 2018** specification.

## Installation

To compile Go programs you will need the [Go installation kit](https://golang.org/dl/).

Issue the following command to install this package:

```
go get -u github.com/christoofar/gowebapi
```

## Getting Started

Here is a sample Go program for retrieving the links from the Web API home page.  

Create a directory under `%GOPATH%` and let's call it `webapitest`.   Then create a new code file with the name `webapitest.go`

This will print all the version numbers of your PI Web API server plugins.  Replace the values in braces with the appropriate values for your environment.

```go
// webapitest.go
// webapitest.go
package main

import (
	"context"
	"fmt"
	"log"

	pi "github.com/christoofar/gowebapi"
)

var cfg = pi.NewConfiguration()

var client *pi.APIClient
var auth context.Context

func Init() {
	cfg.BasePath = "https://{your web api server here}/piwebapi"

	auth = context.WithValue(context.Background(), pi.ContextBasicAuth, pi.BasicAuth{
		UserName: "{user name here}",
		Password: "{password here}",
	})

	client = pi.NewAPIClient(cfg)
}

func main() {
	response, _, fail := client.SystemApi.SystemVersions(auth)
	if fail != nil {
		log.Fatal(fail)
	}

	fmt.Println("Here's all the plugin versions on PI Web API")
	for i := range response {
		fmt.Println(i, response[i].FullVersion)
	}
}
```

You can run the program by issuing the following commands

```
~/go/webapitest $ go build
~/go/webapitest $ ./webapitest
```

Your output should look something like this

```
~/go/webapitest $ ./webapitest 
Here's all the plugin versions on PI Web API
OSIsoft.REST.Documentation 1.11.0.967
OSIsoft.REST.Services 1.11.0.967
OSIsoft.Search.SvcLib 1.8.0.3651
OSIsoft.PIDirectory 1.0.0.0
OSIsoft.REST.Core 1.11.0.967
```

## Documentation for API Endpoints

If you'd like to read up on the Golang-specific data types that come back from this web client, [visit the documentation](/docs/README.md) under the `/docs` folder.

You can also use your own PI Web API server to read documentation by visiting your PI Web API server endpoint at `https://your.webapi.server/piwebapi` in any web browser.

## Licensing
Copyright 2018 OSIsoft, LLC.

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

`
   http://www.apache.org/licenses/LICENSE-2.0
`

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

Please see the file named [LICENSE.md](LICENSE.md).
