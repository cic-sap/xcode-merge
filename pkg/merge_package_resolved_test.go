package pkg

import (
	"github.com/rs/zerolog/log"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestDoMergePackage(t *testing.T) {
	s := `
{
  "object": {
    "pins": [
      {
<<<<<<< HEAD
        "package": "swift-atomics",
        "repositoryURL": "https://github.com/apple/swift-atomics.git",
        "state": {
          "branch": null,
          "revision": "919eb1d83e02121cdb434c7bfc1f0c66ef17febe",
          "version": "1.0.2"
        }
      },
      {
        "package": "SwiftProtobuf",
        "repositoryURL": "https://github.com/apple/swift-protobuf.git",
        "state": {
          "branch": null,
          "revision": "7e2c5f3cbbeea68e004915e3a8961e20bd11d824",
          "version": "1.18.0"
=======
        "package": "AnyCodable",
        "repositoryURL": "https://github.com/Flight-School/AnyCodable",
        "state": {
          "branch": "master",
          "revision": "b1a7a8a6186f2fcb28f7bda67cfc545de48b3c80",
          "version": null
>>>>>>> anycode
        }
      }
    ]
  },
  "version": 1
}`
	f, err := ioutil.TempFile("", "Package.resolved.")
	if err != nil {
		t.Fatal(err)
	}
	_, _ = io.WriteString(f, s)
	_ = f.Close()
	defer func() {
		log.Debug().Msgf("remove:%s", f.Name())
		_ = os.Remove(f.Name())
	}()

	log.Debug().Msgf("write:%s", f.Name())
	log.Debug().Msgf("before merge:%s", s)
	err = DoMergePackage(f.Name())
	if err != nil {
		t.Fatal(err)
	}
	buf, err := ioutil.ReadFile(f.Name())
	if err != nil {
		t.Fatal(err)
	}
	log.Debug().Msgf("after merge:\n%s\n", string(buf))

}
