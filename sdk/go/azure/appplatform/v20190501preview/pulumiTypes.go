


package v20190501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AppResourceProperties struct {
	ActiveDeploymentName *string         `pulumi:"activeDeploymentName"`
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
	ConfigServerProperties *ConfigServerProperties `pulumi:"configServerProperties"`
	NetworkProfile         *NetworkProfile         `pulumi:"networkProfile"`
	Trace                  *TraceProperties        `pulumi:"trace"`
}





type ClusterResourcePropertiesInput interface {
	pulumi.Input

	ToClusterResourcePropertiesOutput() ClusterResourcePropertiesOutput
	ToClusterResourcePropertiesOutputWithContext(context.Context) ClusterResourcePropertiesOutput
}

type ClusterResourcePropertiesArgs struct {
	ConfigServerProperties ConfigServerPropertiesPtrInput `pulumi:"configServerProperties"`
	NetworkProfile         NetworkProfilePtrInput         `pulumi:"networkProfile"`
	Trace                  TracePropertiesPtrInput        `pulumi:"trace"`
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

func (o ClusterResourcePropertiesOutput) ConfigServerProperties() ConfigServerPropertiesPtrOutput {
	return o.ApplyT(func(v ClusterResourceProperties) *ConfigServerProperties { return v.ConfigServerProperties }).(ConfigServerPropertiesPtrOutput)
}

func (o ClusterResourcePropertiesOutput) NetworkProfile() NetworkProfilePtrOutput {
	return o.ApplyT(func(v ClusterResourceProperties) *NetworkProfile { return v.NetworkProfile }).(NetworkProfilePtrOutput)
}

func (o ClusterResourcePropertiesOutput) Trace() TracePropertiesPtrOutput {
	return o.ApplyT(func(v ClusterResourceProperties) *TraceProperties { return v.Trace }).(TracePropertiesPtrOutput)
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

func (o ClusterResourcePropertiesPtrOutput) ConfigServerProperties() ConfigServerPropertiesPtrOutput {
	return o.ApplyT(func(v *ClusterResourceProperties) *ConfigServerProperties {
		if v == nil {
			return nil
		}
		return v.ConfigServerProperties
	}).(ConfigServerPropertiesPtrOutput)
}

func (o ClusterResourcePropertiesPtrOutput) NetworkProfile() NetworkProfilePtrOutput {
	return o.ApplyT(func(v *ClusterResourceProperties) *NetworkProfile {
		if v == nil {
			return nil
		}
		return v.NetworkProfile
	}).(NetworkProfilePtrOutput)
}

func (o ClusterResourcePropertiesPtrOutput) Trace() TracePropertiesPtrOutput {
	return o.ApplyT(func(v *ClusterResourceProperties) *TraceProperties {
		if v == nil {
			return nil
		}
		return v.Trace
	}).(TracePropertiesPtrOutput)
}

type ClusterResourcePropertiesResponse struct {
	ConfigServerProperties *ConfigServerPropertiesResponse `pulumi:"configServerProperties"`
	NetworkProfile         *NetworkProfileResponse         `pulumi:"networkProfile"`
	ProvisioningState      string                          `pulumi:"provisioningState"`
	ServiceId              string                          `pulumi:"serviceId"`
	Trace                  *TracePropertiesResponse        `pulumi:"trace"`
	Version                int                             `pulumi:"version"`
}





type ClusterResourcePropertiesResponseInput interface {
	pulumi.Input

	ToClusterResourcePropertiesResponseOutput() ClusterResourcePropertiesResponseOutput
	ToClusterResourcePropertiesResponseOutputWithContext(context.Context) ClusterResourcePropertiesResponseOutput
}

type ClusterResourcePropertiesResponseArgs struct {
	ConfigServerProperties ConfigServerPropertiesResponsePtrInput `pulumi:"configServerProperties"`
	NetworkProfile         NetworkProfileResponsePtrInput         `pulumi:"networkProfile"`
	ProvisioningState      pulumi.StringInput                     `pulumi:"provisioningState"`
	ServiceId              pulumi.StringInput                     `pulumi:"serviceId"`
	Trace                  TracePropertiesResponsePtrInput        `pulumi:"trace"`
	Version                pulumi.IntInput                        `pulumi:"version"`
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

func (o ClusterResourcePropertiesResponseOutput) ConfigServerProperties() ConfigServerPropertiesResponsePtrOutput {
	return o.ApplyT(func(v ClusterResourcePropertiesResponse) *ConfigServerPropertiesResponse {
		return v.ConfigServerProperties
	}).(ConfigServerPropertiesResponsePtrOutput)
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

func (o ClusterResourcePropertiesResponseOutput) Trace() TracePropertiesResponsePtrOutput {
	return o.ApplyT(func(v ClusterResourcePropertiesResponse) *TracePropertiesResponse { return v.Trace }).(TracePropertiesResponsePtrOutput)
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

func (o ClusterResourcePropertiesResponsePtrOutput) ConfigServerProperties() ConfigServerPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *ClusterResourcePropertiesResponse) *ConfigServerPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.ConfigServerProperties
	}).(ConfigServerPropertiesResponsePtrOutput)
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

func (o ClusterResourcePropertiesResponsePtrOutput) Trace() TracePropertiesResponsePtrOutput {
	return o.ApplyT(func(v *ClusterResourcePropertiesResponse) *TracePropertiesResponse {
		if v == nil {
			return nil
		}
		return v.Trace
	}).(TracePropertiesResponsePtrOutput)
}

func (o ClusterResourcePropertiesResponsePtrOutput) Version() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ClusterResourcePropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return &v.Version
	}).(pulumi.IntPtrOutput)
}

type ConfigServerGitProperty struct {
	HostKey               *string                `pulumi:"hostKey"`
	HostKeyAlgorithm      *string                `pulumi:"hostKeyAlgorithm"`
	Label                 *string                `pulumi:"label"`
	Password              *string                `pulumi:"password"`
	PrivateKey            *string                `pulumi:"privateKey"`
	Repositories          []GitPatternRepository `pulumi:"repositories"`
	SearchPaths           []string               `pulumi:"searchPaths"`
	StrictHostKeyChecking *bool                  `pulumi:"strictHostKeyChecking"`
	Uri                   string                 `pulumi:"uri"`
	Username              *string                `pulumi:"username"`
}





type ConfigServerGitPropertyInput interface {
	pulumi.Input

	ToConfigServerGitPropertyOutput() ConfigServerGitPropertyOutput
	ToConfigServerGitPropertyOutputWithContext(context.Context) ConfigServerGitPropertyOutput
}

type ConfigServerGitPropertyArgs struct {
	HostKey               pulumi.StringPtrInput          `pulumi:"hostKey"`
	HostKeyAlgorithm      pulumi.StringPtrInput          `pulumi:"hostKeyAlgorithm"`
	Label                 pulumi.StringPtrInput          `pulumi:"label"`
	Password              pulumi.StringPtrInput          `pulumi:"password"`
	PrivateKey            pulumi.StringPtrInput          `pulumi:"privateKey"`
	Repositories          GitPatternRepositoryArrayInput `pulumi:"repositories"`
	SearchPaths           pulumi.StringArrayInput        `pulumi:"searchPaths"`
	StrictHostKeyChecking pulumi.BoolPtrInput            `pulumi:"strictHostKeyChecking"`
	Uri                   pulumi.StringInput             `pulumi:"uri"`
	Username              pulumi.StringPtrInput          `pulumi:"username"`
}

func (ConfigServerGitPropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigServerGitProperty)(nil)).Elem()
}

func (i ConfigServerGitPropertyArgs) ToConfigServerGitPropertyOutput() ConfigServerGitPropertyOutput {
	return i.ToConfigServerGitPropertyOutputWithContext(context.Background())
}

func (i ConfigServerGitPropertyArgs) ToConfigServerGitPropertyOutputWithContext(ctx context.Context) ConfigServerGitPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigServerGitPropertyOutput)
}

func (i ConfigServerGitPropertyArgs) ToConfigServerGitPropertyPtrOutput() ConfigServerGitPropertyPtrOutput {
	return i.ToConfigServerGitPropertyPtrOutputWithContext(context.Background())
}

func (i ConfigServerGitPropertyArgs) ToConfigServerGitPropertyPtrOutputWithContext(ctx context.Context) ConfigServerGitPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigServerGitPropertyOutput).ToConfigServerGitPropertyPtrOutputWithContext(ctx)
}









type ConfigServerGitPropertyPtrInput interface {
	pulumi.Input

	ToConfigServerGitPropertyPtrOutput() ConfigServerGitPropertyPtrOutput
	ToConfigServerGitPropertyPtrOutputWithContext(context.Context) ConfigServerGitPropertyPtrOutput
}

type configServerGitPropertyPtrType ConfigServerGitPropertyArgs

func ConfigServerGitPropertyPtr(v *ConfigServerGitPropertyArgs) ConfigServerGitPropertyPtrInput {
	return (*configServerGitPropertyPtrType)(v)
}

func (*configServerGitPropertyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigServerGitProperty)(nil)).Elem()
}

func (i *configServerGitPropertyPtrType) ToConfigServerGitPropertyPtrOutput() ConfigServerGitPropertyPtrOutput {
	return i.ToConfigServerGitPropertyPtrOutputWithContext(context.Background())
}

func (i *configServerGitPropertyPtrType) ToConfigServerGitPropertyPtrOutputWithContext(ctx context.Context) ConfigServerGitPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigServerGitPropertyPtrOutput)
}

type ConfigServerGitPropertyOutput struct{ *pulumi.OutputState }

func (ConfigServerGitPropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigServerGitProperty)(nil)).Elem()
}

func (o ConfigServerGitPropertyOutput) ToConfigServerGitPropertyOutput() ConfigServerGitPropertyOutput {
	return o
}

func (o ConfigServerGitPropertyOutput) ToConfigServerGitPropertyOutputWithContext(ctx context.Context) ConfigServerGitPropertyOutput {
	return o
}

func (o ConfigServerGitPropertyOutput) ToConfigServerGitPropertyPtrOutput() ConfigServerGitPropertyPtrOutput {
	return o.ToConfigServerGitPropertyPtrOutputWithContext(context.Background())
}

func (o ConfigServerGitPropertyOutput) ToConfigServerGitPropertyPtrOutputWithContext(ctx context.Context) ConfigServerGitPropertyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConfigServerGitProperty) *ConfigServerGitProperty {
		return &v
	}).(ConfigServerGitPropertyPtrOutput)
}

func (o ConfigServerGitPropertyOutput) HostKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigServerGitProperty) *string { return v.HostKey }).(pulumi.StringPtrOutput)
}

func (o ConfigServerGitPropertyOutput) HostKeyAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigServerGitProperty) *string { return v.HostKeyAlgorithm }).(pulumi.StringPtrOutput)
}

func (o ConfigServerGitPropertyOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigServerGitProperty) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o ConfigServerGitPropertyOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigServerGitProperty) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o ConfigServerGitPropertyOutput) PrivateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigServerGitProperty) *string { return v.PrivateKey }).(pulumi.StringPtrOutput)
}

func (o ConfigServerGitPropertyOutput) Repositories() GitPatternRepositoryArrayOutput {
	return o.ApplyT(func(v ConfigServerGitProperty) []GitPatternRepository { return v.Repositories }).(GitPatternRepositoryArrayOutput)
}

func (o ConfigServerGitPropertyOutput) SearchPaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ConfigServerGitProperty) []string { return v.SearchPaths }).(pulumi.StringArrayOutput)
}

func (o ConfigServerGitPropertyOutput) StrictHostKeyChecking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ConfigServerGitProperty) *bool { return v.StrictHostKeyChecking }).(pulumi.BoolPtrOutput)
}

func (o ConfigServerGitPropertyOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigServerGitProperty) string { return v.Uri }).(pulumi.StringOutput)
}

func (o ConfigServerGitPropertyOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigServerGitProperty) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type ConfigServerGitPropertyPtrOutput struct{ *pulumi.OutputState }

func (ConfigServerGitPropertyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigServerGitProperty)(nil)).Elem()
}

func (o ConfigServerGitPropertyPtrOutput) ToConfigServerGitPropertyPtrOutput() ConfigServerGitPropertyPtrOutput {
	return o
}

func (o ConfigServerGitPropertyPtrOutput) ToConfigServerGitPropertyPtrOutputWithContext(ctx context.Context) ConfigServerGitPropertyPtrOutput {
	return o
}

func (o ConfigServerGitPropertyPtrOutput) Elem() ConfigServerGitPropertyOutput {
	return o.ApplyT(func(v *ConfigServerGitProperty) ConfigServerGitProperty {
		if v != nil {
			return *v
		}
		var ret ConfigServerGitProperty
		return ret
	}).(ConfigServerGitPropertyOutput)
}

func (o ConfigServerGitPropertyPtrOutput) HostKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigServerGitProperty) *string {
		if v == nil {
			return nil
		}
		return v.HostKey
	}).(pulumi.StringPtrOutput)
}

func (o ConfigServerGitPropertyPtrOutput) HostKeyAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigServerGitProperty) *string {
		if v == nil {
			return nil
		}
		return v.HostKeyAlgorithm
	}).(pulumi.StringPtrOutput)
}

func (o ConfigServerGitPropertyPtrOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigServerGitProperty) *string {
		if v == nil {
			return nil
		}
		return v.Label
	}).(pulumi.StringPtrOutput)
}

func (o ConfigServerGitPropertyPtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigServerGitProperty) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

func (o ConfigServerGitPropertyPtrOutput) PrivateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigServerGitProperty) *string {
		if v == nil {
			return nil
		}
		return v.PrivateKey
	}).(pulumi.StringPtrOutput)
}

func (o ConfigServerGitPropertyPtrOutput) Repositories() GitPatternRepositoryArrayOutput {
	return o.ApplyT(func(v *ConfigServerGitProperty) []GitPatternRepository {
		if v == nil {
			return nil
		}
		return v.Repositories
	}).(GitPatternRepositoryArrayOutput)
}

func (o ConfigServerGitPropertyPtrOutput) SearchPaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ConfigServerGitProperty) []string {
		if v == nil {
			return nil
		}
		return v.SearchPaths
	}).(pulumi.StringArrayOutput)
}

func (o ConfigServerGitPropertyPtrOutput) StrictHostKeyChecking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ConfigServerGitProperty) *bool {
		if v == nil {
			return nil
		}
		return v.StrictHostKeyChecking
	}).(pulumi.BoolPtrOutput)
}

func (o ConfigServerGitPropertyPtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigServerGitProperty) *string {
		if v == nil {
			return nil
		}
		return &v.Uri
	}).(pulumi.StringPtrOutput)
}

func (o ConfigServerGitPropertyPtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigServerGitProperty) *string {
		if v == nil {
			return nil
		}
		return v.Username
	}).(pulumi.StringPtrOutput)
}

type ConfigServerGitPropertyResponse struct {
	HostKey               *string                        `pulumi:"hostKey"`
	HostKeyAlgorithm      *string                        `pulumi:"hostKeyAlgorithm"`
	Label                 *string                        `pulumi:"label"`
	Password              *string                        `pulumi:"password"`
	PrivateKey            *string                        `pulumi:"privateKey"`
	Repositories          []GitPatternRepositoryResponse `pulumi:"repositories"`
	SearchPaths           []string                       `pulumi:"searchPaths"`
	StrictHostKeyChecking *bool                          `pulumi:"strictHostKeyChecking"`
	Uri                   string                         `pulumi:"uri"`
	Username              *string                        `pulumi:"username"`
}





type ConfigServerGitPropertyResponseInput interface {
	pulumi.Input

	ToConfigServerGitPropertyResponseOutput() ConfigServerGitPropertyResponseOutput
	ToConfigServerGitPropertyResponseOutputWithContext(context.Context) ConfigServerGitPropertyResponseOutput
}

type ConfigServerGitPropertyResponseArgs struct {
	HostKey               pulumi.StringPtrInput                  `pulumi:"hostKey"`
	HostKeyAlgorithm      pulumi.StringPtrInput                  `pulumi:"hostKeyAlgorithm"`
	Label                 pulumi.StringPtrInput                  `pulumi:"label"`
	Password              pulumi.StringPtrInput                  `pulumi:"password"`
	PrivateKey            pulumi.StringPtrInput                  `pulumi:"privateKey"`
	Repositories          GitPatternRepositoryResponseArrayInput `pulumi:"repositories"`
	SearchPaths           pulumi.StringArrayInput                `pulumi:"searchPaths"`
	StrictHostKeyChecking pulumi.BoolPtrInput                    `pulumi:"strictHostKeyChecking"`
	Uri                   pulumi.StringInput                     `pulumi:"uri"`
	Username              pulumi.StringPtrInput                  `pulumi:"username"`
}

func (ConfigServerGitPropertyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigServerGitPropertyResponse)(nil)).Elem()
}

func (i ConfigServerGitPropertyResponseArgs) ToConfigServerGitPropertyResponseOutput() ConfigServerGitPropertyResponseOutput {
	return i.ToConfigServerGitPropertyResponseOutputWithContext(context.Background())
}

func (i ConfigServerGitPropertyResponseArgs) ToConfigServerGitPropertyResponseOutputWithContext(ctx context.Context) ConfigServerGitPropertyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigServerGitPropertyResponseOutput)
}

func (i ConfigServerGitPropertyResponseArgs) ToConfigServerGitPropertyResponsePtrOutput() ConfigServerGitPropertyResponsePtrOutput {
	return i.ToConfigServerGitPropertyResponsePtrOutputWithContext(context.Background())
}

func (i ConfigServerGitPropertyResponseArgs) ToConfigServerGitPropertyResponsePtrOutputWithContext(ctx context.Context) ConfigServerGitPropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigServerGitPropertyResponseOutput).ToConfigServerGitPropertyResponsePtrOutputWithContext(ctx)
}









type ConfigServerGitPropertyResponsePtrInput interface {
	pulumi.Input

	ToConfigServerGitPropertyResponsePtrOutput() ConfigServerGitPropertyResponsePtrOutput
	ToConfigServerGitPropertyResponsePtrOutputWithContext(context.Context) ConfigServerGitPropertyResponsePtrOutput
}

type configServerGitPropertyResponsePtrType ConfigServerGitPropertyResponseArgs

func ConfigServerGitPropertyResponsePtr(v *ConfigServerGitPropertyResponseArgs) ConfigServerGitPropertyResponsePtrInput {
	return (*configServerGitPropertyResponsePtrType)(v)
}

func (*configServerGitPropertyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigServerGitPropertyResponse)(nil)).Elem()
}

func (i *configServerGitPropertyResponsePtrType) ToConfigServerGitPropertyResponsePtrOutput() ConfigServerGitPropertyResponsePtrOutput {
	return i.ToConfigServerGitPropertyResponsePtrOutputWithContext(context.Background())
}

func (i *configServerGitPropertyResponsePtrType) ToConfigServerGitPropertyResponsePtrOutputWithContext(ctx context.Context) ConfigServerGitPropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigServerGitPropertyResponsePtrOutput)
}

type ConfigServerGitPropertyResponseOutput struct{ *pulumi.OutputState }

func (ConfigServerGitPropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigServerGitPropertyResponse)(nil)).Elem()
}

func (o ConfigServerGitPropertyResponseOutput) ToConfigServerGitPropertyResponseOutput() ConfigServerGitPropertyResponseOutput {
	return o
}

func (o ConfigServerGitPropertyResponseOutput) ToConfigServerGitPropertyResponseOutputWithContext(ctx context.Context) ConfigServerGitPropertyResponseOutput {
	return o
}

func (o ConfigServerGitPropertyResponseOutput) ToConfigServerGitPropertyResponsePtrOutput() ConfigServerGitPropertyResponsePtrOutput {
	return o.ToConfigServerGitPropertyResponsePtrOutputWithContext(context.Background())
}

func (o ConfigServerGitPropertyResponseOutput) ToConfigServerGitPropertyResponsePtrOutputWithContext(ctx context.Context) ConfigServerGitPropertyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConfigServerGitPropertyResponse) *ConfigServerGitPropertyResponse {
		return &v
	}).(ConfigServerGitPropertyResponsePtrOutput)
}

func (o ConfigServerGitPropertyResponseOutput) HostKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigServerGitPropertyResponse) *string { return v.HostKey }).(pulumi.StringPtrOutput)
}

func (o ConfigServerGitPropertyResponseOutput) HostKeyAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigServerGitPropertyResponse) *string { return v.HostKeyAlgorithm }).(pulumi.StringPtrOutput)
}

func (o ConfigServerGitPropertyResponseOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigServerGitPropertyResponse) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o ConfigServerGitPropertyResponseOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigServerGitPropertyResponse) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o ConfigServerGitPropertyResponseOutput) PrivateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigServerGitPropertyResponse) *string { return v.PrivateKey }).(pulumi.StringPtrOutput)
}

func (o ConfigServerGitPropertyResponseOutput) Repositories() GitPatternRepositoryResponseArrayOutput {
	return o.ApplyT(func(v ConfigServerGitPropertyResponse) []GitPatternRepositoryResponse { return v.Repositories }).(GitPatternRepositoryResponseArrayOutput)
}

func (o ConfigServerGitPropertyResponseOutput) SearchPaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ConfigServerGitPropertyResponse) []string { return v.SearchPaths }).(pulumi.StringArrayOutput)
}

func (o ConfigServerGitPropertyResponseOutput) StrictHostKeyChecking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ConfigServerGitPropertyResponse) *bool { return v.StrictHostKeyChecking }).(pulumi.BoolPtrOutput)
}

func (o ConfigServerGitPropertyResponseOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigServerGitPropertyResponse) string { return v.Uri }).(pulumi.StringOutput)
}

func (o ConfigServerGitPropertyResponseOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigServerGitPropertyResponse) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type ConfigServerGitPropertyResponsePtrOutput struct{ *pulumi.OutputState }

func (ConfigServerGitPropertyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigServerGitPropertyResponse)(nil)).Elem()
}

func (o ConfigServerGitPropertyResponsePtrOutput) ToConfigServerGitPropertyResponsePtrOutput() ConfigServerGitPropertyResponsePtrOutput {
	return o
}

func (o ConfigServerGitPropertyResponsePtrOutput) ToConfigServerGitPropertyResponsePtrOutputWithContext(ctx context.Context) ConfigServerGitPropertyResponsePtrOutput {
	return o
}

func (o ConfigServerGitPropertyResponsePtrOutput) Elem() ConfigServerGitPropertyResponseOutput {
	return o.ApplyT(func(v *ConfigServerGitPropertyResponse) ConfigServerGitPropertyResponse {
		if v != nil {
			return *v
		}
		var ret ConfigServerGitPropertyResponse
		return ret
	}).(ConfigServerGitPropertyResponseOutput)
}

func (o ConfigServerGitPropertyResponsePtrOutput) HostKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigServerGitPropertyResponse) *string {
		if v == nil {
			return nil
		}
		return v.HostKey
	}).(pulumi.StringPtrOutput)
}

func (o ConfigServerGitPropertyResponsePtrOutput) HostKeyAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigServerGitPropertyResponse) *string {
		if v == nil {
			return nil
		}
		return v.HostKeyAlgorithm
	}).(pulumi.StringPtrOutput)
}

func (o ConfigServerGitPropertyResponsePtrOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigServerGitPropertyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Label
	}).(pulumi.StringPtrOutput)
}

func (o ConfigServerGitPropertyResponsePtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigServerGitPropertyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

func (o ConfigServerGitPropertyResponsePtrOutput) PrivateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigServerGitPropertyResponse) *string {
		if v == nil {
			return nil
		}
		return v.PrivateKey
	}).(pulumi.StringPtrOutput)
}

func (o ConfigServerGitPropertyResponsePtrOutput) Repositories() GitPatternRepositoryResponseArrayOutput {
	return o.ApplyT(func(v *ConfigServerGitPropertyResponse) []GitPatternRepositoryResponse {
		if v == nil {
			return nil
		}
		return v.Repositories
	}).(GitPatternRepositoryResponseArrayOutput)
}

func (o ConfigServerGitPropertyResponsePtrOutput) SearchPaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ConfigServerGitPropertyResponse) []string {
		if v == nil {
			return nil
		}
		return v.SearchPaths
	}).(pulumi.StringArrayOutput)
}

func (o ConfigServerGitPropertyResponsePtrOutput) StrictHostKeyChecking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ConfigServerGitPropertyResponse) *bool {
		if v == nil {
			return nil
		}
		return v.StrictHostKeyChecking
	}).(pulumi.BoolPtrOutput)
}

func (o ConfigServerGitPropertyResponsePtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigServerGitPropertyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Uri
	}).(pulumi.StringPtrOutput)
}

func (o ConfigServerGitPropertyResponsePtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigServerGitPropertyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Username
	}).(pulumi.StringPtrOutput)
}

type ConfigServerProperties struct {
	ConfigServer *ConfigServerSettings `pulumi:"configServer"`
	Error        *Error                `pulumi:"error"`
}





type ConfigServerPropertiesInput interface {
	pulumi.Input

	ToConfigServerPropertiesOutput() ConfigServerPropertiesOutput
	ToConfigServerPropertiesOutputWithContext(context.Context) ConfigServerPropertiesOutput
}

type ConfigServerPropertiesArgs struct {
	ConfigServer ConfigServerSettingsPtrInput `pulumi:"configServer"`
	Error        ErrorPtrInput                `pulumi:"error"`
}

func (ConfigServerPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigServerProperties)(nil)).Elem()
}

func (i ConfigServerPropertiesArgs) ToConfigServerPropertiesOutput() ConfigServerPropertiesOutput {
	return i.ToConfigServerPropertiesOutputWithContext(context.Background())
}

func (i ConfigServerPropertiesArgs) ToConfigServerPropertiesOutputWithContext(ctx context.Context) ConfigServerPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigServerPropertiesOutput)
}

func (i ConfigServerPropertiesArgs) ToConfigServerPropertiesPtrOutput() ConfigServerPropertiesPtrOutput {
	return i.ToConfigServerPropertiesPtrOutputWithContext(context.Background())
}

func (i ConfigServerPropertiesArgs) ToConfigServerPropertiesPtrOutputWithContext(ctx context.Context) ConfigServerPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigServerPropertiesOutput).ToConfigServerPropertiesPtrOutputWithContext(ctx)
}









type ConfigServerPropertiesPtrInput interface {
	pulumi.Input

	ToConfigServerPropertiesPtrOutput() ConfigServerPropertiesPtrOutput
	ToConfigServerPropertiesPtrOutputWithContext(context.Context) ConfigServerPropertiesPtrOutput
}

type configServerPropertiesPtrType ConfigServerPropertiesArgs

func ConfigServerPropertiesPtr(v *ConfigServerPropertiesArgs) ConfigServerPropertiesPtrInput {
	return (*configServerPropertiesPtrType)(v)
}

func (*configServerPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigServerProperties)(nil)).Elem()
}

func (i *configServerPropertiesPtrType) ToConfigServerPropertiesPtrOutput() ConfigServerPropertiesPtrOutput {
	return i.ToConfigServerPropertiesPtrOutputWithContext(context.Background())
}

func (i *configServerPropertiesPtrType) ToConfigServerPropertiesPtrOutputWithContext(ctx context.Context) ConfigServerPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigServerPropertiesPtrOutput)
}

type ConfigServerPropertiesOutput struct{ *pulumi.OutputState }

func (ConfigServerPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigServerProperties)(nil)).Elem()
}

func (o ConfigServerPropertiesOutput) ToConfigServerPropertiesOutput() ConfigServerPropertiesOutput {
	return o
}

func (o ConfigServerPropertiesOutput) ToConfigServerPropertiesOutputWithContext(ctx context.Context) ConfigServerPropertiesOutput {
	return o
}

func (o ConfigServerPropertiesOutput) ToConfigServerPropertiesPtrOutput() ConfigServerPropertiesPtrOutput {
	return o.ToConfigServerPropertiesPtrOutputWithContext(context.Background())
}

func (o ConfigServerPropertiesOutput) ToConfigServerPropertiesPtrOutputWithContext(ctx context.Context) ConfigServerPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConfigServerProperties) *ConfigServerProperties {
		return &v
	}).(ConfigServerPropertiesPtrOutput)
}

func (o ConfigServerPropertiesOutput) ConfigServer() ConfigServerSettingsPtrOutput {
	return o.ApplyT(func(v ConfigServerProperties) *ConfigServerSettings { return v.ConfigServer }).(ConfigServerSettingsPtrOutput)
}

func (o ConfigServerPropertiesOutput) Error() ErrorPtrOutput {
	return o.ApplyT(func(v ConfigServerProperties) *Error { return v.Error }).(ErrorPtrOutput)
}

type ConfigServerPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ConfigServerPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigServerProperties)(nil)).Elem()
}

func (o ConfigServerPropertiesPtrOutput) ToConfigServerPropertiesPtrOutput() ConfigServerPropertiesPtrOutput {
	return o
}

func (o ConfigServerPropertiesPtrOutput) ToConfigServerPropertiesPtrOutputWithContext(ctx context.Context) ConfigServerPropertiesPtrOutput {
	return o
}

func (o ConfigServerPropertiesPtrOutput) Elem() ConfigServerPropertiesOutput {
	return o.ApplyT(func(v *ConfigServerProperties) ConfigServerProperties {
		if v != nil {
			return *v
		}
		var ret ConfigServerProperties
		return ret
	}).(ConfigServerPropertiesOutput)
}

func (o ConfigServerPropertiesPtrOutput) ConfigServer() ConfigServerSettingsPtrOutput {
	return o.ApplyT(func(v *ConfigServerProperties) *ConfigServerSettings {
		if v == nil {
			return nil
		}
		return v.ConfigServer
	}).(ConfigServerSettingsPtrOutput)
}

func (o ConfigServerPropertiesPtrOutput) Error() ErrorPtrOutput {
	return o.ApplyT(func(v *ConfigServerProperties) *Error {
		if v == nil {
			return nil
		}
		return v.Error
	}).(ErrorPtrOutput)
}

type ConfigServerPropertiesResponse struct {
	ConfigServer *ConfigServerSettingsResponse `pulumi:"configServer"`
	Error        *ErrorResponse                `pulumi:"error"`
	State        string                        `pulumi:"state"`
}





type ConfigServerPropertiesResponseInput interface {
	pulumi.Input

	ToConfigServerPropertiesResponseOutput() ConfigServerPropertiesResponseOutput
	ToConfigServerPropertiesResponseOutputWithContext(context.Context) ConfigServerPropertiesResponseOutput
}

type ConfigServerPropertiesResponseArgs struct {
	ConfigServer ConfigServerSettingsResponsePtrInput `pulumi:"configServer"`
	Error        ErrorResponsePtrInput                `pulumi:"error"`
	State        pulumi.StringInput                   `pulumi:"state"`
}

func (ConfigServerPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigServerPropertiesResponse)(nil)).Elem()
}

func (i ConfigServerPropertiesResponseArgs) ToConfigServerPropertiesResponseOutput() ConfigServerPropertiesResponseOutput {
	return i.ToConfigServerPropertiesResponseOutputWithContext(context.Background())
}

func (i ConfigServerPropertiesResponseArgs) ToConfigServerPropertiesResponseOutputWithContext(ctx context.Context) ConfigServerPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigServerPropertiesResponseOutput)
}

func (i ConfigServerPropertiesResponseArgs) ToConfigServerPropertiesResponsePtrOutput() ConfigServerPropertiesResponsePtrOutput {
	return i.ToConfigServerPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ConfigServerPropertiesResponseArgs) ToConfigServerPropertiesResponsePtrOutputWithContext(ctx context.Context) ConfigServerPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigServerPropertiesResponseOutput).ToConfigServerPropertiesResponsePtrOutputWithContext(ctx)
}









type ConfigServerPropertiesResponsePtrInput interface {
	pulumi.Input

	ToConfigServerPropertiesResponsePtrOutput() ConfigServerPropertiesResponsePtrOutput
	ToConfigServerPropertiesResponsePtrOutputWithContext(context.Context) ConfigServerPropertiesResponsePtrOutput
}

type configServerPropertiesResponsePtrType ConfigServerPropertiesResponseArgs

func ConfigServerPropertiesResponsePtr(v *ConfigServerPropertiesResponseArgs) ConfigServerPropertiesResponsePtrInput {
	return (*configServerPropertiesResponsePtrType)(v)
}

func (*configServerPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigServerPropertiesResponse)(nil)).Elem()
}

func (i *configServerPropertiesResponsePtrType) ToConfigServerPropertiesResponsePtrOutput() ConfigServerPropertiesResponsePtrOutput {
	return i.ToConfigServerPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *configServerPropertiesResponsePtrType) ToConfigServerPropertiesResponsePtrOutputWithContext(ctx context.Context) ConfigServerPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigServerPropertiesResponsePtrOutput)
}

type ConfigServerPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ConfigServerPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigServerPropertiesResponse)(nil)).Elem()
}

func (o ConfigServerPropertiesResponseOutput) ToConfigServerPropertiesResponseOutput() ConfigServerPropertiesResponseOutput {
	return o
}

func (o ConfigServerPropertiesResponseOutput) ToConfigServerPropertiesResponseOutputWithContext(ctx context.Context) ConfigServerPropertiesResponseOutput {
	return o
}

func (o ConfigServerPropertiesResponseOutput) ToConfigServerPropertiesResponsePtrOutput() ConfigServerPropertiesResponsePtrOutput {
	return o.ToConfigServerPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ConfigServerPropertiesResponseOutput) ToConfigServerPropertiesResponsePtrOutputWithContext(ctx context.Context) ConfigServerPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConfigServerPropertiesResponse) *ConfigServerPropertiesResponse {
		return &v
	}).(ConfigServerPropertiesResponsePtrOutput)
}

func (o ConfigServerPropertiesResponseOutput) ConfigServer() ConfigServerSettingsResponsePtrOutput {
	return o.ApplyT(func(v ConfigServerPropertiesResponse) *ConfigServerSettingsResponse { return v.ConfigServer }).(ConfigServerSettingsResponsePtrOutput)
}

func (o ConfigServerPropertiesResponseOutput) Error() ErrorResponsePtrOutput {
	return o.ApplyT(func(v ConfigServerPropertiesResponse) *ErrorResponse { return v.Error }).(ErrorResponsePtrOutput)
}

func (o ConfigServerPropertiesResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigServerPropertiesResponse) string { return v.State }).(pulumi.StringOutput)
}

type ConfigServerPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ConfigServerPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigServerPropertiesResponse)(nil)).Elem()
}

func (o ConfigServerPropertiesResponsePtrOutput) ToConfigServerPropertiesResponsePtrOutput() ConfigServerPropertiesResponsePtrOutput {
	return o
}

func (o ConfigServerPropertiesResponsePtrOutput) ToConfigServerPropertiesResponsePtrOutputWithContext(ctx context.Context) ConfigServerPropertiesResponsePtrOutput {
	return o
}

func (o ConfigServerPropertiesResponsePtrOutput) Elem() ConfigServerPropertiesResponseOutput {
	return o.ApplyT(func(v *ConfigServerPropertiesResponse) ConfigServerPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ConfigServerPropertiesResponse
		return ret
	}).(ConfigServerPropertiesResponseOutput)
}

func (o ConfigServerPropertiesResponsePtrOutput) ConfigServer() ConfigServerSettingsResponsePtrOutput {
	return o.ApplyT(func(v *ConfigServerPropertiesResponse) *ConfigServerSettingsResponse {
		if v == nil {
			return nil
		}
		return v.ConfigServer
	}).(ConfigServerSettingsResponsePtrOutput)
}

func (o ConfigServerPropertiesResponsePtrOutput) Error() ErrorResponsePtrOutput {
	return o.ApplyT(func(v *ConfigServerPropertiesResponse) *ErrorResponse {
		if v == nil {
			return nil
		}
		return v.Error
	}).(ErrorResponsePtrOutput)
}

func (o ConfigServerPropertiesResponsePtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigServerPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.State
	}).(pulumi.StringPtrOutput)
}

type ConfigServerSettings struct {
	GitProperty *ConfigServerGitProperty `pulumi:"gitProperty"`
}





type ConfigServerSettingsInput interface {
	pulumi.Input

	ToConfigServerSettingsOutput() ConfigServerSettingsOutput
	ToConfigServerSettingsOutputWithContext(context.Context) ConfigServerSettingsOutput
}

type ConfigServerSettingsArgs struct {
	GitProperty ConfigServerGitPropertyPtrInput `pulumi:"gitProperty"`
}

func (ConfigServerSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigServerSettings)(nil)).Elem()
}

func (i ConfigServerSettingsArgs) ToConfigServerSettingsOutput() ConfigServerSettingsOutput {
	return i.ToConfigServerSettingsOutputWithContext(context.Background())
}

func (i ConfigServerSettingsArgs) ToConfigServerSettingsOutputWithContext(ctx context.Context) ConfigServerSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigServerSettingsOutput)
}

func (i ConfigServerSettingsArgs) ToConfigServerSettingsPtrOutput() ConfigServerSettingsPtrOutput {
	return i.ToConfigServerSettingsPtrOutputWithContext(context.Background())
}

func (i ConfigServerSettingsArgs) ToConfigServerSettingsPtrOutputWithContext(ctx context.Context) ConfigServerSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigServerSettingsOutput).ToConfigServerSettingsPtrOutputWithContext(ctx)
}









type ConfigServerSettingsPtrInput interface {
	pulumi.Input

	ToConfigServerSettingsPtrOutput() ConfigServerSettingsPtrOutput
	ToConfigServerSettingsPtrOutputWithContext(context.Context) ConfigServerSettingsPtrOutput
}

type configServerSettingsPtrType ConfigServerSettingsArgs

func ConfigServerSettingsPtr(v *ConfigServerSettingsArgs) ConfigServerSettingsPtrInput {
	return (*configServerSettingsPtrType)(v)
}

func (*configServerSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigServerSettings)(nil)).Elem()
}

func (i *configServerSettingsPtrType) ToConfigServerSettingsPtrOutput() ConfigServerSettingsPtrOutput {
	return i.ToConfigServerSettingsPtrOutputWithContext(context.Background())
}

func (i *configServerSettingsPtrType) ToConfigServerSettingsPtrOutputWithContext(ctx context.Context) ConfigServerSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigServerSettingsPtrOutput)
}

type ConfigServerSettingsOutput struct{ *pulumi.OutputState }

func (ConfigServerSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigServerSettings)(nil)).Elem()
}

func (o ConfigServerSettingsOutput) ToConfigServerSettingsOutput() ConfigServerSettingsOutput {
	return o
}

func (o ConfigServerSettingsOutput) ToConfigServerSettingsOutputWithContext(ctx context.Context) ConfigServerSettingsOutput {
	return o
}

func (o ConfigServerSettingsOutput) ToConfigServerSettingsPtrOutput() ConfigServerSettingsPtrOutput {
	return o.ToConfigServerSettingsPtrOutputWithContext(context.Background())
}

func (o ConfigServerSettingsOutput) ToConfigServerSettingsPtrOutputWithContext(ctx context.Context) ConfigServerSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConfigServerSettings) *ConfigServerSettings {
		return &v
	}).(ConfigServerSettingsPtrOutput)
}

func (o ConfigServerSettingsOutput) GitProperty() ConfigServerGitPropertyPtrOutput {
	return o.ApplyT(func(v ConfigServerSettings) *ConfigServerGitProperty { return v.GitProperty }).(ConfigServerGitPropertyPtrOutput)
}

type ConfigServerSettingsPtrOutput struct{ *pulumi.OutputState }

func (ConfigServerSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigServerSettings)(nil)).Elem()
}

func (o ConfigServerSettingsPtrOutput) ToConfigServerSettingsPtrOutput() ConfigServerSettingsPtrOutput {
	return o
}

func (o ConfigServerSettingsPtrOutput) ToConfigServerSettingsPtrOutputWithContext(ctx context.Context) ConfigServerSettingsPtrOutput {
	return o
}

func (o ConfigServerSettingsPtrOutput) Elem() ConfigServerSettingsOutput {
	return o.ApplyT(func(v *ConfigServerSettings) ConfigServerSettings {
		if v != nil {
			return *v
		}
		var ret ConfigServerSettings
		return ret
	}).(ConfigServerSettingsOutput)
}

func (o ConfigServerSettingsPtrOutput) GitProperty() ConfigServerGitPropertyPtrOutput {
	return o.ApplyT(func(v *ConfigServerSettings) *ConfigServerGitProperty {
		if v == nil {
			return nil
		}
		return v.GitProperty
	}).(ConfigServerGitPropertyPtrOutput)
}

type ConfigServerSettingsResponse struct {
	GitProperty *ConfigServerGitPropertyResponse `pulumi:"gitProperty"`
}





type ConfigServerSettingsResponseInput interface {
	pulumi.Input

	ToConfigServerSettingsResponseOutput() ConfigServerSettingsResponseOutput
	ToConfigServerSettingsResponseOutputWithContext(context.Context) ConfigServerSettingsResponseOutput
}

type ConfigServerSettingsResponseArgs struct {
	GitProperty ConfigServerGitPropertyResponsePtrInput `pulumi:"gitProperty"`
}

func (ConfigServerSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigServerSettingsResponse)(nil)).Elem()
}

func (i ConfigServerSettingsResponseArgs) ToConfigServerSettingsResponseOutput() ConfigServerSettingsResponseOutput {
	return i.ToConfigServerSettingsResponseOutputWithContext(context.Background())
}

func (i ConfigServerSettingsResponseArgs) ToConfigServerSettingsResponseOutputWithContext(ctx context.Context) ConfigServerSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigServerSettingsResponseOutput)
}

func (i ConfigServerSettingsResponseArgs) ToConfigServerSettingsResponsePtrOutput() ConfigServerSettingsResponsePtrOutput {
	return i.ToConfigServerSettingsResponsePtrOutputWithContext(context.Background())
}

func (i ConfigServerSettingsResponseArgs) ToConfigServerSettingsResponsePtrOutputWithContext(ctx context.Context) ConfigServerSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigServerSettingsResponseOutput).ToConfigServerSettingsResponsePtrOutputWithContext(ctx)
}









type ConfigServerSettingsResponsePtrInput interface {
	pulumi.Input

	ToConfigServerSettingsResponsePtrOutput() ConfigServerSettingsResponsePtrOutput
	ToConfigServerSettingsResponsePtrOutputWithContext(context.Context) ConfigServerSettingsResponsePtrOutput
}

type configServerSettingsResponsePtrType ConfigServerSettingsResponseArgs

func ConfigServerSettingsResponsePtr(v *ConfigServerSettingsResponseArgs) ConfigServerSettingsResponsePtrInput {
	return (*configServerSettingsResponsePtrType)(v)
}

func (*configServerSettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigServerSettingsResponse)(nil)).Elem()
}

func (i *configServerSettingsResponsePtrType) ToConfigServerSettingsResponsePtrOutput() ConfigServerSettingsResponsePtrOutput {
	return i.ToConfigServerSettingsResponsePtrOutputWithContext(context.Background())
}

func (i *configServerSettingsResponsePtrType) ToConfigServerSettingsResponsePtrOutputWithContext(ctx context.Context) ConfigServerSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigServerSettingsResponsePtrOutput)
}

type ConfigServerSettingsResponseOutput struct{ *pulumi.OutputState }

func (ConfigServerSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigServerSettingsResponse)(nil)).Elem()
}

func (o ConfigServerSettingsResponseOutput) ToConfigServerSettingsResponseOutput() ConfigServerSettingsResponseOutput {
	return o
}

func (o ConfigServerSettingsResponseOutput) ToConfigServerSettingsResponseOutputWithContext(ctx context.Context) ConfigServerSettingsResponseOutput {
	return o
}

func (o ConfigServerSettingsResponseOutput) ToConfigServerSettingsResponsePtrOutput() ConfigServerSettingsResponsePtrOutput {
	return o.ToConfigServerSettingsResponsePtrOutputWithContext(context.Background())
}

func (o ConfigServerSettingsResponseOutput) ToConfigServerSettingsResponsePtrOutputWithContext(ctx context.Context) ConfigServerSettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConfigServerSettingsResponse) *ConfigServerSettingsResponse {
		return &v
	}).(ConfigServerSettingsResponsePtrOutput)
}

func (o ConfigServerSettingsResponseOutput) GitProperty() ConfigServerGitPropertyResponsePtrOutput {
	return o.ApplyT(func(v ConfigServerSettingsResponse) *ConfigServerGitPropertyResponse { return v.GitProperty }).(ConfigServerGitPropertyResponsePtrOutput)
}

type ConfigServerSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (ConfigServerSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigServerSettingsResponse)(nil)).Elem()
}

func (o ConfigServerSettingsResponsePtrOutput) ToConfigServerSettingsResponsePtrOutput() ConfigServerSettingsResponsePtrOutput {
	return o
}

func (o ConfigServerSettingsResponsePtrOutput) ToConfigServerSettingsResponsePtrOutputWithContext(ctx context.Context) ConfigServerSettingsResponsePtrOutput {
	return o
}

func (o ConfigServerSettingsResponsePtrOutput) Elem() ConfigServerSettingsResponseOutput {
	return o.ApplyT(func(v *ConfigServerSettingsResponse) ConfigServerSettingsResponse {
		if v != nil {
			return *v
		}
		var ret ConfigServerSettingsResponse
		return ret
	}).(ConfigServerSettingsResponseOutput)
}

func (o ConfigServerSettingsResponsePtrOutput) GitProperty() ConfigServerGitPropertyResponsePtrOutput {
	return o.ApplyT(func(v *ConfigServerSettingsResponse) *ConfigServerGitPropertyResponse {
		if v == nil {
			return nil
		}
		return v.GitProperty
	}).(ConfigServerGitPropertyResponsePtrOutput)
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
	InstanceCount        *int              `pulumi:"instanceCount"`
	JvmOptions           *string           `pulumi:"jvmOptions"`
	MemoryInGB           *int              `pulumi:"memoryInGB"`
	NetCoreMainEntryPath *string           `pulumi:"netCoreMainEntryPath"`
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
	if isZero(tmp.InstanceCount) {
		instanceCount_ := 1
		tmp.InstanceCount = &instanceCount_
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
	Cpu                  pulumi.IntPtrInput    `pulumi:"cpu"`
	EnvironmentVariables pulumi.StringMapInput `pulumi:"environmentVariables"`
	InstanceCount        pulumi.IntPtrInput    `pulumi:"instanceCount"`
	JvmOptions           pulumi.StringPtrInput `pulumi:"jvmOptions"`
	MemoryInGB           pulumi.IntPtrInput    `pulumi:"memoryInGB"`
	NetCoreMainEntryPath pulumi.StringPtrInput `pulumi:"netCoreMainEntryPath"`
	RuntimeVersion       pulumi.StringPtrInput `pulumi:"runtimeVersion"`
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

func (o DeploymentSettingsOutput) InstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DeploymentSettings) *int { return v.InstanceCount }).(pulumi.IntPtrOutput)
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

func (o DeploymentSettingsPtrOutput) InstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DeploymentSettings) *int {
		if v == nil {
			return nil
		}
		return v.InstanceCount
	}).(pulumi.IntPtrOutput)
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

func (o DeploymentSettingsPtrOutput) RuntimeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentSettings) *string {
		if v == nil {
			return nil
		}
		return v.RuntimeVersion
	}).(pulumi.StringPtrOutput)
}

type DeploymentSettingsResponse struct {
	Cpu                  *int              `pulumi:"cpu"`
	EnvironmentVariables map[string]string `pulumi:"environmentVariables"`
	InstanceCount        *int              `pulumi:"instanceCount"`
	JvmOptions           *string           `pulumi:"jvmOptions"`
	MemoryInGB           *int              `pulumi:"memoryInGB"`
	NetCoreMainEntryPath *string           `pulumi:"netCoreMainEntryPath"`
	RuntimeVersion       *string           `pulumi:"runtimeVersion"`
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
	if isZero(tmp.InstanceCount) {
		instanceCount_ := 1
		tmp.InstanceCount = &instanceCount_
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
	Cpu                  pulumi.IntPtrInput    `pulumi:"cpu"`
	EnvironmentVariables pulumi.StringMapInput `pulumi:"environmentVariables"`
	InstanceCount        pulumi.IntPtrInput    `pulumi:"instanceCount"`
	JvmOptions           pulumi.StringPtrInput `pulumi:"jvmOptions"`
	MemoryInGB           pulumi.IntPtrInput    `pulumi:"memoryInGB"`
	NetCoreMainEntryPath pulumi.StringPtrInput `pulumi:"netCoreMainEntryPath"`
	RuntimeVersion       pulumi.StringPtrInput `pulumi:"runtimeVersion"`
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

func (o DeploymentSettingsResponseOutput) InstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DeploymentSettingsResponse) *int { return v.InstanceCount }).(pulumi.IntPtrOutput)
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

func (o DeploymentSettingsResponsePtrOutput) InstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DeploymentSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.InstanceCount
	}).(pulumi.IntPtrOutput)
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

func (o DeploymentSettingsResponsePtrOutput) RuntimeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.RuntimeVersion
	}).(pulumi.StringPtrOutput)
}

type Error struct {
	Code    *string `pulumi:"code"`
	Message *string `pulumi:"message"`
}





type ErrorInput interface {
	pulumi.Input

	ToErrorOutput() ErrorOutput
	ToErrorOutputWithContext(context.Context) ErrorOutput
}

type ErrorArgs struct {
	Code    pulumi.StringPtrInput `pulumi:"code"`
	Message pulumi.StringPtrInput `pulumi:"message"`
}

func (ErrorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Error)(nil)).Elem()
}

func (i ErrorArgs) ToErrorOutput() ErrorOutput {
	return i.ToErrorOutputWithContext(context.Background())
}

func (i ErrorArgs) ToErrorOutputWithContext(ctx context.Context) ErrorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorOutput)
}

func (i ErrorArgs) ToErrorPtrOutput() ErrorPtrOutput {
	return i.ToErrorPtrOutputWithContext(context.Background())
}

func (i ErrorArgs) ToErrorPtrOutputWithContext(ctx context.Context) ErrorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorOutput).ToErrorPtrOutputWithContext(ctx)
}









type ErrorPtrInput interface {
	pulumi.Input

	ToErrorPtrOutput() ErrorPtrOutput
	ToErrorPtrOutputWithContext(context.Context) ErrorPtrOutput
}

type errorPtrType ErrorArgs

func ErrorPtr(v *ErrorArgs) ErrorPtrInput {
	return (*errorPtrType)(v)
}

func (*errorPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Error)(nil)).Elem()
}

func (i *errorPtrType) ToErrorPtrOutput() ErrorPtrOutput {
	return i.ToErrorPtrOutputWithContext(context.Background())
}

func (i *errorPtrType) ToErrorPtrOutputWithContext(ctx context.Context) ErrorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorPtrOutput)
}

type ErrorOutput struct{ *pulumi.OutputState }

func (ErrorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Error)(nil)).Elem()
}

func (o ErrorOutput) ToErrorOutput() ErrorOutput {
	return o
}

func (o ErrorOutput) ToErrorOutputWithContext(ctx context.Context) ErrorOutput {
	return o
}

func (o ErrorOutput) ToErrorPtrOutput() ErrorPtrOutput {
	return o.ToErrorPtrOutputWithContext(context.Background())
}

func (o ErrorOutput) ToErrorPtrOutputWithContext(ctx context.Context) ErrorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Error) *Error {
		return &v
	}).(ErrorPtrOutput)
}

func (o ErrorOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Error) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o ErrorOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Error) *string { return v.Message }).(pulumi.StringPtrOutput)
}

type ErrorPtrOutput struct{ *pulumi.OutputState }

func (ErrorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Error)(nil)).Elem()
}

func (o ErrorPtrOutput) ToErrorPtrOutput() ErrorPtrOutput {
	return o
}

func (o ErrorPtrOutput) ToErrorPtrOutputWithContext(ctx context.Context) ErrorPtrOutput {
	return o
}

func (o ErrorPtrOutput) Elem() ErrorOutput {
	return o.ApplyT(func(v *Error) Error {
		if v != nil {
			return *v
		}
		var ret Error
		return ret
	}).(ErrorOutput)
}

func (o ErrorPtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Error) *string {
		if v == nil {
			return nil
		}
		return v.Code
	}).(pulumi.StringPtrOutput)
}

func (o ErrorPtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Error) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

type ErrorResponse struct {
	Code    *string `pulumi:"code"`
	Message *string `pulumi:"message"`
}





type ErrorResponseInput interface {
	pulumi.Input

	ToErrorResponseOutput() ErrorResponseOutput
	ToErrorResponseOutputWithContext(context.Context) ErrorResponseOutput
}

type ErrorResponseArgs struct {
	Code    pulumi.StringPtrInput `pulumi:"code"`
	Message pulumi.StringPtrInput `pulumi:"message"`
}

func (ErrorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorResponse)(nil)).Elem()
}

func (i ErrorResponseArgs) ToErrorResponseOutput() ErrorResponseOutput {
	return i.ToErrorResponseOutputWithContext(context.Background())
}

func (i ErrorResponseArgs) ToErrorResponseOutputWithContext(ctx context.Context) ErrorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorResponseOutput)
}

func (i ErrorResponseArgs) ToErrorResponsePtrOutput() ErrorResponsePtrOutput {
	return i.ToErrorResponsePtrOutputWithContext(context.Background())
}

func (i ErrorResponseArgs) ToErrorResponsePtrOutputWithContext(ctx context.Context) ErrorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorResponseOutput).ToErrorResponsePtrOutputWithContext(ctx)
}









type ErrorResponsePtrInput interface {
	pulumi.Input

	ToErrorResponsePtrOutput() ErrorResponsePtrOutput
	ToErrorResponsePtrOutputWithContext(context.Context) ErrorResponsePtrOutput
}

type errorResponsePtrType ErrorResponseArgs

func ErrorResponsePtr(v *ErrorResponseArgs) ErrorResponsePtrInput {
	return (*errorResponsePtrType)(v)
}

func (*errorResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ErrorResponse)(nil)).Elem()
}

func (i *errorResponsePtrType) ToErrorResponsePtrOutput() ErrorResponsePtrOutput {
	return i.ToErrorResponsePtrOutputWithContext(context.Background())
}

func (i *errorResponsePtrType) ToErrorResponsePtrOutputWithContext(ctx context.Context) ErrorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorResponsePtrOutput)
}

type ErrorResponseOutput struct{ *pulumi.OutputState }

func (ErrorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorResponse)(nil)).Elem()
}

func (o ErrorResponseOutput) ToErrorResponseOutput() ErrorResponseOutput {
	return o
}

func (o ErrorResponseOutput) ToErrorResponseOutputWithContext(ctx context.Context) ErrorResponseOutput {
	return o
}

func (o ErrorResponseOutput) ToErrorResponsePtrOutput() ErrorResponsePtrOutput {
	return o.ToErrorResponsePtrOutputWithContext(context.Background())
}

func (o ErrorResponseOutput) ToErrorResponsePtrOutputWithContext(ctx context.Context) ErrorResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ErrorResponse) *ErrorResponse {
		return &v
	}).(ErrorResponsePtrOutput)
}

func (o ErrorResponseOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ErrorResponse) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o ErrorResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ErrorResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

type ErrorResponsePtrOutput struct{ *pulumi.OutputState }

func (ErrorResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ErrorResponse)(nil)).Elem()
}

func (o ErrorResponsePtrOutput) ToErrorResponsePtrOutput() ErrorResponsePtrOutput {
	return o
}

func (o ErrorResponsePtrOutput) ToErrorResponsePtrOutputWithContext(ctx context.Context) ErrorResponsePtrOutput {
	return o
}

func (o ErrorResponsePtrOutput) Elem() ErrorResponseOutput {
	return o.ApplyT(func(v *ErrorResponse) ErrorResponse {
		if v != nil {
			return *v
		}
		var ret ErrorResponse
		return ret
	}).(ErrorResponseOutput)
}

func (o ErrorResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Code
	}).(pulumi.StringPtrOutput)
}

func (o ErrorResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

type GitPatternRepository struct {
	HostKey               *string  `pulumi:"hostKey"`
	HostKeyAlgorithm      *string  `pulumi:"hostKeyAlgorithm"`
	Label                 *string  `pulumi:"label"`
	Name                  string   `pulumi:"name"`
	Password              *string  `pulumi:"password"`
	Pattern               []string `pulumi:"pattern"`
	PrivateKey            *string  `pulumi:"privateKey"`
	SearchPaths           []string `pulumi:"searchPaths"`
	StrictHostKeyChecking *bool    `pulumi:"strictHostKeyChecking"`
	Uri                   string   `pulumi:"uri"`
	Username              *string  `pulumi:"username"`
}





type GitPatternRepositoryInput interface {
	pulumi.Input

	ToGitPatternRepositoryOutput() GitPatternRepositoryOutput
	ToGitPatternRepositoryOutputWithContext(context.Context) GitPatternRepositoryOutput
}

type GitPatternRepositoryArgs struct {
	HostKey               pulumi.StringPtrInput   `pulumi:"hostKey"`
	HostKeyAlgorithm      pulumi.StringPtrInput   `pulumi:"hostKeyAlgorithm"`
	Label                 pulumi.StringPtrInput   `pulumi:"label"`
	Name                  pulumi.StringInput      `pulumi:"name"`
	Password              pulumi.StringPtrInput   `pulumi:"password"`
	Pattern               pulumi.StringArrayInput `pulumi:"pattern"`
	PrivateKey            pulumi.StringPtrInput   `pulumi:"privateKey"`
	SearchPaths           pulumi.StringArrayInput `pulumi:"searchPaths"`
	StrictHostKeyChecking pulumi.BoolPtrInput     `pulumi:"strictHostKeyChecking"`
	Uri                   pulumi.StringInput      `pulumi:"uri"`
	Username              pulumi.StringPtrInput   `pulumi:"username"`
}

func (GitPatternRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GitPatternRepository)(nil)).Elem()
}

func (i GitPatternRepositoryArgs) ToGitPatternRepositoryOutput() GitPatternRepositoryOutput {
	return i.ToGitPatternRepositoryOutputWithContext(context.Background())
}

func (i GitPatternRepositoryArgs) ToGitPatternRepositoryOutputWithContext(ctx context.Context) GitPatternRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitPatternRepositoryOutput)
}





type GitPatternRepositoryArrayInput interface {
	pulumi.Input

	ToGitPatternRepositoryArrayOutput() GitPatternRepositoryArrayOutput
	ToGitPatternRepositoryArrayOutputWithContext(context.Context) GitPatternRepositoryArrayOutput
}

type GitPatternRepositoryArray []GitPatternRepositoryInput

func (GitPatternRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GitPatternRepository)(nil)).Elem()
}

func (i GitPatternRepositoryArray) ToGitPatternRepositoryArrayOutput() GitPatternRepositoryArrayOutput {
	return i.ToGitPatternRepositoryArrayOutputWithContext(context.Background())
}

func (i GitPatternRepositoryArray) ToGitPatternRepositoryArrayOutputWithContext(ctx context.Context) GitPatternRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitPatternRepositoryArrayOutput)
}

type GitPatternRepositoryOutput struct{ *pulumi.OutputState }

func (GitPatternRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GitPatternRepository)(nil)).Elem()
}

func (o GitPatternRepositoryOutput) ToGitPatternRepositoryOutput() GitPatternRepositoryOutput {
	return o
}

func (o GitPatternRepositoryOutput) ToGitPatternRepositoryOutputWithContext(ctx context.Context) GitPatternRepositoryOutput {
	return o
}

func (o GitPatternRepositoryOutput) HostKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitPatternRepository) *string { return v.HostKey }).(pulumi.StringPtrOutput)
}

func (o GitPatternRepositoryOutput) HostKeyAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitPatternRepository) *string { return v.HostKeyAlgorithm }).(pulumi.StringPtrOutput)
}

func (o GitPatternRepositoryOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitPatternRepository) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o GitPatternRepositoryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GitPatternRepository) string { return v.Name }).(pulumi.StringOutput)
}

func (o GitPatternRepositoryOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitPatternRepository) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o GitPatternRepositoryOutput) Pattern() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GitPatternRepository) []string { return v.Pattern }).(pulumi.StringArrayOutput)
}

func (o GitPatternRepositoryOutput) PrivateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitPatternRepository) *string { return v.PrivateKey }).(pulumi.StringPtrOutput)
}

func (o GitPatternRepositoryOutput) SearchPaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GitPatternRepository) []string { return v.SearchPaths }).(pulumi.StringArrayOutput)
}

func (o GitPatternRepositoryOutput) StrictHostKeyChecking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GitPatternRepository) *bool { return v.StrictHostKeyChecking }).(pulumi.BoolPtrOutput)
}

func (o GitPatternRepositoryOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v GitPatternRepository) string { return v.Uri }).(pulumi.StringOutput)
}

func (o GitPatternRepositoryOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitPatternRepository) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type GitPatternRepositoryArrayOutput struct{ *pulumi.OutputState }

func (GitPatternRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GitPatternRepository)(nil)).Elem()
}

func (o GitPatternRepositoryArrayOutput) ToGitPatternRepositoryArrayOutput() GitPatternRepositoryArrayOutput {
	return o
}

func (o GitPatternRepositoryArrayOutput) ToGitPatternRepositoryArrayOutputWithContext(ctx context.Context) GitPatternRepositoryArrayOutput {
	return o
}

func (o GitPatternRepositoryArrayOutput) Index(i pulumi.IntInput) GitPatternRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GitPatternRepository {
		return vs[0].([]GitPatternRepository)[vs[1].(int)]
	}).(GitPatternRepositoryOutput)
}

type GitPatternRepositoryResponse struct {
	HostKey               *string  `pulumi:"hostKey"`
	HostKeyAlgorithm      *string  `pulumi:"hostKeyAlgorithm"`
	Label                 *string  `pulumi:"label"`
	Name                  string   `pulumi:"name"`
	Password              *string  `pulumi:"password"`
	Pattern               []string `pulumi:"pattern"`
	PrivateKey            *string  `pulumi:"privateKey"`
	SearchPaths           []string `pulumi:"searchPaths"`
	StrictHostKeyChecking *bool    `pulumi:"strictHostKeyChecking"`
	Uri                   string   `pulumi:"uri"`
	Username              *string  `pulumi:"username"`
}





type GitPatternRepositoryResponseInput interface {
	pulumi.Input

	ToGitPatternRepositoryResponseOutput() GitPatternRepositoryResponseOutput
	ToGitPatternRepositoryResponseOutputWithContext(context.Context) GitPatternRepositoryResponseOutput
}

type GitPatternRepositoryResponseArgs struct {
	HostKey               pulumi.StringPtrInput   `pulumi:"hostKey"`
	HostKeyAlgorithm      pulumi.StringPtrInput   `pulumi:"hostKeyAlgorithm"`
	Label                 pulumi.StringPtrInput   `pulumi:"label"`
	Name                  pulumi.StringInput      `pulumi:"name"`
	Password              pulumi.StringPtrInput   `pulumi:"password"`
	Pattern               pulumi.StringArrayInput `pulumi:"pattern"`
	PrivateKey            pulumi.StringPtrInput   `pulumi:"privateKey"`
	SearchPaths           pulumi.StringArrayInput `pulumi:"searchPaths"`
	StrictHostKeyChecking pulumi.BoolPtrInput     `pulumi:"strictHostKeyChecking"`
	Uri                   pulumi.StringInput      `pulumi:"uri"`
	Username              pulumi.StringPtrInput   `pulumi:"username"`
}

func (GitPatternRepositoryResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GitPatternRepositoryResponse)(nil)).Elem()
}

func (i GitPatternRepositoryResponseArgs) ToGitPatternRepositoryResponseOutput() GitPatternRepositoryResponseOutput {
	return i.ToGitPatternRepositoryResponseOutputWithContext(context.Background())
}

func (i GitPatternRepositoryResponseArgs) ToGitPatternRepositoryResponseOutputWithContext(ctx context.Context) GitPatternRepositoryResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitPatternRepositoryResponseOutput)
}





type GitPatternRepositoryResponseArrayInput interface {
	pulumi.Input

	ToGitPatternRepositoryResponseArrayOutput() GitPatternRepositoryResponseArrayOutput
	ToGitPatternRepositoryResponseArrayOutputWithContext(context.Context) GitPatternRepositoryResponseArrayOutput
}

type GitPatternRepositoryResponseArray []GitPatternRepositoryResponseInput

func (GitPatternRepositoryResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GitPatternRepositoryResponse)(nil)).Elem()
}

func (i GitPatternRepositoryResponseArray) ToGitPatternRepositoryResponseArrayOutput() GitPatternRepositoryResponseArrayOutput {
	return i.ToGitPatternRepositoryResponseArrayOutputWithContext(context.Background())
}

func (i GitPatternRepositoryResponseArray) ToGitPatternRepositoryResponseArrayOutputWithContext(ctx context.Context) GitPatternRepositoryResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitPatternRepositoryResponseArrayOutput)
}

type GitPatternRepositoryResponseOutput struct{ *pulumi.OutputState }

func (GitPatternRepositoryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GitPatternRepositoryResponse)(nil)).Elem()
}

func (o GitPatternRepositoryResponseOutput) ToGitPatternRepositoryResponseOutput() GitPatternRepositoryResponseOutput {
	return o
}

func (o GitPatternRepositoryResponseOutput) ToGitPatternRepositoryResponseOutputWithContext(ctx context.Context) GitPatternRepositoryResponseOutput {
	return o
}

func (o GitPatternRepositoryResponseOutput) HostKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitPatternRepositoryResponse) *string { return v.HostKey }).(pulumi.StringPtrOutput)
}

func (o GitPatternRepositoryResponseOutput) HostKeyAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitPatternRepositoryResponse) *string { return v.HostKeyAlgorithm }).(pulumi.StringPtrOutput)
}

func (o GitPatternRepositoryResponseOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitPatternRepositoryResponse) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o GitPatternRepositoryResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GitPatternRepositoryResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o GitPatternRepositoryResponseOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitPatternRepositoryResponse) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o GitPatternRepositoryResponseOutput) Pattern() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GitPatternRepositoryResponse) []string { return v.Pattern }).(pulumi.StringArrayOutput)
}

func (o GitPatternRepositoryResponseOutput) PrivateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitPatternRepositoryResponse) *string { return v.PrivateKey }).(pulumi.StringPtrOutput)
}

func (o GitPatternRepositoryResponseOutput) SearchPaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GitPatternRepositoryResponse) []string { return v.SearchPaths }).(pulumi.StringArrayOutput)
}

func (o GitPatternRepositoryResponseOutput) StrictHostKeyChecking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GitPatternRepositoryResponse) *bool { return v.StrictHostKeyChecking }).(pulumi.BoolPtrOutput)
}

func (o GitPatternRepositoryResponseOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v GitPatternRepositoryResponse) string { return v.Uri }).(pulumi.StringOutput)
}

func (o GitPatternRepositoryResponseOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitPatternRepositoryResponse) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type GitPatternRepositoryResponseArrayOutput struct{ *pulumi.OutputState }

func (GitPatternRepositoryResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GitPatternRepositoryResponse)(nil)).Elem()
}

func (o GitPatternRepositoryResponseArrayOutput) ToGitPatternRepositoryResponseArrayOutput() GitPatternRepositoryResponseArrayOutput {
	return o
}

func (o GitPatternRepositoryResponseArrayOutput) ToGitPatternRepositoryResponseArrayOutputWithContext(ctx context.Context) GitPatternRepositoryResponseArrayOutput {
	return o
}

func (o GitPatternRepositoryResponseArrayOutput) Index(i pulumi.IntInput) GitPatternRepositoryResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GitPatternRepositoryResponse {
		return vs[0].([]GitPatternRepositoryResponse)[vs[1].(int)]
	}).(GitPatternRepositoryResponseOutput)
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

type TraceProperties struct {
	AppInsightInstrumentationKey *string `pulumi:"appInsightInstrumentationKey"`
	Enabled                      *bool   `pulumi:"enabled"`
	Error                        *Error  `pulumi:"error"`
}





type TracePropertiesInput interface {
	pulumi.Input

	ToTracePropertiesOutput() TracePropertiesOutput
	ToTracePropertiesOutputWithContext(context.Context) TracePropertiesOutput
}

type TracePropertiesArgs struct {
	AppInsightInstrumentationKey pulumi.StringPtrInput `pulumi:"appInsightInstrumentationKey"`
	Enabled                      pulumi.BoolPtrInput   `pulumi:"enabled"`
	Error                        ErrorPtrInput         `pulumi:"error"`
}

func (TracePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TraceProperties)(nil)).Elem()
}

func (i TracePropertiesArgs) ToTracePropertiesOutput() TracePropertiesOutput {
	return i.ToTracePropertiesOutputWithContext(context.Background())
}

func (i TracePropertiesArgs) ToTracePropertiesOutputWithContext(ctx context.Context) TracePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TracePropertiesOutput)
}

func (i TracePropertiesArgs) ToTracePropertiesPtrOutput() TracePropertiesPtrOutput {
	return i.ToTracePropertiesPtrOutputWithContext(context.Background())
}

func (i TracePropertiesArgs) ToTracePropertiesPtrOutputWithContext(ctx context.Context) TracePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TracePropertiesOutput).ToTracePropertiesPtrOutputWithContext(ctx)
}









type TracePropertiesPtrInput interface {
	pulumi.Input

	ToTracePropertiesPtrOutput() TracePropertiesPtrOutput
	ToTracePropertiesPtrOutputWithContext(context.Context) TracePropertiesPtrOutput
}

type tracePropertiesPtrType TracePropertiesArgs

func TracePropertiesPtr(v *TracePropertiesArgs) TracePropertiesPtrInput {
	return (*tracePropertiesPtrType)(v)
}

func (*tracePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TraceProperties)(nil)).Elem()
}

func (i *tracePropertiesPtrType) ToTracePropertiesPtrOutput() TracePropertiesPtrOutput {
	return i.ToTracePropertiesPtrOutputWithContext(context.Background())
}

func (i *tracePropertiesPtrType) ToTracePropertiesPtrOutputWithContext(ctx context.Context) TracePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TracePropertiesPtrOutput)
}

type TracePropertiesOutput struct{ *pulumi.OutputState }

func (TracePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TraceProperties)(nil)).Elem()
}

func (o TracePropertiesOutput) ToTracePropertiesOutput() TracePropertiesOutput {
	return o
}

func (o TracePropertiesOutput) ToTracePropertiesOutputWithContext(ctx context.Context) TracePropertiesOutput {
	return o
}

func (o TracePropertiesOutput) ToTracePropertiesPtrOutput() TracePropertiesPtrOutput {
	return o.ToTracePropertiesPtrOutputWithContext(context.Background())
}

func (o TracePropertiesOutput) ToTracePropertiesPtrOutputWithContext(ctx context.Context) TracePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TraceProperties) *TraceProperties {
		return &v
	}).(TracePropertiesPtrOutput)
}

func (o TracePropertiesOutput) AppInsightInstrumentationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TraceProperties) *string { return v.AppInsightInstrumentationKey }).(pulumi.StringPtrOutput)
}

func (o TracePropertiesOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v TraceProperties) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o TracePropertiesOutput) Error() ErrorPtrOutput {
	return o.ApplyT(func(v TraceProperties) *Error { return v.Error }).(ErrorPtrOutput)
}

type TracePropertiesPtrOutput struct{ *pulumi.OutputState }

func (TracePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TraceProperties)(nil)).Elem()
}

func (o TracePropertiesPtrOutput) ToTracePropertiesPtrOutput() TracePropertiesPtrOutput {
	return o
}

func (o TracePropertiesPtrOutput) ToTracePropertiesPtrOutputWithContext(ctx context.Context) TracePropertiesPtrOutput {
	return o
}

func (o TracePropertiesPtrOutput) Elem() TracePropertiesOutput {
	return o.ApplyT(func(v *TraceProperties) TraceProperties {
		if v != nil {
			return *v
		}
		var ret TraceProperties
		return ret
	}).(TracePropertiesOutput)
}

func (o TracePropertiesPtrOutput) AppInsightInstrumentationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TraceProperties) *string {
		if v == nil {
			return nil
		}
		return v.AppInsightInstrumentationKey
	}).(pulumi.StringPtrOutput)
}

func (o TracePropertiesPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TraceProperties) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o TracePropertiesPtrOutput) Error() ErrorPtrOutput {
	return o.ApplyT(func(v *TraceProperties) *Error {
		if v == nil {
			return nil
		}
		return v.Error
	}).(ErrorPtrOutput)
}

type TracePropertiesResponse struct {
	AppInsightInstrumentationKey *string        `pulumi:"appInsightInstrumentationKey"`
	Enabled                      *bool          `pulumi:"enabled"`
	Error                        *ErrorResponse `pulumi:"error"`
	State                        string         `pulumi:"state"`
}





type TracePropertiesResponseInput interface {
	pulumi.Input

	ToTracePropertiesResponseOutput() TracePropertiesResponseOutput
	ToTracePropertiesResponseOutputWithContext(context.Context) TracePropertiesResponseOutput
}

type TracePropertiesResponseArgs struct {
	AppInsightInstrumentationKey pulumi.StringPtrInput `pulumi:"appInsightInstrumentationKey"`
	Enabled                      pulumi.BoolPtrInput   `pulumi:"enabled"`
	Error                        ErrorResponsePtrInput `pulumi:"error"`
	State                        pulumi.StringInput    `pulumi:"state"`
}

func (TracePropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TracePropertiesResponse)(nil)).Elem()
}

func (i TracePropertiesResponseArgs) ToTracePropertiesResponseOutput() TracePropertiesResponseOutput {
	return i.ToTracePropertiesResponseOutputWithContext(context.Background())
}

func (i TracePropertiesResponseArgs) ToTracePropertiesResponseOutputWithContext(ctx context.Context) TracePropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TracePropertiesResponseOutput)
}

func (i TracePropertiesResponseArgs) ToTracePropertiesResponsePtrOutput() TracePropertiesResponsePtrOutput {
	return i.ToTracePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i TracePropertiesResponseArgs) ToTracePropertiesResponsePtrOutputWithContext(ctx context.Context) TracePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TracePropertiesResponseOutput).ToTracePropertiesResponsePtrOutputWithContext(ctx)
}









type TracePropertiesResponsePtrInput interface {
	pulumi.Input

	ToTracePropertiesResponsePtrOutput() TracePropertiesResponsePtrOutput
	ToTracePropertiesResponsePtrOutputWithContext(context.Context) TracePropertiesResponsePtrOutput
}

type tracePropertiesResponsePtrType TracePropertiesResponseArgs

func TracePropertiesResponsePtr(v *TracePropertiesResponseArgs) TracePropertiesResponsePtrInput {
	return (*tracePropertiesResponsePtrType)(v)
}

func (*tracePropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TracePropertiesResponse)(nil)).Elem()
}

func (i *tracePropertiesResponsePtrType) ToTracePropertiesResponsePtrOutput() TracePropertiesResponsePtrOutput {
	return i.ToTracePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *tracePropertiesResponsePtrType) ToTracePropertiesResponsePtrOutputWithContext(ctx context.Context) TracePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TracePropertiesResponsePtrOutput)
}

type TracePropertiesResponseOutput struct{ *pulumi.OutputState }

func (TracePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TracePropertiesResponse)(nil)).Elem()
}

func (o TracePropertiesResponseOutput) ToTracePropertiesResponseOutput() TracePropertiesResponseOutput {
	return o
}

func (o TracePropertiesResponseOutput) ToTracePropertiesResponseOutputWithContext(ctx context.Context) TracePropertiesResponseOutput {
	return o
}

func (o TracePropertiesResponseOutput) ToTracePropertiesResponsePtrOutput() TracePropertiesResponsePtrOutput {
	return o.ToTracePropertiesResponsePtrOutputWithContext(context.Background())
}

func (o TracePropertiesResponseOutput) ToTracePropertiesResponsePtrOutputWithContext(ctx context.Context) TracePropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TracePropertiesResponse) *TracePropertiesResponse {
		return &v
	}).(TracePropertiesResponsePtrOutput)
}

func (o TracePropertiesResponseOutput) AppInsightInstrumentationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TracePropertiesResponse) *string { return v.AppInsightInstrumentationKey }).(pulumi.StringPtrOutput)
}

func (o TracePropertiesResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v TracePropertiesResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o TracePropertiesResponseOutput) Error() ErrorResponsePtrOutput {
	return o.ApplyT(func(v TracePropertiesResponse) *ErrorResponse { return v.Error }).(ErrorResponsePtrOutput)
}

func (o TracePropertiesResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v TracePropertiesResponse) string { return v.State }).(pulumi.StringOutput)
}

type TracePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (TracePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TracePropertiesResponse)(nil)).Elem()
}

func (o TracePropertiesResponsePtrOutput) ToTracePropertiesResponsePtrOutput() TracePropertiesResponsePtrOutput {
	return o
}

func (o TracePropertiesResponsePtrOutput) ToTracePropertiesResponsePtrOutputWithContext(ctx context.Context) TracePropertiesResponsePtrOutput {
	return o
}

func (o TracePropertiesResponsePtrOutput) Elem() TracePropertiesResponseOutput {
	return o.ApplyT(func(v *TracePropertiesResponse) TracePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret TracePropertiesResponse
		return ret
	}).(TracePropertiesResponseOutput)
}

func (o TracePropertiesResponsePtrOutput) AppInsightInstrumentationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TracePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.AppInsightInstrumentationKey
	}).(pulumi.StringPtrOutput)
}

func (o TracePropertiesResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TracePropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o TracePropertiesResponsePtrOutput) Error() ErrorResponsePtrOutput {
	return o.ApplyT(func(v *TracePropertiesResponse) *ErrorResponse {
		if v == nil {
			return nil
		}
		return v.Error
	}).(ErrorResponsePtrOutput)
}

func (o TracePropertiesResponsePtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TracePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.State
	}).(pulumi.StringPtrOutput)
}

type UserSourceInfo struct {
	ArtifactSelector *string `pulumi:"artifactSelector"`
	RelativePath     *string `pulumi:"relativePath"`
	Type             *string `pulumi:"type"`
	Version          *string `pulumi:"version"`
}





type UserSourceInfoInput interface {
	pulumi.Input

	ToUserSourceInfoOutput() UserSourceInfoOutput
	ToUserSourceInfoOutputWithContext(context.Context) UserSourceInfoOutput
}

type UserSourceInfoArgs struct {
	ArtifactSelector pulumi.StringPtrInput `pulumi:"artifactSelector"`
	RelativePath     pulumi.StringPtrInput `pulumi:"relativePath"`
	Type             pulumi.StringPtrInput `pulumi:"type"`
	Version          pulumi.StringPtrInput `pulumi:"version"`
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
	ArtifactSelector *string `pulumi:"artifactSelector"`
	RelativePath     *string `pulumi:"relativePath"`
	Type             *string `pulumi:"type"`
	Version          *string `pulumi:"version"`
}





type UserSourceInfoResponseInput interface {
	pulumi.Input

	ToUserSourceInfoResponseOutput() UserSourceInfoResponseOutput
	ToUserSourceInfoResponseOutputWithContext(context.Context) UserSourceInfoResponseOutput
}

type UserSourceInfoResponseArgs struct {
	ArtifactSelector pulumi.StringPtrInput `pulumi:"artifactSelector"`
	RelativePath     pulumi.StringPtrInput `pulumi:"relativePath"`
	Type             pulumi.StringPtrInput `pulumi:"type"`
	Version          pulumi.StringPtrInput `pulumi:"version"`
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
	pulumi.RegisterOutputType(ConfigServerGitPropertyOutput{})
	pulumi.RegisterOutputType(ConfigServerGitPropertyPtrOutput{})
	pulumi.RegisterOutputType(ConfigServerGitPropertyResponseOutput{})
	pulumi.RegisterOutputType(ConfigServerGitPropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(ConfigServerPropertiesOutput{})
	pulumi.RegisterOutputType(ConfigServerPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ConfigServerPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ConfigServerPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ConfigServerSettingsOutput{})
	pulumi.RegisterOutputType(ConfigServerSettingsPtrOutput{})
	pulumi.RegisterOutputType(ConfigServerSettingsResponseOutput{})
	pulumi.RegisterOutputType(ConfigServerSettingsResponsePtrOutput{})
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
	pulumi.RegisterOutputType(ErrorOutput{})
	pulumi.RegisterOutputType(ErrorPtrOutput{})
	pulumi.RegisterOutputType(ErrorResponseOutput{})
	pulumi.RegisterOutputType(ErrorResponsePtrOutput{})
	pulumi.RegisterOutputType(GitPatternRepositoryOutput{})
	pulumi.RegisterOutputType(GitPatternRepositoryArrayOutput{})
	pulumi.RegisterOutputType(GitPatternRepositoryResponseOutput{})
	pulumi.RegisterOutputType(GitPatternRepositoryResponseArrayOutput{})
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
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(TemporaryDiskOutput{})
	pulumi.RegisterOutputType(TemporaryDiskPtrOutput{})
	pulumi.RegisterOutputType(TemporaryDiskResponseOutput{})
	pulumi.RegisterOutputType(TemporaryDiskResponsePtrOutput{})
	pulumi.RegisterOutputType(TracePropertiesOutput{})
	pulumi.RegisterOutputType(TracePropertiesPtrOutput{})
	pulumi.RegisterOutputType(TracePropertiesResponseOutput{})
	pulumi.RegisterOutputType(TracePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(UserSourceInfoOutput{})
	pulumi.RegisterOutputType(UserSourceInfoPtrOutput{})
	pulumi.RegisterOutputType(UserSourceInfoResponseOutput{})
	pulumi.RegisterOutputType(UserSourceInfoResponsePtrOutput{})
}
