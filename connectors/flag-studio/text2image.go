package FlagStudio

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/soulteary/sparrow/internal/define"
)

type TextToImage struct {
	Prompt          string  `json:"prompt"`
	GuidanceScale   float64 `json:"guidance_scale"`
	Height          int     `json:"height"`
	NegativePrompts string  `json:"negative_prompts"`
	Sampler         string  `json:"sampler"`
	Seed            int     `json:"seed"`
	Steps           int     `json:"steps"`
	Style           string  `json:"style"`
	Upsample        int     `json:"upsample"`
	Width           int     `json:"width"`
}

var FS_STYLES = []string{"国画", "写实主义", "虚幻引擎", "黑白插画", "版绘", "低聚", "工业霓虹", "电影艺术", "史诗大片", "暗黑", "涂鸦", "漫画场景", "特写", "儿童画", "油画", "水彩画", "素描", "卡通画", "浮世绘", "赛博朋克", "吉卜力", "哑光", "现代中式", "相机", "CG渲染", "动漫", "霓虹游戏", "蒸汽波", "宝可梦", "火影忍者", "圣诞老人", "个人特效", "通用漫画", "Momoko", "MJ风格", "剪纸", "齐白石", "张大千", "丰子恺", "毕加索", "梵高", "塞尚", "莫奈", "马克·夏加尔", "丢勒", "米开朗基罗", "高更", "爱德华·蒙克", "托马斯·科尔", "安迪·霍尔", "新海诚", "倪传婧", "村上隆", "黄光剑", "吴冠中", "林风眠", "木内达朗", "萨雷尔", "杜拉克", "比利宾", "布拉德利", "普罗旺森", "莫比乌斯", "格里斯利", "比普", "卡尔·西松", "玛丽·布莱尔", "埃里克·卡尔", "扎哈·哈迪德", "包豪斯", "英格尔斯", "RHADS", "阿泰·盖兰", "俊西", "坎皮恩", "德尚鲍尔", "库沙特", "雷诺阿"}

func GetRandomStyle() string {
	return FS_STYLES[define.GetRandomNumber(0, len(FS_STYLES)-1)]
}

func GenerateImageByText(s string) string {
	data := TextToImage{
		Prompt:          s,
		GuidanceScale:   7.5,
		Width:           512,
		Height:          512,
		NegativePrompts: "",
		Sampler:         "ddim",
		Seed:            1024,
		Steps:           50,
		Style:           GetRandomStyle(),
		Upsample:        1,
	}

	payload, err := define.MakeJSON(data)
	if err != nil {
		return fmt.Sprintf("FlagStudio API, An error occurred while preparing to enter data: %v", err)
	}

	token, err := GetToken(define.FLAGSTUDIO_API_KEY)
	if err != nil {
		return fmt.Sprintf("FlagStudio API, An error occurred while getting the token: %v", err)
	}

	req, err := http.NewRequest("POST", API_TEXT_TO_IMAGE, strings.NewReader(payload))
	if err != nil {
		return fmt.Sprintf("FlagStudio API, An error occurred while initializing network components: %v", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("token", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Sprintf("FlagStudio API, An error occurred while sending request: %v", err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Sprintf("FlagStudio API, An error occurred while reading response: %v", err)
	}

	base64Image, err := parseTextToImage(body)
	if err != nil {
		return fmt.Sprintf("FlagStudio API, An error occurred while parsing response: %v", err)
	}

	return `![](data:image/png;base64,` + base64Image + `)`
}

type ResponseTextToImage struct {
	Code int    `json:"code"`
	Data string `json:"data"`
	Nsfw int    `json:"nsfw"`
}

// parseToken parses the token from the response body
func parseTextToImage(buf []byte) (string, error) {
	var data ResponseTextToImage
	err := json.Unmarshal(buf, &data)
	if err != nil {
		return "", err
	}
	if data.Code != 200 || data.Data == "" {
		return "", fmt.Errorf("FlagStudio API, Get Result error, Code %d", data.Code)
	}

	if data.Nsfw != 0 {
		return "", fmt.Errorf("FlagStudio API, Get Token error, Code %d\n, NSFW: %d", data.Code, data.Nsfw)
	}

	return data.Data, nil
}
