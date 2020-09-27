// Package modl implements encoding and decoding of MODL as defined in modl.uk.
package modl

import "github.com/bign8/modl.go/internal/fs"

// Unmarshal parses the MODL-encoded data and stores the result
// in the value pointed to by v.
func Unmarshal(data []byte, v interface{}, files fs.FS) error {
	return nil // TODO: everything :cry:
}
