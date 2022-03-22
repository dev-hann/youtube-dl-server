package melon

import (
	"github.com/gocolly/colly"
	"log"
	"strings"
)

const _melonTopURL = "https://www.melon.com/chart/index.htm"

type Melon struct {
	topURL string
}

func NewMelon() *Melon {
	return &Melon{
		topURL: _melonTopURL,
	}
}

func (m *Melon) Top() *Top {
	top := &Top{
		ItemList: []*Sing{},
	}
	c := colly.NewCollector(
		colly.MaxDepth(2),
		colly.Async(true),
	)
	err := c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 2})
	if err != nil {
		log.Panicln(err)
	}
	getTitle(c, top)
	getSubTitle(c, top)
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
			top.ItemList = append(top.ItemList, &tmpSing)
		}
	})
	c.Visit(m.topURL)
	c.Wait()
	return top

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

func getTitle(c *colly.Collector, top *Top) {
	c.OnHTML("span", func(element *colly.HTMLElement) {
		if element.Attr("class") == "yyyymmdd" {
			res := ""
			res = strings.TrimSpace(element.Text)
			top.Title = res
			log.Println(res)
		}
	})
}

func getSubTitle(c *colly.Collector, top *Top) {
	c.OnHTML("span.hhmm", func(element *colly.HTMLElement) {
		res := ""
		res = strings.TrimSpace(element.Text)
		top.SubTitle = res
	})
}

func getHeadPhoto(e *colly.HTMLElement) string {
	return e.ChildAttr("img", "src")
}
func getRank(e *colly.HTMLElement) string {
	return e.ChildText("span.rank")
}
