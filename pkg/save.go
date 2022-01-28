package pkg

import (
	"bytes"
	"fmt"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"reflect"
	"strings"
)

func save(f string, data map[string]interface{}) (err error) {

	buf := bytes.NewBuffer(nil)
	buf.WriteString("// !$*UTF8*$!\n")
	buf.WriteString("{\n")

	keys := getKeys(data)
	for _, k := range keys {
		buf.WriteString(fmt.Sprintf("\t%s = %s ;\n", k, dumpData(data[k], 2, []string{k})))
	}
	buf.WriteString("}\n")
	err = ioutil.WriteFile(f, buf.Bytes(), 0644)
	return err
}

func dumpData(v interface{}, indent int, parent []string) string {
	space := strings.Repeat("\t", indent)
	t := reflect.TypeOf(v)
	if t.Name() != "" {
		//fmt.Println("dump type", t.Name())
		buf, err := yaml.Marshal(v)
		if err == nil {
			return strings.TrimSpace(string(buf))
		}
		//error
		log.Error().Err(err).Msgf("get error")
	}
	if obj, ok := v.(map[string]interface{}); ok {
		fmt.Println("parent keys:", parent)
		code := "{\n"
		// build sections /* Begin PBXBuildFile section */
		if parent[0] == "objects" && len(parent) == 1 {
			sections := make(map[string][]string)
			sectionKeys := make(map[string]interface{})
			for k2, v2 := range obj {
				if obj2, ok2 := v2.(map[string]interface{}); ok2 {

					sectionKeys[obj2["isa"].(string)] = true
					sections[obj2["isa"].(string)] = append(sections[obj2["isa"].(string)], k2)
				}
			}
			for _, sectionKey := range getKeys(sectionKeys) {
				code += fmt.Sprintf("\n/* Begin %s section */\n", sectionKey)
				for _, k3 := range sections[sectionKey] {
					v2 := obj[k3]
					code += fmt.Sprintf("%s%s = %s ;\n", space, k3, dumpData(v2, indent+1, appendKeys(parent, sectionKey)))
				}
				code += fmt.Sprintf("\n/* End %s section */\n", sectionKey)
			}
		} else {
			for k2, v2 := range obj {
				code += fmt.Sprintf("%s%s = %s ;\n", space, k2, dumpData(v2, indent+1, appendKeys(parent, k2)))
			}
		}

		code += space + "}"
		return code
	}

	if obj, ok := v.([]interface{}); ok {
		code := "(\n"
		for _, v2 := range obj {
			code += fmt.Sprintf("%s%s ,\n", space, dumpData(v2, indent+1, parent))
		}
		code += space + ")"
		return code
	}
	fmt.Println("get error")
	return ""
}
