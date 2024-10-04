package handler

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/traPtitech/go-traq"
)

func (h *Handler) BotSimplePost(channelID string, content string) (messageid string) {
	q, r, err := h.bot.API().
		MessageApi.
		PostMessage(context.Background(), channelID).
		PostMessageRequest(traq.PostMessageRequest{
			Content: content,
		}).
		Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return ""
	}
	return q.Id
}

func (h *Handler) BotJoiner(channelID string) {
	_, err := h.bot.API().BotApi.LetBotJoinChannel(context.Background(), "019258ac-db2d-7dad-81f5-e206d162f4c5").
		PostBotActionJoinRequest(*traq.NewPostBotActionJoinRequest(channelID)).Execute()
	if err != nil {
		log.Println(err)
	}
	channel, _, _ := h.bot.API().ChannelApi.GetChannel(context.Background(), channelID).Execute()
	log.Println("joined:" + channel.Name)
}

func (h *Handler) BotLeaver(channelID string) {
	_, err := h.bot.API().BotApi.LetBotLeaveChannel(context.Background(), "019258ac-db2d-7dad-81f5-e206d162f4c5").
		PostBotActionLeaveRequest(*traq.NewPostBotActionLeaveRequest(channelID)).Execute()
	if err != nil {
		log.Println(err)
	}
	channel, _, _ := h.bot.API().ChannelApi.GetChannel(context.Background(), channelID).Execute()
	log.Println("left:" + channel.Name)
}
