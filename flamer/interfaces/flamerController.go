package interfaces

import (
	"context"
	"flamingo.me/flamingo/v3/framework/flamingo"
	"flamingo.me/flamingo/v3/framework/web"

	"io/ioutil"
	"net/http"
	"os"
	runtimePprof "runtime/pprof"
	"time"
)

type (
	FlamerController struct {
		responder *web.Responder
		logger    flamingo.Logger
	}

	Result struct {
		Message     string
		MessageCode string
		Success     bool
	}

	CPUProfile struct {
		Duration time.Duration // 30 seconds by default
	}
)

func (cc *FlamerController) Inject(
	responder *web.Responder,
	logger flamingo.Logger,
) {
	cc.responder = responder
	cc.logger = logger.WithField("module", "flamer.flamegraph")
}

func (cc *FlamerController) GetFlameGraphAction(ctx context.Context, r *web.Request) web.Result {
	p := &CPUProfile{}

	file, err := p.Capture(cc)
	if err != nil {
		return cc.responder.Data(Result{
			Message: err.Error(),
			Success: false,
		}).Status(http.StatusInternalServerError)
	}

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return cc.responder.Data(Result{
			Message: err.Error(),
			Success: false,
		}).Status(http.StatusInternalServerError)
	}

	return cc.responder.Data(data).Status(http.StatusOK)
}

func (p CPUProfile) Capture(cc *FlamerController) (string, error) {
	dur := p.Duration
	if dur == 0 {
		dur = 30 * time.Second
	}

	f := cc.newTemp()
	if err := runtimePprof.StartCPUProfile(f); err != nil {
		return "", nil
	}
	time.Sleep(dur)
	runtimePprof.StopCPUProfile()
	if err := f.Close(); err != nil {
		return "", nil
	}
	return f.Name(), nil
}

func (cc *FlamerController) newTemp() (f *os.File) {
	f, err := ioutil.TempFile("", "profile-")
	if err != nil {
		cc.logger.Fatal("Cannot create new temp profile file: %v", err)
	}
	return f
}
