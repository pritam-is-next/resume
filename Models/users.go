package models

import Models "github.com/vrianta/Server/ModelsHandler"

var users = Models.New(
	"users",
	[]Models.Field{
		{
			Name:     "userId",
			Type:     Models.FieldsTypes.VarChar,
			Length:   20,
			Nullable: false,
			Index: Models.Index{
				PrimaryKey: true,
				Unique:     true,
				Index:      true,
			},
		},
		{
			Name:     "userName",
			Type:     Models.FieldsTypes.VarChar,
			Length:   30,
			Nullable: false,
			Index: Models.Index{
				Unique: true,
				Index:  true,
			},
		},
		{
			Name:     "password",
			Type:     Models.FieldsTypes.Text,
			Nullable: false,
		},
		{
			Name:     "firstName",
			Type:     Models.FieldsTypes.Text,
			Nullable: false,
		},
	},
)
