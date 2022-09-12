### Example to ge the list of SMTP Policies

``` go
package example

import (
	"fmt"
	"log"
	"os"
)

func GetObjectList(config *Config) {

	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	apiResponse, r, err := apiClient.SmtpApi.GetSmtpPolicyList(ctx).Execute()
	if err != nil {
		// 		fmt.Fprintf(os.Stderr, "Error when calling `SmtpApi.GetSmtpPolicyList``: %v\n", err)
		// 		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		log.Printf("Error when calling `SmtpApi.GetSmtpPolicyList``: %v\n", err)
		log.Printf("Full HTTP response: %v\n", r)
		return
	}
	fmt.Fprintf(os.Stdout, "Response from `SmtpApi.GetSmtpPolicyList`: %v\n", apiResponse)
}
```
