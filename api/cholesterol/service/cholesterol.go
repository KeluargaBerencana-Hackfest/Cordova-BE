package service

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Ndraaa15/cordova/api/cholesterol/repository"
	"github.com/Ndraaa15/cordova/domain"
	"github.com/Ndraaa15/cordova/utils/enum"
)

type CholesterolServiceImpl interface {
	CalculateCholesterol(c context.Context, userID string, req *domain.CholesterolRequest) (*domain.Cholesterol, error)
	GetCholesterolHistory(c context.Context, userID string) (*domain.Cholesterol, error)
}

type CholesterolService struct {
	cr repository.CholesterolRepositoryImpl
}

func NewCholesterolService(cholesterolRepository repository.CholesterolRepositoryImpl) CholesterolServiceImpl {
	return &CholesterolService{cholesterolRepository}
}

func (cs *CholesterolService) CalculateCholesterol(c context.Context, userID string, req *domain.CholesterolRequest) (*domain.Cholesterol, error) {
	user, err := cs.cr.GetUserByID(c, userID)
	if err != nil {
		log.Printf("[cordova-cholesterol-service] failed to get user from database. Error : %v\n", err)
		return nil, err
	}

	reqPredictMap, err := parseRequestPredict(req, user)
	if err != nil {
		log.Printf("[cordova-cholesterol-service] failed to parse request predict. Error : %v\n", err)
		return nil, err
	}

	reqPredictJson, err := json.Marshal(reqPredictMap)
	if err != nil {
		log.Printf("[cordova-cholesterol-service] failed to marshal json request predict. Error : %v\n", err)
		return nil, err
	}

	reqBody := bytes.NewBuffer(reqPredictJson)
	res, err := http.Post("https://cordova-model-j5ofojnjyq-as.a.run.app/predict", "application/json", reqBody)
	if err != nil {
		log.Printf("[cordova-cholesterol-service] failed to get prediction. Error : %v\n", err)
		return nil, err
	}

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Printf("[cordova-cholesterol-service] failed to decode json response. Error : %v\n", err)
		return nil, err
	}

	percentage := result["prediction"].(float64)
	log.Printf("[cordova-cholesterol-service] prediction percentage. Prediction : %v\n", percentage)

	//Count Cholesterol Record By User ID, Month, and Year
	parsedDate, err := time.Parse("2006-01-02", req.CholesterolTestDate)
	if err != nil {
		log.Printf("[cordova-cholesterol-service] failed to parse time cholesterol test date. Error : %v\n", err)
		return nil, err
	}

	count, err := cs.cr.CountCholesterolRecord(c, userID, int(parsedDate.Month()), parsedDate.Year())
	if err != nil {
		log.Printf("[cordova-cholesterol-service] failed to count cholesterol record. Error : %v\n", err)
		return nil, err
	}

	//Check Count and Create or Update record
	if count == 0 {
		cholesterolNew := &domain.CholesterolDB{
			UserID: userID,
		}

		cholesterolParsed, err := parseCholesterol(req, cholesterolNew)
		if err != nil {
			log.Printf("[cordova-cholesterol-service] failed to parse cholesterol record. Error : %v\n", err)
			return nil, err
		}

		cholesterolParsed.HeartRiskPercentage = percentage
		_, err = cs.cr.SavedRecordCholesterol(c, userID, cholesterolParsed)
		if err != nil {
			log.Printf("[cordova-cholesterol-service] failed to save cholesterol record. Error : %v\n", err)
			return nil, err
		}

	} else {
		cholesterolHistory, err := cs.cr.GetCholesterolHistory(c, userID)
		if err != nil {
			log.Printf("[cordova-cholesterol-service] failed to get cholesterol history. Error : %v\n", err)
			return nil, err
		}

		cholesterolParsed, err := parseCholesterol(req, cholesterolHistory[0])
		if err != nil {
			log.Printf("[cordova-cholesterol-service] failed to parse cholesterol record. Error : %v\n", err)
			return nil, err
		}

		cholesterolHistory[0].HeartRiskPercentage = percentage
		_, err = cs.cr.UpdateRecordCholesterol(c, userID, cholesterolParsed)
		if err != nil {
			log.Printf("[cordova-cholesterol-service] failed to update cholesterol record. Error : %v\n", err)
			return nil, err
		}
	}

	cholesterolHistory, err := cs.GetCholesterolHistory(c, userID)
	if err != nil {
		log.Printf("[cordova-cholesterol-service] failed to get cholesterol history. Error : %v\n", err)
		return nil, err
	}

	return cholesterolHistory, nil
}

func (cs *CholesterolService) GetCholesterolHistory(c context.Context, id string) (*domain.Cholesterol, error) {
	cholesterols := &domain.Cholesterol{}
	cholesterolMap := make(map[int][]*domain.CholesterolDB)

	cholesterol, err := cs.cr.GetCholesterolHistory(c, id)
	if err != nil {
		log.Printf("[cordova-cholesterol-service] failed to get all activity. Error : %v\n", err)
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

func getCholesterolLevel(cholesterol float64) string {
	if cholesterol < 200 {
		return enum.GoodString
	} else if cholesterol >= 200 && cholesterol <= 239 {
		return enum.WarningString
	} else {
		return enum.DangerString
	}
}

func parseCholesterol(req *domain.CholesterolRequest, cholesterol *domain.CholesterolDB) (*domain.CholesterolDB, error) {
	parsedDate, err := time.Parse("2006-01-02", req.CholesterolTestDate)
	if err != nil {
		return nil, err
	}

	if int(parsedDate.Month()) == cholesterol.Month && parsedDate.Year() == cholesterol.Year && cholesterol.AverageCholesterol > 0 {
		cholesterol.AverageCholesterol = (cholesterol.AverageCholesterol + req.Cholesterol) / 2
	} else {
		cholesterol.Month = int(parsedDate.Month())
		cholesterol.Year = parsedDate.Year()
		cholesterol.AverageCholesterol = req.Cholesterol
	}

	cholesterol.LastCholesterolRecord = req.Cholesterol
	cholesterol.CholesterolLevel = getCholesterolLevel(req.Cholesterol)
	cholesterol.Triglycerides = req.Triglycerides
	cholesterol.HeartRate = req.HeartRate
	cholesterol.BloodPressure = req.BloodPressure
	cholesterol.CholesterolTestDate = parsedDate

	return cholesterol, nil
}
