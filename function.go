package a

import (
	"fmt"

	c "github.com/ktateish/go-module-exp-lib-c"
)

//go:generate ./genvers.sh

func F(s string) string {
	t := c.F(s)
	return fmt.Sprintf("A %s: %s", Version, t)
}
