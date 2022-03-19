package youtube_dl

import (
	"os/exec"
	"strconv"
)

//https://github.com/ytdl-org/youtube-dl

type Config struct {
	AudioFormat  string
	AudioQuality int
}

// AudioQuality is Between 0 (better) and 0 (worse), default 5.
type YoutubeDL struct {
	config *Config
}

func NewYoutubeDL(config *Config) *YoutubeDL {
	return &YoutubeDL{
		config: config,
	}
}

// LoadAudio requires ffmpeg/avconv and ffprobe/avprobe.
func (dl *YoutubeDL) LoadAudio(url string) ([]byte, error) {
	format := dl.config.AudioFormat
	quality := strconv.Itoa(dl.config.AudioQuality)
	cmd := exec.Command("youtube-dl", "-x", "--audio-format", format, url, "--get-url", "--audio-quality", quality)
	return cmd.CombinedOutput()
}
