package migration

import (
	"gorm.io/gorm"
)

type Migration struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	sql  string
}

var migrationsTable = `
CREATE TABLE IF NOT EXISTS migrations (
	id SERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
`

var migrations = []Migration{
	{
		Name: "create_users_table_01",
		sql: `
			CREATE TABLE users (
				id SERIAL PRIMARY KEY,
				username VARCHAR(255) NOT NULL,
				password VARCHAR(255) NOT NULL,
				email VARCHAR(255) NOT NULL,
				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				deleted_at TIMESTAMP,

				UNIQUE (username),
				UNIQUE (email)
			);
			`,
	},
	{
		Name: "create_projects_table_02",
		sql: `
			CREATE TABLE projects (
				id SERIAL PRIMARY KEY,
				name VARCHAR(255) NOT NULL,
				user_id BIGINT UNSIGNED NOT NULL,
				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				deleted_at TIMESTAMP,

				FOREIGN KEY (user_id) REFERENCES users (id),
				UNIQUE KEY ` + "`user_id_name`" + ` (user_id, name)
			);
			`,
	},
	{
		Name: "create_lists_table_03",
		sql: `
			CREATE TABLE lists (
				id SERIAL PRIMARY KEY,
				name VARCHAR(255) NOT NULL,
				project_id BIGINT UNSIGNED NOT NULL,
				position INT NOT NULL,
				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				deleted_at TIMESTAMP,

				FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE,
				UNIQUE KEY ` + "`project_id_name`" + ` (project_id, name)
				
			);
			`,
	},
	{
		Name: "create_tasks_table_04",
		sql: `
			CREATE TABLE tasks (
				id SERIAL PRIMARY KEY,
				name VARCHAR(255) NOT NULL,
				description TEXT,
				list_id BIGINT UNSIGNED NOT NULL,
				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				deleted_at TIMESTAMP,

				FOREIGN KEY (list_id) REFERENCES lists (id) ON DELETE CASCADE,
				UNIQUE KEY ` + "`list_id_name`" + ` (list_id, name)
			);
			`,
	},
}

func (Migration) TableName() string {
	return "migrations"
}

func Run(db *gorm.DB) error {

	// Create migrations table
	if err := db.Exec(migrationsTable).Error; err != nil {
		panic(err)
	}

	for _, migration := range migrations {

		// Check if migration already exists
		var m Migration
		db.Where("name = ?", migration.Name).First(&m)

		if m.ID != 0 {
			continue
		}

		if err := db.Exec(migration.sql).Error; err != nil {
			panic(err)
		}
		db.Create(&migration)
	}
	return nil
}
