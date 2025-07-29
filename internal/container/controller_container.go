package container

import "test-backend/internal/controller"

func (c *Container) ControllerProvider() {
	if err := c.Container.Provide(controller.NewUserController); err != nil {
		c.Error = err
	}
}
