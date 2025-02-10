// Package po
// Code generated by ikaiguang. <https://github.com/ikaiguang>
package po

import (
	enumv1 "github.com/go-micro-saas/organization-service/api/org-service/v1/enums"
	datatypes "gorm.io/datatypes"
	time "time"
)

var _ = time.Time{}

var _ = datatypes.JSON{}

// OrgEventHistory ENGINE InnoDB CHARSET utf8mb4 COMMENT '事件历史'
type OrgEventHistory struct {
	Id                uint64                                   `gorm:"column:id;primaryKey" json:"id"`                        // ID
	CreatedTime       time.Time                                `gorm:"column:created_time" json:"created_time"`               // 创建时间
	UpdatedTime       time.Time                                `gorm:"column:updated_time" json:"updated_time"`               // 最后修改时间
	EventName         string                                   `gorm:"column:event_name" json:"event_name"`                   // 事件名称
	EventStatus       enumv1.OrgEventStatusEnum_OrgEventStatus `gorm:"column:event_status" json:"event_status"`               // 事件状态；1：成功，2：失败，3：重试中，4：重试成功，5：重试失败。。。
	EventContent      string                                   `gorm:"column:event_content" json:"event_content"`             // 事件内容
	EventError        string                                   `gorm:"column:event_error" json:"event_error"`                 // 事件错误信息
	RetryEventTime    uint64                                   `gorm:"column:retry_event_time" json:"retry_event_time"`       // 重试事件时间
	RetryEventCounter uint32                                   `gorm:"column:retry_event_counter" json:"retry_event_counter"` // 重试事件次数
	RetryEventResult  string                                   `gorm:"column:retry_event_result" json:"retry_event_result"`   // 重试事件结果
}
