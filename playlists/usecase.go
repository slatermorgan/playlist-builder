package playlists

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
)

var (
	validate *validator.Validate
)

type repository interface {
	Get(ctx context.Context, id string) (*Shift, error)
	GetAll(ctx context.Context) ([]*Shift, error)
	Update(ctx context.Context, id string, shift *UpdateShift) error
	Create(ctx context.Context, shift *Shift) error
	Delete(ctx context.Context, id string) error
}

// Usecase for interacting with shifts
type Usecase struct {
	Repository repository
}

// Get a single shift
func (u *Usecase) Get(ctx context.Context, id string) (*Shift, error) {
	shift, err := u.Repository.Get(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "error fetching a single shift")
	}
	return shift, nil
}

// GetAll gets all shifts
func (u *Usecase) GetAll(ctx context.Context) ([]*Shift, error) {
	shifts, err := u.Repository.GetAll(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error fetching all shifts")
	}
	return shifts, nil
}

// Update a single shift
func (u *Usecase) Update(ctx context.Context, id string, shift *UpdateShift) error {
	validate = validator.New()
	if err := validate.Struct(shift); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return validationErrors
	}

	if err := u.Repository.Update(ctx, id, shift); err != nil {
		return errors.Wrap(err, "error updating shift")
	}
	return nil
}

// Create a single shift
func (u *Usecase) Create(ctx context.Context, shift *Shift) error {
	validate = validator.New()
	if err := validate.Struct(*shift); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		// Here we should create custom returns
		return validationErrors
	}

	shift.ID = u.newID()
	if err := u.Repository.Create(ctx, shift); err != nil {
		return errors.Wrap(err, "error creating new shift")
	}

	return nil
}

// Delete a single shift
func (u *Usecase) Delete(ctx context.Context, id string) error {
	if err := u.Repository.Delete(ctx, id); err != nil {
		return errors.Wrap(err, "error deleting shift")
	}
	return nil
}

func (u *Usecase) newID() string {
	uid := uuid.New()
	return uid.String()
}
