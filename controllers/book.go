package controllers

import "gitlab.dreamdev.cn/ebag/knowtech-api/models"

type BookController struct {
	BaseController
}

type AddBookParam struct {
	UserID   int    `form:"userId" valid:"Required"`
	BookData string `form:"bookData" valid:"Required"`
}

type EditBookParam struct {
	UserID   int    `form:"userId" valid:"Required"`
	BookID   int    `form:"bookId" valid:"Required"`
	BookData string `form:"bookData" valid:"Required"`
}

type GetBookDetailParam struct {
	BookID int `form:"bookId" valid:"Required"`
}

type DeleteBookParam struct {
	BookID int `form:"bookId" valid:"Required"`
}

type GetBookListParam struct {
	UserID int    `form:"userId" valid:"Required"`
	Limit  int    `form:"limit" valid:"Required"`
	Page   int    `form:"page" valid:"Required"`
	Sort   int    `form:"sort"`
	Q      string `form:"q"`
}

// @Title 用户添加新书本
// @Description 用户添加新书本
// @Param userId form int true 用户id
// @Param bookData form string true 书本数据
// @Router /v1/book/add [post]
func (b *BookController) AddBook() {
	datas := b.GetResponseData()
	//get params
	params := &AddBookParam{}
	if b.CheckParams(datas, params) {
		err := models.AddBook(params.UserID, params.BookData)
		b.IfErr(datas, models.RESP_PARAM_ERR, err)
	}
	b.jsonEcho(datas)
}

// @Title 用户编辑书本
// @Description 用户编辑书本
// @Param userId form int true 用户id
// @Param bookId form int true 书本id
// @Param bookData form string true 书本数据
// @Router /v1/book/edit [post]
func (b *BookController) EditBook() {
	datas := b.GetResponseData()
	//get params
	params := &EditBookParam{}
	if b.CheckParams(datas, params) {
		err := models.EditBook(params.UserID, params.BookID, params.BookData)
		b.IfErr(datas, models.RESP_PARAM_ERR, err)
	}
	b.jsonEcho(datas)
}

// @Title 删除对应书本
// @Description 删除对应书本
// @Param bookId form int true 书本id
// @Router /v1/book/delete [post]
func (b *BookController) DeleteBook() {
	datas := b.GetResponseData()
	//get params
	params := &DeleteBookParam{}
	if b.CheckParams(datas, params) {
		err := models.DeleteBook(params.BookID)
		b.IfErr(datas, models.RESP_PARAM_ERR, err)
	}
	b.jsonEcho(datas)
}

// @Title 用户获取书本列表
// @Description 用户获取书本列表
// @Param userId form int true 用户id
// @Param limit form int true 一页个数
// @Param page form int true 当前个数
// @Param sort form int true 是否排序-时间
// @Param q form string true 搜索关键字
// @Router /v1/book/list [post]
func (b *BookController) GetBookList() {
	datas := b.GetResponseData()
	//get params
	params := &GetBookListParam{}
	if b.CheckParams(datas, params) {
		data, count, err := models.GetBookList(params.UserID, params.Limit, params.Page, params.Sort, params.Q)
		datas["F_data"] = data
		datas["F_count"] = count
		b.IfErr(datas, models.RESP_NO_RESULT, err)
	}
	b.jsonEcho(datas)
}

// @Title 用户获取书本详情
// @Description 用户获取书本详情
// @Param userId form int true 用户id
// @Param limit form int true 一页个数
// @Param page form int true 当前个数
// @Param sort form int true 是否排序-时间
// @Param q form string true 搜索关键字
// @Router /v1/book/detail [post]
func (b *BookController) GetBookDetail() {
	datas := b.GetResponseData()
	//get params
	params := &GetBookDetailParam{}
	if b.CheckParams(datas, params) {
		data, err := models.GetBookDetail(params.BookID)
		datas["F_data"] = data
		b.IfErr(datas, models.RESP_NO_RESULT, err)
	}
	b.jsonEcho(datas)
}
