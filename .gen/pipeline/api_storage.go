/*
 * Pipeline API
 *
 * Pipeline is a feature rich application platform, built for containers on top of Kubernetes to automate the DevOps experience, continuous application development and the lifecycle of deployments. 
 *
 * API version: latest
 * Contact: info@banzaicloud.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipeline

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ _context.Context
)

// StorageApiService StorageApi service
type StorageApiService service

/*
CreateObjectStoreBucket Creates a new object store bucket with the given params
Creates a new object store bucket on the Cloud provider referenced by the given secret. The credentials for creating the bucket is taken from the provided secret.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgId Organization identifier
 * @param createObjectStoreBucketRequest
@return CreateObjectStoreBucketResponse
*/
func (a *StorageApiService) CreateObjectStoreBucket(ctx _context.Context, orgId int32, createObjectStoreBucketRequest CreateObjectStoreBucketRequest) (CreateObjectStoreBucketResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CreateObjectStoreBucketResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v1/orgs/{orgId}/buckets"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", _neturl.QueryEscape(parameterToString(orgId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = &createObjectStoreBucketRequest
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v CommonError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// DeleteObjectStoreBucketOpts Optional parameters for the method 'DeleteObjectStoreBucket'
type DeleteObjectStoreBucketOpts struct {
    Force optional.Bool
    ResourceGroup optional.String
    StorageAccount optional.String
    Location optional.String
}

/*
DeleteObjectStoreBucket Deletes the object store bucket with the given name
Deletes the object store bucket identified by the given name. The credentials for deleting the bucket is taken from the provided secret.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgId Organization identifier
 * @param name Bucket identification
 * @param secretId Secret identification
 * @param cloudType Identifies the cloud provider
 * @param optional nil or *DeleteObjectStoreBucketOpts - Optional Parameters:
 * @param "Force" (optional.Bool) -  Is the operation forced
 * @param "ResourceGroup" (optional.String) -  Azure resource group the storage account that holds the bucket (storage container) to be deleted
 * @param "StorageAccount" (optional.String) -  Azure storage account to delete the bucket (storage container) from
 * @param "Location" (optional.String) -  The region to delete the bucket from. Required on Amazon, Oracle and Alibaba cloud providers.
*/
func (a *StorageApiService) DeleteObjectStoreBucket(ctx _context.Context, orgId int32, name string, secretId string, cloudType string, localVarOptionals *DeleteObjectStoreBucketOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v1/orgs/{orgId}/buckets/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", _neturl.QueryEscape(parameterToString(orgId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", _neturl.QueryEscape(parameterToString(name, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	localVarQueryParams.Add("cloudType", parameterToString(cloudType, ""))
	if localVarOptionals != nil && localVarOptionals.Force.IsSet() {
		localVarQueryParams.Add("force", parameterToString(localVarOptionals.Force.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ResourceGroup.IsSet() {
		localVarQueryParams.Add("resourceGroup", parameterToString(localVarOptionals.ResourceGroup.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StorageAccount.IsSet() {
		localVarQueryParams.Add("storageAccount", parameterToString(localVarOptionals.StorageAccount.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Location.IsSet() {
		localVarQueryParams.Add("location", parameterToString(localVarOptionals.Location.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["secretId"] = parameterToString(secretId, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v CommonError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// GetBucketOpts Optional parameters for the method 'GetBucket'
type GetBucketOpts struct {
    SecretId optional.String
    SecretName optional.String
    ResourceGroup optional.String
    StorageAccount optional.String
    Location optional.String
}

/*
GetBucket Get object store bucket details
Retrieves the details of the object store bucket given its name
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgId Organization identifier
 * @param name Bucket identification
 * @param cloudType Identifies the cloud provider
 * @param optional nil or *GetBucketOpts - Optional Parameters:
 * @param "SecretId" (optional.String) -  Secret identification
 * @param "SecretName" (optional.String) -  Secret identification by name
 * @param "ResourceGroup" (optional.String) -  Azure resource group to lookup the bucket(storage container) under. Required only on Azure cloud provider.
 * @param "StorageAccount" (optional.String) -  Azure storage account to lookup the bucket(storage container) under. Required only on Azure cloud provider.
 * @param "Location" (optional.String) -  The region to lookup the bucket under. Required on Amazon, Oracle and Alibaba cloud providers.
@return BucketInfo
*/
func (a *StorageApiService) GetBucket(ctx _context.Context, orgId int32, name string, cloudType string, localVarOptionals *GetBucketOpts) (BucketInfo, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  BucketInfo
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v1/orgs/{orgId}/buckets/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", _neturl.QueryEscape(parameterToString(orgId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", _neturl.QueryEscape(parameterToString(name, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	localVarQueryParams.Add("cloudType", parameterToString(cloudType, ""))
	if localVarOptionals != nil && localVarOptionals.ResourceGroup.IsSet() {
		localVarQueryParams.Add("resourceGroup", parameterToString(localVarOptionals.ResourceGroup.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StorageAccount.IsSet() {
		localVarQueryParams.Add("storageAccount", parameterToString(localVarOptionals.StorageAccount.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Location.IsSet() {
		localVarQueryParams.Add("location", parameterToString(localVarOptionals.Location.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.SecretId.IsSet() {
		localVarHeaderParams["secretId"] = parameterToString(localVarOptionals.SecretId.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.SecretName.IsSet() {
		localVarHeaderParams["secretName"] = parameterToString(localVarOptionals.SecretName.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v CommonError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// GetObjectStoreBucketStatusOpts Optional parameters for the method 'GetObjectStoreBucketStatus'
type GetObjectStoreBucketStatusOpts struct {
    SecretId optional.String
    SecretName optional.String
    ResourceGroup optional.String
    StorageAccount optional.String
    Location optional.String
}

/*
GetObjectStoreBucketStatus Get object store bucket status
Retrieves the status of the object store bucket given its name
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgId Organization identifier
 * @param name Bucket identification
 * @param cloudType Identifies the cloud provider
 * @param optional nil or *GetObjectStoreBucketStatusOpts - Optional Parameters:
 * @param "SecretId" (optional.String) -  Secret identification
 * @param "SecretName" (optional.String) -  Secret identification by name
 * @param "ResourceGroup" (optional.String) -  Azure resource group to lookup the bucket(storage container) under. Required only on Azure cloud provider.
 * @param "StorageAccount" (optional.String) -  Azure storage account to lookup the bucket(storage container) under. Required only on Azure cloud provider.
 * @param "Location" (optional.String) -  The region to lookup the bucket under. Required on Amazon, Oracle and Alibaba cloud providers.
*/
func (a *StorageApiService) GetObjectStoreBucketStatus(ctx _context.Context, orgId int32, name string, cloudType string, localVarOptionals *GetObjectStoreBucketStatusOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodHead
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v1/orgs/{orgId}/buckets/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", _neturl.QueryEscape(parameterToString(orgId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", _neturl.QueryEscape(parameterToString(name, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	localVarQueryParams.Add("cloudType", parameterToString(cloudType, ""))
	if localVarOptionals != nil && localVarOptionals.ResourceGroup.IsSet() {
		localVarQueryParams.Add("resourceGroup", parameterToString(localVarOptionals.ResourceGroup.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StorageAccount.IsSet() {
		localVarQueryParams.Add("storageAccount", parameterToString(localVarOptionals.StorageAccount.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Location.IsSet() {
		localVarQueryParams.Add("location", parameterToString(localVarOptionals.Location.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.SecretId.IsSet() {
		localVarHeaderParams["secretId"] = parameterToString(localVarOptionals.SecretId.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.SecretName.IsSet() {
		localVarHeaderParams["secretName"] = parameterToString(localVarOptionals.SecretName.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v CommonError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// ListObjectStoreBucketsOpts Optional parameters for the method 'ListObjectStoreBuckets'
type ListObjectStoreBucketsOpts struct {
    SecretId optional.String
    CloudType optional.String
    Include optional.String
}

/*
ListObjectStoreBuckets List object storage buckets
List object store buckets accessible by the credentials referenced by the given secret. If no credentials provided all managed buckets are returned for all cloud types.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgId Organization identifier
 * @param optional nil or *ListObjectStoreBucketsOpts - Optional Parameters:
 * @param "SecretId" (optional.String) -  Secret identification. If not provided only the managed buckets (those created via pipeline) are listed
 * @param "CloudType" (optional.String) -  Identifies the cloud provider - mandatory if secretId header is provided
 * @param "Include" (optional.String) -  Signals whether the secret name is to be returned
@return []BucketInfo
*/
func (a *StorageApiService) ListObjectStoreBuckets(ctx _context.Context, orgId int32, localVarOptionals *ListObjectStoreBucketsOpts) ([]BucketInfo, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []BucketInfo
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v1/orgs/{orgId}/buckets"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", _neturl.QueryEscape(parameterToString(orgId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.CloudType.IsSet() {
		localVarQueryParams.Add("cloudType", parameterToString(localVarOptionals.CloudType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Include.IsSet() {
		localVarQueryParams.Add("include", parameterToString(localVarOptionals.Include.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.SecretId.IsSet() {
		localVarHeaderParams["secretId"] = parameterToString(localVarOptionals.SecretId.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v CommonError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
