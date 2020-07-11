package sliceUtils

import "fmt"

func Compare(ma, mb interface{}) bool {
	return fmt.Sprintf("%v", ma) == fmt.Sprintf("%v", mb)
}
