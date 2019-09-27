package flamer

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/framework/web"
	"github.com/i-love-flamingo/flamer/flamer/application"
	"github.com/i-love-flamingo/flamer/flamer/domain"
	"github.com/i-love-flamingo/flamer/flamer/interfaces"
)

// Module struct defines the attribute value renderer module
type Module struct {
}

// Configure configures the attribute value renderer module
func (m *Module) Configure(injector *dingo.Injector) {
	injector.Bind((*domain.Profiler)(nil)).To(application.Profiler{})
	web.BindRoutes(injector, new(routes))
}

type routes struct {
	flamerController *interfaces.FlamerController
}

func (r *routes) Inject(flamerController *interfaces.FlamerController) {
	r.flamerController = flamerController
}

func (r *routes) Routes(registry *web.RouterRegistry) {
	registry.HandleGet("flamer.flamegraph", r.flamerController.GetFlameGraphAction)
	_, _ = registry.Route("/flamer/flamegraph", "flamer.flamegraph")
}
