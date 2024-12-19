package sqlite

import (
	"database/sql"
	"fmt"
)

type Device struct {
	ID   int
	Name string
}

type DeviceModel struct {
	DB *sql.DB
}

func (m *DeviceModel) Add(dev string) error {
	query := "INSERT INTO devices (name) VALUES (?)"
	_, err := m.DB.Exec(query, dev)
	if err != nil {
		return err
	}
	return nil
}
func (m *DeviceModel) Exists(dev string) (bool, error) {
	query := "SELECT EXISTS(SELECT 1 FROM devices WHERE name = ?)"
	var exists bool

	err := m.DB.QueryRow(query, dev).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("checking device existence: %w", err)
	}

	return exists, nil
}

func (m *DeviceModel) Delete(dev string) (bool, error) {
	exists, err := m.Exists(dev)
	if err != nil {
		return false, fmt.Errorf("checking device existence: %w", err)
	}

	if !exists {
		return false, nil
	}

	query := "DELETE FROM devices WHERE name = ?"
	_, err = m.DB.Exec(query, dev)
	if err != nil {
		return false, fmt.Errorf("deleting device: %w", err)
	}
	return true, nil
}

func (m *DeviceModel) DeleteAll() error {
	query := "Delete from devices"
	_, err := m.DB.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (m *DeviceModel) GetAll() ([]Device, error) {
	query := `SELECT id, name FROM devices`
	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
		}
	}(rows)

	var devices []Device
	for rows.Next() {
		var device Device
		err := rows.Scan(&device.ID, &device.Name)
		if err != nil {
			return nil, err
		}
		devices = append(devices, device)
	}
	return devices, nil
}

func (m *DeviceModel) Count() (int, error) {
	var count int
	query := "SELECT count(*) FROM devices"
	err := m.DB.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
