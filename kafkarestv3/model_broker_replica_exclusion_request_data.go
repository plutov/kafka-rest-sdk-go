/*
 * REST Admin API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 3.0.0
 * Contact: kafka-clients-proxy-team@confluent.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package kafkarestv3

// BrokerReplicaExclusionRequestData struct for BrokerReplicaExclusionRequestData
type BrokerReplicaExclusionRequestData struct {
	BrokerId int32  `json:"broker_id"`
	Reason   string `json:"reason"`
}
