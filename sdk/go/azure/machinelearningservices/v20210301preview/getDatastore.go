


package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDatastore(ctx *pulumi.Context, args *LookupDatastoreArgs, opts ...pulumi.InvokeOption) (*LookupDatastoreResult, error) {
	var rv LookupDatastoreResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20210301preview:getDatastore", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatastoreArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupDatastoreResult struct {
	Id         string                      `pulumi:"id"`
	Name       string                      `pulumi:"name"`
	Properties DatastorePropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse          `pulumi:"systemData"`
	Type       string                      `pulumi:"type"`
}

func LookupDatastoreOutput(ctx *pulumi.Context, args LookupDatastoreOutputArgs, opts ...pulumi.InvokeOption) LookupDatastoreResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatastoreResult, error) {
			args := v.(LookupDatastoreArgs)
			r, err := LookupDatastore(ctx, &args, opts...)
			var s LookupDatastoreResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatastoreResultOutput)
}

type LookupDatastoreOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupDatastoreOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatastoreArgs)(nil)).Elem()
}


type LookupDatastoreResultOutput struct{ *pulumi.OutputState }

func (LookupDatastoreResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatastoreResult)(nil)).Elem()
}

func (o LookupDatastoreResultOutput) ToLookupDatastoreResultOutput() LookupDatastoreResultOutput {
	return o
}

func (o LookupDatastoreResultOutput) ToLookupDatastoreResultOutputWithContext(ctx context.Context) LookupDatastoreResultOutput {
	return o
}

func (o LookupDatastoreResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatastoreResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDatastoreResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatastoreResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDatastoreResultOutput) Properties() DatastorePropertiesResponseOutput {
	return o.ApplyT(func(v LookupDatastoreResult) DatastorePropertiesResponse { return v.Properties }).(DatastorePropertiesResponseOutput)
}

func (o LookupDatastoreResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDatastoreResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupDatastoreResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatastoreResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatastoreResultOutput{})
}
