package repository

const GetCholesterolHistory = `
SELECT
	user_id,
	average_cholesterol,
	last_cholesterol_record,
	cholesterol_level,
	triglycerides,
	heart_rate,
	blood_pressure,
	month,
	year,
	heart_risk_percentage,
	cholesterol_test_date,
	created_at,
	updated_at
FROM cholesterols
WHERE user_id = :user_id
ORDER BY year DESC, month DESC
`

const SavedRecordCholesterol = `
INSERT INTO cholesterols (
	user_id,
	average_cholesterol,
	last_cholesterol_record,
	cholesterol_level,
	triglycerides,
	heart_rate,
	blood_pressure,
	month,
	year,
	heart_risk_percentage,
	cholesterol_test_date
	) VALUES (
		:user_id,
		:average_cholesterol,
		:last_cholesterol_record,
		:cholesterol_level,
		:triglycerides,
		:heart_rate,
		:blood_pressure,
		:month,
		:year,
		:heart_risk_percentage,
		:cholesterol_test_date 
	)
`

const GetUserByID = `
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
