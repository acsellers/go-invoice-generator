package generator

import "strings"

type FreeFrom struct {
	Content string
}

func FreeFromString(content string) *FreeFrom {
	return &FreeFrom{Content: content}
}

func (f *FreeFrom) ToString() string {
	return f.Content
}

func (f *FreeFrom) Lines() int {
	return len(strings.Split(f.Content, "\n"))
}
