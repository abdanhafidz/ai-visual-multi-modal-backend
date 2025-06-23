package services

import (
	"errors"
	"time"

	models "github.com/abdanhafidz/ai-visual-multi-modal-backend/models"
	repositories "github.com/abdanhafidz/ai-visual-multi-modal-backend/repositories"
	"gorm.io/gorm"
)

type (
	service[TRepo repositories.Repository] struct {
		repository TRepo
		exception  models.Exception
		errors     error
	}
	Service interface {
		ThrowsException(*bool, string)
		ThrowsError(error)
		Exception() models.Exception
		ThrowsRepoException() bool
		Error() error
	}
)

func (s *service[TRepo]) ThrowsException(status *bool, message string) {

	*status = true
	s.exception.Message = message

}

func (s *service[TRepo]) ThrowsError(err error) {

	s.errors = errors.Join(s.errors, err)

}

func (s *service[TRepo]) Exception() models.Exception {
	return s.exception
}
func (s *service[TRepo]) Error() error {
	return s.errors
}
func CalculateDueTime(duration time.Duration) time.Time {
	return time.Now().Add(duration)
}

func (s *service[TRepo]) ThrowsRepoException() bool {

	if s.repository.RowsError() != nil {

		s.ThrowsException(&s.exception.QueryError, "Database error!")
		s.ThrowsError(s.repository.RowsError())
		return true

	}
	
	if errors.Is(s.repository.RowsError(), gorm.ErrDuplicatedKey) {
		s.ThrowsException(&s.exception.DataDuplicate, "Duplicated data!")

		return true
	}
	if s.repository.IsNoRecord() {
		s.ThrowsException(&s.exception.DataNotFound, "No record found")
		return true
	}

	return false
}
