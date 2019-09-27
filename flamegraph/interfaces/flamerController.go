package interfaces

import (
	"context"
	"flamingo.me/flamingo/v3/framework/web"
	"fmt"
	gpprof "github.com/google/pprof"
	"github.com/google/pprof/driver"
	"net/http/pprof"
	"os"
)

type (
	FlamerController struct {
	}
)

func (cc *FlamerController) GetFlameGraphAction(ctx context.Context, r *web.Request) web.Result {

}
