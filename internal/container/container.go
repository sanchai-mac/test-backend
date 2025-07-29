package container

import (
	"test-backend/internal/config"
	"test-backend/internal/infrastructure/database"
	"test-backend/internal/infrastructure/server"

	"go.uber.org/dig"
)

// Container ...
type Container struct {
	Container *dig.Container
	Error     error
}

// Configure ...
func (c *Container) Configure() {
	c.Container = dig.New()
	//Config
	if err := c.Container.Provide(config.NewConfiguration); err != nil {
		c.Error = err
	}
	//Infrastructure
	if err := c.Container.Provide(database.NewDB); err != nil {
		c.Error = err
	}
	if err := c.Container.Provide(server.NewServer); err != nil {
		c.Error = err
	}
	c.ControllerProvider()
	c.ServiceProvider()
	c.RepositoryProvider()
}

// Run...
func (c *Container) Run() *Container {
	if err := c.Container.Invoke(func(
		g *server.Server,
	) {
		if err := g.StartRestful(); err != nil {
			panic(err)
		}
	}); err != nil {
		panic(err)
	}

	return c
}

// NewContainer...
func NewContainer() *Container {
	c := &Container{}
	c.Configure()
	if c.Error != nil {
		panic(c.Error)
	}

	return c
}
