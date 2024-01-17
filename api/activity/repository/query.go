package repository

const GetAllActivity = `
SELECT
    a.id AS activity_id,
    a.user_id,
    a.activity,
    a.description,
    a.total_sub_activity,
    a.finished_sub_activity,
    a.is_done AS activity_is_done,
    a.created_at AS activity_created_at,
    a.updated_at AS activity_updated_at,
    s.id AS sub_activity_id,
    s.activity_id AS sub_activity_activity_id,
    s.sub_activity,
		s.description AS sub_activity_description,
		s.ingredients,
		s.steps,
		s.duration,
    s.is_done AS sub_activity_is_done,
		s.image,
    s.created_at AS sub_activity_created_at,
    s.updated_at AS sub_activity_updated_at
FROM
    activities a
LEFT JOIN
    sub_activities s ON a.id = s.activity_id
WHERE
		a.user_id = :user_id AND s.id IS NOT NULL
ORDER BY
    a.id, s.id;
`

const UpdateSubActivity = `
UPDATE 
	sub_activities
SET
	is_done = :is_done
WHERE 
	id = :id
`

const UpdateActivity = `
UPDATE
	activities
SET
	activity = :activity,
	description = :description,
	total_sub_activity = :total_sub_activity,
	finished_sub_activity = :finished_sub_activity,
	is_done = :is_done
WHERE
	id = :id
`

const GetSubActivityByID = `
SELECT
	id,
	activity_id,
	sub_activity,
	description,
	ingredients,
	steps,
	duration,
	is_done,
	image,
	created_at,
	updated_at
FROM
	sub_activities
WHERE
	id = :id
`

const GetActivityByID = `
SELECT
	id,
	user_id,
	activity,
	description,
	total_sub_activity,
	finished_sub_activity,
	is_done,
	created_at,
	updated_at
FROM 
	activities
WHERE 
	id = :id
`

const GetUserByID = `
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
AND 
	month = :month 
AND 
	year = :year
`

const SavedActivity = `
INSERT INTO activities (
	user_id,
	activity,
	description,
	total_sub_activity,
	finished_sub_activity,
	is_done
) VALUES (
		:user_id,
		:activity,
		:description,
		:total_sub_activity,
		:finished_sub_activity,
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
	duration,
	is_done,
	image
) VALUES (
		:activity_id,
		:sub_activity,
		:description,
		ARRAY[:ingredients],
		ARRAY[:steps],
		:duration,
		:is_done,
		:image
) RETURNING id
`

const DeleteActivity = `
DELETE FROM
	activities
WHERE
	id = :id
`
