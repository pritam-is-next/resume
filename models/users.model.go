package models

import "github.com/vrianta/agai/v1/model"

var Users = model.New("users", struct {
	UserId    *model.Field
	UserName  *model.Field
	Password  *model.Field
	FirstName *model.Field
	LastName  *model.Field
}{
	UserId: &model.Field{
		Type:     model.FieldTypes.VarChar,
		Length:   100,
		Nullable: false,
		Index: model.Index{
			PrimaryKey: true,
			Unique:     false,
			Index:      true,
		},
	},
	UserName: &model.Field{
		Type:     model.FieldTypes.VarChar,
		Length:   30,
		Nullable: false,
		Index: model.Index{
			Unique: true,
			Index:  true,
		},
	},
	Password: &model.Field{
		Type:     model.FieldTypes.Text,
		Nullable: false,
	},
	FirstName: &model.Field{
		Type:     model.FieldTypes.VarChar,
		Length:   20,
		Nullable: false,
	},
	LastName: &model.Field{
		Type:     model.FieldTypes.VarChar,
		Length:   20,
		Nullable: false,
	},
})
