package ip2addr

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

// 根据IP查询具体的地址 使用的是ip.tool.chinaz.com 查看的结果
func IpToAddr(ip string) (string, error) {
	resp, err := http.Get("http://ip.tool.chinaz.com/" + ip)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("查询出现错误")
		return "", err
	}
	if doc, err := goquery.NewDocumentFromReader(resp.Body); err == nil {
		addr := doc.Find(".Whwtdhalf em")
		return addr.Text(), err
	}
	return "", err
}
