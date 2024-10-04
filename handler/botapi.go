package handler

import (
	"context"
	"fmt"
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
