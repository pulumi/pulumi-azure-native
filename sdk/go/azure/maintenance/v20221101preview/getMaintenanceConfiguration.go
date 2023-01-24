


package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMaintenanceConfiguration(ctx *pulumi.Context, args *LookupMaintenanceConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupMaintenanceConfigurationResult, error) {
	var rv LookupMaintenanceConfigurationResult
	err := ctx.Invoke("azure-native:maintenance/v20221101preview:getMaintenanceConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupMaintenanceConfigurationArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupMaintenanceConfigurationResult struct {
	Duration            *string                                 `pulumi:"duration"`
	ExpirationDateTime  *string                                 `pulumi:"expirationDateTime"`
	ExtensionProperties map[string]string                       `pulumi:"extensionProperties"`
	Id                  string                                  `pulumi:"id"`
	InstallPatches      *InputPatchConfigurationResponse        `pulumi:"installPatches"`
	Location            *string                                 `pulumi:"location"`
	MaintenanceScope    *string                                 `pulumi:"maintenanceScope"`
	Name                string                                  `pulumi:"name"`
	Namespace           *string                                 `pulumi:"namespace"`
	Overrides           []MaintenanceOverridePropertiesResponse `pulumi:"overrides"`
	RecurEvery          *string                                 `pulumi:"recurEvery"`
	StartDateTime       *string                                 `pulumi:"startDateTime"`
	SystemData          SystemDataResponse                      `pulumi:"systemData"`
	Tags                map[string]string                       `pulumi:"tags"`
	TimeZone            *string                                 `pulumi:"timeZone"`
	Type                string                                  `pulumi:"type"`
	Visibility          *string                                 `pulumi:"visibility"`
}


func (val *LookupMaintenanceConfigurationResult) Defaults() *LookupMaintenanceConfigurationResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.InstallPatches = tmp.InstallPatches.Defaults()

	return &tmp
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

func (o LookupMaintenanceConfigurationResultOutput) Duration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) *string { return v.Duration }).(pulumi.StringPtrOutput)
}

func (o LookupMaintenanceConfigurationResultOutput) ExpirationDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) *string { return v.ExpirationDateTime }).(pulumi.StringPtrOutput)
}

func (o LookupMaintenanceConfigurationResultOutput) ExtensionProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) map[string]string { return v.ExtensionProperties }).(pulumi.StringMapOutput)
}

func (o LookupMaintenanceConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMaintenanceConfigurationResultOutput) InstallPatches() InputPatchConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) *InputPatchConfigurationResponse { return v.InstallPatches }).(InputPatchConfigurationResponsePtrOutput)
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

func (o LookupMaintenanceConfigurationResultOutput) Overrides() MaintenanceOverridePropertiesResponseArrayOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) []MaintenanceOverridePropertiesResponse {
		return v.Overrides
	}).(MaintenanceOverridePropertiesResponseArrayOutput)
}

func (o LookupMaintenanceConfigurationResultOutput) RecurEvery() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) *string { return v.RecurEvery }).(pulumi.StringPtrOutput)
}

func (o LookupMaintenanceConfigurationResultOutput) StartDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) *string { return v.StartDateTime }).(pulumi.StringPtrOutput)
}

func (o LookupMaintenanceConfigurationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupMaintenanceConfigurationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupMaintenanceConfigurationResultOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) *string { return v.TimeZone }).(pulumi.StringPtrOutput)
}

func (o LookupMaintenanceConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupMaintenanceConfigurationResultOutput) Visibility() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMaintenanceConfigurationResult) *string { return v.Visibility }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMaintenanceConfigurationResultOutput{})
}
