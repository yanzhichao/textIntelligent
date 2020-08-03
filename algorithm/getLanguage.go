package algorithm

import (
	"github.com/astaxie/beego/logs"
)
import "github.com/chrisport/go-lang-detector/langdet/langdetdef"

type TextInput struct {
	Context string  `json: "context,omitempty"`
}

type LanguageResult struct {
	Status string `json: "status,omitempty"`
	Msg	string `json "msg,omitempty"`
	Data string `json data,omitempty`
}


func GetLanguage(context string)  string {
	detector := langdetdef.NewWithDefaultLanguages()

	result := detector.GetClosestLanguage(context)
	logs.Info("input word is ",context)
	logs.Info("GetClosestLanguage returns:\n", "    ", result)
    return result

	//fullResults := detector.GetLanguages("义勇军进行曲")
	//logs.Debug("GetLanguages for Chinese returns:")
	//var result string
	//for _, r := range fullResults {
	//	logs.Debug("    ", r.Name, r.Confidence, "%")
	//	result = r.Name
	//}
	//return result
}