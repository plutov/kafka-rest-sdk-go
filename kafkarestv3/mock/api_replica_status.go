// Code generated by mocker. DO NOT EDIT.
// github.com/travisjeffery/mocker
// Source: ../kafka-rest-sdk-go/kafkarestv3/api_replica_status.go

package mock

import (
	context "context"
	net_http "net/http"
	sync "sync"

	github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3 "github.com/confluentinc/kafka-rest-sdk-go/kafkarestv3"
)

// ReplicaStatusApi is a mock of ReplicaStatusApi interface
type ReplicaStatusApi struct {
	lockClustersClusterIdTopicsPartitionsReplicaStatusGet sync.Mutex
	ClustersClusterIdTopicsPartitionsReplicaStatusGetFunc func(ctx context.Context, clusterId string) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.ReplicaStatusDataList, *net_http.Response, error)

	lockClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicaStatusGet sync.Mutex
	ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicaStatusGetFunc func(ctx context.Context, clusterId, topicName string, partitionId int32) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.ReplicaStatusDataList, *net_http.Response, error)

	lockClustersClusterIdTopicsTopicNamePartitionsReplicaStatusGet sync.Mutex
	ClustersClusterIdTopicsTopicNamePartitionsReplicaStatusGetFunc func(ctx context.Context, clusterId, topicName string) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.ReplicaStatusDataList, *net_http.Response, error)

	calls struct {
		ClustersClusterIdTopicsPartitionsReplicaStatusGet []struct {
			Ctx       context.Context
			ClusterId string
		}
		ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicaStatusGet []struct {
			Ctx         context.Context
			ClusterId   string
			TopicName   string
			PartitionId int32
		}
		ClustersClusterIdTopicsTopicNamePartitionsReplicaStatusGet []struct {
			Ctx       context.Context
			ClusterId string
			TopicName string
		}
	}
}

// ClustersClusterIdTopicsPartitionsReplicaStatusGet mocks base method by wrapping the associated func.
func (m *ReplicaStatusApi) ClustersClusterIdTopicsPartitionsReplicaStatusGet(ctx context.Context, clusterId string) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.ReplicaStatusDataList, *net_http.Response, error) {
	m.lockClustersClusterIdTopicsPartitionsReplicaStatusGet.Lock()
	defer m.lockClustersClusterIdTopicsPartitionsReplicaStatusGet.Unlock()

	if m.ClustersClusterIdTopicsPartitionsReplicaStatusGetFunc == nil {
		panic("mocker: ReplicaStatusApi.ClustersClusterIdTopicsPartitionsReplicaStatusGetFunc is nil but ReplicaStatusApi.ClustersClusterIdTopicsPartitionsReplicaStatusGet was called.")
	}

	call := struct {
		Ctx       context.Context
		ClusterId string
	}{
		Ctx:       ctx,
		ClusterId: clusterId,
	}

	m.calls.ClustersClusterIdTopicsPartitionsReplicaStatusGet = append(m.calls.ClustersClusterIdTopicsPartitionsReplicaStatusGet, call)

	return m.ClustersClusterIdTopicsPartitionsReplicaStatusGetFunc(ctx, clusterId)
}

// ClustersClusterIdTopicsPartitionsReplicaStatusGetCalled returns true if ClustersClusterIdTopicsPartitionsReplicaStatusGet was called at least once.
func (m *ReplicaStatusApi) ClustersClusterIdTopicsPartitionsReplicaStatusGetCalled() bool {
	m.lockClustersClusterIdTopicsPartitionsReplicaStatusGet.Lock()
	defer m.lockClustersClusterIdTopicsPartitionsReplicaStatusGet.Unlock()

	return len(m.calls.ClustersClusterIdTopicsPartitionsReplicaStatusGet) > 0
}

// ClustersClusterIdTopicsPartitionsReplicaStatusGetCalls returns the calls made to ClustersClusterIdTopicsPartitionsReplicaStatusGet.
func (m *ReplicaStatusApi) ClustersClusterIdTopicsPartitionsReplicaStatusGetCalls() []struct {
	Ctx       context.Context
	ClusterId string
} {
	m.lockClustersClusterIdTopicsPartitionsReplicaStatusGet.Lock()
	defer m.lockClustersClusterIdTopicsPartitionsReplicaStatusGet.Unlock()

	return m.calls.ClustersClusterIdTopicsPartitionsReplicaStatusGet
}

// ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicaStatusGet mocks base method by wrapping the associated func.
func (m *ReplicaStatusApi) ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicaStatusGet(ctx context.Context, clusterId, topicName string, partitionId int32) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.ReplicaStatusDataList, *net_http.Response, error) {
	m.lockClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicaStatusGet.Lock()
	defer m.lockClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicaStatusGet.Unlock()

	if m.ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicaStatusGetFunc == nil {
		panic("mocker: ReplicaStatusApi.ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicaStatusGetFunc is nil but ReplicaStatusApi.ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicaStatusGet was called.")
	}

	call := struct {
		Ctx         context.Context
		ClusterId   string
		TopicName   string
		PartitionId int32
	}{
		Ctx:         ctx,
		ClusterId:   clusterId,
		TopicName:   topicName,
		PartitionId: partitionId,
	}

	m.calls.ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicaStatusGet = append(m.calls.ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicaStatusGet, call)

	return m.ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicaStatusGetFunc(ctx, clusterId, topicName, partitionId)
}

// ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicaStatusGetCalled returns true if ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicaStatusGet was called at least once.
func (m *ReplicaStatusApi) ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicaStatusGetCalled() bool {
	m.lockClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicaStatusGet.Lock()
	defer m.lockClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicaStatusGet.Unlock()

	return len(m.calls.ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicaStatusGet) > 0
}

// ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicaStatusGetCalls returns the calls made to ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicaStatusGet.
func (m *ReplicaStatusApi) ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicaStatusGetCalls() []struct {
	Ctx         context.Context
	ClusterId   string
	TopicName   string
	PartitionId int32
} {
	m.lockClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicaStatusGet.Lock()
	defer m.lockClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicaStatusGet.Unlock()

	return m.calls.ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicaStatusGet
}

// ClustersClusterIdTopicsTopicNamePartitionsReplicaStatusGet mocks base method by wrapping the associated func.
func (m *ReplicaStatusApi) ClustersClusterIdTopicsTopicNamePartitionsReplicaStatusGet(ctx context.Context, clusterId, topicName string) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.ReplicaStatusDataList, *net_http.Response, error) {
	m.lockClustersClusterIdTopicsTopicNamePartitionsReplicaStatusGet.Lock()
	defer m.lockClustersClusterIdTopicsTopicNamePartitionsReplicaStatusGet.Unlock()

	if m.ClustersClusterIdTopicsTopicNamePartitionsReplicaStatusGetFunc == nil {
		panic("mocker: ReplicaStatusApi.ClustersClusterIdTopicsTopicNamePartitionsReplicaStatusGetFunc is nil but ReplicaStatusApi.ClustersClusterIdTopicsTopicNamePartitionsReplicaStatusGet was called.")
	}

	call := struct {
		Ctx       context.Context
		ClusterId string
		TopicName string
	}{
		Ctx:       ctx,
		ClusterId: clusterId,
		TopicName: topicName,
	}

	m.calls.ClustersClusterIdTopicsTopicNamePartitionsReplicaStatusGet = append(m.calls.ClustersClusterIdTopicsTopicNamePartitionsReplicaStatusGet, call)

	return m.ClustersClusterIdTopicsTopicNamePartitionsReplicaStatusGetFunc(ctx, clusterId, topicName)
}

// ClustersClusterIdTopicsTopicNamePartitionsReplicaStatusGetCalled returns true if ClustersClusterIdTopicsTopicNamePartitionsReplicaStatusGet was called at least once.
func (m *ReplicaStatusApi) ClustersClusterIdTopicsTopicNamePartitionsReplicaStatusGetCalled() bool {
	m.lockClustersClusterIdTopicsTopicNamePartitionsReplicaStatusGet.Lock()
	defer m.lockClustersClusterIdTopicsTopicNamePartitionsReplicaStatusGet.Unlock()

	return len(m.calls.ClustersClusterIdTopicsTopicNamePartitionsReplicaStatusGet) > 0
}

// ClustersClusterIdTopicsTopicNamePartitionsReplicaStatusGetCalls returns the calls made to ClustersClusterIdTopicsTopicNamePartitionsReplicaStatusGet.
func (m *ReplicaStatusApi) ClustersClusterIdTopicsTopicNamePartitionsReplicaStatusGetCalls() []struct {
	Ctx       context.Context
	ClusterId string
	TopicName string
} {
	m.lockClustersClusterIdTopicsTopicNamePartitionsReplicaStatusGet.Lock()
	defer m.lockClustersClusterIdTopicsTopicNamePartitionsReplicaStatusGet.Unlock()

	return m.calls.ClustersClusterIdTopicsTopicNamePartitionsReplicaStatusGet
}

// Reset resets the calls made to the mocked methods.
func (m *ReplicaStatusApi) Reset() {
	m.lockClustersClusterIdTopicsPartitionsReplicaStatusGet.Lock()
	m.calls.ClustersClusterIdTopicsPartitionsReplicaStatusGet = nil
	m.lockClustersClusterIdTopicsPartitionsReplicaStatusGet.Unlock()
	m.lockClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicaStatusGet.Lock()
	m.calls.ClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicaStatusGet = nil
	m.lockClustersClusterIdTopicsTopicNamePartitionsPartitionIdReplicaStatusGet.Unlock()
	m.lockClustersClusterIdTopicsTopicNamePartitionsReplicaStatusGet.Lock()
	m.calls.ClustersClusterIdTopicsTopicNamePartitionsReplicaStatusGet = nil
	m.lockClustersClusterIdTopicsTopicNamePartitionsReplicaStatusGet.Unlock()
}
