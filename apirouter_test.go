package apiserver

import (
	"context"
	"github.com/TonyPath/apiserver/router"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestNew(t *testing.T) {

	r := router.NewApiRouter()
	r.Handle(http.MethodGet, "/", func(writer http.ResponseWriter, r *http.Request) {

	})

	apiSrv, err := New(r)
	require.NoError(t, err)
	require.NotNil(t, apiSrv)

	ctx, cnl := context.WithCancel(context.TODO())

	guard := make(chan struct{}, 1)
	go func() {
		_ = apiSrv.Run(ctx)
		guard <- struct{}{}
	}()

	time.Sleep(5 * time.Millisecond)
	cnl()

	<-guard

}
