/*
Elastic Email REST API

Testing FilesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package elasticemail

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/pourcheragh/elasticemail-go"
)

func Test_elasticemail_FilesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FilesAPIService FilesByNameDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		httpRes, err := apiClient.FilesAPI.FilesByNameDelete(context.Background(), name).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FilesAPIService FilesByNameGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.FilesAPI.FilesByNameGet(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FilesAPIService FilesByNameInfoGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.FilesAPI.FilesByNameInfoGet(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FilesAPIService FilesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FilesAPI.FilesGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FilesAPIService FilesPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FilesAPI.FilesPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}