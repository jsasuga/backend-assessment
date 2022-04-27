package data

import (
	"encoding/json"
	"github.com/jsasuga/stryd-backend-challenge/internal/models"
	"io/ioutil"
	"os"
)

var (
	Workouts  []models.Workout
	CurrentId int
)

type WorkoutDataDriver struct {
}

func LoadWorkouts() error {
	jsonFile, err := os.Open("./internal/data/workouts.json")

	if err != nil {
		return err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	if err := json.Unmarshal(byteValue, &Workouts); err != nil {
		return err
	}

	CurrentId = len(Workouts)
	return nil
}
