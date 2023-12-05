package gen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultsForStorageAccount(t *testing.T) {
	def := propertyDefaults("storage", "StorageAccount")
	assert.NotNil(t, def)

	assert.Contains(t, def, "networkRuleSet")
	assert.Equal(t, 1, len(def))
}

func TestDefaultsAreSpecificToStorageAccount(t *testing.T) {
	assert.NotNil(t, propertyDefaults("storage", "StorageAccount"))
	assert.NotNil(t, propertyDefaults("storage/v20210501", "StorageAccount"))

	assert.Nil(t, propertyDefaults("storagecache", "StorageAccount"))

	assert.Nil(t, propertyDefaults("storage", "StorageSomething"))
}
