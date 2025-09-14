package models

import (
	"github.com/vrianta/agai/v1/model"
)

var App_state = model.New("app_state", struct {
	Element *model.Field
	Value   *model.Field
}{
	Element: &model.Field{
		Nullable: false,
		Type:     model.FieldTypes.VarChar,
		Length:   20,
		Index: model.Index{
			PrimaryKey: true,
			Index:      true,
			Unique:     false,
		},
	},
	Value: &model.Field{
		Nullable: false,
		Type:     model.FieldTypes.VarChar,
		Length:   20,
	},
})
