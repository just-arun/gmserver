package postutil

import (
	s "strings"
)

func StringToLink(text string) string {
	lowe := s.ToLower(text)
	arr := s.Split(lowe, " ")
	forma := s.Join(arr, "-")
	return forma
}
