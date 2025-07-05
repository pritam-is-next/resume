package models

import (
	"github.com/vrianta/Server/v1/model"
)

var Settings = model.New("settings", struct {
	Element model.Field
	Value   model.Field
}{
	Element: model.Field{
		Nullable: false,
		Type:     model.FieldTypes.VarChar,
		Length:   20,
		Index: model.Index{
			PrimaryKey: true,
			Index:      true,
			Unique:     false,
		},
	},
	Value: model.Field{
		Nullable: false,
		Type:     model.FieldTypes.VarChar,
		Length:   20,
	},
})
