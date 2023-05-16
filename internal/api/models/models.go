package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soulteary/sparrow/internal/datatypes"
	"github.com/soulteary/sparrow/internal/define"
)

func GetModels(c *gin.Context) {
	var modelList []datatypes.ModelListItem

	if define.ENABLE_CLAUDE {
		model := GetClaudeModel()
		if define.ENABLE_CLAUDE_ONLY {
			c.JSON(http.StatusOK, datatypes.Models{Models: model})
			return
		}
		modelList = append(modelList, model...)
	}

	if define.ENABLE_MIDJOURNEY {
		model := GetMidJourneyModel()
		if define.ENABLE_MIDJOURNEY_ONLY {
			c.JSON(http.StatusOK, datatypes.Models{Models: model})
			return
		}
		modelList = append(modelList, model...)
	}

	if define.ENABLE_FLAGSTUDIO {
		model := GetFlagStudioModel()
		if define.ENABLE_FLAGSTUDIO_ONLY {
			c.JSON(http.StatusOK, datatypes.Models{Models: model})
		}
		modelList = append(modelList, model...)
	}

	if define.ENABLE_OPENAI_API {
		model := GetApiWrapper35()
		if define.ENABLE_OPENAI_API_ONLY {
			c.JSON(http.StatusOK, datatypes.Models{Models: model})
			return
		}
		modelList = append(modelList, model...)
	}

	// modelList = append(modelList, GetCustomModels()...)

	if define.ENABLE_OPENAI_OFFICIAL_MODEL {
		modelList = append(modelList, GetOfficialModels()...)
	}

	if len(modelList) == 0 {
		modelList = GetEmptyPlaceHolder()
	}

	categories := []datatypes.ModelsCategory{datatypes.MODEL_TEXT_DAVINCI_002_RENDER_SHA_CATEGORY, datatypes.MODEL_GPT4_CATEGORY, datatypes.MODEL_OTHER_CATEGORY}
	c.JSON(http.StatusOK, datatypes.Models{Models: modelList, Categories: categories})
}
