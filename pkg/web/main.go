package web

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"xedni/pkg/domain/tokenization"

	"xedni/pkg/configuration"
	"xedni/pkg/domain/document"
	"xedni/pkg/infrastructure/storage/file"
	"xedni/pkg/infrastructure/storage/memory"
	"xedni/pkg/service"
	webdocument "xedni/pkg/web/document"

	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/rs/zerolog"
)

// NewDocumentRepository instantiates a storage repository according to the configuration.
func NewDocumentRepository(ctx context.Context, cfg *configuration.AppConfiguration, logger *zerolog.Logger) (document.DocumentRepository, error) {
	switch cfg.Repository.Adapter {
	case "memory":
		return memory.NewDocumentRepository(ctx, cfg.Repository.Options, logger)
	case "file":
		return file.NewDocumentRepository(ctx, cfg.Repository.Options, logger)
	default:
		return nil, fmt.Errorf("unknown storage adapter: [%s]", cfg.Repository.Adapter)
	}
}

// NewTermRepository instantiates a storage repository according to the configuration.
func NewTermRepository(ctx context.Context, cfg *configuration.AppConfiguration, logger *zerolog.Logger) (tokenization.TermRepository, error) {
	switch cfg.Repository.Adapter {
	case "memory":
		return memory.NewTermRepository(ctx, cfg.Repository.Options, logger)
	case "file":
		return file.NewTermRepository(ctx, cfg.Repository.Options, logger)
	default:
		return nil, fmt.Errorf("unknown storage adapter: [%s]", cfg.Repository.Adapter)
	}
}

// NewIndexService fires up a Index service
func NewIndexService(d document.DocumentRepository, t tokenization.TermRepository, l *zerolog.Logger) (*service.IndexService, error) {
	return &service.IndexService{
		DocumentRepository: d,
		TermRepository:     t,
		Logger:             l,
	}, nil
}

// NewRouter creates a mux with mounted routes and instantiates respective dependencies.
func NewRouter(ctx context.Context, cfg *configuration.AppConfiguration, logger *zerolog.Logger) *chi.Mux {
	documentRepository, err := NewDocumentRepository(ctx, cfg, logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("Could not instantiate the Document repository")
	}

	termRepository, err := NewTermRepository(ctx, cfg, logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("Could not instantiate the term repository")
	}

	indexService, err := NewIndexService(documentRepository, termRepository, logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("Could not instantiate the Document service")
	}

	d := chi.NewRouter()

	d.Use(render.SetContentType(render.ContentTypeJSON))
	d.Use(chimiddleware.Heartbeat("/status"))

	d.Mount("/api", webdocument.Handler{}.Routes(logger, indexService))

	return d
}

// LaunchServer starts a web server and propagates shutdown context.
func LaunchServer(cfg *configuration.AppConfiguration, logger *zerolog.Logger) error {
	var err error

	c := make(chan os.Signal, 1)
	signal.Notify(
		c,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		s := <-c
		logger.Debug().Str("syscall", s.String()).Msg("Intercepted syscall")
		cancel()
	}()

	router := NewRouter(ctx, cfg, logger)
	srv := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%d", cfg.Port),
	}

	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal().Err(err).Msg("Could not launch the web server")
		}
	}()
	logger.Printf("Starting server on port: [%d]", cfg.Port)

	<-ctx.Done()

	logger.Printf("Cleaning up the server")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	if err = srv.Shutdown(ctxShutDown); err != nil {
		logger.Fatal().Err(err).Msg("Error on server shutdown")
	}

	cancel()

	logger.Printf("Server exited successfully")

	if err == http.ErrServerClosed {
		err = nil
	}
	return err
}
