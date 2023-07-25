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

// ListLinksResponseData struct for ListLinksResponseData
type ListLinksResponseData struct {
	Kind                 string           `json:"kind"`
	Metadata             ResourceMetadata `json:"metadata"`
	SourceClusterId      *string          `json:"source_cluster_id,omitempty"`
	DestinationClusterId *string          `json:"destination_cluster_id,omitempty"`
	LinkName             string           `json:"link_name"`
	LinkId               string           `json:"link_id"`
	ClusterLinkId        string           `json:"cluster_link_id"`
	TopicNames           []string         `json:"topic_names"`
}
