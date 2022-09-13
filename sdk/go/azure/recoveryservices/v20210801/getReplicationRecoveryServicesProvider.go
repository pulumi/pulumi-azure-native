


package v20210801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupReplicationRecoveryServicesProvider(ctx *pulumi.Context, args *LookupReplicationRecoveryServicesProviderArgs, opts ...pulumi.InvokeOption) (*LookupReplicationRecoveryServicesProviderResult, error) {
	var rv LookupReplicationRecoveryServicesProviderResult
	err := ctx.Invoke("azure-native:recoveryservices/v20210801:getReplicationRecoveryServicesProvider", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReplicationRecoveryServicesProviderArgs struct {
	FabricName        string `pulumi:"fabricName"`
	ProviderName      string `pulumi:"providerName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupReplicationRecoveryServicesProviderResult struct {
	Id         string                                     `pulumi:"id"`
	Location   *string                                    `pulumi:"location"`
	Name       string                                     `pulumi:"name"`
	Properties RecoveryServicesProviderPropertiesResponse `pulumi:"properties"`
	Type       string                                     `pulumi:"type"`
}

func LookupReplicationRecoveryServicesProviderOutput(ctx *pulumi.Context, args LookupReplicationRecoveryServicesProviderOutputArgs, opts ...pulumi.InvokeOption) LookupReplicationRecoveryServicesProviderResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReplicationRecoveryServicesProviderResult, error) {
			args := v.(LookupReplicationRecoveryServicesProviderArgs)
			r, err := LookupReplicationRecoveryServicesProvider(ctx, &args, opts...)
			var s LookupReplicationRecoveryServicesProviderResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReplicationRecoveryServicesProviderResultOutput)
}

type LookupReplicationRecoveryServicesProviderOutputArgs struct {
	FabricName        pulumi.StringInput `pulumi:"fabricName"`
	ProviderName      pulumi.StringInput `pulumi:"providerName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupReplicationRecoveryServicesProviderOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationRecoveryServicesProviderArgs)(nil)).Elem()
}


type LookupReplicationRecoveryServicesProviderResultOutput struct{ *pulumi.OutputState }

func (LookupReplicationRecoveryServicesProviderResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationRecoveryServicesProviderResult)(nil)).Elem()
}

func (o LookupReplicationRecoveryServicesProviderResultOutput) ToLookupReplicationRecoveryServicesProviderResultOutput() LookupReplicationRecoveryServicesProviderResultOutput {
	return o
}

func (o LookupReplicationRecoveryServicesProviderResultOutput) ToLookupReplicationRecoveryServicesProviderResultOutputWithContext(ctx context.Context) LookupReplicationRecoveryServicesProviderResultOutput {
	return o
}

func (o LookupReplicationRecoveryServicesProviderResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationRecoveryServicesProviderResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupReplicationRecoveryServicesProviderResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicationRecoveryServicesProviderResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupReplicationRecoveryServicesProviderResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationRecoveryServicesProviderResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupReplicationRecoveryServicesProviderResultOutput) Properties() RecoveryServicesProviderPropertiesResponseOutput {
	return o.ApplyT(func(v LookupReplicationRecoveryServicesProviderResult) RecoveryServicesProviderPropertiesResponse {
		return v.Properties
	}).(RecoveryServicesProviderPropertiesResponseOutput)
}

func (o LookupReplicationRecoveryServicesProviderResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationRecoveryServicesProviderResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReplicationRecoveryServicesProviderResultOutput{})
}
