package self

import (
	"context"
	"net/url"
	"testing"

	"github.com/farwydi/shorturl/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSelf(t *testing.T) {
	var cfg domain.Config
	cfg.Gateway.Self.CacheSize = 2
	sf, err := NewSelfDataGateway(cfg)
	require.NoError(t, err)

	rawUrl := "https://www.google.com/search?q=hello"
	u1, err := url.Parse(rawUrl)
	require.NoError(t, err)

	prefix, err := sf.SaveUrl(context.TODO(), u1)
	assert.NoError(t, err)

	u2, err := sf.FindUrl(context.TODO(), prefix)
	assert.NoError(t, err)
	assert.Equal(t, u1, u2)

	assert.Equal(t, rawUrl, u2.String())

	u3, err := sf.FindUrl(context.TODO(), "bad-url")
	assert.Error(t, err, "not found")
	assert.Nil(t, u3)
}
