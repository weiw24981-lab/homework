package test

import (
	dbfactory "practise/dbfactory"
	model "practise/model"
	"testing"
)

func TestBlogDemo(t *testing.T) {
	config := model.Config{
		DataTypeCon: model.DataTypeConfig{
			DataType: "mysql",
		},
		MysqlCon: model.MysqlConfig{
			Host:     "localhost",
			Port:     3306,
			Username: "root",
			Password: "root",
			DBName:   "myapp",
		},
	}

	db := dbfactory.NewTestDB(&config)
	if err := db.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{}, &model.Tag{}); err != nil {
		t.Fatalf("auto migrate: %v", err)
	}

	var us model.User
	if err := db.Where(" name = ? and password = ?", "admin", "admin123").First(&us).Error; err != nil {
		panic(err)
	}
}
