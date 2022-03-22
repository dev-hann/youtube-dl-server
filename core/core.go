package core

import (
	"github.com/youtube-dl-server/config"
	"github.com/youtube-dl-server/core/src/firebase"
	"github.com/youtube-dl-server/core/src/melon"
	"github.com/youtube-dl-server/core/src/ngrok"
	"github.com/youtube-dl-server/core/src/youtube_dl"
)

type Core struct {
	config    *config.Config
	ngrok     *ngrok.Ngrok
	youtubeDl *youtube_dl.YoutubeDL
	melon     *melon.Melon
}

func InitCore(c *config.Config) *Core {
	dl := youtube_dl.NewYoutubeDL(c.YoutubeDlConfig)
	n := ngrok.NewNgrok(c.NgrokConfig)
	f := firebase.NewFirebase(c.FirebaseConfig)
	f.UpdateNgrok(n)
	m := melon.NewMelon()
	return &Core{
		config:    c,
		ngrok:     n,
		youtubeDl: dl,
		melon:     m,
	}
}

func (c *Core) LoadAudioURL(url string) ([]byte, error) {
	return c.youtubeDl.LoadAudio(url)
}

func (c *Core) LoadConfig() interface{} {
	return c.config
}

func (c *Core) LoadMelon() interface{} {
	return c.melon.Top()
}
