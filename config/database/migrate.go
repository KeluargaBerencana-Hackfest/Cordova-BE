package database

import (
	"log"
)

const userTable = `
CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255),
    birthday DATE,
    gender BOOLEAN,
    weight DOUBLE PRECISION,
    height DOUBLE PRECISION,
    exercise DOUBLE PRECISION,
    physical_activity DOUBLE PRECISION,
    sleep_hours DOUBLE PRECISION,
    smoking BOOLEAN,
    alcohol_consumption BOOLEAN,
    sedentary_hours DOUBLE PRECISION,
    diabetes BOOLEAN,
    family_history BOOLEAN,
    previous_heart_problem BOOLEAN,
    medication_use BOOLEAN,
    photo_profile VARCHAR(255),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);`

const cholesterolTable = `
CREATE TABLE IF NOT EXISTS cholesterols (
    user_id VARCHAR(255) REFERENCES users(id),
    cholesterol DOUBLE PRECISION,
    cholesterol_level VARCHAR(255),
    triglycerides DOUBLE PRECISION,
    heart_rate DOUBLE PRECISION,
    blood_pressure VARCHAR(255),
    year BIGINT,
    month BIGINT,
    PRIMARY KEY (user_id, year, month),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);`

func (db *ClientDB) MigrateDatabase() error {
	tables := []string{userTable, cholesterolTable}
	for _, table := range tables {
		_, err := db.Exec(table)
		if err != nil {
			log.Printf("[cordova-database] failed to migrate to postgres database. Error : %v\n", err)
			return err
		}
	}
	return nil
}
