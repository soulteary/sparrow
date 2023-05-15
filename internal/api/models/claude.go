package models

import (
	"github.com/soulteary/sparrow/internal/datatypes"
	"github.com/soulteary/sparrow/internal/define"
)

func GetClaudeModel() (result []datatypes.ModelListItem) {
	model := datatypes.MODEL_CLAUDE

	if define.ENABLE_I18N {
		model.Description = "Claude 是下一代人工智能助手，基于 Anthropic 对培训有用、诚实和无害的人工智能系统的研究。"
	}

	result = append(result, model)
	return result
}
