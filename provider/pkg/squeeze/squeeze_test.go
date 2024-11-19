package squeeze

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApiVersionToDate(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		apiVersion := "v20200101"
		date, err := apiVersionToDate(apiVersion)
		assert.NoError(t, err)
		expected := "2020-01-01"
		actual := date.Format("2006-01-02")
		assert.Equal(t, expected, actual)
	})

	t.Run("preview", func(t *testing.T) {
		apiVersion := "v20200101preview"
		date, err := apiVersionToDate(apiVersion)
		assert.NoError(t, err)
		expected := "2020-01-01"
		actual := date.Format("2006-01-02")
		assert.Equal(t, expected, actual)
	})
}

func TestSortApiVersions(t *testing.T) {
	t.Run("already ordered", func(t *testing.T) {
		versions := []string{"v20200101", "v20210202"}
		sortApiVersions(versions)
		expected := []string{"v20200101", "v20210202"}
		assert.Equal(t, expected, versions)
	})

	t.Run("reversed", func(t *testing.T) {
		versions := []string{"v20210202", "v20200101"}
		sortApiVersions(versions)
		expected := []string{"v20200101", "v20210202"}
		assert.Equal(t, expected, versions)
	})

	t.Run("preview comes before stable", func(t *testing.T) {
		versions := []string{"v20200101", "v20200101preview"}
		sortApiVersions(versions)
		expected := []string{"v20200101preview", "v20200101"}
		assert.Equal(t, expected, versions)
	})

	t.Run("private comes before preview", func(t *testing.T) {
		versions := []string{"v20200101preview", "v20200101privatepreview"}
		sortApiVersions(versions)
		expected := []string{"v20200101privatepreview", "v20200101preview"}
		assert.Equal(t, expected, versions)
	})
}
