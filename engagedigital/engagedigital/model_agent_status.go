/*
 * Engage Digital API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package engagedigital

type AgentStatus struct {
	AgentId      string               `json:"agent_id,omitempty"`
	Channels     []AgentStatusChannel `json:"channels,omitempty"`
	CustomStatus AgentCustomStatus    `json:"custom_status,omitempty"`
}