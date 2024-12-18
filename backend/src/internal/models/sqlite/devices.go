package sqlite

import (
	"database/sql"
)

type Device struct {
	ID   int
	Name string
}

type DeviceModel struct {
	DB *sql.DB
}

func (m *DeviceModel) GetAll() ([]Device, error) {
	query := `SELECT id, name FROM devices`
	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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
