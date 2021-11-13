package playlists

import (
	"context"

	"go.uber.org/zap"
)

// LoggerAdapter wraps the usecase interface
// with a logging adapter which can be swapped out
type LoggerAdapter struct {
	Logger  *zap.Logger
	Usecase ShiftService
}

func (a *LoggerAdapter) logErr(err error) {
	if err != nil {
		a.Logger.Error(err.Error())
	}
}

// Get a single shift
func (a *LoggerAdapter) Get(ctx context.Context, id string) (*Shift, error) {
	defer a.Logger.Sync()
	a.Logger.With(zap.String("id", id))
	a.Logger.Info("getting a single shift")
	shift, err := a.Usecase.Get(ctx, id)
	a.logErr(err)
	return shift, err
}

// GetAll gets all shifts
func (a *LoggerAdapter) GetAll(ctx context.Context) ([]*Shift, error) {
	defer a.Logger.Sync()
	a.Logger.Info("getting all shifts")
	shifts, err := a.Usecase.GetAll(ctx)
	a.logErr(err)
	return shifts, err
}

// Update a single shift
func (a *LoggerAdapter) Update(ctx context.Context, id string, shift *UpdateShift) error {
	defer a.Logger.Sync()
	a.Logger.With(zap.String("id", id))
	a.Logger.Info("updating a single shift")
	err := a.Usecase.Update(ctx, id, shift)
	a.logErr(err)
	return err
}

// Create a single shift
func (a *LoggerAdapter) Create(ctx context.Context, shift *Shift) error {
	defer a.Logger.Sync()
	a.Logger.Info("creating a single shift")
	err := a.Usecase.Create(ctx, shift)
	a.logErr(err)
	return err
}

// Delete a single shift
func (a *LoggerAdapter) Delete(ctx context.Context, id string) error {
	defer a.Logger.Sync()
	a.Logger.Info("deleting a single shift")
	err := a.Usecase.Delete(ctx, id)
	a.logErr(err)
	return err
}
