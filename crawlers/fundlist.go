package crawlers

import (
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"log"
)

/**
基金列表抓取程序
**/
import (
	"fmt"
)

var (
	fundListColly *colly.Collector
)

func init()  {
	fundListColly = colly.NewCollector(colly.Debugger(&debug.LogDebugger{}))
	fmt.Println("初始化fundListColly")

	fundListColly.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	fundListColly.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
		r.Method = "POST"
		r.Headers.Set("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
		r.Headers.Set("Cookie","__utmz=172984700.1584160250.1.1.utmcsr=baidu|utmccn=(organic)|utmcmd=organic; ASP.NET_SessionId=rixqvd55ibx0xanregx5rl45; BIGipServercn=2241222154.20480.0000; Hm_lvt_eca85e284f8b74d1200a42c9faa85464=1584160250,1584160262,1584162711,1584234483; __utmc=172984700; __utma=172984700.122964807.1584160250.1584238296.1584241687.6; __utmt=1; __utmb=172984700.6.10.1584241687; Hm_lpvt_eca85e284f8b74d1200a42c9faa85464=1584243414")
		r.Headers.Set("Referer","http://cn.morningstar.com/fundselect/default.aspx")
		r.Headers.Set("User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.131 Safari/537.36")
	})

	fundListColly.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})
	fundListColly.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))

		fmt.Println("ss")
	})

	fundListColly.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	fundListColly.OnHTML("tr .gridItem", func(element *colly.HTMLElement) {
		fmt.Println("具体项目",element.Text)
	})

	fundListColly.OnHTML("tr .gridAlternateItem", func(element *colly.HTMLElement) {
		fmt.Println("具体项目",element.Text)
	})

	fundListColly.OnHTML("dev #ctl00_cphMain_AspNetPager1", func(element *colly.HTMLElement) {
		fmt.Println("分页",element.Text)
	})
}

func FundListVisit(url string) {
	fundListColly.Visit(url)
}
