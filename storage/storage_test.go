package storage

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitSelectedStorage_InMemory(t *testing.T) {
  storage_type := os.Getenv("STORAGE_TYPE")
  defer os.Setenv("STORAGE_TYPE", storage_type)

  os.Setenv("STORAGE_TYPE", "in_memory")
  err := InitSelectedStorage()
  assert.Nil(t, err)
  assert.IsType(t, &InMemoryStorage{}, SelectedStorage)
  assert.Equal(t, reflect.TypeOf(&InMemoryStorage{}), reflect.TypeOf(SelectedStorage))
}

func TestInitSelectedStorage_Postgres(t *testing.T) {
  storage_type := os.Getenv("STORAGE_TYPE")
  defer os.Setenv("STORAGE_TYPE", storage_type)

  os.Setenv("STORAGE_TYPE", "postgres")
  err := InitSelectedStorage()
  assert.Nil(t, err)
  assert.IsType(t, &PostgresStorage{}, SelectedStorage)
  assert.Equal(t, reflect.TypeOf(&PostgresStorage{}), reflect.TypeOf(SelectedStorage))
}
