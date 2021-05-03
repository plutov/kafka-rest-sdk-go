/*
 * Kafka HTTP API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package kafkarestv3
// AclDataAllOf struct for AclDataAllOf
type AclDataAllOf struct {
	ClusterId string `json:"cluster_id"`
	ResourceType AclResourceType `json:"resource_type"`
	ResourceName string `json:"resource_name"`
	PatternType AclPatternType `json:"pattern_type"`
	Principal string `json:"principal"`
	Host string `json:"host"`
	Operation AclOperation `json:"operation"`
	Permission AclPermission `json:"permission"`
}
