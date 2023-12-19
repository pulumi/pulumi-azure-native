// Copyright 2016-2020, Pulumi Corporation.

package convert

import (
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
)

// convertTypedObject is a callback that performs some kind of arbitrary conversion of an object for which we know its property types.
// Returning nil indicates that the given value is not of the expected type.
type convertTypedObject func(typeName string, props map[string]resources.AzureAPIProperty, values map[string]interface{}) map[string]interface{}
