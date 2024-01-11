package repository

const GetCholesterolHistory = `
SELECT
	user_id,
	cholesterol,
	cholesterol_level,
	triglycerides,
	heart_rate,
	blood_pressure,
	year,
	month,
	created_at,
	updated_at
FROM cholesterols
WHERE user_id = :user_id
ORDER BY year DESC, month DESC
`
