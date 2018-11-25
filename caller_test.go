package loteria_test

import (
	"testing"

	"github.com/MarioCarrion/loteria"
)

func TestCaller_AddPlayer(t *testing.T) {
	var tests = []struct {
		name          string
		callerAction  func(c *loteria.Caller)
		expectedError bool
	}{
		{
			"OK",
			func(c *loteria.Caller) {},
			false,
		},
		{
			"Error: player already exists",
			func(c *loteria.Caller) {
				c.AddPlayer("mario")
			},
			true,
		},
		{
			"Error: game already started",
			func(c *loteria.Caller) {
				c.Announce()
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(ts *testing.T) {
			caller := loteria.NewCaller()
			tt.callerAction(&caller)

			_, err := caller.AddPlayer("mario")
			if (err != nil) != tt.expectedError {
				ts.Errorf("expected error %t, got %s", tt.expectedError, err)
			}
		})
	}
}

func TestCaller_Announce(t *testing.T) {
	caller := loteria.NewCaller()
	caller.AddPlayer("mario")

	for {
		_, err := caller.Announce()
		if err != nil {
			break
		}
	}

	_, err := caller.Announce()
	if err == nil {
		t.Fatalf("expected error, got nothing")
	}
}

func TestCaller_Loteria(t *testing.T) {
}
