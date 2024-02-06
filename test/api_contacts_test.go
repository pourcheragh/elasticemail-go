/*
Elastic Email REST API

Testing ContactsAPIService

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

func Test_elasticemail_ContactsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContactsAPIService ContactsByEmailDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var email string

		httpRes, err := apiClient.ContactsAPI.ContactsByEmailDelete(context.Background(), email).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContactsAPIService ContactsByEmailGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var email string

		resp, httpRes, err := apiClient.ContactsAPI.ContactsByEmailGet(context.Background(), email).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContactsAPIService ContactsByEmailPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var email string

		resp, httpRes, err := apiClient.ContactsAPI.ContactsByEmailPut(context.Background(), email).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContactsAPIService ContactsDeletePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ContactsAPI.ContactsDeletePost(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContactsAPIService ContactsExportByIdStatusGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ContactsAPI.ContactsExportByIdStatusGet(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContactsAPIService ContactsExportPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContactsAPI.ContactsExportPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContactsAPIService ContactsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContactsAPI.ContactsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContactsAPIService ContactsImportPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ContactsAPI.ContactsImportPost(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContactsAPIService ContactsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContactsAPI.ContactsPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}