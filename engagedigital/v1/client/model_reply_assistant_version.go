/*
 * Engage Digital API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package engagedigital

import (
	"time"
)

type ReplyAssistantVersion struct {
	AttachmentsCount int32     `json:"attachments_count,omitempty"`
	Body             string    `json:"body,omitempty"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
	EntryId          string    `json:"entry_id,omitempty"`
	Format           string    `json:"format,omitempty"`
	Id               string    `json:"id"`
	Language         string    `json:"language,omitempty"`
	SourceIds        []string  `json:"source_ids,omitempty"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
}
