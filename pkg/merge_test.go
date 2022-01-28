package pkg

import (
	"errors"
	"testing"
)

func TestStartMergeDebug(t *testing.T) {
	var err error
	err = StartMerge("../../../rao/rao-ios-app/src/xcode/RAO.xcodeproj")
	if err != nil {
		t.Fatal(err)
	}

}
func TestStartMerge(t *testing.T) {
	var err error
	err = StartMerge("error path")
	if err == nil {
		t.Fatal("except error")
	}
	err = StartMerge("error path.xxx")
	t.Log(err)
	if err == nil {
		t.Fatal("except error")
	} else if !errors.Is(err, ErrorFileName) {
		t.Fatal("except error: ErrorFileName")
	}

}
