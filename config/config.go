package config

import (
	"github.com/spf13/viper"
	"github.com/youtube-dl-server/api"
	"github.com/youtube-dl-server/ngrok"
	"github.com/youtube-dl-server/view"
	"github.com/youtube-dl-server/youtube_dl"
	"log"
	"os"
)

type Config struct {
	FirebaseTokenPath string
	NgrokConfig       *ngrok.Config
	YoutubeDlConfig   *youtube_dl.Config
	ApiConfig         *api.Config
	ViewConfig        *view.Config
}

func NewConfig(path string) *Config {
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	viper.SetConfigType("yaml")

	err = viper.ReadConfig(file)
	if err != nil {
		log.Panicln(err)
	}
	return &Config{
		FirebaseTokenPath: viper.GetString("firebase_token_path"),
		NgrokConfig: &ngrok.Config{
			Token: viper.GetString("ngrok.token"),
			Port:  viper.GetString("ngrok.port"),
		},
		YoutubeDlConfig: &youtube_dl.Config{
			AudioFormat:  viper.GetString("youtube_dl.audio_format"),
			AudioQuality: viper.GetInt("youtube_dl.audio_quality"),
		},
		ApiConfig: &api.Config{
			Version:  viper.GetString("API.version"),
			AudioAPI: viper.GetString("API.audioAPI"),
		},
	}
}
