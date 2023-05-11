package StreamResponser

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"regexp"
	"time"

	"github.com/soulteary/sparrow/internal/datatypes"
	"github.com/soulteary/sparrow/internal/define"
)

func ParseConversationBody(r io.Reader) (originBody []byte, result datatypes.Conversation, err error) {
	originBody, err = io.ReadAll(r)
	if err != nil {
		return nil, result, err
	}
	result, err = PromptSerialization(originBody)
	return originBody, result, err
}

func PromptSerialization(buf []byte) (datatypes.Conversation, error) {
	var conversation datatypes.Conversation
	err := json.Unmarshal(buf, &conversation)
	if err != nil {
		return conversation, err
	}
	return conversation, nil
}

func RandomSleep() {
	var k int
	if define.DEV_MODE {
		k = 10
	} else {
		min := 40
		max := 120
		k = rand.Intn(max-min+1) + min
	}
	if k >= 110 {
		k = rand.Intn(500-300+1) + 300
	}
	time.Sleep(time.Millisecond * time.Duration(k/define.RESPONSE_SPEED))
}

func ContainMarkdownImage(str string) bool {
	var re = regexp.MustCompile(`!\[.*?\]\(.*?\)`)
	return len(re.FindAllString(str, -1)) > 0
}

func ContainMarkdownLink(str string) bool {
	var re = regexp.MustCompile(`\[.*?\]\(.*?\)`)
	return len(re.FindAllString(str, -1)) > 0
}

func MakeJSON(data any) (string, error) {
	ret, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(ret), nil
}
