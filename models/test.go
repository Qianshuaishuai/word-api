package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/tealeg/xlsx"
)

func Test() {
	GetDb().Table("t_users").Where("phone = ?", "15602335027").Update("username", "Dashuai")
}

func TranslateWordData() {
	xlsxfile2, err := xlsx.OpenFile("./软件3.xlsx")
	if err != nil {
		fmt.Printf("open failed: %s\n", err)
		return
	}

	for _, sheet := range xlsxfile2.Sheets {
		rowCount := sheet.MaxRow
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
				wordType := row.GetCell(0).String()
				pingYin := row.GetCell(1).String()
				pingYinSingle := row.GetCell(2).String()
				// tone := row.GetCell(3).String()
				word := row.GetCell(4).String()
				word2 := row.GetCell(5).String()
				color, _ := row.GetCell(14).Int()
				size, _ := row.GetCell(15).Int()
				comboType := row.GetCell(19).String()
				strokeCount, _ := row.GetCell(20).Int()
				combo := row.GetCell(21).String()
				newWord.WordType = wordType
				newWord.Pinyin = pingYin
				newWord.PinyinSingle = pingYinSingle
				newWord.Word = word
				newWord.Word2 = word2
				newWord.Color = color
				newWord.Size = size
				newWord.ComboType = comboType
				newWord.StrokeCount = strokeCount
				newWord.Time = time.Now()

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
