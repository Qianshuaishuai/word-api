package models

import (
	"encoding/json"
	"errors"
	"strings"
	"time"
)

//添加题目
func AddQuestion(userId int, questionData string) (err error) {
	var snow *MSnowflakCurl
	tx := GetDb().Begin()

	var receiverQuestion Question
	json.Unmarshal([]byte(questionData), &receiverQuestion)

	var checkCount int
	tx.Table("t_questions").Where("name = ?", receiverQuestion.Name).Count(&checkCount)
	if checkCount > 0 {
		tx.Rollback()
		return errors.New("已有该单道试题")
	}

	//Step1:先上传题目
	if receiverQuestion.ID == 0 {
		receiverQuestion.ID = snow.GetIntId()
	}
	receiverQuestion.UserID = userId
	receiverQuestion.CreateTime = time.Now()

	err = tx.Table("t_questions").Create(&receiverQuestion).Error
	if err != nil {
		tx.Rollback()
		return errors.New("创建单题失败")
	}

	//Step2:后上传题目截图框
	for f := range receiverQuestion.Frames {
		if receiverQuestion.Frames[f].ID == 0 {
			receiverQuestion.Frames[f].ID = snow.GetIntId()
		}
		receiverQuestion.Frames[f].QuestionID = receiverQuestion.ID
		receiverQuestion.Frames[f].CreateTime = time.Now()
		receiverQuestion.Frames[f].UserID = userId

		err = tx.Table("t_question_frames").Create(&receiverQuestion.Frames[f]).Error
		if err != nil {
			tx.Rollback()
			return errors.New("创建题目截图框失败")
		}
	}

	tx.Commit()
	return nil
}

//获取题目列表（带分页）
func GetQuestionList(userId int, limit int, page int, sort int, q string) (datas []Question, count int, err error) {
	db := GetDb().Table("t_questions")

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

//获取题目对应详情信息（包含截框信息）
func GetQuestionDetail(questionId int) (question Question, err error) {
	GetDb().Table("t_questions").Where("id = ?", questionId).Find(&question)
	GetDb().Table("t_question_frames").Where("question_id = ?", questionId).Find(&question.Frames)

	return question, nil
}

//删除题目
func DeleteQuestion(questionId int) (err error) {
	err = GetDb().Table("t_questions").Where("id = ?", questionId).Delete(Question{}).Error
	err = GetDb().Table("t_question_frames").Where("question_id = ?", questionId).Delete(QuestionFrame{}).Error

	if err != nil {
		return errors.New("删除失败")
	}

	return nil
}

//编辑题目
func EditQuestion(userId, questionId int, questionData string) (err error) {
	//step1:删除原有题目数据
	err = DeleteQuestion(questionId)
	if err != nil {
		return errors.New("删除失败")
	}

	//step2:重建题目数据
	err = AddQuestion(userId, questionData)
	if err != nil {
		return errors.New("重建失败")
	}

	return nil
}
