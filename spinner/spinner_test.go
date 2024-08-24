package spinner_test

import (
	"testing"
	"time"

	"github.com/ayato-san/ui-kit/spinner"
)

func TestSpinner(t *testing.T) {
	d := spinner.Init(spinner.Options{Text: "loading ...", EndText: "done"})
	go func() {
		t.Log("start")
		time.Sleep(1 * time.Second)
		d.Stop()
		t.Log("stop")
	}()
	if err := d.Start(); err != nil {
		t.Error(err)
	}
}
