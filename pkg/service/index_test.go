package service_test

import (
	"flag"
	"path/filepath"
	"testing"
	"xedni/pkg/configuration"
	"xedni/pkg/service"
	"xedni/pkg/web"
)

const TEXT1 = "An interactive introduction to Go in three sections The first section covers basic syntax and data structures the second discusses methods and interfaces and the third introduces Go's concurrency primitives Each section concludes with a few exercises so you can practice what you've learned You can take the tour online or install it locally with"

const TEXT2 = "An interactive introduction to Go in three sections The first section covers basic syntax and data structures the second discusses methods and interfaces and the third introduces Go's concurrency primitives Each section concludes with a few exercises so you can practice what you've learned You can take the tour online or install it locally with An interactive introduction to Go in three sections The first section covers basic syntax and data structures the second discusses methods and interfaces and the third introduces Go's concurrency primitives Each section concludes with a few exercises so you can practice what you've learned You can take the tour online or install it locally with"

var configPath = flag.String("config-path", "", "Directory for the specific environment yaml configuration file.")

func initializeIndexService() (*service.IndexService, error) {
	appConfiguration := &configuration.AppConfiguration{}
	absoluteConfigPath, err := filepath.Abs(*configPath)
	if err != nil {
		return nil, err
	}

	err = configuration.LoadYAML(appConfiguration, &absoluteConfigPath, nil, nil)
	if err != nil {
		return nil, err
	}

	d, err := web.NewDocumentRepository(nil, appConfiguration, nil)
	if err != nil {
		return nil, err
	}

	t, err := web.NewTermRepository(nil, appConfiguration, nil)
	if err != nil {
		return nil, err
	}

	return web.NewIndexService(d, t, nil)
}

func benchmarkIndex(text string, b *testing.B) {
	is, err := initializeIndexService()
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		is.Index(text)
	}
}

func benchmarkSearch(words []string, b *testing.B) {
	is, err := initializeIndexService()
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		is.Search(words)
	}
}

func BenchmarkIndexText1(b *testing.B)   { benchmarkIndex(TEXT1, b) }
func BenchmarkIndexText2(b *testing.B)   { benchmarkIndex(TEXT2, b) }
func BenchmarkIndexSearch1(b *testing.B) { benchmarkSearch([]string{"online", "Go"}, b) }
func BenchmarkIndexSearch2(b *testing.B) {
	benchmarkSearch([]string{"interactive", "introduction", "to", "Go", "in", "three", "sections", "The", "first"}, b)
}
