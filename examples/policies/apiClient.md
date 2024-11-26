#### Example to create policy
```
package policy

import (
	"context"
	"fmt"
	"log"
	"os"

	intersight "github.com/CiscoDevNet/intersight-go"
)

func SetInputs(config *Config) (context.Context, error) {
	ctx := context.Background()
	httpSigningInfo := new(intersight.HttpSignatureAuth)
	httpSigningInfo.KeyId = config.ApiKey
	httpSigningInfo.PrivateKeyPath = config.SecretKey
	httpSigningInfo.SigningScheme = intersight.HttpSigningSchemeRsaSha256
	httpSigningInfo.SigningAlgorithm = intersight.HttpSigningAlgorithmRsaPKCS1v15
	httpSigningInfo.SignedHeaders = []string{intersight.HttpSignatureParameterRequestTarget,
		// 		intersight.HttpSignatureParameterCreated,
		// 		intersight.HttpSignatureParameterExpires,
		"Host",
		"Date",
		"Digest",
	}
	if _, err := os.Stat(config.SecretKey); err != nil {
		err = httpSigningInfo.SetPrivateKey(config.SecretKey)
		if err != nil {
			return nil, err
		}
	} else {
		httpSigningInfo.PrivateKeyPath = config.SecretKey
	}

	ctx, err := httpSigningInfo.ContextWithValue(ctx)
	if err != nil {
		fmt.Println(err)
		log.Fatal("Error creating authentication context")
	}

	return ctx, nil
}

func getApiClient(config *Config) *Config {

	ctx, err := SetInputs(config)
	if err != nil {
		log.Printf("Error: ", err)
		log.Fatal("Error in authentication context")
	}
	config.ctx = ctx
	cfg := intersight.NewConfiguration()
	cfg.Host = config.Endpoint
	cfg.Debug = true
	apiClient := intersight.NewAPIClient(cfg)
	config.ApiClient = apiClient
	return config
}
```
