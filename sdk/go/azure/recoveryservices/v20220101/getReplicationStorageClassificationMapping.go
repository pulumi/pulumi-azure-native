


package v20220101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupReplicationStorageClassificationMapping(ctx *pulumi.Context, args *LookupReplicationStorageClassificationMappingArgs, opts ...pulumi.InvokeOption) (*LookupReplicationStorageClassificationMappingResult, error) {
	var rv LookupReplicationStorageClassificationMappingResult
	err := ctx.Invoke("azure-native:recoveryservices/v20220101:getReplicationStorageClassificationMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReplicationStorageClassificationMappingArgs struct {
	FabricName                       string `pulumi:"fabricName"`
	ResourceGroupName                string `pulumi:"resourceGroupName"`
	ResourceName                     string `pulumi:"resourceName"`
	StorageClassificationMappingName string `pulumi:"storageClassificationMappingName"`
	StorageClassificationName        string `pulumi:"storageClassificationName"`
}


type LookupReplicationStorageClassificationMappingResult struct {
	Id         string                                         `pulumi:"id"`
	Location   *string                                        `pulumi:"location"`
	Name       string                                         `pulumi:"name"`
	Properties StorageClassificationMappingPropertiesResponse `pulumi:"properties"`
	Type       string                                         `pulumi:"type"`
}

func LookupReplicationStorageClassificationMappingOutput(ctx *pulumi.Context, args LookupReplicationStorageClassificationMappingOutputArgs, opts ...pulumi.InvokeOption) LookupReplicationStorageClassificationMappingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReplicationStorageClassificationMappingResult, error) {
			args := v.(LookupReplicationStorageClassificationMappingArgs)
			r, err := LookupReplicationStorageClassificationMapping(ctx, &args, opts...)
			var s LookupReplicationStorageClassificationMappingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReplicationStorageClassificationMappingResultOutput)
}

type LookupReplicationStorageClassificationMappingOutputArgs struct {
	FabricName                       pulumi.StringInput `pulumi:"fabricName"`
	ResourceGroupName                pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName                     pulumi.StringInput `pulumi:"resourceName"`
	StorageClassificationMappingName pulumi.StringInput `pulumi:"storageClassificationMappingName"`
	StorageClassificationName        pulumi.StringInput `pulumi:"storageClassificationName"`
}

func (LookupReplicationStorageClassificationMappingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationStorageClassificationMappingArgs)(nil)).Elem()
}


type LookupReplicationStorageClassificationMappingResultOutput struct{ *pulumi.OutputState }

func (LookupReplicationStorageClassificationMappingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationStorageClassificationMappingResult)(nil)).Elem()
}

func (o LookupReplicationStorageClassificationMappingResultOutput) ToLookupReplicationStorageClassificationMappingResultOutput() LookupReplicationStorageClassificationMappingResultOutput {
	return o
}

func (o LookupReplicationStorageClassificationMappingResultOutput) ToLookupReplicationStorageClassificationMappingResultOutputWithContext(ctx context.Context) LookupReplicationStorageClassificationMappingResultOutput {
	return o
}

func (o LookupReplicationStorageClassificationMappingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationStorageClassificationMappingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupReplicationStorageClassificationMappingResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicationStorageClassificationMappingResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupReplicationStorageClassificationMappingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationStorageClassificationMappingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupReplicationStorageClassificationMappingResultOutput) Properties() StorageClassificationMappingPropertiesResponseOutput {
	return o.ApplyT(func(v LookupReplicationStorageClassificationMappingResult) StorageClassificationMappingPropertiesResponse {
		return v.Properties
	}).(StorageClassificationMappingPropertiesResponseOutput)
}

func (o LookupReplicationStorageClassificationMappingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationStorageClassificationMappingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReplicationStorageClassificationMappingResultOutput{})
}
