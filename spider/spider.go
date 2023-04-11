package spider

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/simplifiedchinese"
	"log"
	"net/http"
	"strings"
)

//https://www.yoyone.com/world.asp
//http://www.abcdao.com/yazhou/riben/
//http://site.nihaowang.com/

const (
	Asia         = "亚洲"
	Africa       = "非洲"
	NorthAmerica = "北美洲"
	SouthAmerica = "南美洲"
	Europe       = "欧洲"
	Oceania      = "大洋洲"
)

type Website struct {
	//大洲
	Continent string
	//国家
	Country string
	//类别
	Category string
	//网站名称
	SiteName string
	//网站地址
	SiteUrl string
}

type FetchWebsite interface {
	GetWebsiteUrl() []Website
}

type World68Spider struct {
	BaseUrl      string
	FetchDataUrl string
}

func NewWorld68Spider() *World68Spider {
	spider := &World68Spider{
		BaseUrl:      "http://www.world68.com",
		FetchDataUrl: "http://www.world68.com/country.asp",
	}
	return spider
}

func (w *World68Spider) GetWebsiteUrl() []Website {
	// Make a GET request to the target URL
	// 创建 HTTP GET 请求
	url := w.FetchDataUrl
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error creating HTTP request:", err)
	}

	// 设置 User-Agent 头部字段
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")

	// 发送 HTTP 请求并获取响应
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending HTTP request:", err)
	}
	defer resp.Body.Close()

	// Use goquery to parse the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Println(err)
	}

	// Find all links on the page and print their href attributes
	selection := doc.Find("body > div.w.content.country > div.content_all.r")

	selection.Find("dt").Each(func(i int, s *goquery.Selection) {
		bytes := []byte(s.Text())
		// 解码 GBK 编码的响应数据
		decoder := simplifiedchinese.GB18030.NewDecoder()
		decodedBody, err := decoder.Bytes(bytes)
		if err != nil {
			fmt.Println("Error decoding response data:", err)
		}
		result := string(decodedBody)
		result = strings.ReplaceAll(result, "：", "")
		log.Println(result)
	})

	fmt.Println(selection)
	return nil
}
