package handler

import (
	"fmt"
	"log"
	"strconv"
	"unicode"

	"github.com/google/uuid"
	"github.com/traPtitech/traq-ws-bot/payload"
)

func (am *AnkaManager) AnkaReader(p *payload.MessageCreated) *AnkaViewMessage {
	log.Printf(":::AnkaReader Start:::")
	nums := []int{}
	ankaViewMessage := &AnkaViewMessage{}
	ankaViewMessage.originText = make([]string, 0)
	ankaViewMessage.openedText = make([]string, 0)
	messageRune := []rune(p.Message.Text)

	lastAnkaRuneNum := 0
	for i := 0; i < len(messageRune); i++ {
		fmt.Printf("%c", messageRune[i])
		if messageRune[i] == '↓' {
			log.Printf("::Anka Start::")
			num, length := am.ankaNumReader(messageRune[i+1:])
			if num == 0 {
				continue
			}
			newAnka := &Anka{
				id:                 uuid.New().String(),
				messageID:          p.Message.ID,
				channelID:          p.Message.ChannelID,
				messageCount:       num,
				originmessageCount: num,
				inMessageAnkaOrder: len(nums),
				viewMessage:        ankaViewMessage}
			am.AddAnka(*newAnka)
			nums = append(nums, num)
			ankaViewMessage.originText = append(ankaViewMessage.originText, string(messageRune[lastAnkaRuneNum:i]))
			ankaViewMessage.openedText = append(ankaViewMessage.openedText, string(messageRune[i:i+length+1]))
			lastAnkaRuneNum = i + length + 1
			log.Printf("::Anka detected:: %d", num)
		}
	}
	ankaViewMessage.originText = append(ankaViewMessage.originText, string(messageRune[lastAnkaRuneNum:]))
	fmt.Printf("\n")

	// origintextの中身をforを使ってログ表示
	for i, text := range ankaViewMessage.originText {
		log.Printf("OriginText[%d]: %s", i, text)
	}

	// openedtextの中身をforを使ってログ表示
	for i, text := range ankaViewMessage.openedText {
		log.Printf("OpenedText[%d]: %s", i, text)
	}

	log.Printf(":::AnkaReader End:::")
	return ankaViewMessage

}

func (am *AnkaManager) ankaNumReader(ankaOrigin []rune) (num int, length int) {
	result := "0"
	length = 0
	log.Printf("::Anka Read Start::")
	for i := 0; i < min(12, len(ankaOrigin)); i++ {
		if !unicode.IsDigit(ankaOrigin[i]) {
			break
		}
		length++
		fmt.Printf("%c", ankaOrigin[i])
		result += string(ankaOrigin[i])
	}
	fmt.Printf("\n")
	log.Printf("::Anka Read End::")
	num, err := strconv.Atoi(result)

	if err != nil {
		log.Printf("AnkaNumReadError")
	}
	return num, length
}

// 複数安価メッセージの編集用文の作成
// bool→編集対象か(安価数が3以上か)
func (vm *AnkaViewMessage) ViewMessageMaker() (string, bool) {
	var result string
	if len(vm.openedText) < 3 {
		return "", false
	}
	for i, text := range vm.openedText {
		result += vm.originText[i] + text
	}
	result += vm.originText[len(vm.originText)-1]
	return result, true
}
