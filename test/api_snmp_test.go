/*
Cisco Intersight

Testing SnmpApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package intersight

import (
	"context"
	"testing"

	openapiclient "github.com/CiscoDevNet/intersight-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_intersight_SnmpApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SnmpApiService CreateSnmpPolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SnmpApi.CreateSnmpPolicy(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SnmpApiService DeleteSnmpPolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var moid string

		httpRes, err := apiClient.SnmpApi.DeleteSnmpPolicy(context.Background(), moid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SnmpApiService GetSnmpPolicyByMoid", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var moid string

		resp, httpRes, err := apiClient.SnmpApi.GetSnmpPolicyByMoid(context.Background(), moid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SnmpApiService GetSnmpPolicyInventoryByMoid", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var moid string

		resp, httpRes, err := apiClient.SnmpApi.GetSnmpPolicyInventoryByMoid(context.Background(), moid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SnmpApiService GetSnmpPolicyInventoryList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SnmpApi.GetSnmpPolicyInventoryList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SnmpApiService GetSnmpPolicyList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SnmpApi.GetSnmpPolicyList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SnmpApiService PatchSnmpPolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var moid string

		resp, httpRes, err := apiClient.SnmpApi.PatchSnmpPolicy(context.Background(), moid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SnmpApiService UpdateSnmpPolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var moid string

		resp, httpRes, err := apiClient.SnmpApi.UpdateSnmpPolicy(context.Background(), moid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}