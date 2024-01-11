package repository

const UpdateUser = `
UPDATE users
SET
	name = :name,
	email = :email,
	birthday = :birthday,
	gender = :gender,
	weight = :weight,
	height = :height,
	exercise = :exercise,
	physical_activity = :physical_activity,
	sleep_hours = :sleep_hours,
	smoking = :smoking,
	alcohol_consumption = :alcohol_consumption,
	sedentary_hours = :sedentary_hours,
	diabetes = :diabetes,
	family_history = :family_history,
	previous_heart_problem = :previous_heart_problem,
	medication_use = :medication_use,
	photo_profile = :photo_profile
WHERE id = :id
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
	photo_profile,
	created_at,
	updated_at
FROM users 
WHERE id = :id
`
