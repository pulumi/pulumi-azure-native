


package v20220815

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure-native/sdk/go/azure"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "azure-native:documentdb/v20220815:CassandraCluster":
		r = &CassandraCluster{}
	case "azure-native:documentdb/v20220815:CassandraDataCenter":
		r = &CassandraDataCenter{}
	case "azure-native:documentdb/v20220815:CassandraResourceCassandraKeyspace":
		r = &CassandraResourceCassandraKeyspace{}
	case "azure-native:documentdb/v20220815:CassandraResourceCassandraTable":
		r = &CassandraResourceCassandraTable{}
	case "azure-native:documentdb/v20220815:DatabaseAccount":
		r = &DatabaseAccount{}
	case "azure-native:documentdb/v20220815:GremlinResourceGremlinDatabase":
		r = &GremlinResourceGremlinDatabase{}
	case "azure-native:documentdb/v20220815:GremlinResourceGremlinGraph":
		r = &GremlinResourceGremlinGraph{}
	case "azure-native:documentdb/v20220815:MongoDBResourceMongoDBCollection":
		r = &MongoDBResourceMongoDBCollection{}
	case "azure-native:documentdb/v20220815:MongoDBResourceMongoDBDatabase":
		r = &MongoDBResourceMongoDBDatabase{}
	case "azure-native:documentdb/v20220815:MongoDBResourceMongoRoleDefinition":
		r = &MongoDBResourceMongoRoleDefinition{}
	case "azure-native:documentdb/v20220815:MongoDBResourceMongoUserDefinition":
		r = &MongoDBResourceMongoUserDefinition{}
	case "azure-native:documentdb/v20220815:NotebookWorkspace":
		r = &NotebookWorkspace{}
	case "azure-native:documentdb/v20220815:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:documentdb/v20220815:Service":
		r = &Service{}
	case "azure-native:documentdb/v20220815:SqlResourceSqlContainer":
		r = &SqlResourceSqlContainer{}
	case "azure-native:documentdb/v20220815:SqlResourceSqlDatabase":
		r = &SqlResourceSqlDatabase{}
	case "azure-native:documentdb/v20220815:SqlResourceSqlRoleAssignment":
		r = &SqlResourceSqlRoleAssignment{}
	case "azure-native:documentdb/v20220815:SqlResourceSqlRoleDefinition":
		r = &SqlResourceSqlRoleDefinition{}
	case "azure-native:documentdb/v20220815:SqlResourceSqlStoredProcedure":
		r = &SqlResourceSqlStoredProcedure{}
	case "azure-native:documentdb/v20220815:SqlResourceSqlTrigger":
		r = &SqlResourceSqlTrigger{}
	case "azure-native:documentdb/v20220815:SqlResourceSqlUserDefinedFunction":
		r = &SqlResourceSqlUserDefinedFunction{}
	case "azure-native:documentdb/v20220815:TableResourceTable":
		r = &TableResourceTable{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := azure.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"azure-native",
		"documentdb/v20220815",
		&module{version},
	)
}
