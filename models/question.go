package models

// func AddQuestion(userId int, imageAI int, frameData string) (err error) {
// 	var snow *MSnowflakCurl
// 	tx := GetDb().Begin()

// 	//Step1:先上传单题
// 	var newQuestion Question
// 	newQuestion.ID = snow.GetIntId()
// 	newQuestion.ImageAi = imageAI
// 	newQuestion.UserID = userId
// 	newQuestion.CreateTime = time.Now()

// 	err = tx.Table("t_questions").Create(&newQuestion).Error
// 	if err != nil {
// 		tx.Rollback()
// 		return errors.New("创建单题失败")
// 	}

// 	//Step2:后上传截框
// 	var uploadFrames []UploadFrame
// 	json.Unmarshal([]byte(frameData), &uploadFrames)

// 	//ToDO:一条post多条上传操作，理应前端给适当交互
// 	for u := range uploadFrames {
// 		uploadFrame := uploadFrames[u]
// 		var newFrame Frame

// 		newFrame.UserID = userId
// 		newFrame.ResourceID = newQuestion.ID
// 		newFrame.Type = TYPE_RESOURCE_FRAME_QUESTION
// 		newFrame.CreateTime = time.Now()
// 		newFrame.Position = uploadFrame.Position
// 		newFrame.ResourceURL = uploadFrame.URL
// 		newFrame.Content = "暂定" //ToDo

// 		err = tx.Table("t_frames").Create(&newFrame).Error
// 		if err != nil {
// 			tx.Rollback()
// 			return errors.New("创建截图框失败")
// 		}

// 	}

// 	return nil
// }

// //获取题目列表（带分页）
// func GetQuestionList(userId int, limit int, page int, sort int) (datas []Question, count int, err error) {
// 	db := GetDb().Table("t_questions")

// 	//处理分页参数
// 	var offset int
// 	if limit > 0 && page > 0 {
// 		offset = (page - 1) * limit
// 	}

// 	var sortStr = "DESC" // 默认时间 降序
// 	if sort == 1 {
// 		sortStr = "ASC"
// 	}

// 	db.Count(&count)

// 	db.Limit(limit).
// 		Offset(offset).
// 		Order("create_time " + sortStr).
// 		Scan(&datas)

// 	return datas, count, nil
// }

// //获取题目对应详情信息（包含截框信息）
// func GetQuestionDetail(questionId int) (detail QuestionDetail, err error) {
// 	GetDb().Table("t_questions").Where("id = ?", questionId).Find(&detail.Question)
// 	GetDb().Table("t_frames").Where("resource_id = ?", questionId).Find(&detail.Frames)

// 	return detail, nil
// }

// //删除题目
// func DeleteQuestion(questionId int) (err error) {
// 	err = GetDb().Table("t_questions").Where("id = ?", questionId).Delete(Question{}).Error
// 	err = GetDb().Table("t_frames").Where("resource_id = ?", questionId).Delete(Frame{}).Error

// 	if err != nil {
// 		return errors.New("删除失败")
// 	}

// 	return nil
// }
