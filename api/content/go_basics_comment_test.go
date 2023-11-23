package content_test

import (
	"lightban/api/content"
	"testing"
)

func TestGoBasicsComment(t *testing.T) {
	c := content.GetLessonContent(".", "go", "basics", "comment")

	if len(c.Files) != 2 {
		t.Errorf("Expected 2 files, got %d", len(c.Files))
	}

	for _, f := range c.Files {
		if f.Name == "" {
			t.Errorf("Expected file content to be non-empty")
		}
	}

	foundMain := false
	foundGoMod := false
	for _, f := range c.Files {
		if f.Name == "main.go" {
			foundMain = true
		}

		if f.Name == "go.mod" {
			foundGoMod = true
		}
	}

	if !foundMain {
		t.Errorf("Expected to find main.go")
	}

	if !foundGoMod {
		t.Errorf("Expected to find go.mod")
	}

	if len(c.Dialogs) != 5 {
		t.Errorf("Expected 1 dialog, got %d", len(c.Dialogs))
	}

	if c.Dialogs[0].KeyName != "go_basics_comment_dialog_intro" {
		t.Errorf("Expected dialog to be 'go_basics_comment_dialog_intro', got %s", c.Dialogs[0])
	}

	if c.Dialogs[1].KeyName != "go_basics_comment_dialog_completed" {
		t.Errorf("Expected dialog to be 'go_basics_comment_dialog_completed', got %s", c.Dialogs[1])
	}

	if len(c.Asserts) != 1 {
		t.Errorf("Expected 1 asserts, got %d", len(c.Asserts))
	}

	if c.Asserts[0] != "John" {
		t.Errorf("Expected asserts to be 'John', got %s", c.Asserts[0])
	}
}
