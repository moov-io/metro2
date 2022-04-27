# \Metro2FilesApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Convert**](Metro2FilesApi.md#Convert) | **Post** /convert | Convert metro2 file
[**Health**](Metro2FilesApi.md#Health) | **Get** /health | health metro2 service
[**Print**](Metro2FilesApi.md#Print) | **Post** /print | Print metro2 file with specific format
[**Validator**](Metro2FilesApi.md#Validator) | **Post** /validator | Validate metro2 file



## Convert

> *os.File Convert(ctx).Format(format).Format2(format2).Type_(type_).Generate(generate).Newline(newline).File(file).Execute()

Convert metro2 file



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    format := "format_example" // string | print metro2 file type (optional) (default to "json")
    format2 := "format_example" // string | format of metro file (optional) (default to "json")
    type_ := "type__example" // string | metro file type (optional)
    generate := true // bool | generate new trailer record (optional) (default to false)
    newline := true // bool | has new line (optional)
    file := os.NewFile(1234, "some_file") // *os.File | metro2 file to upload (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Metro2FilesApi.Convert(context.Background()).Format(format).Format2(format2).Type_(type_).Generate(generate).Newline(newline).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Metro2FilesApi.Convert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Convert`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `Metro2FilesApi.Convert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConvertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string** | print metro2 file type | [default to &quot;json&quot;]
 **format2** | **string** | format of metro file | [default to &quot;json&quot;]
 **type_** | **string** | metro file type | 
 **generate** | **bool** | generate new trailer record | [default to false]
 **newline** | **bool** | has new line | 
 **file** | ***os.File** | metro2 file to upload | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data, application/json
- **Accept**: application/octet-stream, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Health

> string Health(ctx).Execute()

health metro2 service



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Metro2FilesApi.Health(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Metro2FilesApi.Health``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Health`: string
    fmt.Fprintf(os.Stdout, "Response from `Metro2FilesApi.Health`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHealthRequest struct via the builder pattern


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Print

> string Print(ctx).Format(format).Format2(format2).File(file).Execute()

Print metro2 file with specific format



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    format := "format_example" // string | print metro2 file type (optional) (default to "json")
    format2 := "format_example" // string | print metro2 file type (optional) (default to "json")
    file := os.NewFile(1234, "some_file") // *os.File | metro2 file to upload (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Metro2FilesApi.Print(context.Background()).Format(format).Format2(format2).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Metro2FilesApi.Print``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Print`: string
    fmt.Fprintf(os.Stdout, "Response from `Metro2FilesApi.Print`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string** | print metro2 file type | [default to &quot;json&quot;]
 **format2** | **string** | print metro2 file type | [default to &quot;json&quot;]
 **file** | ***os.File** | metro2 file to upload | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data, application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Validator

> string Validator(ctx).File(file).Execute()

Validate metro2 file



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    file := os.NewFile(1234, "some_file") // *os.File | metro2 file to upload (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Metro2FilesApi.Validator(context.Background()).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Metro2FilesApi.Validator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Validator`: string
    fmt.Fprintf(os.Stdout, "Response from `Metro2FilesApi.Validator`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | metro2 file to upload | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

