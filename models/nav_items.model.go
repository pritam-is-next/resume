package models

import (
	"github.com/vrianta/agai/v1/model"
)

var Nav_items = model.New("navigation_items", struct {
	Id       *model.Field
	Name     *model.Field
	Href     *model.Field
	Disabled *model.Field
	Dropdown *model.Field
}{
	Id: &model.Field{
		Nullable: false,
		Type:     model.FieldTypes.Int,
		Length:   10,
		Index: model.Index{
			PrimaryKey: true,
			Index:      true,
		},
	},
	Name: &model.Field{
		Nullable: false,
		Type:     model.FieldTypes.VarChar,
		Length:   10,
	},
	Href: &model.Field{
		Nullable: false,
		Type:     model.FieldTypes.Text,
	},
	Disabled: &model.Field{
		Nullable:     false,
		DefaultValue: "0",
		Type:         model.FieldTypes.Bool,
	},
	Dropdown: &model.Field{
		Nullable:     true,
		DefaultValue: "",
		Type:         model.FieldTypes.JSON,
	},
})
