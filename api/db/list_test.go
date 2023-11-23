package db_test

import "testing"

func TestGetListsByProjectID(t *testing.T) {
	db := testDB()

	// create a random list
	list := db.CreateRandomList()

	res, err := db.GetListsByProjectID(list.ProjectID)

	if err != nil {
		t.Errorf("Unexpected error %v", err)
		return
	}

	if len(res) < 1 {
		t.Errorf("Expected at least 1 lang, got %d", len(res))
	}

	if res[0].ID != list.ID {
		t.Errorf("Expected list ID %d, got %d", list.ID, res[0].ID)
	}

	if res[0].Name != list.Name {
		t.Errorf("Expected list name %s, got %s", list.Name, res[0].Name)
	}

	if res[0].ProjectID != list.ProjectID {
		t.Errorf("Expected list project ID %d, got %d", list.ProjectID, res[0].ProjectID)
	}
}

func TestGetList(t *testing.T) {
	db := testDB()

	// create a random list
	list := db.CreateRandomList()

	res, err := db.GetList(list.ID)

	if err != nil {
		t.Errorf("Unexpected error %v", err)
		return
	}

	if res.ID != list.ID {
		t.Errorf("Expected list ID %d, got %d", list.ID, res.ID)
	}

	if res.Name != list.Name {
		t.Errorf("Expected list name %s, got %s", list.Name, res.Name)
	}

	if res.ProjectID != list.ProjectID {
		t.Errorf("Expected list project ID %d, got %d", list.ProjectID, res.ProjectID)
	}

	if res.Project.UserID != list.Project.UserID {
		t.Errorf("Expected list project user ID %d, got %d", list.Project.UserID, res.Project.UserID)
	}
}

func TestDeleteByIDList(t *testing.T) {
	db := testDB()
	list := db.CreateRandomList()

	err := db.DeleteListByID(list.ID)

	if err != nil {
		t.Errorf("Unexpected error %v", err)
		return
	}

	res, err := db.GetList(list.ID)

	if err == nil {
		t.Errorf("Expected error, got nil")
		return
	}

	if res != nil {
		t.Errorf("Expected nil, got %v", res)
	}
}

func TestCreateList(t *testing.T) {
	db := testDB()

	list := db.CreateRandomList()

	if list.ID == 0 {
		t.Errorf("Expected ID > 0, got %d", list.ID)
	}

	if list.Name == "" {
		t.Errorf("Expected name not empty, got %s", list.Name)
	}

	if list.ProjectID == 0 {
		t.Errorf("Expected project ID > 0, got %d", list.ProjectID)
	}
}
