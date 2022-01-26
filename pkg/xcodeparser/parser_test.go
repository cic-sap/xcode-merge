package xcodeparser

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

//go:generate go run golang.org/x/tools/cmd/goyacc  -v y.output -o parser.go parser.y

func TestParser233(t *testing.T) {

	s := `
{  
    aa = (
				C8540FAB20616EA200F75258 /* Debug */,
				
			);

}
`
	s = strings.TrimSpace(s)
	fmt.Println(s)
	got, err := Parse([]byte(s))
	if err != nil {
		log.Error().Err(err).Interface("err", err).Msg("get error")
		t.Fatal(err)
	}
	log.Info().Interface("object", got).Msg("get info")
	t.Log("got:", got)
}
func TestParser23(t *testing.T) {

	buf, err := ioutil.ReadFile("1.pbxproj")
	if err != nil {
		t.Fatal(err)
	}
	s := string(buf)
	s = strings.TrimSpace(s)
	fmt.Println(s)
	got, err := Parse([]byte(s))
	if err != nil {
		t.Fatal(err)
	}
	t.Log("got:", got)
	buf, err = yaml.Marshal(got)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(buf))
}

func TestParser11(t *testing.T) {

	s := `
{
}
`
	s = strings.TrimSpace(s)
	fmt.Println(s)
	got, err := Parse([]byte(s))
	if err != nil {
		t.Fatal(err)
	}
	t.Log("got:", got)
}

func TestParser3(t *testing.T) {

	s := `
// !$*UTF8*$!
{
	archiveVersion = 1;
	classes = {
	};
	objectVersion = 48;
}
`
	s = strings.TrimSpace(s)
	got, err := Parse([]byte(s))
	if err != nil {
		t.Fatal(err)
	}
	t.Log("got:", got)
}

func TestParser(t *testing.T) {
	testcases := []struct {
		input   string
		output  map[string]interface{}
		wantErr string
	}{{
		input:  `{}`,
		output: map[string]interface{}{},
	}, {
		input: `{"a": 1}`,
		output: map[string]interface{}{
			"a": float64(1),
		},
	}, {
		input: `{"a": 1, "b": ["c3333", 2]}`,
		output: map[string]interface{}{
			"a": float64(1),
			"b": []interface{}{"c3333", float64(2)},
		},
	}, {
		input: `{"a": []}`,
		output: map[string]interface{}{
			"a": []interface{}{},
		},
	}, {
		input: `{"a": [1.2]}`,
		output: map[string]interface{}{
			"a": []interface{}{float64(1.2)},
		},
	}, {
		input: `{"a": [1.2, 2.3]}`,
		output: map[string]interface{}{
			"a": []interface{}{float64(1.2), float64(2.3)},
		},
	}, {
		input: `{"a": true, "b": false, "c": null}`,
		output: map[string]interface{}{
			"a": true,
			"b": false,
			"c": nil,
		},
	}, {
		input:   `.1`,
		wantErr: `syntax error`,
	}, {
		input:   `invalid`,
		wantErr: `syntax error`,
	}}
	for _, tc := range testcases {
		got, err := Parse([]byte(tc.input))
		var gotErr string
		if err != nil {
			gotErr = err.Error()
		}
		if gotErr != tc.wantErr {
			t.Errorf(`%s err: %v, want %v`, tc.input, gotErr, tc.wantErr)
		}
		if !reflect.DeepEqual(got, tc.output) {
			t.Errorf(`%s: %#v want %#v`, tc.input, got, tc.output)
		}
	}
}
