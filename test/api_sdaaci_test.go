/*
Cisco Intersight

Testing SdaaciApiService

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

func Test_intersight_SdaaciApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SdaaciApiService CreateSdaaciConnection", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SdaaciApi.CreateSdaaciConnection(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SdaaciApiService CreateSdaaciConnectionDetail", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SdaaciApi.CreateSdaaciConnectionDetail(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SdaaciApiService DeleteSdaaciConnection", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var moid string

		httpRes, err := apiClient.SdaaciApi.DeleteSdaaciConnection(context.Background(), moid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SdaaciApiService DeleteSdaaciConnectionDetail", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var moid string

		httpRes, err := apiClient.SdaaciApi.DeleteSdaaciConnectionDetail(context.Background(), moid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SdaaciApiService GetSdaaciConnectionByMoid", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var moid string

		resp, httpRes, err := apiClient.SdaaciApi.GetSdaaciConnectionByMoid(context.Background(), moid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SdaaciApiService GetSdaaciConnectionDetailByMoid", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var moid string

		resp, httpRes, err := apiClient.SdaaciApi.GetSdaaciConnectionDetailByMoid(context.Background(), moid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SdaaciApiService GetSdaaciConnectionDetailList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SdaaciApi.GetSdaaciConnectionDetailList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SdaaciApiService GetSdaaciConnectionList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SdaaciApi.GetSdaaciConnectionList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SdaaciApiService PatchSdaaciConnection", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var moid string

		resp, httpRes, err := apiClient.SdaaciApi.PatchSdaaciConnection(context.Background(), moid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SdaaciApiService PatchSdaaciConnectionDetail", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var moid string

		resp, httpRes, err := apiClient.SdaaciApi.PatchSdaaciConnectionDetail(context.Background(), moid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SdaaciApiService UpdateSdaaciConnection", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var moid string

		resp, httpRes, err := apiClient.SdaaciApi.UpdateSdaaciConnection(context.Background(), moid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SdaaciApiService UpdateSdaaciConnectionDetail", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var moid string

		resp, httpRes, err := apiClient.SdaaciApi.UpdateSdaaciConnectionDetail(context.Background(), moid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}