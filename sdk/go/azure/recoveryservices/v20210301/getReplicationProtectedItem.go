


package v20210301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupReplicationProtectedItem(ctx *pulumi.Context, args *LookupReplicationProtectedItemArgs, opts ...pulumi.InvokeOption) (*LookupReplicationProtectedItemResult, error) {
	var rv LookupReplicationProtectedItemResult
	err := ctx.Invoke("azure-native:recoveryservices/v20210301:getReplicationProtectedItem", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReplicationProtectedItemArgs struct {
	FabricName                  string `pulumi:"fabricName"`
	ProtectionContainerName     string `pulumi:"protectionContainerName"`
	ReplicatedProtectedItemName string `pulumi:"replicatedProtectedItemName"`
	ResourceGroupName           string `pulumi:"resourceGroupName"`
	ResourceName                string `pulumi:"resourceName"`
}


type LookupReplicationProtectedItemResult struct {
	Id         string                                     `pulumi:"id"`
	Location   *string                                    `pulumi:"location"`
	Name       string                                     `pulumi:"name"`
	Properties ReplicationProtectedItemPropertiesResponse `pulumi:"properties"`
	Type       string                                     `pulumi:"type"`
}

func LookupReplicationProtectedItemOutput(ctx *pulumi.Context, args LookupReplicationProtectedItemOutputArgs, opts ...pulumi.InvokeOption) LookupReplicationProtectedItemResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReplicationProtectedItemResult, error) {
			args := v.(LookupReplicationProtectedItemArgs)
			r, err := LookupReplicationProtectedItem(ctx, &args, opts...)
			var s LookupReplicationProtectedItemResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReplicationProtectedItemResultOutput)
}

type LookupReplicationProtectedItemOutputArgs struct {
	FabricName                  pulumi.StringInput `pulumi:"fabricName"`
	ProtectionContainerName     pulumi.StringInput `pulumi:"protectionContainerName"`
	ReplicatedProtectedItemName pulumi.StringInput `pulumi:"replicatedProtectedItemName"`
	ResourceGroupName           pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName                pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupReplicationProtectedItemOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationProtectedItemArgs)(nil)).Elem()
}


type LookupReplicationProtectedItemResultOutput struct{ *pulumi.OutputState }

func (LookupReplicationProtectedItemResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationProtectedItemResult)(nil)).Elem()
}

func (o LookupReplicationProtectedItemResultOutput) ToLookupReplicationProtectedItemResultOutput() LookupReplicationProtectedItemResultOutput {
	return o
}

func (o LookupReplicationProtectedItemResultOutput) ToLookupReplicationProtectedItemResultOutputWithContext(ctx context.Context) LookupReplicationProtectedItemResultOutput {
	return o
}

func (o LookupReplicationProtectedItemResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationProtectedItemResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupReplicationProtectedItemResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicationProtectedItemResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupReplicationProtectedItemResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationProtectedItemResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupReplicationProtectedItemResultOutput) Properties() ReplicationProtectedItemPropertiesResponseOutput {
	return o.ApplyT(func(v LookupReplicationProtectedItemResult) ReplicationProtectedItemPropertiesResponse {
		return v.Properties
	}).(ReplicationProtectedItemPropertiesResponseOutput)
}

func (o LookupReplicationProtectedItemResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationProtectedItemResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReplicationProtectedItemResultOutput{})
}
