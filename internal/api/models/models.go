package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soulteary/sparrow/internal/datatypes"
	"github.com/soulteary/sparrow/internal/define"
)

func GetModels(c *gin.Context) {
	var modelList []datatypes.ModelListItem

	if define.ENABLE_MIDJOURNEY {
		modelList = append(modelList, GetMidJourneyModel()...)
		if define.ENABLE_ONLY_MIDJOURNEY {
			c.JSON(http.StatusOK, datatypes.Models{Models: modelList})
			return
		}
	}

	if true {
		modelList = append(modelList, GetFlagStudioModel()...)
	}

	// modelList = append(modelList, GetCustomModels()...)

	if define.ENABLE_OPENAI_API {
		modelList = append(modelList, GetApiWrapper35()...)
		if define.ENABLE_OPENAI_ONLY_3_5 {
			c.JSON(http.StatusOK, datatypes.Models{Models: modelList})
			return
		}
	}

	if define.ENABLE_OPENAI_OFFICIAL_MODEL {
		modelList = append(modelList, GetOfficialModels()...)
	}

	if len(modelList) == 0 {
		modelList = GetEmptyPlaceHolder()
	}

	c.JSON(http.StatusOK, datatypes.Models{Models: modelList})
}
