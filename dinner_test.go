package main

import "testing"

func TestDinner(t *testing.T) {
	if dinner([]string{"steamed", "pork", "carrots"}) != "Enjoy your steamed pork with carrots." {
		t.Error("dinner should have been `Enjoy your steamed pork with carrots.`")
	}
}
