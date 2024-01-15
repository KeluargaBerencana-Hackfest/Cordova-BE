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
FROM 
	cholesterols
WHERE 
	user_id = :user_id
ORDER 
	BY year DESC, month DESC
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

const UpdateRecordCholesterol = `
UPDATE 
	cholesterols
SET
	average_cholesterol = :average_cholesterol,
	last_cholesterol_record = :last_cholesterol_record,
	cholesterol_level = :cholesterol_level,
	triglycerides = :triglycerides,
	heart_rate = :heart_rate,
	blood_pressure = :blood_pressure,
	heart_risk_percentage = :heart_risk_percentage,
	cholesterol_test_date = :cholesterol_test_date
WHERE 	
	user_id = :user_id 
AND 
	year = :year 
AND 
	month = :month
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
FROM 
	users 
WHERE 
	id = :id
`

const CountCholesterolRecord = `
SELECT 
	COUNT(user_id)
FROM 
	cholesterols
WHERE 
	user_id = :user_id 
AND 
	year = :year 
AND 
	month = :month
`

const SavedActivity = `
INSERT INTO activities (
	user_id,
	activity,
	description,
	total_sub_activity,
	finished_sub_activity,
	image,
	is_done
) VALUES (
		:user_id,
		:activity,
		:description,
		:total_sub_activity,
		:finished_sub_activity,
		:image,
		:is_done
) RETURNING id
`

const SavedSubActivity = `
INSERT INTO sub_activities (
	activity_id,
	sub_activity,
	description,
	ingredients,
	steps,
	is_done
) VALUES (
		:activity_id,
		:sub_activity,
		:description,
		ARRAY[:ingredients],
		ARRAY[:steps],
		:is_done
) RETURNING id
`

const GetAllActivity = `
SELECT
    a.id AS activity_id,
    a.user_id,
    a.activity,
    a.description,
    a.total_sub_activity,
    a.finished_sub_activity,
    a.image,
    a.is_done AS activity_is_done,
    a.created_at AS activity_created_at,
    a.updated_at AS activity_updated_at,
    s.id AS sub_activity_id,
    s.activity_id AS sub_activity_activity_id,
    s.sub_activity,
		s.description AS sub_activity_description,
		s.ingredients,
		s.steps,
    s.is_done AS sub_activity_is_done,
    s.created_at AS sub_activity_created_at,
    s.updated_at AS sub_activity_updated_at
FROM
    activities a
LEFT JOIN
    sub_activities s ON a.id = s.activity_id
WHERE
		a.user_id = :user_id
ORDER BY
    a.id, s.id;
`
