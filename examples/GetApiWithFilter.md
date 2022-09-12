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
	// 	expand := "Biosunits"
	//	apply := "groupby((Model),aggregate(AvailableMemory with min as MinAvailableMemory))"
	count := false
	inlinecount := "allpages"
	// 	at := "VersionType eq 'Configured'"
	// 	tags := "tags_example"

	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	apiResponse, r, err := apiClient.ComputeApi.GetComputeRackUnitList(ctx).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Count(count).Inlinecount(inlinecount).Execute()
	if err != nil {
		// 		fmt.Fprintf(os.Stderr, "Error when calling `ComputeApi.GetComputeRckUnitList``: %v\n", err)
		// 		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		log.Printf("Error when calling `ComputeApi.GetComputeRckUnitList``: %v\n", err)
		log.Printf("Full HTTP response: %v\n", r)
		return
	}

	fmt.Fprintf(os.Stdout, "Response from `ComputeApi.GetComputeRckUnitList`: %v\n", apiResponse)
}
```
