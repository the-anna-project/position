// Package position implements Service to manage positions of the connection
// space.
package position

import (
	"strconv"
	"strings"
	"sync"

	"github.com/the-anna-project/random"
)

// ServiceConfig represents the configuration used to create a new position
// service.
type ServiceConfig struct {
	// Dependencies.
	RandomService random.Service

	// Settings.
	DimensionCount int
	DimensionDepth int
}

// DefaultServiceConfig provides a default configuration to create a new
// position service by best effort.
func DefaultServiceConfig() ServiceConfig {
	var err error

	var randomService random.Service
	{
		randomConfig := random.DefaultServiceConfig()
		randomService, err = random.NewService(randomConfig)
		if err != nil {
			panic(err)
		}
	}

	config := ServiceConfig{
		// Dependencies.
		RandomService: randomService,

		// Settings.
		DimensionCount: 0,
		DimensionDepth: 0,
	}

	return config
}

// NewService creates a new position service.
func NewService(config ServiceConfig) (Service, error) {
	// Dependencies.
	if config.RandomService == nil {
		return nil, maskAnyf(invalidConfigError, "random service must not be empty")
	}

	// Settings.
	if config.DimensionCount == 0 {
		return nil, maskAnyf(invalidConfigError, "dimension count must not be empty")
	}
	if config.DimensionDepth == 0 {
		return nil, maskAnyf(invalidConfigError, "dimension depth must not be empty")
	}

	newService := &service{
		// Dependencies.
		random: config.RandomService,

		// Internals.
		bootOnce:     sync.Once{},
		closer:       make(chan struct{}, 1),
		shutdownOnce: sync.Once{},

		// Settings.
		dimensionCount: config.DimensionCount,
		dimensionDepth: config.DimensionDepth,
	}

	return newService, nil
}

type service struct {
	// Dependencies.
	random random.Service

	// Internals.
	bootOnce     sync.Once
	closer       chan struct{}
	shutdownOnce sync.Once

	// Settings.
	dimensionCount int
	dimensionDepth int
}

func (s *service) Boot() {
	s.bootOnce.Do(func() {
		// Service specific boot logic goes here.
	})
}

func (s *service) Default() (string, error) {
	nums, err := s.random.CreateNMax(s.dimensionCount, s.dimensionDepth)
	if err != nil {
		return "", maskAny(err)
	}

	coordinates := []string{}
	for _, n := range nums {
		coordinates = append(coordinates, strconv.Itoa(n))
	}
	position := strings.Join(coordinates, ",")

	return position, nil
}

func (s *service) Shutdown() {
	s.shutdownOnce.Do(func() {
		close(s.closer)
	})
}
