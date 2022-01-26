package pkg

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestParseGitConflictsCode(t *testing.T) {
	type args struct {
		buf string
	}
	tests := []struct {
		name    string
		args    args
		want    *ConflictsResult
		wantErr bool
	}{
		{
			name: "c1-emtpy",
			args: args{
				buf: "",
			},
			want: nil,
		},
		{
			name: "c1-one",
			args: args{
				buf: "abc",
			},
			want: &ConflictsResult{IsConflicts: false, ConflictFileA: &ConflictFile{

				Name: "",
				Data: "abc",
			}},
		},
		{
			name: "c1-two-split",
			args: args{
				buf: `abc
<<<<<<< HEAD
v1
=======
v2
>>>>>>> ABC-4993
`,
			},
			want: &ConflictsResult{IsConflicts: true,
				ConflictFileA: &ConflictFile{

					Name: "HEAD",
					Data: `abc
v1
`,
				},
				ConflictFileB: &ConflictFile{
					Name: "ABC-4993",
					Data: `abc
v2
`,
				},
			},
		},
		{
			name: "c1-two-split-2",
			args: args{
				buf: `abc
<<<<<<< HEAD
v1
=======
v2
>>>>>>> ABC-4993
<<<<<<< HEAD
v1
=======
>>>>>>> ABC-4993
`,
			},
			want: &ConflictsResult{IsConflicts: true,
				ConflictFileA: &ConflictFile{

					Name: "HEAD",
					Data: `abc
v1
v1
`,
				},
				ConflictFileB: &ConflictFile{
					Name: "ABC-4993",
					Data: `abc
v2
`,
				},
			},
		},
		{
			name: "c1-two-split-2",
			args: args{
				buf: `abc
<<<<<<< HEAD
v1
v1
v1
v1
v1
=======
v2
>>>>>>> ABC-4993
<<<<<<< HEAD
v13333
=======

>>>>>>> ABC-4993
`,
			},
			want: &ConflictsResult{IsConflicts: true,
				ConflictFileA: &ConflictFile{

					Name: "HEAD",
					Data: `abc
v1
v1
v1
v1
v1
v13333
`,
				},
				ConflictFileB: &ConflictFile{
					Name: "ABC-4993",
					Data: `abc
v2

`,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseGitConflictsCode(tt.args.buf)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseGitConflictsCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseGitConflictsCode() got = %v, want %v", got, tt.want)
				a, _ := json.MarshalIndent(tt.want, "  ", "  ")
				b, _ := json.MarshalIndent(got, "  ", "  ")
				t.Errorf("a:\n%s\n,b:\n%s\n", string(a), string(b))
			}
		})
	}
}
