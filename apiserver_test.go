package apiserver

import (
	"context"
	"github.com/TonyPath/apiserver/router"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestRun(t *testing.T) {

	r := router.NewApiRouter()

	apiSrv, err := New(r, WithPort(8080), WithReadTimeout(1*time.Second))
	require.NoError(t, err)
	require.NotNil(t, apiSrv)

	ctx, cnl := context.WithCancel(context.TODO())

	var errSrv error
	guard := make(chan struct{}, 1)
	go func() {
		errSrv = apiSrv.Run(ctx)
		guard <- struct{}{}
	}()

	time.Sleep(5 * time.Millisecond)
	cnl()
	<-guard

	require.ErrorIs(t, errSrv, context.Canceled)
}
