// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"io"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func encodeCreateTodoResponse(response CreateTodoRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *TodoCreate:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *R409:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		span.SetStatus(codes.Error, http.StatusText(409))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeDeleteTodoResponse(response DeleteTodoRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *DeleteTodoNoContent:
		w.WriteHeader(204)
		span.SetStatus(codes.Ok, http.StatusText(204))

		return nil

	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *R404:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *R409:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		span.SetStatus(codes.Error, http.StatusText(409))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeListTodoResponse(response ListTodoRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *ListTodoOKApplicationJSON:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *R404:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *R409:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		span.SetStatus(codes.Error, http.StatusText(409))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeListTodoChildrenResponse(response ListTodoChildrenRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *ListTodoChildrenOKApplicationJSON:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *R404:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *R409:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		span.SetStatus(codes.Error, http.StatusText(409))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeReadTodoResponse(response ReadTodoRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *TodoRead:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *R404:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *R409:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		span.SetStatus(codes.Error, http.StatusText(409))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeReadTodoParentResponse(response ReadTodoParentRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *TodoParentRead:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *R404:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *R409:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		span.SetStatus(codes.Error, http.StatusText(409))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeServerStatusResponse(response ServerStatusRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *ServerStatusOK:
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		writer := w
		if _, err := io.Copy(writer, response); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *ServerStatusServiceUnavailable:
		w.WriteHeader(503)
		span.SetStatus(codes.Error, http.StatusText(503))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeUpdateTodoResponse(response UpdateTodoRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *TodoUpdate:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *R404:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *R409:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		span.SetStatus(codes.Error, http.StatusText(409))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}
