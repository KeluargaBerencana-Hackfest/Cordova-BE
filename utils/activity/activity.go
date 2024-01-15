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
		Description:  "...",
		Image:        "...",
	}

	cycling := domain.Activity{
		NameActivity: "Cycling",
		Description:  "...",
		Image:        "...",
	}

	healtyFood := domain.Activity{
		NameActivity: "Healthy Food",
		Description:  "...",
		Image:        "...",
	}

	activities := make([]*domain.Activity, 0)
	if isHealhtyFood {
		activities = append(activities, &healtyFood)
		totalGeneraate = totalGeneraate - 1
	}

	activityMap := map[string][]domain.Activity{
		"Good":    {jogging, cycling},
		"Warning": {jogging, cycling},
		"Danger":  {jogging, cycling},
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
				NameSubActivity: "Jogging Day",
				IsSequential:    true,
				Description:     "-",
				Ingredients:     []string{""},
				Steps:           []string{""},
				Count:           5,
			},
		},
		"Healthy Food": {
			{
				NameSubActivity: "Eat Avocado",
				Description:     "...",
				Ingredients:     []string{"Avocado", "Salt", "Pepper", "Lemon", "Bread"},
				Steps:           []string{"Eat", "Eat", "Eat", "Eat", "Eat"},
				IsSequential:    false,
				Count:           1,
			},
		},
		"Cycling": {
			{
				NameSubActivity: "Cycling Day",
				IsSequential:    true,
				Description:     "-",
				Ingredients:     []string{""},
				Steps:           []string{""},
				Count:           5,
			},
		},
	}
	return keyActivity[activity][GenerateRandomNumber(len(keyActivity[activity]))]
}

func GenerateRandomNumber(len int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(len)
}
