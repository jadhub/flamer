package flamegraph

import (
	"flamer/flamegraph/interfaces"
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/framework/web"
)

// Module struct defines the attribute value renderer module
type Module struct {
}

// Configure configures the attribute value renderer module
func (m *Module) Configure(injector *dingo.Injector) {}

type routes struct {
	flamerController *interfaces.FlamerController
}

func (r *routes) Inject(flamerController *interfaces.FlamerController) {
	r.flamerController = flamerController
}

func (r *routes) Routes(registry *web.RouterRegistry) {
	registry.HandleGet("flamer.debug.flamegraph", r.flamerController.GetFlameGraphAction)
	_, _ = registry.Route("/flamer", "flamer.debug.flamegraph")
}
