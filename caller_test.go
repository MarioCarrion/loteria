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

func TestCaller_Loteria(t *testing.T) { //nolint: gocyclo
	t.Run("Err: game has not started", func(ts *testing.T) {
		if caller := loteria.NewCaller(); caller.Loteria("mario") == nil {
			ts.Fatalf("expected error, got nil")
		}
	})

	t.Run("Err: game already finished", func(ts *testing.T) {
		caller := loteria.NewCaller()
		if _, err := caller.AddPlayer("mario"); err != nil {
			ts.Fatalf("expected no error, got %s", err)
		}

		for {
			if _, err := caller.Announce(); err != nil {
				break
			}
		}

		if err := caller.Loteria("mario"); err == nil {
			ts.Fatalf("expected error, got nil")
		}
	})

	t.Run("Err: player not part of the game", func(ts *testing.T) {
		caller := loteria.NewCaller()
		caller.Announce()

		if err := caller.Loteria("mario"); err == nil {
			ts.Fatalf("expected error, got nil")
		}
	})

	t.Run("Err: board is not a winner one", func(ts *testing.T) {
		caller := loteria.NewCaller()
		caller.AddPlayer("mario")
		caller.Announce()

		if err := caller.Loteria("mario"); err == nil {
			ts.Fatalf("expected error, got nil")
		}
	})

	t.Run("OK", func(ts *testing.T) {
		caller := loteria.NewCaller()
		caller.AddPlayer("mario")

		winnerFound := false
		for i := 0; i < 54; i++ {
			if _, err := caller.Announce(); err != nil {
				ts.Fatalf("expected no error got %s", err)
			}

			if err := caller.Loteria("mario"); err == nil {
				winnerFound = true
				break
			}
		}
		if !winnerFound {
			ts.Fatalf("expected to have a winner")
		}
	})
}
