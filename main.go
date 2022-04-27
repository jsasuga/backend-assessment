package main

import (
	"github.com/jsasuga/stryd-backend-challenge/internal/data"
	"github.com/jsasuga/stryd-backend-challenge/internal/handlers"
	"github.com/jsasuga/stryd-backend-challenge/internal/repositories"
	"github.com/jsasuga/stryd-backend-challenge/internal/services"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	workoutRepository := repositories.CreateWorkoutRepository()
	workoutService := services.WorkoutService{
		WorkoutRepository: workoutRepository,
	}
	workoutHandler := handlers.CreateWorkoutHandler(&workoutService)

	if err := data.LoadWorkouts(); err != nil {
		e.Logger.Fatal(err)
		return
	}

	// Routes
	e.GET("/", helloWorld)

	e.GET("/workouts", workoutHandler.GetWorkouts)
	e.POST("/workouts", workoutHandler.RequestWorkout)
	e.PUT("/workouts/:id", workoutHandler.UpdateWorkout)
	e.PUT("/workouts/:id/approve", workoutHandler.ApproveWorkout)
	e.PUT("/workouts/:id/complete", workoutHandler.CompleteWorkout)
	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

func helloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
