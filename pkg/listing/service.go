package listitng

import (
	sdk_v1alpha1 "knative.dev/client/pkg/serving/v1alpha1"
	"knative.dev/serving/pkg/apis/serving/v1alpha1"
)

// Repository provides access to the Serving storage.
type Repository interface {
	GetServing(client sdk_v1alpha1.KnServingClient, serviceName string) (*v1alpha1.Service, error)
	GetAllServing(client sdk_v1alpha1.KnServingClient) (*v1alpha1.ServiceList, error)
}

// Service provides serving listing operations.
type Service interface {
	GetServing(client sdk_v1alpha1.KnServingClient, serviceName string) (*v1alpha1.Service, error)
	GetAllServing(client sdk_v1alpha1.KnServingClient) (*v1alpha1.ServiceList, error)
}

type service struct {
	r Repository
}

// NewService creates a listing service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetServing(client sdk_v1alpha1.KnServingClient, serviceName string) (*v1alpha1.Service, error) {
	return s.r.GetServing(client, serviceName)
}

func (s *service) GetAllServing(client sdk_v1alpha1.KnServingClient) (*v1alpha1.ServiceList, error) {
	return s.r.GetAllServing(client)
}
