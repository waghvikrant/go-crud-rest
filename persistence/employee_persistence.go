package persistence

import (
	"github.com/waghvikrant/go-crud-rest/logger"
	"github.com/waghvikrant/go-crud-rest/model"
	"time"

	logrus "github.com/sirupsen/logrus"

	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	logger.Init()
	db = connect()
	db.AutoMigrate(model.Employee{})
}

func Get(id uint64) (model.Employee, error) {
	e := model.Employee{}

	result := db.First(&e, id)

	if result.Error != nil {
		logrus.Error(result.Error)
		return model.Employee{}, result.Error
	}

	return e, nil
}

func List() ([]model.Employee, error) {
	e := make([]model.Employee, 0)

	result := db.Find(&e)

	if result.Error != nil {
		logrus.Error(result.Error)
		return nil, result.Error
	}

	return e, nil
}

func Add(e model.Employee) error {
	e.CreatedAt = time.Now()
	e.UpdatedAt = time.Now()

	result := db.Create(&e)

	if result.Error != nil {
		logrus.Error(result.Error)
		return result.Error
	}

	return nil
}

func Update(e model.Employee) error {
	e.UpdatedAt = time.Now()

	result := db.Save(&e)

	if result.Error != nil {
		logrus.Error(result.Error)
		return result.Error
	}

	return nil
}

func Delete(id uint64) error {
	e := model.Employee{}

	result := db.Delete(&e, id)

	if result.Error != nil {
		logrus.Error(result.Error)
		return result.Error
	}

	return nil
}
