package schema

import (
	"context"
	"errors"

	"github.com/kyawmyintthein/event-ingestion-gateway/internal/domain"
)

var ErrSchemaNotFound = errors.New("schema not found")

type Definition struct {
	ID       string
	Version  string
	Encoding string
	Content  string
}

type Registry interface {
	Lookup(ctx context.Context, id string, version string) (Definition, error)
}

type FileRegistry struct{ Root string }

func (r *FileRegistry) Lookup(ctx context.Context, id string, version string) (Definition, error) {
	// TODO(student): resolve schema identity and version from repository files.
	return Definition{}, domain.ErrNotImplemented
}
