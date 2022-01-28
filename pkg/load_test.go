package pkg

import (
	"testing"
)

func TestLoad(t *testing.T) {
	f := "xcodeparser/test/example2--project.pbxproj"
	data, err := Load(f)

	if err != nil {
		t.Fatal(err)
	}
	err = save(f+".dump.pbxproj", data)
	if err != nil {
		t.Fatal(err)
	}
}
