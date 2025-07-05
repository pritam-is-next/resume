package models

import (
	"github.com/vrianta/Server/v1/model"
)

// var Nav_items = model.New(
// 	"navigation_items",
// 	model.FieldMap{
// 		"Name": {
// 			Name:     "Name",
// 			Nullable: false,
// 			Type:     model.FieldTypes.VarChar,
// 			Length:   10,
// 			Index: model.Index{
// 				PrimaryKey: true,
// 			},
// 		},
// 		"Href": {
// 			Name:     "Href",
// 			Nullable: false,
// 			Type:     model.FieldTypes.Text,
// 		},
// 		"Disabled": {
// 			Name:         "Disabled",
// 			Nullable:     false,
// 			DefaultValue: "0",
// 			Type:         model.FieldTypes.Bool,
// 		},
// 		"Dropdown": {
// 			Name:         "Dropdown",
// 			Nullable:     true,
// 			DefaultValue: "",
// 			Type:         model.FieldTypes.JSON,
// 		},
// 	},
// )

var Nav_items = model.New("navigation_items", struct {
	Name     model.Field
	Href     model.Field
	Disabled model.Field
	Dropdown model.Field
}{
	Name: model.Field{
		Nullable: false,
		Type:     model.FieldTypes.VarChar,
		Length:   10,
		Index: model.Index{
			PrimaryKey: true,
		},
	},
	Href: model.Field{
		Nullable: false,
		Type:     model.FieldTypes.Text,
	},
	Disabled: model.Field{
		Nullable:     false,
		DefaultValue: "0",
		Type:         model.FieldTypes.Bool,
	},
	Dropdown: model.Field{
		Nullable:     true,
		DefaultValue: "",
		Type:         model.FieldTypes.JSON,
	},
})
