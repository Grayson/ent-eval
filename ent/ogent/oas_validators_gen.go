// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"fmt"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/validate"
)

func (s *CreateTodoReq) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.Status.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "status",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s CreateTodoReqStatus) Validate() error {
	switch s {
	case "IN_PROGRESS":
		return nil
	case "COMPLETED":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s ListTodoChildrenOKApplicationJSON) Validate() error {
	alias := ([]TodoChildrenList)(s)
	if alias == nil {
		return errors.New("nil is invalid value")
	}
	var failures []validate.FieldError
	for i, elem := range alias {
		if err := func() error {
			if err := elem.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			failures = append(failures, validate.FieldError{
				Name:  fmt.Sprintf("[%d]", i),
				Error: err,
			})
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s ListTodoOKApplicationJSON) Validate() error {
	alias := ([]TodoList)(s)
	if alias == nil {
		return errors.New("nil is invalid value")
	}
	var failures []validate.FieldError
	for i, elem := range alias {
		if err := func() error {
			if err := elem.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			failures = append(failures, validate.FieldError{
				Name:  fmt.Sprintf("[%d]", i),
				Error: err,
			})
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *TodoChildrenList) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.Status.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "status",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s TodoChildrenListStatus) Validate() error {
	switch s {
	case "IN_PROGRESS":
		return nil
	case "COMPLETED":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s *TodoCreate) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.Status.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "status",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s TodoCreateStatus) Validate() error {
	switch s {
	case "IN_PROGRESS":
		return nil
	case "COMPLETED":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s *TodoList) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.Status.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "status",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s TodoListStatus) Validate() error {
	switch s {
	case "IN_PROGRESS":
		return nil
	case "COMPLETED":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s *TodoParentRead) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.Status.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "status",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s TodoParentReadStatus) Validate() error {
	switch s {
	case "IN_PROGRESS":
		return nil
	case "COMPLETED":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s *TodoRead) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.Status.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "status",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s TodoReadStatus) Validate() error {
	switch s {
	case "IN_PROGRESS":
		return nil
	case "COMPLETED":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s *TodoUpdate) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := s.Status.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "status",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s TodoUpdateStatus) Validate() error {
	switch s {
	case "IN_PROGRESS":
		return nil
	case "COMPLETED":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s *UpdateTodoReq) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if value, ok := s.Status.Get(); ok {
			if err := func() error {
				if err := value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "status",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s UpdateTodoReqStatus) Validate() error {
	switch s {
	case "IN_PROGRESS":
		return nil
	case "COMPLETED":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
