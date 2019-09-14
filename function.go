package a

import (
	"fmt"

	c "github.com/ktateish/go-module-exp-lib-c"
)

func F(s string) string {
	t := c.F(s)
	return fmt.Sprintf("A v0.0.0: %s", t)
}
