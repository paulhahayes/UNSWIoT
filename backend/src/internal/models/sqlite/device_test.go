package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestDevice_Add(t *testing.T) {
	type testCase struct {
		name     string
		device   string
		expected int
		wantErr  bool
	}

	setupTest := func(t *testing.T) (*DeviceModel, func()) {
		db, err := sql.Open("sqlite3", ":memory:")
		if err != nil {
			t.Fatal(err)
		}

		// Create the devices table
		_, err = db.Exec(`
            CREATE TABLE IF NOT EXISTS devices (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                name TEXT NOT NULL UNIQUE
            )
        `)
		if err != nil {
			t.Fatal(err)
		}

		model := &DeviceModel{DB: db}

		// Return cleanup function
		return model, func() {
			db.Close()
		}
	}

	tests := []testCase{
		{
			name:     "add first device",
			device:   "device1",
			expected: 1,
			wantErr:  false,
		},
		//{
		//	name:     "add duplicate device",
		//	device:   "device1",
		//	expected: 1, // count should remain 1
		//	wantErr:  true,
		//},
		//{
		//	name:     "add empty device name",
		//	device:   "",
		//	expected: 0,
		//	wantErr:  true,
		//},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup fresh database for each test case
			model, cleanup := setupTest(t)
			defer cleanup()

			if tt.name == "add duplicate device" {
				err := model.Add("device1")
				if err != nil {
					t.Fatal(err)
				}
			}

			err := model.Add(tt.device)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
			count, err := model.Count()
			if err != nil {
				t.Fatal(err)
			}
			if count != tt.expected {
				t.Errorf("Add() got count = %d, want %d", count, tt.expected)
			}
		})
	}
}
