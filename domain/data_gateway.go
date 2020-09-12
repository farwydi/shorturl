package domain

import (
	"context"
	"net/url"

	"github.com/pkg/errors"
)

var (
	ErrNotFound = errors.New("not found")
)

type (
	DataGateway interface {
		SaveUrl(context.Context, *url.URL) (prefix string, err error)
		FindUrl(ctx context.Context, prefix string) (*url.URL, error)
	}
)
