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
		model35.Description = "针对速度进行了优化，目前可供 Plus 用户使用"
	}
	result = append(result, model35)

	modelLegacy := datatypes.MODEL_TEXT_DAVINCI_002_RENDER_PAID
	if define.ENABLE_I18N {
		modelLegacy.Title = "经典模型 (GPT-3.5)"
		modelLegacy.Description = "早先时候的 ChatGPT Plus 模型"
	}

	model4 := datatypes.MODEL_GPT4
	if define.ENABLE_I18N {
		model4.Title = "GPT-4"
		model4.Description = "我们最先进的模型，可供 Plus 订阅者使用。\n\nGPT-4 擅长处理需要高级推理、复杂指令理解和更多创造力的任务。"
	}

	result = append(result, []datatypes.ModelListItem{modelLegacy, model4}...)
	return result
}
