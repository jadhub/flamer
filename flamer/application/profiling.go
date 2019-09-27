package application

import (
	"context"
	"flamingo.me/flamingo/v3/framework/flamingo"
	"io/ioutil"
	"os"
	runtimePprof "runtime/pprof"
	"time"
)

type (
	// Profiler logic
	Profiler struct {
		logger flamingo.Logger
	}

	// CPUProfile Data
	CPUProfile struct {
		Duration time.Duration // 5 seconds by default
	}
)

// Inject dependencies
func (p *Profiler) Inject(
	logger flamingo.Logger,
) {
	p.logger = logger.WithField("module", "flamer.application.profiling")
}

// CPUProfile method
func (p *Profiler) CPUProfile(ctx context.Context) ([]byte, error) {
	data, err := p.cpuCapture()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (p *Profiler) cpuCapture() ([]byte, error) {
	cpuProfile := new(CPUProfile)

	dur := cpuProfile.Duration
	if dur == 0 {
		dur = 5 * time.Second
	}

	f := p.newTmpFile()
	if err := runtimePprof.StartCPUProfile(f); err != nil {
		return nil, nil
	}

	time.Sleep(dur)

	runtimePprof.StopCPUProfile()

	if err := f.Close(); err != nil {
		return nil, nil
	}
	return p.getFileContent(f)
}

func (p *Profiler) getFileContent(f *os.File) ([]byte, error) {
	data, err := ioutil.ReadFile(f.Name())
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (p *Profiler) newTmpFile() (f *os.File) {
	f, err := ioutil.TempFile("", "profile-")
	if err != nil {
		p.logger.Fatal("Cannot create new temp profile file: %v", err)
	}

	return f
}
