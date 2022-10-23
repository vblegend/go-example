package source

import (
	"crypto/sha256"
	"fmt"
)

// Sum returns the md5 checksum of the ChangeSet data
func (c *ChangeSet) Sum() string {
	h := sha256.New()
	h.Write(c.Data)
	return fmt.Sprintf("%x", h.Sum(nil))
}
