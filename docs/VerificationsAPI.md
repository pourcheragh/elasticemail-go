# \VerificationsAPI

All URIs are relative to *https://api.elasticemail.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VerificationsByEmailDelete**](VerificationsAPI.md#VerificationsByEmailDelete) | **Delete** /verifications/{email} | Delete Email Verification Result
[**VerificationsByEmailGet**](VerificationsAPI.md#VerificationsByEmailGet) | **Get** /verifications/{email} | Get Email Verification Result
[**VerificationsByEmailPost**](VerificationsAPI.md#VerificationsByEmailPost) | **Post** /verifications/{email} | Verify Email
[**VerificationsFilesByIdDelete**](VerificationsAPI.md#VerificationsFilesByIdDelete) | **Delete** /verifications/files/{id} | Delete File Verification Result
[**VerificationsFilesByIdResultDownloadGet**](VerificationsAPI.md#VerificationsFilesByIdResultDownloadGet) | **Get** /verifications/files/{id}/result/download | Download File Verification Result
[**VerificationsFilesByIdResultGet**](VerificationsAPI.md#VerificationsFilesByIdResultGet) | **Get** /verifications/files/{id}/result | Get Detailed File Verification Result
[**VerificationsFilesByIdVerificationPost**](VerificationsAPI.md#VerificationsFilesByIdVerificationPost) | **Post** /verifications/files/{id}/verification | Start verification
[**VerificationsFilesPost**](VerificationsAPI.md#VerificationsFilesPost) | **Post** /verifications/files | Upload File with Emails
[**VerificationsFilesResultGet**](VerificationsAPI.md#VerificationsFilesResultGet) | **Get** /verifications/files/result | Get Files Verification Results
[**VerificationsGet**](VerificationsAPI.md#VerificationsGet) | **Get** /verifications | Get Emails Verification Results



## VerificationsByEmailDelete

> VerificationsByEmailDelete(ctx, email).Execute()

Delete Email Verification Result



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pourcheragh/elasticemail-go"
)

func main() {
	email := "email_example" // string | Email address to verification

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VerificationsAPI.VerificationsByEmailDelete(context.Background(), email).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerificationsAPI.VerificationsByEmailDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** | Email address to verification | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerificationsByEmailDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerificationsByEmailGet

> EmailValidationResult VerificationsByEmailGet(ctx, email).Execute()

Get Email Verification Result



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pourcheragh/elasticemail-go"
)

func main() {
	email := "email_example" // string | Email address to view verification result of

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerificationsAPI.VerificationsByEmailGet(context.Background(), email).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerificationsAPI.VerificationsByEmailGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerificationsByEmailGet`: EmailValidationResult
	fmt.Fprintf(os.Stdout, "Response from `VerificationsAPI.VerificationsByEmailGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** | Email address to view verification result of | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerificationsByEmailGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EmailValidationResult**](EmailValidationResult.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerificationsByEmailPost

> EmailValidationResult VerificationsByEmailPost(ctx, email).Execute()

Verify Email



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pourcheragh/elasticemail-go"
)

func main() {
	email := "email_example" // string | Email address to verify

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerificationsAPI.VerificationsByEmailPost(context.Background(), email).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerificationsAPI.VerificationsByEmailPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerificationsByEmailPost`: EmailValidationResult
	fmt.Fprintf(os.Stdout, "Response from `VerificationsAPI.VerificationsByEmailPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** | Email address to verify | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerificationsByEmailPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EmailValidationResult**](EmailValidationResult.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerificationsFilesByIdDelete

> VerificationsFilesByIdDelete(ctx, id).Execute()

Delete File Verification Result



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pourcheragh/elasticemail-go"
)

func main() {
	id := "E33EBA7A-C20D-4D3D-8F2F-5EEF42F58E6F" // string | ID of the exported file

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VerificationsAPI.VerificationsFilesByIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerificationsAPI.VerificationsFilesByIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the exported file | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerificationsFilesByIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerificationsFilesByIdResultDownloadGet

> *os.File VerificationsFilesByIdResultDownloadGet(ctx, id).Execute()

Download File Verification Result



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pourcheragh/elasticemail-go"
)

func main() {
	id := "id_example" // string | Verification ID to download

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerificationsAPI.VerificationsFilesByIdResultDownloadGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerificationsAPI.VerificationsFilesByIdResultDownloadGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerificationsFilesByIdResultDownloadGet`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `VerificationsAPI.VerificationsFilesByIdResultDownloadGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Verification ID to download | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerificationsFilesByIdResultDownloadGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerificationsFilesByIdResultGet

> VerificationFileResultDetails VerificationsFilesByIdResultGet(ctx, id).Limit(limit).Offset(offset).Execute()

Get Detailed File Verification Result



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pourcheragh/elasticemail-go"
)

func main() {
	id := "id_example" // string | ID of the Verification to display status of
	limit := int32(56) // int32 | Maximum number of returned email verification results (optional)
	offset := int32(56) // int32 | How many result items should be returned ahead (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerificationsAPI.VerificationsFilesByIdResultGet(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerificationsAPI.VerificationsFilesByIdResultGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerificationsFilesByIdResultGet`: VerificationFileResultDetails
	fmt.Fprintf(os.Stdout, "Response from `VerificationsAPI.VerificationsFilesByIdResultGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Verification to display status of | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerificationsFilesByIdResultGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Maximum number of returned email verification results | 
 **offset** | **int32** | How many result items should be returned ahead | 

### Return type

[**VerificationFileResultDetails**](VerificationFileResultDetails.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerificationsFilesByIdVerificationPost

> VerificationsFilesByIdVerificationPost(ctx, id).Execute()

Start verification



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pourcheragh/elasticemail-go"
)

func main() {
	id := "id_example" // string | File ID to start verification

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VerificationsAPI.VerificationsFilesByIdVerificationPost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerificationsAPI.VerificationsFilesByIdVerificationPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | File ID to start verification | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerificationsFilesByIdVerificationPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerificationsFilesPost

> VerificationFileResult VerificationsFilesPost(ctx).File(file).Execute()

Upload File with Emails



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pourcheragh/elasticemail-go"
)

func main() {
	file := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerificationsAPI.VerificationsFilesPost(context.Background()).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerificationsAPI.VerificationsFilesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerificationsFilesPost`: VerificationFileResult
	fmt.Fprintf(os.Stdout, "Response from `VerificationsAPI.VerificationsFilesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerificationsFilesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 

### Return type

[**VerificationFileResult**](VerificationFileResult.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerificationsFilesResultGet

> []VerificationFileResult VerificationsFilesResultGet(ctx).Execute()

Get Files Verification Results



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pourcheragh/elasticemail-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerificationsAPI.VerificationsFilesResultGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerificationsAPI.VerificationsFilesResultGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerificationsFilesResultGet`: []VerificationFileResult
	fmt.Fprintf(os.Stdout, "Response from `VerificationsAPI.VerificationsFilesResultGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVerificationsFilesResultGetRequest struct via the builder pattern


### Return type

[**[]VerificationFileResult**](VerificationFileResult.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerificationsGet

> []EmailValidationResult VerificationsGet(ctx).Limit(limit).Offset(offset).Execute()

Get Emails Verification Results



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pourcheragh/elasticemail-go"
)

func main() {
	limit := int32(100) // int32 | Maximum number of returned items. (optional)
	offset := int32(20) // int32 | How many items should be returned ahead. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerificationsAPI.VerificationsGet(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerificationsAPI.VerificationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerificationsGet`: []EmailValidationResult
	fmt.Fprintf(os.Stdout, "Response from `VerificationsAPI.VerificationsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerificationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of returned items. | 
 **offset** | **int32** | How many items should be returned ahead. | 

### Return type

[**[]EmailValidationResult**](EmailValidationResult.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

