package controller

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tucnak/climax"
	"github.com/vsco/dcdr/config"
	"github.com/vsco/dcdr/models"
)

const (
	Error   = 1
	Success = 0
)

type MockClient struct {
	Features models.Features
	Feature  *models.Feature
	Error    error
}

func NewMockClient(f *models.Feature, fts models.Features, err error) (m *MockClient) {
	m = &MockClient{
		Features: fts,
		Feature:  f,
		Error:    err,
	}

	return
}

func (m *MockClient) Get(key string, v interface{}) error {
	return m.Error
}

func (m *MockClient) Set(ft *models.Feature) error {
	return m.Error
}

func (m *MockClient) Delete(key string, scope string) error {
	return m.Error
}

func (m *MockClient) Namespace() string {
	return "dcdr"
}

func (m *MockClient) List(prefix string, scope string) (models.Features, error) {
	return m.Features, m.Error
}

func (m *MockClient) GetInfo() (*models.Info, error) {
	return nil, m.Error
}

func (m *MockClient) Commit(ft *models.Feature, deleted bool) error {
	return m.Error
}

func (m *MockClient) Push() error {
	return m.Error
}

func (m *MockClient) UpdateCurrentSHA() (string, error) {
	return "", m.Error
}

func (m *MockClient) InitRepo(create bool) error {
	return m.Error
}

func (m *MockClient) Watch() {}

func TestListEmptyFeatures(t *testing.T) {
	cfg := config.DefaultConfig()
	c := NewMockClient(nil, nil, nil)
	ctl := New(cfg, c)

	ctx := climax.Context{
		Variable: map[string]string{},
	}

	code := ctl.List(ctx)

	assert.Equal(t, Error, code)
}

func TestListFeatures(t *testing.T) {
	cfg := config.DefaultConfig()
	fts := models.Features{
		models.Feature{
			Key:   "test",
			Value: true,
		},
	}
	c := NewMockClient(nil, fts, nil)
	ctl := New(cfg, c)

	ctx := climax.Context{
		Variable: map[string]string{},
	}

	code := ctl.List(ctx)

	assert.Equal(t, Success, code)
}

func TestSet(t *testing.T) {
	cfg := config.DefaultConfig()
	fts := models.Features{
		models.Feature{
			Key:   "test",
			Value: true,
		},
	}
	c := NewMockClient(nil, fts, nil)
	ctl := New(cfg, c)

	ctx := climax.Context{
		Variable: map[string]string{"name": "null-test"},
	}

	code := ctl.Set(ctx)

	assert.Equal(t, Success, code)
}
