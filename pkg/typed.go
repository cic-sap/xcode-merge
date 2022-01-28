package pkg

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"reflect"
	"strings"
)

func dumpFields(name string, value interface{}, indent int) string {

	indentSpace := strings.Repeat(" ", indent)
	t := reflect.TypeOf(value)
	if t.Name() != "" {
		return fmt.Sprintf("%s%s %s `json:\"%s\"`\n",
			indentSpace, strings.Title(name), t.Name(), name)
	}
	if obj, ok := value.(map[string]interface{}); ok {
		code := indentSpace + "struct {\n"
		for k2, v2 := range obj {
			code += indentSpace + dumpFields(k2, v2, indent+2)
		}
		code += "\n}"
		return fmt.Sprintf(" %s%s `json:\"%s\"`\n",
			strings.Title(name), code, name)
	}
	if arr, ok := value.([]interface{}); ok {
		if len(arr) > 0 {
			t := reflect.TypeOf(arr[0])
			if t.Name() != "" {
				return fmt.Sprintf("%s%s %s `json:\"%s\"`\n",
					indentSpace, strings.Title(name), "[]"+t.Name(), name)
			}
		}
		return fmt.Sprintf("%s%s %s `json:\"%s\"`\n",
			indentSpace, strings.Title(name), "[]interface{}", name)
	}
	return ""
}

// todo
func DumpTyped(f string, got map[string]interface{}) error {
	out2, err := json.MarshalIndent(got, "  ", "  ")
	out, err := yaml.Marshal(got)
	if err != nil {
		return err
	}
	dump := map[interface{}]bool{}
	code := strings.Replace(`package models
type BaseItem struct {
	UUID string  'json:"-"'
    Isa string 'json:"isa"'
}
func (item BaseItem)getIsa() string {
	return item.Isa
}

`, "'", "`", -1)

	if obj, ok := got["objects"].(map[string]interface{}); ok {

		for k, v := range obj {
			uuid := k
			_ = uuid
			if isa, ok := v.(map[string]interface{}); ok {

				if !dump[isa["isa"]] {
					fmt.Println("isa", isa["isa"])
					code += fmt.Sprintf("\ntype %s struct { \n", isa["isa"])
					code += "  BaseItem\n"
					for k2, v2 := range isa {
						if k2 == "isa" {
							continue
						}
						code += (dumpFields(k2, v2, 2))
					}
					code += ("\n}")
				}
				dump[isa["isa"]] = true
			}
		}
	} else {

		fmt.Println("get error")
	}

	_ = ioutil.WriteFile(f+".yaml", out, 0777)
	_ = ioutil.WriteFile(f+".json", out2, 0777)
	_ = ioutil.WriteFile("models/code.go", []byte(code), 0777)
	return nil
}
