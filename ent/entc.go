//go:build ignore

package main

import (
	"log"

	"ariga.io/ogent"
	"entgo.io/contrib/entoas"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/ogen-go/ogen"
	"github.com/ogen-go/ogen/gen/ir"
)

func main() {
	spec := new(ogen.Spec)
	oas, err := entoas.NewExtension(entoas.Spec(spec), entoas.Mutations(mutations()...))
	if err != nil {
		log.Fatalf("creating entoas extension: %v", err)
	}

	ogent, err := ogent.NewExtension(spec)
	if err != nil {
		log.Fatalf("creating ogent extension: %v", err)
	}

	err = entc.Generate("./schema", &gen.Config{}, entc.Extensions(ogent, oas))
	if err != nil {
		log.Fatalf("runnint ent codegen: %v", err)
	}
}

func mutations() []entoas.MutateFunc {
	return []entoas.MutateFunc{
		func(g *gen.Graph, s *ogen.Spec) error {
			okResponse := ogen.NewResponse().SetDescription("Server is alive and well").
				AddContent(ir.EncodingTextPlain.String(), ogen.NewSchema().SetType("string"))
			operation := ogen.NewOperation().
				SetOperationID("ServerStatus").
				SetSummary("Get information about the server").
				AddResponse("200", okResponse).
				AddResponse("503", ogen.NewResponse().SetDescription("Server is unavailable"))
			path := ogen.NewPathItem().
				SetDescription("Returns basic data about the server").
				SetGet(operation)
			s.AddPathItem("/status", path)
			return nil
		},
	}
}
