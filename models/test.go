package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/tealeg/xlsx"
)

func Test() {
	GetDb().Table("t_users").Where("phone = ?", "15602335027").Update("username", "Dashuai")
}

func TranslateWordData() {
	xlsxfile2, err := xlsx.OpenFile("./bbbbb.xlsx")
	if err != nil {
		fmt.Printf("open failed: %s\n", err)
		return
	}
	err = GetDb().Table("words").Delete(Word{}).Error
	if err != nil {
		beego.Debug("err:", err)
	}
	for _, sheet := range xlsxfile2.Sheets {
		rowCount := sheet.MaxRow
		nextWordType := ""
		for r := 0; r < rowCount; r++ {
			//固定模板,从第三行开始解析
			if r >= 1 {
				// row := sheet.Row(r)
				row, _ := sheet.Row(r)
				//姓名，学号，考号，年级，班级，考场，座位号
				// studentName := row.Cells[0].String()
				// studentNo := row.Cells[1].String()
				// examNo := row.Cells[2].String()
				var newWord Word
				wordType := row.GetCell(1).String()
				beego.Debug("wordType:", wordType)
				pingYin := row.GetCell(2).String()
				pingYinSingle := row.GetCell(3).String()
				// tone := row.GetCell(3).String()
				word := row.GetCell(4).String()
				word2 := row.GetCell(5).String()
				background, _ := row.GetCell(11).Int()
				color, _ := row.GetCell(12).Int()
				size, _ := row.GetCell(13).Int()
				comboType := row.GetCell(21).String()
				strokeCount, _ := row.GetCell(0).Int()
				combo := row.GetCell(16).String()
				comboColor := row.GetCell(15).String()
				wordType = strings.Replace(wordType, " ", "", -1)
				// 去除换行符
				wordType = strings.Replace(wordType, "\n", "", -1)
				beego.Debug("nextWordType:", nextWordType)
				if wordType == "" {
					newWord.WordType = nextWordType
				} else {
					newWord.WordType = wordType
					nextWordType = newWord.WordType
				}
				newWord.Pinyin = pingYin
				newWord.PinyinSingle = pingYinSingle
				newWord.Word = word
				newWord.Word2 = word2
				newWord.Color = color
				newWord.Size = size
				newWord.ComboType = comboType
				newWord.StrokeCount = strokeCount
				newWord.Background = background
				newWord.Time = time.Now()
				if comboColor != " " {
					if strings.Contains(comboColor, "26") {
						newWord.ComboColor = 26
					}
					if strings.Contains(comboColor, "27") {
						newWord.ComboColor = 27
					}

					strArray := []rune(comboColor)
					newWord.ComboWord = string(strArray[2:3])
				}

				if combo == "" {
					// beego.Debug("word:", word)
				} else {
					if strings.Contains(combo, "\\") {
						comboArray := strings.Split(combo, "\\")
						for c := range comboArray {
							// beego.Debug("c:" + strconv.Itoa(c) + " ," + comboArray[c])
							if c == 0 {
								if strings.Contains(comboArray[c], "(") {

								} else {
									strArray := []rune(comboArray[c])
									for i := 0; i < len(strArray); i++ {
										if i == 0 {
											newWord.Single1 = string(strArray[i])
										}

										if i == 1 {
											newWord.Single2 = string(strArray[i])
										}

										if i == 2 {
											newWord.Single3 = string(strArray[i])
										}
									}
								}

							} else if c == len(comboArray)-1 {
								strArray := []rune(comboArray[c])
								for i := 0; i < len(strArray); i++ {
									var search Search
									search.Search = string(strArray[i])
									search.Word = word
									err = GetDb().Table("searchs").Create(&search).Error
									if err != nil {
										beego.Debug("err:", err)
									}
								}
							}
						}
					} else {
						// beego.Debug("word:", word)
						strArray := []rune(combo)

						for i := 0; i < len(strArray); i++ {
							if i == 0 {
								newWord.Single1 = string(strArray[i])
							}

							if i == 1 {
								newWord.Single2 = string(strArray[i])
							}

							if i == 2 {
								newWord.Single3 = string(strArray[i])
							}
						}
					}
				}
				// if tone != "" {
				// 	rs := []rune(tone)
				// 	rl := len(rs)
				// 	beego.Debug(rl)
				// 	// toneIndexStr := strings(tone[1:4])
				// 	// beego.Debug(toneIndexStr)
				// 	// toneIndex, _ := strconv.Atoi(toneIndexStr)
				// 	// beego.Debug(toneIndex)
				// }
				err = GetDb().Table("words").Create(&newWord).Error
				if err != nil {
					beego.Debug("err:", err)
				}
			}

		}
	}
}

func TranslateWordData2() {
	xlsxfile2, err := xlsx.OpenFile("./6666.xlsx")
	if err != nil {
		fmt.Printf("open failed: %s\n", err)
		return
	}

	if err != nil {
		beego.Debug("err:", err)
	}

	for _, sheet := range xlsxfile2.Sheets {
		rowCount := sheet.MaxRow
		for r := 0; r < rowCount; r++ {
			//固定模板,从第三行开始解析
			if r >= 30014 {
				row, _ := sheet.Row(r)
				word := row.GetCell(1).String()

				var checkWord Word
				GetDb().Table("words").Where("word = ?", word).Find(&checkWord)
				// fmt.Printf("%+v", checkWord)
				word = strings.Replace(word, " ", "", -1)
				// 去除换行符
				word = strings.Replace(word, "\n", "", -1)
				beego.Debug("第" + strconv.Itoa(r) + "行")
				beego.Debug("fixWord:" + word)
				if checkWord.ID != 0 {
					comboType := row.GetCell(3).String()
					beego.Debug("comboType:" + comboType)
					comboType = strings.Replace(comboType, " ", "", -1)
					// 去除换行符
					comboType = strings.Replace(comboType, "\n", "", -1)
					if comboType != "" {
						err = GetDb().Table("words").Where("id = ?", checkWord.ID).Update("combo_type", comboType).Error
						if err != nil {
							beego.Debug("err:", err)
						}
					}
					beego.Debug("enddddddddddddddddddddd")
				}
			}

			// time.Sleep(2 * time.Second)
		}
	}
}

// func TranslateNonData() {
// 	var words []Word
// 	GetDb().Table("words").Find(&words)

// 	for w := range words {
// 		GetDb().Table("words")
// 	}
// }
