package activity

import (
	"math/rand"
	"time"

	"github.com/Ndraaa15/cordova/domain"
)

func GenerateRecommendedActivity(cholesterolLevel int, totalGenereate int, isHealthyFood bool) []*domain.Activity {
	var activities []*domain.Activity

	if cholesterolLevel < 200 {
		activities = listActivity("Good", totalGenereate, isHealthyFood)
	} else if cholesterolLevel >= 200 && cholesterolLevel <= 239 {
		activities = listActivity("Warning", totalGenereate, isHealthyFood)
	} else if cholesterolLevel > 239 {
		activities = listActivity("Danger", totalGenereate, isHealthyFood)
	}

	for _, activity := range activities {
		activity.SubActivities = listSubActivity(activity.NameActivity)
	}

	return activities
}

func listActivity(level string, totalGeneraate int, isHealhtyFood bool) []*domain.Activity {
	jogging := domain.Activity{
		NameActivity: "Jogging",
		Description:  "Regular jogging can help reduce the amount of LDL cholesterol in a person's blood. Jogging is also a great exercise for managing weight, which can further help lower cholesterol levels.",
	}

	cycling := domain.Activity{
		NameActivity: "Cycling",
		Description:  "According to a review of 300 studies, indoor cycling has a positive effect on total cholesterol. Cycling is also a great exercise for managing weight, which can further help lower cholesterol levels.",
	}

	healtyFood := domain.Activity{
		NameActivity: "Healthy Food",
		Description:  "Healthy Food",
	}

	walking := domain.Activity{
		NameActivity: "Walking",
		Description:  "Walking has been scientifically proven to have positive effects on cholesterol levels, promoting an increase in good cholesterol (HDL) and a decrease in bad cholesterol (LDL)",
	}

	swimming := domain.Activity{
		NameActivity: "Swimming",
		Description:  "Swimming is an aerobic exercise that can help reduce total cholesterol and increase good cholesterol (HDL) levels. Swimming for 30 minutes can increase the chances of reducing dangerous cholesterol levels and also raise HDL cholesterol levels. Additionally, swimming is great for overall heart health and can help lower cholesterol. It is also an effective way to reduce the risk of heart disease, as it can reduce coronary heart disease by 30 to 40 percent in women",
	}

	consultDoctor := domain.Activity{
		NameActivity: "Consultation Doctor",
		Description:  "Consultation to Doctor ⇒ Since your cholesterol level considered as (WARNING/DANGER), we advise you to have further consultation with the doctor.",
	}

	activities := make([]*domain.Activity, 0)
	if level == "Danger" {
		activities = append(activities, &consultDoctor)
	}

	if isHealhtyFood {
		activities = append(activities, &healtyFood)
		totalGeneraate = totalGeneraate - 1
	}

	activityMap := map[string][]domain.Activity{
		"Good":    {jogging, cycling, walking, swimming},
		"Warning": {jogging, cycling, walking, swimming},
		"Danger":  {jogging, cycling, walking, swimming},
	}

	choose := make(map[int]bool)
	for i := 0; i < totalGeneraate; i++ {
		randNum := GenerateRandomNumber(len(activityMap[level]))
		if choose[randNum] {
			i--
			continue
		} else {
			choose[randNum] = true
			activities = append(activities, &activityMap[level][randNum])
		}
	}
	return activities
}

func listSubActivity(activity string) domain.SubActivity {
	keyActivity := map[string][]domain.SubActivity{
		"Jogging": {
			{
				NameSubActivity: "Day",
				IsSequential:    true,
				Description:     "-",
				Ingredients: []string{
					"-",
				},
				Steps: []string{
					"-",
				},
				Duration: 45,
				Count:    2,
				Image:    "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/cordova-hackfest-2023/jogging.jpg",
			},
			{
				NameSubActivity: "Day",
				IsSequential:    true,
				Description:     "-",
				Ingredients: []string{
					"-",
				},
				Steps: []string{
					"-",
				},
				Duration: 30,
				Count:    3,
				Image:    "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/cordova-hackfest-2023/jogging%202.jpg",
			},
		},
		"Healthy Food": {
			{
				NameSubActivity: "Edamame",
				Description:     "Edamame has several benefits in lowering cholesterol levels. Edamame is rich in protein, antioxidants, and fiber that may lower circulating cholesterol levels. Research has discovered that consuming 47 grams of edamame a day can lower bad cholesterol by nearly 13 percent. Furthermore, the FDA has approved the claims that soy foods reduce cholesterol, and edamame is a good source of soy protein.",
				Ingredients: []string{
					"Edamame 47 Grams",
				},
				Steps: []string{
					"-",
				},
				IsSequential: false,
				Duration:     5,
				Count:        1,
				Image:        "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/cordova-hackfest-2023/edamame.jpg",
			},
			{
				NameSubActivity: "Mango",
				Description:     "Mangoes have several benefits in lowering cholesterol levels. They are rich in soluble dietary fiber, pectin, and vitamin C, which help to lower blood cholesterol levels, specifically low-density lipoprotein (LDL) cholesterol, a high quantity of which is believed to be a factor in causing coronary heart disease.",
				Ingredients:     []string{"1 Piece Mango"},
				Steps:           []string{"-"},
				IsSequential:    false,
				Duration:        5,
				Count:           1,
				Image:           "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/cordova-hackfest-2023/mango.jpg",
			},
			{
				NameSubActivity: "Apple",
				Description:     "Apples are rich in soluble dietary fiber, pectin, and polyphenols, which are known to help lower blood cholesterol levels, specifically low-density lipoprotein (LDL) cholesterol. Research published in the American Journal of Clinical Nutrition suggests that eating two apples per day can reduce cholesterol levels, with a decrease in total cholesterol levels between 5% and 13%.",
				Ingredients: []string{
					"2 Pieces Apple",
				},
				Steps: []string{
					"-",
				},
				IsSequential: false,
				Duration:     5,
				Count:        1,
				Image:        "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/cordova-hackfest-2023/apple.jpg",
			},
			{
				NameSubActivity: "Oatmeal",
				Description:     "Oatmeal is most effective in lowering LDL cholesterol (bad cholesterol) levels. Oatmeal is rich in beta-glucan, a viscous soluble fiber that moves through the digestive tract slowly and forms a gel-like substance, which helps lower total cholesterol and LDL cholesterol. Research suggests that consuming 40 to 100 grams of dry oats per day can lead to cholesterol-lowering effects.",
				Ingredients: []string{
					"Oatmeal 100 Grams",
				},
				Steps: []string{
					"-",
				},
				IsSequential: false,
				Duration:     5,
				Count:        1,
				Image:        "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/cordova-hackfest-2023/oatmeal.jpg",
			},
			{
				NameSubActivity: "Banana",
				Description:     "Bananas have several benefits in lowering cholesterol levels. They are rich in potassium and fiber, which can help reduce cholesterol and blood pressure. The fiber in bananas, particularly soluble fiber, can remove cholesterol from the body, contributing to improved cardiovascular health. Additionally, a study found that daily consumption of bananas significantly lowered fasting blood glucose and LDL-cholesterol/HDL-cholesterol ratio in hypercholesterolemic subjects. Furthermore, the potassium content in bananas is associated with lowering cholesterol and blood pressure, which are important factors for heart health. A study on hypercholesterolemic and type 2 diabetic subjects found that daily consumption of 250 grams of banana significantly lowered fasting blood glucose and LDL-cholesterol/HDL-cholesterol ratio.",
				Ingredients: []string{
					"3 Pieces Banana (1 piece per day)",
				},
				Steps: []string{
					"-",
				},
				IsSequential: true,
				Duration:     5,
				Count:        3,
				Image:        "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/cordova-hackfest-2023/banana.jpg",
			},
			{
				NameSubActivity: "Banana",
				Description:     "Bananas have several benefits in lowering cholesterol levels. They are rich in potassium and fiber, which can help reduce cholesterol and blood pressure. The fiber in bananas, particularly soluble fiber, can remove cholesterol from the body, contributing to improved cardiovascular health. Additionally, a study found that daily consumption of bananas significantly lowered fasting blood glucose and LDL-cholesterol/HDL-cholesterol ratio in hypercholesterolemic subjects. Furthermore, the potassium content in bananas is associated with lowering cholesterol and blood pressure, which are important factors for heart health. A study on hypercholesterolemic and type 2 diabetic subjects found that daily consumption of 250 grams of banana significantly lowered fasting blood glucose and LDL-cholesterol/HDL-cholesterol ratio.",
				Ingredients: []string{
					"6 Pieces Banana (2 piece per day)",
				},
				Steps: []string{
					"-",
				},
				IsSequential: true,
				Duration:     5,
				Count:        3,
				Image:        "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/cordova-hackfest-2023/banana%202.jpg",
			},
			{
				NameSubActivity: "Avocado",
				Description:     "Avocados contain monounsaturated fat, a type of fat that helps increase cholesterol levels (HDL) and reduce bad cholesterol levels (LDL). Based on the Instituto Mexicano del Seguro Social study, consumption of one avocado/100 gr a day for one week is effective in reducing cholesterol levels by 17%.",
				Ingredients: []string{
					"1 Piece Avocado",
				},
				Steps: []string{
					"-",
				},
				IsSequential: false,
				Duration:     5,
				Count:        1,
				Image:        "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/cordova-hackfest-2023/avocado.jpg",
			},
			{
				NameSubActivity: "Spinach soup with chicken",
				Description:     "Spinach is a leafy green vegetable that has several benefits in lowering cholesterol levels. Spinach is a great source of soluble fiber, which helps lower cholesterol levels by binding to bile acids and removing them from the body. Spinach contains antioxidants that help protect the body from oxidative stress, which can contribute to heart disease. On the other side, chicken is a lean meat with less saturated fat and dietary cholesterol compared to red meats like pork, beef, and lamb. Boiled chicken is a healthier option compared to fried chicken, which can be high in saturated fat and cholesterol.",
				Ingredients: []string{
					"2 bunches of spinach, chopped leaves",
					"150 gr chicken breast (diced)",
					"Salt, pepper and stock powder",
					"Water",
					"1 garlic clove, sliced",
					"2 shallots, sliced",
				},
				Steps: []string{
					"Bring water to a boil in a pot, add chicken, shallots, and garlic.",
					"Add salt, pepper, stock powder",
					"Add spinach and cook until vegetables are cooked (turn occasionally).",
					"Correct the flavor. Serve warm.",
				},
				IsSequential: false,
				Duration:     5,
				Count:        1,
				Image:        "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/cordova-hackfest-2023/Spinach%20soup%20with%20chicken.jpg",
			},
			{
				NameSubActivity: "Spinach soup with corn",
				Description:     "Spinach is a leafy green vegetable that has several benefits in lowering cholesterol levels. Spinach is a great source of soluble fiber, which helps lower cholesterol levels by binding to bile acids and removing them from the body. Spinach contains antioxidants that help protect the body from oxidative stress, which can contribute to heart disease. On the other side, corn is rich in soluble fiber, which can help lower cholesterol levels by binding to bile acids and removing them from the body.",
				Ingredients: []string{
					"2 bunches of fresh spinach",
					"1½ pieces of sweet corn",
					"1½ liters of water",
					"3 cloves of garlic",
					"5 cloves shallots",
					"1 bunch of basil leaves",
					"3½ tsp salt",
					"2 tsp mushroom broth",
				},
				Steps: []string{
					"Pick spinach according to taste, then wash and set aside.",
					"Cut sweet corn to taste, then set aside.",
					"Slice shallots and garlic and set aside.",
					"Shred the basil leaves, then wash and set aside.",
					"Boil the water in a pot, put the cut sweet corn into the pot. Cover the pot and let it boil and the corn cooked for about 15 minutes.",
					"Once the corn is cooked, add shallots, garlic, salt, and mushroom stock, and adjust the flavor. If you prefer a sweeter flavor, you can add sugar to taste.",
					"Once the flavor is right, add the spinach. Stir until the spinach is soft.",
					"Add basil leaves until soft",
					"Turn off the stove, remove from heat and ready to be served.",
				},
				IsSequential: false,
				Duration:     5,
				Count:        1,
				Image:        "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/cordova-hackfest-2023/Spinach%20soup%20with%20corn.jpg",
			},
		},
		"Cycling": {
			{
				NameSubActivity: "Cycling",
				IsSequential:    false,
				Description:     "One time a week",
				Ingredients: []string{
					"-",
				},
				Steps: []string{
					"-",
				},
				Duration: 60,
				Count:    1,
				Image:    "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/cordova-hackfest-2023/cycling.jpg",
			},
			{
				NameSubActivity: "Cycling",
				IsSequential:    false,
				Description:     "One time a week",
				Ingredients: []string{
					"-",
				},
				Steps: []string{
					"-",
				},
				Duration: 90,
				Count:    1,
				Image:    "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/cordova-hackfest-2023/cycling%202.jpg",
			},
			{
				NameSubActivity: "Cycling",
				IsSequential:    false,
				Description:     "One time a week",
				Ingredients: []string{
					"-",
				},
				Steps: []string{
					"-",
				},
				Duration: 120,
				Count:    1,
				Image:    "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/cordova-hackfest-2023/cycling%203.jpg",
			},
		},
		"Walking": {
			{
				NameSubActivity: "Day",
				IsSequential:    true,
				Description:     "-",
				Ingredients: []string{
					"-",
				},
				Steps: []string{
					"-",
				},
				Duration: 45,
				Count:    2,
				Image:    "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/cordova-hackfest-2023/walking.jpg",
			},
			{
				NameSubActivity: "Day",
				IsSequential:    true,
				Description:     "-",
				Ingredients: []string{
					"-",
				},
				Steps: []string{
					"-",
				},
				Duration: 30,
				Count:    3,
				Image:    "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/cordova-hackfest-2023/walking%202.jpg",
			},
		},
		"Swimming": {
			{
				NameSubActivity: "Swimming",
				IsSequential:    false,
				Description:     "One time a week",
				Ingredients: []string{
					"-",
				},
				Steps: []string{
					"-",
				},
				Duration: 30,
				Count:    1,
				Image:    "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/cordova-hackfest-2023/swimming.jpg",
			},
		},
		"Consultation Doctor": {
			{
				NameSubActivity: "Consultation Doctor",
				IsSequential:    false,
				Description:     "Since your cholesterol level considered as (WARNING/DANGER), we advise you to have further consultation with the doctor.",
				Ingredients: []string{
					"-",
				},
				Steps: []string{
					"-",
				},
				Duration: 30,
				Count:    1,
				Image:    "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/cordova-hackfest-2023/doctor.jpg",
			},
		},
	}
	return keyActivity[activity][GenerateRandomNumber(len(keyActivity[activity]))]
}

func GenerateRandomNumber(len int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(len)
}
