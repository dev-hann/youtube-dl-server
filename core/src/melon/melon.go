package melon

import (
	"github.com/gocolly/colly"
	"strings"
)

const _melonTop50URL = "https://www.melon.com/chart/index.htm"

type Melon struct {
	top50URL string
}

func NewMelon() *Melon {
	return &Melon{
		top50URL: _melonTop50URL,
	}
}

func (m *Melon) Top50() *Top {
	top := Top{
		Items: []*Sing{},
	}
	c := colly.NewCollector()
	top.Date = getDate(c)
	top.Time = getTime(c)
	c.OnHTML("tr", func(e *colly.HTMLElement) {
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
			top.Items = append(top.Items, &tmpSing)
		}
	})
	c.Visit(m.top50URL)
	return &top

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

func getDate(c *colly.Collector) string {
	res := ""
	c.OnHTML("span.yyyymmdd", func(element *colly.HTMLElement) {
		res = strings.TrimSpace(element.Text)
	})
	return res
}

func getTime(c *colly.Collector) string {
	res := ""
	c.OnHTML("span.hhmm", func(element *colly.HTMLElement) {
		res = strings.TrimSpace(element.Text)
	})
	return res
}

func getHeadPhoto(e *colly.HTMLElement) string {
	return e.ChildAttr("img", "src")
}
func getRank(e *colly.HTMLElement) string {
	return e.ChildText("span.rank")
}
