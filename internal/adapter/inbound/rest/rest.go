package rest

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	gv1 "github.com/iwanjunaid/hexarch/internal/adapter/inbound/rest/group/v1"
	"github.com/iwanjunaid/hexarch/internal/adapter/inbound/rest/handler"
	hv1 "github.com/iwanjunaid/hexarch/internal/adapter/inbound/rest/handler/v1"
	"github.com/iwanjunaid/hexarch/internal/core/port/inbound/registry"
	"github.com/sirupsen/logrus"
	grace "gitlab.sicepat.tech/platform/golib/httputil"
	myrouter "gitlab.sicepat.tech/platform/golib/router"
)

type RestOptions struct {
	Port            int
	GracefulTimeout time.Duration
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	ApiTimeout      int
}

type Rest struct {
	logger      *logrus.Entry
	options     *RestOptions
	listenErrCh chan error
}

func New(logger *logrus.Entry, options *RestOptions, serviceRegistry registry.ServiceRegistry) *Rest {
	root := myrouter.New(&myrouter.Options{
		Timeout: options.ApiTimeout,
	})

	root.GET("/health", handler.Health)

	v1 := myrouter.New(&myrouter.Options{
		Prefix:  "/v1",
		Timeout: options.ApiTimeout,
	})

	handlerV1 := hv1.New(logger, serviceRegistry)

	gv1.NewGroupBookV1(v1, handlerV1)

	return &Rest{
		logger:  logger,
		options: options,
	}
}

func (r *Rest) Serve() {
	bindPort := fmt.Sprintf(":%d", r.options.Port)
	r.logger.Infof("API listening on %s", bindPort)

	r.listenErrCh <- grace.Serve(bindPort, myrouter.WrapperHandler(), r.options.GracefulTimeout, r.options.ReadTimeout, r.options.WriteTimeout)
}

func (r *Rest) ListernError() <-chan error {
	return r.listenErrCh
}

func (r *Rest) SignalCheck() {
	term := make(chan os.Signal)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)

	select {
	case <-term:
		r.logger.Infoln("Exiting gracefully...")
	case err := <-r.ListernError():
		r.logger.Errorln("Error starting web server, exiting gracefully:", err)
	}
}
