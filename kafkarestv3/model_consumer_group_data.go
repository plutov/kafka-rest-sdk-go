/*
 * Kafka HTTP API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package kafkarestv3
// ConsumerGroupData struct for ConsumerGroupData
type ConsumerGroupData struct {
	Kind string `json:"kind"`
	Metadata ResourceMetadata `json:"metadata"`
	ClusterId string `json:"cluster_id"`
	ConsumerGroupId string `json:"consumer_group_id"`
	IsSimple bool `json:"is_simple"`
	PartitionAssignor string `json:"partition_assignor"`
	State ConsumerGroupState `json:"state"`
	Coordinator Relationship `json:"coordinator"`
	Consumer Relationship `json:"consumer,omitempty"`
}