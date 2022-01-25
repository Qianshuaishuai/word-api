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
	ComboColor   int       `gorm:"column:combo_color" json:"comboColor"`
	ComboWord    string    `gorm:"column:combo_word" json:"comboWord"`
	Background   int       `gorm:"column:background" json:"background"`
	Size         int       `gorm:"column:size" json:"size"`
	StrokeCount  int       `gorm:"column:stroke_count" json:"strokeCount"`
	ComboType    string    `gorm:"column:combo_type" json:"comboType"`
	Time         time.Time `gorm:"column:time" json:"time"`
}

type Search struct {
	ID     int    `gorm:"column:id" json:"id"`
	Word   string `gorm:"column:word" json:"word"`
	Search string `gorm:"column:search" json:"search"`
}

func SearchWord(word string) (data Word, err error) {
	GetDb().Table("words").Where("word = ?", word).Find(&data)

	if data.ID == 0 {
		return data, errors.New("未找到字词")
	}

	return data, nil
}

func SearchList(search string) (datas []Search, err error) {
	var word Word
	GetDb().Table("words").Where("word = ?", search).Find(&word)

	if word.ID != 0 {
		var newSearch Search
		newSearch.Search = search
		newSearch.Word = word.Word
		datas = append(datas, newSearch)
	}

	var searchs []Search
	GetDb().Table("searchs").Where("search = ?", search).Limit(30).Find(&searchs)

	for s := range searchs {
		datas = append(datas, searchs[s])
	}

	return datas, nil
}
