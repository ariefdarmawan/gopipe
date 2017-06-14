package test

import (
	"testing"

	"github.com/ariefdarmawan/gopipe"
	"github.com/eaciit/toolkit"
)

func TestCmd(t *testing.T) {
	name := "Go Pipe"
	cmd := gopipe.NewCommand(gopipe.CommandDirect, "fmt.Println(\"Hello %s\", $name)")
	cmd.AddLib("fmt")
	result := cmd.Exec(toolkit.M{}.Set("name", name))
	if result.Status == toolkit.Status_OK {
		if toolkit.IsNilOrEmpty(result.Data) {
			t.Fatalf("Output is nil")
		}

		resultTxt := result.Data.(string)
		if resultTxt != "Hello "+name {
			t.Errorf("Output %s expect %s", resultTxt, "Hello "+name)
		}
	} else {
		t.Errorf("Unable to invoke: %s", result.Message)
	}
}
