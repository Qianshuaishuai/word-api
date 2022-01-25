package controllers

import "gitlab.dreamdev.cn/ebag/word-api/models"

type WordController struct {
	BaseController
}

type SearchWordParam struct {
	Word string `form:"word" valid:"Required"`
}

type SearchListParam struct {
	Search string `form:"search" valid:"Required"`
}

// @Router /v1/word/search [get]
func (w *WordController) GetSearchWord() {
	datas := w.GetResponseData()
	//get params
	params := &SearchWordParam{}
	if w.CheckParams(datas, params) {
		data, err := models.SearchWord(params.Word)
		datas["F_data"] = data
		w.IfErr(datas, models.RESP_NO_RESULT, err)
	}
	w.jsonEcho(datas)
}

// @Router /v1/word/search/list [get]
func (w *WordController) GetSearchList() {
	datas := w.GetResponseData()
	//get params
	params := &SearchListParam{}
	if w.CheckParams(datas, params) {
		data, err := models.SearchList(params.Search)
		datas["F_data"] = data
		w.IfErr(datas, models.RESP_NO_RESULT, err)
	}
	w.jsonEcho(datas)
}
