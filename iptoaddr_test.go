package ip2addr

import (
	"strconv"
	"testing"
)

/**
  连续只能查询25次
 */
func TestGet(t *testing.T) {
	for i:=0;i <26;i++ {
		resp := IpToAddr("1.41.102.1")
		t.Log(resp +"------------------"+strconv.Itoa(i))
	}
}
