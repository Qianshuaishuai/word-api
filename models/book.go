package models

import (
	"encoding/json"
	"errors"
	"strings"
	"time"
)

const (
	AI_BASE_URL = "http://240225d4j1.qicp.vip:19231"
)

func AddBook(userId int, bookData string) (err error) {
	var snow *MSnowflakCurl
	tx := GetDb().Begin()

	var receiverBook Book
	json.Unmarshal([]byte(bookData), &receiverBook)

	var checkCount int
	tx.Table("t_books").Where("name = ?", receiverBook.Name).Count(&checkCount)
	if checkCount > 0 {
		tx.Rollback()
		return errors.New("已有该同步书本")
	}

	//Step1:先上传书本
	if receiverBook.ID == 0 {
		receiverBook.ID = snow.GetIntId()
	}
	receiverBook.UserID = userId
	receiverBook.CreateTime = time.Now()

	err = tx.Table("t_books").Create(&receiverBook).Error
	if err != nil {
		tx.Rollback()
		return errors.New("创建书本失败")
	}

	//Step2:后上传书页以及截图框
	for p := range receiverBook.Pages {
		if receiverBook.Pages[p].ID == 0 {
			receiverBook.Pages[p].ID = snow.GetIntId()
		}
		receiverBook.Pages[p].BookID = receiverBook.ID
		receiverBook.Pages[p].CreateTime = time.Now()
		receiverBook.Pages[p].UserID = userId

		//如果imageAi为1且page对应的AiContent字段为空（即代表未Ai处理过），则开启Ai处理
		if receiverBook.ImageAi == 1 {

		}

		err = tx.Table("t_book_pages").Create(&receiverBook.Pages[p]).Error
		if err != nil {
			tx.Rollback()
			return errors.New("创建书页失败")
		}

		for f := range receiverBook.Pages[p].Frames {
			if receiverBook.Pages[p].Frames[f].BookID == 0 {
				receiverBook.Pages[p].Frames[f].BookID = receiverBook.ID
			}
			receiverBook.Pages[p].Frames[f].UserID = userId
			receiverBook.Pages[p].Frames[f].PageID = receiverBook.Pages[p].ID
			receiverBook.Pages[p].Frames[f].CreateTime = time.Now()

			err = tx.Table("t_page_frames").Create(&receiverBook.Pages[p].Frames[f]).Error
			if err != nil {
				tx.Rollback()
				return errors.New("创建截图框失败")
			}
		}
	}

	tx.Commit()
	return nil
}

//获取书本列表（带分页）
func GetBookList(userId int, limit int, page int, sort int, q string) (datas []Book, count int, err error) {
	db := GetDb().Table("t_books")

	//处理分页参数
	var offset int
	if limit > 0 && page > 0 {
		offset = (page - 1) * limit
	}

	// 将搜索字符串按空格拆分
	q = strings.TrimSpace(q)
	var qstring string
	if len(q) > 0 {
		qs := strings.Fields(q)
		for _, v := range qs {
			qstring += "%" + v
		}
		qstring += "%"
	}

	if len(qstring) > 0 {
		db = db.Where("name LIKE ?", qstring)
	}

	var sortStr = "DESC" // 默认时间 降序
	if sort == 1 {
		sortStr = "ASC"
	}

	db.Count(&count)

	db.Limit(limit).
		Offset(offset).
		Order("create_time " + sortStr).
		Scan(&datas)

	for d := range datas {
		datas[d].CreateTimeStr = datas[d].CreateTime.Format(timeLayoutStr)
	}

	return datas, count, nil
}

//获取书本对应详情信息（包含截框信息）
func GetBookDetail(bookId int) (book Book, err error) {
	GetDb().Table("t_books").Where("id = ?", bookId).Find(&book)
	GetDb().Table("t_book_pages").Where("book_id = ?", bookId).Find(&book.Pages)

	for p := range book.Pages {
		GetDb().Table("t_page_frames").Where("page_id = ?", book.Pages[p].ID).Find(&book.Pages[p].Frames)
	}

	return book, nil
}

//删除书本
func DeleteBook(bookId int) (err error) {
	err = GetDb().Table("t_books").Where("id = ?", bookId).Delete(Book{}).Error
	err = GetDb().Table("t_book_pages").Where("book_id = ?", bookId).Delete(Page{}).Error
	err = GetDb().Table("t_page_frames").Where("book_id = ?", bookId).Delete(PageFrame{}).Error

	if err != nil {
		return errors.New("删除失败")
	}

	return nil
}

//编辑书本
func EditBook(userId, bookId int, bookData string) (err error) {
	//step1:删除原有书本数据
	err = DeleteBook(bookId)
	if err != nil {
		return errors.New("删除失败")
	}

	//step2:重建书本数据
	err = AddBook(userId, bookData)
	if err != nil {
		return errors.New("重建失败")
	}

	return nil
}
