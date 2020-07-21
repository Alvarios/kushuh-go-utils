package osUtils

import "os"

func ENV(key string, def string) string {
	o := os.Getenv(key)
	if o == "" && def != "" {
		o = def
	}

	return o
}