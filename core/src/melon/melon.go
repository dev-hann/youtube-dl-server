package melon

import (
	"github.com/gocolly/colly"
	"github.com/youtube-dl-server/config"
	"strings"
)

const (
	_TopURL     = "https://www.melon.com/chart/index.htm"
	_genre      = "https://www.melon.com/genre/song_list.htm?gnrCode=GN0"
	_BalladeURL = _genre + "100"
	_DanceURL   = _genre + "200"
	_HiphopURL  = _genre + "300"
	_RnbURL     = _genre + "400"
	_IndieURL   = _genre + "500"
	_RockURL    = _genre + "600"
	_TrotURL    = _genre + "700"
	_FolkURL    = _genre + "800"
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
func (m *Melon) LoadChartList() (*Chart, error) {

	topList, err := parsing(m.topURL, m.config.Top)
	if err != nil {
		return nil, err
	}
	balladeList, err := parsing(m.balladeURL, m.config.Ballade)
	if err != nil {
		return nil, err
	}
	danceList, err := parsing(m.danceURL, m.config.Dance)
	if err != nil {
		return nil, err
	}
	hiphopList, err := parsing(m.hiphopURL, m.config.Hiphop)
	if err != nil {
		return nil, err
	}
	rnbList, err := parsing(m.rnbURL, m.config.Rnb)
	if err != nil {
		return nil, err
	}
	indieList, err := parsing(m.indieURL, m.config.Indie)
	if err != nil {
		return nil, err
	}
	rockList, err := parsing(m.rockURL, m.config.Rock)
	if err != nil {
		return nil, err
	}
	trotList, err := parsing(m.trotURL, m.config.Trot)
	if err != nil {
		return nil, err
	}
	folkList, err := parsing(m.balladeURL, m.config.Folk)
	if err != nil {
		return nil, err
	}

	return &Chart{
		Top:     topList,
		Ballade: balladeList,
		Dance:   danceList,
		Hiphop:  hiphopList,
		Rnb:     rnbList,
		Indie:   indieList,
		Rock:    rockList,
		Trot:    trotList,
		Folk:    folkList,
	}, nil
}

func parsing(url string, length int) ([]*Sing, error) {
	var res []*Sing

	if length == 0 {
		return res, nil
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
		return nil, err
	}
	return res, nil

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
