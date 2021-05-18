
package main

import (
	"encoding/json"
	corev2 "github.com/sensu/sensu-go/api/core/v2"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	//assert := assert.New(t)
	file, _ := ioutil.TempFile(os.TempDir(), "otel-sensu-handler-plugin-")
	defer func() {
		_ = os.Remove(file.Name())
	}()

	event := corev2.FixtureEvent("entity1", "check1")
	event.Check = nil
	event.Metrics = corev2.FixtureMetrics()
	eventJSON, _ := json.Marshal(event)
	log.Printf("%s", string(eventJSON))
	_, _ = file.WriteString(string(eventJSON))
	_, _ = file.Seek(0, 0)
	os.Stdin = file
	//requestReceived := false

	oldArgs := os.Args
	os.Setenv("LS_ACCESS_TOKEN", "foo")
	os.Args = []string{"otel-sensu-handler-plugin"}
	defer func() { os.Args = oldArgs }()
	main()
}