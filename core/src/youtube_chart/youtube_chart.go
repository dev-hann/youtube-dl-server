package youtube_chart

import (
	"github.com/gocolly/colly"
	"github.com/youtube-dl-server/config"
	"github.com/youtube-dl-server/core/src/chart"
	"log"
	"strings"
)

const (
	_TopURL = "https://charts.youtube.com/charts/TopSongs/kr?hl=ko"
)

type YoutubeChart struct {
	config *config.YoutubeChartConfig
	topURL string
}

func NewYoutubeChart(config *config.YoutubeChartConfig) *YoutubeChart {
	return &YoutubeChart{
		config: config,
		topURL: _TopURL,
	}
}

func (y *YoutubeChart) LoadYoutubeChart() *chart.Chart {
	return &chart.Chart{
		Top: y.topList(),
	}
}
func (y *YoutubeChart) topList() []*chart.Sing {
	//https://www.melon.com/chart/index.htm
	return parsing(y.topURL, 100)
}

func parsing(url string, length int) []*chart.Sing {
	var res []*chart.Sing

	//if length == 0 {
	//	return res
	//}
	c := colly.NewCollector()
	c.OnResponse(func(response *colly.Response) {
		response.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	})
	c.OnHTML("class", func(element *colly.HTMLElement) {
		log.Println(element)
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
