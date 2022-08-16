


package v20150521preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupCustomImageResource(ctx *pulumi.Context, args *LookupCustomImageResourceArgs, opts ...pulumi.InvokeOption) (*LookupCustomImageResourceResult, error) {
	var rv LookupCustomImageResourceResult
	err := ctx.Invoke("azure-native:devtestlab/v20150521preview:getCustomImageResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCustomImageResourceArgs struct {
	LabName           string `pulumi:"labName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupCustomImageResourceResult struct {
	Author            *string                              `pulumi:"author"`
	CreationDate      *string                              `pulumi:"creationDate"`
	Description       *string                              `pulumi:"description"`
	Id                *string                              `pulumi:"id"`
	Location          *string                              `pulumi:"location"`
	Name              *string                              `pulumi:"name"`
	OsType            *string                              `pulumi:"osType"`
	ProvisioningState *string                              `pulumi:"provisioningState"`
	Tags              map[string]string                    `pulumi:"tags"`
	Type              *string                              `pulumi:"type"`
	Vhd               *CustomImagePropertiesCustomResponse `pulumi:"vhd"`
	Vm                *CustomImagePropertiesFromVmResponse `pulumi:"vm"`
}

func LookupCustomImageResourceOutput(ctx *pulumi.Context, args LookupCustomImageResourceOutputArgs, opts ...pulumi.InvokeOption) LookupCustomImageResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCustomImageResourceResult, error) {
			args := v.(LookupCustomImageResourceArgs)
			r, err := LookupCustomImageResource(ctx, &args, opts...)
			var s LookupCustomImageResourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCustomImageResourceResultOutput)
}

type LookupCustomImageResourceOutputArgs struct {
	LabName           pulumi.StringInput `pulumi:"labName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCustomImageResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomImageResourceArgs)(nil)).Elem()
}


type LookupCustomImageResourceResultOutput struct{ *pulumi.OutputState }

func (LookupCustomImageResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomImageResourceResult)(nil)).Elem()
}

func (o LookupCustomImageResourceResultOutput) ToLookupCustomImageResourceResultOutput() LookupCustomImageResourceResultOutput {
	return o
}

func (o LookupCustomImageResourceResultOutput) ToLookupCustomImageResourceResultOutputWithContext(ctx context.Context) LookupCustomImageResourceResultOutput {
	return o
}

func (o LookupCustomImageResourceResultOutput) Author() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomImageResourceResult) *string { return v.Author }).(pulumi.StringPtrOutput)
}

func (o LookupCustomImageResourceResultOutput) CreationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomImageResourceResult) *string { return v.CreationDate }).(pulumi.StringPtrOutput)
}

func (o LookupCustomImageResourceResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomImageResourceResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupCustomImageResourceResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomImageResourceResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupCustomImageResourceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomImageResourceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupCustomImageResourceResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomImageResourceResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupCustomImageResourceResultOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomImageResourceResult) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o LookupCustomImageResourceResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomImageResourceResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupCustomImageResourceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCustomImageResourceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupCustomImageResourceResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomImageResourceResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o LookupCustomImageResourceResultOutput) Vhd() CustomImagePropertiesCustomResponsePtrOutput {
	return o.ApplyT(func(v LookupCustomImageResourceResult) *CustomImagePropertiesCustomResponse { return v.Vhd }).(CustomImagePropertiesCustomResponsePtrOutput)
}

func (o LookupCustomImageResourceResultOutput) Vm() CustomImagePropertiesFromVmResponsePtrOutput {
	return o.ApplyT(func(v LookupCustomImageResourceResult) *CustomImagePropertiesFromVmResponse { return v.Vm }).(CustomImagePropertiesFromVmResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCustomImageResourceResultOutput{})
}
