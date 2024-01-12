package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Ndraaa15/cordova/api/cholesterol/repository"
	"github.com/Ndraaa15/cordova/domain"
)

type CholesterolServiceImpl interface {
	CalculateCholesterol(c context.Context, id string, req *domain.CholesterolRequest) (*domain.Cholesterol, error)
	GetCholesterolHistory(c context.Context, id string) (*domain.Cholesterol, error)
}

type CholesterolService struct {
	cr repository.CholesterolRepositoryImpl
}

func NewCholesterolService(cholesterolRepository repository.CholesterolRepositoryImpl) CholesterolServiceImpl {
	return &CholesterolService{cholesterolRepository}
}

func (cs *CholesterolService) CalculateCholesterol(c context.Context, id string, req *domain.CholesterolRequest) (*domain.Cholesterol, error) {
	user, err := cs.cr.GetUserByID(c, id)
	if err != nil {
		return nil, err
	}

	reqPredictMap, err := parseRequestPredict(req, user)
	if err != nil {
		return nil, err
	}

	reqPredictJson, err := json.Marshal(reqPredictMap)
	if err != nil {
		return nil, err
	}

	body := bytes.NewBuffer(reqPredictJson)
	res, err := http.Post("https://cordova-model-j5ofojnjyq-as.a.run.app/predict", "application/json", body)
	if err != nil {
		return nil, err
	}

	//to-do :
	//generate recommended activity
	//saved to database

	var result map[string]interface{}
	json.NewDecoder(res.Body).Decode(&result)
	fmt.Println(result)
	//map[prediction:64]

	return nil, nil
}

func (cs *CholesterolService) GetCholesterolHistory(c context.Context, id string) (*domain.Cholesterol, error) {
	cholesterols := &domain.Cholesterol{}
	cholesterolMap := make(map[uint64][]*domain.CholesterolDB)

	cholesterol, err := cs.cr.GetCholesterolHistory(c, id)
	if err != nil {
		return nil, err
	}

	for _, value := range cholesterol {
		cholesterolMap[value.Year] = append(cholesterolMap[value.Year], value)
	}

	cholesterols.UserID = id
	cholesterols.Cholesterols = cholesterolMap

	return cholesterols, nil
}

func parseRequestPredict(req *domain.CholesterolRequest, user *domain.User) (map[string]interface{}, error) {
	date, err := time.Parse("2006-01-02", user.Birthday.Format("2006-01-02"))
	if err != nil {
		return nil, err
	}

	age := (time.Now().Year() - date.Year())

	res := strings.Split(req.BloodPressure, "/")
	if err != nil {
		return nil, err
	}

	bp_systolic, err := strconv.Atoi(res[0])
	if err != nil {
		return nil, err
	}

	bp_diastolic, err := strconv.Atoi(res[1])
	if err != nil {
		return nil, err
	}

	var obesity int
	if user.Weight/(user.Height*user.Height) > 30 {
		obesity = 1
	} else {
		obesity = 0
	}

	request := map[string]interface{}{
		"age":                             age,
		"sex":                             user.Gender,
		"cholesterol":                     req.Cholesterol,
		"heart_rate":                      req.HeartRate,
		"diabetes":                        user.Diabetes,
		"family_history":                  user.FamilyHistory,
		"smoking":                         user.Smoking,
		"obesity":                         obesity,
		"alcohol_consumption":             user.AlcoholConsumption,
		"exercise_hours_per_week":         user.Exercise,
		"diet":                            0,
		"previous_heart_problems":         user.PreviousHeartProblem,
		"medication_use":                  user.MedicationUse,
		"stress_level":                    user.StressLevel,
		"sedentary_hours_per_day":         user.SedentaryHours,
		"bmi":                             user.Weight / (user.Height * user.Height),
		"triglycerides":                   req.Triglycerides,
		"physical_activity_days_per_week": user.PhysicalActivity,
		"sleep_hours_per_day":             user.SleepHours,
		"continent":                       2,
		"hemisphere":                      2,
		"bp_systolic":                     bp_systolic,
		"bp_diastolic":                    bp_diastolic,
	}
	return request, nil
}
