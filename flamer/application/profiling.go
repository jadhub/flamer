package application

import (
	"context"
	"flamingo.me/flamingo/v3/framework/flamingo"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
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
		Duration time.Duration // 30 seconds by default
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

	runtime.SetCPUProfileRate(500)

	dur := cpuProfile.Duration
	if dur == 0 {
		dur = 30 * time.Second
	}

	f := p.newTmpFile()

	err := runtimePprof.StartCPUProfile(f)
	if err != nil {
		return nil, nil
	}

	time.Sleep(dur)

	runtimePprof.StopCPUProfile()

	err = f.Close()
	if err != nil {
		return nil, nil
	}

	return p.getFileContent(f)
}

func (p *Profiler) getFileContent(f *os.File) ([]byte, error) {
	p.logger.Debugf(fmt.Sprintf("profile filename: %s", f.Name()))

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
