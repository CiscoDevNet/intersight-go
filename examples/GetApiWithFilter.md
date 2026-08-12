### Example for GET API with filter

``` go
package example

import (
	"fmt"
	"log"
	"os"
)

func GetObjectListWithFilter(config *Config) {
	filter := "CreateTime gt 2021-08-29T21:58:33Z"
	orderby := "CreateTime"
	top := int32(5)
	skip := int32(1)
	select_ := "CreateTime,ModTime"
	count := false
	inlinecount := "allpages"

	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	apiResponse, r, err := apiClient.ComputeApi.GetComputeRackUnitList(ctx).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Count(count).Inlinecount(inlinecount).Execute()
	if err != nil {
		log.Printf("Error when calling `ComputeApi.GetComputeRckUnitList``: %v\n", err)
		log.Printf("Full HTTP response: %v\n", r)
		return
	}

	fmt.Fprintf(os.Stdout, "Response from `ComputeApi.GetComputeRckUnitList`: %v\n", apiResponse)
}
```
