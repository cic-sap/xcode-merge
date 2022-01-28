package pkg

import (
	"github.com/cic-sap/xcode-merge/pkg/xcodeparser"
	"io/ioutil"
	"strings"
)

func Load(f string) (map[string]interface{}, error) {
	buf, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	s := string(buf)
	s = strings.TrimSpace(s)
	//fmt.Println(s)
	got, err := xcodeparser.Parse([]byte(s))
	//_ = DumpTyped(f, got)
	return got, err
}
