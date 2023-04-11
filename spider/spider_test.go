package spider

import (
	"fmt"
	"testing"
)

func TestSpider(t *testing.T) {
	spider := NewWorld68Spider()
	url := spider.GetWebsiteUrl()
	fmt.Println(url)
}
