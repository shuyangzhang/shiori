// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameCharactersParties = "charactersparties"

// CharactersParties mapped from table <charactersparties>
type CharactersParties struct {
	ID         int32  `gorm:"column:id;type:integer;primaryKey;autoIncrement:true" json:"id"`
	Characters string `gorm:"column:characters;type:jsonb;not null" json:"characters"`
	IsDeleted  bool   `gorm:"column:is_deleted;type:boolean;not null" json:"is_deleted"`
}

// TableName CharactersParties's table name
func (*CharactersParties) TableName() string {
	return TableNameCharactersParties
}