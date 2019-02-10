package main

import (
	"fmt"
	"strings"
	"time"
)

func dinner(f []string) string {
	var h bool
	ii := make([]string, 0)
	if INTERNAL {
		h = false
	} else {
		h = time.Date(time.Now().Year(), time.Now().Month()+1, 0, 0, 0, 0, 0, time.UTC).Day()/time.Now().Day() > 2
	}

	if h {
		for i, m := range f {
			if m == "chicken" || m == "beef" || m == "pork" || m == "fish" || m == "ham" {
				f[i] = "tofu"
			}
		}
	}

	if f[0] == "-c" {
		ii = addCheese(ii)
	}

	for _, i := range f {
		if i == "-c" {
			continue
		}
		ii = append(ii, i)
	}

	return cook(ii)
}

func addCheese(s []string) []string {
	s = append(s, "cheese")
	return s
}

func cook(f []string) string {
	var m, p string
	var rest []string
	for _, i := range f {
		if i == "steamed" {
			m = i
		} else if i == "fried" {
			m = i
		} else if i == "sautéed" {
			m = i
		} else if i == "chicken" {
			p = i
		} else if i == "beef" {
			p = i
		} else if i == "pork" {
			p = i
		} else if i == "fish" {
			p = i
		} else if i == "ham" {
			p = i
		} else if i == "tofu" {
			p = i
		} else {
			rest = append(rest, i)
		}
	}

	return fmt.Sprintf("Enjoy your %s %s with %s.", m, p, strings.Join(rest, " and "))
}
