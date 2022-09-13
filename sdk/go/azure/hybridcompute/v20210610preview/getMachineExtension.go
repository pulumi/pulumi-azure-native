


package v20210610preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMachineExtension(ctx *pulumi.Context, args *LookupMachineExtensionArgs, opts ...pulumi.InvokeOption) (*LookupMachineExtensionResult, error) {
	var rv LookupMachineExtensionResult
	err := ctx.Invoke("azure-native:hybridcompute/v20210610preview:getMachineExtension", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMachineExtensionArgs struct {
	ExtensionName     string `pulumi:"extensionName"`
	MachineName       string `pulumi:"machineName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupMachineExtensionResult struct {
	Id         string                             `pulumi:"id"`
	Location   string                             `pulumi:"location"`
	Name       string                             `pulumi:"name"`
	Properties MachineExtensionPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse                 `pulumi:"systemData"`
	Tags       map[string]string                  `pulumi:"tags"`
	Type       string                             `pulumi:"type"`
}

func LookupMachineExtensionOutput(ctx *pulumi.Context, args LookupMachineExtensionOutputArgs, opts ...pulumi.InvokeOption) LookupMachineExtensionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMachineExtensionResult, error) {
			args := v.(LookupMachineExtensionArgs)
			r, err := LookupMachineExtension(ctx, &args, opts...)
			var s LookupMachineExtensionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMachineExtensionResultOutput)
}

type LookupMachineExtensionOutputArgs struct {
	ExtensionName     pulumi.StringInput `pulumi:"extensionName"`
	MachineName       pulumi.StringInput `pulumi:"machineName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupMachineExtensionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMachineExtensionArgs)(nil)).Elem()
}


type LookupMachineExtensionResultOutput struct{ *pulumi.OutputState }

func (LookupMachineExtensionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMachineExtensionResult)(nil)).Elem()
}

func (o LookupMachineExtensionResultOutput) ToLookupMachineExtensionResultOutput() LookupMachineExtensionResultOutput {
	return o
}

func (o LookupMachineExtensionResultOutput) ToLookupMachineExtensionResultOutputWithContext(ctx context.Context) LookupMachineExtensionResultOutput {
	return o
}

func (o LookupMachineExtensionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineExtensionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMachineExtensionResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineExtensionResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupMachineExtensionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineExtensionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMachineExtensionResultOutput) Properties() MachineExtensionPropertiesResponseOutput {
	return o.ApplyT(func(v LookupMachineExtensionResult) MachineExtensionPropertiesResponse { return v.Properties }).(MachineExtensionPropertiesResponseOutput)
}

func (o LookupMachineExtensionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMachineExtensionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupMachineExtensionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMachineExtensionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupMachineExtensionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineExtensionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMachineExtensionResultOutput{})
}
