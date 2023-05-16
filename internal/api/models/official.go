package models

import (
	"github.com/soulteary/sparrow/internal/datatypes"
	"github.com/soulteary/sparrow/internal/define"
)

func GetOfficialModels() (result []datatypes.ModelListItem) {
	if define.ENABLE_PLUGIN {
		model := datatypes.MODEL_TEXT_DAVINCI_002_PLUGINS
		if define.ENABLE_I18N {
			model.Title = "插件模型 (Plugins)"
			model.Description = "一个试验中的模型，懂得如何使用插件。"
		}
		result = append(result, model)
	}

	model35 := datatypes.MODEL_TEXT_DAVINCI_002_RENDER_SHA
	if define.ENABLE_I18N {
		model35.Title = "默认模型 (GPT-3.5)"
		model35.Description = "我们最快的型号，非常适合大多数日常任务。"
	}
	result = append(result, model35)

	// discard
	// modelLegacy := datatypes.MODEL_TEXT_DAVINCI_002_RENDER_PAID
	// if define.ENABLE_I18N {
	// 	modelLegacy.Title = "经典模型 (GPT-3.5)"
	// 	modelLegacy.Description = "早先时候的 ChatGPT Plus 模型"
	// }

	model4 := datatypes.MODEL_GPT4
	if define.ENABLE_I18N {
		model4.Title = "GPT-4"
		model4.Description = "我们最强大的模型，非常适合需要创造力和高级推理的任务。"
	}
	result = append(result, model4)

	// gradually open...
	result = append(result, datatypes.MODEL_GPT4_BROWSING)
	result = append(result, datatypes.MODEL_GPT4_CODE)
	result = append(result, datatypes.MODEL_GPT4_PLUGIN)
	result = append(result, datatypes.MODEL_OTHER)

	return result
}
