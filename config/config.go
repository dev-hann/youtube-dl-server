package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	FirebaseConfig     *FirebaseConfig     `json:"firebase_config"`
	NgrokConfig        *NgrokConfig        `json:"ngrok_config"`
	YoutubeDlConfig    *YoutubeDlConfig    `json:"youtube_dl_config"`
	ApiConfig          *ApiConfig          `json:"api_config"`
	ViewConfig         *ViewConfig         `json:"view_config"`
	MelonConfig        *MelonConfig        `json:"melon_config"`
	LoggerConfig       *LoggerConfig       `json:"logger_config"`
	YoutubeChartConfig *YoutubeChartConfig `json:"youtube_chart_config"`
}

type FirebaseConfig struct {
	TokenPath string `json:"token_path"`
}

type NgrokConfig struct {
	Port  string `json:"port"`
	Token string `json:"token"`
}

type YoutubeDlConfig struct {
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

type MelonConfig struct {
	Top     int `json:"top"`
	Ballade int `json:"ballade"`
	Dance   int `json:"dance"`
	Hiphop  int `json:"hiphop"`
	Rnb     int `json:"rnb"`
	Indie   int `json:"indie"`
	Rock    int `json:"rock"`
	Trot    int `json:"trot"`
	Folk    int `json:"folk"`
}

type YoutubeChartConfig struct {
}

type LoggerConfig struct {
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
		YoutubeDlConfig: &YoutubeDlConfig{
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
		MelonConfig: &MelonConfig{
			Top:     viper.GetInt("melon_chart.top"),
			Ballade: viper.GetInt("melon_chart.ballade"),
			Dance:   viper.GetInt("melon_chart.dance"),
			Hiphop:  viper.GetInt("melon_chart.hiphop"),
			Rnb:     viper.GetInt("melon_chart.rnb"),
			Indie:   viper.GetInt("melon_chart.indie"),
			Rock:    viper.GetInt("melon_chart.rock"),
			Trot:    viper.GetInt("melon_chart.trot"),
			Folk:    viper.GetInt("melon_chart.folk"),
		},
		LoggerConfig: &LoggerConfig{
			Path: viper.GetString("logger.path"),
		},
	}
}
