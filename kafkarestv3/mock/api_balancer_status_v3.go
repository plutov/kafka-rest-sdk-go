// Code generated by mocker. DO NOT EDIT.
// github.com/travisjeffery/mocker
// Source: ../kafka-rest-sdk-go/kafkarestv3/api_balancer_status_v3.go

package mock

import (
	context "context"
	net_http "net/http"
	sync "sync"

	github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3 "github.com/confluentinc/kafka-rest-sdk-go/kafkarestv3"
)

// BalancerStatusV3Api is a mock of BalancerStatusV3Api interface
type BalancerStatusV3Api struct {
	lockClustersClusterIdBalancerAnyUnevenLoadGet sync.Mutex
	ClustersClusterIdBalancerAnyUnevenLoadGetFunc func(ctx context.Context, clusterId string) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.AnyUnevenLoadData, *net_http.Response, error)

	lockClustersClusterIdBalancerGet sync.Mutex
	ClustersClusterIdBalancerGetFunc func(ctx context.Context, clusterId string) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.BalancerStatusData, *net_http.Response, error)

	calls struct {
		ClustersClusterIdBalancerAnyUnevenLoadGet []struct {
			Ctx       context.Context
			ClusterId string
		}
		ClustersClusterIdBalancerGet []struct {
			Ctx       context.Context
			ClusterId string
		}
	}
}

// ClustersClusterIdBalancerAnyUnevenLoadGet mocks base method by wrapping the associated func.
func (m *BalancerStatusV3Api) ClustersClusterIdBalancerAnyUnevenLoadGet(ctx context.Context, clusterId string) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.AnyUnevenLoadData, *net_http.Response, error) {
	m.lockClustersClusterIdBalancerAnyUnevenLoadGet.Lock()
	defer m.lockClustersClusterIdBalancerAnyUnevenLoadGet.Unlock()

	if m.ClustersClusterIdBalancerAnyUnevenLoadGetFunc == nil {
		panic("mocker: BalancerStatusV3Api.ClustersClusterIdBalancerAnyUnevenLoadGetFunc is nil but BalancerStatusV3Api.ClustersClusterIdBalancerAnyUnevenLoadGet was called.")
	}

	call := struct {
		Ctx       context.Context
		ClusterId string
	}{
		Ctx:       ctx,
		ClusterId: clusterId,
	}

	m.calls.ClustersClusterIdBalancerAnyUnevenLoadGet = append(m.calls.ClustersClusterIdBalancerAnyUnevenLoadGet, call)

	return m.ClustersClusterIdBalancerAnyUnevenLoadGetFunc(ctx, clusterId)
}

// ClustersClusterIdBalancerAnyUnevenLoadGetCalled returns true if ClustersClusterIdBalancerAnyUnevenLoadGet was called at least once.
func (m *BalancerStatusV3Api) ClustersClusterIdBalancerAnyUnevenLoadGetCalled() bool {
	m.lockClustersClusterIdBalancerAnyUnevenLoadGet.Lock()
	defer m.lockClustersClusterIdBalancerAnyUnevenLoadGet.Unlock()

	return len(m.calls.ClustersClusterIdBalancerAnyUnevenLoadGet) > 0
}

// ClustersClusterIdBalancerAnyUnevenLoadGetCalls returns the calls made to ClustersClusterIdBalancerAnyUnevenLoadGet.
func (m *BalancerStatusV3Api) ClustersClusterIdBalancerAnyUnevenLoadGetCalls() []struct {
	Ctx       context.Context
	ClusterId string
} {
	m.lockClustersClusterIdBalancerAnyUnevenLoadGet.Lock()
	defer m.lockClustersClusterIdBalancerAnyUnevenLoadGet.Unlock()

	return m.calls.ClustersClusterIdBalancerAnyUnevenLoadGet
}

// ClustersClusterIdBalancerGet mocks base method by wrapping the associated func.
func (m *BalancerStatusV3Api) ClustersClusterIdBalancerGet(ctx context.Context, clusterId string) (github_com_confluentinc_kafka_rest_sdk_go_kafkarestv3.BalancerStatusData, *net_http.Response, error) {
	m.lockClustersClusterIdBalancerGet.Lock()
	defer m.lockClustersClusterIdBalancerGet.Unlock()

	if m.ClustersClusterIdBalancerGetFunc == nil {
		panic("mocker: BalancerStatusV3Api.ClustersClusterIdBalancerGetFunc is nil but BalancerStatusV3Api.ClustersClusterIdBalancerGet was called.")
	}

	call := struct {
		Ctx       context.Context
		ClusterId string
	}{
		Ctx:       ctx,
		ClusterId: clusterId,
	}

	m.calls.ClustersClusterIdBalancerGet = append(m.calls.ClustersClusterIdBalancerGet, call)

	return m.ClustersClusterIdBalancerGetFunc(ctx, clusterId)
}

// ClustersClusterIdBalancerGetCalled returns true if ClustersClusterIdBalancerGet was called at least once.
func (m *BalancerStatusV3Api) ClustersClusterIdBalancerGetCalled() bool {
	m.lockClustersClusterIdBalancerGet.Lock()
	defer m.lockClustersClusterIdBalancerGet.Unlock()

	return len(m.calls.ClustersClusterIdBalancerGet) > 0
}

// ClustersClusterIdBalancerGetCalls returns the calls made to ClustersClusterIdBalancerGet.
func (m *BalancerStatusV3Api) ClustersClusterIdBalancerGetCalls() []struct {
	Ctx       context.Context
	ClusterId string
} {
	m.lockClustersClusterIdBalancerGet.Lock()
	defer m.lockClustersClusterIdBalancerGet.Unlock()

	return m.calls.ClustersClusterIdBalancerGet
}

// Reset resets the calls made to the mocked methods.
func (m *BalancerStatusV3Api) Reset() {
	m.lockClustersClusterIdBalancerAnyUnevenLoadGet.Lock()
	m.calls.ClustersClusterIdBalancerAnyUnevenLoadGet = nil
	m.lockClustersClusterIdBalancerAnyUnevenLoadGet.Unlock()
	m.lockClustersClusterIdBalancerGet.Lock()
	m.calls.ClustersClusterIdBalancerGet = nil
	m.lockClustersClusterIdBalancerGet.Unlock()
}
