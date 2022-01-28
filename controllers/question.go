package controllers

import "gitlab.dreamdev.cn/ebag/word-api/models"

type QuestionController struct {
	BaseController
}

type AddQuestionParam struct {
	UserID       int    `form:"userId" valid:"Required"`
	QuestionData string `form:"questionData" valid:"Required"`
}

type EditQuestionParam struct {
	UserID       int    `form:"userId" valid:"Required"`
	QuestionID   int    `form:"questionId" valid:"Required"`
	QuestionData string `form:"questionData" valid:"Required"`
}

type GetQuestionDetailParam struct {
	QuestionID int `form:"questionId" valid:"Required"`
}

type DeleteQuestionParam struct {
	QuestionID int `form:"questionId" valid:"Required"`
}

type GetQuestionListParam struct {
	UserID int    `form:"userId" valid:"Required"`
	Limit  int    `form:"limit" valid:"Required"`
	Page   int    `form:"page" valid:"Required"`
	Sort   int    `form:"sort"`
	Q      string `form:"q"`
}

// @Title 用户添加新题目
// @Description 用户添加新题目
// @Param userId form int true 用户id
// @Param questionData form string true 题目数据
// @Router /v1/question/add [post]
func (q *QuestionController) AddQuestion() {
	datas := q.GetResponseData()
	//get params
	params := &AddQuestionParam{}
	if q.CheckParams(datas, params) {
		err := models.AddQuestion(params.UserID, params.QuestionData)
		q.IfErr(datas, models.RESP_PARAM_ERR, err)
	}
	q.jsonEcho(datas)
}

// @Title 用户编辑题目
// @Description 用户编辑题目
// @Param userId form int true 用户id
// @Param questionId form int true 题目id
// @Param questionData form string true 题目数据
// @Router /v1/question/edit [post]
func (q *QuestionController) EditQuestion() {
	datas := q.GetResponseData()
	//get params
	params := &EditQuestionParam{}
	if q.CheckParams(datas, params) {
		err := models.EditQuestion(params.UserID, params.QuestionID, params.QuestionData)
		q.IfErr(datas, models.RESP_PARAM_ERR, err)
	}
	q.jsonEcho(datas)
}

// @Title 删除对应题目
// @Description 删除对应题目
// @Param questionId form int true 题目id
// @Router /v1/question/delete [post]
func (q *QuestionController) DeleteQuestion() {
	datas := q.GetResponseData()
	//get params
	params := &DeleteQuestionParam{}
	if q.CheckParams(datas, params) {
		err := models.DeleteQuestion(params.QuestionID)
		q.IfErr(datas, models.RESP_PARAM_ERR, err)
	}
	q.jsonEcho(datas)
}

// @Title 用户获取题目列表
// @Description 用户获取题目列表
// @Param userId form int true 用户id
// @Param limit form int true 一页个数
// @Param page form int true 当前个数
// @Param sort form int true 是否排序-时间
// @Param q form string true 搜索关键字
// @Router /v1/question/list [post]
func (q *QuestionController) GetQuestionList() {
	datas := q.GetResponseData()
	//get params
	params := &GetQuestionListParam{}
	if q.CheckParams(datas, params) {
		data, count, err := models.GetQuestionList(params.UserID, params.Limit, params.Page, params.Sort, params.Q)
		datas["F_data"] = data
		datas["F_count"] = count
		q.IfErr(datas, models.RESP_NO_RESULT, err)
	}
	q.jsonEcho(datas)
}

// @Title 用户获取题目详情
// @Description 用户获取题目详情
// @Param userId form int true 用户id
// @Param limit form int true 一页个数
// @Param page form int true 当前个数
// @Param sort form int true 是否排序-时间
// @Param q form string true 搜索关键字
// @Router /v1/question/detail [post]
func (q *QuestionController) GetQuestionDetail() {
	datas := q.GetResponseData()
	//get params
	params := &GetQuestionDetailParam{}
	if q.CheckParams(datas, params) {
		data, err := models.GetQuestionDetail(params.QuestionID)
		datas["F_data"] = data
		q.IfErr(datas, models.RESP_NO_RESULT, err)
	}
	q.jsonEcho(datas)
}
