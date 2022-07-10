// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameArenaSolutions = "arenasolutions"

// ArenaSolutions mapped from table <arenasolutions>
type ArenaSolutions struct {
	ID          int32     `gorm:"column:id;type:integer;primaryKey;autoIncrement:true" json:"id"`
	DefenderID  int32     `gorm:"column:defender_id;type:integer;not null" json:"defender_id"`
	AttackerID  *int32    `gorm:"column:attacker_id;type:integer" json:"attacker_id"`
	IsRandom    bool      `gorm:"column:is_random;type:boolean;not null" json:"is_random"`
	Comment     *string   `gorm:"column:comment;type:text" json:"comment"`
	IsOutdated  bool      `gorm:"column:is_outdated;type:boolean;not null" json:"is_outdated"`
	UpdatedTime time.Time `gorm:"column:updated_time;type:timestamp with time zone;not null" json:"updated_time"`
	IsDeleted   bool      `gorm:"column:is_deleted;type:boolean;not null" json:"is_deleted"`
}

// TableName ArenaSolutions's table name
func (*ArenaSolutions) TableName() string {
	return TableNameArenaSolutions
}