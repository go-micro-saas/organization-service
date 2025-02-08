// Package schemas
// Code generated by ikaiguang. <https://github.com/ikaiguang>
package schemas

import (
	migrationuitl "github.com/ikaiguang/go-srv-kit/data/migration"
	datatypes "gorm.io/datatypes"
	gorm "gorm.io/gorm"
	time "time"
)

var _ = time.Time{}

var _ = datatypes.JSON{}

// OrgEventHistorySchema OrgEventHistory
var OrgEventHistorySchema OrgEventHistory

// NewOrgEventHistory new schema
func NewOrgEventHistory() *OrgEventHistory {
	return &OrgEventHistory{}
}

const (
	TableName = "og_org_event_history"

	FieldId                = "id"
	FieldCreatedTime       = "created_time"
	FieldUpdatedTime       = "updated_time"
	FieldEventName         = "event_name"
	FieldEventStatus       = "event_status"
	FieldEventContent      = "event_content"
	FieldEventError        = "event_error"
	FieldRetryEventTime    = "retry_event_time"
	FieldRetryEventCounter = "retry_event_counter"
	FieldRetryEventResult  = "retry_event_result"
)

// OrgEventHistory ENGINE InnoDB CHARSET utf8mb4 COMMENT '事件历史'
type OrgEventHistory struct {
	Id                uint64    `gorm:"column:id;primaryKey;type:uint;autoIncrement;comment:ID" json:"id"`
	CreatedTime       time.Time `gorm:"column:created_time;type:time;not null;comment:创建时间" json:"created_time"`
	UpdatedTime       time.Time `gorm:"column:updated_time;type:time;not null;comment:最后修改时间" json:"updated_time"`
	EventName         string    `gorm:"column:event_name;index;type:string;size:255;not null;default:'';comment:事件名称" json:"event_name"`
	EventStatus       uint32    `gorm:"column:event_status;index;type:uint;not null;default:0;comment:事件状态；1：成功，2：失败，3：重试中，4：重试成功，5：重试失败。。。" json:"event_status"`
	EventContent      string    `gorm:"column:event_content;type:text;not null;comment:事件内容" json:"event_content"`
	EventError        string    `gorm:"column:event_error;type:text;not null;comment:事件错误信息" json:"event_error"`
	RetryEventTime    uint64    `gorm:"column:retry_event_time;type:uint;not null;default:0;comment:重试事件时间" json:"retry_event_time"`
	RetryEventCounter uint32    `gorm:"column:retry_event_counter;type:uint;not null;default:0;comment:重试事件次数" json:"retry_event_counter"`
	RetryEventResult  string    `gorm:"column:retry_event_result;type:text;not null;default:'';comment:重试事件结果" json:"retry_event_result"`
}

// TableName table name
func (s *OrgEventHistory) TableName() string {
	return TableName
}

// CreateTableMigrator create table migrator
func (s *OrgEventHistory) CreateTableMigrator(migrator gorm.Migrator) migrationuitl.MigrationInterface {
	return migrationuitl.NewCreateTable(migrator, migrationuitl.Version, s)
}

// DropTableMigrator create table migrator
func (s *OrgEventHistory) DropTableMigrator(migrator gorm.Migrator) migrationuitl.MigrationInterface {
	return migrationuitl.NewDropTable(migrator, migrationuitl.Version, s)
}

// TableSQL table SQL
func (s *OrgEventHistory) TableSQL() string {
	return `
create table og_org_event_history (
	id bigint unsigned auto_increment comment 'ID',
	created_time datetime(3) not null comment '创建时间',
	updated_time datetime(3) not null comment '最后修改时间',
	event_name varchar(255) not null default '' comment '事件名称',
	event_status integer unsigned not null default 0 comment '事件状态；1：成功，2：失败，3：重试中，4：重试成功，5：重试失败。。。',
	event_content text not null comment '事件内容',
	event_error text not null comment '事件错误信息',
	retry_event_time bigint unsigned not null default 0 comment '重试事件时间',
	retry_event_counter integer unsigned not null default 0 comment '重试事件次数',
	retry_event_result text not null default '' comment '重试事件结果',
	primary key (id),
	key (event_name),
	key (event_status)
) ENGINE InnoDB,
  CHARSET utf8mb4,
  COMMENT '事件历史'
`
}
