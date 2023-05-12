package models

import (
	"github.com/soulteary/sparrow/internal/datatypes"
	"github.com/soulteary/sparrow/internal/define"
)

func GetFlagStudioModel() (result []datatypes.ModelListItem) {
	model := datatypes.MODEL_FLAGSTUDIO

	if define.ENABLE_I18N {
		model.Description = "FlagStudio 是由 BAAI 旗下的创新应用实验室和 FlagAI 团队开发的文图生成工具。\n\n支持中英等18语的文图生成，旨在为大家提供先进的AI艺术创作体验。"
	}

	result = append(result, model)
	return result
}
