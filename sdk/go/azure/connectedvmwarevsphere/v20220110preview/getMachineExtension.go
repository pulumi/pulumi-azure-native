


package v20220110preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMachineExtension(ctx *pulumi.Context, args *LookupMachineExtensionArgs, opts ...pulumi.InvokeOption) (*LookupMachineExtensionResult, error) {
	var rv LookupMachineExtensionResult
	err := ctx.Invoke("azure-native:connectedvmwarevsphere/v20220110preview:getMachineExtension", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMachineExtensionArgs struct {
	ExtensionName     string `pulumi:"extensionName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupMachineExtensionResult struct {
	AutoUpgradeMinorVersion *bool                                           `pulumi:"autoUpgradeMinorVersion"`
	EnableAutomaticUpgrade  *bool                                           `pulumi:"enableAutomaticUpgrade"`
	ForceUpdateTag          *string                                         `pulumi:"forceUpdateTag"`
	Id                      string                                          `pulumi:"id"`
	InstanceView            *MachineExtensionPropertiesResponseInstanceView `pulumi:"instanceView"`
	Location                *string                                         `pulumi:"location"`
	Name                    string                                          `pulumi:"name"`
	ProtectedSettings       interface{}                                     `pulumi:"protectedSettings"`
	ProvisioningState       string                                          `pulumi:"provisioningState"`
	Publisher               *string                                         `pulumi:"publisher"`
	Settings                interface{}                                     `pulumi:"settings"`
	SystemData              SystemDataResponse                              `pulumi:"systemData"`
	Tags                    map[string]string                               `pulumi:"tags"`
	Type                    string                                          `pulumi:"type"`
	TypeHandlerVersion      *string                                         `pulumi:"typeHandlerVersion"`
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
	Name              pulumi.StringInput `pulumi:"name"`
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

func (o LookupMachineExtensionResultOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupMachineExtensionResult) *bool { return v.AutoUpgradeMinorVersion }).(pulumi.BoolPtrOutput)
}

func (o LookupMachineExtensionResultOutput) EnableAutomaticUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupMachineExtensionResult) *bool { return v.EnableAutomaticUpgrade }).(pulumi.BoolPtrOutput)
}

func (o LookupMachineExtensionResultOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMachineExtensionResult) *string { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

func (o LookupMachineExtensionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineExtensionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMachineExtensionResultOutput) InstanceView() MachineExtensionPropertiesResponseInstanceViewPtrOutput {
	return o.ApplyT(func(v LookupMachineExtensionResult) *MachineExtensionPropertiesResponseInstanceView {
		return v.InstanceView
	}).(MachineExtensionPropertiesResponseInstanceViewPtrOutput)
}

func (o LookupMachineExtensionResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMachineExtensionResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupMachineExtensionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineExtensionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMachineExtensionResultOutput) ProtectedSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupMachineExtensionResult) interface{} { return v.ProtectedSettings }).(pulumi.AnyOutput)
}

func (o LookupMachineExtensionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineExtensionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupMachineExtensionResultOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMachineExtensionResult) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o LookupMachineExtensionResultOutput) Settings() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupMachineExtensionResult) interface{} { return v.Settings }).(pulumi.AnyOutput)
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

func (o LookupMachineExtensionResultOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMachineExtensionResult) *string { return v.TypeHandlerVersion }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMachineExtensionResultOutput{})
}
