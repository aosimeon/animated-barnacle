package models

import (
	"gorm.io/gorm"
)

type Url struct {
	gorm.Model
	ShortUrl string
	LongUrl  string
}

// ShortenUrl Shorten url
func ShortenUrl(db *gorm.DB, Url *Url) (err error) {
	err = db.Create(Url).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUrl(db *gorm.DB, url *Url, id string) (err error) {
	err = db.Where("short_url = ?", id).First(url).Error
	if err != nil {
		return err
	}
	return nil
}
