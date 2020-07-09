package main

import (
	"os"

	"donnybrook-v2/commands"
	"donnybrook-v2/config"

	"github.com/andersfylling/disgord"
	"github.com/pazuzu156/atlas"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

/* Donnybrook v2 is a rewrite of Donnybrook to use
disgord and atlas in place of discord-go. */

var log = &logrus.Logger{
	Out:       os.Stderr,
	Formatter: new(logrus.TextFormatter),
	Hooks:     make(logrus.LevelHooks),
	Level:     logrus.InfoLevel,
}

func main() {

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

	client := atlas.New(&atlas.Options{
		DisgordOptions: disgord.Config{
			BotToken: cfg.Discord.Token,
			Logger:   log,
		},
		OwnerID: cfg.Discord.Owner,
	})

	client.Use(atlas.DefaultLogger())
	client.GetPrefix = func(m *disgord.Message) string {
		return cfg.Bot.Prefix
	}

	if err := client.Init(); err != nil {
		panic(err)
	}
}

func init() {
	atlas.Use(commands.InitHelp().Register())
	atlas.Use(commands.InitTTS().Register())
	atlas.Use(commands.InitCleanUp().Register())
	// atlas.Use(commands.InitSetup().Register())
	// atlas.Use(commands.InitJoin().Register())
	// atlas.Use(commands.InitDone().Register())
	// atlas.Use(commands.InitLeave().Register())
	// atlas.Use(commands.InitReady().Register())
	// atlas.Use(commands.InitUnready().Register())
	// atlas.Use(commands.InitScatter().Register())
	// atlas.Use(commands.InitHonk().Register())
	// atlas.Use(commands.InitSwarm().Register())
	// atlas.Use(commands.InitLick().Register())
	// TODO: Figure out how to get youtube working again after update and add command
	// atlas.Use(commands.InitYT().Register())

}
