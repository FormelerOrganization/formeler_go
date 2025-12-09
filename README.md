# Formeler SDK

Formeler SDK is a cloud language detection api SDK, It's design to make get started quick and easy with the ability to
scale up to complex applications.

## An Example

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

If the API call is successful,the `result` variable will contain the following JSON response:

```json
{
  "result": "success",
  "language": "de",
  "tokenCount": "4"
}
```

### More Example

You can find additional examples in the `examples` directory of this repository.

### API errors

The Formeler API may return an error in the response. You can identify these by checking the `result` field in the
returned object.
An error response will follow this structure:

```json
{
  "result": "failed",
  "message": "unauthorized"
}
```