package commands

import (
	"context"
	"time"

	"github.com/andersfylling/disgord"
	"github.com/pazuzu156/atlas"
	"github.com/sirupsen/logrus"
)

// CleanUp struct
type CleanUp struct{ Command }

// InitCleanUp provides data to help
func InitCleanUp() CleanUp {
	return CleanUp{Init(&CommandItem{
		Name:        "cleanup",
		Description: "Cleans a channel up.",
		Aliases:     []string{"c", "purge"},
		Usage:       ".cleanup <number of messages>, Requires manage messages",
		Parameters: []Parameter{
			{
				Name:        "Count",
				Description: "Count of messages you want to delete if empty will delete 1000",
				Required:    false,
			},
		},
		Admin: true,
	})}
}

// Register CleanUp
func (c CleanUp) Register() *atlas.Command {
	c.CommandInterface.Run = func(ctx atlas.Context) {
		atlas.Disgord.DeleteMessage(ctx.Atlas.Disgord, context.TODO(), ctx.Message.ChannelID, ctx.Message.ID)
		p, err := disgord.Session.GetMemberPermissions(ctx.Atlas.Disgord, context.Background(), ctx.Message.GuildID, ctx.Message.Author.ID)
		if err != nil {
			logrus.Fatal("Error can't find permissions", err)
		}
		if p&disgord.PermissionManageMessages == 0 {
			ctx.Message.Reply(ctx.Context, ctx.Atlas, "Sorry you are missing the ManageMessage Permissions")
		} else {
			if ctx.Args[0] == "" {
				for i := 0; i < 5; i++ {
					message, err := atlas.Disgord.GetMessages(ctx.Atlas.Disgord, context.TODO(), ctx.Message.ChannelID, &disgord.GetMessagesParams{Before: ctx.Message.ID, Limit: 1})
					if err != nil {
						logrus.Error("Error could not retreve messages, Channel Empty? %v", err)
					}
					atlas.Disgord.DeleteMessage(ctx.Atlas.Disgord, context.TODO(), ctx.Message.ChannelID, message[0].ID)
					time.Sleep(500 * time.Millisecond)
				}
			} /* else {
				for i := 0; i < strconv.Atoi(ctx.Args[0]); i++ {
					message, err := atlas.Disgord.GetMessages(ctx.Atlas.Disgord, context.TODO(), ctx.Message.ChannelID, &disgord.GetMessagesParams{Before: ctx.Message.ID, Limit: 1})
					if err != nil {
						logrus.Error("Error could not retreve messages, Channel Empty? %v", err)
					}
					atlas.Disgord.DeleteMessage(ctx.Atlas.Disgord, context.TODO(), ctx.Message.ChannelID, message[0].ID)
					time.Sleep(500 * time.Millisecond)
				}
			} */
		}
	}
	return c.CommandInterface
}
