package docman

import (
	"fmt"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	once sync.Once
)

type Job struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	gorm.Model
}

func NewDB() error {
	arr := []string{Cfg.DbHost, Cfg.DbPort, Cfg.DbUser, Cfg.DbName, Cfg.DbPassword}

	for i, v := range arr {
		if v == "" {
			return fmt.Errorf("DB config error: %d", i)
		}
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", Cfg.DbHost, Cfg.DbPort, Cfg.DbUser, Cfg.DbName, Cfg.DbPassword)
	var err error
	once.Do(func() {
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		DB = db
		db.AutoMigrate(&Job{})
	},
	)
	return err
}

func (j *Job) Create() error {
	if err := DB.Create(&j).Error; err != nil {
		return err
	}
	return nil
}

func (j *Job) Update() error {
	if err := DB.Save(&j).Error; err != nil {
		return err
	}
	return nil
}

func (j *Job) Delete() error {
	if err := DB.Delete(&j).Error; err != nil {
		return err
	}
	return nil
}

func FindByID(id uint) ([]Job, error) {
	jobs := []Job{}
	err := DB.Limit(1).Where("ID=?", id).Find(&jobs).Error
	return jobs, err
}

func FindAll() ([]Job, error) {
	var jobs []Job
	if err := DB.Find(&jobs).Error; err != nil {
		return nil, err
	}
	return jobs, nil
}
