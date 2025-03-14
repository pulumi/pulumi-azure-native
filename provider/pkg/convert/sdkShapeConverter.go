// Copyright 2016-2020, Pulumi Corporation.

package convert

import (
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
)

const body = "body"

// SdkShapeConverter providers functions to convert between HTTP request/response shapes and
// Pulumi SDK shapes (with flattening, renaming, etc.).
type SdkShapeConverter struct {
	Types resources.MapLike[resources.AzureAPIType]
}

func NewSdkShapeConverterPartial(ptypes resources.MapLike[resources.AzureAPIType]) SdkShapeConverter {
	return SdkShapeConverter{
		Types: ptypes,
	}
}

func NewSdkShapeConverterFull(types map[string]resources.AzureAPIType) SdkShapeConverter {
	return SdkShapeConverter{
		Types: resources.GoMap[resources.AzureAPIType](types),
	}
}

func (k *SdkShapeConverter) GetType(name string) (resources.AzureAPIType, bool, error) {
	if k.Types == nil {
		return resources.AzureAPIType{}, false, nil
	}
	return k.Types.Get(name)
}
