/*
 * Kafka HTTP API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package kafkarestv3

// ReplicaStatusDataAllOf struct for ReplicaStatusDataAllOf
type ReplicaStatusDataAllOf struct {
	ClusterId          string `json:"cluster_id"`
	TopicName          string `json:"topic_name"`
	BrokerId           int32  `json:"broker_id"`
	PartitionId        int32  `json:"partition_id"`
	IsLeader           bool   `json:"is_leader"`
	IsObserver         bool   `json:"is_observer"`
	IsIsrEligible      bool   `json:"is_isr_eligible"`
	IsInIsr            bool   `json:"is_in_isr"`
	IsCaughtUp         bool   `json:"is_caught_up"`
	LogStartOffset     int32  `json:"log_start_offset"`
	LogEndOffset       int32  `json:"log_end_offset"`
	LastCaughtUpTimeMs int32  `json:"last_caught_up_time_ms"`
	LastFetchTimeMs    int32  `json:"last_fetch_time_ms"`
	LinkName           string `json:"link_name,omitempty"`
}
