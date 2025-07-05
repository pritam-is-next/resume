package models

import "github.com/vrianta/Server/model"

var Users = model.NewV2("users", struct {
	UserId    model.Field
	UserName  model.Field
	Password  model.Field
	FirstName model.Field
}{
	UserId: model.Field{
		Name:     "userId",
		Type:     model.FieldsTypes.VarChar,
		Length:   20,
		Nullable: false,
		Index: model.Index{
			PrimaryKey: true,
			Unique:     false,
			Index:      true,
		},
	},
	UserName: model.Field{
		Name:     "userName",
		Type:     model.FieldsTypes.VarChar,
		Length:   30,
		Nullable: false,
		Index: model.Index{
			Unique: true,
			Index:  true,
		},
	},
	Password: model.Field{
		Name:     "password",
		Type:     model.FieldsTypes.Text,
		Nullable: false,
	},
	FirstName: model.Field{
		Name:     "firstName",
		Type:     model.FieldsTypes.Text,
		Nullable: false,
	},
})
