package models

import (
	"errors"
	"time"
)

type Word struct {
	ID           int       `gorm:"column:id" json:"id"`
	Word         string    `gorm:"column:word" json:"word"`
	Word2        string    `gorm:"column:word2" json:"word2"`
	WordType     string    `gorm:"column:word_type" json:"wordType"`
	Radical      string    `gorm:"column:radical" json:"radical"`
	Single1      string    `gorm:"column:single1" json:"single1"`
	Single2      string    `gorm:"column:single2" json:"single2"`
	Single3      string    `gorm:"column:single3" json:"single3"`
	Pinyin       string    `gorm:"column:pinyin" json:"pinyin"`
	PinyinSingle string    `gorm:"column:pinyin_single" json:"pinyinSingle"`
	Tone         int       `gorm:"column:tone" json:"tone"`
	Type         int       `gorm:"column:type" json:"type"`
	Color        int       `gorm:"column:color" json:"color"`
	Size         int       `gorm:"column:size" json:"size"`
	StrokeCount  int       `gorm:"column:stroke_count" json:"strokeCount"`
	ComboType    string    `gorm:"column:combo_type" json:"comboType"`
	Time         time.Time `gorm:"column:time" json:"time"`
}

func SearchWord(word string) (data Word, err error) {
	GetDb().Table("words").Where("word = ?", word).Find(&data)

	if data.ID == 0 {
		return data, errors.New("未找到字词")
	}

	return data, nil
}
