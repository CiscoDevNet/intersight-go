/*
Cisco Intersight

Testing InventoryApiService

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

func Test_intersight_InventoryApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test InventoryApiService CreateInventoryRequest", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.InventoryApi.CreateInventoryRequest(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiService GetInventoryDeviceInfoByMoid", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var moid string

		resp, httpRes, err := apiClient.InventoryApi.GetInventoryDeviceInfoByMoid(context.Background(), moid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiService GetInventoryDeviceInfoList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.InventoryApi.GetInventoryDeviceInfoList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiService GetInventoryDnMoBindingByMoid", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var moid string

		resp, httpRes, err := apiClient.InventoryApi.GetInventoryDnMoBindingByMoid(context.Background(), moid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiService GetInventoryDnMoBindingList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.InventoryApi.GetInventoryDnMoBindingList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiService GetInventoryGenericInventoryByMoid", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var moid string

		resp, httpRes, err := apiClient.InventoryApi.GetInventoryGenericInventoryByMoid(context.Background(), moid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiService GetInventoryGenericInventoryHolderByMoid", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var moid string

		resp, httpRes, err := apiClient.InventoryApi.GetInventoryGenericInventoryHolderByMoid(context.Background(), moid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiService GetInventoryGenericInventoryHolderList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.InventoryApi.GetInventoryGenericInventoryHolderList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiService GetInventoryGenericInventoryList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.InventoryApi.GetInventoryGenericInventoryList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiService PatchInventoryGenericInventory", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var moid string

		resp, httpRes, err := apiClient.InventoryApi.PatchInventoryGenericInventory(context.Background(), moid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiService PatchInventoryGenericInventoryHolder", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var moid string

		resp, httpRes, err := apiClient.InventoryApi.PatchInventoryGenericInventoryHolder(context.Background(), moid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiService UpdateInventoryGenericInventory", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var moid string

		resp, httpRes, err := apiClient.InventoryApi.UpdateInventoryGenericInventory(context.Background(), moid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InventoryApiService UpdateInventoryGenericInventoryHolder", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var moid string

		resp, httpRes, err := apiClient.InventoryApi.UpdateInventoryGenericInventoryHolder(context.Background(), moid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
