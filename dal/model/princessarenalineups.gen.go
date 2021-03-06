// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNamePrincessArenaLineups = "princessarenalineups"

// PrincessArenaLineups mapped from table <princessarenalineups>
type PrincessArenaLineups struct {
	ID              int32      `gorm:"column:id;type:integer;primaryKey;autoIncrement:true" json:"id"`
	UserID          string     `gorm:"column:user_id;type:character varying(128);not null" json:"user_id"`
	PrincessArenaID *time.Time `gorm:"column:princess_arena_id;type:date" json:"princess_arena_id"`
	DefenderID      *int32     `gorm:"column:defender_id;type:integer" json:"defender_id"`
	DefenderName    string     `gorm:"column:defender_name;type:character varying(128);not null" json:"defender_name"`
	FirstLineupID   *int32     `gorm:"column:first_lineup_id;type:integer" json:"first_lineup_id"`
	SecondLineupID  *int32     `gorm:"column:second_lineup_id;type:integer" json:"second_lineup_id"`
	ThirdLineupID   *int32     `gorm:"column:third_lineup_id;type:integer" json:"third_lineup_id"`
	IsOutdated      bool       `gorm:"column:is_outdated;type:boolean;not null" json:"is_outdated"`
	UpdatedTime     time.Time  `gorm:"column:updated_time;type:timestamp with time zone;not null" json:"updated_time"`
	IsDeleted       bool       `gorm:"column:is_deleted;type:boolean;not null" json:"is_deleted"`
}

// TableName PrincessArenaLineups's table name
func (*PrincessArenaLineups) TableName() string {
	return TableNamePrincessArenaLineups
}
