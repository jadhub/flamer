package interfaces

import (
	"context"
	"flamingo.me/flamingo/v3/framework/flamingo"
	"flamingo.me/flamingo/v3/framework/web"
	"github.com/i-love-flamingo/flamer/flamer/application"
	"net/http"
)

type (
	// FlamerController controller properties
	FlamerController struct {
		responder *web.Responder
		logger    flamingo.Logger
		profiler  *application.Profiler
	}

	// Result structure of the FlamerController response
	Result struct {
		Message     string
		MessageCode string
		Success     bool
	}
)

// Inject dependencies
func (fc *FlamerController) Inject(
	responder *web.Responder,
	logger flamingo.Logger,
	profiler *application.Profiler,
) {
	fc.responder = responder
	fc.logger = logger.WithField("module", "flamer.interfaces.flamercontroller")
	fc.profiler = profiler
}

// GetFlameGraphAction returns a flame graph
func (fc *FlamerController) GetFlameGraphAction(ctx context.Context, r *web.Request) web.Result {
	result, err := fc.profiler.CPUProfile(ctx)
	if err != nil {
		return fc.responder.Data(Result{
			Message: err.Error(),
			Success: false,
		}).Status(http.StatusInternalServerError)
	}

	return fc.responder.Data(result).Status(http.StatusOK)
}
