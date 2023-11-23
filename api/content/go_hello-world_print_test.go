package content_test

import (
	"lightban/api/content"
	"testing"
)

func TestGoBasicsHelloWorld(t *testing.T) {
	c := content.GetLessonContent(".", "go", "basics", "hello-world")

	if len(c.Files) != 3 {
		t.Errorf("Expected 3 files, got %d", len(c.Files))
	}

	for _, f := range c.Files {
		if f.Name == "" {
			t.Errorf("Expected file content to be non-empty")
		}
	}

	foundMain := false
	foundGoMod := false
	foundMainAnswer := false

	for _, f := range c.Files {
		if f.Name == "main.go" {
			foundMain = true
		}

		if f.Name == "go.mod" {
			foundGoMod = true
		}

		if f.Name == "main.go.answer" {
			foundMainAnswer = true
		}
	}

	if !foundMain {
		t.Errorf("Expected to find main.go")
	}

	if !foundGoMod {
		t.Errorf("Expected to find go.mod")
	}

	if !foundMainAnswer {
		t.Errorf("Expected to find main.go.answer")
	}

	if len(c.Dialogs) != 7 {
		t.Errorf("Expected 7 dialog, got %d", len(c.Dialogs))
	}

	if c.Dialogs[0].KeyName != "go_basics_hello-world_dialog_intro" {
		t.Errorf("Expected dialog to be 'go_basics_hello-world_dialog_intro', got %s", c.Dialogs[0].KeyName)
	}

	if c.Dialogs[0].Trigger != "started" {
		t.Errorf("Expected dialog to be triggered by 'started', got %s", c.Dialogs[0].Trigger)
	}

	if len(c.Asserts) != 1 {
		t.Errorf("Expected 1 asserts, got %d", len(c.Asserts))
	}

	if c.Asserts[0] != "Hello World" {
		t.Errorf("Expected asserts to be 'Hello World', got %s", c.Asserts[0])
	}
}
