


package maintenance

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMaintenanceConfiguration(ctx *pulumi.Context, args *LookupMaintenanceConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupMaintenanceConfigurationResult, error) {
	var rv LookupMaintenanceConfigurationResult
	err := ctx.Invoke("azure-native:maintenance:getMaintenanceConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMaintenanceConfigurationArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupMaintenanceConfigurationResult struct {
	ExtensionProperties map[string]string `pulumi:"extensionProperties"`
	Id                  string            `pulumi:"id"`
	Location            *string           `pulumi:"location"`
	MaintenanceScope    *string           `pulumi:"maintenanceScope"`
	Name                string            `pulumi:"name"`
	Namespace           *string           `pulumi:"namespace"`
	Tags                map[string]string `pulumi:"tags"`
	Type                string            `pulumi:"type"`
}

func LookupMaintenanceConfigurationOutput(ctx *pulumi.Context, args LookupMaintenanceConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupMaintenanceConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMaintenanceConfigurationResult, error) {
			args := v.(LookupMaintenanceConfigurationArgs)
			r, err := LookupMaintenanceConfiguration(ctx, &args, opts...)
			var s LookupMaintenanceConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMaintenanceConfigurationResultOutput)
}

type LookupMaintenanceConfigurationOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupMaintenanceConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMaintenanceConfigurationArgs)(nil)).Elem()
}


type LookupMaintenanceConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupMaintenanceConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMaintenanceConfigurationResult)(nil)).Elem()
}

func (o LookupMaintenanceConfigurationResultOutput) ToLookupMaintenanceConfigurationResultOutput() LookupMaintenanceConfigurationResultOutput {
	return o
}

func (o LookupMaintenanceConfigurationResultOutput) ToLookupMaintenanceConfigurationResultOutputWithContext(ctx context.Context) LookupMaintenanceConfigurationResultOutput {
	return o
}

func (o LookupMaintenanceConfigurationResultOutput) ExtensionProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) map[string]string { return v.ExtensionProperties }).(pulumi.StringMapOutput)
}

func (o LookupMaintenanceConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMaintenanceConfigurationResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupMaintenanceConfigurationResultOutput) MaintenanceScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) *string { return v.MaintenanceScope }).(pulumi.StringPtrOutput)
}

func (o LookupMaintenanceConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMaintenanceConfigurationResultOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

func (o LookupMaintenanceConfigurationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupMaintenanceConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMaintenanceConfigurationResultOutput{})
}
