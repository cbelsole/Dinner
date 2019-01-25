package main

import (
	"fmt"
	"strings"
	"time"
)

// discover the bug if the slice is empty
// extract the hard to understand date logic and rename it something more semantic like is first half of month
// Cannot rename flag as that would be a non passive change for your users
// Test flag can be extracted, even better date can be passed in.
// The date is not a true half since it is only to the day precision, but you are not here to fix the logic just make it more readable
// f contains flags and ingredients. split these out into 2 variables
// extract create isMeat function and extract the `if h` code
// hunting code smells like a truffle pig
// What happens if they pass multiple cooking methods. Make a comment and ask the customer about it.
// If statement at the bottom is a huge red flag since it is all repeatable code and it is mixing areas of concern. 1 set is setting the method of cooking another is setting the protein and another is setting the rest.

func dinner(f []string) interface{} {
	var h bool
	ii := make([]string, 0)
	if INTERNAL {
		h = true
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

	if f[0] == "c" {
		ii = addCheese(ii)
	}

	for _, i := range f {
		if i == "c" {
		} else {
			ii = append(ii, i)
		}
	}

	return cook(ii)
}

func addCheese(s []string) []string {
	s = append(s, "cheese")
	return s
}

func cook(f []string) interface{} {
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
