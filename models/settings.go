package models

import (
	"github.com/vrianta/Server/model"
)

var Settings = model.NewV2("settings", struct {
	Element model.Field
	Value   model.Field
}{
	Element: model.Field{
		Name:     "element",
		Nullable: false,
		Type:     model.FieldsTypes.VarChar,
		Length:   20,
		Index: model.Index{
			PrimaryKey: true,
			Index:      true,
			Unique:     false,
		},
	},
	Value: model.Field{
		Name:     "value",
		Nullable: false,
		Type:     model.FieldsTypes.VarChar,
		Length:   20,
	},
})
