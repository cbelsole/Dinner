package main

import "testing"

func TestDinner(t *testing.T) {
	INTERNAL = true
	if dinner([]string{"steamed", "pork", "carrots"}) != "Enjoy your steamed pork with carrots." {
		t.Errorf(
			"dinner should have been `Enjoy your steamed pork with carrots.` is `%s`",
			dinner([]string{"steamed", "pork", "carrots"}),
		)
	}
}
