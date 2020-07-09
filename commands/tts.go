package commands

import (
	"context"
	"donnybrook-v2/config"

	golang_tts "github.com/leprosus/golang-tts"
	"github.com/pazuzu156/atlas"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// TTS is a Text to speech object
type TTS struct{ Command }

// InitTTS sets up TTS
func InitTTS() TTS {
	return TTS{Init(&CommandItem{
		Name:        "tts",
		Description: "Text To Speech",
		Aliases:     []string{"t", "say"},
		Usage:       ".tts <what you want it to say>",
		Parameters: []Parameter{
			{
				Name:     "phrase",
				Required: true,
			},
		},
		Admin: false,
	})}

}

// Register TTS
func (c TTS) Register() *atlas.Command {
	c.CommandInterface.Run = func(ctx atlas.Context) {

		viper.SetConfigName("config")
		viper.SetConfigType("toml")
		viper.AddConfigPath(".")

		var cfg config.Configuration

		if err := viper.ReadInConfig(); err != nil {
			logrus.Fatal("Error reading config file, %s", err)
		}
		err := viper.Unmarshal(&cfg)
		if err != nil {
			logrus.Fatal("unable to decode into struct, %v", err)
		}

		awsKeyID := cfg.Amazon.KeyID
		awsSecID := cfg.Amazon.KeySecret

		polly := golang_tts.New(awsKeyID, awsSecID)

		polly.Format(golang_tts.MP3)
		polly.Voice(golang_tts.Justin)

		// Actual command goodies
		atlas.Disgord.DeleteMessage(ctx.Atlas.Disgord, context.TODO(), ctx.Message.ChannelID, ctx.Message.ID)

	}
	return c.CommandInterface
}
