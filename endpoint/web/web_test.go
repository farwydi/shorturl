package web

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/farwydi/shorturl/gateway/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestWeb(t *testing.T) {
	dm := new(mocks.DataGateway)
	h := NewRouter(NewGinEngineForTesting(), dm)

	u, _ := url.Parse("https://google.com/?q=hello+world")
	dm.On("FindUrl", mock.Anything, "foo").Return(u, nil)

	req := httptest.NewRequest("GET", "/foo", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)

	assert.Equal(t, http.StatusFound, w.Code)
	assert.Equal(t, u.String(), w.Header().Get("Location"))
}
