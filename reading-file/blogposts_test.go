package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "kentliuqiao.com/go-test/reading-file"
)

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tag: tdd, go
---
Hello
World!`
		secondBody = `Title: Post 2
Description: Description 2
Tag: borrow-checker, rust
---
Hello
Rust`
	)
	fs := fstest.MapFS{
		"hello world.md":  &fstest.MapFile{Data: []byte(firstBody)},
		"hello_world2.md": &fstest.MapFile{Data: []byte(secondBody)},
	}
	posts, err := blogposts.NewBlogPosts(fs)
	if err != nil {
		t.Fatal(err)
	}
	if len(posts) != len(fs) {
		t.Errorf("got %d posts, but want %d posts", len(posts), len(fs))
	}

	got := posts[0]
	want := blogposts.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
		Body: `Hello
World!`,
	}

	assertPost(t, got, want)
}

func assertPost(t *testing.T, got, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

type StubFileSystem struct{}

func (s *StubFileSystem) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, I always fail")
}
