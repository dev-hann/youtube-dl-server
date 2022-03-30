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

	topList, err := m.trotList()
	if err != nil {
		return nil, err
	}
	balladeList, err := m.balladeList()
	if err != nil {
		return nil, err
	}
	danceList, err := m.danceList()
	if err != nil {
		return nil, err
	}
	hiphopList, err := m.hiphopList()
	if err != nil {
		return nil, err
	}
	rnbList, err := m.rnbList()
	if err != nil {
		return nil, err
	}
	indieList, err := m.indieList()
	if err != nil {
		return nil, err
	}
	rockList, err := m.rockList()
	if err != nil {
		return nil, err
	}
	trotList, err := m.trotList()
	if err != nil {
		return nil, err
	}
	folkList, err := m.folkList()
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
func (m *Melon) topList() ([]*Sing, error) {
	return parsing(m.topURL, m.config.Top)
}

func (m *Melon) balladeList() ([]*Sing, error) {
	return parsing(m.balladeURL, m.config.Ballade)
}

func (m *Melon) danceList() ([]*Sing, error) {
	return parsing(m.danceURL, m.config.Dance)
}
func (m *Melon) hiphopList() ([]*Sing, error) {
	return parsing(m.hiphopURL, m.config.Hiphop)
}
func (m *Melon) rnbList() ([]*Sing, error) {
	return parsing(m.rnbURL, m.config.Rnb)
}
func (m *Melon) indieList() ([]*Sing, error) {
	return parsing(m.indieURL, m.config.Indie)
}
func (m *Melon) rockList() ([]*Sing, error) {
	return parsing(m.rockURL, m.config.Rock)
}
func (m *Melon) trotList() ([]*Sing, error) {
	return parsing(m.trotURL, m.config.Trot)
}
func (m *Melon) folkList() ([]*Sing, error) {
	return parsing(m.balladeURL, m.config.Folk)
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
