package services

import (
	"github.com/f97/n/pkg/datastore"
	"github.com/f97/n/pkg/settings"
	"github.com/f97/n/pkg/uuid"
)

// ServiceUsingDB represents a service that need to use db
type ServiceUsingDB struct {
	container *datastore.DataStoreContainer
}

// UserDB returns the datastore which contains user
func (s *ServiceUsingDB) UserDB() *datastore.Database {
	return s.container.UserStore.Choose(0)
}

// TokenDB returns the datastore which contains user token
func (s *ServiceUsingDB) TokenDB(uid int64) *datastore.Database {
	return s.container.TokenStore.Choose(uid)
}

// UserDataDB returns the datastore which contains user data
func (s *ServiceUsingDB) UserDataDB(uid int64) *datastore.Database {
	return s.container.UserDataStore.Choose(uid)
}

// ServiceUsingConfig represents a service that need to use config
type ServiceUsingConfig struct {
	container *settings.ConfigContainer
}

// CurrentConfig returns the current config
func (s *ServiceUsingConfig) CurrentConfig() *settings.Config {
	return s.container.Current
}

// ServiceUsingUuid represents a service that need to use uuid
type ServiceUsingUuid struct {
	container *uuid.UuidContainer
}

// GenerateUuid generates a new uuid according to given uuid type
func (s *ServiceUsingUuid) GenerateUuid(uuidType uuid.UuidType) int64 {
	return s.container.GenerateUuid(uuidType)
}

// GenerateUuids generates new uuids according to given uuid type and count
func (s *ServiceUsingUuid) GenerateUuids(uuidType uuid.UuidType, count uint8) []int64 {
	return s.container.GenerateUuids(uuidType, count)
}
