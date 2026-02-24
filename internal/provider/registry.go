// package provider

// import (
// 	"github.com/premchand11/or/internal/config"
// 	"github.com/rs/zerolog"
// )

// type Registry struct {
// 	// will hold providers later
// }

//	func NewRegistry(cfg *config.Config, log *zerolog.Logger) *Registry {
//		return &Registry{}
//	}
package provider

import (
	"errors"
)

type Registry struct {
	providers map[string]Provider
}

func NewRegistry() *Registry {
	return &Registry{
		providers: make(map[string]Provider),
	}
}

func (r *Registry) Register(p Provider) {
	r.providers[p.Name()] = p
}

func (r *Registry) Get(name string) (Provider, error) {
	p, ok := r.providers[name]
	if !ok {
		return nil, ErrProviderNotFound
	}
	return p, nil
}

var ErrProviderNotFound = errors.New("provider not found")
