package trace

import (
	"testing"
	"bytes"
)

func TestOff(t *testing.T) {
	var silentTracer = Off()
	silentTracer.Trace("something")
}

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)

	if tracer == nil {
		t.Error("Return should not be nil")
	} else {
		tracer.Trace("Hello trace package")
		if buf.String() != "Hello trace package" {
			t.Errorf("Trace should not write '%s'.", buf.String())
		}
	}
}
