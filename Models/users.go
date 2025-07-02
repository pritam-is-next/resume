package Model

import models_handler "github.com/vrianta/Server/models"

var Users = models_handler.New(
	"users",
	map[string]models_handler.Field{
		"userId": {
			Name:     "userId",
			Type:     models_handler.FieldsTypes.VarChar,
			Length:   20,
			Nullable: false,
			Index: models_handler.Index{
				PrimaryKey: true,
				Unique:     false,
				Index:      true,
			},
		},
		"userName": {
			Name:     "userName",
			Type:     models_handler.FieldsTypes.VarChar,
			Length:   30,
			Nullable: false,
			Index: models_handler.Index{
				Unique: true,
				Index:  true,
			},
		},
		"password": {
			Name:     "password",
			Type:     models_handler.FieldsTypes.Text,
			Nullable: false,
		},
		"firstName": {
			Name:     "firstName",
			Type:     models_handler.FieldsTypes.Text,
			Nullable: false,
		},
	},
)
