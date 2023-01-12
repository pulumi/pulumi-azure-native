


package v20220601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDataVersion(ctx *pulumi.Context, args *LookupDataVersionArgs, opts ...pulumi.InvokeOption) (*LookupDataVersionResult, error) {
	var rv LookupDataVersionResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20220601preview:getDataVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataVersionArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Version           string `pulumi:"version"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupDataVersionResult struct {
	DataVersionBaseProperties interface{}        `pulumi:"dataVersionBaseProperties"`
	Id                        string             `pulumi:"id"`
	Name                      string             `pulumi:"name"`
	SystemData                SystemDataResponse `pulumi:"systemData"`
	Type                      string             `pulumi:"type"`
}

func LookupDataVersionOutput(ctx *pulumi.Context, args LookupDataVersionOutputArgs, opts ...pulumi.InvokeOption) LookupDataVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDataVersionResult, error) {
			args := v.(LookupDataVersionArgs)
			r, err := LookupDataVersion(ctx, &args, opts...)
			var s LookupDataVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDataVersionResultOutput)
}

type LookupDataVersionOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Version           pulumi.StringInput `pulumi:"version"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupDataVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataVersionArgs)(nil)).Elem()
}


type LookupDataVersionResultOutput struct{ *pulumi.OutputState }

func (LookupDataVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataVersionResult)(nil)).Elem()
}

func (o LookupDataVersionResultOutput) ToLookupDataVersionResultOutput() LookupDataVersionResultOutput {
	return o
}

func (o LookupDataVersionResultOutput) ToLookupDataVersionResultOutputWithContext(ctx context.Context) LookupDataVersionResultOutput {
	return o
}

func (o LookupDataVersionResultOutput) DataVersionBaseProperties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupDataVersionResult) interface{} { return v.DataVersionBaseProperties }).(pulumi.AnyOutput)
}

func (o LookupDataVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDataVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDataVersionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDataVersionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupDataVersionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataVersionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataVersionResultOutput{})
}
