package models

import (
	"github.com/vrianta/Server/model"
)

var Settings = model.New(
	"settings",
	map[string]model.Field{
		"element": {
			Name:     "element", // element of the settings
			Nullable: false,
			Type:     model.FieldsTypes.VarChar,
			Length:   20,
			Index: model.Index{
				Index:  true,
				Unique: true,
			},
		},
		"value": {
			Name:     "value", // value of the element
			Nullable: false,
			Type:     model.FieldsTypes.VarChar,
			Length:   20,
		},
	},
)
