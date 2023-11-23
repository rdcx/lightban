package db_test

import "testing"

func TestGetTask(t *testing.T) {
	db := testDB()

	// create a random task
	task := db.CreateRandomTask()

	res, err := db.GetTask(task.ID)

	if err != nil {
		t.Errorf("Unexpected error %v", err)
		return
	}

	if res.ID != task.ID {
		t.Errorf("Expected task ID %d, got %d", task.ID, res.ID)
	}

	if res.Name != task.Name {
		t.Errorf("Expected task name %s, got %s", task.Name, res.Name)
	}
}
