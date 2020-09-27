// Package fs is a placeholder while the fs.FS draft is worked into mainline go.
//
// More information can be found here: https://go.googlesource.com/proposal/+/master/design/draft-iofs.md
package fs

import "os"

type FS interface {
	Open(name string) (File, error)
}

type File interface {
	Stat() (os.FileInfo, error)
	Read([]byte) (int, error)
	Close() error
}
