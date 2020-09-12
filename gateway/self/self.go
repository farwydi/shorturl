package self

import (
	"context"
	"encoding/base64"
	"github.com/farwydi/shorturl/domain"
	lru "github.com/hashicorp/golang-lru"
	"hash/adler32"
	"net/url"
	"time"
)

func NewSelfDataGateway(config domain.Config) (domain.DataGateway, error) {
	cache, err := lru.New(config.Gateway.Self.CacheSize)
	if err != nil {
		return nil, err
	}

	return &selfData{
		cache: cache,
	}, nil
}

type selfData struct {
	cache *lru.Cache
}

func (s *selfData) SaveUrl(_ context.Context, url *url.URL) (prefix string, err error) {
	tm, err := time.Now().MarshalBinary()
	if err != nil {
		return "", err
	}

	h := adler32.New()
	_, _ = h.Write(tm)
	prefix = base64.RawURLEncoding.EncodeToString(h.Sum(nil))

	_ = s.cache.Add(prefix, url)
	return prefix, nil
}

func (s *selfData) FindUrl(_ context.Context, prefix string) (*url.URL, error) {
	if fromCache, ok := s.cache.Get(prefix); ok {
		return fromCache.(*url.URL), nil
	}

	return nil, domain.ErrNotFound
}
