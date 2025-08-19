package models

import (
	"github.com/vrianta/agai/v1/model"
)

var User_details = model.New("user_details", struct {
	Id          *model.Field
	UserId      *model.Field
	FullName    *model.Field
	AboutMe     *model.Field
	Email       *model.Field
	Phone       *model.Field
	Dob         *model.Field
	Gender      *model.Field
	Bio         *model.Field
	Avatar      *model.Field
	Country     *model.Field
	State       *model.Field
	City        *model.Field
	AddressLine *model.Field
	ZipCode     *model.Field
	CreatedAt   *model.Field
	UpdatedAt   *model.Field
}{
	Id: &model.Field{
		Nullable: false,
		Type:     model.FieldTypes.Int,
		Length:   10,
		Index: model.Index{
			PrimaryKey: true,
			Index:      true,
		},
	},

	UserId: Users.Definition.UserId.ToForeignKey("CASCADE", "CASCADE", false, true, true),

	FullName: &model.Field{
		Nullable: false,
		Type:     model.FieldTypes.VarChar,
		Length:   100,
	},

	AboutMe: &model.Field{
		Nullable: true,
		Type:     model.FieldTypes.Text,
	},

	Email: &model.Field{
		Nullable: true,
		Type:     model.FieldTypes.VarChar,
		Length:   150,
		Index: model.Index{
			Unique: true,
		},
	},

	Phone: &model.Field{
		Nullable: true,
		Type:     model.FieldTypes.VarChar,
		Length:   20,
	},

	Dob: &model.Field{
		Nullable: true,
		Type:     model.FieldTypes.Date,
	},

	Gender: &model.Field{
		Nullable:     true,
		Type:         model.FieldTypes.Enum,
		Definition:   []any{"Male", "Female", "Other"},
		DefaultValue: "Male",
	},

	Bio: &model.Field{
		Nullable: true,
		Type:     model.FieldTypes.Text,
	},

	Avatar: &model.Field{
		Nullable: true,
		Type:     model.FieldTypes.VarChar,
		Length:   255,
	},

	Country: &model.Field{
		Nullable: true,
		Type:     model.FieldTypes.VarChar,
		Length:   100,
	},

	State: &model.Field{
		Nullable: true,
		Type:     model.FieldTypes.VarChar,
		Length:   100,
	},

	City: &model.Field{
		Nullable: true,
		Type:     model.FieldTypes.VarChar,
		Length:   100,
	},

	AddressLine: &model.Field{
		Nullable: true,
		Type:     model.FieldTypes.VarChar,
		Length:   255,
	},

	ZipCode: &model.Field{
		Nullable: true,
		Type:     model.FieldTypes.VarChar,
		Length:   20,
	},

	CreatedAt: &model.Field{
		Nullable:     false,
		Type:         model.FieldTypes.Timestamp,
		DefaultValue: "CURRENT_TIMESTAMP",
	},

	UpdatedAt: &model.Field{
		Nullable:     false,
		Type:         model.FieldTypes.Timestamp,
		DefaultValue: "CURRENT_TIMESTAMP",
	},
})
