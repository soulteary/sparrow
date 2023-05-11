package models

import (
	"github.com/soulteary/sparrow/internal/datatypes"
	"github.com/soulteary/sparrow/internal/define"
)

func GetMidJourneyModel() (result []datatypes.ModelListItem) {
	model := datatypes.MODEL_MIDJOURNEY

	if define.ENABLE_I18N {
		model.Description = "目前效果最好的绘图模型之一。\n\n由同名研究实验室开发的人工智能程序，可根据文本生成图像，于2022年7月12日进入公开测试阶段，用户可透过Discord的机器人指令进行操作。\n\n该研究实验室由Leap Motion的创办人大卫·霍尔兹负责领导。"
	}

	result = append(result, model)
	return result
}
