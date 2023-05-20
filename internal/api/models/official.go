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

	model35Mobile := datatypes.MODEL_TEXT_DAVINCI_002_RENDER_SHA_MOBILE
	if define.ENABLE_I18N {
		model35Mobile.Title = "默认模型 (GPT-3.5) - 移动版"
		model35Mobile.Description = "我们最快的型号，非常适合大多数日常任务。"
	}
	result = append(result, model35Mobile)

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

	model4Code := datatypes.MODEL_GPT4_CODE
	if define.ENABLE_I18N {
		model4Code.Title = "GPT-4 代码解释器"
		model4Code.Description = "我们最强大的模型，非常适合需要创造力和高级推理的任务。"
	}
	result = append(result, model4Code)

	model4Browsing := datatypes.MODEL_GPT4_BROWSING
	if define.ENABLE_I18N {
		model4Browsing.Title = "GPT-4 浏览器"
		model4Browsing.Description = "知道何时以及如何浏览互联网的实验模型。"
	}
	result = append(result, model4Browsing)

	if define.ENABLE_PLUGIN {
		model4Plugin := datatypes.MODEL_GPT4_PLUGIN
		if define.ENABLE_I18N {
			model4Plugin.Title = "GPT-4 插件"
			model4Plugin.Description = "一个知道何时以及如何使用插件的实验模型"
		}
		result = append(result, model4Plugin)
	}

	model4BrowsingMobile := datatypes.MODEL_GPT4_MOBILE
	if define.ENABLE_I18N {
		model4BrowsingMobile.Title = "GPT-4 - 移动版 v2"
		model4BrowsingMobile.Description = "我们最强大的模型，非常适合需要创造力和高级推理的任务。"
	}
	result = append(result, model4BrowsingMobile)

	// gradually open...
	result = append(result, datatypes.MODEL_OTHER)

	return result
}
