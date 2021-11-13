package playlists

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCanGetShift(t *testing.T) {
	expected := &Shift{Name: "Ewan"}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockrepository(ctrl)
	repo.EXPECT().Get(context.Background(), "abc123").Return(expected, nil)

	uc := Usecase{repo}

	shift, err := uc.Get(context.Background(), "abc123")

	assert.NoError(t, err)
	assert.Equal(t, expected, shift)
}

func TestCanGetAllShifts(t *testing.T) {
	expected := []*Shift{
		&Shift{Name: "test1", Age: 1},
		&Shift{Name: "test2", Age: 2},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockrepository(ctrl)
	repo.EXPECT().GetAll(context.Background()).Return(expected, nil)

	uc := Usecase{repo}

	shifts, err := uc.GetAll(context.Background())
	assert.NoError(t, err)
	assert.Len(t, shifts, 2)
	assert.Equal(t, expected, shifts)
}

func TestCanCreateShift(t *testing.T) {
	expected := &Shift{
		Name:  "testing",
		Email: "test@test.com",
		Age:   30,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockrepository(ctrl)
	repo.EXPECT().Create(context.Background(), expected).Return(nil)

	uc := Usecase{repo}
	err := uc.Create(context.Background(), expected)

	assert.NoError(t, err)
}

func TestCanValidateShift(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockrepository(ctrl)

	uc := Usecase{repo}

	shifts := []*Shift{
		&Shift{},                      // No required fields
		&Shift{Name: "", Age: 0},      // Blank name
		&Shift{Name: "123", Age: 200}, // Integers as name, age too high
		&Shift{Email: "nope"},
	}
	for _, val := range shifts {
		err := uc.Create(context.Background(), val)
		assert.Error(t, err)
	}
}

func TestCanUpdateShift(t *testing.T) {
	shift := &UpdateShift{
		Name:  "new name",
		Email: "test@test.com",
		Age:   20,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockrepository(ctrl)
	repo.EXPECT().Update(context.Background(), "abc123", shift).Return(nil)
	uc := Usecase{repo}
	err := uc.Update(context.Background(), "abc123", shift)
	assert.NoError(t, err)
}

func TestCanDeleteShift(t *testing.T) {
	shift := &Shift{
		ID: "abc123",
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockrepository(ctrl)
	repo.EXPECT().Delete(context.Background(), shift.ID).Return(nil)
	uc := Usecase{repo}
	err := uc.Delete(context.Background(), shift.ID)
	assert.NoError(t, err)
}
