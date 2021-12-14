


package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AppResourceProperties struct {
	ActiveDeploymentName *string         `pulumi:"activeDeploymentName"`
	EnableEndToEndTLS    *bool           `pulumi:"enableEndToEndTLS"`
	Fqdn                 *string         `pulumi:"fqdn"`
	HttpsOnly            *bool           `pulumi:"httpsOnly"`
	PersistentDisk       *PersistentDisk `pulumi:"persistentDisk"`
	Public               *bool           `pulumi:"public"`
	TemporaryDisk        *TemporaryDisk  `pulumi:"temporaryDisk"`
}


func (val *AppResourceProperties) Defaults() *AppResourceProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EnableEndToEndTLS) {
		enableEndToEndTLS_ := false
		tmp.EnableEndToEndTLS = &enableEndToEndTLS_
	}
	if isZero(tmp.HttpsOnly) {
		httpsOnly_ := false
		tmp.HttpsOnly = &httpsOnly_
	}
	tmp.TemporaryDisk = tmp.TemporaryDisk.Defaults()

	return &tmp
}





type AppResourcePropertiesInput interface {
	pulumi.Input

	ToAppResourcePropertiesOutput() AppResourcePropertiesOutput
	ToAppResourcePropertiesOutputWithContext(context.Context) AppResourcePropertiesOutput
}

type AppResourcePropertiesArgs struct {
	ActiveDeploymentName pulumi.StringPtrInput  `pulumi:"activeDeploymentName"`
	EnableEndToEndTLS    pulumi.BoolPtrInput    `pulumi:"enableEndToEndTLS"`
	Fqdn                 pulumi.StringPtrInput  `pulumi:"fqdn"`
	HttpsOnly            pulumi.BoolPtrInput    `pulumi:"httpsOnly"`
	PersistentDisk       PersistentDiskPtrInput `pulumi:"persistentDisk"`
	Public               pulumi.BoolPtrInput    `pulumi:"public"`
	TemporaryDisk        TemporaryDiskPtrInput  `pulumi:"temporaryDisk"`
}

func (AppResourcePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AppResourceProperties)(nil)).Elem()
}

func (i AppResourcePropertiesArgs) ToAppResourcePropertiesOutput() AppResourcePropertiesOutput {
	return i.ToAppResourcePropertiesOutputWithContext(context.Background())
}

func (i AppResourcePropertiesArgs) ToAppResourcePropertiesOutputWithContext(ctx context.Context) AppResourcePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppResourcePropertiesOutput)
}

func (i AppResourcePropertiesArgs) ToAppResourcePropertiesPtrOutput() AppResourcePropertiesPtrOutput {
	return i.ToAppResourcePropertiesPtrOutputWithContext(context.Background())
}

func (i AppResourcePropertiesArgs) ToAppResourcePropertiesPtrOutputWithContext(ctx context.Context) AppResourcePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppResourcePropertiesOutput).ToAppResourcePropertiesPtrOutputWithContext(ctx)
}









type AppResourcePropertiesPtrInput interface {
	pulumi.Input

	ToAppResourcePropertiesPtrOutput() AppResourcePropertiesPtrOutput
	ToAppResourcePropertiesPtrOutputWithContext(context.Context) AppResourcePropertiesPtrOutput
}

type appResourcePropertiesPtrType AppResourcePropertiesArgs

func AppResourcePropertiesPtr(v *AppResourcePropertiesArgs) AppResourcePropertiesPtrInput {
	return (*appResourcePropertiesPtrType)(v)
}

func (*appResourcePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AppResourceProperties)(nil)).Elem()
}

func (i *appResourcePropertiesPtrType) ToAppResourcePropertiesPtrOutput() AppResourcePropertiesPtrOutput {
	return i.ToAppResourcePropertiesPtrOutputWithContext(context.Background())
}

func (i *appResourcePropertiesPtrType) ToAppResourcePropertiesPtrOutputWithContext(ctx context.Context) AppResourcePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppResourcePropertiesPtrOutput)
}

type AppResourcePropertiesOutput struct{ *pulumi.OutputState }

func (AppResourcePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppResourceProperties)(nil)).Elem()
}

func (o AppResourcePropertiesOutput) ToAppResourcePropertiesOutput() AppResourcePropertiesOutput {
	return o
}

func (o AppResourcePropertiesOutput) ToAppResourcePropertiesOutputWithContext(ctx context.Context) AppResourcePropertiesOutput {
	return o
}

func (o AppResourcePropertiesOutput) ToAppResourcePropertiesPtrOutput() AppResourcePropertiesPtrOutput {
	return o.ToAppResourcePropertiesPtrOutputWithContext(context.Background())
}

func (o AppResourcePropertiesOutput) ToAppResourcePropertiesPtrOutputWithContext(ctx context.Context) AppResourcePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AppResourceProperties) *AppResourceProperties {
		return &v
	}).(AppResourcePropertiesPtrOutput)
}

func (o AppResourcePropertiesOutput) ActiveDeploymentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppResourceProperties) *string { return v.ActiveDeploymentName }).(pulumi.StringPtrOutput)
}

func (o AppResourcePropertiesOutput) EnableEndToEndTLS() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AppResourceProperties) *bool { return v.EnableEndToEndTLS }).(pulumi.BoolPtrOutput)
}

func (o AppResourcePropertiesOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppResourceProperties) *string { return v.Fqdn }).(pulumi.StringPtrOutput)
}

func (o AppResourcePropertiesOutput) HttpsOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AppResourceProperties) *bool { return v.HttpsOnly }).(pulumi.BoolPtrOutput)
}

func (o AppResourcePropertiesOutput) PersistentDisk() PersistentDiskPtrOutput {
	return o.ApplyT(func(v AppResourceProperties) *PersistentDisk { return v.PersistentDisk }).(PersistentDiskPtrOutput)
}

func (o AppResourcePropertiesOutput) Public() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AppResourceProperties) *bool { return v.Public }).(pulumi.BoolPtrOutput)
}

func (o AppResourcePropertiesOutput) TemporaryDisk() TemporaryDiskPtrOutput {
	return o.ApplyT(func(v AppResourceProperties) *TemporaryDisk { return v.TemporaryDisk }).(TemporaryDiskPtrOutput)
}

type AppResourcePropertiesPtrOutput struct{ *pulumi.OutputState }

func (AppResourcePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppResourceProperties)(nil)).Elem()
}

func (o AppResourcePropertiesPtrOutput) ToAppResourcePropertiesPtrOutput() AppResourcePropertiesPtrOutput {
	return o
}

func (o AppResourcePropertiesPtrOutput) ToAppResourcePropertiesPtrOutputWithContext(ctx context.Context) AppResourcePropertiesPtrOutput {
	return o
}

func (o AppResourcePropertiesPtrOutput) Elem() AppResourcePropertiesOutput {
	return o.ApplyT(func(v *AppResourceProperties) AppResourceProperties {
		if v != nil {
			return *v
		}
		var ret AppResourceProperties
		return ret
	}).(AppResourcePropertiesOutput)
}

func (o AppResourcePropertiesPtrOutput) ActiveDeploymentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppResourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.ActiveDeploymentName
	}).(pulumi.StringPtrOutput)
}

func (o AppResourcePropertiesPtrOutput) EnableEndToEndTLS() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AppResourceProperties) *bool {
		if v == nil {
			return nil
		}
		return v.EnableEndToEndTLS
	}).(pulumi.BoolPtrOutput)
}

func (o AppResourcePropertiesPtrOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppResourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.Fqdn
	}).(pulumi.StringPtrOutput)
}

func (o AppResourcePropertiesPtrOutput) HttpsOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AppResourceProperties) *bool {
		if v == nil {
			return nil
		}
		return v.HttpsOnly
	}).(pulumi.BoolPtrOutput)
}

func (o AppResourcePropertiesPtrOutput) PersistentDisk() PersistentDiskPtrOutput {
	return o.ApplyT(func(v *AppResourceProperties) *PersistentDisk {
		if v == nil {
			return nil
		}
		return v.PersistentDisk
	}).(PersistentDiskPtrOutput)
}

func (o AppResourcePropertiesPtrOutput) Public() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AppResourceProperties) *bool {
		if v == nil {
			return nil
		}
		return v.Public
	}).(pulumi.BoolPtrOutput)
}

func (o AppResourcePropertiesPtrOutput) TemporaryDisk() TemporaryDiskPtrOutput {
	return o.ApplyT(func(v *AppResourceProperties) *TemporaryDisk {
		if v == nil {
			return nil
		}
		return v.TemporaryDisk
	}).(TemporaryDiskPtrOutput)
}

type AppResourcePropertiesResponse struct {
	ActiveDeploymentName *string                 `pulumi:"activeDeploymentName"`
	CreatedTime          string                  `pulumi:"createdTime"`
	EnableEndToEndTLS    *bool                   `pulumi:"enableEndToEndTLS"`
	Fqdn                 *string                 `pulumi:"fqdn"`
	HttpsOnly            *bool                   `pulumi:"httpsOnly"`
	PersistentDisk       *PersistentDiskResponse `pulumi:"persistentDisk"`
	ProvisioningState    string                  `pulumi:"provisioningState"`
	Public               *bool                   `pulumi:"public"`
	TemporaryDisk        *TemporaryDiskResponse  `pulumi:"temporaryDisk"`
	Url                  string                  `pulumi:"url"`
}


func (val *AppResourcePropertiesResponse) Defaults() *AppResourcePropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EnableEndToEndTLS) {
		enableEndToEndTLS_ := false
		tmp.EnableEndToEndTLS = &enableEndToEndTLS_
	}
	if isZero(tmp.HttpsOnly) {
		httpsOnly_ := false
		tmp.HttpsOnly = &httpsOnly_
	}
	tmp.TemporaryDisk = tmp.TemporaryDisk.Defaults()

	return &tmp
}





type AppResourcePropertiesResponseInput interface {
	pulumi.Input

	ToAppResourcePropertiesResponseOutput() AppResourcePropertiesResponseOutput
	ToAppResourcePropertiesResponseOutputWithContext(context.Context) AppResourcePropertiesResponseOutput
}

type AppResourcePropertiesResponseArgs struct {
	ActiveDeploymentName pulumi.StringPtrInput          `pulumi:"activeDeploymentName"`
	CreatedTime          pulumi.StringInput             `pulumi:"createdTime"`
	EnableEndToEndTLS    pulumi.BoolPtrInput            `pulumi:"enableEndToEndTLS"`
	Fqdn                 pulumi.StringPtrInput          `pulumi:"fqdn"`
	HttpsOnly            pulumi.BoolPtrInput            `pulumi:"httpsOnly"`
	PersistentDisk       PersistentDiskResponsePtrInput `pulumi:"persistentDisk"`
	ProvisioningState    pulumi.StringInput             `pulumi:"provisioningState"`
	Public               pulumi.BoolPtrInput            `pulumi:"public"`
	TemporaryDisk        TemporaryDiskResponsePtrInput  `pulumi:"temporaryDisk"`
	Url                  pulumi.StringInput             `pulumi:"url"`
}

func (AppResourcePropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AppResourcePropertiesResponse)(nil)).Elem()
}

func (i AppResourcePropertiesResponseArgs) ToAppResourcePropertiesResponseOutput() AppResourcePropertiesResponseOutput {
	return i.ToAppResourcePropertiesResponseOutputWithContext(context.Background())
}

func (i AppResourcePropertiesResponseArgs) ToAppResourcePropertiesResponseOutputWithContext(ctx context.Context) AppResourcePropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppResourcePropertiesResponseOutput)
}

func (i AppResourcePropertiesResponseArgs) ToAppResourcePropertiesResponsePtrOutput() AppResourcePropertiesResponsePtrOutput {
	return i.ToAppResourcePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i AppResourcePropertiesResponseArgs) ToAppResourcePropertiesResponsePtrOutputWithContext(ctx context.Context) AppResourcePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppResourcePropertiesResponseOutput).ToAppResourcePropertiesResponsePtrOutputWithContext(ctx)
}









type AppResourcePropertiesResponsePtrInput interface {
	pulumi.Input

	ToAppResourcePropertiesResponsePtrOutput() AppResourcePropertiesResponsePtrOutput
	ToAppResourcePropertiesResponsePtrOutputWithContext(context.Context) AppResourcePropertiesResponsePtrOutput
}

type appResourcePropertiesResponsePtrType AppResourcePropertiesResponseArgs

func AppResourcePropertiesResponsePtr(v *AppResourcePropertiesResponseArgs) AppResourcePropertiesResponsePtrInput {
	return (*appResourcePropertiesResponsePtrType)(v)
}

func (*appResourcePropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AppResourcePropertiesResponse)(nil)).Elem()
}

func (i *appResourcePropertiesResponsePtrType) ToAppResourcePropertiesResponsePtrOutput() AppResourcePropertiesResponsePtrOutput {
	return i.ToAppResourcePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *appResourcePropertiesResponsePtrType) ToAppResourcePropertiesResponsePtrOutputWithContext(ctx context.Context) AppResourcePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppResourcePropertiesResponsePtrOutput)
}

type AppResourcePropertiesResponseOutput struct{ *pulumi.OutputState }

func (AppResourcePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppResourcePropertiesResponse)(nil)).Elem()
}

func (o AppResourcePropertiesResponseOutput) ToAppResourcePropertiesResponseOutput() AppResourcePropertiesResponseOutput {
	return o
}

func (o AppResourcePropertiesResponseOutput) ToAppResourcePropertiesResponseOutputWithContext(ctx context.Context) AppResourcePropertiesResponseOutput {
	return o
}

func (o AppResourcePropertiesResponseOutput) ToAppResourcePropertiesResponsePtrOutput() AppResourcePropertiesResponsePtrOutput {
	return o.ToAppResourcePropertiesResponsePtrOutputWithContext(context.Background())
}

func (o AppResourcePropertiesResponseOutput) ToAppResourcePropertiesResponsePtrOutputWithContext(ctx context.Context) AppResourcePropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AppResourcePropertiesResponse) *AppResourcePropertiesResponse {
		return &v
	}).(AppResourcePropertiesResponsePtrOutput)
}

func (o AppResourcePropertiesResponseOutput) ActiveDeploymentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppResourcePropertiesResponse) *string { return v.ActiveDeploymentName }).(pulumi.StringPtrOutput)
}

func (o AppResourcePropertiesResponseOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v AppResourcePropertiesResponse) string { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o AppResourcePropertiesResponseOutput) EnableEndToEndTLS() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AppResourcePropertiesResponse) *bool { return v.EnableEndToEndTLS }).(pulumi.BoolPtrOutput)
}

func (o AppResourcePropertiesResponseOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppResourcePropertiesResponse) *string { return v.Fqdn }).(pulumi.StringPtrOutput)
}

func (o AppResourcePropertiesResponseOutput) HttpsOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AppResourcePropertiesResponse) *bool { return v.HttpsOnly }).(pulumi.BoolPtrOutput)
}

func (o AppResourcePropertiesResponseOutput) PersistentDisk() PersistentDiskResponsePtrOutput {
	return o.ApplyT(func(v AppResourcePropertiesResponse) *PersistentDiskResponse { return v.PersistentDisk }).(PersistentDiskResponsePtrOutput)
}

func (o AppResourcePropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v AppResourcePropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AppResourcePropertiesResponseOutput) Public() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AppResourcePropertiesResponse) *bool { return v.Public }).(pulumi.BoolPtrOutput)
}

func (o AppResourcePropertiesResponseOutput) TemporaryDisk() TemporaryDiskResponsePtrOutput {
	return o.ApplyT(func(v AppResourcePropertiesResponse) *TemporaryDiskResponse { return v.TemporaryDisk }).(TemporaryDiskResponsePtrOutput)
}

func (o AppResourcePropertiesResponseOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v AppResourcePropertiesResponse) string { return v.Url }).(pulumi.StringOutput)
}

type AppResourcePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (AppResourcePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppResourcePropertiesResponse)(nil)).Elem()
}

func (o AppResourcePropertiesResponsePtrOutput) ToAppResourcePropertiesResponsePtrOutput() AppResourcePropertiesResponsePtrOutput {
	return o
}

func (o AppResourcePropertiesResponsePtrOutput) ToAppResourcePropertiesResponsePtrOutputWithContext(ctx context.Context) AppResourcePropertiesResponsePtrOutput {
	return o
}

func (o AppResourcePropertiesResponsePtrOutput) Elem() AppResourcePropertiesResponseOutput {
	return o.ApplyT(func(v *AppResourcePropertiesResponse) AppResourcePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret AppResourcePropertiesResponse
		return ret
	}).(AppResourcePropertiesResponseOutput)
}

func (o AppResourcePropertiesResponsePtrOutput) ActiveDeploymentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppResourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ActiveDeploymentName
	}).(pulumi.StringPtrOutput)
}

func (o AppResourcePropertiesResponsePtrOutput) CreatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppResourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CreatedTime
	}).(pulumi.StringPtrOutput)
}

func (o AppResourcePropertiesResponsePtrOutput) EnableEndToEndTLS() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AppResourcePropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableEndToEndTLS
	}).(pulumi.BoolPtrOutput)
}

func (o AppResourcePropertiesResponsePtrOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppResourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Fqdn
	}).(pulumi.StringPtrOutput)
}

func (o AppResourcePropertiesResponsePtrOutput) HttpsOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AppResourcePropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.HttpsOnly
	}).(pulumi.BoolPtrOutput)
}

func (o AppResourcePropertiesResponsePtrOutput) PersistentDisk() PersistentDiskResponsePtrOutput {
	return o.ApplyT(func(v *AppResourcePropertiesResponse) *PersistentDiskResponse {
		if v == nil {
			return nil
		}
		return v.PersistentDisk
	}).(PersistentDiskResponsePtrOutput)
}

func (o AppResourcePropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppResourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o AppResourcePropertiesResponsePtrOutput) Public() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AppResourcePropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Public
	}).(pulumi.BoolPtrOutput)
}

func (o AppResourcePropertiesResponsePtrOutput) TemporaryDisk() TemporaryDiskResponsePtrOutput {
	return o.ApplyT(func(v *AppResourcePropertiesResponse) *TemporaryDiskResponse {
		if v == nil {
			return nil
		}
		return v.TemporaryDisk
	}).(TemporaryDiskResponsePtrOutput)
}

func (o AppResourcePropertiesResponsePtrOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppResourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Url
	}).(pulumi.StringPtrOutput)
}

type BindingResourceProperties struct {
	BindingParameters map[string]interface{} `pulumi:"bindingParameters"`
	Key               *string                `pulumi:"key"`
	ResourceId        *string                `pulumi:"resourceId"`
}





type BindingResourcePropertiesInput interface {
	pulumi.Input

	ToBindingResourcePropertiesOutput() BindingResourcePropertiesOutput
	ToBindingResourcePropertiesOutputWithContext(context.Context) BindingResourcePropertiesOutput
}

type BindingResourcePropertiesArgs struct {
	BindingParameters pulumi.MapInput       `pulumi:"bindingParameters"`
	Key               pulumi.StringPtrInput `pulumi:"key"`
	ResourceId        pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (BindingResourcePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BindingResourceProperties)(nil)).Elem()
}

func (i BindingResourcePropertiesArgs) ToBindingResourcePropertiesOutput() BindingResourcePropertiesOutput {
	return i.ToBindingResourcePropertiesOutputWithContext(context.Background())
}

func (i BindingResourcePropertiesArgs) ToBindingResourcePropertiesOutputWithContext(ctx context.Context) BindingResourcePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BindingResourcePropertiesOutput)
}

func (i BindingResourcePropertiesArgs) ToBindingResourcePropertiesPtrOutput() BindingResourcePropertiesPtrOutput {
	return i.ToBindingResourcePropertiesPtrOutputWithContext(context.Background())
}

func (i BindingResourcePropertiesArgs) ToBindingResourcePropertiesPtrOutputWithContext(ctx context.Context) BindingResourcePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BindingResourcePropertiesOutput).ToBindingResourcePropertiesPtrOutputWithContext(ctx)
}









type BindingResourcePropertiesPtrInput interface {
	pulumi.Input

	ToBindingResourcePropertiesPtrOutput() BindingResourcePropertiesPtrOutput
	ToBindingResourcePropertiesPtrOutputWithContext(context.Context) BindingResourcePropertiesPtrOutput
}

type bindingResourcePropertiesPtrType BindingResourcePropertiesArgs

func BindingResourcePropertiesPtr(v *BindingResourcePropertiesArgs) BindingResourcePropertiesPtrInput {
	return (*bindingResourcePropertiesPtrType)(v)
}

func (*bindingResourcePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BindingResourceProperties)(nil)).Elem()
}

func (i *bindingResourcePropertiesPtrType) ToBindingResourcePropertiesPtrOutput() BindingResourcePropertiesPtrOutput {
	return i.ToBindingResourcePropertiesPtrOutputWithContext(context.Background())
}

func (i *bindingResourcePropertiesPtrType) ToBindingResourcePropertiesPtrOutputWithContext(ctx context.Context) BindingResourcePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BindingResourcePropertiesPtrOutput)
}

type BindingResourcePropertiesOutput struct{ *pulumi.OutputState }

func (BindingResourcePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BindingResourceProperties)(nil)).Elem()
}

func (o BindingResourcePropertiesOutput) ToBindingResourcePropertiesOutput() BindingResourcePropertiesOutput {
	return o
}

func (o BindingResourcePropertiesOutput) ToBindingResourcePropertiesOutputWithContext(ctx context.Context) BindingResourcePropertiesOutput {
	return o
}

func (o BindingResourcePropertiesOutput) ToBindingResourcePropertiesPtrOutput() BindingResourcePropertiesPtrOutput {
	return o.ToBindingResourcePropertiesPtrOutputWithContext(context.Background())
}

func (o BindingResourcePropertiesOutput) ToBindingResourcePropertiesPtrOutputWithContext(ctx context.Context) BindingResourcePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BindingResourceProperties) *BindingResourceProperties {
		return &v
	}).(BindingResourcePropertiesPtrOutput)
}

func (o BindingResourcePropertiesOutput) BindingParameters() pulumi.MapOutput {
	return o.ApplyT(func(v BindingResourceProperties) map[string]interface{} { return v.BindingParameters }).(pulumi.MapOutput)
}

func (o BindingResourcePropertiesOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BindingResourceProperties) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o BindingResourcePropertiesOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BindingResourceProperties) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type BindingResourcePropertiesPtrOutput struct{ *pulumi.OutputState }

func (BindingResourcePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BindingResourceProperties)(nil)).Elem()
}

func (o BindingResourcePropertiesPtrOutput) ToBindingResourcePropertiesPtrOutput() BindingResourcePropertiesPtrOutput {
	return o
}

func (o BindingResourcePropertiesPtrOutput) ToBindingResourcePropertiesPtrOutputWithContext(ctx context.Context) BindingResourcePropertiesPtrOutput {
	return o
}

func (o BindingResourcePropertiesPtrOutput) Elem() BindingResourcePropertiesOutput {
	return o.ApplyT(func(v *BindingResourceProperties) BindingResourceProperties {
		if v != nil {
			return *v
		}
		var ret BindingResourceProperties
		return ret
	}).(BindingResourcePropertiesOutput)
}

func (o BindingResourcePropertiesPtrOutput) BindingParameters() pulumi.MapOutput {
	return o.ApplyT(func(v *BindingResourceProperties) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.BindingParameters
	}).(pulumi.MapOutput)
}

func (o BindingResourcePropertiesPtrOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BindingResourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.Key
	}).(pulumi.StringPtrOutput)
}

func (o BindingResourcePropertiesPtrOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BindingResourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.ResourceId
	}).(pulumi.StringPtrOutput)
}

type BindingResourcePropertiesResponse struct {
	BindingParameters   map[string]interface{} `pulumi:"bindingParameters"`
	CreatedAt           string                 `pulumi:"createdAt"`
	GeneratedProperties string                 `pulumi:"generatedProperties"`
	Key                 *string                `pulumi:"key"`
	ResourceId          *string                `pulumi:"resourceId"`
	ResourceName        string                 `pulumi:"resourceName"`
	ResourceType        string                 `pulumi:"resourceType"`
	UpdatedAt           string                 `pulumi:"updatedAt"`
}





type BindingResourcePropertiesResponseInput interface {
	pulumi.Input

	ToBindingResourcePropertiesResponseOutput() BindingResourcePropertiesResponseOutput
	ToBindingResourcePropertiesResponseOutputWithContext(context.Context) BindingResourcePropertiesResponseOutput
}

type BindingResourcePropertiesResponseArgs struct {
	BindingParameters   pulumi.MapInput       `pulumi:"bindingParameters"`
	CreatedAt           pulumi.StringInput    `pulumi:"createdAt"`
	GeneratedProperties pulumi.StringInput    `pulumi:"generatedProperties"`
	Key                 pulumi.StringPtrInput `pulumi:"key"`
	ResourceId          pulumi.StringPtrInput `pulumi:"resourceId"`
	ResourceName        pulumi.StringInput    `pulumi:"resourceName"`
	ResourceType        pulumi.StringInput    `pulumi:"resourceType"`
	UpdatedAt           pulumi.StringInput    `pulumi:"updatedAt"`
}

func (BindingResourcePropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BindingResourcePropertiesResponse)(nil)).Elem()
}

func (i BindingResourcePropertiesResponseArgs) ToBindingResourcePropertiesResponseOutput() BindingResourcePropertiesResponseOutput {
	return i.ToBindingResourcePropertiesResponseOutputWithContext(context.Background())
}

func (i BindingResourcePropertiesResponseArgs) ToBindingResourcePropertiesResponseOutputWithContext(ctx context.Context) BindingResourcePropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BindingResourcePropertiesResponseOutput)
}

func (i BindingResourcePropertiesResponseArgs) ToBindingResourcePropertiesResponsePtrOutput() BindingResourcePropertiesResponsePtrOutput {
	return i.ToBindingResourcePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i BindingResourcePropertiesResponseArgs) ToBindingResourcePropertiesResponsePtrOutputWithContext(ctx context.Context) BindingResourcePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BindingResourcePropertiesResponseOutput).ToBindingResourcePropertiesResponsePtrOutputWithContext(ctx)
}









type BindingResourcePropertiesResponsePtrInput interface {
	pulumi.Input

	ToBindingResourcePropertiesResponsePtrOutput() BindingResourcePropertiesResponsePtrOutput
	ToBindingResourcePropertiesResponsePtrOutputWithContext(context.Context) BindingResourcePropertiesResponsePtrOutput
}

type bindingResourcePropertiesResponsePtrType BindingResourcePropertiesResponseArgs

func BindingResourcePropertiesResponsePtr(v *BindingResourcePropertiesResponseArgs) BindingResourcePropertiesResponsePtrInput {
	return (*bindingResourcePropertiesResponsePtrType)(v)
}

func (*bindingResourcePropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BindingResourcePropertiesResponse)(nil)).Elem()
}

func (i *bindingResourcePropertiesResponsePtrType) ToBindingResourcePropertiesResponsePtrOutput() BindingResourcePropertiesResponsePtrOutput {
	return i.ToBindingResourcePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *bindingResourcePropertiesResponsePtrType) ToBindingResourcePropertiesResponsePtrOutputWithContext(ctx context.Context) BindingResourcePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BindingResourcePropertiesResponsePtrOutput)
}

type BindingResourcePropertiesResponseOutput struct{ *pulumi.OutputState }

func (BindingResourcePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BindingResourcePropertiesResponse)(nil)).Elem()
}

func (o BindingResourcePropertiesResponseOutput) ToBindingResourcePropertiesResponseOutput() BindingResourcePropertiesResponseOutput {
	return o
}

func (o BindingResourcePropertiesResponseOutput) ToBindingResourcePropertiesResponseOutputWithContext(ctx context.Context) BindingResourcePropertiesResponseOutput {
	return o
}

func (o BindingResourcePropertiesResponseOutput) ToBindingResourcePropertiesResponsePtrOutput() BindingResourcePropertiesResponsePtrOutput {
	return o.ToBindingResourcePropertiesResponsePtrOutputWithContext(context.Background())
}

func (o BindingResourcePropertiesResponseOutput) ToBindingResourcePropertiesResponsePtrOutputWithContext(ctx context.Context) BindingResourcePropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BindingResourcePropertiesResponse) *BindingResourcePropertiesResponse {
		return &v
	}).(BindingResourcePropertiesResponsePtrOutput)
}

func (o BindingResourcePropertiesResponseOutput) BindingParameters() pulumi.MapOutput {
	return o.ApplyT(func(v BindingResourcePropertiesResponse) map[string]interface{} { return v.BindingParameters }).(pulumi.MapOutput)
}

func (o BindingResourcePropertiesResponseOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v BindingResourcePropertiesResponse) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o BindingResourcePropertiesResponseOutput) GeneratedProperties() pulumi.StringOutput {
	return o.ApplyT(func(v BindingResourcePropertiesResponse) string { return v.GeneratedProperties }).(pulumi.StringOutput)
}

func (o BindingResourcePropertiesResponseOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BindingResourcePropertiesResponse) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o BindingResourcePropertiesResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BindingResourcePropertiesResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o BindingResourcePropertiesResponseOutput) ResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v BindingResourcePropertiesResponse) string { return v.ResourceName }).(pulumi.StringOutput)
}

func (o BindingResourcePropertiesResponseOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v BindingResourcePropertiesResponse) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o BindingResourcePropertiesResponseOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v BindingResourcePropertiesResponse) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

type BindingResourcePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (BindingResourcePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BindingResourcePropertiesResponse)(nil)).Elem()
}

func (o BindingResourcePropertiesResponsePtrOutput) ToBindingResourcePropertiesResponsePtrOutput() BindingResourcePropertiesResponsePtrOutput {
	return o
}

func (o BindingResourcePropertiesResponsePtrOutput) ToBindingResourcePropertiesResponsePtrOutputWithContext(ctx context.Context) BindingResourcePropertiesResponsePtrOutput {
	return o
}

func (o BindingResourcePropertiesResponsePtrOutput) Elem() BindingResourcePropertiesResponseOutput {
	return o.ApplyT(func(v *BindingResourcePropertiesResponse) BindingResourcePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret BindingResourcePropertiesResponse
		return ret
	}).(BindingResourcePropertiesResponseOutput)
}

func (o BindingResourcePropertiesResponsePtrOutput) BindingParameters() pulumi.MapOutput {
	return o.ApplyT(func(v *BindingResourcePropertiesResponse) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.BindingParameters
	}).(pulumi.MapOutput)
}

func (o BindingResourcePropertiesResponsePtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BindingResourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

func (o BindingResourcePropertiesResponsePtrOutput) GeneratedProperties() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BindingResourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.GeneratedProperties
	}).(pulumi.StringPtrOutput)
}

func (o BindingResourcePropertiesResponsePtrOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BindingResourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Key
	}).(pulumi.StringPtrOutput)
}

func (o BindingResourcePropertiesResponsePtrOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BindingResourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceId
	}).(pulumi.StringPtrOutput)
}

func (o BindingResourcePropertiesResponsePtrOutput) ResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BindingResourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ResourceName
	}).(pulumi.StringPtrOutput)
}

func (o BindingResourcePropertiesResponsePtrOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BindingResourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ResourceType
	}).(pulumi.StringPtrOutput)
}

func (o BindingResourcePropertiesResponsePtrOutput) UpdatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BindingResourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.UpdatedAt
	}).(pulumi.StringPtrOutput)
}

type CertificateProperties struct {
	CertVersion      *string `pulumi:"certVersion"`
	KeyVaultCertName string  `pulumi:"keyVaultCertName"`
	VaultUri         string  `pulumi:"vaultUri"`
}





type CertificatePropertiesInput interface {
	pulumi.Input

	ToCertificatePropertiesOutput() CertificatePropertiesOutput
	ToCertificatePropertiesOutputWithContext(context.Context) CertificatePropertiesOutput
}

type CertificatePropertiesArgs struct {
	CertVersion      pulumi.StringPtrInput `pulumi:"certVersion"`
	KeyVaultCertName pulumi.StringInput    `pulumi:"keyVaultCertName"`
	VaultUri         pulumi.StringInput    `pulumi:"vaultUri"`
}

func (CertificatePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateProperties)(nil)).Elem()
}

func (i CertificatePropertiesArgs) ToCertificatePropertiesOutput() CertificatePropertiesOutput {
	return i.ToCertificatePropertiesOutputWithContext(context.Background())
}

func (i CertificatePropertiesArgs) ToCertificatePropertiesOutputWithContext(ctx context.Context) CertificatePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificatePropertiesOutput)
}

func (i CertificatePropertiesArgs) ToCertificatePropertiesPtrOutput() CertificatePropertiesPtrOutput {
	return i.ToCertificatePropertiesPtrOutputWithContext(context.Background())
}

func (i CertificatePropertiesArgs) ToCertificatePropertiesPtrOutputWithContext(ctx context.Context) CertificatePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificatePropertiesOutput).ToCertificatePropertiesPtrOutputWithContext(ctx)
}









type CertificatePropertiesPtrInput interface {
	pulumi.Input

	ToCertificatePropertiesPtrOutput() CertificatePropertiesPtrOutput
	ToCertificatePropertiesPtrOutputWithContext(context.Context) CertificatePropertiesPtrOutput
}

type certificatePropertiesPtrType CertificatePropertiesArgs

func CertificatePropertiesPtr(v *CertificatePropertiesArgs) CertificatePropertiesPtrInput {
	return (*certificatePropertiesPtrType)(v)
}

func (*certificatePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateProperties)(nil)).Elem()
}

func (i *certificatePropertiesPtrType) ToCertificatePropertiesPtrOutput() CertificatePropertiesPtrOutput {
	return i.ToCertificatePropertiesPtrOutputWithContext(context.Background())
}

func (i *certificatePropertiesPtrType) ToCertificatePropertiesPtrOutputWithContext(ctx context.Context) CertificatePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificatePropertiesPtrOutput)
}

type CertificatePropertiesOutput struct{ *pulumi.OutputState }

func (CertificatePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateProperties)(nil)).Elem()
}

func (o CertificatePropertiesOutput) ToCertificatePropertiesOutput() CertificatePropertiesOutput {
	return o
}

func (o CertificatePropertiesOutput) ToCertificatePropertiesOutputWithContext(ctx context.Context) CertificatePropertiesOutput {
	return o
}

func (o CertificatePropertiesOutput) ToCertificatePropertiesPtrOutput() CertificatePropertiesPtrOutput {
	return o.ToCertificatePropertiesPtrOutputWithContext(context.Background())
}

func (o CertificatePropertiesOutput) ToCertificatePropertiesPtrOutputWithContext(ctx context.Context) CertificatePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CertificateProperties) *CertificateProperties {
		return &v
	}).(CertificatePropertiesPtrOutput)
}

func (o CertificatePropertiesOutput) CertVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertificateProperties) *string { return v.CertVersion }).(pulumi.StringPtrOutput)
}

func (o CertificatePropertiesOutput) KeyVaultCertName() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateProperties) string { return v.KeyVaultCertName }).(pulumi.StringOutput)
}

func (o CertificatePropertiesOutput) VaultUri() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateProperties) string { return v.VaultUri }).(pulumi.StringOutput)
}

type CertificatePropertiesPtrOutput struct{ *pulumi.OutputState }

func (CertificatePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateProperties)(nil)).Elem()
}

func (o CertificatePropertiesPtrOutput) ToCertificatePropertiesPtrOutput() CertificatePropertiesPtrOutput {
	return o
}

func (o CertificatePropertiesPtrOutput) ToCertificatePropertiesPtrOutputWithContext(ctx context.Context) CertificatePropertiesPtrOutput {
	return o
}

func (o CertificatePropertiesPtrOutput) Elem() CertificatePropertiesOutput {
	return o.ApplyT(func(v *CertificateProperties) CertificateProperties {
		if v != nil {
			return *v
		}
		var ret CertificateProperties
		return ret
	}).(CertificatePropertiesOutput)
}

func (o CertificatePropertiesPtrOutput) CertVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateProperties) *string {
		if v == nil {
			return nil
		}
		return v.CertVersion
	}).(pulumi.StringPtrOutput)
}

func (o CertificatePropertiesPtrOutput) KeyVaultCertName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateProperties) *string {
		if v == nil {
			return nil
		}
		return &v.KeyVaultCertName
	}).(pulumi.StringPtrOutput)
}

func (o CertificatePropertiesPtrOutput) VaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificateProperties) *string {
		if v == nil {
			return nil
		}
		return &v.VaultUri
	}).(pulumi.StringPtrOutput)
}

type CertificatePropertiesResponse struct {
	ActivateDate     string   `pulumi:"activateDate"`
	CertVersion      *string  `pulumi:"certVersion"`
	DnsNames         []string `pulumi:"dnsNames"`
	ExpirationDate   string   `pulumi:"expirationDate"`
	IssuedDate       string   `pulumi:"issuedDate"`
	Issuer           string   `pulumi:"issuer"`
	KeyVaultCertName string   `pulumi:"keyVaultCertName"`
	SubjectName      string   `pulumi:"subjectName"`
	Thumbprint       string   `pulumi:"thumbprint"`
	VaultUri         string   `pulumi:"vaultUri"`
}





type CertificatePropertiesResponseInput interface {
	pulumi.Input

	ToCertificatePropertiesResponseOutput() CertificatePropertiesResponseOutput
	ToCertificatePropertiesResponseOutputWithContext(context.Context) CertificatePropertiesResponseOutput
}

type CertificatePropertiesResponseArgs struct {
	ActivateDate     pulumi.StringInput      `pulumi:"activateDate"`
	CertVersion      pulumi.StringPtrInput   `pulumi:"certVersion"`
	DnsNames         pulumi.StringArrayInput `pulumi:"dnsNames"`
	ExpirationDate   pulumi.StringInput      `pulumi:"expirationDate"`
	IssuedDate       pulumi.StringInput      `pulumi:"issuedDate"`
	Issuer           pulumi.StringInput      `pulumi:"issuer"`
	KeyVaultCertName pulumi.StringInput      `pulumi:"keyVaultCertName"`
	SubjectName      pulumi.StringInput      `pulumi:"subjectName"`
	Thumbprint       pulumi.StringInput      `pulumi:"thumbprint"`
	VaultUri         pulumi.StringInput      `pulumi:"vaultUri"`
}

func (CertificatePropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificatePropertiesResponse)(nil)).Elem()
}

func (i CertificatePropertiesResponseArgs) ToCertificatePropertiesResponseOutput() CertificatePropertiesResponseOutput {
	return i.ToCertificatePropertiesResponseOutputWithContext(context.Background())
}

func (i CertificatePropertiesResponseArgs) ToCertificatePropertiesResponseOutputWithContext(ctx context.Context) CertificatePropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificatePropertiesResponseOutput)
}

func (i CertificatePropertiesResponseArgs) ToCertificatePropertiesResponsePtrOutput() CertificatePropertiesResponsePtrOutput {
	return i.ToCertificatePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i CertificatePropertiesResponseArgs) ToCertificatePropertiesResponsePtrOutputWithContext(ctx context.Context) CertificatePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificatePropertiesResponseOutput).ToCertificatePropertiesResponsePtrOutputWithContext(ctx)
}









type CertificatePropertiesResponsePtrInput interface {
	pulumi.Input

	ToCertificatePropertiesResponsePtrOutput() CertificatePropertiesResponsePtrOutput
	ToCertificatePropertiesResponsePtrOutputWithContext(context.Context) CertificatePropertiesResponsePtrOutput
}

type certificatePropertiesResponsePtrType CertificatePropertiesResponseArgs

func CertificatePropertiesResponsePtr(v *CertificatePropertiesResponseArgs) CertificatePropertiesResponsePtrInput {
	return (*certificatePropertiesResponsePtrType)(v)
}

func (*certificatePropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificatePropertiesResponse)(nil)).Elem()
}

func (i *certificatePropertiesResponsePtrType) ToCertificatePropertiesResponsePtrOutput() CertificatePropertiesResponsePtrOutput {
	return i.ToCertificatePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *certificatePropertiesResponsePtrType) ToCertificatePropertiesResponsePtrOutputWithContext(ctx context.Context) CertificatePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificatePropertiesResponsePtrOutput)
}

type CertificatePropertiesResponseOutput struct{ *pulumi.OutputState }

func (CertificatePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificatePropertiesResponse)(nil)).Elem()
}

func (o CertificatePropertiesResponseOutput) ToCertificatePropertiesResponseOutput() CertificatePropertiesResponseOutput {
	return o
}

func (o CertificatePropertiesResponseOutput) ToCertificatePropertiesResponseOutputWithContext(ctx context.Context) CertificatePropertiesResponseOutput {
	return o
}

func (o CertificatePropertiesResponseOutput) ToCertificatePropertiesResponsePtrOutput() CertificatePropertiesResponsePtrOutput {
	return o.ToCertificatePropertiesResponsePtrOutputWithContext(context.Background())
}

func (o CertificatePropertiesResponseOutput) ToCertificatePropertiesResponsePtrOutputWithContext(ctx context.Context) CertificatePropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CertificatePropertiesResponse) *CertificatePropertiesResponse {
		return &v
	}).(CertificatePropertiesResponsePtrOutput)
}

func (o CertificatePropertiesResponseOutput) ActivateDate() pulumi.StringOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) string { return v.ActivateDate }).(pulumi.StringOutput)
}

func (o CertificatePropertiesResponseOutput) CertVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) *string { return v.CertVersion }).(pulumi.StringPtrOutput)
}

func (o CertificatePropertiesResponseOutput) DnsNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) []string { return v.DnsNames }).(pulumi.StringArrayOutput)
}

func (o CertificatePropertiesResponseOutput) ExpirationDate() pulumi.StringOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) string { return v.ExpirationDate }).(pulumi.StringOutput)
}

func (o CertificatePropertiesResponseOutput) IssuedDate() pulumi.StringOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) string { return v.IssuedDate }).(pulumi.StringOutput)
}

func (o CertificatePropertiesResponseOutput) Issuer() pulumi.StringOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) string { return v.Issuer }).(pulumi.StringOutput)
}

func (o CertificatePropertiesResponseOutput) KeyVaultCertName() pulumi.StringOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) string { return v.KeyVaultCertName }).(pulumi.StringOutput)
}

func (o CertificatePropertiesResponseOutput) SubjectName() pulumi.StringOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) string { return v.SubjectName }).(pulumi.StringOutput)
}

func (o CertificatePropertiesResponseOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) string { return v.Thumbprint }).(pulumi.StringOutput)
}

func (o CertificatePropertiesResponseOutput) VaultUri() pulumi.StringOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) string { return v.VaultUri }).(pulumi.StringOutput)
}

type CertificatePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (CertificatePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificatePropertiesResponse)(nil)).Elem()
}

func (o CertificatePropertiesResponsePtrOutput) ToCertificatePropertiesResponsePtrOutput() CertificatePropertiesResponsePtrOutput {
	return o
}

func (o CertificatePropertiesResponsePtrOutput) ToCertificatePropertiesResponsePtrOutputWithContext(ctx context.Context) CertificatePropertiesResponsePtrOutput {
	return o
}

func (o CertificatePropertiesResponsePtrOutput) Elem() CertificatePropertiesResponseOutput {
	return o.ApplyT(func(v *CertificatePropertiesResponse) CertificatePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret CertificatePropertiesResponse
		return ret
	}).(CertificatePropertiesResponseOutput)
}

func (o CertificatePropertiesResponsePtrOutput) ActivateDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificatePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ActivateDate
	}).(pulumi.StringPtrOutput)
}

func (o CertificatePropertiesResponsePtrOutput) CertVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificatePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.CertVersion
	}).(pulumi.StringPtrOutput)
}

func (o CertificatePropertiesResponsePtrOutput) DnsNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CertificatePropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.DnsNames
	}).(pulumi.StringArrayOutput)
}

func (o CertificatePropertiesResponsePtrOutput) ExpirationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificatePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ExpirationDate
	}).(pulumi.StringPtrOutput)
}

func (o CertificatePropertiesResponsePtrOutput) IssuedDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificatePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.IssuedDate
	}).(pulumi.StringPtrOutput)
}

func (o CertificatePropertiesResponsePtrOutput) Issuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificatePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Issuer
	}).(pulumi.StringPtrOutput)
}

func (o CertificatePropertiesResponsePtrOutput) KeyVaultCertName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificatePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.KeyVaultCertName
	}).(pulumi.StringPtrOutput)
}

func (o CertificatePropertiesResponsePtrOutput) SubjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificatePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SubjectName
	}).(pulumi.StringPtrOutput)
}

func (o CertificatePropertiesResponsePtrOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificatePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Thumbprint
	}).(pulumi.StringPtrOutput)
}

func (o CertificatePropertiesResponsePtrOutput) VaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificatePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.VaultUri
	}).(pulumi.StringPtrOutput)
}

type ClusterResourceProperties struct {
	NetworkProfile *NetworkProfile `pulumi:"networkProfile"`
}





type ClusterResourcePropertiesInput interface {
	pulumi.Input

	ToClusterResourcePropertiesOutput() ClusterResourcePropertiesOutput
	ToClusterResourcePropertiesOutputWithContext(context.Context) ClusterResourcePropertiesOutput
}

type ClusterResourcePropertiesArgs struct {
	NetworkProfile NetworkProfilePtrInput `pulumi:"networkProfile"`
}

func (ClusterResourcePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterResourceProperties)(nil)).Elem()
}

func (i ClusterResourcePropertiesArgs) ToClusterResourcePropertiesOutput() ClusterResourcePropertiesOutput {
	return i.ToClusterResourcePropertiesOutputWithContext(context.Background())
}

func (i ClusterResourcePropertiesArgs) ToClusterResourcePropertiesOutputWithContext(ctx context.Context) ClusterResourcePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterResourcePropertiesOutput)
}

func (i ClusterResourcePropertiesArgs) ToClusterResourcePropertiesPtrOutput() ClusterResourcePropertiesPtrOutput {
	return i.ToClusterResourcePropertiesPtrOutputWithContext(context.Background())
}

func (i ClusterResourcePropertiesArgs) ToClusterResourcePropertiesPtrOutputWithContext(ctx context.Context) ClusterResourcePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterResourcePropertiesOutput).ToClusterResourcePropertiesPtrOutputWithContext(ctx)
}









type ClusterResourcePropertiesPtrInput interface {
	pulumi.Input

	ToClusterResourcePropertiesPtrOutput() ClusterResourcePropertiesPtrOutput
	ToClusterResourcePropertiesPtrOutputWithContext(context.Context) ClusterResourcePropertiesPtrOutput
}

type clusterResourcePropertiesPtrType ClusterResourcePropertiesArgs

func ClusterResourcePropertiesPtr(v *ClusterResourcePropertiesArgs) ClusterResourcePropertiesPtrInput {
	return (*clusterResourcePropertiesPtrType)(v)
}

func (*clusterResourcePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterResourceProperties)(nil)).Elem()
}

func (i *clusterResourcePropertiesPtrType) ToClusterResourcePropertiesPtrOutput() ClusterResourcePropertiesPtrOutput {
	return i.ToClusterResourcePropertiesPtrOutputWithContext(context.Background())
}

func (i *clusterResourcePropertiesPtrType) ToClusterResourcePropertiesPtrOutputWithContext(ctx context.Context) ClusterResourcePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterResourcePropertiesPtrOutput)
}

type ClusterResourcePropertiesOutput struct{ *pulumi.OutputState }

func (ClusterResourcePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterResourceProperties)(nil)).Elem()
}

func (o ClusterResourcePropertiesOutput) ToClusterResourcePropertiesOutput() ClusterResourcePropertiesOutput {
	return o
}

func (o ClusterResourcePropertiesOutput) ToClusterResourcePropertiesOutputWithContext(ctx context.Context) ClusterResourcePropertiesOutput {
	return o
}

func (o ClusterResourcePropertiesOutput) ToClusterResourcePropertiesPtrOutput() ClusterResourcePropertiesPtrOutput {
	return o.ToClusterResourcePropertiesPtrOutputWithContext(context.Background())
}

func (o ClusterResourcePropertiesOutput) ToClusterResourcePropertiesPtrOutputWithContext(ctx context.Context) ClusterResourcePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterResourceProperties) *ClusterResourceProperties {
		return &v
	}).(ClusterResourcePropertiesPtrOutput)
}

func (o ClusterResourcePropertiesOutput) NetworkProfile() NetworkProfilePtrOutput {
	return o.ApplyT(func(v ClusterResourceProperties) *NetworkProfile { return v.NetworkProfile }).(NetworkProfilePtrOutput)
}

type ClusterResourcePropertiesPtrOutput struct{ *pulumi.OutputState }

func (ClusterResourcePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterResourceProperties)(nil)).Elem()
}

func (o ClusterResourcePropertiesPtrOutput) ToClusterResourcePropertiesPtrOutput() ClusterResourcePropertiesPtrOutput {
	return o
}

func (o ClusterResourcePropertiesPtrOutput) ToClusterResourcePropertiesPtrOutputWithContext(ctx context.Context) ClusterResourcePropertiesPtrOutput {
	return o
}

func (o ClusterResourcePropertiesPtrOutput) Elem() ClusterResourcePropertiesOutput {
	return o.ApplyT(func(v *ClusterResourceProperties) ClusterResourceProperties {
		if v != nil {
			return *v
		}
		var ret ClusterResourceProperties
		return ret
	}).(ClusterResourcePropertiesOutput)
}

func (o ClusterResourcePropertiesPtrOutput) NetworkProfile() NetworkProfilePtrOutput {
	return o.ApplyT(func(v *ClusterResourceProperties) *NetworkProfile {
		if v == nil {
			return nil
		}
		return v.NetworkProfile
	}).(NetworkProfilePtrOutput)
}

type ClusterResourcePropertiesResponse struct {
	NetworkProfile    *NetworkProfileResponse `pulumi:"networkProfile"`
	ProvisioningState string                  `pulumi:"provisioningState"`
	ServiceId         string                  `pulumi:"serviceId"`
	Version           int                     `pulumi:"version"`
}





type ClusterResourcePropertiesResponseInput interface {
	pulumi.Input

	ToClusterResourcePropertiesResponseOutput() ClusterResourcePropertiesResponseOutput
	ToClusterResourcePropertiesResponseOutputWithContext(context.Context) ClusterResourcePropertiesResponseOutput
}

type ClusterResourcePropertiesResponseArgs struct {
	NetworkProfile    NetworkProfileResponsePtrInput `pulumi:"networkProfile"`
	ProvisioningState pulumi.StringInput             `pulumi:"provisioningState"`
	ServiceId         pulumi.StringInput             `pulumi:"serviceId"`
	Version           pulumi.IntInput                `pulumi:"version"`
}

func (ClusterResourcePropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterResourcePropertiesResponse)(nil)).Elem()
}

func (i ClusterResourcePropertiesResponseArgs) ToClusterResourcePropertiesResponseOutput() ClusterResourcePropertiesResponseOutput {
	return i.ToClusterResourcePropertiesResponseOutputWithContext(context.Background())
}

func (i ClusterResourcePropertiesResponseArgs) ToClusterResourcePropertiesResponseOutputWithContext(ctx context.Context) ClusterResourcePropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterResourcePropertiesResponseOutput)
}

func (i ClusterResourcePropertiesResponseArgs) ToClusterResourcePropertiesResponsePtrOutput() ClusterResourcePropertiesResponsePtrOutput {
	return i.ToClusterResourcePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ClusterResourcePropertiesResponseArgs) ToClusterResourcePropertiesResponsePtrOutputWithContext(ctx context.Context) ClusterResourcePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterResourcePropertiesResponseOutput).ToClusterResourcePropertiesResponsePtrOutputWithContext(ctx)
}









type ClusterResourcePropertiesResponsePtrInput interface {
	pulumi.Input

	ToClusterResourcePropertiesResponsePtrOutput() ClusterResourcePropertiesResponsePtrOutput
	ToClusterResourcePropertiesResponsePtrOutputWithContext(context.Context) ClusterResourcePropertiesResponsePtrOutput
}

type clusterResourcePropertiesResponsePtrType ClusterResourcePropertiesResponseArgs

func ClusterResourcePropertiesResponsePtr(v *ClusterResourcePropertiesResponseArgs) ClusterResourcePropertiesResponsePtrInput {
	return (*clusterResourcePropertiesResponsePtrType)(v)
}

func (*clusterResourcePropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterResourcePropertiesResponse)(nil)).Elem()
}

func (i *clusterResourcePropertiesResponsePtrType) ToClusterResourcePropertiesResponsePtrOutput() ClusterResourcePropertiesResponsePtrOutput {
	return i.ToClusterResourcePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *clusterResourcePropertiesResponsePtrType) ToClusterResourcePropertiesResponsePtrOutputWithContext(ctx context.Context) ClusterResourcePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterResourcePropertiesResponsePtrOutput)
}

type ClusterResourcePropertiesResponseOutput struct{ *pulumi.OutputState }

func (ClusterResourcePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterResourcePropertiesResponse)(nil)).Elem()
}

func (o ClusterResourcePropertiesResponseOutput) ToClusterResourcePropertiesResponseOutput() ClusterResourcePropertiesResponseOutput {
	return o
}

func (o ClusterResourcePropertiesResponseOutput) ToClusterResourcePropertiesResponseOutputWithContext(ctx context.Context) ClusterResourcePropertiesResponseOutput {
	return o
}

func (o ClusterResourcePropertiesResponseOutput) ToClusterResourcePropertiesResponsePtrOutput() ClusterResourcePropertiesResponsePtrOutput {
	return o.ToClusterResourcePropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ClusterResourcePropertiesResponseOutput) ToClusterResourcePropertiesResponsePtrOutputWithContext(ctx context.Context) ClusterResourcePropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterResourcePropertiesResponse) *ClusterResourcePropertiesResponse {
		return &v
	}).(ClusterResourcePropertiesResponsePtrOutput)
}

func (o ClusterResourcePropertiesResponseOutput) NetworkProfile() NetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v ClusterResourcePropertiesResponse) *NetworkProfileResponse { return v.NetworkProfile }).(NetworkProfileResponsePtrOutput)
}

func (o ClusterResourcePropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterResourcePropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ClusterResourcePropertiesResponseOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterResourcePropertiesResponse) string { return v.ServiceId }).(pulumi.StringOutput)
}

func (o ClusterResourcePropertiesResponseOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v ClusterResourcePropertiesResponse) int { return v.Version }).(pulumi.IntOutput)
}

type ClusterResourcePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ClusterResourcePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterResourcePropertiesResponse)(nil)).Elem()
}

func (o ClusterResourcePropertiesResponsePtrOutput) ToClusterResourcePropertiesResponsePtrOutput() ClusterResourcePropertiesResponsePtrOutput {
	return o
}

func (o ClusterResourcePropertiesResponsePtrOutput) ToClusterResourcePropertiesResponsePtrOutputWithContext(ctx context.Context) ClusterResourcePropertiesResponsePtrOutput {
	return o
}

func (o ClusterResourcePropertiesResponsePtrOutput) Elem() ClusterResourcePropertiesResponseOutput {
	return o.ApplyT(func(v *ClusterResourcePropertiesResponse) ClusterResourcePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ClusterResourcePropertiesResponse
		return ret
	}).(ClusterResourcePropertiesResponseOutput)
}

func (o ClusterResourcePropertiesResponsePtrOutput) NetworkProfile() NetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v *ClusterResourcePropertiesResponse) *NetworkProfileResponse {
		if v == nil {
			return nil
		}
		return v.NetworkProfile
	}).(NetworkProfileResponsePtrOutput)
}

func (o ClusterResourcePropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterResourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o ClusterResourcePropertiesResponsePtrOutput) ServiceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterResourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ServiceId
	}).(pulumi.StringPtrOutput)
}

func (o ClusterResourcePropertiesResponsePtrOutput) Version() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ClusterResourcePropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return &v.Version
	}).(pulumi.IntPtrOutput)
}

type CustomContainer struct {
	Args                    []string                 `pulumi:"args"`
	Command                 []string                 `pulumi:"command"`
	ContainerImage          *string                  `pulumi:"containerImage"`
	ImageRegistryCredential *ImageRegistryCredential `pulumi:"imageRegistryCredential"`
	Server                  *string                  `pulumi:"server"`
}





type CustomContainerInput interface {
	pulumi.Input

	ToCustomContainerOutput() CustomContainerOutput
	ToCustomContainerOutputWithContext(context.Context) CustomContainerOutput
}

type CustomContainerArgs struct {
	Args                    pulumi.StringArrayInput         `pulumi:"args"`
	Command                 pulumi.StringArrayInput         `pulumi:"command"`
	ContainerImage          pulumi.StringPtrInput           `pulumi:"containerImage"`
	ImageRegistryCredential ImageRegistryCredentialPtrInput `pulumi:"imageRegistryCredential"`
	Server                  pulumi.StringPtrInput           `pulumi:"server"`
}

func (CustomContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomContainer)(nil)).Elem()
}

func (i CustomContainerArgs) ToCustomContainerOutput() CustomContainerOutput {
	return i.ToCustomContainerOutputWithContext(context.Background())
}

func (i CustomContainerArgs) ToCustomContainerOutputWithContext(ctx context.Context) CustomContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomContainerOutput)
}

func (i CustomContainerArgs) ToCustomContainerPtrOutput() CustomContainerPtrOutput {
	return i.ToCustomContainerPtrOutputWithContext(context.Background())
}

func (i CustomContainerArgs) ToCustomContainerPtrOutputWithContext(ctx context.Context) CustomContainerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomContainerOutput).ToCustomContainerPtrOutputWithContext(ctx)
}









type CustomContainerPtrInput interface {
	pulumi.Input

	ToCustomContainerPtrOutput() CustomContainerPtrOutput
	ToCustomContainerPtrOutputWithContext(context.Context) CustomContainerPtrOutput
}

type customContainerPtrType CustomContainerArgs

func CustomContainerPtr(v *CustomContainerArgs) CustomContainerPtrInput {
	return (*customContainerPtrType)(v)
}

func (*customContainerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomContainer)(nil)).Elem()
}

func (i *customContainerPtrType) ToCustomContainerPtrOutput() CustomContainerPtrOutput {
	return i.ToCustomContainerPtrOutputWithContext(context.Background())
}

func (i *customContainerPtrType) ToCustomContainerPtrOutputWithContext(ctx context.Context) CustomContainerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomContainerPtrOutput)
}

type CustomContainerOutput struct{ *pulumi.OutputState }

func (CustomContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomContainer)(nil)).Elem()
}

func (o CustomContainerOutput) ToCustomContainerOutput() CustomContainerOutput {
	return o
}

func (o CustomContainerOutput) ToCustomContainerOutputWithContext(ctx context.Context) CustomContainerOutput {
	return o
}

func (o CustomContainerOutput) ToCustomContainerPtrOutput() CustomContainerPtrOutput {
	return o.ToCustomContainerPtrOutputWithContext(context.Background())
}

func (o CustomContainerOutput) ToCustomContainerPtrOutputWithContext(ctx context.Context) CustomContainerPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomContainer) *CustomContainer {
		return &v
	}).(CustomContainerPtrOutput)
}

func (o CustomContainerOutput) Args() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CustomContainer) []string { return v.Args }).(pulumi.StringArrayOutput)
}

func (o CustomContainerOutput) Command() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CustomContainer) []string { return v.Command }).(pulumi.StringArrayOutput)
}

func (o CustomContainerOutput) ContainerImage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomContainer) *string { return v.ContainerImage }).(pulumi.StringPtrOutput)
}

func (o CustomContainerOutput) ImageRegistryCredential() ImageRegistryCredentialPtrOutput {
	return o.ApplyT(func(v CustomContainer) *ImageRegistryCredential { return v.ImageRegistryCredential }).(ImageRegistryCredentialPtrOutput)
}

func (o CustomContainerOutput) Server() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomContainer) *string { return v.Server }).(pulumi.StringPtrOutput)
}

type CustomContainerPtrOutput struct{ *pulumi.OutputState }

func (CustomContainerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomContainer)(nil)).Elem()
}

func (o CustomContainerPtrOutput) ToCustomContainerPtrOutput() CustomContainerPtrOutput {
	return o
}

func (o CustomContainerPtrOutput) ToCustomContainerPtrOutputWithContext(ctx context.Context) CustomContainerPtrOutput {
	return o
}

func (o CustomContainerPtrOutput) Elem() CustomContainerOutput {
	return o.ApplyT(func(v *CustomContainer) CustomContainer {
		if v != nil {
			return *v
		}
		var ret CustomContainer
		return ret
	}).(CustomContainerOutput)
}

func (o CustomContainerPtrOutput) Args() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CustomContainer) []string {
		if v == nil {
			return nil
		}
		return v.Args
	}).(pulumi.StringArrayOutput)
}

func (o CustomContainerPtrOutput) Command() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CustomContainer) []string {
		if v == nil {
			return nil
		}
		return v.Command
	}).(pulumi.StringArrayOutput)
}

func (o CustomContainerPtrOutput) ContainerImage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomContainer) *string {
		if v == nil {
			return nil
		}
		return v.ContainerImage
	}).(pulumi.StringPtrOutput)
}

func (o CustomContainerPtrOutput) ImageRegistryCredential() ImageRegistryCredentialPtrOutput {
	return o.ApplyT(func(v *CustomContainer) *ImageRegistryCredential {
		if v == nil {
			return nil
		}
		return v.ImageRegistryCredential
	}).(ImageRegistryCredentialPtrOutput)
}

func (o CustomContainerPtrOutput) Server() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomContainer) *string {
		if v == nil {
			return nil
		}
		return v.Server
	}).(pulumi.StringPtrOutput)
}

type CustomContainerResponse struct {
	Args                    []string                         `pulumi:"args"`
	Command                 []string                         `pulumi:"command"`
	ContainerImage          *string                          `pulumi:"containerImage"`
	ImageRegistryCredential *ImageRegistryCredentialResponse `pulumi:"imageRegistryCredential"`
	Server                  *string                          `pulumi:"server"`
}





type CustomContainerResponseInput interface {
	pulumi.Input

	ToCustomContainerResponseOutput() CustomContainerResponseOutput
	ToCustomContainerResponseOutputWithContext(context.Context) CustomContainerResponseOutput
}

type CustomContainerResponseArgs struct {
	Args                    pulumi.StringArrayInput                 `pulumi:"args"`
	Command                 pulumi.StringArrayInput                 `pulumi:"command"`
	ContainerImage          pulumi.StringPtrInput                   `pulumi:"containerImage"`
	ImageRegistryCredential ImageRegistryCredentialResponsePtrInput `pulumi:"imageRegistryCredential"`
	Server                  pulumi.StringPtrInput                   `pulumi:"server"`
}

func (CustomContainerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomContainerResponse)(nil)).Elem()
}

func (i CustomContainerResponseArgs) ToCustomContainerResponseOutput() CustomContainerResponseOutput {
	return i.ToCustomContainerResponseOutputWithContext(context.Background())
}

func (i CustomContainerResponseArgs) ToCustomContainerResponseOutputWithContext(ctx context.Context) CustomContainerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomContainerResponseOutput)
}

func (i CustomContainerResponseArgs) ToCustomContainerResponsePtrOutput() CustomContainerResponsePtrOutput {
	return i.ToCustomContainerResponsePtrOutputWithContext(context.Background())
}

func (i CustomContainerResponseArgs) ToCustomContainerResponsePtrOutputWithContext(ctx context.Context) CustomContainerResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomContainerResponseOutput).ToCustomContainerResponsePtrOutputWithContext(ctx)
}









type CustomContainerResponsePtrInput interface {
	pulumi.Input

	ToCustomContainerResponsePtrOutput() CustomContainerResponsePtrOutput
	ToCustomContainerResponsePtrOutputWithContext(context.Context) CustomContainerResponsePtrOutput
}

type customContainerResponsePtrType CustomContainerResponseArgs

func CustomContainerResponsePtr(v *CustomContainerResponseArgs) CustomContainerResponsePtrInput {
	return (*customContainerResponsePtrType)(v)
}

func (*customContainerResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomContainerResponse)(nil)).Elem()
}

func (i *customContainerResponsePtrType) ToCustomContainerResponsePtrOutput() CustomContainerResponsePtrOutput {
	return i.ToCustomContainerResponsePtrOutputWithContext(context.Background())
}

func (i *customContainerResponsePtrType) ToCustomContainerResponsePtrOutputWithContext(ctx context.Context) CustomContainerResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomContainerResponsePtrOutput)
}

type CustomContainerResponseOutput struct{ *pulumi.OutputState }

func (CustomContainerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomContainerResponse)(nil)).Elem()
}

func (o CustomContainerResponseOutput) ToCustomContainerResponseOutput() CustomContainerResponseOutput {
	return o
}

func (o CustomContainerResponseOutput) ToCustomContainerResponseOutputWithContext(ctx context.Context) CustomContainerResponseOutput {
	return o
}

func (o CustomContainerResponseOutput) ToCustomContainerResponsePtrOutput() CustomContainerResponsePtrOutput {
	return o.ToCustomContainerResponsePtrOutputWithContext(context.Background())
}

func (o CustomContainerResponseOutput) ToCustomContainerResponsePtrOutputWithContext(ctx context.Context) CustomContainerResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomContainerResponse) *CustomContainerResponse {
		return &v
	}).(CustomContainerResponsePtrOutput)
}

func (o CustomContainerResponseOutput) Args() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CustomContainerResponse) []string { return v.Args }).(pulumi.StringArrayOutput)
}

func (o CustomContainerResponseOutput) Command() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CustomContainerResponse) []string { return v.Command }).(pulumi.StringArrayOutput)
}

func (o CustomContainerResponseOutput) ContainerImage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomContainerResponse) *string { return v.ContainerImage }).(pulumi.StringPtrOutput)
}

func (o CustomContainerResponseOutput) ImageRegistryCredential() ImageRegistryCredentialResponsePtrOutput {
	return o.ApplyT(func(v CustomContainerResponse) *ImageRegistryCredentialResponse { return v.ImageRegistryCredential }).(ImageRegistryCredentialResponsePtrOutput)
}

func (o CustomContainerResponseOutput) Server() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomContainerResponse) *string { return v.Server }).(pulumi.StringPtrOutput)
}

type CustomContainerResponsePtrOutput struct{ *pulumi.OutputState }

func (CustomContainerResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomContainerResponse)(nil)).Elem()
}

func (o CustomContainerResponsePtrOutput) ToCustomContainerResponsePtrOutput() CustomContainerResponsePtrOutput {
	return o
}

func (o CustomContainerResponsePtrOutput) ToCustomContainerResponsePtrOutputWithContext(ctx context.Context) CustomContainerResponsePtrOutput {
	return o
}

func (o CustomContainerResponsePtrOutput) Elem() CustomContainerResponseOutput {
	return o.ApplyT(func(v *CustomContainerResponse) CustomContainerResponse {
		if v != nil {
			return *v
		}
		var ret CustomContainerResponse
		return ret
	}).(CustomContainerResponseOutput)
}

func (o CustomContainerResponsePtrOutput) Args() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CustomContainerResponse) []string {
		if v == nil {
			return nil
		}
		return v.Args
	}).(pulumi.StringArrayOutput)
}

func (o CustomContainerResponsePtrOutput) Command() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CustomContainerResponse) []string {
		if v == nil {
			return nil
		}
		return v.Command
	}).(pulumi.StringArrayOutput)
}

func (o CustomContainerResponsePtrOutput) ContainerImage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomContainerResponse) *string {
		if v == nil {
			return nil
		}
		return v.ContainerImage
	}).(pulumi.StringPtrOutput)
}

func (o CustomContainerResponsePtrOutput) ImageRegistryCredential() ImageRegistryCredentialResponsePtrOutput {
	return o.ApplyT(func(v *CustomContainerResponse) *ImageRegistryCredentialResponse {
		if v == nil {
			return nil
		}
		return v.ImageRegistryCredential
	}).(ImageRegistryCredentialResponsePtrOutput)
}

func (o CustomContainerResponsePtrOutput) Server() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomContainerResponse) *string {
		if v == nil {
			return nil
		}
		return v.Server
	}).(pulumi.StringPtrOutput)
}

type CustomDomainProperties struct {
	CertName   *string `pulumi:"certName"`
	Thumbprint *string `pulumi:"thumbprint"`
}





type CustomDomainPropertiesInput interface {
	pulumi.Input

	ToCustomDomainPropertiesOutput() CustomDomainPropertiesOutput
	ToCustomDomainPropertiesOutputWithContext(context.Context) CustomDomainPropertiesOutput
}

type CustomDomainPropertiesArgs struct {
	CertName   pulumi.StringPtrInput `pulumi:"certName"`
	Thumbprint pulumi.StringPtrInput `pulumi:"thumbprint"`
}

func (CustomDomainPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomDomainProperties)(nil)).Elem()
}

func (i CustomDomainPropertiesArgs) ToCustomDomainPropertiesOutput() CustomDomainPropertiesOutput {
	return i.ToCustomDomainPropertiesOutputWithContext(context.Background())
}

func (i CustomDomainPropertiesArgs) ToCustomDomainPropertiesOutputWithContext(ctx context.Context) CustomDomainPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainPropertiesOutput)
}

func (i CustomDomainPropertiesArgs) ToCustomDomainPropertiesPtrOutput() CustomDomainPropertiesPtrOutput {
	return i.ToCustomDomainPropertiesPtrOutputWithContext(context.Background())
}

func (i CustomDomainPropertiesArgs) ToCustomDomainPropertiesPtrOutputWithContext(ctx context.Context) CustomDomainPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainPropertiesOutput).ToCustomDomainPropertiesPtrOutputWithContext(ctx)
}









type CustomDomainPropertiesPtrInput interface {
	pulumi.Input

	ToCustomDomainPropertiesPtrOutput() CustomDomainPropertiesPtrOutput
	ToCustomDomainPropertiesPtrOutputWithContext(context.Context) CustomDomainPropertiesPtrOutput
}

type customDomainPropertiesPtrType CustomDomainPropertiesArgs

func CustomDomainPropertiesPtr(v *CustomDomainPropertiesArgs) CustomDomainPropertiesPtrInput {
	return (*customDomainPropertiesPtrType)(v)
}

func (*customDomainPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDomainProperties)(nil)).Elem()
}

func (i *customDomainPropertiesPtrType) ToCustomDomainPropertiesPtrOutput() CustomDomainPropertiesPtrOutput {
	return i.ToCustomDomainPropertiesPtrOutputWithContext(context.Background())
}

func (i *customDomainPropertiesPtrType) ToCustomDomainPropertiesPtrOutputWithContext(ctx context.Context) CustomDomainPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainPropertiesPtrOutput)
}

type CustomDomainPropertiesOutput struct{ *pulumi.OutputState }

func (CustomDomainPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomDomainProperties)(nil)).Elem()
}

func (o CustomDomainPropertiesOutput) ToCustomDomainPropertiesOutput() CustomDomainPropertiesOutput {
	return o
}

func (o CustomDomainPropertiesOutput) ToCustomDomainPropertiesOutputWithContext(ctx context.Context) CustomDomainPropertiesOutput {
	return o
}

func (o CustomDomainPropertiesOutput) ToCustomDomainPropertiesPtrOutput() CustomDomainPropertiesPtrOutput {
	return o.ToCustomDomainPropertiesPtrOutputWithContext(context.Background())
}

func (o CustomDomainPropertiesOutput) ToCustomDomainPropertiesPtrOutputWithContext(ctx context.Context) CustomDomainPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomDomainProperties) *CustomDomainProperties {
		return &v
	}).(CustomDomainPropertiesPtrOutput)
}

func (o CustomDomainPropertiesOutput) CertName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomDomainProperties) *string { return v.CertName }).(pulumi.StringPtrOutput)
}

func (o CustomDomainPropertiesOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomDomainProperties) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type CustomDomainPropertiesPtrOutput struct{ *pulumi.OutputState }

func (CustomDomainPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDomainProperties)(nil)).Elem()
}

func (o CustomDomainPropertiesPtrOutput) ToCustomDomainPropertiesPtrOutput() CustomDomainPropertiesPtrOutput {
	return o
}

func (o CustomDomainPropertiesPtrOutput) ToCustomDomainPropertiesPtrOutputWithContext(ctx context.Context) CustomDomainPropertiesPtrOutput {
	return o
}

func (o CustomDomainPropertiesPtrOutput) Elem() CustomDomainPropertiesOutput {
	return o.ApplyT(func(v *CustomDomainProperties) CustomDomainProperties {
		if v != nil {
			return *v
		}
		var ret CustomDomainProperties
		return ret
	}).(CustomDomainPropertiesOutput)
}

func (o CustomDomainPropertiesPtrOutput) CertName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomDomainProperties) *string {
		if v == nil {
			return nil
		}
		return v.CertName
	}).(pulumi.StringPtrOutput)
}

func (o CustomDomainPropertiesPtrOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomDomainProperties) *string {
		if v == nil {
			return nil
		}
		return v.Thumbprint
	}).(pulumi.StringPtrOutput)
}

type CustomDomainPropertiesResponse struct {
	AppName    string  `pulumi:"appName"`
	CertName   *string `pulumi:"certName"`
	Thumbprint *string `pulumi:"thumbprint"`
}





type CustomDomainPropertiesResponseInput interface {
	pulumi.Input

	ToCustomDomainPropertiesResponseOutput() CustomDomainPropertiesResponseOutput
	ToCustomDomainPropertiesResponseOutputWithContext(context.Context) CustomDomainPropertiesResponseOutput
}

type CustomDomainPropertiesResponseArgs struct {
	AppName    pulumi.StringInput    `pulumi:"appName"`
	CertName   pulumi.StringPtrInput `pulumi:"certName"`
	Thumbprint pulumi.StringPtrInput `pulumi:"thumbprint"`
}

func (CustomDomainPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomDomainPropertiesResponse)(nil)).Elem()
}

func (i CustomDomainPropertiesResponseArgs) ToCustomDomainPropertiesResponseOutput() CustomDomainPropertiesResponseOutput {
	return i.ToCustomDomainPropertiesResponseOutputWithContext(context.Background())
}

func (i CustomDomainPropertiesResponseArgs) ToCustomDomainPropertiesResponseOutputWithContext(ctx context.Context) CustomDomainPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainPropertiesResponseOutput)
}

func (i CustomDomainPropertiesResponseArgs) ToCustomDomainPropertiesResponsePtrOutput() CustomDomainPropertiesResponsePtrOutput {
	return i.ToCustomDomainPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i CustomDomainPropertiesResponseArgs) ToCustomDomainPropertiesResponsePtrOutputWithContext(ctx context.Context) CustomDomainPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainPropertiesResponseOutput).ToCustomDomainPropertiesResponsePtrOutputWithContext(ctx)
}









type CustomDomainPropertiesResponsePtrInput interface {
	pulumi.Input

	ToCustomDomainPropertiesResponsePtrOutput() CustomDomainPropertiesResponsePtrOutput
	ToCustomDomainPropertiesResponsePtrOutputWithContext(context.Context) CustomDomainPropertiesResponsePtrOutput
}

type customDomainPropertiesResponsePtrType CustomDomainPropertiesResponseArgs

func CustomDomainPropertiesResponsePtr(v *CustomDomainPropertiesResponseArgs) CustomDomainPropertiesResponsePtrInput {
	return (*customDomainPropertiesResponsePtrType)(v)
}

func (*customDomainPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDomainPropertiesResponse)(nil)).Elem()
}

func (i *customDomainPropertiesResponsePtrType) ToCustomDomainPropertiesResponsePtrOutput() CustomDomainPropertiesResponsePtrOutput {
	return i.ToCustomDomainPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *customDomainPropertiesResponsePtrType) ToCustomDomainPropertiesResponsePtrOutputWithContext(ctx context.Context) CustomDomainPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainPropertiesResponsePtrOutput)
}

type CustomDomainPropertiesResponseOutput struct{ *pulumi.OutputState }

func (CustomDomainPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomDomainPropertiesResponse)(nil)).Elem()
}

func (o CustomDomainPropertiesResponseOutput) ToCustomDomainPropertiesResponseOutput() CustomDomainPropertiesResponseOutput {
	return o
}

func (o CustomDomainPropertiesResponseOutput) ToCustomDomainPropertiesResponseOutputWithContext(ctx context.Context) CustomDomainPropertiesResponseOutput {
	return o
}

func (o CustomDomainPropertiesResponseOutput) ToCustomDomainPropertiesResponsePtrOutput() CustomDomainPropertiesResponsePtrOutput {
	return o.ToCustomDomainPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o CustomDomainPropertiesResponseOutput) ToCustomDomainPropertiesResponsePtrOutputWithContext(ctx context.Context) CustomDomainPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomDomainPropertiesResponse) *CustomDomainPropertiesResponse {
		return &v
	}).(CustomDomainPropertiesResponsePtrOutput)
}

func (o CustomDomainPropertiesResponseOutput) AppName() pulumi.StringOutput {
	return o.ApplyT(func(v CustomDomainPropertiesResponse) string { return v.AppName }).(pulumi.StringOutput)
}

func (o CustomDomainPropertiesResponseOutput) CertName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomDomainPropertiesResponse) *string { return v.CertName }).(pulumi.StringPtrOutput)
}

func (o CustomDomainPropertiesResponseOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomDomainPropertiesResponse) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type CustomDomainPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (CustomDomainPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDomainPropertiesResponse)(nil)).Elem()
}

func (o CustomDomainPropertiesResponsePtrOutput) ToCustomDomainPropertiesResponsePtrOutput() CustomDomainPropertiesResponsePtrOutput {
	return o
}

func (o CustomDomainPropertiesResponsePtrOutput) ToCustomDomainPropertiesResponsePtrOutputWithContext(ctx context.Context) CustomDomainPropertiesResponsePtrOutput {
	return o
}

func (o CustomDomainPropertiesResponsePtrOutput) Elem() CustomDomainPropertiesResponseOutput {
	return o.ApplyT(func(v *CustomDomainPropertiesResponse) CustomDomainPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret CustomDomainPropertiesResponse
		return ret
	}).(CustomDomainPropertiesResponseOutput)
}

func (o CustomDomainPropertiesResponsePtrOutput) AppName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomDomainPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AppName
	}).(pulumi.StringPtrOutput)
}

func (o CustomDomainPropertiesResponsePtrOutput) CertName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomDomainPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.CertName
	}).(pulumi.StringPtrOutput)
}

func (o CustomDomainPropertiesResponsePtrOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomDomainPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Thumbprint
	}).(pulumi.StringPtrOutput)
}

type DeploymentInstanceResponse struct {
	DiscoveryStatus string `pulumi:"discoveryStatus"`
	Name            string `pulumi:"name"`
	Reason          string `pulumi:"reason"`
	StartTime       string `pulumi:"startTime"`
	Status          string `pulumi:"status"`
}





type DeploymentInstanceResponseInput interface {
	pulumi.Input

	ToDeploymentInstanceResponseOutput() DeploymentInstanceResponseOutput
	ToDeploymentInstanceResponseOutputWithContext(context.Context) DeploymentInstanceResponseOutput
}

type DeploymentInstanceResponseArgs struct {
	DiscoveryStatus pulumi.StringInput `pulumi:"discoveryStatus"`
	Name            pulumi.StringInput `pulumi:"name"`
	Reason          pulumi.StringInput `pulumi:"reason"`
	StartTime       pulumi.StringInput `pulumi:"startTime"`
	Status          pulumi.StringInput `pulumi:"status"`
}

func (DeploymentInstanceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentInstanceResponse)(nil)).Elem()
}

func (i DeploymentInstanceResponseArgs) ToDeploymentInstanceResponseOutput() DeploymentInstanceResponseOutput {
	return i.ToDeploymentInstanceResponseOutputWithContext(context.Background())
}

func (i DeploymentInstanceResponseArgs) ToDeploymentInstanceResponseOutputWithContext(ctx context.Context) DeploymentInstanceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentInstanceResponseOutput)
}





type DeploymentInstanceResponseArrayInput interface {
	pulumi.Input

	ToDeploymentInstanceResponseArrayOutput() DeploymentInstanceResponseArrayOutput
	ToDeploymentInstanceResponseArrayOutputWithContext(context.Context) DeploymentInstanceResponseArrayOutput
}

type DeploymentInstanceResponseArray []DeploymentInstanceResponseInput

func (DeploymentInstanceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DeploymentInstanceResponse)(nil)).Elem()
}

func (i DeploymentInstanceResponseArray) ToDeploymentInstanceResponseArrayOutput() DeploymentInstanceResponseArrayOutput {
	return i.ToDeploymentInstanceResponseArrayOutputWithContext(context.Background())
}

func (i DeploymentInstanceResponseArray) ToDeploymentInstanceResponseArrayOutputWithContext(ctx context.Context) DeploymentInstanceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentInstanceResponseArrayOutput)
}

type DeploymentInstanceResponseOutput struct{ *pulumi.OutputState }

func (DeploymentInstanceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentInstanceResponse)(nil)).Elem()
}

func (o DeploymentInstanceResponseOutput) ToDeploymentInstanceResponseOutput() DeploymentInstanceResponseOutput {
	return o
}

func (o DeploymentInstanceResponseOutput) ToDeploymentInstanceResponseOutputWithContext(ctx context.Context) DeploymentInstanceResponseOutput {
	return o
}

func (o DeploymentInstanceResponseOutput) DiscoveryStatus() pulumi.StringOutput {
	return o.ApplyT(func(v DeploymentInstanceResponse) string { return v.DiscoveryStatus }).(pulumi.StringOutput)
}

func (o DeploymentInstanceResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DeploymentInstanceResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o DeploymentInstanceResponseOutput) Reason() pulumi.StringOutput {
	return o.ApplyT(func(v DeploymentInstanceResponse) string { return v.Reason }).(pulumi.StringOutput)
}

func (o DeploymentInstanceResponseOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v DeploymentInstanceResponse) string { return v.StartTime }).(pulumi.StringOutput)
}

func (o DeploymentInstanceResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v DeploymentInstanceResponse) string { return v.Status }).(pulumi.StringOutput)
}

type DeploymentInstanceResponseArrayOutput struct{ *pulumi.OutputState }

func (DeploymentInstanceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DeploymentInstanceResponse)(nil)).Elem()
}

func (o DeploymentInstanceResponseArrayOutput) ToDeploymentInstanceResponseArrayOutput() DeploymentInstanceResponseArrayOutput {
	return o
}

func (o DeploymentInstanceResponseArrayOutput) ToDeploymentInstanceResponseArrayOutputWithContext(ctx context.Context) DeploymentInstanceResponseArrayOutput {
	return o
}

func (o DeploymentInstanceResponseArrayOutput) Index(i pulumi.IntInput) DeploymentInstanceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DeploymentInstanceResponse {
		return vs[0].([]DeploymentInstanceResponse)[vs[1].(int)]
	}).(DeploymentInstanceResponseOutput)
}

type DeploymentResourceProperties struct {
	DeploymentSettings *DeploymentSettings `pulumi:"deploymentSettings"`
	Source             *UserSourceInfo     `pulumi:"source"`
}


func (val *DeploymentResourceProperties) Defaults() *DeploymentResourceProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.DeploymentSettings = tmp.DeploymentSettings.Defaults()

	return &tmp
}





type DeploymentResourcePropertiesInput interface {
	pulumi.Input

	ToDeploymentResourcePropertiesOutput() DeploymentResourcePropertiesOutput
	ToDeploymentResourcePropertiesOutputWithContext(context.Context) DeploymentResourcePropertiesOutput
}

type DeploymentResourcePropertiesArgs struct {
	DeploymentSettings DeploymentSettingsPtrInput `pulumi:"deploymentSettings"`
	Source             UserSourceInfoPtrInput     `pulumi:"source"`
}

func (DeploymentResourcePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentResourceProperties)(nil)).Elem()
}

func (i DeploymentResourcePropertiesArgs) ToDeploymentResourcePropertiesOutput() DeploymentResourcePropertiesOutput {
	return i.ToDeploymentResourcePropertiesOutputWithContext(context.Background())
}

func (i DeploymentResourcePropertiesArgs) ToDeploymentResourcePropertiesOutputWithContext(ctx context.Context) DeploymentResourcePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentResourcePropertiesOutput)
}

func (i DeploymentResourcePropertiesArgs) ToDeploymentResourcePropertiesPtrOutput() DeploymentResourcePropertiesPtrOutput {
	return i.ToDeploymentResourcePropertiesPtrOutputWithContext(context.Background())
}

func (i DeploymentResourcePropertiesArgs) ToDeploymentResourcePropertiesPtrOutputWithContext(ctx context.Context) DeploymentResourcePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentResourcePropertiesOutput).ToDeploymentResourcePropertiesPtrOutputWithContext(ctx)
}









type DeploymentResourcePropertiesPtrInput interface {
	pulumi.Input

	ToDeploymentResourcePropertiesPtrOutput() DeploymentResourcePropertiesPtrOutput
	ToDeploymentResourcePropertiesPtrOutputWithContext(context.Context) DeploymentResourcePropertiesPtrOutput
}

type deploymentResourcePropertiesPtrType DeploymentResourcePropertiesArgs

func DeploymentResourcePropertiesPtr(v *DeploymentResourcePropertiesArgs) DeploymentResourcePropertiesPtrInput {
	return (*deploymentResourcePropertiesPtrType)(v)
}

func (*deploymentResourcePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentResourceProperties)(nil)).Elem()
}

func (i *deploymentResourcePropertiesPtrType) ToDeploymentResourcePropertiesPtrOutput() DeploymentResourcePropertiesPtrOutput {
	return i.ToDeploymentResourcePropertiesPtrOutputWithContext(context.Background())
}

func (i *deploymentResourcePropertiesPtrType) ToDeploymentResourcePropertiesPtrOutputWithContext(ctx context.Context) DeploymentResourcePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentResourcePropertiesPtrOutput)
}

type DeploymentResourcePropertiesOutput struct{ *pulumi.OutputState }

func (DeploymentResourcePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentResourceProperties)(nil)).Elem()
}

func (o DeploymentResourcePropertiesOutput) ToDeploymentResourcePropertiesOutput() DeploymentResourcePropertiesOutput {
	return o
}

func (o DeploymentResourcePropertiesOutput) ToDeploymentResourcePropertiesOutputWithContext(ctx context.Context) DeploymentResourcePropertiesOutput {
	return o
}

func (o DeploymentResourcePropertiesOutput) ToDeploymentResourcePropertiesPtrOutput() DeploymentResourcePropertiesPtrOutput {
	return o.ToDeploymentResourcePropertiesPtrOutputWithContext(context.Background())
}

func (o DeploymentResourcePropertiesOutput) ToDeploymentResourcePropertiesPtrOutputWithContext(ctx context.Context) DeploymentResourcePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeploymentResourceProperties) *DeploymentResourceProperties {
		return &v
	}).(DeploymentResourcePropertiesPtrOutput)
}

func (o DeploymentResourcePropertiesOutput) DeploymentSettings() DeploymentSettingsPtrOutput {
	return o.ApplyT(func(v DeploymentResourceProperties) *DeploymentSettings { return v.DeploymentSettings }).(DeploymentSettingsPtrOutput)
}

func (o DeploymentResourcePropertiesOutput) Source() UserSourceInfoPtrOutput {
	return o.ApplyT(func(v DeploymentResourceProperties) *UserSourceInfo { return v.Source }).(UserSourceInfoPtrOutput)
}

type DeploymentResourcePropertiesPtrOutput struct{ *pulumi.OutputState }

func (DeploymentResourcePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentResourceProperties)(nil)).Elem()
}

func (o DeploymentResourcePropertiesPtrOutput) ToDeploymentResourcePropertiesPtrOutput() DeploymentResourcePropertiesPtrOutput {
	return o
}

func (o DeploymentResourcePropertiesPtrOutput) ToDeploymentResourcePropertiesPtrOutputWithContext(ctx context.Context) DeploymentResourcePropertiesPtrOutput {
	return o
}

func (o DeploymentResourcePropertiesPtrOutput) Elem() DeploymentResourcePropertiesOutput {
	return o.ApplyT(func(v *DeploymentResourceProperties) DeploymentResourceProperties {
		if v != nil {
			return *v
		}
		var ret DeploymentResourceProperties
		return ret
	}).(DeploymentResourcePropertiesOutput)
}

func (o DeploymentResourcePropertiesPtrOutput) DeploymentSettings() DeploymentSettingsPtrOutput {
	return o.ApplyT(func(v *DeploymentResourceProperties) *DeploymentSettings {
		if v == nil {
			return nil
		}
		return v.DeploymentSettings
	}).(DeploymentSettingsPtrOutput)
}

func (o DeploymentResourcePropertiesPtrOutput) Source() UserSourceInfoPtrOutput {
	return o.ApplyT(func(v *DeploymentResourceProperties) *UserSourceInfo {
		if v == nil {
			return nil
		}
		return v.Source
	}).(UserSourceInfoPtrOutput)
}

type DeploymentResourcePropertiesResponse struct {
	Active             bool                         `pulumi:"active"`
	AppName            string                       `pulumi:"appName"`
	CreatedTime        string                       `pulumi:"createdTime"`
	DeploymentSettings *DeploymentSettingsResponse  `pulumi:"deploymentSettings"`
	Instances          []DeploymentInstanceResponse `pulumi:"instances"`
	ProvisioningState  string                       `pulumi:"provisioningState"`
	Source             *UserSourceInfoResponse      `pulumi:"source"`
	Status             string                       `pulumi:"status"`
}


func (val *DeploymentResourcePropertiesResponse) Defaults() *DeploymentResourcePropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.DeploymentSettings = tmp.DeploymentSettings.Defaults()

	return &tmp
}





type DeploymentResourcePropertiesResponseInput interface {
	pulumi.Input

	ToDeploymentResourcePropertiesResponseOutput() DeploymentResourcePropertiesResponseOutput
	ToDeploymentResourcePropertiesResponseOutputWithContext(context.Context) DeploymentResourcePropertiesResponseOutput
}

type DeploymentResourcePropertiesResponseArgs struct {
	Active             pulumi.BoolInput                     `pulumi:"active"`
	AppName            pulumi.StringInput                   `pulumi:"appName"`
	CreatedTime        pulumi.StringInput                   `pulumi:"createdTime"`
	DeploymentSettings DeploymentSettingsResponsePtrInput   `pulumi:"deploymentSettings"`
	Instances          DeploymentInstanceResponseArrayInput `pulumi:"instances"`
	ProvisioningState  pulumi.StringInput                   `pulumi:"provisioningState"`
	Source             UserSourceInfoResponsePtrInput       `pulumi:"source"`
	Status             pulumi.StringInput                   `pulumi:"status"`
}

func (DeploymentResourcePropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentResourcePropertiesResponse)(nil)).Elem()
}

func (i DeploymentResourcePropertiesResponseArgs) ToDeploymentResourcePropertiesResponseOutput() DeploymentResourcePropertiesResponseOutput {
	return i.ToDeploymentResourcePropertiesResponseOutputWithContext(context.Background())
}

func (i DeploymentResourcePropertiesResponseArgs) ToDeploymentResourcePropertiesResponseOutputWithContext(ctx context.Context) DeploymentResourcePropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentResourcePropertiesResponseOutput)
}

func (i DeploymentResourcePropertiesResponseArgs) ToDeploymentResourcePropertiesResponsePtrOutput() DeploymentResourcePropertiesResponsePtrOutput {
	return i.ToDeploymentResourcePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i DeploymentResourcePropertiesResponseArgs) ToDeploymentResourcePropertiesResponsePtrOutputWithContext(ctx context.Context) DeploymentResourcePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentResourcePropertiesResponseOutput).ToDeploymentResourcePropertiesResponsePtrOutputWithContext(ctx)
}









type DeploymentResourcePropertiesResponsePtrInput interface {
	pulumi.Input

	ToDeploymentResourcePropertiesResponsePtrOutput() DeploymentResourcePropertiesResponsePtrOutput
	ToDeploymentResourcePropertiesResponsePtrOutputWithContext(context.Context) DeploymentResourcePropertiesResponsePtrOutput
}

type deploymentResourcePropertiesResponsePtrType DeploymentResourcePropertiesResponseArgs

func DeploymentResourcePropertiesResponsePtr(v *DeploymentResourcePropertiesResponseArgs) DeploymentResourcePropertiesResponsePtrInput {
	return (*deploymentResourcePropertiesResponsePtrType)(v)
}

func (*deploymentResourcePropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentResourcePropertiesResponse)(nil)).Elem()
}

func (i *deploymentResourcePropertiesResponsePtrType) ToDeploymentResourcePropertiesResponsePtrOutput() DeploymentResourcePropertiesResponsePtrOutput {
	return i.ToDeploymentResourcePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *deploymentResourcePropertiesResponsePtrType) ToDeploymentResourcePropertiesResponsePtrOutputWithContext(ctx context.Context) DeploymentResourcePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentResourcePropertiesResponsePtrOutput)
}

type DeploymentResourcePropertiesResponseOutput struct{ *pulumi.OutputState }

func (DeploymentResourcePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentResourcePropertiesResponse)(nil)).Elem()
}

func (o DeploymentResourcePropertiesResponseOutput) ToDeploymentResourcePropertiesResponseOutput() DeploymentResourcePropertiesResponseOutput {
	return o
}

func (o DeploymentResourcePropertiesResponseOutput) ToDeploymentResourcePropertiesResponseOutputWithContext(ctx context.Context) DeploymentResourcePropertiesResponseOutput {
	return o
}

func (o DeploymentResourcePropertiesResponseOutput) ToDeploymentResourcePropertiesResponsePtrOutput() DeploymentResourcePropertiesResponsePtrOutput {
	return o.ToDeploymentResourcePropertiesResponsePtrOutputWithContext(context.Background())
}

func (o DeploymentResourcePropertiesResponseOutput) ToDeploymentResourcePropertiesResponsePtrOutputWithContext(ctx context.Context) DeploymentResourcePropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeploymentResourcePropertiesResponse) *DeploymentResourcePropertiesResponse {
		return &v
	}).(DeploymentResourcePropertiesResponsePtrOutput)
}

func (o DeploymentResourcePropertiesResponseOutput) Active() pulumi.BoolOutput {
	return o.ApplyT(func(v DeploymentResourcePropertiesResponse) bool { return v.Active }).(pulumi.BoolOutput)
}

func (o DeploymentResourcePropertiesResponseOutput) AppName() pulumi.StringOutput {
	return o.ApplyT(func(v DeploymentResourcePropertiesResponse) string { return v.AppName }).(pulumi.StringOutput)
}

func (o DeploymentResourcePropertiesResponseOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v DeploymentResourcePropertiesResponse) string { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o DeploymentResourcePropertiesResponseOutput) DeploymentSettings() DeploymentSettingsResponsePtrOutput {
	return o.ApplyT(func(v DeploymentResourcePropertiesResponse) *DeploymentSettingsResponse { return v.DeploymentSettings }).(DeploymentSettingsResponsePtrOutput)
}

func (o DeploymentResourcePropertiesResponseOutput) Instances() DeploymentInstanceResponseArrayOutput {
	return o.ApplyT(func(v DeploymentResourcePropertiesResponse) []DeploymentInstanceResponse { return v.Instances }).(DeploymentInstanceResponseArrayOutput)
}

func (o DeploymentResourcePropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v DeploymentResourcePropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DeploymentResourcePropertiesResponseOutput) Source() UserSourceInfoResponsePtrOutput {
	return o.ApplyT(func(v DeploymentResourcePropertiesResponse) *UserSourceInfoResponse { return v.Source }).(UserSourceInfoResponsePtrOutput)
}

func (o DeploymentResourcePropertiesResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v DeploymentResourcePropertiesResponse) string { return v.Status }).(pulumi.StringOutput)
}

type DeploymentResourcePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (DeploymentResourcePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentResourcePropertiesResponse)(nil)).Elem()
}

func (o DeploymentResourcePropertiesResponsePtrOutput) ToDeploymentResourcePropertiesResponsePtrOutput() DeploymentResourcePropertiesResponsePtrOutput {
	return o
}

func (o DeploymentResourcePropertiesResponsePtrOutput) ToDeploymentResourcePropertiesResponsePtrOutputWithContext(ctx context.Context) DeploymentResourcePropertiesResponsePtrOutput {
	return o
}

func (o DeploymentResourcePropertiesResponsePtrOutput) Elem() DeploymentResourcePropertiesResponseOutput {
	return o.ApplyT(func(v *DeploymentResourcePropertiesResponse) DeploymentResourcePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret DeploymentResourcePropertiesResponse
		return ret
	}).(DeploymentResourcePropertiesResponseOutput)
}

func (o DeploymentResourcePropertiesResponsePtrOutput) Active() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DeploymentResourcePropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.Active
	}).(pulumi.BoolPtrOutput)
}

func (o DeploymentResourcePropertiesResponsePtrOutput) AppName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentResourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AppName
	}).(pulumi.StringPtrOutput)
}

func (o DeploymentResourcePropertiesResponsePtrOutput) CreatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentResourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CreatedTime
	}).(pulumi.StringPtrOutput)
}

func (o DeploymentResourcePropertiesResponsePtrOutput) DeploymentSettings() DeploymentSettingsResponsePtrOutput {
	return o.ApplyT(func(v *DeploymentResourcePropertiesResponse) *DeploymentSettingsResponse {
		if v == nil {
			return nil
		}
		return v.DeploymentSettings
	}).(DeploymentSettingsResponsePtrOutput)
}

func (o DeploymentResourcePropertiesResponsePtrOutput) Instances() DeploymentInstanceResponseArrayOutput {
	return o.ApplyT(func(v *DeploymentResourcePropertiesResponse) []DeploymentInstanceResponse {
		if v == nil {
			return nil
		}
		return v.Instances
	}).(DeploymentInstanceResponseArrayOutput)
}

func (o DeploymentResourcePropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentResourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o DeploymentResourcePropertiesResponsePtrOutput) Source() UserSourceInfoResponsePtrOutput {
	return o.ApplyT(func(v *DeploymentResourcePropertiesResponse) *UserSourceInfoResponse {
		if v == nil {
			return nil
		}
		return v.Source
	}).(UserSourceInfoResponsePtrOutput)
}

func (o DeploymentResourcePropertiesResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentResourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

type DeploymentSettings struct {
	Cpu                  *int              `pulumi:"cpu"`
	EnvironmentVariables map[string]string `pulumi:"environmentVariables"`
	JvmOptions           *string           `pulumi:"jvmOptions"`
	MemoryInGB           *int              `pulumi:"memoryInGB"`
	NetCoreMainEntryPath *string           `pulumi:"netCoreMainEntryPath"`
	ResourceRequests     *ResourceRequests `pulumi:"resourceRequests"`
	RuntimeVersion       *string           `pulumi:"runtimeVersion"`
}


func (val *DeploymentSettings) Defaults() *DeploymentSettings {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Cpu) {
		cpu_ := 1
		tmp.Cpu = &cpu_
	}
	if isZero(tmp.MemoryInGB) {
		memoryInGB_ := 1
		tmp.MemoryInGB = &memoryInGB_
	}
	if isZero(tmp.RuntimeVersion) {
		runtimeVersion_ := "Java_8"
		tmp.RuntimeVersion = &runtimeVersion_
	}
	return &tmp
}





type DeploymentSettingsInput interface {
	pulumi.Input

	ToDeploymentSettingsOutput() DeploymentSettingsOutput
	ToDeploymentSettingsOutputWithContext(context.Context) DeploymentSettingsOutput
}

type DeploymentSettingsArgs struct {
	Cpu                  pulumi.IntPtrInput       `pulumi:"cpu"`
	EnvironmentVariables pulumi.StringMapInput    `pulumi:"environmentVariables"`
	JvmOptions           pulumi.StringPtrInput    `pulumi:"jvmOptions"`
	MemoryInGB           pulumi.IntPtrInput       `pulumi:"memoryInGB"`
	NetCoreMainEntryPath pulumi.StringPtrInput    `pulumi:"netCoreMainEntryPath"`
	ResourceRequests     ResourceRequestsPtrInput `pulumi:"resourceRequests"`
	RuntimeVersion       pulumi.StringPtrInput    `pulumi:"runtimeVersion"`
}

func (DeploymentSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentSettings)(nil)).Elem()
}

func (i DeploymentSettingsArgs) ToDeploymentSettingsOutput() DeploymentSettingsOutput {
	return i.ToDeploymentSettingsOutputWithContext(context.Background())
}

func (i DeploymentSettingsArgs) ToDeploymentSettingsOutputWithContext(ctx context.Context) DeploymentSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentSettingsOutput)
}

func (i DeploymentSettingsArgs) ToDeploymentSettingsPtrOutput() DeploymentSettingsPtrOutput {
	return i.ToDeploymentSettingsPtrOutputWithContext(context.Background())
}

func (i DeploymentSettingsArgs) ToDeploymentSettingsPtrOutputWithContext(ctx context.Context) DeploymentSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentSettingsOutput).ToDeploymentSettingsPtrOutputWithContext(ctx)
}









type DeploymentSettingsPtrInput interface {
	pulumi.Input

	ToDeploymentSettingsPtrOutput() DeploymentSettingsPtrOutput
	ToDeploymentSettingsPtrOutputWithContext(context.Context) DeploymentSettingsPtrOutput
}

type deploymentSettingsPtrType DeploymentSettingsArgs

func DeploymentSettingsPtr(v *DeploymentSettingsArgs) DeploymentSettingsPtrInput {
	return (*deploymentSettingsPtrType)(v)
}

func (*deploymentSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentSettings)(nil)).Elem()
}

func (i *deploymentSettingsPtrType) ToDeploymentSettingsPtrOutput() DeploymentSettingsPtrOutput {
	return i.ToDeploymentSettingsPtrOutputWithContext(context.Background())
}

func (i *deploymentSettingsPtrType) ToDeploymentSettingsPtrOutputWithContext(ctx context.Context) DeploymentSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentSettingsPtrOutput)
}

type DeploymentSettingsOutput struct{ *pulumi.OutputState }

func (DeploymentSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentSettings)(nil)).Elem()
}

func (o DeploymentSettingsOutput) ToDeploymentSettingsOutput() DeploymentSettingsOutput {
	return o
}

func (o DeploymentSettingsOutput) ToDeploymentSettingsOutputWithContext(ctx context.Context) DeploymentSettingsOutput {
	return o
}

func (o DeploymentSettingsOutput) ToDeploymentSettingsPtrOutput() DeploymentSettingsPtrOutput {
	return o.ToDeploymentSettingsPtrOutputWithContext(context.Background())
}

func (o DeploymentSettingsOutput) ToDeploymentSettingsPtrOutputWithContext(ctx context.Context) DeploymentSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeploymentSettings) *DeploymentSettings {
		return &v
	}).(DeploymentSettingsPtrOutput)
}

func (o DeploymentSettingsOutput) Cpu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DeploymentSettings) *int { return v.Cpu }).(pulumi.IntPtrOutput)
}

func (o DeploymentSettingsOutput) EnvironmentVariables() pulumi.StringMapOutput {
	return o.ApplyT(func(v DeploymentSettings) map[string]string { return v.EnvironmentVariables }).(pulumi.StringMapOutput)
}

func (o DeploymentSettingsOutput) JvmOptions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentSettings) *string { return v.JvmOptions }).(pulumi.StringPtrOutput)
}

func (o DeploymentSettingsOutput) MemoryInGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DeploymentSettings) *int { return v.MemoryInGB }).(pulumi.IntPtrOutput)
}

func (o DeploymentSettingsOutput) NetCoreMainEntryPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentSettings) *string { return v.NetCoreMainEntryPath }).(pulumi.StringPtrOutput)
}

func (o DeploymentSettingsOutput) ResourceRequests() ResourceRequestsPtrOutput {
	return o.ApplyT(func(v DeploymentSettings) *ResourceRequests { return v.ResourceRequests }).(ResourceRequestsPtrOutput)
}

func (o DeploymentSettingsOutput) RuntimeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentSettings) *string { return v.RuntimeVersion }).(pulumi.StringPtrOutput)
}

type DeploymentSettingsPtrOutput struct{ *pulumi.OutputState }

func (DeploymentSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentSettings)(nil)).Elem()
}

func (o DeploymentSettingsPtrOutput) ToDeploymentSettingsPtrOutput() DeploymentSettingsPtrOutput {
	return o
}

func (o DeploymentSettingsPtrOutput) ToDeploymentSettingsPtrOutputWithContext(ctx context.Context) DeploymentSettingsPtrOutput {
	return o
}

func (o DeploymentSettingsPtrOutput) Elem() DeploymentSettingsOutput {
	return o.ApplyT(func(v *DeploymentSettings) DeploymentSettings {
		if v != nil {
			return *v
		}
		var ret DeploymentSettings
		return ret
	}).(DeploymentSettingsOutput)
}

func (o DeploymentSettingsPtrOutput) Cpu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DeploymentSettings) *int {
		if v == nil {
			return nil
		}
		return v.Cpu
	}).(pulumi.IntPtrOutput)
}

func (o DeploymentSettingsPtrOutput) EnvironmentVariables() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DeploymentSettings) map[string]string {
		if v == nil {
			return nil
		}
		return v.EnvironmentVariables
	}).(pulumi.StringMapOutput)
}

func (o DeploymentSettingsPtrOutput) JvmOptions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentSettings) *string {
		if v == nil {
			return nil
		}
		return v.JvmOptions
	}).(pulumi.StringPtrOutput)
}

func (o DeploymentSettingsPtrOutput) MemoryInGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DeploymentSettings) *int {
		if v == nil {
			return nil
		}
		return v.MemoryInGB
	}).(pulumi.IntPtrOutput)
}

func (o DeploymentSettingsPtrOutput) NetCoreMainEntryPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentSettings) *string {
		if v == nil {
			return nil
		}
		return v.NetCoreMainEntryPath
	}).(pulumi.StringPtrOutput)
}

func (o DeploymentSettingsPtrOutput) ResourceRequests() ResourceRequestsPtrOutput {
	return o.ApplyT(func(v *DeploymentSettings) *ResourceRequests {
		if v == nil {
			return nil
		}
		return v.ResourceRequests
	}).(ResourceRequestsPtrOutput)
}

func (o DeploymentSettingsPtrOutput) RuntimeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentSettings) *string {
		if v == nil {
			return nil
		}
		return v.RuntimeVersion
	}).(pulumi.StringPtrOutput)
}

type DeploymentSettingsResponse struct {
	Cpu                  *int                      `pulumi:"cpu"`
	EnvironmentVariables map[string]string         `pulumi:"environmentVariables"`
	JvmOptions           *string                   `pulumi:"jvmOptions"`
	MemoryInGB           *int                      `pulumi:"memoryInGB"`
	NetCoreMainEntryPath *string                   `pulumi:"netCoreMainEntryPath"`
	ResourceRequests     *ResourceRequestsResponse `pulumi:"resourceRequests"`
	RuntimeVersion       *string                   `pulumi:"runtimeVersion"`
}


func (val *DeploymentSettingsResponse) Defaults() *DeploymentSettingsResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Cpu) {
		cpu_ := 1
		tmp.Cpu = &cpu_
	}
	if isZero(tmp.MemoryInGB) {
		memoryInGB_ := 1
		tmp.MemoryInGB = &memoryInGB_
	}
	if isZero(tmp.RuntimeVersion) {
		runtimeVersion_ := "Java_8"
		tmp.RuntimeVersion = &runtimeVersion_
	}
	return &tmp
}





type DeploymentSettingsResponseInput interface {
	pulumi.Input

	ToDeploymentSettingsResponseOutput() DeploymentSettingsResponseOutput
	ToDeploymentSettingsResponseOutputWithContext(context.Context) DeploymentSettingsResponseOutput
}

type DeploymentSettingsResponseArgs struct {
	Cpu                  pulumi.IntPtrInput               `pulumi:"cpu"`
	EnvironmentVariables pulumi.StringMapInput            `pulumi:"environmentVariables"`
	JvmOptions           pulumi.StringPtrInput            `pulumi:"jvmOptions"`
	MemoryInGB           pulumi.IntPtrInput               `pulumi:"memoryInGB"`
	NetCoreMainEntryPath pulumi.StringPtrInput            `pulumi:"netCoreMainEntryPath"`
	ResourceRequests     ResourceRequestsResponsePtrInput `pulumi:"resourceRequests"`
	RuntimeVersion       pulumi.StringPtrInput            `pulumi:"runtimeVersion"`
}

func (DeploymentSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentSettingsResponse)(nil)).Elem()
}

func (i DeploymentSettingsResponseArgs) ToDeploymentSettingsResponseOutput() DeploymentSettingsResponseOutput {
	return i.ToDeploymentSettingsResponseOutputWithContext(context.Background())
}

func (i DeploymentSettingsResponseArgs) ToDeploymentSettingsResponseOutputWithContext(ctx context.Context) DeploymentSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentSettingsResponseOutput)
}

func (i DeploymentSettingsResponseArgs) ToDeploymentSettingsResponsePtrOutput() DeploymentSettingsResponsePtrOutput {
	return i.ToDeploymentSettingsResponsePtrOutputWithContext(context.Background())
}

func (i DeploymentSettingsResponseArgs) ToDeploymentSettingsResponsePtrOutputWithContext(ctx context.Context) DeploymentSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentSettingsResponseOutput).ToDeploymentSettingsResponsePtrOutputWithContext(ctx)
}









type DeploymentSettingsResponsePtrInput interface {
	pulumi.Input

	ToDeploymentSettingsResponsePtrOutput() DeploymentSettingsResponsePtrOutput
	ToDeploymentSettingsResponsePtrOutputWithContext(context.Context) DeploymentSettingsResponsePtrOutput
}

type deploymentSettingsResponsePtrType DeploymentSettingsResponseArgs

func DeploymentSettingsResponsePtr(v *DeploymentSettingsResponseArgs) DeploymentSettingsResponsePtrInput {
	return (*deploymentSettingsResponsePtrType)(v)
}

func (*deploymentSettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentSettingsResponse)(nil)).Elem()
}

func (i *deploymentSettingsResponsePtrType) ToDeploymentSettingsResponsePtrOutput() DeploymentSettingsResponsePtrOutput {
	return i.ToDeploymentSettingsResponsePtrOutputWithContext(context.Background())
}

func (i *deploymentSettingsResponsePtrType) ToDeploymentSettingsResponsePtrOutputWithContext(ctx context.Context) DeploymentSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentSettingsResponsePtrOutput)
}

type DeploymentSettingsResponseOutput struct{ *pulumi.OutputState }

func (DeploymentSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentSettingsResponse)(nil)).Elem()
}

func (o DeploymentSettingsResponseOutput) ToDeploymentSettingsResponseOutput() DeploymentSettingsResponseOutput {
	return o
}

func (o DeploymentSettingsResponseOutput) ToDeploymentSettingsResponseOutputWithContext(ctx context.Context) DeploymentSettingsResponseOutput {
	return o
}

func (o DeploymentSettingsResponseOutput) ToDeploymentSettingsResponsePtrOutput() DeploymentSettingsResponsePtrOutput {
	return o.ToDeploymentSettingsResponsePtrOutputWithContext(context.Background())
}

func (o DeploymentSettingsResponseOutput) ToDeploymentSettingsResponsePtrOutputWithContext(ctx context.Context) DeploymentSettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeploymentSettingsResponse) *DeploymentSettingsResponse {
		return &v
	}).(DeploymentSettingsResponsePtrOutput)
}

func (o DeploymentSettingsResponseOutput) Cpu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DeploymentSettingsResponse) *int { return v.Cpu }).(pulumi.IntPtrOutput)
}

func (o DeploymentSettingsResponseOutput) EnvironmentVariables() pulumi.StringMapOutput {
	return o.ApplyT(func(v DeploymentSettingsResponse) map[string]string { return v.EnvironmentVariables }).(pulumi.StringMapOutput)
}

func (o DeploymentSettingsResponseOutput) JvmOptions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentSettingsResponse) *string { return v.JvmOptions }).(pulumi.StringPtrOutput)
}

func (o DeploymentSettingsResponseOutput) MemoryInGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DeploymentSettingsResponse) *int { return v.MemoryInGB }).(pulumi.IntPtrOutput)
}

func (o DeploymentSettingsResponseOutput) NetCoreMainEntryPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentSettingsResponse) *string { return v.NetCoreMainEntryPath }).(pulumi.StringPtrOutput)
}

func (o DeploymentSettingsResponseOutput) ResourceRequests() ResourceRequestsResponsePtrOutput {
	return o.ApplyT(func(v DeploymentSettingsResponse) *ResourceRequestsResponse { return v.ResourceRequests }).(ResourceRequestsResponsePtrOutput)
}

func (o DeploymentSettingsResponseOutput) RuntimeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentSettingsResponse) *string { return v.RuntimeVersion }).(pulumi.StringPtrOutput)
}

type DeploymentSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (DeploymentSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentSettingsResponse)(nil)).Elem()
}

func (o DeploymentSettingsResponsePtrOutput) ToDeploymentSettingsResponsePtrOutput() DeploymentSettingsResponsePtrOutput {
	return o
}

func (o DeploymentSettingsResponsePtrOutput) ToDeploymentSettingsResponsePtrOutputWithContext(ctx context.Context) DeploymentSettingsResponsePtrOutput {
	return o
}

func (o DeploymentSettingsResponsePtrOutput) Elem() DeploymentSettingsResponseOutput {
	return o.ApplyT(func(v *DeploymentSettingsResponse) DeploymentSettingsResponse {
		if v != nil {
			return *v
		}
		var ret DeploymentSettingsResponse
		return ret
	}).(DeploymentSettingsResponseOutput)
}

func (o DeploymentSettingsResponsePtrOutput) Cpu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DeploymentSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.Cpu
	}).(pulumi.IntPtrOutput)
}

func (o DeploymentSettingsResponsePtrOutput) EnvironmentVariables() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DeploymentSettingsResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.EnvironmentVariables
	}).(pulumi.StringMapOutput)
}

func (o DeploymentSettingsResponsePtrOutput) JvmOptions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.JvmOptions
	}).(pulumi.StringPtrOutput)
}

func (o DeploymentSettingsResponsePtrOutput) MemoryInGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DeploymentSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.MemoryInGB
	}).(pulumi.IntPtrOutput)
}

func (o DeploymentSettingsResponsePtrOutput) NetCoreMainEntryPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.NetCoreMainEntryPath
	}).(pulumi.StringPtrOutput)
}

func (o DeploymentSettingsResponsePtrOutput) ResourceRequests() ResourceRequestsResponsePtrOutput {
	return o.ApplyT(func(v *DeploymentSettingsResponse) *ResourceRequestsResponse {
		if v == nil {
			return nil
		}
		return v.ResourceRequests
	}).(ResourceRequestsResponsePtrOutput)
}

func (o DeploymentSettingsResponsePtrOutput) RuntimeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.RuntimeVersion
	}).(pulumi.StringPtrOutput)
}

type ImageRegistryCredential struct {
	Password *string `pulumi:"password"`
	Username *string `pulumi:"username"`
}





type ImageRegistryCredentialInput interface {
	pulumi.Input

	ToImageRegistryCredentialOutput() ImageRegistryCredentialOutput
	ToImageRegistryCredentialOutputWithContext(context.Context) ImageRegistryCredentialOutput
}

type ImageRegistryCredentialArgs struct {
	Password pulumi.StringPtrInput `pulumi:"password"`
	Username pulumi.StringPtrInput `pulumi:"username"`
}

func (ImageRegistryCredentialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageRegistryCredential)(nil)).Elem()
}

func (i ImageRegistryCredentialArgs) ToImageRegistryCredentialOutput() ImageRegistryCredentialOutput {
	return i.ToImageRegistryCredentialOutputWithContext(context.Background())
}

func (i ImageRegistryCredentialArgs) ToImageRegistryCredentialOutputWithContext(ctx context.Context) ImageRegistryCredentialOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageRegistryCredentialOutput)
}

func (i ImageRegistryCredentialArgs) ToImageRegistryCredentialPtrOutput() ImageRegistryCredentialPtrOutput {
	return i.ToImageRegistryCredentialPtrOutputWithContext(context.Background())
}

func (i ImageRegistryCredentialArgs) ToImageRegistryCredentialPtrOutputWithContext(ctx context.Context) ImageRegistryCredentialPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageRegistryCredentialOutput).ToImageRegistryCredentialPtrOutputWithContext(ctx)
}









type ImageRegistryCredentialPtrInput interface {
	pulumi.Input

	ToImageRegistryCredentialPtrOutput() ImageRegistryCredentialPtrOutput
	ToImageRegistryCredentialPtrOutputWithContext(context.Context) ImageRegistryCredentialPtrOutput
}

type imageRegistryCredentialPtrType ImageRegistryCredentialArgs

func ImageRegistryCredentialPtr(v *ImageRegistryCredentialArgs) ImageRegistryCredentialPtrInput {
	return (*imageRegistryCredentialPtrType)(v)
}

func (*imageRegistryCredentialPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageRegistryCredential)(nil)).Elem()
}

func (i *imageRegistryCredentialPtrType) ToImageRegistryCredentialPtrOutput() ImageRegistryCredentialPtrOutput {
	return i.ToImageRegistryCredentialPtrOutputWithContext(context.Background())
}

func (i *imageRegistryCredentialPtrType) ToImageRegistryCredentialPtrOutputWithContext(ctx context.Context) ImageRegistryCredentialPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageRegistryCredentialPtrOutput)
}

type ImageRegistryCredentialOutput struct{ *pulumi.OutputState }

func (ImageRegistryCredentialOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageRegistryCredential)(nil)).Elem()
}

func (o ImageRegistryCredentialOutput) ToImageRegistryCredentialOutput() ImageRegistryCredentialOutput {
	return o
}

func (o ImageRegistryCredentialOutput) ToImageRegistryCredentialOutputWithContext(ctx context.Context) ImageRegistryCredentialOutput {
	return o
}

func (o ImageRegistryCredentialOutput) ToImageRegistryCredentialPtrOutput() ImageRegistryCredentialPtrOutput {
	return o.ToImageRegistryCredentialPtrOutputWithContext(context.Background())
}

func (o ImageRegistryCredentialOutput) ToImageRegistryCredentialPtrOutputWithContext(ctx context.Context) ImageRegistryCredentialPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImageRegistryCredential) *ImageRegistryCredential {
		return &v
	}).(ImageRegistryCredentialPtrOutput)
}

func (o ImageRegistryCredentialOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageRegistryCredential) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o ImageRegistryCredentialOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageRegistryCredential) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type ImageRegistryCredentialPtrOutput struct{ *pulumi.OutputState }

func (ImageRegistryCredentialPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageRegistryCredential)(nil)).Elem()
}

func (o ImageRegistryCredentialPtrOutput) ToImageRegistryCredentialPtrOutput() ImageRegistryCredentialPtrOutput {
	return o
}

func (o ImageRegistryCredentialPtrOutput) ToImageRegistryCredentialPtrOutputWithContext(ctx context.Context) ImageRegistryCredentialPtrOutput {
	return o
}

func (o ImageRegistryCredentialPtrOutput) Elem() ImageRegistryCredentialOutput {
	return o.ApplyT(func(v *ImageRegistryCredential) ImageRegistryCredential {
		if v != nil {
			return *v
		}
		var ret ImageRegistryCredential
		return ret
	}).(ImageRegistryCredentialOutput)
}

func (o ImageRegistryCredentialPtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageRegistryCredential) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

func (o ImageRegistryCredentialPtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageRegistryCredential) *string {
		if v == nil {
			return nil
		}
		return v.Username
	}).(pulumi.StringPtrOutput)
}

type ImageRegistryCredentialResponse struct {
	Password *string `pulumi:"password"`
	Username *string `pulumi:"username"`
}





type ImageRegistryCredentialResponseInput interface {
	pulumi.Input

	ToImageRegistryCredentialResponseOutput() ImageRegistryCredentialResponseOutput
	ToImageRegistryCredentialResponseOutputWithContext(context.Context) ImageRegistryCredentialResponseOutput
}

type ImageRegistryCredentialResponseArgs struct {
	Password pulumi.StringPtrInput `pulumi:"password"`
	Username pulumi.StringPtrInput `pulumi:"username"`
}

func (ImageRegistryCredentialResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageRegistryCredentialResponse)(nil)).Elem()
}

func (i ImageRegistryCredentialResponseArgs) ToImageRegistryCredentialResponseOutput() ImageRegistryCredentialResponseOutput {
	return i.ToImageRegistryCredentialResponseOutputWithContext(context.Background())
}

func (i ImageRegistryCredentialResponseArgs) ToImageRegistryCredentialResponseOutputWithContext(ctx context.Context) ImageRegistryCredentialResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageRegistryCredentialResponseOutput)
}

func (i ImageRegistryCredentialResponseArgs) ToImageRegistryCredentialResponsePtrOutput() ImageRegistryCredentialResponsePtrOutput {
	return i.ToImageRegistryCredentialResponsePtrOutputWithContext(context.Background())
}

func (i ImageRegistryCredentialResponseArgs) ToImageRegistryCredentialResponsePtrOutputWithContext(ctx context.Context) ImageRegistryCredentialResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageRegistryCredentialResponseOutput).ToImageRegistryCredentialResponsePtrOutputWithContext(ctx)
}









type ImageRegistryCredentialResponsePtrInput interface {
	pulumi.Input

	ToImageRegistryCredentialResponsePtrOutput() ImageRegistryCredentialResponsePtrOutput
	ToImageRegistryCredentialResponsePtrOutputWithContext(context.Context) ImageRegistryCredentialResponsePtrOutput
}

type imageRegistryCredentialResponsePtrType ImageRegistryCredentialResponseArgs

func ImageRegistryCredentialResponsePtr(v *ImageRegistryCredentialResponseArgs) ImageRegistryCredentialResponsePtrInput {
	return (*imageRegistryCredentialResponsePtrType)(v)
}

func (*imageRegistryCredentialResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageRegistryCredentialResponse)(nil)).Elem()
}

func (i *imageRegistryCredentialResponsePtrType) ToImageRegistryCredentialResponsePtrOutput() ImageRegistryCredentialResponsePtrOutput {
	return i.ToImageRegistryCredentialResponsePtrOutputWithContext(context.Background())
}

func (i *imageRegistryCredentialResponsePtrType) ToImageRegistryCredentialResponsePtrOutputWithContext(ctx context.Context) ImageRegistryCredentialResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageRegistryCredentialResponsePtrOutput)
}

type ImageRegistryCredentialResponseOutput struct{ *pulumi.OutputState }

func (ImageRegistryCredentialResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageRegistryCredentialResponse)(nil)).Elem()
}

func (o ImageRegistryCredentialResponseOutput) ToImageRegistryCredentialResponseOutput() ImageRegistryCredentialResponseOutput {
	return o
}

func (o ImageRegistryCredentialResponseOutput) ToImageRegistryCredentialResponseOutputWithContext(ctx context.Context) ImageRegistryCredentialResponseOutput {
	return o
}

func (o ImageRegistryCredentialResponseOutput) ToImageRegistryCredentialResponsePtrOutput() ImageRegistryCredentialResponsePtrOutput {
	return o.ToImageRegistryCredentialResponsePtrOutputWithContext(context.Background())
}

func (o ImageRegistryCredentialResponseOutput) ToImageRegistryCredentialResponsePtrOutputWithContext(ctx context.Context) ImageRegistryCredentialResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImageRegistryCredentialResponse) *ImageRegistryCredentialResponse {
		return &v
	}).(ImageRegistryCredentialResponsePtrOutput)
}

func (o ImageRegistryCredentialResponseOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageRegistryCredentialResponse) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o ImageRegistryCredentialResponseOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageRegistryCredentialResponse) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type ImageRegistryCredentialResponsePtrOutput struct{ *pulumi.OutputState }

func (ImageRegistryCredentialResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageRegistryCredentialResponse)(nil)).Elem()
}

func (o ImageRegistryCredentialResponsePtrOutput) ToImageRegistryCredentialResponsePtrOutput() ImageRegistryCredentialResponsePtrOutput {
	return o
}

func (o ImageRegistryCredentialResponsePtrOutput) ToImageRegistryCredentialResponsePtrOutputWithContext(ctx context.Context) ImageRegistryCredentialResponsePtrOutput {
	return o
}

func (o ImageRegistryCredentialResponsePtrOutput) Elem() ImageRegistryCredentialResponseOutput {
	return o.ApplyT(func(v *ImageRegistryCredentialResponse) ImageRegistryCredentialResponse {
		if v != nil {
			return *v
		}
		var ret ImageRegistryCredentialResponse
		return ret
	}).(ImageRegistryCredentialResponseOutput)
}

func (o ImageRegistryCredentialResponsePtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageRegistryCredentialResponse) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

func (o ImageRegistryCredentialResponsePtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageRegistryCredentialResponse) *string {
		if v == nil {
			return nil
		}
		return v.Username
	}).(pulumi.StringPtrOutput)
}

type ManagedIdentityProperties struct {
	PrincipalId *string `pulumi:"principalId"`
	TenantId    *string `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}





type ManagedIdentityPropertiesInput interface {
	pulumi.Input

	ToManagedIdentityPropertiesOutput() ManagedIdentityPropertiesOutput
	ToManagedIdentityPropertiesOutputWithContext(context.Context) ManagedIdentityPropertiesOutput
}

type ManagedIdentityPropertiesArgs struct {
	PrincipalId pulumi.StringPtrInput `pulumi:"principalId"`
	TenantId    pulumi.StringPtrInput `pulumi:"tenantId"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
}

func (ManagedIdentityPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentityProperties)(nil)).Elem()
}

func (i ManagedIdentityPropertiesArgs) ToManagedIdentityPropertiesOutput() ManagedIdentityPropertiesOutput {
	return i.ToManagedIdentityPropertiesOutputWithContext(context.Background())
}

func (i ManagedIdentityPropertiesArgs) ToManagedIdentityPropertiesOutputWithContext(ctx context.Context) ManagedIdentityPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentityPropertiesOutput)
}

func (i ManagedIdentityPropertiesArgs) ToManagedIdentityPropertiesPtrOutput() ManagedIdentityPropertiesPtrOutput {
	return i.ToManagedIdentityPropertiesPtrOutputWithContext(context.Background())
}

func (i ManagedIdentityPropertiesArgs) ToManagedIdentityPropertiesPtrOutputWithContext(ctx context.Context) ManagedIdentityPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentityPropertiesOutput).ToManagedIdentityPropertiesPtrOutputWithContext(ctx)
}









type ManagedIdentityPropertiesPtrInput interface {
	pulumi.Input

	ToManagedIdentityPropertiesPtrOutput() ManagedIdentityPropertiesPtrOutput
	ToManagedIdentityPropertiesPtrOutputWithContext(context.Context) ManagedIdentityPropertiesPtrOutput
}

type managedIdentityPropertiesPtrType ManagedIdentityPropertiesArgs

func ManagedIdentityPropertiesPtr(v *ManagedIdentityPropertiesArgs) ManagedIdentityPropertiesPtrInput {
	return (*managedIdentityPropertiesPtrType)(v)
}

func (*managedIdentityPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentityProperties)(nil)).Elem()
}

func (i *managedIdentityPropertiesPtrType) ToManagedIdentityPropertiesPtrOutput() ManagedIdentityPropertiesPtrOutput {
	return i.ToManagedIdentityPropertiesPtrOutputWithContext(context.Background())
}

func (i *managedIdentityPropertiesPtrType) ToManagedIdentityPropertiesPtrOutputWithContext(ctx context.Context) ManagedIdentityPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentityPropertiesPtrOutput)
}

type ManagedIdentityPropertiesOutput struct{ *pulumi.OutputState }

func (ManagedIdentityPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentityProperties)(nil)).Elem()
}

func (o ManagedIdentityPropertiesOutput) ToManagedIdentityPropertiesOutput() ManagedIdentityPropertiesOutput {
	return o
}

func (o ManagedIdentityPropertiesOutput) ToManagedIdentityPropertiesOutputWithContext(ctx context.Context) ManagedIdentityPropertiesOutput {
	return o
}

func (o ManagedIdentityPropertiesOutput) ToManagedIdentityPropertiesPtrOutput() ManagedIdentityPropertiesPtrOutput {
	return o.ToManagedIdentityPropertiesPtrOutputWithContext(context.Background())
}

func (o ManagedIdentityPropertiesOutput) ToManagedIdentityPropertiesPtrOutputWithContext(ctx context.Context) ManagedIdentityPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedIdentityProperties) *ManagedIdentityProperties {
		return &v
	}).(ManagedIdentityPropertiesPtrOutput)
}

func (o ManagedIdentityPropertiesOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedIdentityProperties) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityPropertiesOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedIdentityProperties) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityPropertiesOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedIdentityProperties) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ManagedIdentityPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ManagedIdentityPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentityProperties)(nil)).Elem()
}

func (o ManagedIdentityPropertiesPtrOutput) ToManagedIdentityPropertiesPtrOutput() ManagedIdentityPropertiesPtrOutput {
	return o
}

func (o ManagedIdentityPropertiesPtrOutput) ToManagedIdentityPropertiesPtrOutputWithContext(ctx context.Context) ManagedIdentityPropertiesPtrOutput {
	return o
}

func (o ManagedIdentityPropertiesPtrOutput) Elem() ManagedIdentityPropertiesOutput {
	return o.ApplyT(func(v *ManagedIdentityProperties) ManagedIdentityProperties {
		if v != nil {
			return *v
		}
		var ret ManagedIdentityProperties
		return ret
	}).(ManagedIdentityPropertiesOutput)
}

func (o ManagedIdentityPropertiesPtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedIdentityProperties) *string {
		if v == nil {
			return nil
		}
		return v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityPropertiesPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedIdentityProperties) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityPropertiesPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedIdentityProperties) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ManagedIdentityPropertiesResponse struct {
	PrincipalId *string `pulumi:"principalId"`
	TenantId    *string `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}





type ManagedIdentityPropertiesResponseInput interface {
	pulumi.Input

	ToManagedIdentityPropertiesResponseOutput() ManagedIdentityPropertiesResponseOutput
	ToManagedIdentityPropertiesResponseOutputWithContext(context.Context) ManagedIdentityPropertiesResponseOutput
}

type ManagedIdentityPropertiesResponseArgs struct {
	PrincipalId pulumi.StringPtrInput `pulumi:"principalId"`
	TenantId    pulumi.StringPtrInput `pulumi:"tenantId"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
}

func (ManagedIdentityPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentityPropertiesResponse)(nil)).Elem()
}

func (i ManagedIdentityPropertiesResponseArgs) ToManagedIdentityPropertiesResponseOutput() ManagedIdentityPropertiesResponseOutput {
	return i.ToManagedIdentityPropertiesResponseOutputWithContext(context.Background())
}

func (i ManagedIdentityPropertiesResponseArgs) ToManagedIdentityPropertiesResponseOutputWithContext(ctx context.Context) ManagedIdentityPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentityPropertiesResponseOutput)
}

func (i ManagedIdentityPropertiesResponseArgs) ToManagedIdentityPropertiesResponsePtrOutput() ManagedIdentityPropertiesResponsePtrOutput {
	return i.ToManagedIdentityPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ManagedIdentityPropertiesResponseArgs) ToManagedIdentityPropertiesResponsePtrOutputWithContext(ctx context.Context) ManagedIdentityPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentityPropertiesResponseOutput).ToManagedIdentityPropertiesResponsePtrOutputWithContext(ctx)
}









type ManagedIdentityPropertiesResponsePtrInput interface {
	pulumi.Input

	ToManagedIdentityPropertiesResponsePtrOutput() ManagedIdentityPropertiesResponsePtrOutput
	ToManagedIdentityPropertiesResponsePtrOutputWithContext(context.Context) ManagedIdentityPropertiesResponsePtrOutput
}

type managedIdentityPropertiesResponsePtrType ManagedIdentityPropertiesResponseArgs

func ManagedIdentityPropertiesResponsePtr(v *ManagedIdentityPropertiesResponseArgs) ManagedIdentityPropertiesResponsePtrInput {
	return (*managedIdentityPropertiesResponsePtrType)(v)
}

func (*managedIdentityPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentityPropertiesResponse)(nil)).Elem()
}

func (i *managedIdentityPropertiesResponsePtrType) ToManagedIdentityPropertiesResponsePtrOutput() ManagedIdentityPropertiesResponsePtrOutput {
	return i.ToManagedIdentityPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *managedIdentityPropertiesResponsePtrType) ToManagedIdentityPropertiesResponsePtrOutputWithContext(ctx context.Context) ManagedIdentityPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentityPropertiesResponsePtrOutput)
}

type ManagedIdentityPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ManagedIdentityPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentityPropertiesResponse)(nil)).Elem()
}

func (o ManagedIdentityPropertiesResponseOutput) ToManagedIdentityPropertiesResponseOutput() ManagedIdentityPropertiesResponseOutput {
	return o
}

func (o ManagedIdentityPropertiesResponseOutput) ToManagedIdentityPropertiesResponseOutputWithContext(ctx context.Context) ManagedIdentityPropertiesResponseOutput {
	return o
}

func (o ManagedIdentityPropertiesResponseOutput) ToManagedIdentityPropertiesResponsePtrOutput() ManagedIdentityPropertiesResponsePtrOutput {
	return o.ToManagedIdentityPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ManagedIdentityPropertiesResponseOutput) ToManagedIdentityPropertiesResponsePtrOutputWithContext(ctx context.Context) ManagedIdentityPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedIdentityPropertiesResponse) *ManagedIdentityPropertiesResponse {
		return &v
	}).(ManagedIdentityPropertiesResponsePtrOutput)
}

func (o ManagedIdentityPropertiesResponseOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedIdentityPropertiesResponse) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityPropertiesResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedIdentityPropertiesResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityPropertiesResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedIdentityPropertiesResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ManagedIdentityPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedIdentityPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentityPropertiesResponse)(nil)).Elem()
}

func (o ManagedIdentityPropertiesResponsePtrOutput) ToManagedIdentityPropertiesResponsePtrOutput() ManagedIdentityPropertiesResponsePtrOutput {
	return o
}

func (o ManagedIdentityPropertiesResponsePtrOutput) ToManagedIdentityPropertiesResponsePtrOutputWithContext(ctx context.Context) ManagedIdentityPropertiesResponsePtrOutput {
	return o
}

func (o ManagedIdentityPropertiesResponsePtrOutput) Elem() ManagedIdentityPropertiesResponseOutput {
	return o.ApplyT(func(v *ManagedIdentityPropertiesResponse) ManagedIdentityPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ManagedIdentityPropertiesResponse
		return ret
	}).(ManagedIdentityPropertiesResponseOutput)
}

func (o ManagedIdentityPropertiesResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedIdentityPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityPropertiesResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedIdentityPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityPropertiesResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedIdentityPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type NetworkProfile struct {
	AppNetworkResourceGroup            *string `pulumi:"appNetworkResourceGroup"`
	AppSubnetId                        *string `pulumi:"appSubnetId"`
	ServiceCidr                        *string `pulumi:"serviceCidr"`
	ServiceRuntimeNetworkResourceGroup *string `pulumi:"serviceRuntimeNetworkResourceGroup"`
	ServiceRuntimeSubnetId             *string `pulumi:"serviceRuntimeSubnetId"`
}





type NetworkProfileInput interface {
	pulumi.Input

	ToNetworkProfileOutput() NetworkProfileOutput
	ToNetworkProfileOutputWithContext(context.Context) NetworkProfileOutput
}

type NetworkProfileArgs struct {
	AppNetworkResourceGroup            pulumi.StringPtrInput `pulumi:"appNetworkResourceGroup"`
	AppSubnetId                        pulumi.StringPtrInput `pulumi:"appSubnetId"`
	ServiceCidr                        pulumi.StringPtrInput `pulumi:"serviceCidr"`
	ServiceRuntimeNetworkResourceGroup pulumi.StringPtrInput `pulumi:"serviceRuntimeNetworkResourceGroup"`
	ServiceRuntimeSubnetId             pulumi.StringPtrInput `pulumi:"serviceRuntimeSubnetId"`
}

func (NetworkProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkProfile)(nil)).Elem()
}

func (i NetworkProfileArgs) ToNetworkProfileOutput() NetworkProfileOutput {
	return i.ToNetworkProfileOutputWithContext(context.Background())
}

func (i NetworkProfileArgs) ToNetworkProfileOutputWithContext(ctx context.Context) NetworkProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfileOutput)
}

func (i NetworkProfileArgs) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return i.ToNetworkProfilePtrOutputWithContext(context.Background())
}

func (i NetworkProfileArgs) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfileOutput).ToNetworkProfilePtrOutputWithContext(ctx)
}









type NetworkProfilePtrInput interface {
	pulumi.Input

	ToNetworkProfilePtrOutput() NetworkProfilePtrOutput
	ToNetworkProfilePtrOutputWithContext(context.Context) NetworkProfilePtrOutput
}

type networkProfilePtrType NetworkProfileArgs

func NetworkProfilePtr(v *NetworkProfileArgs) NetworkProfilePtrInput {
	return (*networkProfilePtrType)(v)
}

func (*networkProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfile)(nil)).Elem()
}

func (i *networkProfilePtrType) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return i.ToNetworkProfilePtrOutputWithContext(context.Background())
}

func (i *networkProfilePtrType) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfilePtrOutput)
}

type NetworkProfileOutput struct{ *pulumi.OutputState }

func (NetworkProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkProfile)(nil)).Elem()
}

func (o NetworkProfileOutput) ToNetworkProfileOutput() NetworkProfileOutput {
	return o
}

func (o NetworkProfileOutput) ToNetworkProfileOutputWithContext(ctx context.Context) NetworkProfileOutput {
	return o
}

func (o NetworkProfileOutput) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return o.ToNetworkProfilePtrOutputWithContext(context.Background())
}

func (o NetworkProfileOutput) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkProfile) *NetworkProfile {
		return &v
	}).(NetworkProfilePtrOutput)
}

func (o NetworkProfileOutput) AppNetworkResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfile) *string { return v.AppNetworkResourceGroup }).(pulumi.StringPtrOutput)
}

func (o NetworkProfileOutput) AppSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfile) *string { return v.AppSubnetId }).(pulumi.StringPtrOutput)
}

func (o NetworkProfileOutput) ServiceCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfile) *string { return v.ServiceCidr }).(pulumi.StringPtrOutput)
}

func (o NetworkProfileOutput) ServiceRuntimeNetworkResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfile) *string { return v.ServiceRuntimeNetworkResourceGroup }).(pulumi.StringPtrOutput)
}

func (o NetworkProfileOutput) ServiceRuntimeSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfile) *string { return v.ServiceRuntimeSubnetId }).(pulumi.StringPtrOutput)
}

type NetworkProfilePtrOutput struct{ *pulumi.OutputState }

func (NetworkProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfile)(nil)).Elem()
}

func (o NetworkProfilePtrOutput) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return o
}

func (o NetworkProfilePtrOutput) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return o
}

func (o NetworkProfilePtrOutput) Elem() NetworkProfileOutput {
	return o.ApplyT(func(v *NetworkProfile) NetworkProfile {
		if v != nil {
			return *v
		}
		var ret NetworkProfile
		return ret
	}).(NetworkProfileOutput)
}

func (o NetworkProfilePtrOutput) AppNetworkResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.AppNetworkResourceGroup
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfilePtrOutput) AppSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.AppSubnetId
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfilePtrOutput) ServiceCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.ServiceCidr
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfilePtrOutput) ServiceRuntimeNetworkResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.ServiceRuntimeNetworkResourceGroup
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfilePtrOutput) ServiceRuntimeSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.ServiceRuntimeSubnetId
	}).(pulumi.StringPtrOutput)
}

type NetworkProfileResponse struct {
	AppNetworkResourceGroup            *string                           `pulumi:"appNetworkResourceGroup"`
	AppSubnetId                        *string                           `pulumi:"appSubnetId"`
	OutboundIPs                        NetworkProfileResponseOutboundIPs `pulumi:"outboundIPs"`
	RequiredTraffics                   []RequiredTrafficResponse         `pulumi:"requiredTraffics"`
	ServiceCidr                        *string                           `pulumi:"serviceCidr"`
	ServiceRuntimeNetworkResourceGroup *string                           `pulumi:"serviceRuntimeNetworkResourceGroup"`
	ServiceRuntimeSubnetId             *string                           `pulumi:"serviceRuntimeSubnetId"`
}





type NetworkProfileResponseInput interface {
	pulumi.Input

	ToNetworkProfileResponseOutput() NetworkProfileResponseOutput
	ToNetworkProfileResponseOutputWithContext(context.Context) NetworkProfileResponseOutput
}

type NetworkProfileResponseArgs struct {
	AppNetworkResourceGroup            pulumi.StringPtrInput                  `pulumi:"appNetworkResourceGroup"`
	AppSubnetId                        pulumi.StringPtrInput                  `pulumi:"appSubnetId"`
	OutboundIPs                        NetworkProfileResponseOutboundIPsInput `pulumi:"outboundIPs"`
	RequiredTraffics                   RequiredTrafficResponseArrayInput      `pulumi:"requiredTraffics"`
	ServiceCidr                        pulumi.StringPtrInput                  `pulumi:"serviceCidr"`
	ServiceRuntimeNetworkResourceGroup pulumi.StringPtrInput                  `pulumi:"serviceRuntimeNetworkResourceGroup"`
	ServiceRuntimeSubnetId             pulumi.StringPtrInput                  `pulumi:"serviceRuntimeSubnetId"`
}

func (NetworkProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkProfileResponse)(nil)).Elem()
}

func (i NetworkProfileResponseArgs) ToNetworkProfileResponseOutput() NetworkProfileResponseOutput {
	return i.ToNetworkProfileResponseOutputWithContext(context.Background())
}

func (i NetworkProfileResponseArgs) ToNetworkProfileResponseOutputWithContext(ctx context.Context) NetworkProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfileResponseOutput)
}

func (i NetworkProfileResponseArgs) ToNetworkProfileResponsePtrOutput() NetworkProfileResponsePtrOutput {
	return i.ToNetworkProfileResponsePtrOutputWithContext(context.Background())
}

func (i NetworkProfileResponseArgs) ToNetworkProfileResponsePtrOutputWithContext(ctx context.Context) NetworkProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfileResponseOutput).ToNetworkProfileResponsePtrOutputWithContext(ctx)
}









type NetworkProfileResponsePtrInput interface {
	pulumi.Input

	ToNetworkProfileResponsePtrOutput() NetworkProfileResponsePtrOutput
	ToNetworkProfileResponsePtrOutputWithContext(context.Context) NetworkProfileResponsePtrOutput
}

type networkProfileResponsePtrType NetworkProfileResponseArgs

func NetworkProfileResponsePtr(v *NetworkProfileResponseArgs) NetworkProfileResponsePtrInput {
	return (*networkProfileResponsePtrType)(v)
}

func (*networkProfileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfileResponse)(nil)).Elem()
}

func (i *networkProfileResponsePtrType) ToNetworkProfileResponsePtrOutput() NetworkProfileResponsePtrOutput {
	return i.ToNetworkProfileResponsePtrOutputWithContext(context.Background())
}

func (i *networkProfileResponsePtrType) ToNetworkProfileResponsePtrOutputWithContext(ctx context.Context) NetworkProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfileResponsePtrOutput)
}

type NetworkProfileResponseOutput struct{ *pulumi.OutputState }

func (NetworkProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkProfileResponse)(nil)).Elem()
}

func (o NetworkProfileResponseOutput) ToNetworkProfileResponseOutput() NetworkProfileResponseOutput {
	return o
}

func (o NetworkProfileResponseOutput) ToNetworkProfileResponseOutputWithContext(ctx context.Context) NetworkProfileResponseOutput {
	return o
}

func (o NetworkProfileResponseOutput) ToNetworkProfileResponsePtrOutput() NetworkProfileResponsePtrOutput {
	return o.ToNetworkProfileResponsePtrOutputWithContext(context.Background())
}

func (o NetworkProfileResponseOutput) ToNetworkProfileResponsePtrOutputWithContext(ctx context.Context) NetworkProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkProfileResponse) *NetworkProfileResponse {
		return &v
	}).(NetworkProfileResponsePtrOutput)
}

func (o NetworkProfileResponseOutput) AppNetworkResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfileResponse) *string { return v.AppNetworkResourceGroup }).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponseOutput) AppSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfileResponse) *string { return v.AppSubnetId }).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponseOutput) OutboundIPs() NetworkProfileResponseOutboundIPsOutput {
	return o.ApplyT(func(v NetworkProfileResponse) NetworkProfileResponseOutboundIPs { return v.OutboundIPs }).(NetworkProfileResponseOutboundIPsOutput)
}

func (o NetworkProfileResponseOutput) RequiredTraffics() RequiredTrafficResponseArrayOutput {
	return o.ApplyT(func(v NetworkProfileResponse) []RequiredTrafficResponse { return v.RequiredTraffics }).(RequiredTrafficResponseArrayOutput)
}

func (o NetworkProfileResponseOutput) ServiceCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfileResponse) *string { return v.ServiceCidr }).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponseOutput) ServiceRuntimeNetworkResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfileResponse) *string { return v.ServiceRuntimeNetworkResourceGroup }).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponseOutput) ServiceRuntimeSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfileResponse) *string { return v.ServiceRuntimeSubnetId }).(pulumi.StringPtrOutput)
}

type NetworkProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (NetworkProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfileResponse)(nil)).Elem()
}

func (o NetworkProfileResponsePtrOutput) ToNetworkProfileResponsePtrOutput() NetworkProfileResponsePtrOutput {
	return o
}

func (o NetworkProfileResponsePtrOutput) ToNetworkProfileResponsePtrOutputWithContext(ctx context.Context) NetworkProfileResponsePtrOutput {
	return o
}

func (o NetworkProfileResponsePtrOutput) Elem() NetworkProfileResponseOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) NetworkProfileResponse {
		if v != nil {
			return *v
		}
		var ret NetworkProfileResponse
		return ret
	}).(NetworkProfileResponseOutput)
}

func (o NetworkProfileResponsePtrOutput) AppNetworkResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.AppNetworkResourceGroup
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponsePtrOutput) AppSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.AppSubnetId
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponsePtrOutput) OutboundIPs() NetworkProfileResponseOutboundIPsPtrOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) *NetworkProfileResponseOutboundIPs {
		if v == nil {
			return nil
		}
		return &v.OutboundIPs
	}).(NetworkProfileResponseOutboundIPsPtrOutput)
}

func (o NetworkProfileResponsePtrOutput) RequiredTraffics() RequiredTrafficResponseArrayOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) []RequiredTrafficResponse {
		if v == nil {
			return nil
		}
		return v.RequiredTraffics
	}).(RequiredTrafficResponseArrayOutput)
}

func (o NetworkProfileResponsePtrOutput) ServiceCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServiceCidr
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponsePtrOutput) ServiceRuntimeNetworkResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServiceRuntimeNetworkResourceGroup
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponsePtrOutput) ServiceRuntimeSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServiceRuntimeSubnetId
	}).(pulumi.StringPtrOutput)
}

type NetworkProfileResponseOutboundIPs struct {
	PublicIPs []string `pulumi:"publicIPs"`
}





type NetworkProfileResponseOutboundIPsInput interface {
	pulumi.Input

	ToNetworkProfileResponseOutboundIPsOutput() NetworkProfileResponseOutboundIPsOutput
	ToNetworkProfileResponseOutboundIPsOutputWithContext(context.Context) NetworkProfileResponseOutboundIPsOutput
}

type NetworkProfileResponseOutboundIPsArgs struct {
	PublicIPs pulumi.StringArrayInput `pulumi:"publicIPs"`
}

func (NetworkProfileResponseOutboundIPsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkProfileResponseOutboundIPs)(nil)).Elem()
}

func (i NetworkProfileResponseOutboundIPsArgs) ToNetworkProfileResponseOutboundIPsOutput() NetworkProfileResponseOutboundIPsOutput {
	return i.ToNetworkProfileResponseOutboundIPsOutputWithContext(context.Background())
}

func (i NetworkProfileResponseOutboundIPsArgs) ToNetworkProfileResponseOutboundIPsOutputWithContext(ctx context.Context) NetworkProfileResponseOutboundIPsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfileResponseOutboundIPsOutput)
}

func (i NetworkProfileResponseOutboundIPsArgs) ToNetworkProfileResponseOutboundIPsPtrOutput() NetworkProfileResponseOutboundIPsPtrOutput {
	return i.ToNetworkProfileResponseOutboundIPsPtrOutputWithContext(context.Background())
}

func (i NetworkProfileResponseOutboundIPsArgs) ToNetworkProfileResponseOutboundIPsPtrOutputWithContext(ctx context.Context) NetworkProfileResponseOutboundIPsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfileResponseOutboundIPsOutput).ToNetworkProfileResponseOutboundIPsPtrOutputWithContext(ctx)
}









type NetworkProfileResponseOutboundIPsPtrInput interface {
	pulumi.Input

	ToNetworkProfileResponseOutboundIPsPtrOutput() NetworkProfileResponseOutboundIPsPtrOutput
	ToNetworkProfileResponseOutboundIPsPtrOutputWithContext(context.Context) NetworkProfileResponseOutboundIPsPtrOutput
}

type networkProfileResponseOutboundIPsPtrType NetworkProfileResponseOutboundIPsArgs

func NetworkProfileResponseOutboundIPsPtr(v *NetworkProfileResponseOutboundIPsArgs) NetworkProfileResponseOutboundIPsPtrInput {
	return (*networkProfileResponseOutboundIPsPtrType)(v)
}

func (*networkProfileResponseOutboundIPsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfileResponseOutboundIPs)(nil)).Elem()
}

func (i *networkProfileResponseOutboundIPsPtrType) ToNetworkProfileResponseOutboundIPsPtrOutput() NetworkProfileResponseOutboundIPsPtrOutput {
	return i.ToNetworkProfileResponseOutboundIPsPtrOutputWithContext(context.Background())
}

func (i *networkProfileResponseOutboundIPsPtrType) ToNetworkProfileResponseOutboundIPsPtrOutputWithContext(ctx context.Context) NetworkProfileResponseOutboundIPsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfileResponseOutboundIPsPtrOutput)
}

type NetworkProfileResponseOutboundIPsOutput struct{ *pulumi.OutputState }

func (NetworkProfileResponseOutboundIPsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkProfileResponseOutboundIPs)(nil)).Elem()
}

func (o NetworkProfileResponseOutboundIPsOutput) ToNetworkProfileResponseOutboundIPsOutput() NetworkProfileResponseOutboundIPsOutput {
	return o
}

func (o NetworkProfileResponseOutboundIPsOutput) ToNetworkProfileResponseOutboundIPsOutputWithContext(ctx context.Context) NetworkProfileResponseOutboundIPsOutput {
	return o
}

func (o NetworkProfileResponseOutboundIPsOutput) ToNetworkProfileResponseOutboundIPsPtrOutput() NetworkProfileResponseOutboundIPsPtrOutput {
	return o.ToNetworkProfileResponseOutboundIPsPtrOutputWithContext(context.Background())
}

func (o NetworkProfileResponseOutboundIPsOutput) ToNetworkProfileResponseOutboundIPsPtrOutputWithContext(ctx context.Context) NetworkProfileResponseOutboundIPsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkProfileResponseOutboundIPs) *NetworkProfileResponseOutboundIPs {
		return &v
	}).(NetworkProfileResponseOutboundIPsPtrOutput)
}

func (o NetworkProfileResponseOutboundIPsOutput) PublicIPs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkProfileResponseOutboundIPs) []string { return v.PublicIPs }).(pulumi.StringArrayOutput)
}

type NetworkProfileResponseOutboundIPsPtrOutput struct{ *pulumi.OutputState }

func (NetworkProfileResponseOutboundIPsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfileResponseOutboundIPs)(nil)).Elem()
}

func (o NetworkProfileResponseOutboundIPsPtrOutput) ToNetworkProfileResponseOutboundIPsPtrOutput() NetworkProfileResponseOutboundIPsPtrOutput {
	return o
}

func (o NetworkProfileResponseOutboundIPsPtrOutput) ToNetworkProfileResponseOutboundIPsPtrOutputWithContext(ctx context.Context) NetworkProfileResponseOutboundIPsPtrOutput {
	return o
}

func (o NetworkProfileResponseOutboundIPsPtrOutput) Elem() NetworkProfileResponseOutboundIPsOutput {
	return o.ApplyT(func(v *NetworkProfileResponseOutboundIPs) NetworkProfileResponseOutboundIPs {
		if v != nil {
			return *v
		}
		var ret NetworkProfileResponseOutboundIPs
		return ret
	}).(NetworkProfileResponseOutboundIPsOutput)
}

func (o NetworkProfileResponseOutboundIPsPtrOutput) PublicIPs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkProfileResponseOutboundIPs) []string {
		if v == nil {
			return nil
		}
		return v.PublicIPs
	}).(pulumi.StringArrayOutput)
}

type PersistentDisk struct {
	MountPath *string `pulumi:"mountPath"`
	SizeInGB  *int    `pulumi:"sizeInGB"`
}





type PersistentDiskInput interface {
	pulumi.Input

	ToPersistentDiskOutput() PersistentDiskOutput
	ToPersistentDiskOutputWithContext(context.Context) PersistentDiskOutput
}

type PersistentDiskArgs struct {
	MountPath pulumi.StringPtrInput `pulumi:"mountPath"`
	SizeInGB  pulumi.IntPtrInput    `pulumi:"sizeInGB"`
}

func (PersistentDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PersistentDisk)(nil)).Elem()
}

func (i PersistentDiskArgs) ToPersistentDiskOutput() PersistentDiskOutput {
	return i.ToPersistentDiskOutputWithContext(context.Background())
}

func (i PersistentDiskArgs) ToPersistentDiskOutputWithContext(ctx context.Context) PersistentDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistentDiskOutput)
}

func (i PersistentDiskArgs) ToPersistentDiskPtrOutput() PersistentDiskPtrOutput {
	return i.ToPersistentDiskPtrOutputWithContext(context.Background())
}

func (i PersistentDiskArgs) ToPersistentDiskPtrOutputWithContext(ctx context.Context) PersistentDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistentDiskOutput).ToPersistentDiskPtrOutputWithContext(ctx)
}









type PersistentDiskPtrInput interface {
	pulumi.Input

	ToPersistentDiskPtrOutput() PersistentDiskPtrOutput
	ToPersistentDiskPtrOutputWithContext(context.Context) PersistentDiskPtrOutput
}

type persistentDiskPtrType PersistentDiskArgs

func PersistentDiskPtr(v *PersistentDiskArgs) PersistentDiskPtrInput {
	return (*persistentDiskPtrType)(v)
}

func (*persistentDiskPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PersistentDisk)(nil)).Elem()
}

func (i *persistentDiskPtrType) ToPersistentDiskPtrOutput() PersistentDiskPtrOutput {
	return i.ToPersistentDiskPtrOutputWithContext(context.Background())
}

func (i *persistentDiskPtrType) ToPersistentDiskPtrOutputWithContext(ctx context.Context) PersistentDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistentDiskPtrOutput)
}

type PersistentDiskOutput struct{ *pulumi.OutputState }

func (PersistentDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PersistentDisk)(nil)).Elem()
}

func (o PersistentDiskOutput) ToPersistentDiskOutput() PersistentDiskOutput {
	return o
}

func (o PersistentDiskOutput) ToPersistentDiskOutputWithContext(ctx context.Context) PersistentDiskOutput {
	return o
}

func (o PersistentDiskOutput) ToPersistentDiskPtrOutput() PersistentDiskPtrOutput {
	return o.ToPersistentDiskPtrOutputWithContext(context.Background())
}

func (o PersistentDiskOutput) ToPersistentDiskPtrOutputWithContext(ctx context.Context) PersistentDiskPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PersistentDisk) *PersistentDisk {
		return &v
	}).(PersistentDiskPtrOutput)
}

func (o PersistentDiskOutput) MountPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PersistentDisk) *string { return v.MountPath }).(pulumi.StringPtrOutput)
}

func (o PersistentDiskOutput) SizeInGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PersistentDisk) *int { return v.SizeInGB }).(pulumi.IntPtrOutput)
}

type PersistentDiskPtrOutput struct{ *pulumi.OutputState }

func (PersistentDiskPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PersistentDisk)(nil)).Elem()
}

func (o PersistentDiskPtrOutput) ToPersistentDiskPtrOutput() PersistentDiskPtrOutput {
	return o
}

func (o PersistentDiskPtrOutput) ToPersistentDiskPtrOutputWithContext(ctx context.Context) PersistentDiskPtrOutput {
	return o
}

func (o PersistentDiskPtrOutput) Elem() PersistentDiskOutput {
	return o.ApplyT(func(v *PersistentDisk) PersistentDisk {
		if v != nil {
			return *v
		}
		var ret PersistentDisk
		return ret
	}).(PersistentDiskOutput)
}

func (o PersistentDiskPtrOutput) MountPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PersistentDisk) *string {
		if v == nil {
			return nil
		}
		return v.MountPath
	}).(pulumi.StringPtrOutput)
}

func (o PersistentDiskPtrOutput) SizeInGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PersistentDisk) *int {
		if v == nil {
			return nil
		}
		return v.SizeInGB
	}).(pulumi.IntPtrOutput)
}

type PersistentDiskResponse struct {
	MountPath *string `pulumi:"mountPath"`
	SizeInGB  *int    `pulumi:"sizeInGB"`
	UsedInGB  int     `pulumi:"usedInGB"`
}





type PersistentDiskResponseInput interface {
	pulumi.Input

	ToPersistentDiskResponseOutput() PersistentDiskResponseOutput
	ToPersistentDiskResponseOutputWithContext(context.Context) PersistentDiskResponseOutput
}

type PersistentDiskResponseArgs struct {
	MountPath pulumi.StringPtrInput `pulumi:"mountPath"`
	SizeInGB  pulumi.IntPtrInput    `pulumi:"sizeInGB"`
	UsedInGB  pulumi.IntInput       `pulumi:"usedInGB"`
}

func (PersistentDiskResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PersistentDiskResponse)(nil)).Elem()
}

func (i PersistentDiskResponseArgs) ToPersistentDiskResponseOutput() PersistentDiskResponseOutput {
	return i.ToPersistentDiskResponseOutputWithContext(context.Background())
}

func (i PersistentDiskResponseArgs) ToPersistentDiskResponseOutputWithContext(ctx context.Context) PersistentDiskResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistentDiskResponseOutput)
}

func (i PersistentDiskResponseArgs) ToPersistentDiskResponsePtrOutput() PersistentDiskResponsePtrOutput {
	return i.ToPersistentDiskResponsePtrOutputWithContext(context.Background())
}

func (i PersistentDiskResponseArgs) ToPersistentDiskResponsePtrOutputWithContext(ctx context.Context) PersistentDiskResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistentDiskResponseOutput).ToPersistentDiskResponsePtrOutputWithContext(ctx)
}









type PersistentDiskResponsePtrInput interface {
	pulumi.Input

	ToPersistentDiskResponsePtrOutput() PersistentDiskResponsePtrOutput
	ToPersistentDiskResponsePtrOutputWithContext(context.Context) PersistentDiskResponsePtrOutput
}

type persistentDiskResponsePtrType PersistentDiskResponseArgs

func PersistentDiskResponsePtr(v *PersistentDiskResponseArgs) PersistentDiskResponsePtrInput {
	return (*persistentDiskResponsePtrType)(v)
}

func (*persistentDiskResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PersistentDiskResponse)(nil)).Elem()
}

func (i *persistentDiskResponsePtrType) ToPersistentDiskResponsePtrOutput() PersistentDiskResponsePtrOutput {
	return i.ToPersistentDiskResponsePtrOutputWithContext(context.Background())
}

func (i *persistentDiskResponsePtrType) ToPersistentDiskResponsePtrOutputWithContext(ctx context.Context) PersistentDiskResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistentDiskResponsePtrOutput)
}

type PersistentDiskResponseOutput struct{ *pulumi.OutputState }

func (PersistentDiskResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PersistentDiskResponse)(nil)).Elem()
}

func (o PersistentDiskResponseOutput) ToPersistentDiskResponseOutput() PersistentDiskResponseOutput {
	return o
}

func (o PersistentDiskResponseOutput) ToPersistentDiskResponseOutputWithContext(ctx context.Context) PersistentDiskResponseOutput {
	return o
}

func (o PersistentDiskResponseOutput) ToPersistentDiskResponsePtrOutput() PersistentDiskResponsePtrOutput {
	return o.ToPersistentDiskResponsePtrOutputWithContext(context.Background())
}

func (o PersistentDiskResponseOutput) ToPersistentDiskResponsePtrOutputWithContext(ctx context.Context) PersistentDiskResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PersistentDiskResponse) *PersistentDiskResponse {
		return &v
	}).(PersistentDiskResponsePtrOutput)
}

func (o PersistentDiskResponseOutput) MountPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PersistentDiskResponse) *string { return v.MountPath }).(pulumi.StringPtrOutput)
}

func (o PersistentDiskResponseOutput) SizeInGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PersistentDiskResponse) *int { return v.SizeInGB }).(pulumi.IntPtrOutput)
}

func (o PersistentDiskResponseOutput) UsedInGB() pulumi.IntOutput {
	return o.ApplyT(func(v PersistentDiskResponse) int { return v.UsedInGB }).(pulumi.IntOutput)
}

type PersistentDiskResponsePtrOutput struct{ *pulumi.OutputState }

func (PersistentDiskResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PersistentDiskResponse)(nil)).Elem()
}

func (o PersistentDiskResponsePtrOutput) ToPersistentDiskResponsePtrOutput() PersistentDiskResponsePtrOutput {
	return o
}

func (o PersistentDiskResponsePtrOutput) ToPersistentDiskResponsePtrOutputWithContext(ctx context.Context) PersistentDiskResponsePtrOutput {
	return o
}

func (o PersistentDiskResponsePtrOutput) Elem() PersistentDiskResponseOutput {
	return o.ApplyT(func(v *PersistentDiskResponse) PersistentDiskResponse {
		if v != nil {
			return *v
		}
		var ret PersistentDiskResponse
		return ret
	}).(PersistentDiskResponseOutput)
}

func (o PersistentDiskResponsePtrOutput) MountPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PersistentDiskResponse) *string {
		if v == nil {
			return nil
		}
		return v.MountPath
	}).(pulumi.StringPtrOutput)
}

func (o PersistentDiskResponsePtrOutput) SizeInGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PersistentDiskResponse) *int {
		if v == nil {
			return nil
		}
		return v.SizeInGB
	}).(pulumi.IntPtrOutput)
}

func (o PersistentDiskResponsePtrOutput) UsedInGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PersistentDiskResponse) *int {
		if v == nil {
			return nil
		}
		return &v.UsedInGB
	}).(pulumi.IntPtrOutput)
}

type RequiredTrafficResponse struct {
	Direction string   `pulumi:"direction"`
	Fqdns     []string `pulumi:"fqdns"`
	Ips       []string `pulumi:"ips"`
	Port      int      `pulumi:"port"`
	Protocol  string   `pulumi:"protocol"`
}





type RequiredTrafficResponseInput interface {
	pulumi.Input

	ToRequiredTrafficResponseOutput() RequiredTrafficResponseOutput
	ToRequiredTrafficResponseOutputWithContext(context.Context) RequiredTrafficResponseOutput
}

type RequiredTrafficResponseArgs struct {
	Direction pulumi.StringInput      `pulumi:"direction"`
	Fqdns     pulumi.StringArrayInput `pulumi:"fqdns"`
	Ips       pulumi.StringArrayInput `pulumi:"ips"`
	Port      pulumi.IntInput         `pulumi:"port"`
	Protocol  pulumi.StringInput      `pulumi:"protocol"`
}

func (RequiredTrafficResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RequiredTrafficResponse)(nil)).Elem()
}

func (i RequiredTrafficResponseArgs) ToRequiredTrafficResponseOutput() RequiredTrafficResponseOutput {
	return i.ToRequiredTrafficResponseOutputWithContext(context.Background())
}

func (i RequiredTrafficResponseArgs) ToRequiredTrafficResponseOutputWithContext(ctx context.Context) RequiredTrafficResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequiredTrafficResponseOutput)
}





type RequiredTrafficResponseArrayInput interface {
	pulumi.Input

	ToRequiredTrafficResponseArrayOutput() RequiredTrafficResponseArrayOutput
	ToRequiredTrafficResponseArrayOutputWithContext(context.Context) RequiredTrafficResponseArrayOutput
}

type RequiredTrafficResponseArray []RequiredTrafficResponseInput

func (RequiredTrafficResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RequiredTrafficResponse)(nil)).Elem()
}

func (i RequiredTrafficResponseArray) ToRequiredTrafficResponseArrayOutput() RequiredTrafficResponseArrayOutput {
	return i.ToRequiredTrafficResponseArrayOutputWithContext(context.Background())
}

func (i RequiredTrafficResponseArray) ToRequiredTrafficResponseArrayOutputWithContext(ctx context.Context) RequiredTrafficResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequiredTrafficResponseArrayOutput)
}

type RequiredTrafficResponseOutput struct{ *pulumi.OutputState }

func (RequiredTrafficResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RequiredTrafficResponse)(nil)).Elem()
}

func (o RequiredTrafficResponseOutput) ToRequiredTrafficResponseOutput() RequiredTrafficResponseOutput {
	return o
}

func (o RequiredTrafficResponseOutput) ToRequiredTrafficResponseOutputWithContext(ctx context.Context) RequiredTrafficResponseOutput {
	return o
}

func (o RequiredTrafficResponseOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v RequiredTrafficResponse) string { return v.Direction }).(pulumi.StringOutput)
}

func (o RequiredTrafficResponseOutput) Fqdns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RequiredTrafficResponse) []string { return v.Fqdns }).(pulumi.StringArrayOutput)
}

func (o RequiredTrafficResponseOutput) Ips() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RequiredTrafficResponse) []string { return v.Ips }).(pulumi.StringArrayOutput)
}

func (o RequiredTrafficResponseOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v RequiredTrafficResponse) int { return v.Port }).(pulumi.IntOutput)
}

func (o RequiredTrafficResponseOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v RequiredTrafficResponse) string { return v.Protocol }).(pulumi.StringOutput)
}

type RequiredTrafficResponseArrayOutput struct{ *pulumi.OutputState }

func (RequiredTrafficResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RequiredTrafficResponse)(nil)).Elem()
}

func (o RequiredTrafficResponseArrayOutput) ToRequiredTrafficResponseArrayOutput() RequiredTrafficResponseArrayOutput {
	return o
}

func (o RequiredTrafficResponseArrayOutput) ToRequiredTrafficResponseArrayOutputWithContext(ctx context.Context) RequiredTrafficResponseArrayOutput {
	return o
}

func (o RequiredTrafficResponseArrayOutput) Index(i pulumi.IntInput) RequiredTrafficResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RequiredTrafficResponse {
		return vs[0].([]RequiredTrafficResponse)[vs[1].(int)]
	}).(RequiredTrafficResponseOutput)
}

type ResourceRequests struct {
	Cpu    *string `pulumi:"cpu"`
	Memory *string `pulumi:"memory"`
}





type ResourceRequestsInput interface {
	pulumi.Input

	ToResourceRequestsOutput() ResourceRequestsOutput
	ToResourceRequestsOutputWithContext(context.Context) ResourceRequestsOutput
}

type ResourceRequestsArgs struct {
	Cpu    pulumi.StringPtrInput `pulumi:"cpu"`
	Memory pulumi.StringPtrInput `pulumi:"memory"`
}

func (ResourceRequestsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceRequests)(nil)).Elem()
}

func (i ResourceRequestsArgs) ToResourceRequestsOutput() ResourceRequestsOutput {
	return i.ToResourceRequestsOutputWithContext(context.Background())
}

func (i ResourceRequestsArgs) ToResourceRequestsOutputWithContext(ctx context.Context) ResourceRequestsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceRequestsOutput)
}

func (i ResourceRequestsArgs) ToResourceRequestsPtrOutput() ResourceRequestsPtrOutput {
	return i.ToResourceRequestsPtrOutputWithContext(context.Background())
}

func (i ResourceRequestsArgs) ToResourceRequestsPtrOutputWithContext(ctx context.Context) ResourceRequestsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceRequestsOutput).ToResourceRequestsPtrOutputWithContext(ctx)
}









type ResourceRequestsPtrInput interface {
	pulumi.Input

	ToResourceRequestsPtrOutput() ResourceRequestsPtrOutput
	ToResourceRequestsPtrOutputWithContext(context.Context) ResourceRequestsPtrOutput
}

type resourceRequestsPtrType ResourceRequestsArgs

func ResourceRequestsPtr(v *ResourceRequestsArgs) ResourceRequestsPtrInput {
	return (*resourceRequestsPtrType)(v)
}

func (*resourceRequestsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceRequests)(nil)).Elem()
}

func (i *resourceRequestsPtrType) ToResourceRequestsPtrOutput() ResourceRequestsPtrOutput {
	return i.ToResourceRequestsPtrOutputWithContext(context.Background())
}

func (i *resourceRequestsPtrType) ToResourceRequestsPtrOutputWithContext(ctx context.Context) ResourceRequestsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceRequestsPtrOutput)
}

type ResourceRequestsOutput struct{ *pulumi.OutputState }

func (ResourceRequestsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceRequests)(nil)).Elem()
}

func (o ResourceRequestsOutput) ToResourceRequestsOutput() ResourceRequestsOutput {
	return o
}

func (o ResourceRequestsOutput) ToResourceRequestsOutputWithContext(ctx context.Context) ResourceRequestsOutput {
	return o
}

func (o ResourceRequestsOutput) ToResourceRequestsPtrOutput() ResourceRequestsPtrOutput {
	return o.ToResourceRequestsPtrOutputWithContext(context.Background())
}

func (o ResourceRequestsOutput) ToResourceRequestsPtrOutputWithContext(ctx context.Context) ResourceRequestsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceRequests) *ResourceRequests {
		return &v
	}).(ResourceRequestsPtrOutput)
}

func (o ResourceRequestsOutput) Cpu() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceRequests) *string { return v.Cpu }).(pulumi.StringPtrOutput)
}

func (o ResourceRequestsOutput) Memory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceRequests) *string { return v.Memory }).(pulumi.StringPtrOutput)
}

type ResourceRequestsPtrOutput struct{ *pulumi.OutputState }

func (ResourceRequestsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceRequests)(nil)).Elem()
}

func (o ResourceRequestsPtrOutput) ToResourceRequestsPtrOutput() ResourceRequestsPtrOutput {
	return o
}

func (o ResourceRequestsPtrOutput) ToResourceRequestsPtrOutputWithContext(ctx context.Context) ResourceRequestsPtrOutput {
	return o
}

func (o ResourceRequestsPtrOutput) Elem() ResourceRequestsOutput {
	return o.ApplyT(func(v *ResourceRequests) ResourceRequests {
		if v != nil {
			return *v
		}
		var ret ResourceRequests
		return ret
	}).(ResourceRequestsOutput)
}

func (o ResourceRequestsPtrOutput) Cpu() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceRequests) *string {
		if v == nil {
			return nil
		}
		return v.Cpu
	}).(pulumi.StringPtrOutput)
}

func (o ResourceRequestsPtrOutput) Memory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceRequests) *string {
		if v == nil {
			return nil
		}
		return v.Memory
	}).(pulumi.StringPtrOutput)
}

type ResourceRequestsResponse struct {
	Cpu    *string `pulumi:"cpu"`
	Memory *string `pulumi:"memory"`
}





type ResourceRequestsResponseInput interface {
	pulumi.Input

	ToResourceRequestsResponseOutput() ResourceRequestsResponseOutput
	ToResourceRequestsResponseOutputWithContext(context.Context) ResourceRequestsResponseOutput
}

type ResourceRequestsResponseArgs struct {
	Cpu    pulumi.StringPtrInput `pulumi:"cpu"`
	Memory pulumi.StringPtrInput `pulumi:"memory"`
}

func (ResourceRequestsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceRequestsResponse)(nil)).Elem()
}

func (i ResourceRequestsResponseArgs) ToResourceRequestsResponseOutput() ResourceRequestsResponseOutput {
	return i.ToResourceRequestsResponseOutputWithContext(context.Background())
}

func (i ResourceRequestsResponseArgs) ToResourceRequestsResponseOutputWithContext(ctx context.Context) ResourceRequestsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceRequestsResponseOutput)
}

func (i ResourceRequestsResponseArgs) ToResourceRequestsResponsePtrOutput() ResourceRequestsResponsePtrOutput {
	return i.ToResourceRequestsResponsePtrOutputWithContext(context.Background())
}

func (i ResourceRequestsResponseArgs) ToResourceRequestsResponsePtrOutputWithContext(ctx context.Context) ResourceRequestsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceRequestsResponseOutput).ToResourceRequestsResponsePtrOutputWithContext(ctx)
}









type ResourceRequestsResponsePtrInput interface {
	pulumi.Input

	ToResourceRequestsResponsePtrOutput() ResourceRequestsResponsePtrOutput
	ToResourceRequestsResponsePtrOutputWithContext(context.Context) ResourceRequestsResponsePtrOutput
}

type resourceRequestsResponsePtrType ResourceRequestsResponseArgs

func ResourceRequestsResponsePtr(v *ResourceRequestsResponseArgs) ResourceRequestsResponsePtrInput {
	return (*resourceRequestsResponsePtrType)(v)
}

func (*resourceRequestsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceRequestsResponse)(nil)).Elem()
}

func (i *resourceRequestsResponsePtrType) ToResourceRequestsResponsePtrOutput() ResourceRequestsResponsePtrOutput {
	return i.ToResourceRequestsResponsePtrOutputWithContext(context.Background())
}

func (i *resourceRequestsResponsePtrType) ToResourceRequestsResponsePtrOutputWithContext(ctx context.Context) ResourceRequestsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceRequestsResponsePtrOutput)
}

type ResourceRequestsResponseOutput struct{ *pulumi.OutputState }

func (ResourceRequestsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceRequestsResponse)(nil)).Elem()
}

func (o ResourceRequestsResponseOutput) ToResourceRequestsResponseOutput() ResourceRequestsResponseOutput {
	return o
}

func (o ResourceRequestsResponseOutput) ToResourceRequestsResponseOutputWithContext(ctx context.Context) ResourceRequestsResponseOutput {
	return o
}

func (o ResourceRequestsResponseOutput) ToResourceRequestsResponsePtrOutput() ResourceRequestsResponsePtrOutput {
	return o.ToResourceRequestsResponsePtrOutputWithContext(context.Background())
}

func (o ResourceRequestsResponseOutput) ToResourceRequestsResponsePtrOutputWithContext(ctx context.Context) ResourceRequestsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceRequestsResponse) *ResourceRequestsResponse {
		return &v
	}).(ResourceRequestsResponsePtrOutput)
}

func (o ResourceRequestsResponseOutput) Cpu() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceRequestsResponse) *string { return v.Cpu }).(pulumi.StringPtrOutput)
}

func (o ResourceRequestsResponseOutput) Memory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceRequestsResponse) *string { return v.Memory }).(pulumi.StringPtrOutput)
}

type ResourceRequestsResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceRequestsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceRequestsResponse)(nil)).Elem()
}

func (o ResourceRequestsResponsePtrOutput) ToResourceRequestsResponsePtrOutput() ResourceRequestsResponsePtrOutput {
	return o
}

func (o ResourceRequestsResponsePtrOutput) ToResourceRequestsResponsePtrOutputWithContext(ctx context.Context) ResourceRequestsResponsePtrOutput {
	return o
}

func (o ResourceRequestsResponsePtrOutput) Elem() ResourceRequestsResponseOutput {
	return o.ApplyT(func(v *ResourceRequestsResponse) ResourceRequestsResponse {
		if v != nil {
			return *v
		}
		var ret ResourceRequestsResponse
		return ret
	}).(ResourceRequestsResponseOutput)
}

func (o ResourceRequestsResponsePtrOutput) Cpu() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceRequestsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Cpu
	}).(pulumi.StringPtrOutput)
}

func (o ResourceRequestsResponsePtrOutput) Memory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceRequestsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Memory
	}).(pulumi.StringPtrOutput)
}

type Sku struct {
	Capacity *int    `pulumi:"capacity"`
	Name     *string `pulumi:"name"`
	Tier     *string `pulumi:"tier"`
}


func (val *Sku) Defaults() *Sku {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Name) {
		name_ := "S0"
		tmp.Name = &name_
	}
	if isZero(tmp.Tier) {
		tier_ := "Standard"
		tmp.Tier = &tier_
	}
	return &tmp
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}









type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
}

type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

func (o SkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Sku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku {
		if v != nil {
			return *v
		}
		var ret Sku
		return ret
	}).(SkuOutput)
}

func (o SkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Sku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Name     *string `pulumi:"name"`
	Tier     *string `pulumi:"tier"`
}


func (val *SkuResponse) Defaults() *SkuResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Name) {
		name_ := "S0"
		tmp.Name = &name_
	}
	if isZero(tmp.Tier) {
		tier_ := "Standard"
		tmp.Tier = &tier_
	}
	return &tmp
}





type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

type SkuResponseArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
}

func (SkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (i SkuResponseArgs) ToSkuResponseOutput() SkuResponseOutput {
	return i.ToSkuResponseOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput)
}

func (i SkuResponseArgs) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput).ToSkuResponsePtrOutputWithContext(ctx)
}









type SkuResponsePtrInput interface {
	pulumi.Input

	ToSkuResponsePtrOutput() SkuResponsePtrOutput
	ToSkuResponsePtrOutputWithContext(context.Context) SkuResponsePtrOutput
}

type skuResponsePtrType SkuResponseArgs

func SkuResponsePtr(v *SkuResponseArgs) SkuResponsePtrInput {
	return (*skuResponsePtrType)(v)
}

func (*skuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponsePtrOutput)
}

type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (o SkuResponseOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuResponse) *SkuResponse {
		return &v
	}).(SkuResponsePtrOutput)
}

func (o SkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse {
		if v != nil {
			return *v
		}
		var ret SkuResponse
		return ret
	}).(SkuResponseOutput)
}

func (o SkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type TemporaryDisk struct {
	MountPath *string `pulumi:"mountPath"`
	SizeInGB  *int    `pulumi:"sizeInGB"`
}


func (val *TemporaryDisk) Defaults() *TemporaryDisk {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MountPath) {
		mountPath_ := "/tmp"
		tmp.MountPath = &mountPath_
	}
	return &tmp
}





type TemporaryDiskInput interface {
	pulumi.Input

	ToTemporaryDiskOutput() TemporaryDiskOutput
	ToTemporaryDiskOutputWithContext(context.Context) TemporaryDiskOutput
}

type TemporaryDiskArgs struct {
	MountPath pulumi.StringPtrInput `pulumi:"mountPath"`
	SizeInGB  pulumi.IntPtrInput    `pulumi:"sizeInGB"`
}

func (TemporaryDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TemporaryDisk)(nil)).Elem()
}

func (i TemporaryDiskArgs) ToTemporaryDiskOutput() TemporaryDiskOutput {
	return i.ToTemporaryDiskOutputWithContext(context.Background())
}

func (i TemporaryDiskArgs) ToTemporaryDiskOutputWithContext(ctx context.Context) TemporaryDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemporaryDiskOutput)
}

func (i TemporaryDiskArgs) ToTemporaryDiskPtrOutput() TemporaryDiskPtrOutput {
	return i.ToTemporaryDiskPtrOutputWithContext(context.Background())
}

func (i TemporaryDiskArgs) ToTemporaryDiskPtrOutputWithContext(ctx context.Context) TemporaryDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemporaryDiskOutput).ToTemporaryDiskPtrOutputWithContext(ctx)
}









type TemporaryDiskPtrInput interface {
	pulumi.Input

	ToTemporaryDiskPtrOutput() TemporaryDiskPtrOutput
	ToTemporaryDiskPtrOutputWithContext(context.Context) TemporaryDiskPtrOutput
}

type temporaryDiskPtrType TemporaryDiskArgs

func TemporaryDiskPtr(v *TemporaryDiskArgs) TemporaryDiskPtrInput {
	return (*temporaryDiskPtrType)(v)
}

func (*temporaryDiskPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TemporaryDisk)(nil)).Elem()
}

func (i *temporaryDiskPtrType) ToTemporaryDiskPtrOutput() TemporaryDiskPtrOutput {
	return i.ToTemporaryDiskPtrOutputWithContext(context.Background())
}

func (i *temporaryDiskPtrType) ToTemporaryDiskPtrOutputWithContext(ctx context.Context) TemporaryDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemporaryDiskPtrOutput)
}

type TemporaryDiskOutput struct{ *pulumi.OutputState }

func (TemporaryDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TemporaryDisk)(nil)).Elem()
}

func (o TemporaryDiskOutput) ToTemporaryDiskOutput() TemporaryDiskOutput {
	return o
}

func (o TemporaryDiskOutput) ToTemporaryDiskOutputWithContext(ctx context.Context) TemporaryDiskOutput {
	return o
}

func (o TemporaryDiskOutput) ToTemporaryDiskPtrOutput() TemporaryDiskPtrOutput {
	return o.ToTemporaryDiskPtrOutputWithContext(context.Background())
}

func (o TemporaryDiskOutput) ToTemporaryDiskPtrOutputWithContext(ctx context.Context) TemporaryDiskPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TemporaryDisk) *TemporaryDisk {
		return &v
	}).(TemporaryDiskPtrOutput)
}

func (o TemporaryDiskOutput) MountPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TemporaryDisk) *string { return v.MountPath }).(pulumi.StringPtrOutput)
}

func (o TemporaryDiskOutput) SizeInGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v TemporaryDisk) *int { return v.SizeInGB }).(pulumi.IntPtrOutput)
}

type TemporaryDiskPtrOutput struct{ *pulumi.OutputState }

func (TemporaryDiskPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TemporaryDisk)(nil)).Elem()
}

func (o TemporaryDiskPtrOutput) ToTemporaryDiskPtrOutput() TemporaryDiskPtrOutput {
	return o
}

func (o TemporaryDiskPtrOutput) ToTemporaryDiskPtrOutputWithContext(ctx context.Context) TemporaryDiskPtrOutput {
	return o
}

func (o TemporaryDiskPtrOutput) Elem() TemporaryDiskOutput {
	return o.ApplyT(func(v *TemporaryDisk) TemporaryDisk {
		if v != nil {
			return *v
		}
		var ret TemporaryDisk
		return ret
	}).(TemporaryDiskOutput)
}

func (o TemporaryDiskPtrOutput) MountPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TemporaryDisk) *string {
		if v == nil {
			return nil
		}
		return v.MountPath
	}).(pulumi.StringPtrOutput)
}

func (o TemporaryDiskPtrOutput) SizeInGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TemporaryDisk) *int {
		if v == nil {
			return nil
		}
		return v.SizeInGB
	}).(pulumi.IntPtrOutput)
}

type TemporaryDiskResponse struct {
	MountPath *string `pulumi:"mountPath"`
	SizeInGB  *int    `pulumi:"sizeInGB"`
}


func (val *TemporaryDiskResponse) Defaults() *TemporaryDiskResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MountPath) {
		mountPath_ := "/tmp"
		tmp.MountPath = &mountPath_
	}
	return &tmp
}





type TemporaryDiskResponseInput interface {
	pulumi.Input

	ToTemporaryDiskResponseOutput() TemporaryDiskResponseOutput
	ToTemporaryDiskResponseOutputWithContext(context.Context) TemporaryDiskResponseOutput
}

type TemporaryDiskResponseArgs struct {
	MountPath pulumi.StringPtrInput `pulumi:"mountPath"`
	SizeInGB  pulumi.IntPtrInput    `pulumi:"sizeInGB"`
}

func (TemporaryDiskResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TemporaryDiskResponse)(nil)).Elem()
}

func (i TemporaryDiskResponseArgs) ToTemporaryDiskResponseOutput() TemporaryDiskResponseOutput {
	return i.ToTemporaryDiskResponseOutputWithContext(context.Background())
}

func (i TemporaryDiskResponseArgs) ToTemporaryDiskResponseOutputWithContext(ctx context.Context) TemporaryDiskResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemporaryDiskResponseOutput)
}

func (i TemporaryDiskResponseArgs) ToTemporaryDiskResponsePtrOutput() TemporaryDiskResponsePtrOutput {
	return i.ToTemporaryDiskResponsePtrOutputWithContext(context.Background())
}

func (i TemporaryDiskResponseArgs) ToTemporaryDiskResponsePtrOutputWithContext(ctx context.Context) TemporaryDiskResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemporaryDiskResponseOutput).ToTemporaryDiskResponsePtrOutputWithContext(ctx)
}









type TemporaryDiskResponsePtrInput interface {
	pulumi.Input

	ToTemporaryDiskResponsePtrOutput() TemporaryDiskResponsePtrOutput
	ToTemporaryDiskResponsePtrOutputWithContext(context.Context) TemporaryDiskResponsePtrOutput
}

type temporaryDiskResponsePtrType TemporaryDiskResponseArgs

func TemporaryDiskResponsePtr(v *TemporaryDiskResponseArgs) TemporaryDiskResponsePtrInput {
	return (*temporaryDiskResponsePtrType)(v)
}

func (*temporaryDiskResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TemporaryDiskResponse)(nil)).Elem()
}

func (i *temporaryDiskResponsePtrType) ToTemporaryDiskResponsePtrOutput() TemporaryDiskResponsePtrOutput {
	return i.ToTemporaryDiskResponsePtrOutputWithContext(context.Background())
}

func (i *temporaryDiskResponsePtrType) ToTemporaryDiskResponsePtrOutputWithContext(ctx context.Context) TemporaryDiskResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemporaryDiskResponsePtrOutput)
}

type TemporaryDiskResponseOutput struct{ *pulumi.OutputState }

func (TemporaryDiskResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TemporaryDiskResponse)(nil)).Elem()
}

func (o TemporaryDiskResponseOutput) ToTemporaryDiskResponseOutput() TemporaryDiskResponseOutput {
	return o
}

func (o TemporaryDiskResponseOutput) ToTemporaryDiskResponseOutputWithContext(ctx context.Context) TemporaryDiskResponseOutput {
	return o
}

func (o TemporaryDiskResponseOutput) ToTemporaryDiskResponsePtrOutput() TemporaryDiskResponsePtrOutput {
	return o.ToTemporaryDiskResponsePtrOutputWithContext(context.Background())
}

func (o TemporaryDiskResponseOutput) ToTemporaryDiskResponsePtrOutputWithContext(ctx context.Context) TemporaryDiskResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TemporaryDiskResponse) *TemporaryDiskResponse {
		return &v
	}).(TemporaryDiskResponsePtrOutput)
}

func (o TemporaryDiskResponseOutput) MountPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TemporaryDiskResponse) *string { return v.MountPath }).(pulumi.StringPtrOutput)
}

func (o TemporaryDiskResponseOutput) SizeInGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v TemporaryDiskResponse) *int { return v.SizeInGB }).(pulumi.IntPtrOutput)
}

type TemporaryDiskResponsePtrOutput struct{ *pulumi.OutputState }

func (TemporaryDiskResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TemporaryDiskResponse)(nil)).Elem()
}

func (o TemporaryDiskResponsePtrOutput) ToTemporaryDiskResponsePtrOutput() TemporaryDiskResponsePtrOutput {
	return o
}

func (o TemporaryDiskResponsePtrOutput) ToTemporaryDiskResponsePtrOutputWithContext(ctx context.Context) TemporaryDiskResponsePtrOutput {
	return o
}

func (o TemporaryDiskResponsePtrOutput) Elem() TemporaryDiskResponseOutput {
	return o.ApplyT(func(v *TemporaryDiskResponse) TemporaryDiskResponse {
		if v != nil {
			return *v
		}
		var ret TemporaryDiskResponse
		return ret
	}).(TemporaryDiskResponseOutput)
}

func (o TemporaryDiskResponsePtrOutput) MountPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TemporaryDiskResponse) *string {
		if v == nil {
			return nil
		}
		return v.MountPath
	}).(pulumi.StringPtrOutput)
}

func (o TemporaryDiskResponsePtrOutput) SizeInGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TemporaryDiskResponse) *int {
		if v == nil {
			return nil
		}
		return v.SizeInGB
	}).(pulumi.IntPtrOutput)
}

type UserSourceInfo struct {
	ArtifactSelector *string          `pulumi:"artifactSelector"`
	CustomContainer  *CustomContainer `pulumi:"customContainer"`
	RelativePath     *string          `pulumi:"relativePath"`
	Type             *string          `pulumi:"type"`
	Version          *string          `pulumi:"version"`
}





type UserSourceInfoInput interface {
	pulumi.Input

	ToUserSourceInfoOutput() UserSourceInfoOutput
	ToUserSourceInfoOutputWithContext(context.Context) UserSourceInfoOutput
}

type UserSourceInfoArgs struct {
	ArtifactSelector pulumi.StringPtrInput   `pulumi:"artifactSelector"`
	CustomContainer  CustomContainerPtrInput `pulumi:"customContainer"`
	RelativePath     pulumi.StringPtrInput   `pulumi:"relativePath"`
	Type             pulumi.StringPtrInput   `pulumi:"type"`
	Version          pulumi.StringPtrInput   `pulumi:"version"`
}

func (UserSourceInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserSourceInfo)(nil)).Elem()
}

func (i UserSourceInfoArgs) ToUserSourceInfoOutput() UserSourceInfoOutput {
	return i.ToUserSourceInfoOutputWithContext(context.Background())
}

func (i UserSourceInfoArgs) ToUserSourceInfoOutputWithContext(ctx context.Context) UserSourceInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserSourceInfoOutput)
}

func (i UserSourceInfoArgs) ToUserSourceInfoPtrOutput() UserSourceInfoPtrOutput {
	return i.ToUserSourceInfoPtrOutputWithContext(context.Background())
}

func (i UserSourceInfoArgs) ToUserSourceInfoPtrOutputWithContext(ctx context.Context) UserSourceInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserSourceInfoOutput).ToUserSourceInfoPtrOutputWithContext(ctx)
}









type UserSourceInfoPtrInput interface {
	pulumi.Input

	ToUserSourceInfoPtrOutput() UserSourceInfoPtrOutput
	ToUserSourceInfoPtrOutputWithContext(context.Context) UserSourceInfoPtrOutput
}

type userSourceInfoPtrType UserSourceInfoArgs

func UserSourceInfoPtr(v *UserSourceInfoArgs) UserSourceInfoPtrInput {
	return (*userSourceInfoPtrType)(v)
}

func (*userSourceInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserSourceInfo)(nil)).Elem()
}

func (i *userSourceInfoPtrType) ToUserSourceInfoPtrOutput() UserSourceInfoPtrOutput {
	return i.ToUserSourceInfoPtrOutputWithContext(context.Background())
}

func (i *userSourceInfoPtrType) ToUserSourceInfoPtrOutputWithContext(ctx context.Context) UserSourceInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserSourceInfoPtrOutput)
}

type UserSourceInfoOutput struct{ *pulumi.OutputState }

func (UserSourceInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserSourceInfo)(nil)).Elem()
}

func (o UserSourceInfoOutput) ToUserSourceInfoOutput() UserSourceInfoOutput {
	return o
}

func (o UserSourceInfoOutput) ToUserSourceInfoOutputWithContext(ctx context.Context) UserSourceInfoOutput {
	return o
}

func (o UserSourceInfoOutput) ToUserSourceInfoPtrOutput() UserSourceInfoPtrOutput {
	return o.ToUserSourceInfoPtrOutputWithContext(context.Background())
}

func (o UserSourceInfoOutput) ToUserSourceInfoPtrOutputWithContext(ctx context.Context) UserSourceInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserSourceInfo) *UserSourceInfo {
		return &v
	}).(UserSourceInfoPtrOutput)
}

func (o UserSourceInfoOutput) ArtifactSelector() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserSourceInfo) *string { return v.ArtifactSelector }).(pulumi.StringPtrOutput)
}

func (o UserSourceInfoOutput) CustomContainer() CustomContainerPtrOutput {
	return o.ApplyT(func(v UserSourceInfo) *CustomContainer { return v.CustomContainer }).(CustomContainerPtrOutput)
}

func (o UserSourceInfoOutput) RelativePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserSourceInfo) *string { return v.RelativePath }).(pulumi.StringPtrOutput)
}

func (o UserSourceInfoOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserSourceInfo) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o UserSourceInfoOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserSourceInfo) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type UserSourceInfoPtrOutput struct{ *pulumi.OutputState }

func (UserSourceInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserSourceInfo)(nil)).Elem()
}

func (o UserSourceInfoPtrOutput) ToUserSourceInfoPtrOutput() UserSourceInfoPtrOutput {
	return o
}

func (o UserSourceInfoPtrOutput) ToUserSourceInfoPtrOutputWithContext(ctx context.Context) UserSourceInfoPtrOutput {
	return o
}

func (o UserSourceInfoPtrOutput) Elem() UserSourceInfoOutput {
	return o.ApplyT(func(v *UserSourceInfo) UserSourceInfo {
		if v != nil {
			return *v
		}
		var ret UserSourceInfo
		return ret
	}).(UserSourceInfoOutput)
}

func (o UserSourceInfoPtrOutput) ArtifactSelector() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserSourceInfo) *string {
		if v == nil {
			return nil
		}
		return v.ArtifactSelector
	}).(pulumi.StringPtrOutput)
}

func (o UserSourceInfoPtrOutput) CustomContainer() CustomContainerPtrOutput {
	return o.ApplyT(func(v *UserSourceInfo) *CustomContainer {
		if v == nil {
			return nil
		}
		return v.CustomContainer
	}).(CustomContainerPtrOutput)
}

func (o UserSourceInfoPtrOutput) RelativePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserSourceInfo) *string {
		if v == nil {
			return nil
		}
		return v.RelativePath
	}).(pulumi.StringPtrOutput)
}

func (o UserSourceInfoPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserSourceInfo) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o UserSourceInfoPtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserSourceInfo) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type UserSourceInfoResponse struct {
	ArtifactSelector *string                  `pulumi:"artifactSelector"`
	CustomContainer  *CustomContainerResponse `pulumi:"customContainer"`
	RelativePath     *string                  `pulumi:"relativePath"`
	Type             *string                  `pulumi:"type"`
	Version          *string                  `pulumi:"version"`
}





type UserSourceInfoResponseInput interface {
	pulumi.Input

	ToUserSourceInfoResponseOutput() UserSourceInfoResponseOutput
	ToUserSourceInfoResponseOutputWithContext(context.Context) UserSourceInfoResponseOutput
}

type UserSourceInfoResponseArgs struct {
	ArtifactSelector pulumi.StringPtrInput           `pulumi:"artifactSelector"`
	CustomContainer  CustomContainerResponsePtrInput `pulumi:"customContainer"`
	RelativePath     pulumi.StringPtrInput           `pulumi:"relativePath"`
	Type             pulumi.StringPtrInput           `pulumi:"type"`
	Version          pulumi.StringPtrInput           `pulumi:"version"`
}

func (UserSourceInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserSourceInfoResponse)(nil)).Elem()
}

func (i UserSourceInfoResponseArgs) ToUserSourceInfoResponseOutput() UserSourceInfoResponseOutput {
	return i.ToUserSourceInfoResponseOutputWithContext(context.Background())
}

func (i UserSourceInfoResponseArgs) ToUserSourceInfoResponseOutputWithContext(ctx context.Context) UserSourceInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserSourceInfoResponseOutput)
}

func (i UserSourceInfoResponseArgs) ToUserSourceInfoResponsePtrOutput() UserSourceInfoResponsePtrOutput {
	return i.ToUserSourceInfoResponsePtrOutputWithContext(context.Background())
}

func (i UserSourceInfoResponseArgs) ToUserSourceInfoResponsePtrOutputWithContext(ctx context.Context) UserSourceInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserSourceInfoResponseOutput).ToUserSourceInfoResponsePtrOutputWithContext(ctx)
}









type UserSourceInfoResponsePtrInput interface {
	pulumi.Input

	ToUserSourceInfoResponsePtrOutput() UserSourceInfoResponsePtrOutput
	ToUserSourceInfoResponsePtrOutputWithContext(context.Context) UserSourceInfoResponsePtrOutput
}

type userSourceInfoResponsePtrType UserSourceInfoResponseArgs

func UserSourceInfoResponsePtr(v *UserSourceInfoResponseArgs) UserSourceInfoResponsePtrInput {
	return (*userSourceInfoResponsePtrType)(v)
}

func (*userSourceInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserSourceInfoResponse)(nil)).Elem()
}

func (i *userSourceInfoResponsePtrType) ToUserSourceInfoResponsePtrOutput() UserSourceInfoResponsePtrOutput {
	return i.ToUserSourceInfoResponsePtrOutputWithContext(context.Background())
}

func (i *userSourceInfoResponsePtrType) ToUserSourceInfoResponsePtrOutputWithContext(ctx context.Context) UserSourceInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserSourceInfoResponsePtrOutput)
}

type UserSourceInfoResponseOutput struct{ *pulumi.OutputState }

func (UserSourceInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserSourceInfoResponse)(nil)).Elem()
}

func (o UserSourceInfoResponseOutput) ToUserSourceInfoResponseOutput() UserSourceInfoResponseOutput {
	return o
}

func (o UserSourceInfoResponseOutput) ToUserSourceInfoResponseOutputWithContext(ctx context.Context) UserSourceInfoResponseOutput {
	return o
}

func (o UserSourceInfoResponseOutput) ToUserSourceInfoResponsePtrOutput() UserSourceInfoResponsePtrOutput {
	return o.ToUserSourceInfoResponsePtrOutputWithContext(context.Background())
}

func (o UserSourceInfoResponseOutput) ToUserSourceInfoResponsePtrOutputWithContext(ctx context.Context) UserSourceInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserSourceInfoResponse) *UserSourceInfoResponse {
		return &v
	}).(UserSourceInfoResponsePtrOutput)
}

func (o UserSourceInfoResponseOutput) ArtifactSelector() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserSourceInfoResponse) *string { return v.ArtifactSelector }).(pulumi.StringPtrOutput)
}

func (o UserSourceInfoResponseOutput) CustomContainer() CustomContainerResponsePtrOutput {
	return o.ApplyT(func(v UserSourceInfoResponse) *CustomContainerResponse { return v.CustomContainer }).(CustomContainerResponsePtrOutput)
}

func (o UserSourceInfoResponseOutput) RelativePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserSourceInfoResponse) *string { return v.RelativePath }).(pulumi.StringPtrOutput)
}

func (o UserSourceInfoResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserSourceInfoResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o UserSourceInfoResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserSourceInfoResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type UserSourceInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (UserSourceInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserSourceInfoResponse)(nil)).Elem()
}

func (o UserSourceInfoResponsePtrOutput) ToUserSourceInfoResponsePtrOutput() UserSourceInfoResponsePtrOutput {
	return o
}

func (o UserSourceInfoResponsePtrOutput) ToUserSourceInfoResponsePtrOutputWithContext(ctx context.Context) UserSourceInfoResponsePtrOutput {
	return o
}

func (o UserSourceInfoResponsePtrOutput) Elem() UserSourceInfoResponseOutput {
	return o.ApplyT(func(v *UserSourceInfoResponse) UserSourceInfoResponse {
		if v != nil {
			return *v
		}
		var ret UserSourceInfoResponse
		return ret
	}).(UserSourceInfoResponseOutput)
}

func (o UserSourceInfoResponsePtrOutput) ArtifactSelector() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserSourceInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.ArtifactSelector
	}).(pulumi.StringPtrOutput)
}

func (o UserSourceInfoResponsePtrOutput) CustomContainer() CustomContainerResponsePtrOutput {
	return o.ApplyT(func(v *UserSourceInfoResponse) *CustomContainerResponse {
		if v == nil {
			return nil
		}
		return v.CustomContainer
	}).(CustomContainerResponsePtrOutput)
}

func (o UserSourceInfoResponsePtrOutput) RelativePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserSourceInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.RelativePath
	}).(pulumi.StringPtrOutput)
}

func (o UserSourceInfoResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserSourceInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o UserSourceInfoResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserSourceInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AppResourcePropertiesOutput{})
	pulumi.RegisterOutputType(AppResourcePropertiesPtrOutput{})
	pulumi.RegisterOutputType(AppResourcePropertiesResponseOutput{})
	pulumi.RegisterOutputType(AppResourcePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(BindingResourcePropertiesOutput{})
	pulumi.RegisterOutputType(BindingResourcePropertiesPtrOutput{})
	pulumi.RegisterOutputType(BindingResourcePropertiesResponseOutput{})
	pulumi.RegisterOutputType(BindingResourcePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(CertificatePropertiesOutput{})
	pulumi.RegisterOutputType(CertificatePropertiesPtrOutput{})
	pulumi.RegisterOutputType(CertificatePropertiesResponseOutput{})
	pulumi.RegisterOutputType(CertificatePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ClusterResourcePropertiesOutput{})
	pulumi.RegisterOutputType(ClusterResourcePropertiesPtrOutput{})
	pulumi.RegisterOutputType(ClusterResourcePropertiesResponseOutput{})
	pulumi.RegisterOutputType(ClusterResourcePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(CustomContainerOutput{})
	pulumi.RegisterOutputType(CustomContainerPtrOutput{})
	pulumi.RegisterOutputType(CustomContainerResponseOutput{})
	pulumi.RegisterOutputType(CustomContainerResponsePtrOutput{})
	pulumi.RegisterOutputType(CustomDomainPropertiesOutput{})
	pulumi.RegisterOutputType(CustomDomainPropertiesPtrOutput{})
	pulumi.RegisterOutputType(CustomDomainPropertiesResponseOutput{})
	pulumi.RegisterOutputType(CustomDomainPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(DeploymentInstanceResponseOutput{})
	pulumi.RegisterOutputType(DeploymentInstanceResponseArrayOutput{})
	pulumi.RegisterOutputType(DeploymentResourcePropertiesOutput{})
	pulumi.RegisterOutputType(DeploymentResourcePropertiesPtrOutput{})
	pulumi.RegisterOutputType(DeploymentResourcePropertiesResponseOutput{})
	pulumi.RegisterOutputType(DeploymentResourcePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(DeploymentSettingsOutput{})
	pulumi.RegisterOutputType(DeploymentSettingsPtrOutput{})
	pulumi.RegisterOutputType(DeploymentSettingsResponseOutput{})
	pulumi.RegisterOutputType(DeploymentSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(ImageRegistryCredentialOutput{})
	pulumi.RegisterOutputType(ImageRegistryCredentialPtrOutput{})
	pulumi.RegisterOutputType(ImageRegistryCredentialResponseOutput{})
	pulumi.RegisterOutputType(ImageRegistryCredentialResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedIdentityPropertiesOutput{})
	pulumi.RegisterOutputType(ManagedIdentityPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ManagedIdentityPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ManagedIdentityPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(NetworkProfileOutput{})
	pulumi.RegisterOutputType(NetworkProfilePtrOutput{})
	pulumi.RegisterOutputType(NetworkProfileResponseOutput{})
	pulumi.RegisterOutputType(NetworkProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(NetworkProfileResponseOutboundIPsOutput{})
	pulumi.RegisterOutputType(NetworkProfileResponseOutboundIPsPtrOutput{})
	pulumi.RegisterOutputType(PersistentDiskOutput{})
	pulumi.RegisterOutputType(PersistentDiskPtrOutput{})
	pulumi.RegisterOutputType(PersistentDiskResponseOutput{})
	pulumi.RegisterOutputType(PersistentDiskResponsePtrOutput{})
	pulumi.RegisterOutputType(RequiredTrafficResponseOutput{})
	pulumi.RegisterOutputType(RequiredTrafficResponseArrayOutput{})
	pulumi.RegisterOutputType(ResourceRequestsOutput{})
	pulumi.RegisterOutputType(ResourceRequestsPtrOutput{})
	pulumi.RegisterOutputType(ResourceRequestsResponseOutput{})
	pulumi.RegisterOutputType(ResourceRequestsResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(TemporaryDiskOutput{})
	pulumi.RegisterOutputType(TemporaryDiskPtrOutput{})
	pulumi.RegisterOutputType(TemporaryDiskResponseOutput{})
	pulumi.RegisterOutputType(TemporaryDiskResponsePtrOutput{})
	pulumi.RegisterOutputType(UserSourceInfoOutput{})
	pulumi.RegisterOutputType(UserSourceInfoPtrOutput{})
	pulumi.RegisterOutputType(UserSourceInfoResponseOutput{})
	pulumi.RegisterOutputType(UserSourceInfoResponsePtrOutput{})
}
