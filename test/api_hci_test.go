/*
Cisco Intersight

Testing HciApiService

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

func Test_intersight_HciApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test HciApiService GetHciAlarmByMoid", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var moid string

		resp, httpRes, err := apiClient.HciApi.GetHciAlarmByMoid(context.Background(), moid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HciApiService GetHciAlarmList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.HciApi.GetHciAlarmList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HciApiService GetHciClusterByMoid", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var moid string

		resp, httpRes, err := apiClient.HciApi.GetHciClusterByMoid(context.Background(), moid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HciApiService GetHciClusterList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.HciApi.GetHciClusterList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HciApiService GetHciComplianceByMoid", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var moid string

		resp, httpRes, err := apiClient.HciApi.GetHciComplianceByMoid(context.Background(), moid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HciApiService GetHciComplianceList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.HciApi.GetHciComplianceList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HciApiService GetHciDiskByMoid", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var moid string

		resp, httpRes, err := apiClient.HciApi.GetHciDiskByMoid(context.Background(), moid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HciApiService GetHciDiskList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.HciApi.GetHciDiskList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HciApiService GetHciDomainManagerByMoid", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var moid string

		resp, httpRes, err := apiClient.HciApi.GetHciDomainManagerByMoid(context.Background(), moid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HciApiService GetHciDomainManagerList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.HciApi.GetHciDomainManagerList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HciApiService GetHciEntitlementByMoid", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var moid string

		resp, httpRes, err := apiClient.HciApi.GetHciEntitlementByMoid(context.Background(), moid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HciApiService GetHciEntitlementList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.HciApi.GetHciEntitlementList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HciApiService GetHciGpuByMoid", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var moid string

		resp, httpRes, err := apiClient.HciApi.GetHciGpuByMoid(context.Background(), moid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HciApiService GetHciGpuList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.HciApi.GetHciGpuList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HciApiService GetHciLicenseByMoid", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var moid string

		resp, httpRes, err := apiClient.HciApi.GetHciLicenseByMoid(context.Background(), moid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HciApiService GetHciLicenseList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.HciApi.GetHciLicenseList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HciApiService GetHciNodeByMoid", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var moid string

		resp, httpRes, err := apiClient.HciApi.GetHciNodeByMoid(context.Background(), moid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HciApiService GetHciNodeList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.HciApi.GetHciNodeList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HciApiService GetHciViolationByMoid", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var moid string

		resp, httpRes, err := apiClient.HciApi.GetHciViolationByMoid(context.Background(), moid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HciApiService GetHciViolationList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.HciApi.GetHciViolationList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HciApiService PatchHciCluster", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var moid string

		resp, httpRes, err := apiClient.HciApi.PatchHciCluster(context.Background(), moid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HciApiService PatchHciNode", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var moid string

		resp, httpRes, err := apiClient.HciApi.PatchHciNode(context.Background(), moid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HciApiService UpdateHciCluster", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var moid string

		resp, httpRes, err := apiClient.HciApi.UpdateHciCluster(context.Background(), moid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HciApiService UpdateHciNode", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var moid string

		resp, httpRes, err := apiClient.HciApi.UpdateHciNode(context.Background(), moid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
