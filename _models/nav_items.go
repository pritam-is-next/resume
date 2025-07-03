package models

import (
	"github.com/vrianta/Server/model"
)

var Nav_items = model.New(
	"navigation_items",
	map[string]model.Field{
		"Name": {
			Name:     "Name",
			Nullable: false,
			Type:     model.FieldsTypes.VarChar,
			Length:   10,
			Index: model.Index{
				PrimaryKey: true,
			},
		},
		"Href": {
			Name:     "Href",
			Nullable: false,
			Type:     model.FieldsTypes.Text,
		},
	},
)
