package youtube_dl

import (
	"github.com/youtube-dl-server/config"
	"os/exec"
	"strconv"
	"strings"
)

//https://github.com/ytdl-org/youtube-dl

// AudioQuality is Between 0 (better) and 0 (worse), default 5.
type YoutubeDL struct {
	version string
	config  *config.YoutubeConfig
}

func NewYoutubeDL(config *config.YoutubeConfig) *YoutubeDL {
	return &YoutubeDL{
		version: loadVersion(),
		config:  config,
	}
}

func loadVersion() string {
	cmd := exec.Command("youtube-dl", "--version")
	data, err := cmd.CombinedOutput()
	if err != nil {
		return "Version Error"
	}
	return strings.Trim(string(data), "\n")
}

// LoadAudio requires ffmpeg/avconv and ffprobe/avprobe.
func (dl *YoutubeDL) LoadAudio(url string) ([]byte, error) {
	format := dl.config.AudioFormat
	quality := strconv.Itoa(dl.config.AudioQuality)
	cmd := exec.Command("youtube-dl", "-x", "--audio-format", format, url, "--get-url", "--audio-quality", quality)
	return cmd.CombinedOutput()
}
