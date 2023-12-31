// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"io"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
)

type CreateTodoReq struct {
	Text      string              `json:"text"`
	CreatedAt time.Time           `json:"created_at"`
	Status    CreateTodoReqStatus `json:"status"`
	Priority  int                 `json:"priority"`
	Children  []int               `json:"children"`
	Parent    OptInt              `json:"parent"`
}

// GetText returns the value of Text.
func (s *CreateTodoReq) GetText() string {
	return s.Text
}

// GetCreatedAt returns the value of CreatedAt.
func (s *CreateTodoReq) GetCreatedAt() time.Time {
	return s.CreatedAt
}

// GetStatus returns the value of Status.
func (s *CreateTodoReq) GetStatus() CreateTodoReqStatus {
	return s.Status
}

// GetPriority returns the value of Priority.
func (s *CreateTodoReq) GetPriority() int {
	return s.Priority
}

// GetChildren returns the value of Children.
func (s *CreateTodoReq) GetChildren() []int {
	return s.Children
}

// GetParent returns the value of Parent.
func (s *CreateTodoReq) GetParent() OptInt {
	return s.Parent
}

// SetText sets the value of Text.
func (s *CreateTodoReq) SetText(val string) {
	s.Text = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *CreateTodoReq) SetCreatedAt(val time.Time) {
	s.CreatedAt = val
}

// SetStatus sets the value of Status.
func (s *CreateTodoReq) SetStatus(val CreateTodoReqStatus) {
	s.Status = val
}

// SetPriority sets the value of Priority.
func (s *CreateTodoReq) SetPriority(val int) {
	s.Priority = val
}

// SetChildren sets the value of Children.
func (s *CreateTodoReq) SetChildren(val []int) {
	s.Children = val
}

// SetParent sets the value of Parent.
func (s *CreateTodoReq) SetParent(val OptInt) {
	s.Parent = val
}

type CreateTodoReqStatus string

const (
	CreateTodoReqStatusINPROGRESS CreateTodoReqStatus = "IN_PROGRESS"
	CreateTodoReqStatusCOMPLETED  CreateTodoReqStatus = "COMPLETED"
)

// AllValues returns all CreateTodoReqStatus values.
func (CreateTodoReqStatus) AllValues() []CreateTodoReqStatus {
	return []CreateTodoReqStatus{
		CreateTodoReqStatusINPROGRESS,
		CreateTodoReqStatusCOMPLETED,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s CreateTodoReqStatus) MarshalText() ([]byte, error) {
	switch s {
	case CreateTodoReqStatusINPROGRESS:
		return []byte(s), nil
	case CreateTodoReqStatusCOMPLETED:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *CreateTodoReqStatus) UnmarshalText(data []byte) error {
	switch CreateTodoReqStatus(data) {
	case CreateTodoReqStatusINPROGRESS:
		*s = CreateTodoReqStatusINPROGRESS
		return nil
	case CreateTodoReqStatusCOMPLETED:
		*s = CreateTodoReqStatusCOMPLETED
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// DeleteTodoNoContent is response for DeleteTodo operation.
type DeleteTodoNoContent struct{}

func (*DeleteTodoNoContent) deleteTodoRes() {}

type ListTodoChildrenOKApplicationJSON []TodoChildrenList

func (*ListTodoChildrenOKApplicationJSON) listTodoChildrenRes() {}

type ListTodoOKApplicationJSON []TodoList

func (*ListTodoOKApplicationJSON) listTodoRes() {}

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptUpdateTodoReqStatus returns new OptUpdateTodoReqStatus with value set to v.
func NewOptUpdateTodoReqStatus(v UpdateTodoReqStatus) OptUpdateTodoReqStatus {
	return OptUpdateTodoReqStatus{
		Value: v,
		Set:   true,
	}
}

// OptUpdateTodoReqStatus is optional UpdateTodoReqStatus.
type OptUpdateTodoReqStatus struct {
	Value UpdateTodoReqStatus
	Set   bool
}

// IsSet returns true if OptUpdateTodoReqStatus was set.
func (o OptUpdateTodoReqStatus) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptUpdateTodoReqStatus) Reset() {
	var v UpdateTodoReqStatus
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptUpdateTodoReqStatus) SetTo(v UpdateTodoReqStatus) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptUpdateTodoReqStatus) Get() (v UpdateTodoReqStatus, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptUpdateTodoReqStatus) Or(d UpdateTodoReqStatus) UpdateTodoReqStatus {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

type R400 struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Errors jx.Raw `json:"errors"`
}

// GetCode returns the value of Code.
func (s *R400) GetCode() int {
	return s.Code
}

// GetStatus returns the value of Status.
func (s *R400) GetStatus() string {
	return s.Status
}

// GetErrors returns the value of Errors.
func (s *R400) GetErrors() jx.Raw {
	return s.Errors
}

// SetCode sets the value of Code.
func (s *R400) SetCode(val int) {
	s.Code = val
}

// SetStatus sets the value of Status.
func (s *R400) SetStatus(val string) {
	s.Status = val
}

// SetErrors sets the value of Errors.
func (s *R400) SetErrors(val jx.Raw) {
	s.Errors = val
}

func (*R400) createTodoRes()       {}
func (*R400) deleteTodoRes()       {}
func (*R400) listTodoChildrenRes() {}
func (*R400) listTodoRes()         {}
func (*R400) readTodoParentRes()   {}
func (*R400) readTodoRes()         {}
func (*R400) updateTodoRes()       {}

type R404 struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Errors jx.Raw `json:"errors"`
}

// GetCode returns the value of Code.
func (s *R404) GetCode() int {
	return s.Code
}

// GetStatus returns the value of Status.
func (s *R404) GetStatus() string {
	return s.Status
}

// GetErrors returns the value of Errors.
func (s *R404) GetErrors() jx.Raw {
	return s.Errors
}

// SetCode sets the value of Code.
func (s *R404) SetCode(val int) {
	s.Code = val
}

// SetStatus sets the value of Status.
func (s *R404) SetStatus(val string) {
	s.Status = val
}

// SetErrors sets the value of Errors.
func (s *R404) SetErrors(val jx.Raw) {
	s.Errors = val
}

func (*R404) deleteTodoRes()       {}
func (*R404) listTodoChildrenRes() {}
func (*R404) listTodoRes()         {}
func (*R404) readTodoParentRes()   {}
func (*R404) readTodoRes()         {}
func (*R404) updateTodoRes()       {}

type R409 struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Errors jx.Raw `json:"errors"`
}

// GetCode returns the value of Code.
func (s *R409) GetCode() int {
	return s.Code
}

// GetStatus returns the value of Status.
func (s *R409) GetStatus() string {
	return s.Status
}

// GetErrors returns the value of Errors.
func (s *R409) GetErrors() jx.Raw {
	return s.Errors
}

// SetCode sets the value of Code.
func (s *R409) SetCode(val int) {
	s.Code = val
}

// SetStatus sets the value of Status.
func (s *R409) SetStatus(val string) {
	s.Status = val
}

// SetErrors sets the value of Errors.
func (s *R409) SetErrors(val jx.Raw) {
	s.Errors = val
}

func (*R409) createTodoRes()       {}
func (*R409) deleteTodoRes()       {}
func (*R409) listTodoChildrenRes() {}
func (*R409) listTodoRes()         {}
func (*R409) readTodoParentRes()   {}
func (*R409) readTodoRes()         {}
func (*R409) updateTodoRes()       {}

type R500 struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Errors jx.Raw `json:"errors"`
}

// GetCode returns the value of Code.
func (s *R500) GetCode() int {
	return s.Code
}

// GetStatus returns the value of Status.
func (s *R500) GetStatus() string {
	return s.Status
}

// GetErrors returns the value of Errors.
func (s *R500) GetErrors() jx.Raw {
	return s.Errors
}

// SetCode sets the value of Code.
func (s *R500) SetCode(val int) {
	s.Code = val
}

// SetStatus sets the value of Status.
func (s *R500) SetStatus(val string) {
	s.Status = val
}

// SetErrors sets the value of Errors.
func (s *R500) SetErrors(val jx.Raw) {
	s.Errors = val
}

func (*R500) createTodoRes()       {}
func (*R500) deleteTodoRes()       {}
func (*R500) listTodoChildrenRes() {}
func (*R500) listTodoRes()         {}
func (*R500) readTodoParentRes()   {}
func (*R500) readTodoRes()         {}
func (*R500) updateTodoRes()       {}

type ServerStatusOK struct {
	Data io.Reader
}

// Read reads data from the Data reader.
//
// Kept to satisfy the io.Reader interface.
func (s ServerStatusOK) Read(p []byte) (n int, err error) {
	if s.Data == nil {
		return 0, io.EOF
	}
	return s.Data.Read(p)
}

func (*ServerStatusOK) serverStatusRes() {}

// ServerStatusServiceUnavailable is response for ServerStatus operation.
type ServerStatusServiceUnavailable struct{}

func (*ServerStatusServiceUnavailable) serverStatusRes() {}

// Ref: #/components/schemas/Todo_ChildrenList
type TodoChildrenList struct {
	ID        int                    `json:"id"`
	Text      string                 `json:"text"`
	CreatedAt time.Time              `json:"created_at"`
	Status    TodoChildrenListStatus `json:"status"`
	Priority  int                    `json:"priority"`
}

// GetID returns the value of ID.
func (s *TodoChildrenList) GetID() int {
	return s.ID
}

// GetText returns the value of Text.
func (s *TodoChildrenList) GetText() string {
	return s.Text
}

// GetCreatedAt returns the value of CreatedAt.
func (s *TodoChildrenList) GetCreatedAt() time.Time {
	return s.CreatedAt
}

// GetStatus returns the value of Status.
func (s *TodoChildrenList) GetStatus() TodoChildrenListStatus {
	return s.Status
}

// GetPriority returns the value of Priority.
func (s *TodoChildrenList) GetPriority() int {
	return s.Priority
}

// SetID sets the value of ID.
func (s *TodoChildrenList) SetID(val int) {
	s.ID = val
}

// SetText sets the value of Text.
func (s *TodoChildrenList) SetText(val string) {
	s.Text = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *TodoChildrenList) SetCreatedAt(val time.Time) {
	s.CreatedAt = val
}

// SetStatus sets the value of Status.
func (s *TodoChildrenList) SetStatus(val TodoChildrenListStatus) {
	s.Status = val
}

// SetPriority sets the value of Priority.
func (s *TodoChildrenList) SetPriority(val int) {
	s.Priority = val
}

type TodoChildrenListStatus string

const (
	TodoChildrenListStatusINPROGRESS TodoChildrenListStatus = "IN_PROGRESS"
	TodoChildrenListStatusCOMPLETED  TodoChildrenListStatus = "COMPLETED"
)

// AllValues returns all TodoChildrenListStatus values.
func (TodoChildrenListStatus) AllValues() []TodoChildrenListStatus {
	return []TodoChildrenListStatus{
		TodoChildrenListStatusINPROGRESS,
		TodoChildrenListStatusCOMPLETED,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s TodoChildrenListStatus) MarshalText() ([]byte, error) {
	switch s {
	case TodoChildrenListStatusINPROGRESS:
		return []byte(s), nil
	case TodoChildrenListStatusCOMPLETED:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *TodoChildrenListStatus) UnmarshalText(data []byte) error {
	switch TodoChildrenListStatus(data) {
	case TodoChildrenListStatusINPROGRESS:
		*s = TodoChildrenListStatusINPROGRESS
		return nil
	case TodoChildrenListStatusCOMPLETED:
		*s = TodoChildrenListStatusCOMPLETED
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// Ref: #/components/schemas/TodoCreate
type TodoCreate struct {
	ID        int              `json:"id"`
	Text      string           `json:"text"`
	CreatedAt time.Time        `json:"created_at"`
	Status    TodoCreateStatus `json:"status"`
	Priority  int              `json:"priority"`
}

// GetID returns the value of ID.
func (s *TodoCreate) GetID() int {
	return s.ID
}

// GetText returns the value of Text.
func (s *TodoCreate) GetText() string {
	return s.Text
}

// GetCreatedAt returns the value of CreatedAt.
func (s *TodoCreate) GetCreatedAt() time.Time {
	return s.CreatedAt
}

// GetStatus returns the value of Status.
func (s *TodoCreate) GetStatus() TodoCreateStatus {
	return s.Status
}

// GetPriority returns the value of Priority.
func (s *TodoCreate) GetPriority() int {
	return s.Priority
}

// SetID sets the value of ID.
func (s *TodoCreate) SetID(val int) {
	s.ID = val
}

// SetText sets the value of Text.
func (s *TodoCreate) SetText(val string) {
	s.Text = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *TodoCreate) SetCreatedAt(val time.Time) {
	s.CreatedAt = val
}

// SetStatus sets the value of Status.
func (s *TodoCreate) SetStatus(val TodoCreateStatus) {
	s.Status = val
}

// SetPriority sets the value of Priority.
func (s *TodoCreate) SetPriority(val int) {
	s.Priority = val
}

func (*TodoCreate) createTodoRes() {}

type TodoCreateStatus string

const (
	TodoCreateStatusINPROGRESS TodoCreateStatus = "IN_PROGRESS"
	TodoCreateStatusCOMPLETED  TodoCreateStatus = "COMPLETED"
)

// AllValues returns all TodoCreateStatus values.
func (TodoCreateStatus) AllValues() []TodoCreateStatus {
	return []TodoCreateStatus{
		TodoCreateStatusINPROGRESS,
		TodoCreateStatusCOMPLETED,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s TodoCreateStatus) MarshalText() ([]byte, error) {
	switch s {
	case TodoCreateStatusINPROGRESS:
		return []byte(s), nil
	case TodoCreateStatusCOMPLETED:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *TodoCreateStatus) UnmarshalText(data []byte) error {
	switch TodoCreateStatus(data) {
	case TodoCreateStatusINPROGRESS:
		*s = TodoCreateStatusINPROGRESS
		return nil
	case TodoCreateStatusCOMPLETED:
		*s = TodoCreateStatusCOMPLETED
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// Ref: #/components/schemas/TodoList
type TodoList struct {
	ID        int            `json:"id"`
	Text      string         `json:"text"`
	CreatedAt time.Time      `json:"created_at"`
	Status    TodoListStatus `json:"status"`
	Priority  int            `json:"priority"`
}

// GetID returns the value of ID.
func (s *TodoList) GetID() int {
	return s.ID
}

// GetText returns the value of Text.
func (s *TodoList) GetText() string {
	return s.Text
}

// GetCreatedAt returns the value of CreatedAt.
func (s *TodoList) GetCreatedAt() time.Time {
	return s.CreatedAt
}

// GetStatus returns the value of Status.
func (s *TodoList) GetStatus() TodoListStatus {
	return s.Status
}

// GetPriority returns the value of Priority.
func (s *TodoList) GetPriority() int {
	return s.Priority
}

// SetID sets the value of ID.
func (s *TodoList) SetID(val int) {
	s.ID = val
}

// SetText sets the value of Text.
func (s *TodoList) SetText(val string) {
	s.Text = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *TodoList) SetCreatedAt(val time.Time) {
	s.CreatedAt = val
}

// SetStatus sets the value of Status.
func (s *TodoList) SetStatus(val TodoListStatus) {
	s.Status = val
}

// SetPriority sets the value of Priority.
func (s *TodoList) SetPriority(val int) {
	s.Priority = val
}

type TodoListStatus string

const (
	TodoListStatusINPROGRESS TodoListStatus = "IN_PROGRESS"
	TodoListStatusCOMPLETED  TodoListStatus = "COMPLETED"
)

// AllValues returns all TodoListStatus values.
func (TodoListStatus) AllValues() []TodoListStatus {
	return []TodoListStatus{
		TodoListStatusINPROGRESS,
		TodoListStatusCOMPLETED,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s TodoListStatus) MarshalText() ([]byte, error) {
	switch s {
	case TodoListStatusINPROGRESS:
		return []byte(s), nil
	case TodoListStatusCOMPLETED:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *TodoListStatus) UnmarshalText(data []byte) error {
	switch TodoListStatus(data) {
	case TodoListStatusINPROGRESS:
		*s = TodoListStatusINPROGRESS
		return nil
	case TodoListStatusCOMPLETED:
		*s = TodoListStatusCOMPLETED
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// Ref: #/components/schemas/Todo_ParentRead
type TodoParentRead struct {
	ID        int                  `json:"id"`
	Text      string               `json:"text"`
	CreatedAt time.Time            `json:"created_at"`
	Status    TodoParentReadStatus `json:"status"`
	Priority  int                  `json:"priority"`
}

// GetID returns the value of ID.
func (s *TodoParentRead) GetID() int {
	return s.ID
}

// GetText returns the value of Text.
func (s *TodoParentRead) GetText() string {
	return s.Text
}

// GetCreatedAt returns the value of CreatedAt.
func (s *TodoParentRead) GetCreatedAt() time.Time {
	return s.CreatedAt
}

// GetStatus returns the value of Status.
func (s *TodoParentRead) GetStatus() TodoParentReadStatus {
	return s.Status
}

// GetPriority returns the value of Priority.
func (s *TodoParentRead) GetPriority() int {
	return s.Priority
}

// SetID sets the value of ID.
func (s *TodoParentRead) SetID(val int) {
	s.ID = val
}

// SetText sets the value of Text.
func (s *TodoParentRead) SetText(val string) {
	s.Text = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *TodoParentRead) SetCreatedAt(val time.Time) {
	s.CreatedAt = val
}

// SetStatus sets the value of Status.
func (s *TodoParentRead) SetStatus(val TodoParentReadStatus) {
	s.Status = val
}

// SetPriority sets the value of Priority.
func (s *TodoParentRead) SetPriority(val int) {
	s.Priority = val
}

func (*TodoParentRead) readTodoParentRes() {}

type TodoParentReadStatus string

const (
	TodoParentReadStatusINPROGRESS TodoParentReadStatus = "IN_PROGRESS"
	TodoParentReadStatusCOMPLETED  TodoParentReadStatus = "COMPLETED"
)

// AllValues returns all TodoParentReadStatus values.
func (TodoParentReadStatus) AllValues() []TodoParentReadStatus {
	return []TodoParentReadStatus{
		TodoParentReadStatusINPROGRESS,
		TodoParentReadStatusCOMPLETED,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s TodoParentReadStatus) MarshalText() ([]byte, error) {
	switch s {
	case TodoParentReadStatusINPROGRESS:
		return []byte(s), nil
	case TodoParentReadStatusCOMPLETED:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *TodoParentReadStatus) UnmarshalText(data []byte) error {
	switch TodoParentReadStatus(data) {
	case TodoParentReadStatusINPROGRESS:
		*s = TodoParentReadStatusINPROGRESS
		return nil
	case TodoParentReadStatusCOMPLETED:
		*s = TodoParentReadStatusCOMPLETED
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// Ref: #/components/schemas/TodoRead
type TodoRead struct {
	ID        int            `json:"id"`
	Text      string         `json:"text"`
	CreatedAt time.Time      `json:"created_at"`
	Status    TodoReadStatus `json:"status"`
	Priority  int            `json:"priority"`
}

// GetID returns the value of ID.
func (s *TodoRead) GetID() int {
	return s.ID
}

// GetText returns the value of Text.
func (s *TodoRead) GetText() string {
	return s.Text
}

// GetCreatedAt returns the value of CreatedAt.
func (s *TodoRead) GetCreatedAt() time.Time {
	return s.CreatedAt
}

// GetStatus returns the value of Status.
func (s *TodoRead) GetStatus() TodoReadStatus {
	return s.Status
}

// GetPriority returns the value of Priority.
func (s *TodoRead) GetPriority() int {
	return s.Priority
}

// SetID sets the value of ID.
func (s *TodoRead) SetID(val int) {
	s.ID = val
}

// SetText sets the value of Text.
func (s *TodoRead) SetText(val string) {
	s.Text = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *TodoRead) SetCreatedAt(val time.Time) {
	s.CreatedAt = val
}

// SetStatus sets the value of Status.
func (s *TodoRead) SetStatus(val TodoReadStatus) {
	s.Status = val
}

// SetPriority sets the value of Priority.
func (s *TodoRead) SetPriority(val int) {
	s.Priority = val
}

func (*TodoRead) readTodoRes() {}

type TodoReadStatus string

const (
	TodoReadStatusINPROGRESS TodoReadStatus = "IN_PROGRESS"
	TodoReadStatusCOMPLETED  TodoReadStatus = "COMPLETED"
)

// AllValues returns all TodoReadStatus values.
func (TodoReadStatus) AllValues() []TodoReadStatus {
	return []TodoReadStatus{
		TodoReadStatusINPROGRESS,
		TodoReadStatusCOMPLETED,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s TodoReadStatus) MarshalText() ([]byte, error) {
	switch s {
	case TodoReadStatusINPROGRESS:
		return []byte(s), nil
	case TodoReadStatusCOMPLETED:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *TodoReadStatus) UnmarshalText(data []byte) error {
	switch TodoReadStatus(data) {
	case TodoReadStatusINPROGRESS:
		*s = TodoReadStatusINPROGRESS
		return nil
	case TodoReadStatusCOMPLETED:
		*s = TodoReadStatusCOMPLETED
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// Ref: #/components/schemas/TodoUpdate
type TodoUpdate struct {
	ID        int              `json:"id"`
	Text      string           `json:"text"`
	CreatedAt time.Time        `json:"created_at"`
	Status    TodoUpdateStatus `json:"status"`
	Priority  int              `json:"priority"`
}

// GetID returns the value of ID.
func (s *TodoUpdate) GetID() int {
	return s.ID
}

// GetText returns the value of Text.
func (s *TodoUpdate) GetText() string {
	return s.Text
}

// GetCreatedAt returns the value of CreatedAt.
func (s *TodoUpdate) GetCreatedAt() time.Time {
	return s.CreatedAt
}

// GetStatus returns the value of Status.
func (s *TodoUpdate) GetStatus() TodoUpdateStatus {
	return s.Status
}

// GetPriority returns the value of Priority.
func (s *TodoUpdate) GetPriority() int {
	return s.Priority
}

// SetID sets the value of ID.
func (s *TodoUpdate) SetID(val int) {
	s.ID = val
}

// SetText sets the value of Text.
func (s *TodoUpdate) SetText(val string) {
	s.Text = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *TodoUpdate) SetCreatedAt(val time.Time) {
	s.CreatedAt = val
}

// SetStatus sets the value of Status.
func (s *TodoUpdate) SetStatus(val TodoUpdateStatus) {
	s.Status = val
}

// SetPriority sets the value of Priority.
func (s *TodoUpdate) SetPriority(val int) {
	s.Priority = val
}

func (*TodoUpdate) updateTodoRes() {}

type TodoUpdateStatus string

const (
	TodoUpdateStatusINPROGRESS TodoUpdateStatus = "IN_PROGRESS"
	TodoUpdateStatusCOMPLETED  TodoUpdateStatus = "COMPLETED"
)

// AllValues returns all TodoUpdateStatus values.
func (TodoUpdateStatus) AllValues() []TodoUpdateStatus {
	return []TodoUpdateStatus{
		TodoUpdateStatusINPROGRESS,
		TodoUpdateStatusCOMPLETED,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s TodoUpdateStatus) MarshalText() ([]byte, error) {
	switch s {
	case TodoUpdateStatusINPROGRESS:
		return []byte(s), nil
	case TodoUpdateStatusCOMPLETED:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *TodoUpdateStatus) UnmarshalText(data []byte) error {
	switch TodoUpdateStatus(data) {
	case TodoUpdateStatusINPROGRESS:
		*s = TodoUpdateStatusINPROGRESS
		return nil
	case TodoUpdateStatusCOMPLETED:
		*s = TodoUpdateStatusCOMPLETED
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

type UpdateTodoReq struct {
	Text     OptString              `json:"text"`
	Status   OptUpdateTodoReqStatus `json:"status"`
	Priority OptInt                 `json:"priority"`
	Children []int                  `json:"children"`
	Parent   OptInt                 `json:"parent"`
}

// GetText returns the value of Text.
func (s *UpdateTodoReq) GetText() OptString {
	return s.Text
}

// GetStatus returns the value of Status.
func (s *UpdateTodoReq) GetStatus() OptUpdateTodoReqStatus {
	return s.Status
}

// GetPriority returns the value of Priority.
func (s *UpdateTodoReq) GetPriority() OptInt {
	return s.Priority
}

// GetChildren returns the value of Children.
func (s *UpdateTodoReq) GetChildren() []int {
	return s.Children
}

// GetParent returns the value of Parent.
func (s *UpdateTodoReq) GetParent() OptInt {
	return s.Parent
}

// SetText sets the value of Text.
func (s *UpdateTodoReq) SetText(val OptString) {
	s.Text = val
}

// SetStatus sets the value of Status.
func (s *UpdateTodoReq) SetStatus(val OptUpdateTodoReqStatus) {
	s.Status = val
}

// SetPriority sets the value of Priority.
func (s *UpdateTodoReq) SetPriority(val OptInt) {
	s.Priority = val
}

// SetChildren sets the value of Children.
func (s *UpdateTodoReq) SetChildren(val []int) {
	s.Children = val
}

// SetParent sets the value of Parent.
func (s *UpdateTodoReq) SetParent(val OptInt) {
	s.Parent = val
}

type UpdateTodoReqStatus string

const (
	UpdateTodoReqStatusINPROGRESS UpdateTodoReqStatus = "IN_PROGRESS"
	UpdateTodoReqStatusCOMPLETED  UpdateTodoReqStatus = "COMPLETED"
)

// AllValues returns all UpdateTodoReqStatus values.
func (UpdateTodoReqStatus) AllValues() []UpdateTodoReqStatus {
	return []UpdateTodoReqStatus{
		UpdateTodoReqStatusINPROGRESS,
		UpdateTodoReqStatusCOMPLETED,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s UpdateTodoReqStatus) MarshalText() ([]byte, error) {
	switch s {
	case UpdateTodoReqStatusINPROGRESS:
		return []byte(s), nil
	case UpdateTodoReqStatusCOMPLETED:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *UpdateTodoReqStatus) UnmarshalText(data []byte) error {
	switch UpdateTodoReqStatus(data) {
	case UpdateTodoReqStatusINPROGRESS:
		*s = UpdateTodoReqStatusINPROGRESS
		return nil
	case UpdateTodoReqStatusCOMPLETED:
		*s = UpdateTodoReqStatusCOMPLETED
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}
