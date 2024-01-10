package database

import "log"

const userTable = `
DROP TABLE IF EXISTS users;
CREATE TABLE users (
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
    diet_level VARCHAR(255),
    sedentary_hours DOUBLE PRECISION,
    heart_rate DOUBLE PRECISION,
    diabetes BOOLEAN,
    triglycerides DOUBLE PRECISION,
    family_history BOOLEAN,
    previous_heart_problem BOOLEAN,
    medication_use BOOLEAN
);`

const cholesterolTable = `
DROP TABLE IF EXISTS cholesterols;
CREATE TABLE cholesterols (
    user_id VARCHAR(255) REFERENCES users(id),
    cholesterol_level DOUBLE PRECISION,
    year BIGINT,
    month VARCHAR(255),
    PRIMARY KEY (user_id, year, month)
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
