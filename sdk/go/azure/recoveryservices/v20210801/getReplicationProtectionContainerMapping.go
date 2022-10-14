


package v20210801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupReplicationProtectionContainerMapping(ctx *pulumi.Context, args *LookupReplicationProtectionContainerMappingArgs, opts ...pulumi.InvokeOption) (*LookupReplicationProtectionContainerMappingResult, error) {
	var rv LookupReplicationProtectionContainerMappingResult
	err := ctx.Invoke("azure-native:recoveryservices/v20210801:getReplicationProtectionContainerMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReplicationProtectionContainerMappingArgs struct {
	FabricName              string `pulumi:"fabricName"`
	MappingName             string `pulumi:"mappingName"`
	ProtectionContainerName string `pulumi:"protectionContainerName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
	ResourceName            string `pulumi:"resourceName"`
}


type LookupReplicationProtectionContainerMappingResult struct {
	Id         string                                       `pulumi:"id"`
	Location   *string                                      `pulumi:"location"`
	Name       string                                       `pulumi:"name"`
	Properties ProtectionContainerMappingPropertiesResponse `pulumi:"properties"`
	Type       string                                       `pulumi:"type"`
}

func LookupReplicationProtectionContainerMappingOutput(ctx *pulumi.Context, args LookupReplicationProtectionContainerMappingOutputArgs, opts ...pulumi.InvokeOption) LookupReplicationProtectionContainerMappingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReplicationProtectionContainerMappingResult, error) {
			args := v.(LookupReplicationProtectionContainerMappingArgs)
			r, err := LookupReplicationProtectionContainerMapping(ctx, &args, opts...)
			var s LookupReplicationProtectionContainerMappingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReplicationProtectionContainerMappingResultOutput)
}

type LookupReplicationProtectionContainerMappingOutputArgs struct {
	FabricName              pulumi.StringInput `pulumi:"fabricName"`
	MappingName             pulumi.StringInput `pulumi:"mappingName"`
	ProtectionContainerName pulumi.StringInput `pulumi:"protectionContainerName"`
	ResourceGroupName       pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName            pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupReplicationProtectionContainerMappingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationProtectionContainerMappingArgs)(nil)).Elem()
}


type LookupReplicationProtectionContainerMappingResultOutput struct{ *pulumi.OutputState }

func (LookupReplicationProtectionContainerMappingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationProtectionContainerMappingResult)(nil)).Elem()
}

func (o LookupReplicationProtectionContainerMappingResultOutput) ToLookupReplicationProtectionContainerMappingResultOutput() LookupReplicationProtectionContainerMappingResultOutput {
	return o
}

func (o LookupReplicationProtectionContainerMappingResultOutput) ToLookupReplicationProtectionContainerMappingResultOutputWithContext(ctx context.Context) LookupReplicationProtectionContainerMappingResultOutput {
	return o
}

func (o LookupReplicationProtectionContainerMappingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationProtectionContainerMappingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupReplicationProtectionContainerMappingResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicationProtectionContainerMappingResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupReplicationProtectionContainerMappingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationProtectionContainerMappingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupReplicationProtectionContainerMappingResultOutput) Properties() ProtectionContainerMappingPropertiesResponseOutput {
	return o.ApplyT(func(v LookupReplicationProtectionContainerMappingResult) ProtectionContainerMappingPropertiesResponse {
		return v.Properties
	}).(ProtectionContainerMappingPropertiesResponseOutput)
}

func (o LookupReplicationProtectionContainerMappingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationProtectionContainerMappingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReplicationProtectionContainerMappingResultOutput{})
}
