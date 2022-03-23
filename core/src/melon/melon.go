package melon

import (
	"github.com/gocolly/colly"
	"github.com/youtube-dl-server/config"
	"log"
	"strings"
)

const (
	_TopURL     = "https://www.melon.com/chart/index.htm"
	_BalladeURL = "https://www.melon.com/genre/song_list.htm?gnrCode=GN0100"
	_DanceURL   = "https://www.melon.com/genre/song_list.htm?gnrCode=GN0200"
	_HiphopURL  = "https://www.melon.com/genre/song_list.htm?gnrCode=GN0300"
	_RnbURL     = "https://www.melon.com/genre/song_list.htm?gnrCode=GN0400"
	_IndieURL   = "https://www.melon.com/genre/song_list.htm?gnrCode=GN0500"
	_RockURL    = "https://www.melon.com/genre/song_list.htm?gnrCode=GN0600"
	_TrotURL    = "https://www.melon.com/genre/song_list.htm?gnrCode=GN0700"
	_FolkURL    = "https://www.melon.com/genre/song_list.htm?gnrCode=GN0800"
)

type Melon struct {
	config     *config.MelonConfig
	topURL     string
	balladeURL string
	danceURL   string
	hiphopURL  string
	rnbURL     string
	indieURL   string
	rockURL    string
	trotURL    string
	folkURL    string
}

func NewMelon(config *config.MelonConfig) *Melon {
	return &Melon{
		config:     config,
		topURL:     _TopURL,
		balladeURL: _BalladeURL,
		danceURL:   _DanceURL,
		hiphopURL:  _HiphopURL,
		rnbURL:     _RnbURL,
		indieURL:   _IndieURL,
		rockURL:    _RockURL,
		trotURL:    _TrotURL,
		folkURL:    _FolkURL,
	}
}
func (m *Melon) LoadChartList() *Chart {
	return &Chart{
		Top:     m.topList(),
		Ballade: m.balladeList(),
		Dance:   m.danceList(),
		Hiphop:  m.hiphopList(),
		Rnb:     m.rnbList(),
		Indie:   m.indieList(),
		Rock:    m.rockList(),
		Trot:    m.trotList(),
		Folk:    m.folkList(),
	}
}
func (m *Melon) topList() []*Sing {
	return parsing(m.topURL, m.config.Top)
}

func (m *Melon) balladeList() []*Sing {
	return parsing(m.balladeURL, m.config.Ballade)
}

func (m *Melon) danceList() []*Sing {
	return parsing(m.danceURL, m.config.Dance)
}
func (m *Melon) hiphopList() []*Sing {
	return parsing(m.hiphopURL, m.config.Hiphop)
}
func (m *Melon) rnbList() []*Sing {
	return parsing(m.rnbURL, m.config.Rnb)
}
func (m *Melon) indieList() []*Sing {
	return parsing(m.indieURL, m.config.Indie)
}
func (m *Melon) rockList() []*Sing {
	return parsing(m.rockURL, m.config.Rock)
}
func (m *Melon) trotList() []*Sing {
	return parsing(m.trotURL, m.config.Trot)
}
func (m *Melon) folkList() []*Sing {
	return parsing(m.balladeURL, m.config.Folk)
}

func parsing(url string, length int) []*Sing {
	var res []*Sing

	if length == 0 {
		return res
	}

	c := colly.NewCollector()
	c.OnHTML("tr", func(e *colly.HTMLElement) {
		if len(res) >= length {
			return
		}
		var tmpSing Sing
		e.ForEach("td", func(i int, td *colly.HTMLElement) {
			td.ForEach("div", func(i int, div *colly.HTMLElement) {
				rank := getRank(div)
				if rank != "" {
					tmpSing.Rank = rank
				}
				img := getHeadPhoto(div)
				if img != "" {
					tmpSing.HeadPhoto = img
				}
				titleClass := div.Attr("class")
				if strings.Contains(titleClass, "ellipsis") {
					div.ForEach("a", func(i int, element *colly.HTMLElement) {
						content := strings.TrimSpace(element.Text)
						if isTitleClass(titleClass) {
							tmpSing.Title = content
						}
						if isArtistClass(titleClass) {
							tmpSing.Artist = content
						}
						if isAlbumNameClass(titleClass) {
							tmpSing.AlbumName = content
						}
					})
				}
			})
		})
		if tmpSing.Rank != "" {
			res = append(res, &tmpSing)
		}
	})
	err := c.Visit(url)
	if err != nil {
		log.Panicln(err)
	}
	return res

}

func isContains(str string, target string) bool {
	return strings.Contains(str, target)
}
func isTitleClass(className string) bool {
	return isContains(className, "01")
}
func isArtistClass(className string) bool {
	return isContains(className, "02")
}
func isAlbumNameClass(className string) bool {
	return isContains(className, "03")
}

func getHeadPhoto(e *colly.HTMLElement) string {
	return e.ChildAttr("img", "src")
}

func getRank(e *colly.HTMLElement) string {
	return e.ChildText("span.rank")
}
