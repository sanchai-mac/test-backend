package container

import "test-backend/internal/repository"

func (c *Container) RepositoryProvider() {

	if err := c.Container.Provide(repository.NewUserRepository); err != nil {
		c.Error = err
	}
}
