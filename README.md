![Banner](.github/github-banner@2x.png)
# Formeler SDK

First, please check the documentation (https://formeler.com/lang-id/doc/) page to familiarize yourself with the basics of the API.
## Installation

simply import `formeler` in your code and Go will automatically fetch it during build:

```go
import "github.com/FormelerOrganization/formeler_go"
```

To get started, create a new instance of `Formeler` and pass your API key as the first argument to the constructor. You can obtain a fresh API key from the [website](https://formeler.com/lang-id/overview/) that includes 10M tokens of free credit.
Once initialized, you can use the available `Formeler` methods. Currently, language detection via the LangID module is supported. Below are examples of how to use the LangID methods.
## Example

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/FormelerOrganization/formeler_go"
)

type ResultResponse struct {
	Result     string `json:"result"`
	Message    string `json:"message"`
	Language   string `json:"language"`
	TokenCount int64  `json:"token_count"`
}

func detectExample() {
	formeler := formeler.NewFormeler("YOUR-API-KEY")
	result, err := formeler.LangID.Detect("Dies ist ein Test")
	if err != nil {
		panic(err)
	}
	var resultJSON ResultResponse
	err = json.Unmarshal([]byte(result), &resultJSON)
	if err != nil {
		panic(err)
	}
	fmt.Println(resultJSON.Result, resultJSON.Language, resultJSON.TokenCount)
}

func main() {
	detectExample()
}

```

If the API call is successful, the `result` variable will contain the following JSON response:

```json
{
  "result": "success",
  "language": "de",
  "tokenCount": "4"
}
```

### More Examples

You can find additional examples in the `examples` directory of this repository.

### API Errors

The Formeler API may return an error in the response. You can identify these by checking the `result` field in the returned object.
An error response will follow this structure:

```json
{
  "result": "failed",
  "message": "unauthorized"
}
```
