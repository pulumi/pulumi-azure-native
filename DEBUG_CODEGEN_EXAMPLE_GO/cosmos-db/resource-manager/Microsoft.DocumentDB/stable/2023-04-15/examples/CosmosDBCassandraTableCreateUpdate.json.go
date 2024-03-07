package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/documentdb/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := documentdb.NewCassandraResourceCassandraTable(ctx, "cassandraResourceCassandraTable", &documentdb.CassandraResourceCassandraTableArgs{
AccountName: pulumi.String("ddb1"),
KeyspaceName: pulumi.String("keyspaceName"),
Location: pulumi.String("West US"),
Options: nil,
Resource: documentdb.CassandraTableGetPropertiesResponseResource{
DefaultTtl: pulumi.Int(100),
Id: pulumi.String("tableName"),
Schema: interface{}{
ClusterKeys: documentdb.ClusterKeyArray{
&documentdb.ClusterKeyArgs{
Name: pulumi.String("columnA"),
OrderBy: pulumi.String("Asc"),
},
},
Columns: documentdb.ColumnArray{
&documentdb.ColumnArgs{
Name: pulumi.String("columnA"),
Type: pulumi.String("Ascii"),
},
},
PartitionKeys: documentdb.CassandraPartitionKeyArray{
&documentdb.CassandraPartitionKeyArgs{
Name: pulumi.String("columnA"),
},
},
},
},
ResourceGroupName: pulumi.String("rg1"),
TableName: pulumi.String("tableName"),
Tags: nil,
})
if err != nil {
return err
}
return nil
})
}
