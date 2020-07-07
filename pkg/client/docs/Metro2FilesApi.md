# \Metro2FilesApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Convert**](Metro2FilesApi.md#Convert) | **Post** /convert | Convert metro2 file
[**Health**](Metro2FilesApi.md#Health) | **Get** /health | health metro2 service
[**Print**](Metro2FilesApi.md#Print) | **Post** /print | Print metro2 file with specific format
[**Validator**](Metro2FilesApi.md#Validator) | **Post** /validator | Validate metro2 file



## Convert

> *os.File Convert(ctx, optional)

Convert metro2 file

Convert from original metro2 file to new metro2 file

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConvertOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ConvertOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **optional.String**| print metro2 file type | [default to json]
 **generate** | **optional.Bool**| generate new trailer record | [default to false]
 **file** | **optional.Interface of *os.File****optional.*os.File**| metro2 file to upload | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/octet-stream, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Health

> string Health(ctx, )

health metro2 service

Check the metro2 service to check if running

### Required Parameters

This endpoint does not need any parameter.

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

> string Print(ctx, optional)

Print metro2 file with specific format

Print metro2 file with requested file format.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PrintOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PrintOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **optional.String**| print metro2 file type | [default to json]
 **file** | **optional.Interface of *os.File****optional.*os.File**| metro2 file to upload | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Validator

> string Validator(ctx, optional)

Validate metro2 file

Validation of metro2 file.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ValidatorOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidatorOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | **optional.Interface of *os.File****optional.*os.File**| metro2 file to upload | 

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

