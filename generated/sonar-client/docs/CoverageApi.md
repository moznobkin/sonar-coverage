# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeasuresComponentGet**](CoverageApi.md#MeasuresComponentGet) | **Get** /measures/component | returns component metrics

# **MeasuresComponentGet**
> ProjectMetrics MeasuresComponentGet(ctx, component, optional)
returns component metrics

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **component** | **string**| component id | 
 **optional** | ***CoverageApiMeasuresComponentGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CoverageApiMeasuresComponentGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **branch** | **optional.String**|  | [default to main]
 **metricKeys** | **optional.String**|  | [default to ncloc,complexity,violations,coverage,lines_to_cover]

### Return type

[**ProjectMetrics**](ProjectMetrics.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

