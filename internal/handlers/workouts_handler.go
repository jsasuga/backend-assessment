package handlers

import (
	"github.com/jsasuga/stryd-backend-challenge/internal/models"
	"github.com/jsasuga/stryd-backend-challenge/internal/services"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type WorkoutHandler struct {
	workoutService services.Workouts
}

func CreateWorkoutHandler(ws services.Workouts) *WorkoutHandler {
	return &WorkoutHandler{
		workoutService: ws,
	}
}

func (h *WorkoutHandler) GetWorkouts(c echo.Context) error {
	var workouts []models.Workout
	c.Response().Header().Set("Content-Type", "application/json")
	athlete := c.QueryParam("athlete")
	coach := c.QueryParam("coach")

	if athlete != "" {
		workouts = h.workoutService.GetByAthlete(athlete)
	} else if coach != "" {
		workouts = h.workoutService.GetByCoach(coach)
	} else {
		workouts = h.workoutService.All()
	}

	if err := c.JSON(http.StatusOK, workouts); err != nil {
		c.Logger().Error("failed on parsing response", err.Error())
		return err
	}
	return nil
}

func (h *WorkoutHandler) RequestWorkout(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "application/json")
	rw := new(models.RequestNewWorkout)
	if err := c.Bind(rw); err != nil {
		c.Logger().Error("failed on parsing request", err.Error())
		return err
	}

	w := h.workoutService.Request(*rw)

	if err := c.JSON(http.StatusCreated, w); err != nil {
		c.Logger().Error("failed on parsing response", err.Error())
		return err
	}
	return nil
}

func (h *WorkoutHandler) UpdateWorkout(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Error("invalid id", err.Error())
		c.JSON(http.StatusBadRequest, err)
		return nil
	}

	uw := new(models.UpdateWorkout)
	if err := c.Bind(uw); err != nil {
		c.Logger().Error("failed on parsing request", err.Error())
		return err
	}

	w, err := h.workoutService.Update(id, *uw)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return err
	}

	if err := c.JSON(http.StatusOK, w); err != nil {
		c.Logger().Error("failed on parsing response", err.Error())
		return err
	}
	return nil
}

func (h *WorkoutHandler) ApproveWorkout(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Error("invalid id", err.Error())
		c.JSON(http.StatusBadRequest, err)
		return nil
	}

	if err := h.workoutService.Approve(id); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return err
	}

	if err := c.NoContent(http.StatusNoContent); err != nil {
		c.Logger().Error("failed on parsing response", err.Error())
		return err
	}
	return nil
}

func (h *WorkoutHandler) CompleteWorkout(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Error("invalid id", err.Error())
		c.JSON(http.StatusBadRequest, err)
		return nil
	}

	if err := h.workoutService.Approve(id); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return err
	}

	if err := c.NoContent(http.StatusNoContent); err != nil {
		c.Logger().Error("failed on parsing response", err.Error())
		return err
	}
	return nil
}
