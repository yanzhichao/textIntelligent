package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"textIntelligent/algorithm"
)

type  TextIntelligentController struct {
 beego.Controller
}

type Result struct {
	Data interface{} `json:"data"`
}

func (c *TextIntelligentController) Post() {

	var textInput algorithm.TextInput
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &textInput)
    if err != nil {
		logs.Critical(err)
    	return
	}

	var languageResult = algorithm.LanguageResult{}
	reslut := algorithm.GetLanguage(textInput.Context)

	if reslut != "undefined" {

		languageResult.Data = reslut
		languageResult.Status = "200"
		languageResult.Msg = "识别成功"

	} else {
		languageResult.Data = reslut
		languageResult.Status = "200"
		languageResult.Msg = "识别失败，不支持该语种"
	}

	c.Data["json"] = Result{Data: languageResult}
	c.ServeJSON()
}