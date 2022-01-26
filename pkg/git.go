package pkg

import (
	"io/ioutil"
	"os"
	"strings"
)

type ConflictsResult struct {
	IsConflicts   bool
	ConflictFileA *ConflictFile
	ConflictFileB *ConflictFile
}

type ConflictFile struct {
	Name string
	Data string
}

func ParseGitConflicts(f string) (*ConflictsResult, error) {
	fp, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = fp.Close()
	}()
	buf, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	return ParseGitConflictsCode(string(buf))
}

func ParseGitConflictsCode(code string) (*ConflictsResult, error) {
	if code == "" {
		return nil, nil
	}
	ret := &ConflictsResult{}
	flagStart := "<<<<<<<"
	flagSplit := "======="
	flagEnd := ">>>>>>>"
	// not found
	if strings.Index(code, flagStart) == -1 ||
		strings.Index(code, flagSplit) == -1 ||
		strings.Index(code, flagEnd) == -1 {
		ret.ConflictFileA = &ConflictFile{Name: "", Data: code}
		return ret, nil
	}

	lines := strings.Split(code, "\n")
	var codes []string
	var codesA []string
	var codesB []string
	var nameA, nameB string
	var state = 0
	for _, line := range lines {
		switch state {

		case 0:
			if strings.HasPrefix(line, flagStart) {
				if nameA == "" {
					arr := strings.Split(strings.TrimSpace(line), " ")
					if len(arr) == 2 {
						nameA = strings.TrimSpace(arr[1])
					}
				}
				state = 1
			} else {
				codes = append(codes, line)
				codesA = append(codesA, line)
				codesB = append(codesB, line)
			}
		case 1:
			if strings.HasPrefix(line, flagSplit) {
				state = 2
			} else {
				codesA = append(codesA, line)
			}
		case 2:
			if strings.HasPrefix(line, flagEnd) {
				state = 0
				if nameB == "" {
					arr := strings.Split(strings.TrimSpace(line), " ")
					if len(arr) == 2 {
						nameB = strings.TrimSpace(arr[1])
					}
				}
			} else {
				codesB = append(codesB, line)
			}
		}
		ret.IsConflicts = true
		ret.ConflictFileA = &ConflictFile{Name: nameA, Data: strings.Join(codesA, "\n")}
		ret.ConflictFileB = &ConflictFile{Name: nameB, Data: strings.Join(codesB, "\n")}
	}
	return ret, nil
}
