package services

import (
	"github.com/jsasuga/stryd-backend-challenge/internal/models"
	"github.com/jsasuga/stryd-backend-challenge/internal/repositories"
)

type Workouts interface {
	All() []models.Workout
	GetByAthlete(athlete string) []models.Workout
	GetByCoach(athlete string) []models.Workout
	Request(r models.RequestNewWorkout) models.Workout
	Update(id int, u models.UpdateWorkout) (models.Workout, error)
	Approve(id int) error
	Complete(id int) error
}

type WorkoutService struct {
	WorkoutRepository repositories.WorkoutReceiver
}

func (s *WorkoutService) All() []models.Workout {
	workouts := s.WorkoutRepository.FetchWorkouts()
	return workouts
}

func (s *WorkoutService) GetByAthlete(athlete string) []models.Workout {
	workouts := s.WorkoutRepository.FilterWorkoutsByAthlete(athlete)
	return workouts
}

func (s *WorkoutService) GetByCoach(coach string) []models.Workout {
	workouts := s.WorkoutRepository.FilterWorkoutsByCoach(coach)
	return workouts
}

func (s *WorkoutService) Request(r models.RequestNewWorkout) models.Workout {
	/*
		things to consider:
		* a workout hasn't been completed, do we allow it to request a new workout?
		* a workout hasn't been approved, do we allow it to request a new workout?
		* validate scheduled time to see if coach & athlete are free
		* validate scheduled time not in the past

		* we know both the athlete and the coach are pretty busy but since we don't have their full schedules we have to work only with the workouts that we manage
	*/

	w := models.Workout{
		Athlete:   r.Athlete,
		Coach:     r.Coach,
		Scheduled: r.Scheduled,
	}
	w = s.WorkoutRepository.NewWorkout(w)

	// todo: add notifying layer - coach
	return w
}

func (s *WorkoutService) Update(id int, u models.UpdateWorkout) (models.Workout, error) {
	/*
		things to consider:
		* a workout hasn't been completed, do we allow it to request a new workout?
		* a workout hasn't been approved, do we allow it to request a new workout?
		* validate scheduled time to see if coach & athlete are free
		* validate scheduled time not in the past

		* we know both the athlete and the coach are pretty busy but since we don't have their full schedules we have to work only with the workouts that we manage
	*/

	w := models.Workout{
		Scheduled:   u.Scheduled,
		Description: u.Description,
	}

	w, err := s.WorkoutRepository.UpdateWorkout(id, w)
	if err != nil {
		return models.Workout{}, err
	}

	// todo: add notifying layer - both
	return w, nil
}

func (s *WorkoutService) Approve(id int) error {
	if err := s.WorkoutRepository.ApproveWorkout(id); err != nil {
		return err
	}

	// todo: add notifying layer - both
	return nil
}

func (s *WorkoutService) Complete(id int) error {
	if err := s.WorkoutRepository.CompleteWorkout(id); err != nil {
		return err
	}
	return nil
}
