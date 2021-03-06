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

type Role struct {
	AccessHelpCenter          bool      `json:"access_help_center,omitempty"`
	AccessPreviousMessages    bool      `json:"access_previous_messages,omitempty"`
	AccessPullMode            bool      `json:"access_pull_mode,omitempty"`
	AdminStampAnswer          bool      `json:"admin_stamp_answer,omitempty"`
	AnonymizeIdentity         bool      `json:"anonymize_identity,omitempty"`
	ApproveContent            bool      `json:"approve_content,omitempty"`
	AssignIntervention        bool      `json:"assign_intervention,omitempty"`
	AuthorBlockContent        bool      `json:"author_block_content,omitempty"`
	CloseContentThread        bool      `json:"close_content_thread,omitempty"`
	CreateAndDestroyExtension bool      `json:"create_and_destroy_extension,omitempty"`
	CreateCommunity           bool      `json:"create_community,omitempty"`
	CreateContentSource       bool      `json:"create_content_source,omitempty"`
	CreateUser                bool      `json:"create_user,omitempty"`
	CreatedAt                 time.Time `json:"created_at,omitempty"`
	DelayExportContent        bool      `json:"delay_export_content,omitempty"`
	DeleteContentThread       bool      `json:"delete_content_thread,omitempty"`
	ExportIdentity            bool      `json:"export_identity,omitempty"`
	Id                        string    `json:"id,omitempty"`
	ImpersonateUser           bool      `json:"impersonate_user,omitempty"`
	InviteUser                bool      `json:"invite_user,omitempty"`
	Label                     string    `json:"label,omitempty"`
	LockIdentity              bool      `json:"lock_identity,omitempty"`
	ManageApiAccessTokens     bool      `json:"manage_api_access_tokens,omitempty"`
	ManageAppSdkApplications  bool      `json:"manage_app_sdk_applications,omitempty"`
	ManageCategories          bool      `json:"manage_categories,omitempty"`
	ManageChat                bool      `json:"manage_chat,omitempty"`
	ManageCustomFields        bool      `json:"manage_custom_fields,omitempty"`
	ManageCustomNotifications bool      `json:"manage_custom_notifications,omitempty"`
	ManageEmailsTemplates     bool      `json:"manage_emails_templates,omitempty"`
	ManageFolders             bool      `json:"manage_folders,omitempty"`
	ManageIce                 bool      `json:"manage_ice,omitempty"`
	ManageIdentities          bool      `json:"manage_identities,omitempty"`
	ManageMessaging           bool      `json:"manage_messaging,omitempty"`
	ManageOwnNotifications    bool      `json:"manage_own_notifications,omitempty"`
	ManageReplyAssistant      bool      `json:"manage_reply_assistant,omitempty"`
	ManageRoles               bool      `json:"manage_roles,omitempty"`
	ManageRulesEngineRules    bool      `json:"manage_rules_engine_rules,omitempty"`
	ManageTags                bool      `json:"manage_tags,omitempty"`
	ManageTeams               bool      `json:"manage_teams,omitempty"`
	ManageTopologies          bool      `json:"manage_topologies,omitempty"`
	ManageUsersOfMyTeams      bool      `json:"manage_users_of_my_teams,omitempty"`
	MonitorTasks              bool      `json:"monitor_tasks,omitempty"`
	MonitorTeamTasks          bool      `json:"monitor_team_tasks,omitempty"`
	MuteContent               bool      `json:"mute_content,omitempty"`
	OpenContentThread         bool      `json:"open_content_thread,omitempty"`
	PublishContent            bool      `json:"publish_content,omitempty"`
	ReadCommunity             bool      `json:"read_community,omitempty"`
	ReadContentSource         bool      `json:"read_content_source,omitempty"`
	ReadEvent                 bool      `json:"read_event,omitempty"`
	ReadExport                bool      `json:"read_export,omitempty"`
	ReadIdentity              bool      `json:"read_identity,omitempty"`
	ReadOwnStats              bool      `json:"read_own_stats,omitempty"`
	ReadPresence              bool      `json:"read_presence,omitempty"`
	ReadStats                 bool      `json:"read_stats,omitempty"`
	ReadUser                  bool      `json:"read_user,omitempty"`
	ReceiveTasks              bool      `json:"receive_tasks,omitempty"`
	ReplyWithAssistant        bool      `json:"reply_with_assistant,omitempty"`
	SearchContents            bool      `json:"search_contents,omitempty"`
	SearchEvent               bool      `json:"search_event,omitempty"`
	UpdateCommunity           bool      `json:"update_community,omitempty"`
	UpdateContentSource       bool      `json:"update_content_source,omitempty"`
	UpdateExtension           bool      `json:"update_extension,omitempty"`
	UpdateIdentity            bool      `json:"update_identity,omitempty"`
	UpdateIntervention        bool      `json:"update_intervention,omitempty"`
	UpdateOwnIntervention     bool      `json:"update_own_intervention,omitempty"`
	UpdateSettings            bool      `json:"update_settings,omitempty"`
	UpdateTimeSheet           bool      `json:"update_time_sheet,omitempty"`
	UpdateUser                bool      `json:"update_user,omitempty"`
	UpdatedAt                 time.Time `json:"updated_at,omitempty"`
	UseCobrowsing             bool      `json:"use_cobrowsing,omitempty"`
	UseEmoji                  bool      `json:"use_emoji,omitempty"`
}
