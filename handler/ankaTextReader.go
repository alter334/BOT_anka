package handler

import (
	"fmt"
	"log"
	"strconv"
	"unicode"
)

func (am *AnkaManager) AnkaReader(message string) {
	log.Printf(":::AnkaReader Start:::")
	messageRune := []rune(message)

	for i := 0; i < len(messageRune); i++ {
		fmt.Printf("%c", messageRune[i])
		if messageRune[i] == '↓' {
			log.Printf("::Anka Start::")
			num := am.ankaNumReader(messageRune[i+1:])
			log.Printf("::Anka detected:: %d", num)
		}
	}
	fmt.Printf("\n")

	log.Printf(":::AnkaReader End:::")
}

func (am *AnkaManager) ankaNumReader(ankaOrigin []rune) int {
	result := "0"
	log.Printf("::Anka Read Start::")
	for i := 0; i < 12; i++ {
		if !unicode.IsDigit(ankaOrigin[i]) {
			break
		}
		fmt.Printf("%c", ankaOrigin[i])
		result += string(ankaOrigin[i])
	}
	fmt.Printf("\n")
	log.Printf("::Anka Read End::")
	num, err := strconv.Atoi(result)

	if err != nil {
		log.Printf("AnkaNumReadError")
	}
	return num
}
