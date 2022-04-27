package repositories

import (
	"errors"
	"github.com/jsasuga/stryd-backend-challenge/internal/data"
	"github.com/jsasuga/stryd-backend-challenge/internal/models"
	"time"
)

var errWorkoutNotFound = errors.New("workout not found")

type WorkoutRepository struct {
}

type WorkoutReceiver interface {
	FetchWorkouts() []models.Workout

	FilterWorkoutsByAthlete(athlete string) []models.Workout
	FilterWorkoutsByCoach(coach string) []models.Workout

	NewWorkout(w models.Workout) models.Workout
	UpdateWorkout(id int, w models.Workout) (models.Workout, error)
	ApproveWorkout(id int) error
	CompleteWorkout(id int) error
}

func CreateWorkoutRepository() *WorkoutRepository {
	return &WorkoutRepository{}
}

func (r *WorkoutRepository) FetchWorkouts() []models.Workout {
	return data.Workouts
}

func (r *WorkoutRepository) FilterWorkoutsByAthlete(athlete string) []models.Workout {
	var workouts []models.Workout
	for i := 0; i < len(data.Workouts); i++ {
		if athlete == data.Workouts[i].Athlete {
			workouts = append(workouts, data.Workouts[i])
		}
	}
	return workouts
}

func (r *WorkoutRepository) FilterWorkoutsByCoach(coach string) []models.Workout {
	var workouts []models.Workout
	for i := 0; i < len(data.Workouts); i++ {
		if coach == data.Workouts[i].Coach {
			workouts = append(workouts, data.Workouts[i])
		}
	}
	return workouts
}

func (r *WorkoutRepository) NewWorkout(w models.Workout) models.Workout {
	data.CurrentId = data.CurrentId + 1
	w.ID = data.CurrentId
	w.Submitted = time.Now()
	w.Approved = false
	w.Completed = false

	data.Workouts = append(data.Workouts, w)
	return w
}

func (r *WorkoutRepository) UpdateWorkout(id int, w models.Workout) (models.Workout, error) {
	index := indexOf(id)
	if index == 0 {
		return models.Workout{}, errWorkoutNotFound
	}

	data.Workouts[index].Scheduled = w.Scheduled
	data.Workouts[index].Description = w.Description
	return data.Workouts[index], nil
}

func (r *WorkoutRepository) ApproveWorkout(id int) error {
	index := indexOf(id)
	if index == 0 {
		return errWorkoutNotFound
	}

	data.Workouts[index].Approved = true
	return nil
}

func (r *WorkoutRepository) CompleteWorkout(id int) error {
	index := indexOf(id)
	if index == 0 {
		return errWorkoutNotFound
	}

	data.Workouts[index].Completed = true
	return nil
}

func indexOf(id int) int {
	for i := 0; i < len(data.Workouts); i++ {
		if id == data.Workouts[i].ID {
			return i
		}
	}
	return 0
}
