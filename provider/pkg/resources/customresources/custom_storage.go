// Copyright 2021, Pulumi Corporation.  All rights reserved.

package customresources

import (
	"regexp"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

// Resource property names.
const (
	staticWebsitePath = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/staticWebsite"
	subscriptionId    = "subscriptionId"
	resourceGroupName = "resourceGroupName"
	accountName       = "accountName"
	containerName     = "containerName"
	indexDocument     = "indexDocument"
	error404Document  = "error404Document"
)

// Resource property names.
const (
	blobPath    = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers/{containerName}/blobs/{blobName}"
	accessTier  = "accessTier"
	blobName    = "blobName"
	contentMd5  = "contentMd5"
	contentType = "contentType"
	metadata    = "metadata"
	nameProp    = "name"
	source      = "source"
	typeProp    = "type"
	url         = "url"
)

var storageAccountPathRegexStr = `(?i)^/subscriptions/(.+)/resourceGroups/(.+)/providers/Microsoft.Storage/storageAccounts/(.+?)/`
var storageAccountPathRegex = regexp.MustCompile(storageAccountPathRegexStr)
var blobIDPattern = regexp.MustCompile(storageAccountPathRegexStr + `blobServices/default/containers/(.+)/blobs/(.+)$`)

// parseBlobIdProperties parses an ID of a Blob resource to its identified properties.
// For instance, it will convert
// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myrg/providers/Microsoft.Storage/storageAccounts/mysa/blobServices/default/containers/myc/blobs/log.txt
// to a map of
// resourceGroupName=myrg,accountName=mysa,containerName=myc,blobName=log.txt.
func parseBlobIdProperties(id string) (resource.PropertyMap, bool) {
	match := blobIDPattern.FindStringSubmatch(id)
	if len(match) != 6 {
		return nil, false
	}

	clientProperties := resource.PropertyMap{}
	clientProperties[subscriptionId] = resource.NewStringProperty(match[1])
	clientProperties[resourceGroupName] = resource.NewStringProperty(match[2])
	clientProperties[accountName] = resource.NewStringProperty(match[3])
	clientProperties[containerName] = resource.NewStringProperty(match[4])
	clientProperties[blobName] = resource.NewStringProperty(match[5])
	return clientProperties, true
}

func readAssetBytes(s resource.PropertyValue) ([]byte, error) {
	switch {
	case s.IsAsset():
		asset := s.AssetValue()
		assetBytes, err := asset.Bytes()
		if err != nil {
			return nil, errors.Wrap(err, "reading asset bytes")
		}
		return assetBytes, nil
	case s.IsArchive():
		archive := s.ArchiveValue()
		archiveBytes, err := archive.Bytes(resource.ZIPArchive)
		if err != nil {
			return nil, errors.Wrap(err, "reading archive bytes")
		}
		return archiveBytes, nil
	default:
		return nil, errors.Errorf("unknown asset type %q", s.TypeString())
	}
}
