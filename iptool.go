package ip2addr

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

// 根据IP查询具体的地址 使用的是ip.tool.chinaz.com 查看的结果
func IpToAddr(ip string) string {
	resp, err := http.Get("http://ip.tool.chinaz.com/" + ip)
	if err != nil {
		fmt.Println("查询出现错误")
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	addr := doc.Find(".Whwtdhalf em")
	fmt.Println(addr.Text())
	return addr.Text()
}
