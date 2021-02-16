/*
 * Kafka HTTP API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package kafkarestv3

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

type ClusterLinkingApi interface {

    /*
     * ClustersClusterIdLinksLinkNameMirrorsDestinationTopicNameGet Describe the mirror topic
     *
     * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
     * @param clusterId The Kafka cluster ID.
     * @param linkName The link name
     * @param destinationTopicName Cluster Linking destination topic name
     * @return ListMirrorTopicsResponseData
     */
    ClustersClusterIdLinksLinkNameMirrorsDestinationTopicNameGet(ctx _context.Context, clusterId string, linkName string, destinationTopicName string) (ListMirrorTopicsResponseData, *_nethttp.Response, error)

    /*
     * ClustersClusterIdLinksLinkNameMirrorsFailoverPost Failover the mirror topics
     *
     * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
     * @param clusterId The Kafka cluster ID.
     * @param linkName The link name
     * @param optional nil or *ClustersClusterIdLinksLinkNameMirrorsFailoverPostOpts - Optional Parameters:
     * @param "AlterMirrorsRequestData" (optional.Interface of AlterMirrorsRequestData) -  Name of the topics to apply the changes
     * @return AlterMirrorStatusResponseDataList
     */
    ClustersClusterIdLinksLinkNameMirrorsFailoverPost(ctx _context.Context, clusterId string, linkName string, localVarOptionals *ClustersClusterIdLinksLinkNameMirrorsFailoverPostOpts) (AlterMirrorStatusResponseDataList, *_nethttp.Response, error)

    /*
     * ClustersClusterIdLinksLinkNameMirrorsGet List mirror topics
     *
     * List all mirror topics under the link
     *
     * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
     * @param clusterId The Kafka cluster ID.
     * @param linkName The link name
     * @param optional nil or *ClustersClusterIdLinksLinkNameMirrorsGetOpts - Optional Parameters:
     * @param "MirrorStatus" (optional.Interface of MirrorTopicStatus) -  The status of the mirror topic. If not specified, all mirror topics will be returned.
     * @return ListMirrorTopicsResponseDataList
     */
    ClustersClusterIdLinksLinkNameMirrorsGet(ctx _context.Context, clusterId string, linkName string, localVarOptionals *ClustersClusterIdLinksLinkNameMirrorsGetOpts) (ListMirrorTopicsResponseDataList, *_nethttp.Response, error)

    /*
     * ClustersClusterIdLinksLinkNameMirrorsPausePost Pause the mirror topics
     *
     * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
     * @param clusterId The Kafka cluster ID.
     * @param linkName The link name
     * @param optional nil or *ClustersClusterIdLinksLinkNameMirrorsPausePostOpts - Optional Parameters:
     * @param "AlterMirrorsRequestData" (optional.Interface of AlterMirrorsRequestData) -  Name of the topics to apply the changes
     * @return AlterMirrorStatusResponseDataList
     */
    ClustersClusterIdLinksLinkNameMirrorsPausePost(ctx _context.Context, clusterId string, linkName string, localVarOptionals *ClustersClusterIdLinksLinkNameMirrorsPausePostOpts) (AlterMirrorStatusResponseDataList, *_nethttp.Response, error)

    /*
     * ClustersClusterIdLinksLinkNameMirrorsPromotePost Promote the mirror topics
     *
     * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
     * @param clusterId The Kafka cluster ID.
     * @param linkName The link name
     * @param optional nil or *ClustersClusterIdLinksLinkNameMirrorsPromotePostOpts - Optional Parameters:
     * @param "AlterMirrorsRequestData" (optional.Interface of AlterMirrorsRequestData) -  Name of the topics to apply the changes
     * @return AlterMirrorStatusResponseDataList
     */
    ClustersClusterIdLinksLinkNameMirrorsPromotePost(ctx _context.Context, clusterId string, linkName string, localVarOptionals *ClustersClusterIdLinksLinkNameMirrorsPromotePostOpts) (AlterMirrorStatusResponseDataList, *_nethttp.Response, error)

    /*
     * ClustersClusterIdLinksLinkNameMirrorsPut Create a mirror topic
     *
     * Create a topic in the destination cluster mirroring a topic in the source cluster
     *
     * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
     * @param clusterId The Kafka cluster ID.
     * @param linkName The link name
     * @param optional nil or *ClustersClusterIdLinksLinkNameMirrorsPutOpts - Optional Parameters:
     * @param "CreateMirrorTopicRequestData" (optional.Interface of CreateMirrorTopicRequestData) -  Name of the topics mirroring from and mirroring to
     */
    ClustersClusterIdLinksLinkNameMirrorsPut(ctx _context.Context, clusterId string, linkName string, localVarOptionals *ClustersClusterIdLinksLinkNameMirrorsPutOpts) (*_nethttp.Response, error)

    /*
     * ClustersClusterIdLinksLinkNameMirrorsResumePost Resume the mirror topics
     *
     * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
     * @param clusterId The Kafka cluster ID.
     * @param linkName The link name
     * @param optional nil or *ClustersClusterIdLinksLinkNameMirrorsResumePostOpts - Optional Parameters:
     * @param "AlterMirrorsRequestData" (optional.Interface of AlterMirrorsRequestData) -  Name of the topics to apply the changes
     * @return AlterMirrorStatusResponseDataList
     */
    ClustersClusterIdLinksLinkNameMirrorsResumePost(ctx _context.Context, clusterId string, linkName string, localVarOptionals *ClustersClusterIdLinksLinkNameMirrorsResumePostOpts) (AlterMirrorStatusResponseDataList, *_nethttp.Response, error)
}

// ClusterLinkingApiService ClusterLinkingApi service
type ClusterLinkingApiService service

/*
 * ClustersClusterIdLinksLinkNameMirrorsDestinationTopicNameGet Describe the mirror topic
 *
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterId The Kafka cluster ID.
 * @param linkName The link name
 * @param destinationTopicName Cluster Linking destination topic name
 * @return ListMirrorTopicsResponseData
 */
func (a *ClusterLinkingApiService) ClustersClusterIdLinksLinkNameMirrorsDestinationTopicNameGet(ctx _context.Context, clusterId string, linkName string, destinationTopicName string) (ListMirrorTopicsResponseData, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ListMirrorTopicsResponseData
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/clusters/{cluster_id}/links/{link_name}/mirrors/{destination_topic_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"cluster_id"+"}", _neturl.QueryEscape(parameterToString(clusterId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"link_name"+"}", _neturl.QueryEscape(parameterToString(linkName, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"destination_topic_name"+"}", _neturl.QueryEscape(parameterToString(destinationTopicName, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
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

// ClustersClusterIdLinksLinkNameMirrorsFailoverPostOpts Optional parameters for the method 'ClustersClusterIdLinksLinkNameMirrorsFailoverPost'
type ClustersClusterIdLinksLinkNameMirrorsFailoverPostOpts struct {
    AlterMirrorsRequestData optional.Interface
}

/*
 * ClustersClusterIdLinksLinkNameMirrorsFailoverPost Failover the mirror topics
 *
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterId The Kafka cluster ID.
 * @param linkName The link name
 * @param optional nil or *ClustersClusterIdLinksLinkNameMirrorsFailoverPostOpts - Optional Parameters:
 * @param "AlterMirrorsRequestData" (optional.Interface of AlterMirrorsRequestData) -  Name of the topics to apply the changes
 * @return AlterMirrorStatusResponseDataList
 */
func (a *ClusterLinkingApiService) ClustersClusterIdLinksLinkNameMirrorsFailoverPost(ctx _context.Context, clusterId string, linkName string, localVarOptionals *ClustersClusterIdLinksLinkNameMirrorsFailoverPostOpts) (AlterMirrorStatusResponseDataList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AlterMirrorStatusResponseDataList
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/clusters/{cluster_id}/links/{link_name}/mirrors/failover"
	localVarPath = strings.Replace(localVarPath, "{"+"cluster_id"+"}", _neturl.QueryEscape(parameterToString(clusterId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"link_name"+"}", _neturl.QueryEscape(parameterToString(linkName, "")) , -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.AlterMirrorsRequestData.IsSet() {
		localVarOptionalAlterMirrorsRequestData, localVarOptionalAlterMirrorsRequestDataok := localVarOptionals.AlterMirrorsRequestData.Value().(AlterMirrorsRequestData)
		if !localVarOptionalAlterMirrorsRequestDataok {
			return localVarReturnValue, nil, reportError("alterMirrorsRequestData should be AlterMirrorsRequestData")
		}
		localVarPostBody = &localVarOptionalAlterMirrorsRequestData
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

// ClustersClusterIdLinksLinkNameMirrorsGetOpts Optional parameters for the method 'ClustersClusterIdLinksLinkNameMirrorsGet'
type ClustersClusterIdLinksLinkNameMirrorsGetOpts struct {
    MirrorStatus optional.Interface
}

/*
 * ClustersClusterIdLinksLinkNameMirrorsGet List mirror topics
 *
 * List all mirror topics under the link
 *
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterId The Kafka cluster ID.
 * @param linkName The link name
 * @param optional nil or *ClustersClusterIdLinksLinkNameMirrorsGetOpts - Optional Parameters:
 * @param "MirrorStatus" (optional.Interface of MirrorTopicStatus) -  The status of the mirror topic. If not specified, all mirror topics will be returned.
 * @return ListMirrorTopicsResponseDataList
 */
func (a *ClusterLinkingApiService) ClustersClusterIdLinksLinkNameMirrorsGet(ctx _context.Context, clusterId string, linkName string, localVarOptionals *ClustersClusterIdLinksLinkNameMirrorsGetOpts) (ListMirrorTopicsResponseDataList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ListMirrorTopicsResponseDataList
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/clusters/{cluster_id}/links/{link_name}/mirrors"
	localVarPath = strings.Replace(localVarPath, "{"+"cluster_id"+"}", _neturl.QueryEscape(parameterToString(clusterId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"link_name"+"}", _neturl.QueryEscape(parameterToString(linkName, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.MirrorStatus.IsSet() {
		localVarQueryParams.Add("mirror_status", parameterToString(localVarOptionals.MirrorStatus.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
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

// ClustersClusterIdLinksLinkNameMirrorsPausePostOpts Optional parameters for the method 'ClustersClusterIdLinksLinkNameMirrorsPausePost'
type ClustersClusterIdLinksLinkNameMirrorsPausePostOpts struct {
    AlterMirrorsRequestData optional.Interface
}

/*
 * ClustersClusterIdLinksLinkNameMirrorsPausePost Pause the mirror topics
 *
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterId The Kafka cluster ID.
 * @param linkName The link name
 * @param optional nil or *ClustersClusterIdLinksLinkNameMirrorsPausePostOpts - Optional Parameters:
 * @param "AlterMirrorsRequestData" (optional.Interface of AlterMirrorsRequestData) -  Name of the topics to apply the changes
 * @return AlterMirrorStatusResponseDataList
 */
func (a *ClusterLinkingApiService) ClustersClusterIdLinksLinkNameMirrorsPausePost(ctx _context.Context, clusterId string, linkName string, localVarOptionals *ClustersClusterIdLinksLinkNameMirrorsPausePostOpts) (AlterMirrorStatusResponseDataList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AlterMirrorStatusResponseDataList
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/clusters/{cluster_id}/links/{link_name}/mirrors/pause"
	localVarPath = strings.Replace(localVarPath, "{"+"cluster_id"+"}", _neturl.QueryEscape(parameterToString(clusterId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"link_name"+"}", _neturl.QueryEscape(parameterToString(linkName, "")) , -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.AlterMirrorsRequestData.IsSet() {
		localVarOptionalAlterMirrorsRequestData, localVarOptionalAlterMirrorsRequestDataok := localVarOptionals.AlterMirrorsRequestData.Value().(AlterMirrorsRequestData)
		if !localVarOptionalAlterMirrorsRequestDataok {
			return localVarReturnValue, nil, reportError("alterMirrorsRequestData should be AlterMirrorsRequestData")
		}
		localVarPostBody = &localVarOptionalAlterMirrorsRequestData
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

// ClustersClusterIdLinksLinkNameMirrorsPromotePostOpts Optional parameters for the method 'ClustersClusterIdLinksLinkNameMirrorsPromotePost'
type ClustersClusterIdLinksLinkNameMirrorsPromotePostOpts struct {
    AlterMirrorsRequestData optional.Interface
}

/*
 * ClustersClusterIdLinksLinkNameMirrorsPromotePost Promote the mirror topics
 *
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterId The Kafka cluster ID.
 * @param linkName The link name
 * @param optional nil or *ClustersClusterIdLinksLinkNameMirrorsPromotePostOpts - Optional Parameters:
 * @param "AlterMirrorsRequestData" (optional.Interface of AlterMirrorsRequestData) -  Name of the topics to apply the changes
 * @return AlterMirrorStatusResponseDataList
 */
func (a *ClusterLinkingApiService) ClustersClusterIdLinksLinkNameMirrorsPromotePost(ctx _context.Context, clusterId string, linkName string, localVarOptionals *ClustersClusterIdLinksLinkNameMirrorsPromotePostOpts) (AlterMirrorStatusResponseDataList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AlterMirrorStatusResponseDataList
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/clusters/{cluster_id}/links/{link_name}/mirrors/promote"
	localVarPath = strings.Replace(localVarPath, "{"+"cluster_id"+"}", _neturl.QueryEscape(parameterToString(clusterId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"link_name"+"}", _neturl.QueryEscape(parameterToString(linkName, "")) , -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.AlterMirrorsRequestData.IsSet() {
		localVarOptionalAlterMirrorsRequestData, localVarOptionalAlterMirrorsRequestDataok := localVarOptionals.AlterMirrorsRequestData.Value().(AlterMirrorsRequestData)
		if !localVarOptionalAlterMirrorsRequestDataok {
			return localVarReturnValue, nil, reportError("alterMirrorsRequestData should be AlterMirrorsRequestData")
		}
		localVarPostBody = &localVarOptionalAlterMirrorsRequestData
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

// ClustersClusterIdLinksLinkNameMirrorsPutOpts Optional parameters for the method 'ClustersClusterIdLinksLinkNameMirrorsPut'
type ClustersClusterIdLinksLinkNameMirrorsPutOpts struct {
    CreateMirrorTopicRequestData optional.Interface
}

/*
 * ClustersClusterIdLinksLinkNameMirrorsPut Create a mirror topic
 *
 * Create a topic in the destination cluster mirroring a topic in the source cluster
 *
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterId The Kafka cluster ID.
 * @param linkName The link name
 * @param optional nil or *ClustersClusterIdLinksLinkNameMirrorsPutOpts - Optional Parameters:
 * @param "CreateMirrorTopicRequestData" (optional.Interface of CreateMirrorTopicRequestData) -  Name of the topics mirroring from and mirroring to
 */
func (a *ClusterLinkingApiService) ClustersClusterIdLinksLinkNameMirrorsPut(ctx _context.Context, clusterId string, linkName string, localVarOptionals *ClustersClusterIdLinksLinkNameMirrorsPutOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/clusters/{cluster_id}/links/{link_name}/mirrors"
	localVarPath = strings.Replace(localVarPath, "{"+"cluster_id"+"}", _neturl.QueryEscape(parameterToString(clusterId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"link_name"+"}", _neturl.QueryEscape(parameterToString(linkName, "")) , -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.CreateMirrorTopicRequestData.IsSet() {
		localVarOptionalCreateMirrorTopicRequestData, localVarOptionalCreateMirrorTopicRequestDataok := localVarOptionals.CreateMirrorTopicRequestData.Value().(CreateMirrorTopicRequestData)
		if !localVarOptionalCreateMirrorTopicRequestDataok {
			return nil, reportError("createMirrorTopicRequestData should be CreateMirrorTopicRequestData")
		}
		localVarPostBody = &localVarOptionalCreateMirrorTopicRequestData
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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// ClustersClusterIdLinksLinkNameMirrorsResumePostOpts Optional parameters for the method 'ClustersClusterIdLinksLinkNameMirrorsResumePost'
type ClustersClusterIdLinksLinkNameMirrorsResumePostOpts struct {
    AlterMirrorsRequestData optional.Interface
}

/*
 * ClustersClusterIdLinksLinkNameMirrorsResumePost Resume the mirror topics
 *
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterId The Kafka cluster ID.
 * @param linkName The link name
 * @param optional nil or *ClustersClusterIdLinksLinkNameMirrorsResumePostOpts - Optional Parameters:
 * @param "AlterMirrorsRequestData" (optional.Interface of AlterMirrorsRequestData) -  Name of the topics to apply the changes
 * @return AlterMirrorStatusResponseDataList
 */
func (a *ClusterLinkingApiService) ClustersClusterIdLinksLinkNameMirrorsResumePost(ctx _context.Context, clusterId string, linkName string, localVarOptionals *ClustersClusterIdLinksLinkNameMirrorsResumePostOpts) (AlterMirrorStatusResponseDataList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AlterMirrorStatusResponseDataList
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/clusters/{cluster_id}/links/{link_name}/mirrors/resume"
	localVarPath = strings.Replace(localVarPath, "{"+"cluster_id"+"}", _neturl.QueryEscape(parameterToString(clusterId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"link_name"+"}", _neturl.QueryEscape(parameterToString(linkName, "")) , -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.AlterMirrorsRequestData.IsSet() {
		localVarOptionalAlterMirrorsRequestData, localVarOptionalAlterMirrorsRequestDataok := localVarOptionals.AlterMirrorsRequestData.Value().(AlterMirrorsRequestData)
		if !localVarOptionalAlterMirrorsRequestDataok {
			return localVarReturnValue, nil, reportError("alterMirrorsRequestData should be AlterMirrorsRequestData")
		}
		localVarPostBody = &localVarOptionalAlterMirrorsRequestData
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
