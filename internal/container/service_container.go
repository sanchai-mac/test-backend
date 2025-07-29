package container

import "test-backend/internal/service"

func (c *Container) ServiceProvider() {
	if err := c.Container.Provide(service.NewUserService); err != nil {
		c.Error = err
	}
}
