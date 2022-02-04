package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagSeparator         = "Tag: "
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func newPost(f io.Reader) (Post, error) {
	scanner := bufio.NewScanner(f)

	reaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	title := reaLine(titleSeparator)
	desc := reaLine(descriptionSeparator)
	tags := strings.Split(reaLine(tagSeparator), ", ")

	return Post{
		Title:       title,
		Description: desc,
		Tags:        tags,
		Body:        readBody(scanner),
	}, nil
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan()

	buff := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buff, scanner.Text())
	}

	return strings.Trim(buff.String(), "\n")
}
