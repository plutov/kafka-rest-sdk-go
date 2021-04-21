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
)

// Linger please
var (
	_ _context.Context
)

type BrokerTaskApi interface {

    /*
     * ClustersClusterIdBrokersBrokerIdTasksGet List Broker Tasks of a specific Broker
     *
     * Returns a list of all broker tasks for broker specified with &#x60;&#x60;broker_id&#x60;&#x60; in the cluster specified with &#x60;&#x60;cluster_id&#x60;&#x60;.
     *
     * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
     * @param clusterId The Kafka cluster ID.
     * @param brokerId The Kafka broker ID.
     * @return BrokerTaskDataList
     */
    ClustersClusterIdBrokersBrokerIdTasksGet(ctx _context.Context, clusterId string, brokerId int32) (BrokerTaskDataList, *_nethttp.Response, error)

    /*
     * ClustersClusterIdBrokersBrokerIdTasksTaskTypeGet Get single Broker Task.
     *
     * Returns a single Broker Task specified with &#x60;&#x60;task_type&#x60;&#x60; for broker specified with &#x60;&#x60;broker_id&#x60;&#x60; in the cluster specified with &#x60;&#x60;cluster_id&#x60;&#x60;.
     *
     * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
     * @param clusterId The Kafka cluster ID.
     * @param brokerId The Kafka broker ID.
     * @param taskType The Kafka broker task type.
     * @return BrokerTaskData
     */
    ClustersClusterIdBrokersBrokerIdTasksTaskTypeGet(ctx _context.Context, clusterId string, brokerId int32, taskType BrokerTaskType) (BrokerTaskData, *_nethttp.Response, error)

    /*
     * ClustersClusterIdBrokersTasksGet List Broker Tasks
     *
     * Returns a list of all tasks for all brokers in the cluster specified with &#x60;&#x60;cluster_id&#x60;&#x60;.
     *
     * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
     * @param clusterId The Kafka cluster ID.
     * @return BrokerTaskDataList
     */
    ClustersClusterIdBrokersTasksGet(ctx _context.Context, clusterId string) (BrokerTaskDataList, *_nethttp.Response, error)

    /*
     * ClustersClusterIdBrokersTasksTaskTypeGet List Broker Tasks of a specific TaskType
     *
     * Returns a list of all broker tasks of specified &#x60;&#x60;task_type&#x60;&#x60; in the cluster specified with &#x60;&#x60;cluster_id&#x60;&#x60;.
     *
     * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
     * @param clusterId The Kafka cluster ID.
     * @param taskType The Kafka broker task type.
     * @return BrokerTaskDataList
     */
    ClustersClusterIdBrokersTasksTaskTypeGet(ctx _context.Context, clusterId string, taskType BrokerTaskType) (BrokerTaskDataList, *_nethttp.Response, error)
}

// BrokerTaskApiService BrokerTaskApi service
type BrokerTaskApiService service

/*
 * ClustersClusterIdBrokersBrokerIdTasksGet List Broker Tasks of a specific Broker
 *
 * Returns a list of all broker tasks for broker specified with &#x60;&#x60;broker_id&#x60;&#x60; in the cluster specified with &#x60;&#x60;cluster_id&#x60;&#x60;.
 *
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterId The Kafka cluster ID.
 * @param brokerId The Kafka broker ID.
 * @return BrokerTaskDataList
 */
func (a *BrokerTaskApiService) ClustersClusterIdBrokersBrokerIdTasksGet(ctx _context.Context, clusterId string, brokerId int32) (BrokerTaskDataList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  BrokerTaskDataList
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/clusters/{cluster_id}/brokers/{broker_id}/tasks"
	localVarPath = strings.Replace(localVarPath, "{"+"cluster_id"+"}", _neturl.QueryEscape(parameterToString(clusterId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"broker_id"+"}", _neturl.QueryEscape(parameterToString(brokerId, "")) , -1)

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

/*
 * ClustersClusterIdBrokersBrokerIdTasksTaskTypeGet Get single Broker Task.
 *
 * Returns a single Broker Task specified with &#x60;&#x60;task_type&#x60;&#x60; for broker specified with &#x60;&#x60;broker_id&#x60;&#x60; in the cluster specified with &#x60;&#x60;cluster_id&#x60;&#x60;.
 *
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterId The Kafka cluster ID.
 * @param brokerId The Kafka broker ID.
 * @param taskType The Kafka broker task type.
 * @return BrokerTaskData
 */
func (a *BrokerTaskApiService) ClustersClusterIdBrokersBrokerIdTasksTaskTypeGet(ctx _context.Context, clusterId string, brokerId int32, taskType BrokerTaskType) (BrokerTaskData, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  BrokerTaskData
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/clusters/{cluster_id}/brokers/{broker_id}/tasks/{task_type}"
	localVarPath = strings.Replace(localVarPath, "{"+"cluster_id"+"}", _neturl.QueryEscape(parameterToString(clusterId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"broker_id"+"}", _neturl.QueryEscape(parameterToString(brokerId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"task_type"+"}", _neturl.QueryEscape(parameterToString(taskType, "")) , -1)

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

/*
 * ClustersClusterIdBrokersTasksGet List Broker Tasks
 *
 * Returns a list of all tasks for all brokers in the cluster specified with &#x60;&#x60;cluster_id&#x60;&#x60;.
 *
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterId The Kafka cluster ID.
 * @return BrokerTaskDataList
 */
func (a *BrokerTaskApiService) ClustersClusterIdBrokersTasksGet(ctx _context.Context, clusterId string) (BrokerTaskDataList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  BrokerTaskDataList
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/clusters/{cluster_id}/brokers/-/tasks"
	localVarPath = strings.Replace(localVarPath, "{"+"cluster_id"+"}", _neturl.QueryEscape(parameterToString(clusterId, "")) , -1)

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

/*
 * ClustersClusterIdBrokersTasksTaskTypeGet List Broker Tasks of a specific TaskType
 *
 * Returns a list of all broker tasks of specified &#x60;&#x60;task_type&#x60;&#x60; in the cluster specified with &#x60;&#x60;cluster_id&#x60;&#x60;.
 *
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param clusterId The Kafka cluster ID.
 * @param taskType The Kafka broker task type.
 * @return BrokerTaskDataList
 */
func (a *BrokerTaskApiService) ClustersClusterIdBrokersTasksTaskTypeGet(ctx _context.Context, clusterId string, taskType BrokerTaskType) (BrokerTaskDataList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  BrokerTaskDataList
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/clusters/{cluster_id}/brokers/-/tasks/{task_type}"
	localVarPath = strings.Replace(localVarPath, "{"+"cluster_id"+"}", _neturl.QueryEscape(parameterToString(clusterId, "")) , -1)

	localVarPath = strings.Replace(localVarPath, "{"+"task_type"+"}", _neturl.QueryEscape(parameterToString(taskType, "")) , -1)

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
