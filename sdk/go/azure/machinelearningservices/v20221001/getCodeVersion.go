


package v20221001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCodeVersion(ctx *pulumi.Context, args *LookupCodeVersionArgs, opts ...pulumi.InvokeOption) (*LookupCodeVersionResult, error) {
	var rv LookupCodeVersionResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20221001:getCodeVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupCodeVersionArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Version           string `pulumi:"version"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupCodeVersionResult struct {
	CodeVersionProperties CodeVersionResponse `pulumi:"codeVersionProperties"`
	Id                    string              `pulumi:"id"`
	Name                  string              `pulumi:"name"`
	SystemData            SystemDataResponse  `pulumi:"systemData"`
	Type                  string              `pulumi:"type"`
}


func (val *LookupCodeVersionResult) Defaults() *LookupCodeVersionResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.CodeVersionProperties = *tmp.CodeVersionProperties.Defaults()

	return &tmp
}

func LookupCodeVersionOutput(ctx *pulumi.Context, args LookupCodeVersionOutputArgs, opts ...pulumi.InvokeOption) LookupCodeVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCodeVersionResult, error) {
			args := v.(LookupCodeVersionArgs)
			r, err := LookupCodeVersion(ctx, &args, opts...)
			var s LookupCodeVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCodeVersionResultOutput)
}

type LookupCodeVersionOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Version           pulumi.StringInput `pulumi:"version"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupCodeVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCodeVersionArgs)(nil)).Elem()
}


type LookupCodeVersionResultOutput struct{ *pulumi.OutputState }

func (LookupCodeVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCodeVersionResult)(nil)).Elem()
}

func (o LookupCodeVersionResultOutput) ToLookupCodeVersionResultOutput() LookupCodeVersionResultOutput {
	return o
}

func (o LookupCodeVersionResultOutput) ToLookupCodeVersionResultOutputWithContext(ctx context.Context) LookupCodeVersionResultOutput {
	return o
}

func (o LookupCodeVersionResultOutput) CodeVersionProperties() CodeVersionResponseOutput {
	return o.ApplyT(func(v LookupCodeVersionResult) CodeVersionResponse { return v.CodeVersionProperties }).(CodeVersionResponseOutput)
}

func (o LookupCodeVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCodeVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCodeVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCodeVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCodeVersionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCodeVersionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupCodeVersionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCodeVersionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCodeVersionResultOutput{})
}
