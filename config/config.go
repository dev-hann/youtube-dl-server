package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	FirebaseConfig  *FirebaseConfig `json:"firebase_config"`
	NgrokConfig     *NgrokConfig    `json:"ngrok_config"`
	YoutubeDlConfig *YoutubeConfig  `json:"youtube_dl_config"`
	ApiConfig       *ApiConfig      `json:"api_config"`
	ViewConfig      *ViewConfig     `json:"view_config"`
}

type FirebaseConfig struct {
	TokenPath string `json:"token_path"`
}

type NgrokConfig struct {
	Port  string `json:"port"`
	Token string `json:"token"`
}

type YoutubeConfig struct {
	AudioFormat  string `json:"audio_format"`
	AudioQuality int    `json:"audio_quality"`
}

type ApiConfig struct {
	Version   string `json:"version"`
	ConfigApi string `json:"config_api"`
	AudioApi  string `json:"audio_api"`
	MelonApi  string `json:"melon_api"`
}

type ViewConfig struct {
	Path string `json:"path"`
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
		FirebaseConfig: &FirebaseConfig{
			TokenPath: viper.GetString("firebase.token_path"),
		},
		NgrokConfig: &NgrokConfig{
			Token: viper.GetString("ngrok.token"),
			Port:  viper.GetString("ngrok.port"),
		},
		YoutubeDlConfig: &YoutubeConfig{
			AudioFormat:  viper.GetString("youtube_dl.audio_format"),
			AudioQuality: viper.GetInt("youtube_dl.audio_quality"),
		},
		ApiConfig: &ApiConfig{
			Version:   viper.GetString("api.version"),
			ConfigApi: viper.GetString("api.config_api"),
			AudioApi:  viper.GetString("api.audio_api"),
			MelonApi:  viper.GetString("api.melon_api"),
		},
		ViewConfig: &ViewConfig{
			Path: viper.GetString("view.path"),
		},
	}
}
