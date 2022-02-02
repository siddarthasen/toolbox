package main

import "strings"

type note struct {
	content string
}

func title(n *note) string {
	if n.content == "" {
		return "Unititled"
	}

	return strings.SplitN(n.content, "\n", 2)[0]
}
