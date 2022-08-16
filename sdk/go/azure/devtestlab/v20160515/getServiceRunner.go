


package v20160515

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupServiceRunner(ctx *pulumi.Context, args *LookupServiceRunnerArgs, opts ...pulumi.InvokeOption) (*LookupServiceRunnerResult, error) {
	var rv LookupServiceRunnerResult
	err := ctx.Invoke("azure-native:devtestlab/v20160515:getServiceRunner", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceRunnerArgs struct {
	LabName           string `pulumi:"labName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupServiceRunnerResult struct {
	Id       string                      `pulumi:"id"`
	Identity *IdentityPropertiesResponse `pulumi:"identity"`
	Location *string                     `pulumi:"location"`
	Name     string                      `pulumi:"name"`
	Tags     map[string]string           `pulumi:"tags"`
	Type     string                      `pulumi:"type"`
}

func LookupServiceRunnerOutput(ctx *pulumi.Context, args LookupServiceRunnerOutputArgs, opts ...pulumi.InvokeOption) LookupServiceRunnerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServiceRunnerResult, error) {
			args := v.(LookupServiceRunnerArgs)
			r, err := LookupServiceRunner(ctx, &args, opts...)
			var s LookupServiceRunnerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServiceRunnerResultOutput)
}

type LookupServiceRunnerOutputArgs struct {
	LabName           pulumi.StringInput `pulumi:"labName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupServiceRunnerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceRunnerArgs)(nil)).Elem()
}


type LookupServiceRunnerResultOutput struct{ *pulumi.OutputState }

func (LookupServiceRunnerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceRunnerResult)(nil)).Elem()
}

func (o LookupServiceRunnerResultOutput) ToLookupServiceRunnerResultOutput() LookupServiceRunnerResultOutput {
	return o
}

func (o LookupServiceRunnerResultOutput) ToLookupServiceRunnerResultOutputWithContext(ctx context.Context) LookupServiceRunnerResultOutput {
	return o
}

func (o LookupServiceRunnerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceRunnerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupServiceRunnerResultOutput) Identity() IdentityPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupServiceRunnerResult) *IdentityPropertiesResponse { return v.Identity }).(IdentityPropertiesResponsePtrOutput)
}

func (o LookupServiceRunnerResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceRunnerResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupServiceRunnerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceRunnerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupServiceRunnerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupServiceRunnerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupServiceRunnerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceRunnerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceRunnerResultOutput{})
}
