package ip2addr

import (
	"fmt"
	"strconv"
	"testing"
)

/**
  连续只能查询25次
*/
func TestGet(t *testing.T) {
	for i := 0; i < 26; i++ {
		resp, _ := IpToAddr("1.41.102.1")
		t.Log(resp + "------------------" + strconv.Itoa(i))
	}
}

/**
  查询自己的外网IP
*/
func TestGetExternalIP(t *testing.T) {
	r, i := GetExternalIP()
	fmt.Println(r, i)
}
