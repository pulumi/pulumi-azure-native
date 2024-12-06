package openapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApiVersionToDate(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		date, err := ApiVersionToDate("2020-01-01")
		assert.NoError(t, err)
		actual := date.Format("2006-01-02")
		assert.Equal(t, "2020-01-01", actual)
	})

	t.Run("preview", func(t *testing.T) {
		date, err := ApiVersionToDate("2020-01-01-preview")
		assert.NoError(t, err)
		actual := date.Format("2006-01-02")
		assert.Equal(t, "2020-01-01", actual)
	})
}

func TestSortApiVersions(t *testing.T) {
	t.Run("already ordered", func(t *testing.T) {
		versions := []ApiVersion{"2020-01-01", "2021-02-02"}
		SortApiVersions(versions)
		expected := []ApiVersion{"2020-01-01", "2021-02-02"}
		assert.Equal(t, expected, versions)
	})

	t.Run("reversed", func(t *testing.T) {
		versions := []ApiVersion{"2021-02-02", "2020-01-01"}
		SortApiVersions(versions)
		expected := []ApiVersion{"2020-01-01", "2021-02-02"}
		assert.Equal(t, expected, versions)
	})

	t.Run("preview comes before stable", func(t *testing.T) {
		versions := []ApiVersion{"2020-01-01", "2020-01-01-preview"}
		SortApiVersions(versions)
		expected := []ApiVersion{"2020-01-01-preview", "2020-01-01"}
		assert.Equal(t, expected, versions)
	})

	t.Run("private comes before preview", func(t *testing.T) {
		versions := []ApiVersion{"2020-01-01-preview", "2020-01-01-privatepreview"}
		SortApiVersions(versions)
		expected := []ApiVersion{"2020-01-01-privatepreview", "2020-01-01-preview"}
		assert.Equal(t, expected, versions)
	})
}
