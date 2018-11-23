package comm

import (
	"fmt"
)

func CustomeError(e error) string {
	return fmt.Sprintf("%s", e)
}
