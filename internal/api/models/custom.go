package models

import (
	"github.com/soulteary/sparrow/internal/datatypes"
	"github.com/soulteary/sparrow/internal/define"
)

func GetApiWrapper35() (result []datatypes.ModelListItem) {
	modelAPI35 := datatypes.MODEL_OPENAI_API_3_5
	if define.ENABLE_I18N {
		modelAPI35.Title = "OpenAI API (GPT-3.5)"
		modelAPI35.Description = "基于官方接口封装的版本，具备更稳定的服务能力。"
	}
	result = append(result, modelAPI35)
	return result
}

func GetEmptyPlaceHolder() []datatypes.ModelListItem {
	model := datatypes.MODEL_NO_MODELS
	if define.ENABLE_I18N {
		model.Title = "暂无启用的模型"
		model.Description = "请联系管理员启用模型。"
	}
	return []datatypes.ModelListItem{model, model, model}
}
