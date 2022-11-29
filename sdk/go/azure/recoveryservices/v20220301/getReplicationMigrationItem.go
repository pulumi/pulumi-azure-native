


package v20220301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupReplicationMigrationItem(ctx *pulumi.Context, args *LookupReplicationMigrationItemArgs, opts ...pulumi.InvokeOption) (*LookupReplicationMigrationItemResult, error) {
	var rv LookupReplicationMigrationItemResult
	err := ctx.Invoke("azure-native:recoveryservices/v20220301:getReplicationMigrationItem", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReplicationMigrationItemArgs struct {
	FabricName              string `pulumi:"fabricName"`
	MigrationItemName       string `pulumi:"migrationItemName"`
	ProtectionContainerName string `pulumi:"protectionContainerName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
	ResourceName            string `pulumi:"resourceName"`
}


type LookupReplicationMigrationItemResult struct {
	Id         string                          `pulumi:"id"`
	Location   *string                         `pulumi:"location"`
	Name       string                          `pulumi:"name"`
	Properties MigrationItemPropertiesResponse `pulumi:"properties"`
	Type       string                          `pulumi:"type"`
}

func LookupReplicationMigrationItemOutput(ctx *pulumi.Context, args LookupReplicationMigrationItemOutputArgs, opts ...pulumi.InvokeOption) LookupReplicationMigrationItemResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReplicationMigrationItemResult, error) {
			args := v.(LookupReplicationMigrationItemArgs)
			r, err := LookupReplicationMigrationItem(ctx, &args, opts...)
			var s LookupReplicationMigrationItemResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReplicationMigrationItemResultOutput)
}

type LookupReplicationMigrationItemOutputArgs struct {
	FabricName              pulumi.StringInput `pulumi:"fabricName"`
	MigrationItemName       pulumi.StringInput `pulumi:"migrationItemName"`
	ProtectionContainerName pulumi.StringInput `pulumi:"protectionContainerName"`
	ResourceGroupName       pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName            pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupReplicationMigrationItemOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationMigrationItemArgs)(nil)).Elem()
}


type LookupReplicationMigrationItemResultOutput struct{ *pulumi.OutputState }

func (LookupReplicationMigrationItemResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationMigrationItemResult)(nil)).Elem()
}

func (o LookupReplicationMigrationItemResultOutput) ToLookupReplicationMigrationItemResultOutput() LookupReplicationMigrationItemResultOutput {
	return o
}

func (o LookupReplicationMigrationItemResultOutput) ToLookupReplicationMigrationItemResultOutputWithContext(ctx context.Context) LookupReplicationMigrationItemResultOutput {
	return o
}

func (o LookupReplicationMigrationItemResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationMigrationItemResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupReplicationMigrationItemResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicationMigrationItemResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupReplicationMigrationItemResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationMigrationItemResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupReplicationMigrationItemResultOutput) Properties() MigrationItemPropertiesResponseOutput {
	return o.ApplyT(func(v LookupReplicationMigrationItemResult) MigrationItemPropertiesResponse { return v.Properties }).(MigrationItemPropertiesResponseOutput)
}

func (o LookupReplicationMigrationItemResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationMigrationItemResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReplicationMigrationItemResultOutput{})
}
