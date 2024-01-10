package database

const schema = `
CREATE TABLE user (
    first_name text,
    last_name text,
    email text
);

CREATE TABLE place (
    country text,
    city text NULL,
    telcode integer
)`

func (db *ClientDB) MigrateDatabase() error {
	db.MustExec(schema)

	return nil
}
