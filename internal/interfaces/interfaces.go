package interfaces

import (
	"context"

	"github.com/DJANGO-JANE/models"
)

type IStaffRepository interface {
	RegisterNew(context context.Context, staff *models.Staff) (*models.Staff, error)
	SignIn(context context.Context, staff *models.Staff) (*models.Staff, error)
	ChangePassword(context context.Context, staff *models.Staff) (bool, error)
	GetTeacher(context context.Context, staffId int) (*models.Staff, error)
	UpdateInformation(context context.Context, staff *models.Staff) (*models.Staff, error)
	GetAllStaff(context context.Context) ([]models.Staff, error)
	//ValidatePassword(validateThis string, compareToThis string)(bool, error)

}
