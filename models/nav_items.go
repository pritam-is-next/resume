package models

import (
	"github.com/vrianta/Server/model"
)

// var Nav_items = model.New(
// 	"navigation_items",
// 	model.FieldMap{
// 		"Name": {
// 			Name:     "Name",
// 			Nullable: false,
// 			Type:     model.FieldsTypes.VarChar,
// 			Length:   10,
// 			Index: model.Index{
// 				PrimaryKey: true,
// 			},
// 		},
// 		"Href": {
// 			Name:     "Href",
// 			Nullable: false,
// 			Type:     model.FieldsTypes.Text,
// 		},
// 		"Disabled": {
// 			Name:         "Disabled",
// 			Nullable:     false,
// 			DefaultValue: "0",
// 			Type:         model.FieldsTypes.Bool,
// 		},
// 		"Dropdown": {
// 			Name:         "Dropdown",
// 			Nullable:     true,
// 			DefaultValue: "",
// 			Type:         model.FieldsTypes.JSON,
// 		},
// 	},
// )

var Nav_items = model.NewV2("navigation_items", struct {
	Name     model.Field
	Herf     model.Field
	Disabled model.Field
	Dropdown model.Field
}{
	Name: model.Field{
		Name:     "Name",
		Nullable: false,
		Type:     model.FieldsTypes.VarChar,
		Length:   10,
		Index: model.Index{
			PrimaryKey: true,
		},
	},
	Herf: model.Field{
		Name:     "Href",
		Nullable: false,
		Type:     model.FieldsTypes.Text,
	},
	Disabled: model.Field{
		Name:         "Disabled",
		Nullable:     false,
		DefaultValue: "0",
		Type:         model.FieldsTypes.Bool,
	},
	Dropdown: model.Field{
		Name:         "Dropdown",
		Nullable:     true,
		DefaultValue: "",
		Type:         model.FieldsTypes.JSON,
	},
})
