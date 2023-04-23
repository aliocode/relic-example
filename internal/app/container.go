package app

import (
	"context"
	"net/http"

	"github.com/aliocode/relic-example/internal/api/httphandler"
	"github.com/aliocode/relic-example/internal/api/httpserver"
	"github.com/newrelic/go-agent/v3/newrelic"
	"go.uber.org/zap"
)

type Container struct {
	ctx    context.Context
	cancel context.CancelFunc
	log    *zap.Logger
	cfg    Config
	once   once

	relic      *newrelic.Application
	httpserver http.Handler
}

type once struct{}

func New() (*Container, error) {
	cfg, err := NewConfig()
	if err != nil {
		return nil, err
	}

	log, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())
	_ = cancel

	relic, err := newrelic.NewApplication(
		newrelic.ConfigAppName(cfg.RelicAppName),
		newrelic.ConfigLicense(cfg.RelicLicense),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if err != nil {
		return nil, err
	}

	server := httpserver.New(httphandler.New(), relic)

	return &Container{
		ctx:        ctx,
		cancel:     cancel,
		log:        log,
		cfg:        cfg,
		once:       once{},
		relic:      relic,
		httpserver: server,
	}, nil
}

func (s *Container) Run() error {
	addr := ":" + s.cfg.HttpPort
	if err := http.ListenAndServe(addr, s.httpserver); err != nil {
		return err
	}
	return nil
}

func (s *Container) Close() {
	s.cancel()
}
