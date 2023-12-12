// Copyright 2016-2020, Pulumi Corporation.

package convert

import (
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
)

const body = "body"

// SdkShapeConverter providers functions to convert between HTTP request/response shapes and
// Pulumi SDK shapes (with flattening, renaming, etc.).
type SdkShapeConverter struct {
	Types        map[string]resources.AzureAPIType
	PartialTypes resources.PartialMap[resources.AzureAPIType]
}

func NewSdkShapeConverterPartial(ptypes resources.PartialMap[resources.AzureAPIType]) SdkShapeConverter {
	return SdkShapeConverter{
		Types:        nil,
		PartialTypes: ptypes,
	}
}

func NewSdkShapeConverterFull(types map[string]resources.AzureAPIType) SdkShapeConverter {
	return SdkShapeConverter{
		Types:        types,
		PartialTypes: resources.NewPartialMap[resources.AzureAPIType](),
	}
}

func (k *SdkShapeConverter) GetType(name string) (resources.AzureAPIType, bool, error) {
	if k.Types != nil {
		typ, ok := k.Types[name]
		if ok {
			return typ, true, nil
		}
		return resources.AzureAPIType{}, false, nil
	}

	return k.PartialTypes.Get(name)
}
