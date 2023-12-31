// Code generated by ent, DO NOT EDIT.

package ogent

import (
	"context"
	"net/http"

	"github.com/Grayson/ent-eval/ent"
	"github.com/Grayson/ent-eval/ent/todo"
	"github.com/go-faster/jx"
)

// OgentHandler implements the ogen generated Handler interface and uses Ent as data layer.
type OgentHandler struct {
	client *ent.Client
}

// NewOgentHandler returns a new OgentHandler.
func NewOgentHandler(c *ent.Client) *OgentHandler { return &OgentHandler{c} }

// rawError renders err as json string.
func rawError(err error) jx.Raw {
	var e jx.Encoder
	e.Str(err.Error())
	return e.Bytes()
}

// CreateTodo handles POST /todos requests.
func (h *OgentHandler) CreateTodo(ctx context.Context, req *CreateTodoReq) (CreateTodoRes, error) {
	b := h.client.Todo.Create()
	// Add all fields.
	b.SetText(req.Text)
	b.SetCreatedAt(req.CreatedAt)
	b.SetStatus(todo.Status(req.Status))
	b.SetPriority(req.Priority)
	// Add all edges.
	b.AddChildIDs(req.Children...)
	if v, ok := req.Parent.Get(); ok {
		b.SetParentID(v)
	}
	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		switch {
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	// Reload the entity to attach all eager-loaded edges.
	q := h.client.Todo.Query().Where(todo.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return nil, err
	}
	return NewTodoCreate(e), nil
}

// ReadTodo handles GET /todos/{id} requests.
func (h *OgentHandler) ReadTodo(ctx context.Context, params ReadTodoParams) (ReadTodoRes, error) {
	q := h.client.Todo.Query().Where(todo.IDEQ(params.ID))
	e, err := q.Only(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	return NewTodoRead(e), nil
}

// UpdateTodo handles PATCH /todos/{id} requests.
func (h *OgentHandler) UpdateTodo(ctx context.Context, req *UpdateTodoReq, params UpdateTodoParams) (UpdateTodoRes, error) {
	b := h.client.Todo.UpdateOneID(params.ID)
	// Add all fields.
	if v, ok := req.Text.Get(); ok {
		b.SetText(v)
	}
	if v, ok := req.Status.Get(); ok {
		b.SetStatus(todo.Status(v))
	}
	if v, ok := req.Priority.Get(); ok {
		b.SetPriority(v)
	}
	// Add all edges.
	if req.Children != nil {
		b.ClearChildren().AddChildIDs(req.Children...)
	}
	if v, ok := req.Parent.Get(); ok {
		b.SetParentID(v)
	}
	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	// Reload the entity to attach all eager-loaded edges.
	q := h.client.Todo.Query().Where(todo.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return nil, err
	}
	return NewTodoUpdate(e), nil
}

// DeleteTodo handles DELETE /todos/{id} requests.
func (h *OgentHandler) DeleteTodo(ctx context.Context, params DeleteTodoParams) (DeleteTodoRes, error) {
	err := h.client.Todo.DeleteOneID(params.ID).Exec(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	return new(DeleteTodoNoContent), nil

}

// ListTodo handles GET /todos requests.
func (h *OgentHandler) ListTodo(ctx context.Context, params ListTodoParams) (ListTodoRes, error) {
	q := h.client.Todo.Query()
	page := 1
	if v, ok := params.Page.Get(); ok {
		page = v
	}
	itemsPerPage := 30
	if v, ok := params.ItemsPerPage.Get(); ok {
		itemsPerPage = v
	}
	q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage)

	es, err := q.All(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	r := NewTodoLists(es)
	return (*ListTodoOKApplicationJSON)(&r), nil
}

// ListTodoChildren handles GET /todos/{id}/children requests.
func (h *OgentHandler) ListTodoChildren(ctx context.Context, params ListTodoChildrenParams) (ListTodoChildrenRes, error) {
	q := h.client.Todo.Query().Where(todo.IDEQ(params.ID)).QueryChildren()
	page := 1
	if v, ok := params.Page.Get(); ok {
		page = v
	}
	itemsPerPage := 30
	if v, ok := params.ItemsPerPage.Get(); ok {
		itemsPerPage = v
	}
	q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage)
	es, err := q.All(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	r := NewTodoChildrenLists(es)
	return (*ListTodoChildrenOKApplicationJSON)(&r), nil
}

// ReadTodoParent handles GET /todos/{id}/parent requests.
func (h *OgentHandler) ReadTodoParent(ctx context.Context, params ReadTodoParentParams) (ReadTodoParentRes, error) {
	q := h.client.Todo.Query().Where(todo.IDEQ(params.ID)).QueryParent()
	e, err := q.Only(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	return NewTodoParentRead(e), nil
}
