package arm2pulumi

import (
	"github.com/pulumi/pulumi-azure-nextgen-provider/provider/pkg/resources"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvert(t *testing.T) {
	metadata := resources.AzureAPIMetadata{
		Resources: map[string]resources.AzureAPIResource{
			"azure-nextgen:machinelearningservices/v20200501preview:MachineLearningDatastore": {
				Path:       "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/datastores/{datastoreName}",
				APIVersion: "2020-05-01-preview",
			},
			"azure-nextgen:sql/v20180601preview:ManagedInstance": {
				Path:       "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}",
				APIVersion: "2018-06-01-preview",
			},
			"azure-nextgen:sql/v20200202preview:ManagedInstance": {
				Path:       "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}",
				APIVersion: "2020-02-02-preview",
			},
			"azure-nextgen:documentdb/v20200301:MongoDBResourceMongoDBDatabase": {
				Path:       "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases/{databaseName}",
				APIVersion: "2020-03-01",
			},
			"azure-nextgen:documentdb/v20190801:MongoDBResourceMongoDBDatabase": {
				Path:       "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases/{databaseName}",
				APIVersion: "2019-08-01",
			},
			"azure-nextgen:documentdb:MongoDBResourceMongoDBDatabase": {
				Path:       "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases/{databaseName}",
				APIVersion: "2020-04-01",
			},
			"azure-nextgen:documentdb/v20200601preview:MongoDBResourceMongoDBCollection": {
				Path:       "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases/{databaseName}/collections/{collectionName}",
				APIVersion: "2020-06-01-preview",
			},
			"azure-nextgen:documentdb/v20200401:MongoDBResourceMongoDBCollection": {
				Path:       "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases/{databaseName}/collections/{collectionName}",
				APIVersion: "2020-04-01",
			},
			"azure-nextgen:resources/v20190801:ResourceGroup": {
				Path:       "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}",
				APIVersion: "2019-08-01",
			},
			"azure-nextgen:network/v20150501preview:PublicIpAddress": {
				Path:       "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses/{publicIpAddressName}/",
				APIVersion: "2015-05-01-preview",
			},
			"azure-nextgen:documentdb/v20200301:SqlResourceSqlStoredProcedure": {
				Path:       "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/storedProcedures/{storedProcedureName}",
				APIVersion: "2020-03-01",
			},
		},
	}

	resoureTokenConverter := newResourceTokenConverter(&metadata)

	for _, test := range []struct {
		name         string
		resourceType string
		apiVersion   string
		expected     string
	}{
		{
			name:         "SingleResourceToken",
			resourceType: "Microsoft.MachineLearningServices/workspaces/datastores",
			expected:     "azure-nextgen:machinelearningservices/v20200501preview:MachineLearningDatastore",
			apiVersion:   "2020-05-01-preview",
		},
		{
			name:         "MultipleResoureTokensFirstExactHit",
			resourceType: "Microsoft.Sql/managedInstances",
			expected:     "azure-nextgen:sql/v20180601preview:ManagedInstance",
			apiVersion:   "2018-06-01-preview",
		},
		{
			name:         "MultipleResoureTokensSecondExactHit",
			resourceType: "Microsoft.Sql/managedInstances",
			expected:     "azure-nextgen:sql/v20200202preview:ManagedInstance",
			apiVersion:   "2020-02-02-preview",
		},
		{
			name:         "MultipleResoureTokensFirstSoftMatch",
			resourceType: "Microsoft.Sql/managedInstances",
			expected:     "azure-nextgen:sql/v20180601preview:ManagedInstance",
			apiVersion:   "2018-01-01-preview",
		},
		{
			name:         "MultipleResoureTokensSecondSoftMatch",
			resourceType: "Microsoft.Sql/managedInstances",
			expected:     "azure-nextgen:sql/v20200202preview:ManagedInstance",
			apiVersion:   "2019-02-02-preview",
		},
		{
			name:         "MissingResourceType",
			resourceType: "Microsoft.Sql/managedInstances/blah",
			apiVersion:   "2019-02-02-preview",
		},
		{
			name:         "StableMultipleVersions",
			resourceType: "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases",
			expected:     "azure-nextgen:documentdb/v20200301:MongoDBResourceMongoDBDatabase",
			apiVersion:   "2020-03-01",
		},
		{
			name:         "PreviewAndStableMixed",
			resourceType: "Microsoft.DocumentDb/databaseAccounts/mongodbDatabases/collections",
			expected:     "azure-nextgen:documentdb/v20200401:MongoDBResourceMongoDBCollection",
			apiVersion:   "2020-03-01",
		},
		{
			name:         "PreviewAndStableMixedPreview",
			resourceType: "Microsoft.DocumentDb/databaseAccounts/mongodbDatabases/collections",
			expected:     "azure-nextgen:documentdb/v20200601preview:MongoDBResourceMongoDBCollection",
			apiVersion:   "2020-06-01-preview",
		},
		{
			name:         "Nested",
			resourceType: "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/storedProcedures",
			expected:     "azure-nextgen:documentdb/v20200301:SqlResourceSqlStoredProcedure",
			apiVersion:   "2020-03-01",
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, resoureTokenConverter.convert(test.resourceType, test.apiVersion), test.name)
		})
	}
}
