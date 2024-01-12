package repository

const SavedAccount = `
INSERT INTO users (
	id,
	name,
	email
) VALUES (
	:id,
	:name,
	:email
)
`

const GetAccountByID = `
SELECT
	id,
	name,
	email,
	birthday,
	gender,
	weight,
	height,
	exercise,
	physical_activity,
	sleep_hours,
	smoking,
	alcohol_consumption,
	sedentary_hours,
	diabetes,
	family_history,
	previous_heart_problem,
	medication_use,
	stress_level,
	photo_profile,
	created_at,
	updated_at
FROM users 
WHERE id = :id
`

const CountEmail = `
SELECT COUNT(email) 
FROM users 
WHERE email = :email
`
