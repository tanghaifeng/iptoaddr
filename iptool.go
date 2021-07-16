package ip2addr

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

// 根据IP查询具体的地址 使用的是ip.tool.chinaz.com 查看的结果
func IpToAddr(ip string) (string, error) {
	resp, err := http.Get("http://ip.tool.chinaz.com/" + ip)
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}
	if doc, err := goquery.NewDocumentFromReader(resp.Body); err == nil {
		addr := doc.Find(".Whwtdhalf em")
		return addr.Text(), err
	}
	return "", err
}

//  查询自己的外网IP
func GetExternalIP() (string, string) {
	addr := ""
	ip := ""
	resp, err := http.Get("http://ip.tool.chinaz.com/")
	defer resp.Body.Close()
	if err != nil {
		return addr, ip
	}
	if doc, err := goquery.NewDocumentFromReader(resp.Body); err == nil {
		addr = doc.Find(".Whwtdhalf em").Text()
		ip = doc.Find(".fz24").Text()
	}
	return addr, ip
}
