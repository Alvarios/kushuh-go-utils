package slack

import "testing"

func TestEscapeMessage(t *testing.T) {
	o := EscapeMessage("apple & <bananas>")
	e := "apple &amp; &lt;bananas&gt;"

	if o != e {
		t.Errorf("unexpected escaped srring : expected %s, got %s",e , o)
	}
}
