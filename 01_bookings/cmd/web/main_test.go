package main

import "testing"

func TestRun(t *testing.T) {
	err := run()

	if err != nil {
		t.Errorf("error aoccured at run func : %f", err)
	}

}
