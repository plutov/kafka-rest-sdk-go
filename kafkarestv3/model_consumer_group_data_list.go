/*
 * Kafka HTTP API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package kafkarestv3
// ConsumerGroupDataList struct for ConsumerGroupDataList
type ConsumerGroupDataList struct {
	Kind string `json:"kind"`
	Metadata ResourceCollectionMetadata `json:"metadata"`
	Data []ConsumerGroupData `json:"data"`
}
