package gosms

import (
	"fmt"
	"sync"
)

type ProviderFactory func(cfg *SMSConfig) (SMSProvider, error)

var (
	providerRegistry = make(map[string]ProviderFactory)
	registryLock     sync.RWMutex
)

func RegisterProvider(name string, factory ProviderFactory) {
	registryLock.Lock()
	defer registryLock.Unlock()

	if name == "" {
		panic("gosms: provider name cannot be empty")
	}
	if factory == nil {
		panic("gosms: provider factory cannot be nil")
	}
	if _, exists := providerRegistry[name]; exists {
		panic("gosms: provider already registered: " + name)
	}

	providerRegistry[name] = factory
}

func NewSMSProvider(cfg *SMSConfig) (SMSProvider, error) {
	if cfg == nil {
		return nil, fmt.Errorf("gosms: SMSConfig is nil")
	}

	registryLock.RLock()
	factory, ok := providerRegistry[cfg.Provider]
	registryLock.RUnlock()

	if !ok {
		return nil, fmt.Errorf("gosms: unsupported provider: %s", cfg.Provider)
	}

	return factory(cfg)
}
