package cbErrors

import (
	"fmt"
	"testing"
)

func TestPretty(t *testing.T) {
	cerr := "errType | {\"key1\":\"value1\",\"key2\":{\"key21\":12982398,\"key22\":[]}}"
	exerr := "errType\n{\n    \"key1\": \"value1\",\n    \"key2\": {\n        \"key21\": 12982398,\n        \"key22\": []\n    }\n}"
	parsed, err := Pretty(fmt.Errorf(cerr))
	if err != nil {
		t.Errorf("unexpected error : %s", err.Error())
		return
	}

	if parsed != exerr {
		t.Errorf("unexpected output : expected %s, got %s", exerr, parsed)
		return
	}
}
