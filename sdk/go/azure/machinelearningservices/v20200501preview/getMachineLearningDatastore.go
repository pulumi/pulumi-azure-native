


package v20200501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMachineLearningDatastore(ctx *pulumi.Context, args *LookupMachineLearningDatastoreArgs, opts ...pulumi.InvokeOption) (*LookupMachineLearningDatastoreResult, error) {
	var rv LookupMachineLearningDatastoreResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20200501preview:getMachineLearningDatastore", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupMachineLearningDatastoreArgs struct {
	DatastoreName     string `pulumi:"datastoreName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupMachineLearningDatastoreResult struct {
	Id         string            `pulumi:"id"`
	Identity   *IdentityResponse `pulumi:"identity"`
	Location   *string           `pulumi:"location"`
	Name       string            `pulumi:"name"`
	Properties DatastoreResponse `pulumi:"properties"`
	Sku        *SkuResponse      `pulumi:"sku"`
	Tags       map[string]string `pulumi:"tags"`
	Type       string            `pulumi:"type"`
}


func (val *LookupMachineLearningDatastoreResult) Defaults() *LookupMachineLearningDatastoreResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}

func LookupMachineLearningDatastoreOutput(ctx *pulumi.Context, args LookupMachineLearningDatastoreOutputArgs, opts ...pulumi.InvokeOption) LookupMachineLearningDatastoreResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMachineLearningDatastoreResult, error) {
			args := v.(LookupMachineLearningDatastoreArgs)
			r, err := LookupMachineLearningDatastore(ctx, &args, opts...)
			var s LookupMachineLearningDatastoreResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMachineLearningDatastoreResultOutput)
}

type LookupMachineLearningDatastoreOutputArgs struct {
	DatastoreName     pulumi.StringInput `pulumi:"datastoreName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupMachineLearningDatastoreOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMachineLearningDatastoreArgs)(nil)).Elem()
}


type LookupMachineLearningDatastoreResultOutput struct{ *pulumi.OutputState }

func (LookupMachineLearningDatastoreResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMachineLearningDatastoreResult)(nil)).Elem()
}

func (o LookupMachineLearningDatastoreResultOutput) ToLookupMachineLearningDatastoreResultOutput() LookupMachineLearningDatastoreResultOutput {
	return o
}

func (o LookupMachineLearningDatastoreResultOutput) ToLookupMachineLearningDatastoreResultOutputWithContext(ctx context.Context) LookupMachineLearningDatastoreResultOutput {
	return o
}

func (o LookupMachineLearningDatastoreResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineLearningDatastoreResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMachineLearningDatastoreResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupMachineLearningDatastoreResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o LookupMachineLearningDatastoreResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMachineLearningDatastoreResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupMachineLearningDatastoreResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineLearningDatastoreResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMachineLearningDatastoreResultOutput) Properties() DatastoreResponseOutput {
	return o.ApplyT(func(v LookupMachineLearningDatastoreResult) DatastoreResponse { return v.Properties }).(DatastoreResponseOutput)
}

func (o LookupMachineLearningDatastoreResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupMachineLearningDatastoreResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupMachineLearningDatastoreResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMachineLearningDatastoreResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupMachineLearningDatastoreResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineLearningDatastoreResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMachineLearningDatastoreResultOutput{})
}
