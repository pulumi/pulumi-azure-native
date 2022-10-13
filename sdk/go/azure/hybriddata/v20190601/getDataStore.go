


package v20190601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDataStore(ctx *pulumi.Context, args *LookupDataStoreArgs, opts ...pulumi.InvokeOption) (*LookupDataStoreResult, error) {
	var rv LookupDataStoreResult
	err := ctx.Invoke("azure-native:hybriddata/v20190601:getDataStore", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataStoreArgs struct {
	DataManagerName   string `pulumi:"dataManagerName"`
	DataStoreName     string `pulumi:"dataStoreName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDataStoreResult struct {
	CustomerSecrets    []CustomerSecretResponse `pulumi:"customerSecrets"`
	DataStoreTypeId    string                   `pulumi:"dataStoreTypeId"`
	ExtendedProperties interface{}              `pulumi:"extendedProperties"`
	Id                 string                   `pulumi:"id"`
	Name               string                   `pulumi:"name"`
	RepositoryId       *string                  `pulumi:"repositoryId"`
	State              string                   `pulumi:"state"`
	Type               string                   `pulumi:"type"`
}

func LookupDataStoreOutput(ctx *pulumi.Context, args LookupDataStoreOutputArgs, opts ...pulumi.InvokeOption) LookupDataStoreResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDataStoreResult, error) {
			args := v.(LookupDataStoreArgs)
			r, err := LookupDataStore(ctx, &args, opts...)
			var s LookupDataStoreResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDataStoreResultOutput)
}

type LookupDataStoreOutputArgs struct {
	DataManagerName   pulumi.StringInput `pulumi:"dataManagerName"`
	DataStoreName     pulumi.StringInput `pulumi:"dataStoreName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDataStoreOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataStoreArgs)(nil)).Elem()
}


type LookupDataStoreResultOutput struct{ *pulumi.OutputState }

func (LookupDataStoreResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataStoreResult)(nil)).Elem()
}

func (o LookupDataStoreResultOutput) ToLookupDataStoreResultOutput() LookupDataStoreResultOutput {
	return o
}

func (o LookupDataStoreResultOutput) ToLookupDataStoreResultOutputWithContext(ctx context.Context) LookupDataStoreResultOutput {
	return o
}

func (o LookupDataStoreResultOutput) CustomerSecrets() CustomerSecretResponseArrayOutput {
	return o.ApplyT(func(v LookupDataStoreResult) []CustomerSecretResponse { return v.CustomerSecrets }).(CustomerSecretResponseArrayOutput)
}

func (o LookupDataStoreResultOutput) DataStoreTypeId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataStoreResult) string { return v.DataStoreTypeId }).(pulumi.StringOutput)
}

func (o LookupDataStoreResultOutput) ExtendedProperties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupDataStoreResult) interface{} { return v.ExtendedProperties }).(pulumi.AnyOutput)
}

func (o LookupDataStoreResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataStoreResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDataStoreResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataStoreResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDataStoreResultOutput) RepositoryId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataStoreResult) *string { return v.RepositoryId }).(pulumi.StringPtrOutput)
}

func (o LookupDataStoreResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataStoreResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupDataStoreResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataStoreResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataStoreResultOutput{})
}
