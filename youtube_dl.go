package main

import "os/exec"

//https://github.com/ytdl-org/youtube-dl

// AudioQuality is Between 0 (better) and 0 (worse), default 5.
type YoutubeDL struct {
	AudioFormat  string
	AudioQuality int
}

func NewYoutubeDL() *YoutubeDL {
	return &YoutubeDL{
		AudioFormat:  "mp3",
		AudioQuality: 5,
	}
}

// LoadAudio requires ffmpeg/avconv and ffprobe/avprobe.
func (dl *YoutubeDL) LoadAudio(url string) ([]byte, error) {
	cmd := exec.Command("youtube-dl", "-x", "--audio-format", "mp3", url, "--get-url")
	return cmd.CombinedOutput()
}
