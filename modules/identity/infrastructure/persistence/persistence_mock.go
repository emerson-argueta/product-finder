package persistence

// IdentityReposMock implementation
var IdentityReposMock = newMock()

// ServicesMock represents the services
type ServicesMock struct {
	User UserMock
}

func newMock() *ServicesMock {
	services := &ServicesMock{}

	return services
}
