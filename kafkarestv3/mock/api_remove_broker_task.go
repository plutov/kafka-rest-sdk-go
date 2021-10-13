// Code generated by mocker. DO NOT EDIT.
// github.com/travisjeffery/mocker
// Source: ../kafka-rest-sdk-go/kafkarestv3/api_remove_broker_task.go

package mock

import (
	context "context"
	net_http "net/http"
	sync "sync"

	github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3 "github.com/confluentinc/kafka-rest-sdk-go/kafkarestv3"
)

// RemoveBrokerTaskApi is a mock of RemoveBrokerTaskApi interface
type RemoveBrokerTaskApi struct {
	lockClustersClusterIdRemoveBrokerTasksBrokerIdGet sync.Mutex
	ClustersClusterIdRemoveBrokerTasksBrokerIdGetFunc func(ctx context.Context, clusterId string, brokerId int32) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.RemoveBrokerTaskData, *net_http.Response, error)

	lockClustersClusterIdRemoveBrokerTasksGet sync.Mutex
	ClustersClusterIdRemoveBrokerTasksGetFunc func(ctx context.Context, clusterId string) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.RemoveBrokerTaskDataList, *net_http.Response, error)

	calls struct {
		ClustersClusterIdRemoveBrokerTasksBrokerIdGet []struct {
			Ctx       context.Context
			ClusterId string
			BrokerId  int32
		}
		ClustersClusterIdRemoveBrokerTasksGet []struct {
			Ctx       context.Context
			ClusterId string
		}
	}
}

// ClustersClusterIdRemoveBrokerTasksBrokerIdGet mocks base method by wrapping the associated func.
func (m *RemoveBrokerTaskApi) ClustersClusterIdRemoveBrokerTasksBrokerIdGet(ctx context.Context, clusterId string, brokerId int32) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.RemoveBrokerTaskData, *net_http.Response, error) {
	m.lockClustersClusterIdRemoveBrokerTasksBrokerIdGet.Lock()
	defer m.lockClustersClusterIdRemoveBrokerTasksBrokerIdGet.Unlock()

	if m.ClustersClusterIdRemoveBrokerTasksBrokerIdGetFunc == nil {
		panic("mocker: RemoveBrokerTaskApi.ClustersClusterIdRemoveBrokerTasksBrokerIdGetFunc is nil but RemoveBrokerTaskApi.ClustersClusterIdRemoveBrokerTasksBrokerIdGet was called.")
	}

	call := struct {
		Ctx       context.Context
		ClusterId string
		BrokerId  int32
	}{
		Ctx:       ctx,
		ClusterId: clusterId,
		BrokerId:  brokerId,
	}

	m.calls.ClustersClusterIdRemoveBrokerTasksBrokerIdGet = append(m.calls.ClustersClusterIdRemoveBrokerTasksBrokerIdGet, call)

	return m.ClustersClusterIdRemoveBrokerTasksBrokerIdGetFunc(ctx, clusterId, brokerId)
}

// ClustersClusterIdRemoveBrokerTasksBrokerIdGetCalled returns true if ClustersClusterIdRemoveBrokerTasksBrokerIdGet was called at least once.
func (m *RemoveBrokerTaskApi) ClustersClusterIdRemoveBrokerTasksBrokerIdGetCalled() bool {
	m.lockClustersClusterIdRemoveBrokerTasksBrokerIdGet.Lock()
	defer m.lockClustersClusterIdRemoveBrokerTasksBrokerIdGet.Unlock()

	return len(m.calls.ClustersClusterIdRemoveBrokerTasksBrokerIdGet) > 0
}

// ClustersClusterIdRemoveBrokerTasksBrokerIdGetCalls returns the calls made to ClustersClusterIdRemoveBrokerTasksBrokerIdGet.
func (m *RemoveBrokerTaskApi) ClustersClusterIdRemoveBrokerTasksBrokerIdGetCalls() []struct {
	Ctx       context.Context
	ClusterId string
	BrokerId  int32
} {
	m.lockClustersClusterIdRemoveBrokerTasksBrokerIdGet.Lock()
	defer m.lockClustersClusterIdRemoveBrokerTasksBrokerIdGet.Unlock()

	return m.calls.ClustersClusterIdRemoveBrokerTasksBrokerIdGet
}

// ClustersClusterIdRemoveBrokerTasksGet mocks base method by wrapping the associated func.
func (m *RemoveBrokerTaskApi) ClustersClusterIdRemoveBrokerTasksGet(ctx context.Context, clusterId string) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.RemoveBrokerTaskDataList, *net_http.Response, error) {
	m.lockClustersClusterIdRemoveBrokerTasksGet.Lock()
	defer m.lockClustersClusterIdRemoveBrokerTasksGet.Unlock()

	if m.ClustersClusterIdRemoveBrokerTasksGetFunc == nil {
		panic("mocker: RemoveBrokerTaskApi.ClustersClusterIdRemoveBrokerTasksGetFunc is nil but RemoveBrokerTaskApi.ClustersClusterIdRemoveBrokerTasksGet was called.")
	}

	call := struct {
		Ctx       context.Context
		ClusterId string
	}{
		Ctx:       ctx,
		ClusterId: clusterId,
	}

	m.calls.ClustersClusterIdRemoveBrokerTasksGet = append(m.calls.ClustersClusterIdRemoveBrokerTasksGet, call)

	return m.ClustersClusterIdRemoveBrokerTasksGetFunc(ctx, clusterId)
}

// ClustersClusterIdRemoveBrokerTasksGetCalled returns true if ClustersClusterIdRemoveBrokerTasksGet was called at least once.
func (m *RemoveBrokerTaskApi) ClustersClusterIdRemoveBrokerTasksGetCalled() bool {
	m.lockClustersClusterIdRemoveBrokerTasksGet.Lock()
	defer m.lockClustersClusterIdRemoveBrokerTasksGet.Unlock()

	return len(m.calls.ClustersClusterIdRemoveBrokerTasksGet) > 0
}

// ClustersClusterIdRemoveBrokerTasksGetCalls returns the calls made to ClustersClusterIdRemoveBrokerTasksGet.
func (m *RemoveBrokerTaskApi) ClustersClusterIdRemoveBrokerTasksGetCalls() []struct {
	Ctx       context.Context
	ClusterId string
} {
	m.lockClustersClusterIdRemoveBrokerTasksGet.Lock()
	defer m.lockClustersClusterIdRemoveBrokerTasksGet.Unlock()

	return m.calls.ClustersClusterIdRemoveBrokerTasksGet
}

// Reset resets the calls made to the mocked methods.
func (m *RemoveBrokerTaskApi) Reset() {
	m.lockClustersClusterIdRemoveBrokerTasksBrokerIdGet.Lock()
	m.calls.ClustersClusterIdRemoveBrokerTasksBrokerIdGet = nil
	m.lockClustersClusterIdRemoveBrokerTasksBrokerIdGet.Unlock()
	m.lockClustersClusterIdRemoveBrokerTasksGet.Lock()
	m.calls.ClustersClusterIdRemoveBrokerTasksGet = nil
	m.lockClustersClusterIdRemoveBrokerTasksGet.Unlock()
}
