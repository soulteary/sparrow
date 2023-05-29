package models

import (
	"github.com/soulteary/sparrow/internal/datatypes"
	"github.com/soulteary/sparrow/internal/define"
)

func GetGitHubTopModel() (result []datatypes.ModelListItem) {
	model := datatypes.MODEL_GITHUB_TOP

	if define.ENABLE_I18N {
		model.Description = "GitHub Top 是一个为了演示如何快速接入外部数据源而存在的示例。"
	}

	result = append(result, model)
	return result
}
