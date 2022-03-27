package youtube_chart

import (
	"encoding/json"
	"github.com/gocolly/colly"
	"github.com/youtube-dl-server/config"
	"log"
)

const (
	_ChartApi = "https://charts.youtube.com/youtubei/v1/browse?alt=json&key=AIzaSyCzEW7JUJdSql0-2V4tHUb6laYm4iAE_dM"
)

type YoutubeChart struct {
	config   *config.YoutubeChartConfig
	chartApi string
}

func NewYoutubeChart(config *config.YoutubeChartConfig) *YoutubeChart {
	return &YoutubeChart{
		config:   config,
		chartApi: _ChartApi,
	}
}

func (y *YoutubeChart) LoadYoutubeChart() *YoutubeChart {
	var res *YoutubeChart

	c := colly.NewCollector()
	c.OnRequest(func(req *colly.Request) {
		setHeader(req)
	})

	c.OnResponse(func(response *colly.Response) {
		parsingRequest(response, res)
	})

	data := NewPayload()
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		log.Panicln(err)
	}
	err = c.PostRaw(y.chartApi, payloadBytes)
	if err != nil {
		log.Panicln(err)
	}
	return res
}

func setHeader(req *colly.Request) {
	req.Headers.Set("Authority", "charts.youtube.com")
	req.Headers.Set("Sec-Ch-Ua", "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"99\", \"Google Chrome\";v=\"99\"")
	req.Headers.Set("Sec-Ch-Ua-Arch", "\"x86\"")
	req.Headers.Set("Sec-Ch-Ua-Platform-Version", "\"5.13.0\"")
	req.Headers.Set("X-Youtube-Utc-Offset", "480")
	req.Headers.Set("Sec-Ch-Ua-Full-Version-List", "\" Not A;Brand\";v=\"99.0.0.0\", \"Chromium\";v=\"99.0.4844.82\", \"Google Chrome\";v=\"99.0.4844.82\"")
	req.Headers.Set("Sec-Ch-Ua-Model", "")
	req.Headers.Set("X-Youtube-Time-Zone", "Asia/Taipei")
	req.Headers.Set("Sec-Ch-Ua-Platform", "\"Linux\"")
	req.Headers.Set("Sec-Ch-Ua-Bitness", "\"64\"")
	req.Headers.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Headers.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.82 Safari/537.36")
	req.Headers.Set("Sec-Ch-Ua-Full-Version", "\"99.0.4844.82\"")
	req.Headers.Set("X-Youtube-Client-Name", "31")
	req.Headers.Set("X-Youtube-Client-Version", "0.2")
	req.Headers.Set("Content-Type", "application/json")
	req.Headers.Set("X-Goog-Visitor-Id", "")
	req.Headers.Set("X-Youtube-Ad-Signals", "dt=1648336208865&flash=0&frm&u_tz=480&u_his=5&u_h=1080&u_w=1920&u_ah=1053&u_aw=1920&u_cd=24&bc=31&bih=980&biw=412&brdim=0%2C27%2C0%2C27%2C1920%2C27%2C1920%2C1053%2C412%2C980&vis=1&wgl=true&ca_type=image")
	req.Headers.Set("Accept", "*/*")
	req.Headers.Set("Origin", "https://charts.youtube.com")
	req.Headers.Set("Sec-Fetch-Site", "same-origin")
	req.Headers.Set("Sec-Fetch-Mode", "cors")
	req.Headers.Set("Sec-Fetch-Dest", "empty")
	req.Headers.Set("Referer", "https://charts.youtube.com/charts/TopArtists/kr")
	req.Headers.Set("Accept-Language", "en-US,en;q=0.9")
	req.Headers.Set("Cookie", "YSC=obLgQ-1K_m0; VISITOR_INFO1_LIVE=PoKHd4kZe5g; _ga=GA1.2.1805250081.1648336179; _gid=GA1.2.1425415850.1648336179; _gat=1")

}
func parsingRequest(response *colly.Response, chart *YoutubeChart) {
	var data map[string]interface{}
	err := json.Unmarshal(response.Body, &data)
	if err != nil {
		log.Panicln(err)
	}
	log.Println(data["contents"])
}
