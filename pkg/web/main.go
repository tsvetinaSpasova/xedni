package web

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

// NewDocumentService fires up a Document service
func NewDocumentService(r document.DocumentRepository, l *zerolog.Logger) (*service.DocumentService, error) {
	return &service.DocumentService{
		Repository: r,
		Logger:     l,
	}, nil
}

// NewRouter creates a mux with mounted routes and instantiates respective dependencies.
func NewRouter(ctx context.Context, cfg *configuration.AppConfiguration, logger *zerolog.Logger) *chi.Mux {
	documentRepository, err := NewDocumentRepository(ctx, cfg, logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("Could not instantiate the Document repository")
	}

	documentService, err := NewDocumentService(documentRepository, logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("Could not instantiate the Document service")
	}

	r := chi.NewRouter()

	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(chimiddleware.Heartbeat("/status"))

	r.Mount("/api", webdocument.Handler{}.Routes(logger, documentService))

	return r
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
