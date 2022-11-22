


package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRegistryCodeVersion(ctx *pulumi.Context, args *LookupRegistryCodeVersionArgs, opts ...pulumi.InvokeOption) (*LookupRegistryCodeVersionResult, error) {
	var rv LookupRegistryCodeVersionResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20221001preview:getRegistryCodeVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupRegistryCodeVersionArgs struct {
	CodeName          string `pulumi:"codeName"`
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Version           string `pulumi:"version"`
}


type LookupRegistryCodeVersionResult struct {
	CodeVersionProperties CodeVersionResponse `pulumi:"codeVersionProperties"`
	Id                    string              `pulumi:"id"`
	Name                  string              `pulumi:"name"`
	SystemData            SystemDataResponse  `pulumi:"systemData"`
	Type                  string              `pulumi:"type"`
}


func (val *LookupRegistryCodeVersionResult) Defaults() *LookupRegistryCodeVersionResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.CodeVersionProperties = *tmp.CodeVersionProperties.Defaults()

	return &tmp
}

func LookupRegistryCodeVersionOutput(ctx *pulumi.Context, args LookupRegistryCodeVersionOutputArgs, opts ...pulumi.InvokeOption) LookupRegistryCodeVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegistryCodeVersionResult, error) {
			args := v.(LookupRegistryCodeVersionArgs)
			r, err := LookupRegistryCodeVersion(ctx, &args, opts...)
			var s LookupRegistryCodeVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRegistryCodeVersionResultOutput)
}

type LookupRegistryCodeVersionOutputArgs struct {
	CodeName          pulumi.StringInput `pulumi:"codeName"`
	RegistryName      pulumi.StringInput `pulumi:"registryName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Version           pulumi.StringInput `pulumi:"version"`
}

func (LookupRegistryCodeVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryCodeVersionArgs)(nil)).Elem()
}


type LookupRegistryCodeVersionResultOutput struct{ *pulumi.OutputState }

func (LookupRegistryCodeVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryCodeVersionResult)(nil)).Elem()
}

func (o LookupRegistryCodeVersionResultOutput) ToLookupRegistryCodeVersionResultOutput() LookupRegistryCodeVersionResultOutput {
	return o
}

func (o LookupRegistryCodeVersionResultOutput) ToLookupRegistryCodeVersionResultOutputWithContext(ctx context.Context) LookupRegistryCodeVersionResultOutput {
	return o
}

func (o LookupRegistryCodeVersionResultOutput) CodeVersionProperties() CodeVersionResponseOutput {
	return o.ApplyT(func(v LookupRegistryCodeVersionResult) CodeVersionResponse { return v.CodeVersionProperties }).(CodeVersionResponseOutput)
}

func (o LookupRegistryCodeVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryCodeVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRegistryCodeVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryCodeVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRegistryCodeVersionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupRegistryCodeVersionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupRegistryCodeVersionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryCodeVersionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegistryCodeVersionResultOutput{})
}
