package controllers

import "gitlab.dreamdev.cn/ebag/word-api/models"

type WordController struct {
	BaseController
}

type SearchWordParam struct {
	Word string `form:"word" valid:"Required"`
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
