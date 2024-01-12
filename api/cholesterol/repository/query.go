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
