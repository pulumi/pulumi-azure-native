


package v20221101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AzureBlobDefinition struct {
	AccountKey            *string                     `pulumi:"accountKey"`
	ContainerName         *string                     `pulumi:"containerName"`
	LocalAuthRef          *string                     `pulumi:"localAuthRef"`
	ManagedIdentity       *ManagedIdentityDefinition  `pulumi:"managedIdentity"`
	SasToken              *string                     `pulumi:"sasToken"`
	ServicePrincipal      *ServicePrincipalDefinition `pulumi:"servicePrincipal"`
	SyncIntervalInSeconds *float64                    `pulumi:"syncIntervalInSeconds"`
	TimeoutInSeconds      *float64                    `pulumi:"timeoutInSeconds"`
	Url                   *string                     `pulumi:"url"`
}


func (val *AzureBlobDefinition) Defaults() *AzureBlobDefinition {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ServicePrincipal = tmp.ServicePrincipal.Defaults()

	if isZero(tmp.SyncIntervalInSeconds) {
		syncIntervalInSeconds_ := 600.0
		tmp.SyncIntervalInSeconds = &syncIntervalInSeconds_
	}
	if isZero(tmp.TimeoutInSeconds) {
		timeoutInSeconds_ := 600.0
		tmp.TimeoutInSeconds = &timeoutInSeconds_
	}
	return &tmp
}





type AzureBlobDefinitionInput interface {
	pulumi.Input

	ToAzureBlobDefinitionOutput() AzureBlobDefinitionOutput
	ToAzureBlobDefinitionOutputWithContext(context.Context) AzureBlobDefinitionOutput
}

type AzureBlobDefinitionArgs struct {
	AccountKey            pulumi.StringPtrInput              `pulumi:"accountKey"`
	ContainerName         pulumi.StringPtrInput              `pulumi:"containerName"`
	LocalAuthRef          pulumi.StringPtrInput              `pulumi:"localAuthRef"`
	ManagedIdentity       ManagedIdentityDefinitionPtrInput  `pulumi:"managedIdentity"`
	SasToken              pulumi.StringPtrInput              `pulumi:"sasToken"`
	ServicePrincipal      ServicePrincipalDefinitionPtrInput `pulumi:"servicePrincipal"`
	SyncIntervalInSeconds pulumi.Float64PtrInput             `pulumi:"syncIntervalInSeconds"`
	TimeoutInSeconds      pulumi.Float64PtrInput             `pulumi:"timeoutInSeconds"`
	Url                   pulumi.StringPtrInput              `pulumi:"url"`
}


func (val *AzureBlobDefinitionArgs) Defaults() *AzureBlobDefinitionArgs {
	if val == nil {
		return nil
	}
	tmp := *val

	if isZero(tmp.SyncIntervalInSeconds) {
		tmp.SyncIntervalInSeconds = pulumi.Float64Ptr(600.0)
	}
	if isZero(tmp.TimeoutInSeconds) {
		tmp.TimeoutInSeconds = pulumi.Float64Ptr(600.0)
	}
	return &tmp
}
func (AzureBlobDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureBlobDefinition)(nil)).Elem()
}

func (i AzureBlobDefinitionArgs) ToAzureBlobDefinitionOutput() AzureBlobDefinitionOutput {
	return i.ToAzureBlobDefinitionOutputWithContext(context.Background())
}

func (i AzureBlobDefinitionArgs) ToAzureBlobDefinitionOutputWithContext(ctx context.Context) AzureBlobDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureBlobDefinitionOutput)
}

func (i AzureBlobDefinitionArgs) ToAzureBlobDefinitionPtrOutput() AzureBlobDefinitionPtrOutput {
	return i.ToAzureBlobDefinitionPtrOutputWithContext(context.Background())
}

func (i AzureBlobDefinitionArgs) ToAzureBlobDefinitionPtrOutputWithContext(ctx context.Context) AzureBlobDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureBlobDefinitionOutput).ToAzureBlobDefinitionPtrOutputWithContext(ctx)
}









type AzureBlobDefinitionPtrInput interface {
	pulumi.Input

	ToAzureBlobDefinitionPtrOutput() AzureBlobDefinitionPtrOutput
	ToAzureBlobDefinitionPtrOutputWithContext(context.Context) AzureBlobDefinitionPtrOutput
}

type azureBlobDefinitionPtrType AzureBlobDefinitionArgs

func AzureBlobDefinitionPtr(v *AzureBlobDefinitionArgs) AzureBlobDefinitionPtrInput {
	return (*azureBlobDefinitionPtrType)(v)
}

func (*azureBlobDefinitionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureBlobDefinition)(nil)).Elem()
}

func (i *azureBlobDefinitionPtrType) ToAzureBlobDefinitionPtrOutput() AzureBlobDefinitionPtrOutput {
	return i.ToAzureBlobDefinitionPtrOutputWithContext(context.Background())
}

func (i *azureBlobDefinitionPtrType) ToAzureBlobDefinitionPtrOutputWithContext(ctx context.Context) AzureBlobDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureBlobDefinitionPtrOutput)
}

type AzureBlobDefinitionOutput struct{ *pulumi.OutputState }

func (AzureBlobDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureBlobDefinition)(nil)).Elem()
}

func (o AzureBlobDefinitionOutput) ToAzureBlobDefinitionOutput() AzureBlobDefinitionOutput {
	return o
}

func (o AzureBlobDefinitionOutput) ToAzureBlobDefinitionOutputWithContext(ctx context.Context) AzureBlobDefinitionOutput {
	return o
}

func (o AzureBlobDefinitionOutput) ToAzureBlobDefinitionPtrOutput() AzureBlobDefinitionPtrOutput {
	return o.ToAzureBlobDefinitionPtrOutputWithContext(context.Background())
}

func (o AzureBlobDefinitionOutput) ToAzureBlobDefinitionPtrOutputWithContext(ctx context.Context) AzureBlobDefinitionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureBlobDefinition) *AzureBlobDefinition {
		return &v
	}).(AzureBlobDefinitionPtrOutput)
}

func (o AzureBlobDefinitionOutput) AccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureBlobDefinition) *string { return v.AccountKey }).(pulumi.StringPtrOutput)
}

func (o AzureBlobDefinitionOutput) ContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureBlobDefinition) *string { return v.ContainerName }).(pulumi.StringPtrOutput)
}

func (o AzureBlobDefinitionOutput) LocalAuthRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureBlobDefinition) *string { return v.LocalAuthRef }).(pulumi.StringPtrOutput)
}

func (o AzureBlobDefinitionOutput) ManagedIdentity() ManagedIdentityDefinitionPtrOutput {
	return o.ApplyT(func(v AzureBlobDefinition) *ManagedIdentityDefinition { return v.ManagedIdentity }).(ManagedIdentityDefinitionPtrOutput)
}

func (o AzureBlobDefinitionOutput) SasToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureBlobDefinition) *string { return v.SasToken }).(pulumi.StringPtrOutput)
}

func (o AzureBlobDefinitionOutput) ServicePrincipal() ServicePrincipalDefinitionPtrOutput {
	return o.ApplyT(func(v AzureBlobDefinition) *ServicePrincipalDefinition { return v.ServicePrincipal }).(ServicePrincipalDefinitionPtrOutput)
}

func (o AzureBlobDefinitionOutput) SyncIntervalInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v AzureBlobDefinition) *float64 { return v.SyncIntervalInSeconds }).(pulumi.Float64PtrOutput)
}

func (o AzureBlobDefinitionOutput) TimeoutInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v AzureBlobDefinition) *float64 { return v.TimeoutInSeconds }).(pulumi.Float64PtrOutput)
}

func (o AzureBlobDefinitionOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureBlobDefinition) *string { return v.Url }).(pulumi.StringPtrOutput)
}

type AzureBlobDefinitionPtrOutput struct{ *pulumi.OutputState }

func (AzureBlobDefinitionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureBlobDefinition)(nil)).Elem()
}

func (o AzureBlobDefinitionPtrOutput) ToAzureBlobDefinitionPtrOutput() AzureBlobDefinitionPtrOutput {
	return o
}

func (o AzureBlobDefinitionPtrOutput) ToAzureBlobDefinitionPtrOutputWithContext(ctx context.Context) AzureBlobDefinitionPtrOutput {
	return o
}

func (o AzureBlobDefinitionPtrOutput) Elem() AzureBlobDefinitionOutput {
	return o.ApplyT(func(v *AzureBlobDefinition) AzureBlobDefinition {
		if v != nil {
			return *v
		}
		var ret AzureBlobDefinition
		return ret
	}).(AzureBlobDefinitionOutput)
}

func (o AzureBlobDefinitionPtrOutput) AccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureBlobDefinition) *string {
		if v == nil {
			return nil
		}
		return v.AccountKey
	}).(pulumi.StringPtrOutput)
}

func (o AzureBlobDefinitionPtrOutput) ContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureBlobDefinition) *string {
		if v == nil {
			return nil
		}
		return v.ContainerName
	}).(pulumi.StringPtrOutput)
}

func (o AzureBlobDefinitionPtrOutput) LocalAuthRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureBlobDefinition) *string {
		if v == nil {
			return nil
		}
		return v.LocalAuthRef
	}).(pulumi.StringPtrOutput)
}

func (o AzureBlobDefinitionPtrOutput) ManagedIdentity() ManagedIdentityDefinitionPtrOutput {
	return o.ApplyT(func(v *AzureBlobDefinition) *ManagedIdentityDefinition {
		if v == nil {
			return nil
		}
		return v.ManagedIdentity
	}).(ManagedIdentityDefinitionPtrOutput)
}

func (o AzureBlobDefinitionPtrOutput) SasToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureBlobDefinition) *string {
		if v == nil {
			return nil
		}
		return v.SasToken
	}).(pulumi.StringPtrOutput)
}

func (o AzureBlobDefinitionPtrOutput) ServicePrincipal() ServicePrincipalDefinitionPtrOutput {
	return o.ApplyT(func(v *AzureBlobDefinition) *ServicePrincipalDefinition {
		if v == nil {
			return nil
		}
		return v.ServicePrincipal
	}).(ServicePrincipalDefinitionPtrOutput)
}

func (o AzureBlobDefinitionPtrOutput) SyncIntervalInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *AzureBlobDefinition) *float64 {
		if v == nil {
			return nil
		}
		return v.SyncIntervalInSeconds
	}).(pulumi.Float64PtrOutput)
}

func (o AzureBlobDefinitionPtrOutput) TimeoutInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *AzureBlobDefinition) *float64 {
		if v == nil {
			return nil
		}
		return v.TimeoutInSeconds
	}).(pulumi.Float64PtrOutput)
}

func (o AzureBlobDefinitionPtrOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureBlobDefinition) *string {
		if v == nil {
			return nil
		}
		return v.Url
	}).(pulumi.StringPtrOutput)
}

type AzureBlobDefinitionResponse struct {
	AccountKey            *string                             `pulumi:"accountKey"`
	ContainerName         *string                             `pulumi:"containerName"`
	LocalAuthRef          *string                             `pulumi:"localAuthRef"`
	ManagedIdentity       *ManagedIdentityDefinitionResponse  `pulumi:"managedIdentity"`
	SasToken              *string                             `pulumi:"sasToken"`
	ServicePrincipal      *ServicePrincipalDefinitionResponse `pulumi:"servicePrincipal"`
	SyncIntervalInSeconds *float64                            `pulumi:"syncIntervalInSeconds"`
	TimeoutInSeconds      *float64                            `pulumi:"timeoutInSeconds"`
	Url                   *string                             `pulumi:"url"`
}


func (val *AzureBlobDefinitionResponse) Defaults() *AzureBlobDefinitionResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ServicePrincipal = tmp.ServicePrincipal.Defaults()

	if isZero(tmp.SyncIntervalInSeconds) {
		syncIntervalInSeconds_ := 600.0
		tmp.SyncIntervalInSeconds = &syncIntervalInSeconds_
	}
	if isZero(tmp.TimeoutInSeconds) {
		timeoutInSeconds_ := 600.0
		tmp.TimeoutInSeconds = &timeoutInSeconds_
	}
	return &tmp
}

type AzureBlobDefinitionResponseOutput struct{ *pulumi.OutputState }

func (AzureBlobDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureBlobDefinitionResponse)(nil)).Elem()
}

func (o AzureBlobDefinitionResponseOutput) ToAzureBlobDefinitionResponseOutput() AzureBlobDefinitionResponseOutput {
	return o
}

func (o AzureBlobDefinitionResponseOutput) ToAzureBlobDefinitionResponseOutputWithContext(ctx context.Context) AzureBlobDefinitionResponseOutput {
	return o
}

func (o AzureBlobDefinitionResponseOutput) AccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureBlobDefinitionResponse) *string { return v.AccountKey }).(pulumi.StringPtrOutput)
}

func (o AzureBlobDefinitionResponseOutput) ContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureBlobDefinitionResponse) *string { return v.ContainerName }).(pulumi.StringPtrOutput)
}

func (o AzureBlobDefinitionResponseOutput) LocalAuthRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureBlobDefinitionResponse) *string { return v.LocalAuthRef }).(pulumi.StringPtrOutput)
}

func (o AzureBlobDefinitionResponseOutput) ManagedIdentity() ManagedIdentityDefinitionResponsePtrOutput {
	return o.ApplyT(func(v AzureBlobDefinitionResponse) *ManagedIdentityDefinitionResponse { return v.ManagedIdentity }).(ManagedIdentityDefinitionResponsePtrOutput)
}

func (o AzureBlobDefinitionResponseOutput) SasToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureBlobDefinitionResponse) *string { return v.SasToken }).(pulumi.StringPtrOutput)
}

func (o AzureBlobDefinitionResponseOutput) ServicePrincipal() ServicePrincipalDefinitionResponsePtrOutput {
	return o.ApplyT(func(v AzureBlobDefinitionResponse) *ServicePrincipalDefinitionResponse { return v.ServicePrincipal }).(ServicePrincipalDefinitionResponsePtrOutput)
}

func (o AzureBlobDefinitionResponseOutput) SyncIntervalInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v AzureBlobDefinitionResponse) *float64 { return v.SyncIntervalInSeconds }).(pulumi.Float64PtrOutput)
}

func (o AzureBlobDefinitionResponseOutput) TimeoutInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v AzureBlobDefinitionResponse) *float64 { return v.TimeoutInSeconds }).(pulumi.Float64PtrOutput)
}

func (o AzureBlobDefinitionResponseOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureBlobDefinitionResponse) *string { return v.Url }).(pulumi.StringPtrOutput)
}

type AzureBlobDefinitionResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureBlobDefinitionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureBlobDefinitionResponse)(nil)).Elem()
}

func (o AzureBlobDefinitionResponsePtrOutput) ToAzureBlobDefinitionResponsePtrOutput() AzureBlobDefinitionResponsePtrOutput {
	return o
}

func (o AzureBlobDefinitionResponsePtrOutput) ToAzureBlobDefinitionResponsePtrOutputWithContext(ctx context.Context) AzureBlobDefinitionResponsePtrOutput {
	return o
}

func (o AzureBlobDefinitionResponsePtrOutput) Elem() AzureBlobDefinitionResponseOutput {
	return o.ApplyT(func(v *AzureBlobDefinitionResponse) AzureBlobDefinitionResponse {
		if v != nil {
			return *v
		}
		var ret AzureBlobDefinitionResponse
		return ret
	}).(AzureBlobDefinitionResponseOutput)
}

func (o AzureBlobDefinitionResponsePtrOutput) AccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureBlobDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.AccountKey
	}).(pulumi.StringPtrOutput)
}

func (o AzureBlobDefinitionResponsePtrOutput) ContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureBlobDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.ContainerName
	}).(pulumi.StringPtrOutput)
}

func (o AzureBlobDefinitionResponsePtrOutput) LocalAuthRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureBlobDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.LocalAuthRef
	}).(pulumi.StringPtrOutput)
}

func (o AzureBlobDefinitionResponsePtrOutput) ManagedIdentity() ManagedIdentityDefinitionResponsePtrOutput {
	return o.ApplyT(func(v *AzureBlobDefinitionResponse) *ManagedIdentityDefinitionResponse {
		if v == nil {
			return nil
		}
		return v.ManagedIdentity
	}).(ManagedIdentityDefinitionResponsePtrOutput)
}

func (o AzureBlobDefinitionResponsePtrOutput) SasToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureBlobDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.SasToken
	}).(pulumi.StringPtrOutput)
}

func (o AzureBlobDefinitionResponsePtrOutput) ServicePrincipal() ServicePrincipalDefinitionResponsePtrOutput {
	return o.ApplyT(func(v *AzureBlobDefinitionResponse) *ServicePrincipalDefinitionResponse {
		if v == nil {
			return nil
		}
		return v.ServicePrincipal
	}).(ServicePrincipalDefinitionResponsePtrOutput)
}

func (o AzureBlobDefinitionResponsePtrOutput) SyncIntervalInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *AzureBlobDefinitionResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.SyncIntervalInSeconds
	}).(pulumi.Float64PtrOutput)
}

func (o AzureBlobDefinitionResponsePtrOutput) TimeoutInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *AzureBlobDefinitionResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.TimeoutInSeconds
	}).(pulumi.Float64PtrOutput)
}

func (o AzureBlobDefinitionResponsePtrOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureBlobDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Url
	}).(pulumi.StringPtrOutput)
}

type BucketDefinition struct {
	AccessKey             *string  `pulumi:"accessKey"`
	BucketName            *string  `pulumi:"bucketName"`
	Insecure              *bool    `pulumi:"insecure"`
	LocalAuthRef          *string  `pulumi:"localAuthRef"`
	SyncIntervalInSeconds *float64 `pulumi:"syncIntervalInSeconds"`
	TimeoutInSeconds      *float64 `pulumi:"timeoutInSeconds"`
	Url                   *string  `pulumi:"url"`
}


func (val *BucketDefinition) Defaults() *BucketDefinition {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Insecure) {
		insecure_ := true
		tmp.Insecure = &insecure_
	}
	if isZero(tmp.SyncIntervalInSeconds) {
		syncIntervalInSeconds_ := 600.0
		tmp.SyncIntervalInSeconds = &syncIntervalInSeconds_
	}
	if isZero(tmp.TimeoutInSeconds) {
		timeoutInSeconds_ := 600.0
		tmp.TimeoutInSeconds = &timeoutInSeconds_
	}
	return &tmp
}





type BucketDefinitionInput interface {
	pulumi.Input

	ToBucketDefinitionOutput() BucketDefinitionOutput
	ToBucketDefinitionOutputWithContext(context.Context) BucketDefinitionOutput
}

type BucketDefinitionArgs struct {
	AccessKey             pulumi.StringPtrInput  `pulumi:"accessKey"`
	BucketName            pulumi.StringPtrInput  `pulumi:"bucketName"`
	Insecure              pulumi.BoolPtrInput    `pulumi:"insecure"`
	LocalAuthRef          pulumi.StringPtrInput  `pulumi:"localAuthRef"`
	SyncIntervalInSeconds pulumi.Float64PtrInput `pulumi:"syncIntervalInSeconds"`
	TimeoutInSeconds      pulumi.Float64PtrInput `pulumi:"timeoutInSeconds"`
	Url                   pulumi.StringPtrInput  `pulumi:"url"`
}


func (val *BucketDefinitionArgs) Defaults() *BucketDefinitionArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Insecure) {
		tmp.Insecure = pulumi.BoolPtr(true)
	}
	if isZero(tmp.SyncIntervalInSeconds) {
		tmp.SyncIntervalInSeconds = pulumi.Float64Ptr(600.0)
	}
	if isZero(tmp.TimeoutInSeconds) {
		tmp.TimeoutInSeconds = pulumi.Float64Ptr(600.0)
	}
	return &tmp
}
func (BucketDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketDefinition)(nil)).Elem()
}

func (i BucketDefinitionArgs) ToBucketDefinitionOutput() BucketDefinitionOutput {
	return i.ToBucketDefinitionOutputWithContext(context.Background())
}

func (i BucketDefinitionArgs) ToBucketDefinitionOutputWithContext(ctx context.Context) BucketDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketDefinitionOutput)
}

func (i BucketDefinitionArgs) ToBucketDefinitionPtrOutput() BucketDefinitionPtrOutput {
	return i.ToBucketDefinitionPtrOutputWithContext(context.Background())
}

func (i BucketDefinitionArgs) ToBucketDefinitionPtrOutputWithContext(ctx context.Context) BucketDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketDefinitionOutput).ToBucketDefinitionPtrOutputWithContext(ctx)
}









type BucketDefinitionPtrInput interface {
	pulumi.Input

	ToBucketDefinitionPtrOutput() BucketDefinitionPtrOutput
	ToBucketDefinitionPtrOutputWithContext(context.Context) BucketDefinitionPtrOutput
}

type bucketDefinitionPtrType BucketDefinitionArgs

func BucketDefinitionPtr(v *BucketDefinitionArgs) BucketDefinitionPtrInput {
	return (*bucketDefinitionPtrType)(v)
}

func (*bucketDefinitionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketDefinition)(nil)).Elem()
}

func (i *bucketDefinitionPtrType) ToBucketDefinitionPtrOutput() BucketDefinitionPtrOutput {
	return i.ToBucketDefinitionPtrOutputWithContext(context.Background())
}

func (i *bucketDefinitionPtrType) ToBucketDefinitionPtrOutputWithContext(ctx context.Context) BucketDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketDefinitionPtrOutput)
}

type BucketDefinitionOutput struct{ *pulumi.OutputState }

func (BucketDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketDefinition)(nil)).Elem()
}

func (o BucketDefinitionOutput) ToBucketDefinitionOutput() BucketDefinitionOutput {
	return o
}

func (o BucketDefinitionOutput) ToBucketDefinitionOutputWithContext(ctx context.Context) BucketDefinitionOutput {
	return o
}

func (o BucketDefinitionOutput) ToBucketDefinitionPtrOutput() BucketDefinitionPtrOutput {
	return o.ToBucketDefinitionPtrOutputWithContext(context.Background())
}

func (o BucketDefinitionOutput) ToBucketDefinitionPtrOutputWithContext(ctx context.Context) BucketDefinitionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BucketDefinition) *BucketDefinition {
		return &v
	}).(BucketDefinitionPtrOutput)
}

func (o BucketDefinitionOutput) AccessKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BucketDefinition) *string { return v.AccessKey }).(pulumi.StringPtrOutput)
}

func (o BucketDefinitionOutput) BucketName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BucketDefinition) *string { return v.BucketName }).(pulumi.StringPtrOutput)
}

func (o BucketDefinitionOutput) Insecure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v BucketDefinition) *bool { return v.Insecure }).(pulumi.BoolPtrOutput)
}

func (o BucketDefinitionOutput) LocalAuthRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BucketDefinition) *string { return v.LocalAuthRef }).(pulumi.StringPtrOutput)
}

func (o BucketDefinitionOutput) SyncIntervalInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v BucketDefinition) *float64 { return v.SyncIntervalInSeconds }).(pulumi.Float64PtrOutput)
}

func (o BucketDefinitionOutput) TimeoutInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v BucketDefinition) *float64 { return v.TimeoutInSeconds }).(pulumi.Float64PtrOutput)
}

func (o BucketDefinitionOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BucketDefinition) *string { return v.Url }).(pulumi.StringPtrOutput)
}

type BucketDefinitionPtrOutput struct{ *pulumi.OutputState }

func (BucketDefinitionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketDefinition)(nil)).Elem()
}

func (o BucketDefinitionPtrOutput) ToBucketDefinitionPtrOutput() BucketDefinitionPtrOutput {
	return o
}

func (o BucketDefinitionPtrOutput) ToBucketDefinitionPtrOutputWithContext(ctx context.Context) BucketDefinitionPtrOutput {
	return o
}

func (o BucketDefinitionPtrOutput) Elem() BucketDefinitionOutput {
	return o.ApplyT(func(v *BucketDefinition) BucketDefinition {
		if v != nil {
			return *v
		}
		var ret BucketDefinition
		return ret
	}).(BucketDefinitionOutput)
}

func (o BucketDefinitionPtrOutput) AccessKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BucketDefinition) *string {
		if v == nil {
			return nil
		}
		return v.AccessKey
	}).(pulumi.StringPtrOutput)
}

func (o BucketDefinitionPtrOutput) BucketName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BucketDefinition) *string {
		if v == nil {
			return nil
		}
		return v.BucketName
	}).(pulumi.StringPtrOutput)
}

func (o BucketDefinitionPtrOutput) Insecure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BucketDefinition) *bool {
		if v == nil {
			return nil
		}
		return v.Insecure
	}).(pulumi.BoolPtrOutput)
}

func (o BucketDefinitionPtrOutput) LocalAuthRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BucketDefinition) *string {
		if v == nil {
			return nil
		}
		return v.LocalAuthRef
	}).(pulumi.StringPtrOutput)
}

func (o BucketDefinitionPtrOutput) SyncIntervalInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *BucketDefinition) *float64 {
		if v == nil {
			return nil
		}
		return v.SyncIntervalInSeconds
	}).(pulumi.Float64PtrOutput)
}

func (o BucketDefinitionPtrOutput) TimeoutInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *BucketDefinition) *float64 {
		if v == nil {
			return nil
		}
		return v.TimeoutInSeconds
	}).(pulumi.Float64PtrOutput)
}

func (o BucketDefinitionPtrOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BucketDefinition) *string {
		if v == nil {
			return nil
		}
		return v.Url
	}).(pulumi.StringPtrOutput)
}

type BucketDefinitionResponse struct {
	AccessKey             *string  `pulumi:"accessKey"`
	BucketName            *string  `pulumi:"bucketName"`
	Insecure              *bool    `pulumi:"insecure"`
	LocalAuthRef          *string  `pulumi:"localAuthRef"`
	SyncIntervalInSeconds *float64 `pulumi:"syncIntervalInSeconds"`
	TimeoutInSeconds      *float64 `pulumi:"timeoutInSeconds"`
	Url                   *string  `pulumi:"url"`
}


func (val *BucketDefinitionResponse) Defaults() *BucketDefinitionResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Insecure) {
		insecure_ := true
		tmp.Insecure = &insecure_
	}
	if isZero(tmp.SyncIntervalInSeconds) {
		syncIntervalInSeconds_ := 600.0
		tmp.SyncIntervalInSeconds = &syncIntervalInSeconds_
	}
	if isZero(tmp.TimeoutInSeconds) {
		timeoutInSeconds_ := 600.0
		tmp.TimeoutInSeconds = &timeoutInSeconds_
	}
	return &tmp
}

type BucketDefinitionResponseOutput struct{ *pulumi.OutputState }

func (BucketDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketDefinitionResponse)(nil)).Elem()
}

func (o BucketDefinitionResponseOutput) ToBucketDefinitionResponseOutput() BucketDefinitionResponseOutput {
	return o
}

func (o BucketDefinitionResponseOutput) ToBucketDefinitionResponseOutputWithContext(ctx context.Context) BucketDefinitionResponseOutput {
	return o
}

func (o BucketDefinitionResponseOutput) AccessKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BucketDefinitionResponse) *string { return v.AccessKey }).(pulumi.StringPtrOutput)
}

func (o BucketDefinitionResponseOutput) BucketName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BucketDefinitionResponse) *string { return v.BucketName }).(pulumi.StringPtrOutput)
}

func (o BucketDefinitionResponseOutput) Insecure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v BucketDefinitionResponse) *bool { return v.Insecure }).(pulumi.BoolPtrOutput)
}

func (o BucketDefinitionResponseOutput) LocalAuthRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BucketDefinitionResponse) *string { return v.LocalAuthRef }).(pulumi.StringPtrOutput)
}

func (o BucketDefinitionResponseOutput) SyncIntervalInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v BucketDefinitionResponse) *float64 { return v.SyncIntervalInSeconds }).(pulumi.Float64PtrOutput)
}

func (o BucketDefinitionResponseOutput) TimeoutInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v BucketDefinitionResponse) *float64 { return v.TimeoutInSeconds }).(pulumi.Float64PtrOutput)
}

func (o BucketDefinitionResponseOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BucketDefinitionResponse) *string { return v.Url }).(pulumi.StringPtrOutput)
}

type BucketDefinitionResponsePtrOutput struct{ *pulumi.OutputState }

func (BucketDefinitionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketDefinitionResponse)(nil)).Elem()
}

func (o BucketDefinitionResponsePtrOutput) ToBucketDefinitionResponsePtrOutput() BucketDefinitionResponsePtrOutput {
	return o
}

func (o BucketDefinitionResponsePtrOutput) ToBucketDefinitionResponsePtrOutputWithContext(ctx context.Context) BucketDefinitionResponsePtrOutput {
	return o
}

func (o BucketDefinitionResponsePtrOutput) Elem() BucketDefinitionResponseOutput {
	return o.ApplyT(func(v *BucketDefinitionResponse) BucketDefinitionResponse {
		if v != nil {
			return *v
		}
		var ret BucketDefinitionResponse
		return ret
	}).(BucketDefinitionResponseOutput)
}

func (o BucketDefinitionResponsePtrOutput) AccessKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BucketDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.AccessKey
	}).(pulumi.StringPtrOutput)
}

func (o BucketDefinitionResponsePtrOutput) BucketName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BucketDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.BucketName
	}).(pulumi.StringPtrOutput)
}

func (o BucketDefinitionResponsePtrOutput) Insecure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BucketDefinitionResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Insecure
	}).(pulumi.BoolPtrOutput)
}

func (o BucketDefinitionResponsePtrOutput) LocalAuthRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BucketDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.LocalAuthRef
	}).(pulumi.StringPtrOutput)
}

func (o BucketDefinitionResponsePtrOutput) SyncIntervalInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *BucketDefinitionResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.SyncIntervalInSeconds
	}).(pulumi.Float64PtrOutput)
}

func (o BucketDefinitionResponsePtrOutput) TimeoutInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *BucketDefinitionResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.TimeoutInSeconds
	}).(pulumi.Float64PtrOutput)
}

func (o BucketDefinitionResponsePtrOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BucketDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Url
	}).(pulumi.StringPtrOutput)
}

type ComplianceStatusResponse struct {
	ComplianceState   string  `pulumi:"complianceState"`
	LastConfigApplied *string `pulumi:"lastConfigApplied"`
	Message           *string `pulumi:"message"`
	MessageLevel      *string `pulumi:"messageLevel"`
}

type ComplianceStatusResponseOutput struct{ *pulumi.OutputState }

func (ComplianceStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComplianceStatusResponse)(nil)).Elem()
}

func (o ComplianceStatusResponseOutput) ToComplianceStatusResponseOutput() ComplianceStatusResponseOutput {
	return o
}

func (o ComplianceStatusResponseOutput) ToComplianceStatusResponseOutputWithContext(ctx context.Context) ComplianceStatusResponseOutput {
	return o
}

func (o ComplianceStatusResponseOutput) ComplianceState() pulumi.StringOutput {
	return o.ApplyT(func(v ComplianceStatusResponse) string { return v.ComplianceState }).(pulumi.StringOutput)
}

func (o ComplianceStatusResponseOutput) LastConfigApplied() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComplianceStatusResponse) *string { return v.LastConfigApplied }).(pulumi.StringPtrOutput)
}

func (o ComplianceStatusResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComplianceStatusResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o ComplianceStatusResponseOutput) MessageLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComplianceStatusResponse) *string { return v.MessageLevel }).(pulumi.StringPtrOutput)
}

type ErrorAdditionalInfoResponse struct {
	Info interface{} `pulumi:"info"`
	Type string      `pulumi:"type"`
}

type ErrorAdditionalInfoResponseOutput struct{ *pulumi.OutputState }

func (ErrorAdditionalInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorAdditionalInfoResponse)(nil)).Elem()
}

func (o ErrorAdditionalInfoResponseOutput) ToErrorAdditionalInfoResponseOutput() ErrorAdditionalInfoResponseOutput {
	return o
}

func (o ErrorAdditionalInfoResponseOutput) ToErrorAdditionalInfoResponseOutputWithContext(ctx context.Context) ErrorAdditionalInfoResponseOutput {
	return o
}

func (o ErrorAdditionalInfoResponseOutput) Info() pulumi.AnyOutput {
	return o.ApplyT(func(v ErrorAdditionalInfoResponse) interface{} { return v.Info }).(pulumi.AnyOutput)
}

func (o ErrorAdditionalInfoResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorAdditionalInfoResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ErrorAdditionalInfoResponseArrayOutput struct{ *pulumi.OutputState }

func (ErrorAdditionalInfoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorAdditionalInfoResponse)(nil)).Elem()
}

func (o ErrorAdditionalInfoResponseArrayOutput) ToErrorAdditionalInfoResponseArrayOutput() ErrorAdditionalInfoResponseArrayOutput {
	return o
}

func (o ErrorAdditionalInfoResponseArrayOutput) ToErrorAdditionalInfoResponseArrayOutputWithContext(ctx context.Context) ErrorAdditionalInfoResponseArrayOutput {
	return o
}

func (o ErrorAdditionalInfoResponseArrayOutput) Index(i pulumi.IntInput) ErrorAdditionalInfoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ErrorAdditionalInfoResponse {
		return vs[0].([]ErrorAdditionalInfoResponse)[vs[1].(int)]
	}).(ErrorAdditionalInfoResponseOutput)
}

type ErrorDetailResponse struct {
	AdditionalInfo []ErrorAdditionalInfoResponse `pulumi:"additionalInfo"`
	Code           string                        `pulumi:"code"`
	Details        []ErrorDetailResponse         `pulumi:"details"`
	Message        string                        `pulumi:"message"`
	Target         string                        `pulumi:"target"`
}

type ErrorDetailResponseOutput struct{ *pulumi.OutputState }

func (ErrorDetailResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorDetailResponse)(nil)).Elem()
}

func (o ErrorDetailResponseOutput) ToErrorDetailResponseOutput() ErrorDetailResponseOutput {
	return o
}

func (o ErrorDetailResponseOutput) ToErrorDetailResponseOutputWithContext(ctx context.Context) ErrorDetailResponseOutput {
	return o
}

func (o ErrorDetailResponseOutput) AdditionalInfo() ErrorAdditionalInfoResponseArrayOutput {
	return o.ApplyT(func(v ErrorDetailResponse) []ErrorAdditionalInfoResponse { return v.AdditionalInfo }).(ErrorAdditionalInfoResponseArrayOutput)
}

func (o ErrorDetailResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDetailResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o ErrorDetailResponseOutput) Details() ErrorDetailResponseArrayOutput {
	return o.ApplyT(func(v ErrorDetailResponse) []ErrorDetailResponse { return v.Details }).(ErrorDetailResponseArrayOutput)
}

func (o ErrorDetailResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDetailResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o ErrorDetailResponseOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDetailResponse) string { return v.Target }).(pulumi.StringOutput)
}

type ErrorDetailResponseArrayOutput struct{ *pulumi.OutputState }

func (ErrorDetailResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorDetailResponse)(nil)).Elem()
}

func (o ErrorDetailResponseArrayOutput) ToErrorDetailResponseArrayOutput() ErrorDetailResponseArrayOutput {
	return o
}

func (o ErrorDetailResponseArrayOutput) ToErrorDetailResponseArrayOutputWithContext(ctx context.Context) ErrorDetailResponseArrayOutput {
	return o
}

func (o ErrorDetailResponseArrayOutput) Index(i pulumi.IntInput) ErrorDetailResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ErrorDetailResponse {
		return vs[0].([]ErrorDetailResponse)[vs[1].(int)]
	}).(ErrorDetailResponseOutput)
}

type ExtensionAksAssignedIdentity struct {
	Type *AKSIdentityType `pulumi:"type"`
}





type ExtensionAksAssignedIdentityInput interface {
	pulumi.Input

	ToExtensionAksAssignedIdentityOutput() ExtensionAksAssignedIdentityOutput
	ToExtensionAksAssignedIdentityOutputWithContext(context.Context) ExtensionAksAssignedIdentityOutput
}

type ExtensionAksAssignedIdentityArgs struct {
	Type AKSIdentityTypePtrInput `pulumi:"type"`
}

func (ExtensionAksAssignedIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtensionAksAssignedIdentity)(nil)).Elem()
}

func (i ExtensionAksAssignedIdentityArgs) ToExtensionAksAssignedIdentityOutput() ExtensionAksAssignedIdentityOutput {
	return i.ToExtensionAksAssignedIdentityOutputWithContext(context.Background())
}

func (i ExtensionAksAssignedIdentityArgs) ToExtensionAksAssignedIdentityOutputWithContext(ctx context.Context) ExtensionAksAssignedIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionAksAssignedIdentityOutput)
}

func (i ExtensionAksAssignedIdentityArgs) ToExtensionAksAssignedIdentityPtrOutput() ExtensionAksAssignedIdentityPtrOutput {
	return i.ToExtensionAksAssignedIdentityPtrOutputWithContext(context.Background())
}

func (i ExtensionAksAssignedIdentityArgs) ToExtensionAksAssignedIdentityPtrOutputWithContext(ctx context.Context) ExtensionAksAssignedIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionAksAssignedIdentityOutput).ToExtensionAksAssignedIdentityPtrOutputWithContext(ctx)
}









type ExtensionAksAssignedIdentityPtrInput interface {
	pulumi.Input

	ToExtensionAksAssignedIdentityPtrOutput() ExtensionAksAssignedIdentityPtrOutput
	ToExtensionAksAssignedIdentityPtrOutputWithContext(context.Context) ExtensionAksAssignedIdentityPtrOutput
}

type extensionAksAssignedIdentityPtrType ExtensionAksAssignedIdentityArgs

func ExtensionAksAssignedIdentityPtr(v *ExtensionAksAssignedIdentityArgs) ExtensionAksAssignedIdentityPtrInput {
	return (*extensionAksAssignedIdentityPtrType)(v)
}

func (*extensionAksAssignedIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtensionAksAssignedIdentity)(nil)).Elem()
}

func (i *extensionAksAssignedIdentityPtrType) ToExtensionAksAssignedIdentityPtrOutput() ExtensionAksAssignedIdentityPtrOutput {
	return i.ToExtensionAksAssignedIdentityPtrOutputWithContext(context.Background())
}

func (i *extensionAksAssignedIdentityPtrType) ToExtensionAksAssignedIdentityPtrOutputWithContext(ctx context.Context) ExtensionAksAssignedIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionAksAssignedIdentityPtrOutput)
}

type ExtensionAksAssignedIdentityOutput struct{ *pulumi.OutputState }

func (ExtensionAksAssignedIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtensionAksAssignedIdentity)(nil)).Elem()
}

func (o ExtensionAksAssignedIdentityOutput) ToExtensionAksAssignedIdentityOutput() ExtensionAksAssignedIdentityOutput {
	return o
}

func (o ExtensionAksAssignedIdentityOutput) ToExtensionAksAssignedIdentityOutputWithContext(ctx context.Context) ExtensionAksAssignedIdentityOutput {
	return o
}

func (o ExtensionAksAssignedIdentityOutput) ToExtensionAksAssignedIdentityPtrOutput() ExtensionAksAssignedIdentityPtrOutput {
	return o.ToExtensionAksAssignedIdentityPtrOutputWithContext(context.Background())
}

func (o ExtensionAksAssignedIdentityOutput) ToExtensionAksAssignedIdentityPtrOutputWithContext(ctx context.Context) ExtensionAksAssignedIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExtensionAksAssignedIdentity) *ExtensionAksAssignedIdentity {
		return &v
	}).(ExtensionAksAssignedIdentityPtrOutput)
}

func (o ExtensionAksAssignedIdentityOutput) Type() AKSIdentityTypePtrOutput {
	return o.ApplyT(func(v ExtensionAksAssignedIdentity) *AKSIdentityType { return v.Type }).(AKSIdentityTypePtrOutput)
}

type ExtensionAksAssignedIdentityPtrOutput struct{ *pulumi.OutputState }

func (ExtensionAksAssignedIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtensionAksAssignedIdentity)(nil)).Elem()
}

func (o ExtensionAksAssignedIdentityPtrOutput) ToExtensionAksAssignedIdentityPtrOutput() ExtensionAksAssignedIdentityPtrOutput {
	return o
}

func (o ExtensionAksAssignedIdentityPtrOutput) ToExtensionAksAssignedIdentityPtrOutputWithContext(ctx context.Context) ExtensionAksAssignedIdentityPtrOutput {
	return o
}

func (o ExtensionAksAssignedIdentityPtrOutput) Elem() ExtensionAksAssignedIdentityOutput {
	return o.ApplyT(func(v *ExtensionAksAssignedIdentity) ExtensionAksAssignedIdentity {
		if v != nil {
			return *v
		}
		var ret ExtensionAksAssignedIdentity
		return ret
	}).(ExtensionAksAssignedIdentityOutput)
}

func (o ExtensionAksAssignedIdentityPtrOutput) Type() AKSIdentityTypePtrOutput {
	return o.ApplyT(func(v *ExtensionAksAssignedIdentity) *AKSIdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(AKSIdentityTypePtrOutput)
}

type ExtensionResponseAksAssignedIdentity struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}

type ExtensionResponseAksAssignedIdentityOutput struct{ *pulumi.OutputState }

func (ExtensionResponseAksAssignedIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtensionResponseAksAssignedIdentity)(nil)).Elem()
}

func (o ExtensionResponseAksAssignedIdentityOutput) ToExtensionResponseAksAssignedIdentityOutput() ExtensionResponseAksAssignedIdentityOutput {
	return o
}

func (o ExtensionResponseAksAssignedIdentityOutput) ToExtensionResponseAksAssignedIdentityOutputWithContext(ctx context.Context) ExtensionResponseAksAssignedIdentityOutput {
	return o
}

func (o ExtensionResponseAksAssignedIdentityOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ExtensionResponseAksAssignedIdentity) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ExtensionResponseAksAssignedIdentityOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ExtensionResponseAksAssignedIdentity) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ExtensionResponseAksAssignedIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionResponseAksAssignedIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ExtensionResponseAksAssignedIdentityPtrOutput struct{ *pulumi.OutputState }

func (ExtensionResponseAksAssignedIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtensionResponseAksAssignedIdentity)(nil)).Elem()
}

func (o ExtensionResponseAksAssignedIdentityPtrOutput) ToExtensionResponseAksAssignedIdentityPtrOutput() ExtensionResponseAksAssignedIdentityPtrOutput {
	return o
}

func (o ExtensionResponseAksAssignedIdentityPtrOutput) ToExtensionResponseAksAssignedIdentityPtrOutputWithContext(ctx context.Context) ExtensionResponseAksAssignedIdentityPtrOutput {
	return o
}

func (o ExtensionResponseAksAssignedIdentityPtrOutput) Elem() ExtensionResponseAksAssignedIdentityOutput {
	return o.ApplyT(func(v *ExtensionResponseAksAssignedIdentity) ExtensionResponseAksAssignedIdentity {
		if v != nil {
			return *v
		}
		var ret ExtensionResponseAksAssignedIdentity
		return ret
	}).(ExtensionResponseAksAssignedIdentityOutput)
}

func (o ExtensionResponseAksAssignedIdentityPtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtensionResponseAksAssignedIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ExtensionResponseAksAssignedIdentityPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtensionResponseAksAssignedIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ExtensionResponseAksAssignedIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtensionResponseAksAssignedIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ExtensionStatus struct {
	Code          *string `pulumi:"code"`
	DisplayStatus *string `pulumi:"displayStatus"`
	Level         *string `pulumi:"level"`
	Message       *string `pulumi:"message"`
	Time          *string `pulumi:"time"`
}


func (val *ExtensionStatus) Defaults() *ExtensionStatus {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Level) {
		level_ := "Information"
		tmp.Level = &level_
	}
	return &tmp
}





type ExtensionStatusInput interface {
	pulumi.Input

	ToExtensionStatusOutput() ExtensionStatusOutput
	ToExtensionStatusOutputWithContext(context.Context) ExtensionStatusOutput
}

type ExtensionStatusArgs struct {
	Code          pulumi.StringPtrInput `pulumi:"code"`
	DisplayStatus pulumi.StringPtrInput `pulumi:"displayStatus"`
	Level         pulumi.StringPtrInput `pulumi:"level"`
	Message       pulumi.StringPtrInput `pulumi:"message"`
	Time          pulumi.StringPtrInput `pulumi:"time"`
}


func (val *ExtensionStatusArgs) Defaults() *ExtensionStatusArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Level) {
		tmp.Level = pulumi.StringPtr("Information")
	}
	return &tmp
}
func (ExtensionStatusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtensionStatus)(nil)).Elem()
}

func (i ExtensionStatusArgs) ToExtensionStatusOutput() ExtensionStatusOutput {
	return i.ToExtensionStatusOutputWithContext(context.Background())
}

func (i ExtensionStatusArgs) ToExtensionStatusOutputWithContext(ctx context.Context) ExtensionStatusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionStatusOutput)
}





type ExtensionStatusArrayInput interface {
	pulumi.Input

	ToExtensionStatusArrayOutput() ExtensionStatusArrayOutput
	ToExtensionStatusArrayOutputWithContext(context.Context) ExtensionStatusArrayOutput
}

type ExtensionStatusArray []ExtensionStatusInput

func (ExtensionStatusArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExtensionStatus)(nil)).Elem()
}

func (i ExtensionStatusArray) ToExtensionStatusArrayOutput() ExtensionStatusArrayOutput {
	return i.ToExtensionStatusArrayOutputWithContext(context.Background())
}

func (i ExtensionStatusArray) ToExtensionStatusArrayOutputWithContext(ctx context.Context) ExtensionStatusArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionStatusArrayOutput)
}

type ExtensionStatusOutput struct{ *pulumi.OutputState }

func (ExtensionStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtensionStatus)(nil)).Elem()
}

func (o ExtensionStatusOutput) ToExtensionStatusOutput() ExtensionStatusOutput {
	return o
}

func (o ExtensionStatusOutput) ToExtensionStatusOutputWithContext(ctx context.Context) ExtensionStatusOutput {
	return o
}

func (o ExtensionStatusOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionStatus) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o ExtensionStatusOutput) DisplayStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionStatus) *string { return v.DisplayStatus }).(pulumi.StringPtrOutput)
}

func (o ExtensionStatusOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionStatus) *string { return v.Level }).(pulumi.StringPtrOutput)
}

func (o ExtensionStatusOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionStatus) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o ExtensionStatusOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionStatus) *string { return v.Time }).(pulumi.StringPtrOutput)
}

type ExtensionStatusArrayOutput struct{ *pulumi.OutputState }

func (ExtensionStatusArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExtensionStatus)(nil)).Elem()
}

func (o ExtensionStatusArrayOutput) ToExtensionStatusArrayOutput() ExtensionStatusArrayOutput {
	return o
}

func (o ExtensionStatusArrayOutput) ToExtensionStatusArrayOutputWithContext(ctx context.Context) ExtensionStatusArrayOutput {
	return o
}

func (o ExtensionStatusArrayOutput) Index(i pulumi.IntInput) ExtensionStatusOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExtensionStatus {
		return vs[0].([]ExtensionStatus)[vs[1].(int)]
	}).(ExtensionStatusOutput)
}

type ExtensionStatusResponse struct {
	Code          *string `pulumi:"code"`
	DisplayStatus *string `pulumi:"displayStatus"`
	Level         *string `pulumi:"level"`
	Message       *string `pulumi:"message"`
	Time          *string `pulumi:"time"`
}


func (val *ExtensionStatusResponse) Defaults() *ExtensionStatusResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Level) {
		level_ := "Information"
		tmp.Level = &level_
	}
	return &tmp
}

type ExtensionStatusResponseOutput struct{ *pulumi.OutputState }

func (ExtensionStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtensionStatusResponse)(nil)).Elem()
}

func (o ExtensionStatusResponseOutput) ToExtensionStatusResponseOutput() ExtensionStatusResponseOutput {
	return o
}

func (o ExtensionStatusResponseOutput) ToExtensionStatusResponseOutputWithContext(ctx context.Context) ExtensionStatusResponseOutput {
	return o
}

func (o ExtensionStatusResponseOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionStatusResponse) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o ExtensionStatusResponseOutput) DisplayStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionStatusResponse) *string { return v.DisplayStatus }).(pulumi.StringPtrOutput)
}

func (o ExtensionStatusResponseOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionStatusResponse) *string { return v.Level }).(pulumi.StringPtrOutput)
}

func (o ExtensionStatusResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionStatusResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o ExtensionStatusResponseOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtensionStatusResponse) *string { return v.Time }).(pulumi.StringPtrOutput)
}

type ExtensionStatusResponseArrayOutput struct{ *pulumi.OutputState }

func (ExtensionStatusResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExtensionStatusResponse)(nil)).Elem()
}

func (o ExtensionStatusResponseArrayOutput) ToExtensionStatusResponseArrayOutput() ExtensionStatusResponseArrayOutput {
	return o
}

func (o ExtensionStatusResponseArrayOutput) ToExtensionStatusResponseArrayOutputWithContext(ctx context.Context) ExtensionStatusResponseArrayOutput {
	return o
}

func (o ExtensionStatusResponseArrayOutput) Index(i pulumi.IntInput) ExtensionStatusResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExtensionStatusResponse {
		return vs[0].([]ExtensionStatusResponse)[vs[1].(int)]
	}).(ExtensionStatusResponseOutput)
}

type GitRepositoryDefinition struct {
	HttpsCACert           *string                  `pulumi:"httpsCACert"`
	HttpsUser             *string                  `pulumi:"httpsUser"`
	LocalAuthRef          *string                  `pulumi:"localAuthRef"`
	RepositoryRef         *RepositoryRefDefinition `pulumi:"repositoryRef"`
	SshKnownHosts         *string                  `pulumi:"sshKnownHosts"`
	SyncIntervalInSeconds *float64                 `pulumi:"syncIntervalInSeconds"`
	TimeoutInSeconds      *float64                 `pulumi:"timeoutInSeconds"`
	Url                   *string                  `pulumi:"url"`
}


func (val *GitRepositoryDefinition) Defaults() *GitRepositoryDefinition {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.SyncIntervalInSeconds) {
		syncIntervalInSeconds_ := 600.0
		tmp.SyncIntervalInSeconds = &syncIntervalInSeconds_
	}
	if isZero(tmp.TimeoutInSeconds) {
		timeoutInSeconds_ := 600.0
		tmp.TimeoutInSeconds = &timeoutInSeconds_
	}
	return &tmp
}





type GitRepositoryDefinitionInput interface {
	pulumi.Input

	ToGitRepositoryDefinitionOutput() GitRepositoryDefinitionOutput
	ToGitRepositoryDefinitionOutputWithContext(context.Context) GitRepositoryDefinitionOutput
}

type GitRepositoryDefinitionArgs struct {
	HttpsCACert           pulumi.StringPtrInput           `pulumi:"httpsCACert"`
	HttpsUser             pulumi.StringPtrInput           `pulumi:"httpsUser"`
	LocalAuthRef          pulumi.StringPtrInput           `pulumi:"localAuthRef"`
	RepositoryRef         RepositoryRefDefinitionPtrInput `pulumi:"repositoryRef"`
	SshKnownHosts         pulumi.StringPtrInput           `pulumi:"sshKnownHosts"`
	SyncIntervalInSeconds pulumi.Float64PtrInput          `pulumi:"syncIntervalInSeconds"`
	TimeoutInSeconds      pulumi.Float64PtrInput          `pulumi:"timeoutInSeconds"`
	Url                   pulumi.StringPtrInput           `pulumi:"url"`
}


func (val *GitRepositoryDefinitionArgs) Defaults() *GitRepositoryDefinitionArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.SyncIntervalInSeconds) {
		tmp.SyncIntervalInSeconds = pulumi.Float64Ptr(600.0)
	}
	if isZero(tmp.TimeoutInSeconds) {
		tmp.TimeoutInSeconds = pulumi.Float64Ptr(600.0)
	}
	return &tmp
}
func (GitRepositoryDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GitRepositoryDefinition)(nil)).Elem()
}

func (i GitRepositoryDefinitionArgs) ToGitRepositoryDefinitionOutput() GitRepositoryDefinitionOutput {
	return i.ToGitRepositoryDefinitionOutputWithContext(context.Background())
}

func (i GitRepositoryDefinitionArgs) ToGitRepositoryDefinitionOutputWithContext(ctx context.Context) GitRepositoryDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitRepositoryDefinitionOutput)
}

func (i GitRepositoryDefinitionArgs) ToGitRepositoryDefinitionPtrOutput() GitRepositoryDefinitionPtrOutput {
	return i.ToGitRepositoryDefinitionPtrOutputWithContext(context.Background())
}

func (i GitRepositoryDefinitionArgs) ToGitRepositoryDefinitionPtrOutputWithContext(ctx context.Context) GitRepositoryDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitRepositoryDefinitionOutput).ToGitRepositoryDefinitionPtrOutputWithContext(ctx)
}









type GitRepositoryDefinitionPtrInput interface {
	pulumi.Input

	ToGitRepositoryDefinitionPtrOutput() GitRepositoryDefinitionPtrOutput
	ToGitRepositoryDefinitionPtrOutputWithContext(context.Context) GitRepositoryDefinitionPtrOutput
}

type gitRepositoryDefinitionPtrType GitRepositoryDefinitionArgs

func GitRepositoryDefinitionPtr(v *GitRepositoryDefinitionArgs) GitRepositoryDefinitionPtrInput {
	return (*gitRepositoryDefinitionPtrType)(v)
}

func (*gitRepositoryDefinitionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GitRepositoryDefinition)(nil)).Elem()
}

func (i *gitRepositoryDefinitionPtrType) ToGitRepositoryDefinitionPtrOutput() GitRepositoryDefinitionPtrOutput {
	return i.ToGitRepositoryDefinitionPtrOutputWithContext(context.Background())
}

func (i *gitRepositoryDefinitionPtrType) ToGitRepositoryDefinitionPtrOutputWithContext(ctx context.Context) GitRepositoryDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitRepositoryDefinitionPtrOutput)
}

type GitRepositoryDefinitionOutput struct{ *pulumi.OutputState }

func (GitRepositoryDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GitRepositoryDefinition)(nil)).Elem()
}

func (o GitRepositoryDefinitionOutput) ToGitRepositoryDefinitionOutput() GitRepositoryDefinitionOutput {
	return o
}

func (o GitRepositoryDefinitionOutput) ToGitRepositoryDefinitionOutputWithContext(ctx context.Context) GitRepositoryDefinitionOutput {
	return o
}

func (o GitRepositoryDefinitionOutput) ToGitRepositoryDefinitionPtrOutput() GitRepositoryDefinitionPtrOutput {
	return o.ToGitRepositoryDefinitionPtrOutputWithContext(context.Background())
}

func (o GitRepositoryDefinitionOutput) ToGitRepositoryDefinitionPtrOutputWithContext(ctx context.Context) GitRepositoryDefinitionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GitRepositoryDefinition) *GitRepositoryDefinition {
		return &v
	}).(GitRepositoryDefinitionPtrOutput)
}

func (o GitRepositoryDefinitionOutput) HttpsCACert() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitRepositoryDefinition) *string { return v.HttpsCACert }).(pulumi.StringPtrOutput)
}

func (o GitRepositoryDefinitionOutput) HttpsUser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitRepositoryDefinition) *string { return v.HttpsUser }).(pulumi.StringPtrOutput)
}

func (o GitRepositoryDefinitionOutput) LocalAuthRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitRepositoryDefinition) *string { return v.LocalAuthRef }).(pulumi.StringPtrOutput)
}

func (o GitRepositoryDefinitionOutput) RepositoryRef() RepositoryRefDefinitionPtrOutput {
	return o.ApplyT(func(v GitRepositoryDefinition) *RepositoryRefDefinition { return v.RepositoryRef }).(RepositoryRefDefinitionPtrOutput)
}

func (o GitRepositoryDefinitionOutput) SshKnownHosts() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitRepositoryDefinition) *string { return v.SshKnownHosts }).(pulumi.StringPtrOutput)
}

func (o GitRepositoryDefinitionOutput) SyncIntervalInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v GitRepositoryDefinition) *float64 { return v.SyncIntervalInSeconds }).(pulumi.Float64PtrOutput)
}

func (o GitRepositoryDefinitionOutput) TimeoutInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v GitRepositoryDefinition) *float64 { return v.TimeoutInSeconds }).(pulumi.Float64PtrOutput)
}

func (o GitRepositoryDefinitionOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitRepositoryDefinition) *string { return v.Url }).(pulumi.StringPtrOutput)
}

type GitRepositoryDefinitionPtrOutput struct{ *pulumi.OutputState }

func (GitRepositoryDefinitionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GitRepositoryDefinition)(nil)).Elem()
}

func (o GitRepositoryDefinitionPtrOutput) ToGitRepositoryDefinitionPtrOutput() GitRepositoryDefinitionPtrOutput {
	return o
}

func (o GitRepositoryDefinitionPtrOutput) ToGitRepositoryDefinitionPtrOutputWithContext(ctx context.Context) GitRepositoryDefinitionPtrOutput {
	return o
}

func (o GitRepositoryDefinitionPtrOutput) Elem() GitRepositoryDefinitionOutput {
	return o.ApplyT(func(v *GitRepositoryDefinition) GitRepositoryDefinition {
		if v != nil {
			return *v
		}
		var ret GitRepositoryDefinition
		return ret
	}).(GitRepositoryDefinitionOutput)
}

func (o GitRepositoryDefinitionPtrOutput) HttpsCACert() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitRepositoryDefinition) *string {
		if v == nil {
			return nil
		}
		return v.HttpsCACert
	}).(pulumi.StringPtrOutput)
}

func (o GitRepositoryDefinitionPtrOutput) HttpsUser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitRepositoryDefinition) *string {
		if v == nil {
			return nil
		}
		return v.HttpsUser
	}).(pulumi.StringPtrOutput)
}

func (o GitRepositoryDefinitionPtrOutput) LocalAuthRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitRepositoryDefinition) *string {
		if v == nil {
			return nil
		}
		return v.LocalAuthRef
	}).(pulumi.StringPtrOutput)
}

func (o GitRepositoryDefinitionPtrOutput) RepositoryRef() RepositoryRefDefinitionPtrOutput {
	return o.ApplyT(func(v *GitRepositoryDefinition) *RepositoryRefDefinition {
		if v == nil {
			return nil
		}
		return v.RepositoryRef
	}).(RepositoryRefDefinitionPtrOutput)
}

func (o GitRepositoryDefinitionPtrOutput) SshKnownHosts() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitRepositoryDefinition) *string {
		if v == nil {
			return nil
		}
		return v.SshKnownHosts
	}).(pulumi.StringPtrOutput)
}

func (o GitRepositoryDefinitionPtrOutput) SyncIntervalInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *GitRepositoryDefinition) *float64 {
		if v == nil {
			return nil
		}
		return v.SyncIntervalInSeconds
	}).(pulumi.Float64PtrOutput)
}

func (o GitRepositoryDefinitionPtrOutput) TimeoutInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *GitRepositoryDefinition) *float64 {
		if v == nil {
			return nil
		}
		return v.TimeoutInSeconds
	}).(pulumi.Float64PtrOutput)
}

func (o GitRepositoryDefinitionPtrOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitRepositoryDefinition) *string {
		if v == nil {
			return nil
		}
		return v.Url
	}).(pulumi.StringPtrOutput)
}

type GitRepositoryDefinitionResponse struct {
	HttpsCACert           *string                          `pulumi:"httpsCACert"`
	HttpsUser             *string                          `pulumi:"httpsUser"`
	LocalAuthRef          *string                          `pulumi:"localAuthRef"`
	RepositoryRef         *RepositoryRefDefinitionResponse `pulumi:"repositoryRef"`
	SshKnownHosts         *string                          `pulumi:"sshKnownHosts"`
	SyncIntervalInSeconds *float64                         `pulumi:"syncIntervalInSeconds"`
	TimeoutInSeconds      *float64                         `pulumi:"timeoutInSeconds"`
	Url                   *string                          `pulumi:"url"`
}


func (val *GitRepositoryDefinitionResponse) Defaults() *GitRepositoryDefinitionResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.SyncIntervalInSeconds) {
		syncIntervalInSeconds_ := 600.0
		tmp.SyncIntervalInSeconds = &syncIntervalInSeconds_
	}
	if isZero(tmp.TimeoutInSeconds) {
		timeoutInSeconds_ := 600.0
		tmp.TimeoutInSeconds = &timeoutInSeconds_
	}
	return &tmp
}

type GitRepositoryDefinitionResponseOutput struct{ *pulumi.OutputState }

func (GitRepositoryDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GitRepositoryDefinitionResponse)(nil)).Elem()
}

func (o GitRepositoryDefinitionResponseOutput) ToGitRepositoryDefinitionResponseOutput() GitRepositoryDefinitionResponseOutput {
	return o
}

func (o GitRepositoryDefinitionResponseOutput) ToGitRepositoryDefinitionResponseOutputWithContext(ctx context.Context) GitRepositoryDefinitionResponseOutput {
	return o
}

func (o GitRepositoryDefinitionResponseOutput) HttpsCACert() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitRepositoryDefinitionResponse) *string { return v.HttpsCACert }).(pulumi.StringPtrOutput)
}

func (o GitRepositoryDefinitionResponseOutput) HttpsUser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitRepositoryDefinitionResponse) *string { return v.HttpsUser }).(pulumi.StringPtrOutput)
}

func (o GitRepositoryDefinitionResponseOutput) LocalAuthRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitRepositoryDefinitionResponse) *string { return v.LocalAuthRef }).(pulumi.StringPtrOutput)
}

func (o GitRepositoryDefinitionResponseOutput) RepositoryRef() RepositoryRefDefinitionResponsePtrOutput {
	return o.ApplyT(func(v GitRepositoryDefinitionResponse) *RepositoryRefDefinitionResponse { return v.RepositoryRef }).(RepositoryRefDefinitionResponsePtrOutput)
}

func (o GitRepositoryDefinitionResponseOutput) SshKnownHosts() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitRepositoryDefinitionResponse) *string { return v.SshKnownHosts }).(pulumi.StringPtrOutput)
}

func (o GitRepositoryDefinitionResponseOutput) SyncIntervalInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v GitRepositoryDefinitionResponse) *float64 { return v.SyncIntervalInSeconds }).(pulumi.Float64PtrOutput)
}

func (o GitRepositoryDefinitionResponseOutput) TimeoutInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v GitRepositoryDefinitionResponse) *float64 { return v.TimeoutInSeconds }).(pulumi.Float64PtrOutput)
}

func (o GitRepositoryDefinitionResponseOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitRepositoryDefinitionResponse) *string { return v.Url }).(pulumi.StringPtrOutput)
}

type GitRepositoryDefinitionResponsePtrOutput struct{ *pulumi.OutputState }

func (GitRepositoryDefinitionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GitRepositoryDefinitionResponse)(nil)).Elem()
}

func (o GitRepositoryDefinitionResponsePtrOutput) ToGitRepositoryDefinitionResponsePtrOutput() GitRepositoryDefinitionResponsePtrOutput {
	return o
}

func (o GitRepositoryDefinitionResponsePtrOutput) ToGitRepositoryDefinitionResponsePtrOutputWithContext(ctx context.Context) GitRepositoryDefinitionResponsePtrOutput {
	return o
}

func (o GitRepositoryDefinitionResponsePtrOutput) Elem() GitRepositoryDefinitionResponseOutput {
	return o.ApplyT(func(v *GitRepositoryDefinitionResponse) GitRepositoryDefinitionResponse {
		if v != nil {
			return *v
		}
		var ret GitRepositoryDefinitionResponse
		return ret
	}).(GitRepositoryDefinitionResponseOutput)
}

func (o GitRepositoryDefinitionResponsePtrOutput) HttpsCACert() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitRepositoryDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.HttpsCACert
	}).(pulumi.StringPtrOutput)
}

func (o GitRepositoryDefinitionResponsePtrOutput) HttpsUser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitRepositoryDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.HttpsUser
	}).(pulumi.StringPtrOutput)
}

func (o GitRepositoryDefinitionResponsePtrOutput) LocalAuthRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitRepositoryDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.LocalAuthRef
	}).(pulumi.StringPtrOutput)
}

func (o GitRepositoryDefinitionResponsePtrOutput) RepositoryRef() RepositoryRefDefinitionResponsePtrOutput {
	return o.ApplyT(func(v *GitRepositoryDefinitionResponse) *RepositoryRefDefinitionResponse {
		if v == nil {
			return nil
		}
		return v.RepositoryRef
	}).(RepositoryRefDefinitionResponsePtrOutput)
}

func (o GitRepositoryDefinitionResponsePtrOutput) SshKnownHosts() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitRepositoryDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.SshKnownHosts
	}).(pulumi.StringPtrOutput)
}

func (o GitRepositoryDefinitionResponsePtrOutput) SyncIntervalInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *GitRepositoryDefinitionResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.SyncIntervalInSeconds
	}).(pulumi.Float64PtrOutput)
}

func (o GitRepositoryDefinitionResponsePtrOutput) TimeoutInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *GitRepositoryDefinitionResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.TimeoutInSeconds
	}).(pulumi.Float64PtrOutput)
}

func (o GitRepositoryDefinitionResponsePtrOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitRepositoryDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Url
	}).(pulumi.StringPtrOutput)
}

type HelmOperatorProperties struct {
	ChartValues  *string `pulumi:"chartValues"`
	ChartVersion *string `pulumi:"chartVersion"`
}





type HelmOperatorPropertiesInput interface {
	pulumi.Input

	ToHelmOperatorPropertiesOutput() HelmOperatorPropertiesOutput
	ToHelmOperatorPropertiesOutputWithContext(context.Context) HelmOperatorPropertiesOutput
}

type HelmOperatorPropertiesArgs struct {
	ChartValues  pulumi.StringPtrInput `pulumi:"chartValues"`
	ChartVersion pulumi.StringPtrInput `pulumi:"chartVersion"`
}

func (HelmOperatorPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HelmOperatorProperties)(nil)).Elem()
}

func (i HelmOperatorPropertiesArgs) ToHelmOperatorPropertiesOutput() HelmOperatorPropertiesOutput {
	return i.ToHelmOperatorPropertiesOutputWithContext(context.Background())
}

func (i HelmOperatorPropertiesArgs) ToHelmOperatorPropertiesOutputWithContext(ctx context.Context) HelmOperatorPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HelmOperatorPropertiesOutput)
}

func (i HelmOperatorPropertiesArgs) ToHelmOperatorPropertiesPtrOutput() HelmOperatorPropertiesPtrOutput {
	return i.ToHelmOperatorPropertiesPtrOutputWithContext(context.Background())
}

func (i HelmOperatorPropertiesArgs) ToHelmOperatorPropertiesPtrOutputWithContext(ctx context.Context) HelmOperatorPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HelmOperatorPropertiesOutput).ToHelmOperatorPropertiesPtrOutputWithContext(ctx)
}









type HelmOperatorPropertiesPtrInput interface {
	pulumi.Input

	ToHelmOperatorPropertiesPtrOutput() HelmOperatorPropertiesPtrOutput
	ToHelmOperatorPropertiesPtrOutputWithContext(context.Context) HelmOperatorPropertiesPtrOutput
}

type helmOperatorPropertiesPtrType HelmOperatorPropertiesArgs

func HelmOperatorPropertiesPtr(v *HelmOperatorPropertiesArgs) HelmOperatorPropertiesPtrInput {
	return (*helmOperatorPropertiesPtrType)(v)
}

func (*helmOperatorPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HelmOperatorProperties)(nil)).Elem()
}

func (i *helmOperatorPropertiesPtrType) ToHelmOperatorPropertiesPtrOutput() HelmOperatorPropertiesPtrOutput {
	return i.ToHelmOperatorPropertiesPtrOutputWithContext(context.Background())
}

func (i *helmOperatorPropertiesPtrType) ToHelmOperatorPropertiesPtrOutputWithContext(ctx context.Context) HelmOperatorPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HelmOperatorPropertiesPtrOutput)
}

type HelmOperatorPropertiesOutput struct{ *pulumi.OutputState }

func (HelmOperatorPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HelmOperatorProperties)(nil)).Elem()
}

func (o HelmOperatorPropertiesOutput) ToHelmOperatorPropertiesOutput() HelmOperatorPropertiesOutput {
	return o
}

func (o HelmOperatorPropertiesOutput) ToHelmOperatorPropertiesOutputWithContext(ctx context.Context) HelmOperatorPropertiesOutput {
	return o
}

func (o HelmOperatorPropertiesOutput) ToHelmOperatorPropertiesPtrOutput() HelmOperatorPropertiesPtrOutput {
	return o.ToHelmOperatorPropertiesPtrOutputWithContext(context.Background())
}

func (o HelmOperatorPropertiesOutput) ToHelmOperatorPropertiesPtrOutputWithContext(ctx context.Context) HelmOperatorPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HelmOperatorProperties) *HelmOperatorProperties {
		return &v
	}).(HelmOperatorPropertiesPtrOutput)
}

func (o HelmOperatorPropertiesOutput) ChartValues() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HelmOperatorProperties) *string { return v.ChartValues }).(pulumi.StringPtrOutput)
}

func (o HelmOperatorPropertiesOutput) ChartVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HelmOperatorProperties) *string { return v.ChartVersion }).(pulumi.StringPtrOutput)
}

type HelmOperatorPropertiesPtrOutput struct{ *pulumi.OutputState }

func (HelmOperatorPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HelmOperatorProperties)(nil)).Elem()
}

func (o HelmOperatorPropertiesPtrOutput) ToHelmOperatorPropertiesPtrOutput() HelmOperatorPropertiesPtrOutput {
	return o
}

func (o HelmOperatorPropertiesPtrOutput) ToHelmOperatorPropertiesPtrOutputWithContext(ctx context.Context) HelmOperatorPropertiesPtrOutput {
	return o
}

func (o HelmOperatorPropertiesPtrOutput) Elem() HelmOperatorPropertiesOutput {
	return o.ApplyT(func(v *HelmOperatorProperties) HelmOperatorProperties {
		if v != nil {
			return *v
		}
		var ret HelmOperatorProperties
		return ret
	}).(HelmOperatorPropertiesOutput)
}

func (o HelmOperatorPropertiesPtrOutput) ChartValues() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HelmOperatorProperties) *string {
		if v == nil {
			return nil
		}
		return v.ChartValues
	}).(pulumi.StringPtrOutput)
}

func (o HelmOperatorPropertiesPtrOutput) ChartVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HelmOperatorProperties) *string {
		if v == nil {
			return nil
		}
		return v.ChartVersion
	}).(pulumi.StringPtrOutput)
}

type HelmOperatorPropertiesResponse struct {
	ChartValues  *string `pulumi:"chartValues"`
	ChartVersion *string `pulumi:"chartVersion"`
}

type HelmOperatorPropertiesResponseOutput struct{ *pulumi.OutputState }

func (HelmOperatorPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HelmOperatorPropertiesResponse)(nil)).Elem()
}

func (o HelmOperatorPropertiesResponseOutput) ToHelmOperatorPropertiesResponseOutput() HelmOperatorPropertiesResponseOutput {
	return o
}

func (o HelmOperatorPropertiesResponseOutput) ToHelmOperatorPropertiesResponseOutputWithContext(ctx context.Context) HelmOperatorPropertiesResponseOutput {
	return o
}

func (o HelmOperatorPropertiesResponseOutput) ChartValues() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HelmOperatorPropertiesResponse) *string { return v.ChartValues }).(pulumi.StringPtrOutput)
}

func (o HelmOperatorPropertiesResponseOutput) ChartVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HelmOperatorPropertiesResponse) *string { return v.ChartVersion }).(pulumi.StringPtrOutput)
}

type HelmOperatorPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (HelmOperatorPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HelmOperatorPropertiesResponse)(nil)).Elem()
}

func (o HelmOperatorPropertiesResponsePtrOutput) ToHelmOperatorPropertiesResponsePtrOutput() HelmOperatorPropertiesResponsePtrOutput {
	return o
}

func (o HelmOperatorPropertiesResponsePtrOutput) ToHelmOperatorPropertiesResponsePtrOutputWithContext(ctx context.Context) HelmOperatorPropertiesResponsePtrOutput {
	return o
}

func (o HelmOperatorPropertiesResponsePtrOutput) Elem() HelmOperatorPropertiesResponseOutput {
	return o.ApplyT(func(v *HelmOperatorPropertiesResponse) HelmOperatorPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret HelmOperatorPropertiesResponse
		return ret
	}).(HelmOperatorPropertiesResponseOutput)
}

func (o HelmOperatorPropertiesResponsePtrOutput) ChartValues() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HelmOperatorPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ChartValues
	}).(pulumi.StringPtrOutput)
}

func (o HelmOperatorPropertiesResponsePtrOutput) ChartVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HelmOperatorPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ChartVersion
	}).(pulumi.StringPtrOutput)
}

type HelmReleasePropertiesDefinitionResponse struct {
	FailureCount        *float64                           `pulumi:"failureCount"`
	HelmChartRef        *ObjectReferenceDefinitionResponse `pulumi:"helmChartRef"`
	InstallFailureCount *float64                           `pulumi:"installFailureCount"`
	LastRevisionApplied *float64                           `pulumi:"lastRevisionApplied"`
	UpgradeFailureCount *float64                           `pulumi:"upgradeFailureCount"`
}

type HelmReleasePropertiesDefinitionResponseOutput struct{ *pulumi.OutputState }

func (HelmReleasePropertiesDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HelmReleasePropertiesDefinitionResponse)(nil)).Elem()
}

func (o HelmReleasePropertiesDefinitionResponseOutput) ToHelmReleasePropertiesDefinitionResponseOutput() HelmReleasePropertiesDefinitionResponseOutput {
	return o
}

func (o HelmReleasePropertiesDefinitionResponseOutput) ToHelmReleasePropertiesDefinitionResponseOutputWithContext(ctx context.Context) HelmReleasePropertiesDefinitionResponseOutput {
	return o
}

func (o HelmReleasePropertiesDefinitionResponseOutput) FailureCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v HelmReleasePropertiesDefinitionResponse) *float64 { return v.FailureCount }).(pulumi.Float64PtrOutput)
}

func (o HelmReleasePropertiesDefinitionResponseOutput) HelmChartRef() ObjectReferenceDefinitionResponsePtrOutput {
	return o.ApplyT(func(v HelmReleasePropertiesDefinitionResponse) *ObjectReferenceDefinitionResponse {
		return v.HelmChartRef
	}).(ObjectReferenceDefinitionResponsePtrOutput)
}

func (o HelmReleasePropertiesDefinitionResponseOutput) InstallFailureCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v HelmReleasePropertiesDefinitionResponse) *float64 { return v.InstallFailureCount }).(pulumi.Float64PtrOutput)
}

func (o HelmReleasePropertiesDefinitionResponseOutput) LastRevisionApplied() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v HelmReleasePropertiesDefinitionResponse) *float64 { return v.LastRevisionApplied }).(pulumi.Float64PtrOutput)
}

func (o HelmReleasePropertiesDefinitionResponseOutput) UpgradeFailureCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v HelmReleasePropertiesDefinitionResponse) *float64 { return v.UpgradeFailureCount }).(pulumi.Float64PtrOutput)
}

type HelmReleasePropertiesDefinitionResponsePtrOutput struct{ *pulumi.OutputState }

func (HelmReleasePropertiesDefinitionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HelmReleasePropertiesDefinitionResponse)(nil)).Elem()
}

func (o HelmReleasePropertiesDefinitionResponsePtrOutput) ToHelmReleasePropertiesDefinitionResponsePtrOutput() HelmReleasePropertiesDefinitionResponsePtrOutput {
	return o
}

func (o HelmReleasePropertiesDefinitionResponsePtrOutput) ToHelmReleasePropertiesDefinitionResponsePtrOutputWithContext(ctx context.Context) HelmReleasePropertiesDefinitionResponsePtrOutput {
	return o
}

func (o HelmReleasePropertiesDefinitionResponsePtrOutput) Elem() HelmReleasePropertiesDefinitionResponseOutput {
	return o.ApplyT(func(v *HelmReleasePropertiesDefinitionResponse) HelmReleasePropertiesDefinitionResponse {
		if v != nil {
			return *v
		}
		var ret HelmReleasePropertiesDefinitionResponse
		return ret
	}).(HelmReleasePropertiesDefinitionResponseOutput)
}

func (o HelmReleasePropertiesDefinitionResponsePtrOutput) FailureCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *HelmReleasePropertiesDefinitionResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.FailureCount
	}).(pulumi.Float64PtrOutput)
}

func (o HelmReleasePropertiesDefinitionResponsePtrOutput) HelmChartRef() ObjectReferenceDefinitionResponsePtrOutput {
	return o.ApplyT(func(v *HelmReleasePropertiesDefinitionResponse) *ObjectReferenceDefinitionResponse {
		if v == nil {
			return nil
		}
		return v.HelmChartRef
	}).(ObjectReferenceDefinitionResponsePtrOutput)
}

func (o HelmReleasePropertiesDefinitionResponsePtrOutput) InstallFailureCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *HelmReleasePropertiesDefinitionResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.InstallFailureCount
	}).(pulumi.Float64PtrOutput)
}

func (o HelmReleasePropertiesDefinitionResponsePtrOutput) LastRevisionApplied() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *HelmReleasePropertiesDefinitionResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.LastRevisionApplied
	}).(pulumi.Float64PtrOutput)
}

func (o HelmReleasePropertiesDefinitionResponsePtrOutput) UpgradeFailureCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *HelmReleasePropertiesDefinitionResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.UpgradeFailureCount
	}).(pulumi.Float64PtrOutput)
}

type Identity struct {
	Type *ResourceIdentityType `pulumi:"type"`
}





type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(context.Context) IdentityOutput
}

type IdentityArgs struct {
	Type ResourceIdentityTypePtrInput `pulumi:"type"`
}

func (IdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (i IdentityArgs) ToIdentityOutput() IdentityOutput {
	return i.ToIdentityOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput)
}

func (i IdentityArgs) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput).ToIdentityPtrOutputWithContext(ctx)
}









type IdentityPtrInput interface {
	pulumi.Input

	ToIdentityPtrOutput() IdentityPtrOutput
	ToIdentityPtrOutputWithContext(context.Context) IdentityPtrOutput
}

type identityPtrType IdentityArgs

func IdentityPtr(v *IdentityArgs) IdentityPtrInput {
	return (*identityPtrType)(v)
}

func (*identityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (i *identityPtrType) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i *identityPtrType) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPtrOutput)
}

type IdentityOutput struct{ *pulumi.OutputState }

func (IdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (o IdentityOutput) ToIdentityOutput() IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o.ToIdentityPtrOutputWithContext(context.Background())
}

func (o IdentityOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Identity) *Identity {
		return &v
	}).(IdentityPtrOutput)
}

func (o IdentityOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v Identity) *ResourceIdentityType { return v.Type }).(ResourceIdentityTypePtrOutput)
}

type IdentityPtrOutput struct{ *pulumi.OutputState }

func (IdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (o IdentityPtrOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) Elem() IdentityOutput {
	return o.ApplyT(func(v *Identity) Identity {
		if v != nil {
			return *v
		}
		var ret Identity
		return ret
	}).(IdentityOutput)
}

func (o IdentityPtrOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v *Identity) *ResourceIdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(ResourceIdentityTypePtrOutput)
}

type IdentityResponse struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}

type IdentityResponseOutput struct{ *pulumi.OutputState }

func (IdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponse)(nil)).Elem()
}

func (o IdentityResponseOutput) ToIdentityResponseOutput() IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) ToIdentityResponseOutputWithContext(ctx context.Context) IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type IdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityResponse)(nil)).Elem()
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) Elem() IdentityResponseOutput {
	return o.ApplyT(func(v *IdentityResponse) IdentityResponse {
		if v != nil {
			return *v
		}
		var ret IdentityResponse
		return ret
	}).(IdentityResponseOutput)
}

func (o IdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type KustomizationDefinition struct {
	DependsOn              []string `pulumi:"dependsOn"`
	Force                  *bool    `pulumi:"force"`
	Path                   *string  `pulumi:"path"`
	Prune                  *bool    `pulumi:"prune"`
	RetryIntervalInSeconds *float64 `pulumi:"retryIntervalInSeconds"`
	SyncIntervalInSeconds  *float64 `pulumi:"syncIntervalInSeconds"`
	TimeoutInSeconds       *float64 `pulumi:"timeoutInSeconds"`
}


func (val *KustomizationDefinition) Defaults() *KustomizationDefinition {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Force) {
		force_ := false
		tmp.Force = &force_
	}
	if isZero(tmp.Path) {
		path_ := ""
		tmp.Path = &path_
	}
	if isZero(tmp.Prune) {
		prune_ := false
		tmp.Prune = &prune_
	}
	if isZero(tmp.SyncIntervalInSeconds) {
		syncIntervalInSeconds_ := 600.0
		tmp.SyncIntervalInSeconds = &syncIntervalInSeconds_
	}
	if isZero(tmp.TimeoutInSeconds) {
		timeoutInSeconds_ := 600.0
		tmp.TimeoutInSeconds = &timeoutInSeconds_
	}
	return &tmp
}





type KustomizationDefinitionInput interface {
	pulumi.Input

	ToKustomizationDefinitionOutput() KustomizationDefinitionOutput
	ToKustomizationDefinitionOutputWithContext(context.Context) KustomizationDefinitionOutput
}

type KustomizationDefinitionArgs struct {
	DependsOn              pulumi.StringArrayInput `pulumi:"dependsOn"`
	Force                  pulumi.BoolPtrInput     `pulumi:"force"`
	Path                   pulumi.StringPtrInput   `pulumi:"path"`
	Prune                  pulumi.BoolPtrInput     `pulumi:"prune"`
	RetryIntervalInSeconds pulumi.Float64PtrInput  `pulumi:"retryIntervalInSeconds"`
	SyncIntervalInSeconds  pulumi.Float64PtrInput  `pulumi:"syncIntervalInSeconds"`
	TimeoutInSeconds       pulumi.Float64PtrInput  `pulumi:"timeoutInSeconds"`
}


func (val *KustomizationDefinitionArgs) Defaults() *KustomizationDefinitionArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Force) {
		tmp.Force = pulumi.BoolPtr(false)
	}
	if isZero(tmp.Path) {
		tmp.Path = pulumi.StringPtr("")
	}
	if isZero(tmp.Prune) {
		tmp.Prune = pulumi.BoolPtr(false)
	}
	if isZero(tmp.SyncIntervalInSeconds) {
		tmp.SyncIntervalInSeconds = pulumi.Float64Ptr(600.0)
	}
	if isZero(tmp.TimeoutInSeconds) {
		tmp.TimeoutInSeconds = pulumi.Float64Ptr(600.0)
	}
	return &tmp
}
func (KustomizationDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KustomizationDefinition)(nil)).Elem()
}

func (i KustomizationDefinitionArgs) ToKustomizationDefinitionOutput() KustomizationDefinitionOutput {
	return i.ToKustomizationDefinitionOutputWithContext(context.Background())
}

func (i KustomizationDefinitionArgs) ToKustomizationDefinitionOutputWithContext(ctx context.Context) KustomizationDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KustomizationDefinitionOutput)
}





type KustomizationDefinitionMapInput interface {
	pulumi.Input

	ToKustomizationDefinitionMapOutput() KustomizationDefinitionMapOutput
	ToKustomizationDefinitionMapOutputWithContext(context.Context) KustomizationDefinitionMapOutput
}

type KustomizationDefinitionMap map[string]KustomizationDefinitionInput

func (KustomizationDefinitionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]KustomizationDefinition)(nil)).Elem()
}

func (i KustomizationDefinitionMap) ToKustomizationDefinitionMapOutput() KustomizationDefinitionMapOutput {
	return i.ToKustomizationDefinitionMapOutputWithContext(context.Background())
}

func (i KustomizationDefinitionMap) ToKustomizationDefinitionMapOutputWithContext(ctx context.Context) KustomizationDefinitionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KustomizationDefinitionMapOutput)
}

type KustomizationDefinitionOutput struct{ *pulumi.OutputState }

func (KustomizationDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KustomizationDefinition)(nil)).Elem()
}

func (o KustomizationDefinitionOutput) ToKustomizationDefinitionOutput() KustomizationDefinitionOutput {
	return o
}

func (o KustomizationDefinitionOutput) ToKustomizationDefinitionOutputWithContext(ctx context.Context) KustomizationDefinitionOutput {
	return o
}

func (o KustomizationDefinitionOutput) DependsOn() pulumi.StringArrayOutput {
	return o.ApplyT(func(v KustomizationDefinition) []string { return v.DependsOn }).(pulumi.StringArrayOutput)
}

func (o KustomizationDefinitionOutput) Force() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KustomizationDefinition) *bool { return v.Force }).(pulumi.BoolPtrOutput)
}

func (o KustomizationDefinitionOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KustomizationDefinition) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o KustomizationDefinitionOutput) Prune() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KustomizationDefinition) *bool { return v.Prune }).(pulumi.BoolPtrOutput)
}

func (o KustomizationDefinitionOutput) RetryIntervalInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v KustomizationDefinition) *float64 { return v.RetryIntervalInSeconds }).(pulumi.Float64PtrOutput)
}

func (o KustomizationDefinitionOutput) SyncIntervalInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v KustomizationDefinition) *float64 { return v.SyncIntervalInSeconds }).(pulumi.Float64PtrOutput)
}

func (o KustomizationDefinitionOutput) TimeoutInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v KustomizationDefinition) *float64 { return v.TimeoutInSeconds }).(pulumi.Float64PtrOutput)
}

type KustomizationDefinitionMapOutput struct{ *pulumi.OutputState }

func (KustomizationDefinitionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]KustomizationDefinition)(nil)).Elem()
}

func (o KustomizationDefinitionMapOutput) ToKustomizationDefinitionMapOutput() KustomizationDefinitionMapOutput {
	return o
}

func (o KustomizationDefinitionMapOutput) ToKustomizationDefinitionMapOutputWithContext(ctx context.Context) KustomizationDefinitionMapOutput {
	return o
}

func (o KustomizationDefinitionMapOutput) MapIndex(k pulumi.StringInput) KustomizationDefinitionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) KustomizationDefinition {
		return vs[0].(map[string]KustomizationDefinition)[vs[1].(string)]
	}).(KustomizationDefinitionOutput)
}

type KustomizationDefinitionResponse struct {
	DependsOn              []string `pulumi:"dependsOn"`
	Force                  *bool    `pulumi:"force"`
	Name                   string   `pulumi:"name"`
	Path                   *string  `pulumi:"path"`
	Prune                  *bool    `pulumi:"prune"`
	RetryIntervalInSeconds *float64 `pulumi:"retryIntervalInSeconds"`
	SyncIntervalInSeconds  *float64 `pulumi:"syncIntervalInSeconds"`
	TimeoutInSeconds       *float64 `pulumi:"timeoutInSeconds"`
}


func (val *KustomizationDefinitionResponse) Defaults() *KustomizationDefinitionResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Force) {
		force_ := false
		tmp.Force = &force_
	}
	if isZero(tmp.Path) {
		path_ := ""
		tmp.Path = &path_
	}
	if isZero(tmp.Prune) {
		prune_ := false
		tmp.Prune = &prune_
	}
	if isZero(tmp.SyncIntervalInSeconds) {
		syncIntervalInSeconds_ := 600.0
		tmp.SyncIntervalInSeconds = &syncIntervalInSeconds_
	}
	if isZero(tmp.TimeoutInSeconds) {
		timeoutInSeconds_ := 600.0
		tmp.TimeoutInSeconds = &timeoutInSeconds_
	}
	return &tmp
}

type KustomizationDefinitionResponseOutput struct{ *pulumi.OutputState }

func (KustomizationDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KustomizationDefinitionResponse)(nil)).Elem()
}

func (o KustomizationDefinitionResponseOutput) ToKustomizationDefinitionResponseOutput() KustomizationDefinitionResponseOutput {
	return o
}

func (o KustomizationDefinitionResponseOutput) ToKustomizationDefinitionResponseOutputWithContext(ctx context.Context) KustomizationDefinitionResponseOutput {
	return o
}

func (o KustomizationDefinitionResponseOutput) DependsOn() pulumi.StringArrayOutput {
	return o.ApplyT(func(v KustomizationDefinitionResponse) []string { return v.DependsOn }).(pulumi.StringArrayOutput)
}

func (o KustomizationDefinitionResponseOutput) Force() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KustomizationDefinitionResponse) *bool { return v.Force }).(pulumi.BoolPtrOutput)
}

func (o KustomizationDefinitionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v KustomizationDefinitionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o KustomizationDefinitionResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KustomizationDefinitionResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o KustomizationDefinitionResponseOutput) Prune() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KustomizationDefinitionResponse) *bool { return v.Prune }).(pulumi.BoolPtrOutput)
}

func (o KustomizationDefinitionResponseOutput) RetryIntervalInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v KustomizationDefinitionResponse) *float64 { return v.RetryIntervalInSeconds }).(pulumi.Float64PtrOutput)
}

func (o KustomizationDefinitionResponseOutput) SyncIntervalInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v KustomizationDefinitionResponse) *float64 { return v.SyncIntervalInSeconds }).(pulumi.Float64PtrOutput)
}

func (o KustomizationDefinitionResponseOutput) TimeoutInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v KustomizationDefinitionResponse) *float64 { return v.TimeoutInSeconds }).(pulumi.Float64PtrOutput)
}

type KustomizationDefinitionResponseMapOutput struct{ *pulumi.OutputState }

func (KustomizationDefinitionResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]KustomizationDefinitionResponse)(nil)).Elem()
}

func (o KustomizationDefinitionResponseMapOutput) ToKustomizationDefinitionResponseMapOutput() KustomizationDefinitionResponseMapOutput {
	return o
}

func (o KustomizationDefinitionResponseMapOutput) ToKustomizationDefinitionResponseMapOutputWithContext(ctx context.Context) KustomizationDefinitionResponseMapOutput {
	return o
}

func (o KustomizationDefinitionResponseMapOutput) MapIndex(k pulumi.StringInput) KustomizationDefinitionResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) KustomizationDefinitionResponse {
		return vs[0].(map[string]KustomizationDefinitionResponse)[vs[1].(string)]
	}).(KustomizationDefinitionResponseOutput)
}

type ManagedIdentityDefinition struct {
	ClientId *string `pulumi:"clientId"`
}





type ManagedIdentityDefinitionInput interface {
	pulumi.Input

	ToManagedIdentityDefinitionOutput() ManagedIdentityDefinitionOutput
	ToManagedIdentityDefinitionOutputWithContext(context.Context) ManagedIdentityDefinitionOutput
}

type ManagedIdentityDefinitionArgs struct {
	ClientId pulumi.StringPtrInput `pulumi:"clientId"`
}

func (ManagedIdentityDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentityDefinition)(nil)).Elem()
}

func (i ManagedIdentityDefinitionArgs) ToManagedIdentityDefinitionOutput() ManagedIdentityDefinitionOutput {
	return i.ToManagedIdentityDefinitionOutputWithContext(context.Background())
}

func (i ManagedIdentityDefinitionArgs) ToManagedIdentityDefinitionOutputWithContext(ctx context.Context) ManagedIdentityDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentityDefinitionOutput)
}

func (i ManagedIdentityDefinitionArgs) ToManagedIdentityDefinitionPtrOutput() ManagedIdentityDefinitionPtrOutput {
	return i.ToManagedIdentityDefinitionPtrOutputWithContext(context.Background())
}

func (i ManagedIdentityDefinitionArgs) ToManagedIdentityDefinitionPtrOutputWithContext(ctx context.Context) ManagedIdentityDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentityDefinitionOutput).ToManagedIdentityDefinitionPtrOutputWithContext(ctx)
}









type ManagedIdentityDefinitionPtrInput interface {
	pulumi.Input

	ToManagedIdentityDefinitionPtrOutput() ManagedIdentityDefinitionPtrOutput
	ToManagedIdentityDefinitionPtrOutputWithContext(context.Context) ManagedIdentityDefinitionPtrOutput
}

type managedIdentityDefinitionPtrType ManagedIdentityDefinitionArgs

func ManagedIdentityDefinitionPtr(v *ManagedIdentityDefinitionArgs) ManagedIdentityDefinitionPtrInput {
	return (*managedIdentityDefinitionPtrType)(v)
}

func (*managedIdentityDefinitionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentityDefinition)(nil)).Elem()
}

func (i *managedIdentityDefinitionPtrType) ToManagedIdentityDefinitionPtrOutput() ManagedIdentityDefinitionPtrOutput {
	return i.ToManagedIdentityDefinitionPtrOutputWithContext(context.Background())
}

func (i *managedIdentityDefinitionPtrType) ToManagedIdentityDefinitionPtrOutputWithContext(ctx context.Context) ManagedIdentityDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentityDefinitionPtrOutput)
}

type ManagedIdentityDefinitionOutput struct{ *pulumi.OutputState }

func (ManagedIdentityDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentityDefinition)(nil)).Elem()
}

func (o ManagedIdentityDefinitionOutput) ToManagedIdentityDefinitionOutput() ManagedIdentityDefinitionOutput {
	return o
}

func (o ManagedIdentityDefinitionOutput) ToManagedIdentityDefinitionOutputWithContext(ctx context.Context) ManagedIdentityDefinitionOutput {
	return o
}

func (o ManagedIdentityDefinitionOutput) ToManagedIdentityDefinitionPtrOutput() ManagedIdentityDefinitionPtrOutput {
	return o.ToManagedIdentityDefinitionPtrOutputWithContext(context.Background())
}

func (o ManagedIdentityDefinitionOutput) ToManagedIdentityDefinitionPtrOutputWithContext(ctx context.Context) ManagedIdentityDefinitionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedIdentityDefinition) *ManagedIdentityDefinition {
		return &v
	}).(ManagedIdentityDefinitionPtrOutput)
}

func (o ManagedIdentityDefinitionOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedIdentityDefinition) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

type ManagedIdentityDefinitionPtrOutput struct{ *pulumi.OutputState }

func (ManagedIdentityDefinitionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentityDefinition)(nil)).Elem()
}

func (o ManagedIdentityDefinitionPtrOutput) ToManagedIdentityDefinitionPtrOutput() ManagedIdentityDefinitionPtrOutput {
	return o
}

func (o ManagedIdentityDefinitionPtrOutput) ToManagedIdentityDefinitionPtrOutputWithContext(ctx context.Context) ManagedIdentityDefinitionPtrOutput {
	return o
}

func (o ManagedIdentityDefinitionPtrOutput) Elem() ManagedIdentityDefinitionOutput {
	return o.ApplyT(func(v *ManagedIdentityDefinition) ManagedIdentityDefinition {
		if v != nil {
			return *v
		}
		var ret ManagedIdentityDefinition
		return ret
	}).(ManagedIdentityDefinitionOutput)
}

func (o ManagedIdentityDefinitionPtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedIdentityDefinition) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

type ManagedIdentityDefinitionResponse struct {
	ClientId *string `pulumi:"clientId"`
}

type ManagedIdentityDefinitionResponseOutput struct{ *pulumi.OutputState }

func (ManagedIdentityDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentityDefinitionResponse)(nil)).Elem()
}

func (o ManagedIdentityDefinitionResponseOutput) ToManagedIdentityDefinitionResponseOutput() ManagedIdentityDefinitionResponseOutput {
	return o
}

func (o ManagedIdentityDefinitionResponseOutput) ToManagedIdentityDefinitionResponseOutputWithContext(ctx context.Context) ManagedIdentityDefinitionResponseOutput {
	return o
}

func (o ManagedIdentityDefinitionResponseOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedIdentityDefinitionResponse) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

type ManagedIdentityDefinitionResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedIdentityDefinitionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentityDefinitionResponse)(nil)).Elem()
}

func (o ManagedIdentityDefinitionResponsePtrOutput) ToManagedIdentityDefinitionResponsePtrOutput() ManagedIdentityDefinitionResponsePtrOutput {
	return o
}

func (o ManagedIdentityDefinitionResponsePtrOutput) ToManagedIdentityDefinitionResponsePtrOutputWithContext(ctx context.Context) ManagedIdentityDefinitionResponsePtrOutput {
	return o
}

func (o ManagedIdentityDefinitionResponsePtrOutput) Elem() ManagedIdentityDefinitionResponseOutput {
	return o.ApplyT(func(v *ManagedIdentityDefinitionResponse) ManagedIdentityDefinitionResponse {
		if v != nil {
			return *v
		}
		var ret ManagedIdentityDefinitionResponse
		return ret
	}).(ManagedIdentityDefinitionResponseOutput)
}

func (o ManagedIdentityDefinitionResponsePtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedIdentityDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

type ObjectReferenceDefinitionResponse struct {
	Name      *string `pulumi:"name"`
	Namespace *string `pulumi:"namespace"`
}

type ObjectReferenceDefinitionResponseOutput struct{ *pulumi.OutputState }

func (ObjectReferenceDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ObjectReferenceDefinitionResponse)(nil)).Elem()
}

func (o ObjectReferenceDefinitionResponseOutput) ToObjectReferenceDefinitionResponseOutput() ObjectReferenceDefinitionResponseOutput {
	return o
}

func (o ObjectReferenceDefinitionResponseOutput) ToObjectReferenceDefinitionResponseOutputWithContext(ctx context.Context) ObjectReferenceDefinitionResponseOutput {
	return o
}

func (o ObjectReferenceDefinitionResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ObjectReferenceDefinitionResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ObjectReferenceDefinitionResponseOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ObjectReferenceDefinitionResponse) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

type ObjectReferenceDefinitionResponsePtrOutput struct{ *pulumi.OutputState }

func (ObjectReferenceDefinitionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectReferenceDefinitionResponse)(nil)).Elem()
}

func (o ObjectReferenceDefinitionResponsePtrOutput) ToObjectReferenceDefinitionResponsePtrOutput() ObjectReferenceDefinitionResponsePtrOutput {
	return o
}

func (o ObjectReferenceDefinitionResponsePtrOutput) ToObjectReferenceDefinitionResponsePtrOutputWithContext(ctx context.Context) ObjectReferenceDefinitionResponsePtrOutput {
	return o
}

func (o ObjectReferenceDefinitionResponsePtrOutput) Elem() ObjectReferenceDefinitionResponseOutput {
	return o.ApplyT(func(v *ObjectReferenceDefinitionResponse) ObjectReferenceDefinitionResponse {
		if v != nil {
			return *v
		}
		var ret ObjectReferenceDefinitionResponse
		return ret
	}).(ObjectReferenceDefinitionResponseOutput)
}

func (o ObjectReferenceDefinitionResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ObjectReferenceDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ObjectReferenceDefinitionResponsePtrOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ObjectReferenceDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Namespace
	}).(pulumi.StringPtrOutput)
}

type ObjectStatusConditionDefinitionResponse struct {
	LastTransitionTime *string `pulumi:"lastTransitionTime"`
	Message            *string `pulumi:"message"`
	Reason             *string `pulumi:"reason"`
	Status             *string `pulumi:"status"`
	Type               *string `pulumi:"type"`
}

type ObjectStatusConditionDefinitionResponseOutput struct{ *pulumi.OutputState }

func (ObjectStatusConditionDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ObjectStatusConditionDefinitionResponse)(nil)).Elem()
}

func (o ObjectStatusConditionDefinitionResponseOutput) ToObjectStatusConditionDefinitionResponseOutput() ObjectStatusConditionDefinitionResponseOutput {
	return o
}

func (o ObjectStatusConditionDefinitionResponseOutput) ToObjectStatusConditionDefinitionResponseOutputWithContext(ctx context.Context) ObjectStatusConditionDefinitionResponseOutput {
	return o
}

func (o ObjectStatusConditionDefinitionResponseOutput) LastTransitionTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ObjectStatusConditionDefinitionResponse) *string { return v.LastTransitionTime }).(pulumi.StringPtrOutput)
}

func (o ObjectStatusConditionDefinitionResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ObjectStatusConditionDefinitionResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o ObjectStatusConditionDefinitionResponseOutput) Reason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ObjectStatusConditionDefinitionResponse) *string { return v.Reason }).(pulumi.StringPtrOutput)
}

func (o ObjectStatusConditionDefinitionResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ObjectStatusConditionDefinitionResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o ObjectStatusConditionDefinitionResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ObjectStatusConditionDefinitionResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ObjectStatusConditionDefinitionResponseArrayOutput struct{ *pulumi.OutputState }

func (ObjectStatusConditionDefinitionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ObjectStatusConditionDefinitionResponse)(nil)).Elem()
}

func (o ObjectStatusConditionDefinitionResponseArrayOutput) ToObjectStatusConditionDefinitionResponseArrayOutput() ObjectStatusConditionDefinitionResponseArrayOutput {
	return o
}

func (o ObjectStatusConditionDefinitionResponseArrayOutput) ToObjectStatusConditionDefinitionResponseArrayOutputWithContext(ctx context.Context) ObjectStatusConditionDefinitionResponseArrayOutput {
	return o
}

func (o ObjectStatusConditionDefinitionResponseArrayOutput) Index(i pulumi.IntInput) ObjectStatusConditionDefinitionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ObjectStatusConditionDefinitionResponse {
		return vs[0].([]ObjectStatusConditionDefinitionResponse)[vs[1].(int)]
	}).(ObjectStatusConditionDefinitionResponseOutput)
}

type ObjectStatusDefinitionResponse struct {
	AppliedBy             *ObjectReferenceDefinitionResponse        `pulumi:"appliedBy"`
	ComplianceState       *string                                   `pulumi:"complianceState"`
	HelmReleaseProperties *HelmReleasePropertiesDefinitionResponse  `pulumi:"helmReleaseProperties"`
	Kind                  *string                                   `pulumi:"kind"`
	Name                  *string                                   `pulumi:"name"`
	Namespace             *string                                   `pulumi:"namespace"`
	StatusConditions      []ObjectStatusConditionDefinitionResponse `pulumi:"statusConditions"`
}

type ObjectStatusDefinitionResponseOutput struct{ *pulumi.OutputState }

func (ObjectStatusDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ObjectStatusDefinitionResponse)(nil)).Elem()
}

func (o ObjectStatusDefinitionResponseOutput) ToObjectStatusDefinitionResponseOutput() ObjectStatusDefinitionResponseOutput {
	return o
}

func (o ObjectStatusDefinitionResponseOutput) ToObjectStatusDefinitionResponseOutputWithContext(ctx context.Context) ObjectStatusDefinitionResponseOutput {
	return o
}

func (o ObjectStatusDefinitionResponseOutput) AppliedBy() ObjectReferenceDefinitionResponsePtrOutput {
	return o.ApplyT(func(v ObjectStatusDefinitionResponse) *ObjectReferenceDefinitionResponse { return v.AppliedBy }).(ObjectReferenceDefinitionResponsePtrOutput)
}

func (o ObjectStatusDefinitionResponseOutput) ComplianceState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ObjectStatusDefinitionResponse) *string { return v.ComplianceState }).(pulumi.StringPtrOutput)
}

func (o ObjectStatusDefinitionResponseOutput) HelmReleaseProperties() HelmReleasePropertiesDefinitionResponsePtrOutput {
	return o.ApplyT(func(v ObjectStatusDefinitionResponse) *HelmReleasePropertiesDefinitionResponse {
		return v.HelmReleaseProperties
	}).(HelmReleasePropertiesDefinitionResponsePtrOutput)
}

func (o ObjectStatusDefinitionResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ObjectStatusDefinitionResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ObjectStatusDefinitionResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ObjectStatusDefinitionResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ObjectStatusDefinitionResponseOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ObjectStatusDefinitionResponse) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

func (o ObjectStatusDefinitionResponseOutput) StatusConditions() ObjectStatusConditionDefinitionResponseArrayOutput {
	return o.ApplyT(func(v ObjectStatusDefinitionResponse) []ObjectStatusConditionDefinitionResponse {
		return v.StatusConditions
	}).(ObjectStatusConditionDefinitionResponseArrayOutput)
}

type ObjectStatusDefinitionResponseArrayOutput struct{ *pulumi.OutputState }

func (ObjectStatusDefinitionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ObjectStatusDefinitionResponse)(nil)).Elem()
}

func (o ObjectStatusDefinitionResponseArrayOutput) ToObjectStatusDefinitionResponseArrayOutput() ObjectStatusDefinitionResponseArrayOutput {
	return o
}

func (o ObjectStatusDefinitionResponseArrayOutput) ToObjectStatusDefinitionResponseArrayOutputWithContext(ctx context.Context) ObjectStatusDefinitionResponseArrayOutput {
	return o
}

func (o ObjectStatusDefinitionResponseArrayOutput) Index(i pulumi.IntInput) ObjectStatusDefinitionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ObjectStatusDefinitionResponse {
		return vs[0].([]ObjectStatusDefinitionResponse)[vs[1].(int)]
	}).(ObjectStatusDefinitionResponseOutput)
}

type Plan struct {
	Name          string  `pulumi:"name"`
	Product       string  `pulumi:"product"`
	PromotionCode *string `pulumi:"promotionCode"`
	Publisher     string  `pulumi:"publisher"`
	Version       *string `pulumi:"version"`
}





type PlanInput interface {
	pulumi.Input

	ToPlanOutput() PlanOutput
	ToPlanOutputWithContext(context.Context) PlanOutput
}

type PlanArgs struct {
	Name          pulumi.StringInput    `pulumi:"name"`
	Product       pulumi.StringInput    `pulumi:"product"`
	PromotionCode pulumi.StringPtrInput `pulumi:"promotionCode"`
	Publisher     pulumi.StringInput    `pulumi:"publisher"`
	Version       pulumi.StringPtrInput `pulumi:"version"`
}

func (PlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Plan)(nil)).Elem()
}

func (i PlanArgs) ToPlanOutput() PlanOutput {
	return i.ToPlanOutputWithContext(context.Background())
}

func (i PlanArgs) ToPlanOutputWithContext(ctx context.Context) PlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanOutput)
}

func (i PlanArgs) ToPlanPtrOutput() PlanPtrOutput {
	return i.ToPlanPtrOutputWithContext(context.Background())
}

func (i PlanArgs) ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanOutput).ToPlanPtrOutputWithContext(ctx)
}









type PlanPtrInput interface {
	pulumi.Input

	ToPlanPtrOutput() PlanPtrOutput
	ToPlanPtrOutputWithContext(context.Context) PlanPtrOutput
}

type planPtrType PlanArgs

func PlanPtr(v *PlanArgs) PlanPtrInput {
	return (*planPtrType)(v)
}

func (*planPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Plan)(nil)).Elem()
}

func (i *planPtrType) ToPlanPtrOutput() PlanPtrOutput {
	return i.ToPlanPtrOutputWithContext(context.Background())
}

func (i *planPtrType) ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanPtrOutput)
}

type PlanOutput struct{ *pulumi.OutputState }

func (PlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Plan)(nil)).Elem()
}

func (o PlanOutput) ToPlanOutput() PlanOutput {
	return o
}

func (o PlanOutput) ToPlanOutputWithContext(ctx context.Context) PlanOutput {
	return o
}

func (o PlanOutput) ToPlanPtrOutput() PlanPtrOutput {
	return o.ToPlanPtrOutputWithContext(context.Background())
}

func (o PlanOutput) ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Plan) *Plan {
		return &v
	}).(PlanPtrOutput)
}

func (o PlanOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Plan) string { return v.Name }).(pulumi.StringOutput)
}

func (o PlanOutput) Product() pulumi.StringOutput {
	return o.ApplyT(func(v Plan) string { return v.Product }).(pulumi.StringOutput)
}

func (o PlanOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Plan) *string { return v.PromotionCode }).(pulumi.StringPtrOutput)
}

func (o PlanOutput) Publisher() pulumi.StringOutput {
	return o.ApplyT(func(v Plan) string { return v.Publisher }).(pulumi.StringOutput)
}

func (o PlanOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Plan) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type PlanPtrOutput struct{ *pulumi.OutputState }

func (PlanPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Plan)(nil)).Elem()
}

func (o PlanPtrOutput) ToPlanPtrOutput() PlanPtrOutput {
	return o
}

func (o PlanPtrOutput) ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput {
	return o
}

func (o PlanPtrOutput) Elem() PlanOutput {
	return o.ApplyT(func(v *Plan) Plan {
		if v != nil {
			return *v
		}
		var ret Plan
		return ret
	}).(PlanOutput)
}

func (o PlanPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Plan) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PlanPtrOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Plan) *string {
		if v == nil {
			return nil
		}
		return &v.Product
	}).(pulumi.StringPtrOutput)
}

func (o PlanPtrOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Plan) *string {
		if v == nil {
			return nil
		}
		return v.PromotionCode
	}).(pulumi.StringPtrOutput)
}

func (o PlanPtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Plan) *string {
		if v == nil {
			return nil
		}
		return &v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o PlanPtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Plan) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type PlanResponse struct {
	Name          string  `pulumi:"name"`
	Product       string  `pulumi:"product"`
	PromotionCode *string `pulumi:"promotionCode"`
	Publisher     string  `pulumi:"publisher"`
	Version       *string `pulumi:"version"`
}

type PlanResponseOutput struct{ *pulumi.OutputState }

func (PlanResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PlanResponse)(nil)).Elem()
}

func (o PlanResponseOutput) ToPlanResponseOutput() PlanResponseOutput {
	return o
}

func (o PlanResponseOutput) ToPlanResponseOutputWithContext(ctx context.Context) PlanResponseOutput {
	return o
}

func (o PlanResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PlanResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PlanResponseOutput) Product() pulumi.StringOutput {
	return o.ApplyT(func(v PlanResponse) string { return v.Product }).(pulumi.StringOutput)
}

func (o PlanResponseOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlanResponse) *string { return v.PromotionCode }).(pulumi.StringPtrOutput)
}

func (o PlanResponseOutput) Publisher() pulumi.StringOutput {
	return o.ApplyT(func(v PlanResponse) string { return v.Publisher }).(pulumi.StringOutput)
}

func (o PlanResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlanResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type PlanResponsePtrOutput struct{ *pulumi.OutputState }

func (PlanResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PlanResponse)(nil)).Elem()
}

func (o PlanResponsePtrOutput) ToPlanResponsePtrOutput() PlanResponsePtrOutput {
	return o
}

func (o PlanResponsePtrOutput) ToPlanResponsePtrOutputWithContext(ctx context.Context) PlanResponsePtrOutput {
	return o
}

func (o PlanResponsePtrOutput) Elem() PlanResponseOutput {
	return o.ApplyT(func(v *PlanResponse) PlanResponse {
		if v != nil {
			return *v
		}
		var ret PlanResponse
		return ret
	}).(PlanResponseOutput)
}

func (o PlanResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PlanResponsePtrOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Product
	}).(pulumi.StringPtrOutput)
}

func (o PlanResponsePtrOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.PromotionCode
	}).(pulumi.StringPtrOutput)
}

func (o PlanResponsePtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o PlanResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type RepositoryRefDefinition struct {
	Branch *string `pulumi:"branch"`
	Commit *string `pulumi:"commit"`
	Semver *string `pulumi:"semver"`
	Tag    *string `pulumi:"tag"`
}





type RepositoryRefDefinitionInput interface {
	pulumi.Input

	ToRepositoryRefDefinitionOutput() RepositoryRefDefinitionOutput
	ToRepositoryRefDefinitionOutputWithContext(context.Context) RepositoryRefDefinitionOutput
}

type RepositoryRefDefinitionArgs struct {
	Branch pulumi.StringPtrInput `pulumi:"branch"`
	Commit pulumi.StringPtrInput `pulumi:"commit"`
	Semver pulumi.StringPtrInput `pulumi:"semver"`
	Tag    pulumi.StringPtrInput `pulumi:"tag"`
}

func (RepositoryRefDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RepositoryRefDefinition)(nil)).Elem()
}

func (i RepositoryRefDefinitionArgs) ToRepositoryRefDefinitionOutput() RepositoryRefDefinitionOutput {
	return i.ToRepositoryRefDefinitionOutputWithContext(context.Background())
}

func (i RepositoryRefDefinitionArgs) ToRepositoryRefDefinitionOutputWithContext(ctx context.Context) RepositoryRefDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryRefDefinitionOutput)
}

func (i RepositoryRefDefinitionArgs) ToRepositoryRefDefinitionPtrOutput() RepositoryRefDefinitionPtrOutput {
	return i.ToRepositoryRefDefinitionPtrOutputWithContext(context.Background())
}

func (i RepositoryRefDefinitionArgs) ToRepositoryRefDefinitionPtrOutputWithContext(ctx context.Context) RepositoryRefDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryRefDefinitionOutput).ToRepositoryRefDefinitionPtrOutputWithContext(ctx)
}









type RepositoryRefDefinitionPtrInput interface {
	pulumi.Input

	ToRepositoryRefDefinitionPtrOutput() RepositoryRefDefinitionPtrOutput
	ToRepositoryRefDefinitionPtrOutputWithContext(context.Context) RepositoryRefDefinitionPtrOutput
}

type repositoryRefDefinitionPtrType RepositoryRefDefinitionArgs

func RepositoryRefDefinitionPtr(v *RepositoryRefDefinitionArgs) RepositoryRefDefinitionPtrInput {
	return (*repositoryRefDefinitionPtrType)(v)
}

func (*repositoryRefDefinitionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryRefDefinition)(nil)).Elem()
}

func (i *repositoryRefDefinitionPtrType) ToRepositoryRefDefinitionPtrOutput() RepositoryRefDefinitionPtrOutput {
	return i.ToRepositoryRefDefinitionPtrOutputWithContext(context.Background())
}

func (i *repositoryRefDefinitionPtrType) ToRepositoryRefDefinitionPtrOutputWithContext(ctx context.Context) RepositoryRefDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryRefDefinitionPtrOutput)
}

type RepositoryRefDefinitionOutput struct{ *pulumi.OutputState }

func (RepositoryRefDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RepositoryRefDefinition)(nil)).Elem()
}

func (o RepositoryRefDefinitionOutput) ToRepositoryRefDefinitionOutput() RepositoryRefDefinitionOutput {
	return o
}

func (o RepositoryRefDefinitionOutput) ToRepositoryRefDefinitionOutputWithContext(ctx context.Context) RepositoryRefDefinitionOutput {
	return o
}

func (o RepositoryRefDefinitionOutput) ToRepositoryRefDefinitionPtrOutput() RepositoryRefDefinitionPtrOutput {
	return o.ToRepositoryRefDefinitionPtrOutputWithContext(context.Background())
}

func (o RepositoryRefDefinitionOutput) ToRepositoryRefDefinitionPtrOutputWithContext(ctx context.Context) RepositoryRefDefinitionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RepositoryRefDefinition) *RepositoryRefDefinition {
		return &v
	}).(RepositoryRefDefinitionPtrOutput)
}

func (o RepositoryRefDefinitionOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RepositoryRefDefinition) *string { return v.Branch }).(pulumi.StringPtrOutput)
}

func (o RepositoryRefDefinitionOutput) Commit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RepositoryRefDefinition) *string { return v.Commit }).(pulumi.StringPtrOutput)
}

func (o RepositoryRefDefinitionOutput) Semver() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RepositoryRefDefinition) *string { return v.Semver }).(pulumi.StringPtrOutput)
}

func (o RepositoryRefDefinitionOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RepositoryRefDefinition) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

type RepositoryRefDefinitionPtrOutput struct{ *pulumi.OutputState }

func (RepositoryRefDefinitionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryRefDefinition)(nil)).Elem()
}

func (o RepositoryRefDefinitionPtrOutput) ToRepositoryRefDefinitionPtrOutput() RepositoryRefDefinitionPtrOutput {
	return o
}

func (o RepositoryRefDefinitionPtrOutput) ToRepositoryRefDefinitionPtrOutputWithContext(ctx context.Context) RepositoryRefDefinitionPtrOutput {
	return o
}

func (o RepositoryRefDefinitionPtrOutput) Elem() RepositoryRefDefinitionOutput {
	return o.ApplyT(func(v *RepositoryRefDefinition) RepositoryRefDefinition {
		if v != nil {
			return *v
		}
		var ret RepositoryRefDefinition
		return ret
	}).(RepositoryRefDefinitionOutput)
}

func (o RepositoryRefDefinitionPtrOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryRefDefinition) *string {
		if v == nil {
			return nil
		}
		return v.Branch
	}).(pulumi.StringPtrOutput)
}

func (o RepositoryRefDefinitionPtrOutput) Commit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryRefDefinition) *string {
		if v == nil {
			return nil
		}
		return v.Commit
	}).(pulumi.StringPtrOutput)
}

func (o RepositoryRefDefinitionPtrOutput) Semver() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryRefDefinition) *string {
		if v == nil {
			return nil
		}
		return v.Semver
	}).(pulumi.StringPtrOutput)
}

func (o RepositoryRefDefinitionPtrOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryRefDefinition) *string {
		if v == nil {
			return nil
		}
		return v.Tag
	}).(pulumi.StringPtrOutput)
}

type RepositoryRefDefinitionResponse struct {
	Branch *string `pulumi:"branch"`
	Commit *string `pulumi:"commit"`
	Semver *string `pulumi:"semver"`
	Tag    *string `pulumi:"tag"`
}

type RepositoryRefDefinitionResponseOutput struct{ *pulumi.OutputState }

func (RepositoryRefDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RepositoryRefDefinitionResponse)(nil)).Elem()
}

func (o RepositoryRefDefinitionResponseOutput) ToRepositoryRefDefinitionResponseOutput() RepositoryRefDefinitionResponseOutput {
	return o
}

func (o RepositoryRefDefinitionResponseOutput) ToRepositoryRefDefinitionResponseOutputWithContext(ctx context.Context) RepositoryRefDefinitionResponseOutput {
	return o
}

func (o RepositoryRefDefinitionResponseOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RepositoryRefDefinitionResponse) *string { return v.Branch }).(pulumi.StringPtrOutput)
}

func (o RepositoryRefDefinitionResponseOutput) Commit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RepositoryRefDefinitionResponse) *string { return v.Commit }).(pulumi.StringPtrOutput)
}

func (o RepositoryRefDefinitionResponseOutput) Semver() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RepositoryRefDefinitionResponse) *string { return v.Semver }).(pulumi.StringPtrOutput)
}

func (o RepositoryRefDefinitionResponseOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RepositoryRefDefinitionResponse) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

type RepositoryRefDefinitionResponsePtrOutput struct{ *pulumi.OutputState }

func (RepositoryRefDefinitionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryRefDefinitionResponse)(nil)).Elem()
}

func (o RepositoryRefDefinitionResponsePtrOutput) ToRepositoryRefDefinitionResponsePtrOutput() RepositoryRefDefinitionResponsePtrOutput {
	return o
}

func (o RepositoryRefDefinitionResponsePtrOutput) ToRepositoryRefDefinitionResponsePtrOutputWithContext(ctx context.Context) RepositoryRefDefinitionResponsePtrOutput {
	return o
}

func (o RepositoryRefDefinitionResponsePtrOutput) Elem() RepositoryRefDefinitionResponseOutput {
	return o.ApplyT(func(v *RepositoryRefDefinitionResponse) RepositoryRefDefinitionResponse {
		if v != nil {
			return *v
		}
		var ret RepositoryRefDefinitionResponse
		return ret
	}).(RepositoryRefDefinitionResponseOutput)
}

func (o RepositoryRefDefinitionResponsePtrOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryRefDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Branch
	}).(pulumi.StringPtrOutput)
}

func (o RepositoryRefDefinitionResponsePtrOutput) Commit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryRefDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Commit
	}).(pulumi.StringPtrOutput)
}

func (o RepositoryRefDefinitionResponsePtrOutput) Semver() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryRefDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Semver
	}).(pulumi.StringPtrOutput)
}

func (o RepositoryRefDefinitionResponsePtrOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryRefDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tag
	}).(pulumi.StringPtrOutput)
}

type Scope struct {
	Cluster   *ScopeCluster   `pulumi:"cluster"`
	Namespace *ScopeNamespace `pulumi:"namespace"`
}





type ScopeInput interface {
	pulumi.Input

	ToScopeOutput() ScopeOutput
	ToScopeOutputWithContext(context.Context) ScopeOutput
}

type ScopeArgs struct {
	Cluster   ScopeClusterPtrInput   `pulumi:"cluster"`
	Namespace ScopeNamespacePtrInput `pulumi:"namespace"`
}

func (ScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Scope)(nil)).Elem()
}

func (i ScopeArgs) ToScopeOutput() ScopeOutput {
	return i.ToScopeOutputWithContext(context.Background())
}

func (i ScopeArgs) ToScopeOutputWithContext(ctx context.Context) ScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeOutput)
}

func (i ScopeArgs) ToScopePtrOutput() ScopePtrOutput {
	return i.ToScopePtrOutputWithContext(context.Background())
}

func (i ScopeArgs) ToScopePtrOutputWithContext(ctx context.Context) ScopePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeOutput).ToScopePtrOutputWithContext(ctx)
}









type ScopePtrInput interface {
	pulumi.Input

	ToScopePtrOutput() ScopePtrOutput
	ToScopePtrOutputWithContext(context.Context) ScopePtrOutput
}

type scopePtrType ScopeArgs

func ScopePtr(v *ScopeArgs) ScopePtrInput {
	return (*scopePtrType)(v)
}

func (*scopePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Scope)(nil)).Elem()
}

func (i *scopePtrType) ToScopePtrOutput() ScopePtrOutput {
	return i.ToScopePtrOutputWithContext(context.Background())
}

func (i *scopePtrType) ToScopePtrOutputWithContext(ctx context.Context) ScopePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopePtrOutput)
}

type ScopeOutput struct{ *pulumi.OutputState }

func (ScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Scope)(nil)).Elem()
}

func (o ScopeOutput) ToScopeOutput() ScopeOutput {
	return o
}

func (o ScopeOutput) ToScopeOutputWithContext(ctx context.Context) ScopeOutput {
	return o
}

func (o ScopeOutput) ToScopePtrOutput() ScopePtrOutput {
	return o.ToScopePtrOutputWithContext(context.Background())
}

func (o ScopeOutput) ToScopePtrOutputWithContext(ctx context.Context) ScopePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Scope) *Scope {
		return &v
	}).(ScopePtrOutput)
}

func (o ScopeOutput) Cluster() ScopeClusterPtrOutput {
	return o.ApplyT(func(v Scope) *ScopeCluster { return v.Cluster }).(ScopeClusterPtrOutput)
}

func (o ScopeOutput) Namespace() ScopeNamespacePtrOutput {
	return o.ApplyT(func(v Scope) *ScopeNamespace { return v.Namespace }).(ScopeNamespacePtrOutput)
}

type ScopePtrOutput struct{ *pulumi.OutputState }

func (ScopePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Scope)(nil)).Elem()
}

func (o ScopePtrOutput) ToScopePtrOutput() ScopePtrOutput {
	return o
}

func (o ScopePtrOutput) ToScopePtrOutputWithContext(ctx context.Context) ScopePtrOutput {
	return o
}

func (o ScopePtrOutput) Elem() ScopeOutput {
	return o.ApplyT(func(v *Scope) Scope {
		if v != nil {
			return *v
		}
		var ret Scope
		return ret
	}).(ScopeOutput)
}

func (o ScopePtrOutput) Cluster() ScopeClusterPtrOutput {
	return o.ApplyT(func(v *Scope) *ScopeCluster {
		if v == nil {
			return nil
		}
		return v.Cluster
	}).(ScopeClusterPtrOutput)
}

func (o ScopePtrOutput) Namespace() ScopeNamespacePtrOutput {
	return o.ApplyT(func(v *Scope) *ScopeNamespace {
		if v == nil {
			return nil
		}
		return v.Namespace
	}).(ScopeNamespacePtrOutput)
}

type ScopeCluster struct {
	ReleaseNamespace *string `pulumi:"releaseNamespace"`
}





type ScopeClusterInput interface {
	pulumi.Input

	ToScopeClusterOutput() ScopeClusterOutput
	ToScopeClusterOutputWithContext(context.Context) ScopeClusterOutput
}

type ScopeClusterArgs struct {
	ReleaseNamespace pulumi.StringPtrInput `pulumi:"releaseNamespace"`
}

func (ScopeClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScopeCluster)(nil)).Elem()
}

func (i ScopeClusterArgs) ToScopeClusterOutput() ScopeClusterOutput {
	return i.ToScopeClusterOutputWithContext(context.Background())
}

func (i ScopeClusterArgs) ToScopeClusterOutputWithContext(ctx context.Context) ScopeClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeClusterOutput)
}

func (i ScopeClusterArgs) ToScopeClusterPtrOutput() ScopeClusterPtrOutput {
	return i.ToScopeClusterPtrOutputWithContext(context.Background())
}

func (i ScopeClusterArgs) ToScopeClusterPtrOutputWithContext(ctx context.Context) ScopeClusterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeClusterOutput).ToScopeClusterPtrOutputWithContext(ctx)
}









type ScopeClusterPtrInput interface {
	pulumi.Input

	ToScopeClusterPtrOutput() ScopeClusterPtrOutput
	ToScopeClusterPtrOutputWithContext(context.Context) ScopeClusterPtrOutput
}

type scopeClusterPtrType ScopeClusterArgs

func ScopeClusterPtr(v *ScopeClusterArgs) ScopeClusterPtrInput {
	return (*scopeClusterPtrType)(v)
}

func (*scopeClusterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopeCluster)(nil)).Elem()
}

func (i *scopeClusterPtrType) ToScopeClusterPtrOutput() ScopeClusterPtrOutput {
	return i.ToScopeClusterPtrOutputWithContext(context.Background())
}

func (i *scopeClusterPtrType) ToScopeClusterPtrOutputWithContext(ctx context.Context) ScopeClusterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeClusterPtrOutput)
}

type ScopeClusterOutput struct{ *pulumi.OutputState }

func (ScopeClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScopeCluster)(nil)).Elem()
}

func (o ScopeClusterOutput) ToScopeClusterOutput() ScopeClusterOutput {
	return o
}

func (o ScopeClusterOutput) ToScopeClusterOutputWithContext(ctx context.Context) ScopeClusterOutput {
	return o
}

func (o ScopeClusterOutput) ToScopeClusterPtrOutput() ScopeClusterPtrOutput {
	return o.ToScopeClusterPtrOutputWithContext(context.Background())
}

func (o ScopeClusterOutput) ToScopeClusterPtrOutputWithContext(ctx context.Context) ScopeClusterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScopeCluster) *ScopeCluster {
		return &v
	}).(ScopeClusterPtrOutput)
}

func (o ScopeClusterOutput) ReleaseNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScopeCluster) *string { return v.ReleaseNamespace }).(pulumi.StringPtrOutput)
}

type ScopeClusterPtrOutput struct{ *pulumi.OutputState }

func (ScopeClusterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopeCluster)(nil)).Elem()
}

func (o ScopeClusterPtrOutput) ToScopeClusterPtrOutput() ScopeClusterPtrOutput {
	return o
}

func (o ScopeClusterPtrOutput) ToScopeClusterPtrOutputWithContext(ctx context.Context) ScopeClusterPtrOutput {
	return o
}

func (o ScopeClusterPtrOutput) Elem() ScopeClusterOutput {
	return o.ApplyT(func(v *ScopeCluster) ScopeCluster {
		if v != nil {
			return *v
		}
		var ret ScopeCluster
		return ret
	}).(ScopeClusterOutput)
}

func (o ScopeClusterPtrOutput) ReleaseNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScopeCluster) *string {
		if v == nil {
			return nil
		}
		return v.ReleaseNamespace
	}).(pulumi.StringPtrOutput)
}

type ScopeClusterResponse struct {
	ReleaseNamespace *string `pulumi:"releaseNamespace"`
}

type ScopeClusterResponseOutput struct{ *pulumi.OutputState }

func (ScopeClusterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScopeClusterResponse)(nil)).Elem()
}

func (o ScopeClusterResponseOutput) ToScopeClusterResponseOutput() ScopeClusterResponseOutput {
	return o
}

func (o ScopeClusterResponseOutput) ToScopeClusterResponseOutputWithContext(ctx context.Context) ScopeClusterResponseOutput {
	return o
}

func (o ScopeClusterResponseOutput) ReleaseNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScopeClusterResponse) *string { return v.ReleaseNamespace }).(pulumi.StringPtrOutput)
}

type ScopeClusterResponsePtrOutput struct{ *pulumi.OutputState }

func (ScopeClusterResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopeClusterResponse)(nil)).Elem()
}

func (o ScopeClusterResponsePtrOutput) ToScopeClusterResponsePtrOutput() ScopeClusterResponsePtrOutput {
	return o
}

func (o ScopeClusterResponsePtrOutput) ToScopeClusterResponsePtrOutputWithContext(ctx context.Context) ScopeClusterResponsePtrOutput {
	return o
}

func (o ScopeClusterResponsePtrOutput) Elem() ScopeClusterResponseOutput {
	return o.ApplyT(func(v *ScopeClusterResponse) ScopeClusterResponse {
		if v != nil {
			return *v
		}
		var ret ScopeClusterResponse
		return ret
	}).(ScopeClusterResponseOutput)
}

func (o ScopeClusterResponsePtrOutput) ReleaseNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScopeClusterResponse) *string {
		if v == nil {
			return nil
		}
		return v.ReleaseNamespace
	}).(pulumi.StringPtrOutput)
}

type ScopeNamespace struct {
	TargetNamespace *string `pulumi:"targetNamespace"`
}





type ScopeNamespaceInput interface {
	pulumi.Input

	ToScopeNamespaceOutput() ScopeNamespaceOutput
	ToScopeNamespaceOutputWithContext(context.Context) ScopeNamespaceOutput
}

type ScopeNamespaceArgs struct {
	TargetNamespace pulumi.StringPtrInput `pulumi:"targetNamespace"`
}

func (ScopeNamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScopeNamespace)(nil)).Elem()
}

func (i ScopeNamespaceArgs) ToScopeNamespaceOutput() ScopeNamespaceOutput {
	return i.ToScopeNamespaceOutputWithContext(context.Background())
}

func (i ScopeNamespaceArgs) ToScopeNamespaceOutputWithContext(ctx context.Context) ScopeNamespaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeNamespaceOutput)
}

func (i ScopeNamespaceArgs) ToScopeNamespacePtrOutput() ScopeNamespacePtrOutput {
	return i.ToScopeNamespacePtrOutputWithContext(context.Background())
}

func (i ScopeNamespaceArgs) ToScopeNamespacePtrOutputWithContext(ctx context.Context) ScopeNamespacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeNamespaceOutput).ToScopeNamespacePtrOutputWithContext(ctx)
}









type ScopeNamespacePtrInput interface {
	pulumi.Input

	ToScopeNamespacePtrOutput() ScopeNamespacePtrOutput
	ToScopeNamespacePtrOutputWithContext(context.Context) ScopeNamespacePtrOutput
}

type scopeNamespacePtrType ScopeNamespaceArgs

func ScopeNamespacePtr(v *ScopeNamespaceArgs) ScopeNamespacePtrInput {
	return (*scopeNamespacePtrType)(v)
}

func (*scopeNamespacePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopeNamespace)(nil)).Elem()
}

func (i *scopeNamespacePtrType) ToScopeNamespacePtrOutput() ScopeNamespacePtrOutput {
	return i.ToScopeNamespacePtrOutputWithContext(context.Background())
}

func (i *scopeNamespacePtrType) ToScopeNamespacePtrOutputWithContext(ctx context.Context) ScopeNamespacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeNamespacePtrOutput)
}

type ScopeNamespaceOutput struct{ *pulumi.OutputState }

func (ScopeNamespaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScopeNamespace)(nil)).Elem()
}

func (o ScopeNamespaceOutput) ToScopeNamespaceOutput() ScopeNamespaceOutput {
	return o
}

func (o ScopeNamespaceOutput) ToScopeNamespaceOutputWithContext(ctx context.Context) ScopeNamespaceOutput {
	return o
}

func (o ScopeNamespaceOutput) ToScopeNamespacePtrOutput() ScopeNamespacePtrOutput {
	return o.ToScopeNamespacePtrOutputWithContext(context.Background())
}

func (o ScopeNamespaceOutput) ToScopeNamespacePtrOutputWithContext(ctx context.Context) ScopeNamespacePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScopeNamespace) *ScopeNamespace {
		return &v
	}).(ScopeNamespacePtrOutput)
}

func (o ScopeNamespaceOutput) TargetNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScopeNamespace) *string { return v.TargetNamespace }).(pulumi.StringPtrOutput)
}

type ScopeNamespacePtrOutput struct{ *pulumi.OutputState }

func (ScopeNamespacePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopeNamespace)(nil)).Elem()
}

func (o ScopeNamespacePtrOutput) ToScopeNamespacePtrOutput() ScopeNamespacePtrOutput {
	return o
}

func (o ScopeNamespacePtrOutput) ToScopeNamespacePtrOutputWithContext(ctx context.Context) ScopeNamespacePtrOutput {
	return o
}

func (o ScopeNamespacePtrOutput) Elem() ScopeNamespaceOutput {
	return o.ApplyT(func(v *ScopeNamespace) ScopeNamespace {
		if v != nil {
			return *v
		}
		var ret ScopeNamespace
		return ret
	}).(ScopeNamespaceOutput)
}

func (o ScopeNamespacePtrOutput) TargetNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScopeNamespace) *string {
		if v == nil {
			return nil
		}
		return v.TargetNamespace
	}).(pulumi.StringPtrOutput)
}

type ScopeNamespaceResponse struct {
	TargetNamespace *string `pulumi:"targetNamespace"`
}

type ScopeNamespaceResponseOutput struct{ *pulumi.OutputState }

func (ScopeNamespaceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScopeNamespaceResponse)(nil)).Elem()
}

func (o ScopeNamespaceResponseOutput) ToScopeNamespaceResponseOutput() ScopeNamespaceResponseOutput {
	return o
}

func (o ScopeNamespaceResponseOutput) ToScopeNamespaceResponseOutputWithContext(ctx context.Context) ScopeNamespaceResponseOutput {
	return o
}

func (o ScopeNamespaceResponseOutput) TargetNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScopeNamespaceResponse) *string { return v.TargetNamespace }).(pulumi.StringPtrOutput)
}

type ScopeNamespaceResponsePtrOutput struct{ *pulumi.OutputState }

func (ScopeNamespaceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopeNamespaceResponse)(nil)).Elem()
}

func (o ScopeNamespaceResponsePtrOutput) ToScopeNamespaceResponsePtrOutput() ScopeNamespaceResponsePtrOutput {
	return o
}

func (o ScopeNamespaceResponsePtrOutput) ToScopeNamespaceResponsePtrOutputWithContext(ctx context.Context) ScopeNamespaceResponsePtrOutput {
	return o
}

func (o ScopeNamespaceResponsePtrOutput) Elem() ScopeNamespaceResponseOutput {
	return o.ApplyT(func(v *ScopeNamespaceResponse) ScopeNamespaceResponse {
		if v != nil {
			return *v
		}
		var ret ScopeNamespaceResponse
		return ret
	}).(ScopeNamespaceResponseOutput)
}

func (o ScopeNamespaceResponsePtrOutput) TargetNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScopeNamespaceResponse) *string {
		if v == nil {
			return nil
		}
		return v.TargetNamespace
	}).(pulumi.StringPtrOutput)
}

type ScopeResponse struct {
	Cluster   *ScopeClusterResponse   `pulumi:"cluster"`
	Namespace *ScopeNamespaceResponse `pulumi:"namespace"`
}

type ScopeResponseOutput struct{ *pulumi.OutputState }

func (ScopeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScopeResponse)(nil)).Elem()
}

func (o ScopeResponseOutput) ToScopeResponseOutput() ScopeResponseOutput {
	return o
}

func (o ScopeResponseOutput) ToScopeResponseOutputWithContext(ctx context.Context) ScopeResponseOutput {
	return o
}

func (o ScopeResponseOutput) Cluster() ScopeClusterResponsePtrOutput {
	return o.ApplyT(func(v ScopeResponse) *ScopeClusterResponse { return v.Cluster }).(ScopeClusterResponsePtrOutput)
}

func (o ScopeResponseOutput) Namespace() ScopeNamespaceResponsePtrOutput {
	return o.ApplyT(func(v ScopeResponse) *ScopeNamespaceResponse { return v.Namespace }).(ScopeNamespaceResponsePtrOutput)
}

type ScopeResponsePtrOutput struct{ *pulumi.OutputState }

func (ScopeResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopeResponse)(nil)).Elem()
}

func (o ScopeResponsePtrOutput) ToScopeResponsePtrOutput() ScopeResponsePtrOutput {
	return o
}

func (o ScopeResponsePtrOutput) ToScopeResponsePtrOutputWithContext(ctx context.Context) ScopeResponsePtrOutput {
	return o
}

func (o ScopeResponsePtrOutput) Elem() ScopeResponseOutput {
	return o.ApplyT(func(v *ScopeResponse) ScopeResponse {
		if v != nil {
			return *v
		}
		var ret ScopeResponse
		return ret
	}).(ScopeResponseOutput)
}

func (o ScopeResponsePtrOutput) Cluster() ScopeClusterResponsePtrOutput {
	return o.ApplyT(func(v *ScopeResponse) *ScopeClusterResponse {
		if v == nil {
			return nil
		}
		return v.Cluster
	}).(ScopeClusterResponsePtrOutput)
}

func (o ScopeResponsePtrOutput) Namespace() ScopeNamespaceResponsePtrOutput {
	return o.ApplyT(func(v *ScopeResponse) *ScopeNamespaceResponse {
		if v == nil {
			return nil
		}
		return v.Namespace
	}).(ScopeNamespaceResponsePtrOutput)
}

type ServicePrincipalDefinition struct {
	ClientCertificate          *string `pulumi:"clientCertificate"`
	ClientCertificatePassword  *string `pulumi:"clientCertificatePassword"`
	ClientCertificateSendChain *bool   `pulumi:"clientCertificateSendChain"`
	ClientId                   *string `pulumi:"clientId"`
	ClientSecret               *string `pulumi:"clientSecret"`
	TenantId                   *string `pulumi:"tenantId"`
}


func (val *ServicePrincipalDefinition) Defaults() *ServicePrincipalDefinition {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ClientCertificateSendChain) {
		clientCertificateSendChain_ := false
		tmp.ClientCertificateSendChain = &clientCertificateSendChain_
	}
	return &tmp
}





type ServicePrincipalDefinitionInput interface {
	pulumi.Input

	ToServicePrincipalDefinitionOutput() ServicePrincipalDefinitionOutput
	ToServicePrincipalDefinitionOutputWithContext(context.Context) ServicePrincipalDefinitionOutput
}

type ServicePrincipalDefinitionArgs struct {
	ClientCertificate          pulumi.StringPtrInput `pulumi:"clientCertificate"`
	ClientCertificatePassword  pulumi.StringPtrInput `pulumi:"clientCertificatePassword"`
	ClientCertificateSendChain pulumi.BoolPtrInput   `pulumi:"clientCertificateSendChain"`
	ClientId                   pulumi.StringPtrInput `pulumi:"clientId"`
	ClientSecret               pulumi.StringPtrInput `pulumi:"clientSecret"`
	TenantId                   pulumi.StringPtrInput `pulumi:"tenantId"`
}


func (val *ServicePrincipalDefinitionArgs) Defaults() *ServicePrincipalDefinitionArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ClientCertificateSendChain) {
		tmp.ClientCertificateSendChain = pulumi.BoolPtr(false)
	}
	return &tmp
}
func (ServicePrincipalDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePrincipalDefinition)(nil)).Elem()
}

func (i ServicePrincipalDefinitionArgs) ToServicePrincipalDefinitionOutput() ServicePrincipalDefinitionOutput {
	return i.ToServicePrincipalDefinitionOutputWithContext(context.Background())
}

func (i ServicePrincipalDefinitionArgs) ToServicePrincipalDefinitionOutputWithContext(ctx context.Context) ServicePrincipalDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalDefinitionOutput)
}

func (i ServicePrincipalDefinitionArgs) ToServicePrincipalDefinitionPtrOutput() ServicePrincipalDefinitionPtrOutput {
	return i.ToServicePrincipalDefinitionPtrOutputWithContext(context.Background())
}

func (i ServicePrincipalDefinitionArgs) ToServicePrincipalDefinitionPtrOutputWithContext(ctx context.Context) ServicePrincipalDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalDefinitionOutput).ToServicePrincipalDefinitionPtrOutputWithContext(ctx)
}









type ServicePrincipalDefinitionPtrInput interface {
	pulumi.Input

	ToServicePrincipalDefinitionPtrOutput() ServicePrincipalDefinitionPtrOutput
	ToServicePrincipalDefinitionPtrOutputWithContext(context.Context) ServicePrincipalDefinitionPtrOutput
}

type servicePrincipalDefinitionPtrType ServicePrincipalDefinitionArgs

func ServicePrincipalDefinitionPtr(v *ServicePrincipalDefinitionArgs) ServicePrincipalDefinitionPtrInput {
	return (*servicePrincipalDefinitionPtrType)(v)
}

func (*servicePrincipalDefinitionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePrincipalDefinition)(nil)).Elem()
}

func (i *servicePrincipalDefinitionPtrType) ToServicePrincipalDefinitionPtrOutput() ServicePrincipalDefinitionPtrOutput {
	return i.ToServicePrincipalDefinitionPtrOutputWithContext(context.Background())
}

func (i *servicePrincipalDefinitionPtrType) ToServicePrincipalDefinitionPtrOutputWithContext(ctx context.Context) ServicePrincipalDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalDefinitionPtrOutput)
}

type ServicePrincipalDefinitionOutput struct{ *pulumi.OutputState }

func (ServicePrincipalDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePrincipalDefinition)(nil)).Elem()
}

func (o ServicePrincipalDefinitionOutput) ToServicePrincipalDefinitionOutput() ServicePrincipalDefinitionOutput {
	return o
}

func (o ServicePrincipalDefinitionOutput) ToServicePrincipalDefinitionOutputWithContext(ctx context.Context) ServicePrincipalDefinitionOutput {
	return o
}

func (o ServicePrincipalDefinitionOutput) ToServicePrincipalDefinitionPtrOutput() ServicePrincipalDefinitionPtrOutput {
	return o.ToServicePrincipalDefinitionPtrOutputWithContext(context.Background())
}

func (o ServicePrincipalDefinitionOutput) ToServicePrincipalDefinitionPtrOutputWithContext(ctx context.Context) ServicePrincipalDefinitionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServicePrincipalDefinition) *ServicePrincipalDefinition {
		return &v
	}).(ServicePrincipalDefinitionPtrOutput)
}

func (o ServicePrincipalDefinitionOutput) ClientCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServicePrincipalDefinition) *string { return v.ClientCertificate }).(pulumi.StringPtrOutput)
}

func (o ServicePrincipalDefinitionOutput) ClientCertificatePassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServicePrincipalDefinition) *string { return v.ClientCertificatePassword }).(pulumi.StringPtrOutput)
}

func (o ServicePrincipalDefinitionOutput) ClientCertificateSendChain() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ServicePrincipalDefinition) *bool { return v.ClientCertificateSendChain }).(pulumi.BoolPtrOutput)
}

func (o ServicePrincipalDefinitionOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServicePrincipalDefinition) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o ServicePrincipalDefinitionOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServicePrincipalDefinition) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

func (o ServicePrincipalDefinitionOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServicePrincipalDefinition) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type ServicePrincipalDefinitionPtrOutput struct{ *pulumi.OutputState }

func (ServicePrincipalDefinitionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePrincipalDefinition)(nil)).Elem()
}

func (o ServicePrincipalDefinitionPtrOutput) ToServicePrincipalDefinitionPtrOutput() ServicePrincipalDefinitionPtrOutput {
	return o
}

func (o ServicePrincipalDefinitionPtrOutput) ToServicePrincipalDefinitionPtrOutputWithContext(ctx context.Context) ServicePrincipalDefinitionPtrOutput {
	return o
}

func (o ServicePrincipalDefinitionPtrOutput) Elem() ServicePrincipalDefinitionOutput {
	return o.ApplyT(func(v *ServicePrincipalDefinition) ServicePrincipalDefinition {
		if v != nil {
			return *v
		}
		var ret ServicePrincipalDefinition
		return ret
	}).(ServicePrincipalDefinitionOutput)
}

func (o ServicePrincipalDefinitionPtrOutput) ClientCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicePrincipalDefinition) *string {
		if v == nil {
			return nil
		}
		return v.ClientCertificate
	}).(pulumi.StringPtrOutput)
}

func (o ServicePrincipalDefinitionPtrOutput) ClientCertificatePassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicePrincipalDefinition) *string {
		if v == nil {
			return nil
		}
		return v.ClientCertificatePassword
	}).(pulumi.StringPtrOutput)
}

func (o ServicePrincipalDefinitionPtrOutput) ClientCertificateSendChain() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServicePrincipalDefinition) *bool {
		if v == nil {
			return nil
		}
		return v.ClientCertificateSendChain
	}).(pulumi.BoolPtrOutput)
}

func (o ServicePrincipalDefinitionPtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicePrincipalDefinition) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o ServicePrincipalDefinitionPtrOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicePrincipalDefinition) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecret
	}).(pulumi.StringPtrOutput)
}

func (o ServicePrincipalDefinitionPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicePrincipalDefinition) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

type ServicePrincipalDefinitionResponse struct {
	ClientCertificate          *string `pulumi:"clientCertificate"`
	ClientCertificatePassword  *string `pulumi:"clientCertificatePassword"`
	ClientCertificateSendChain *bool   `pulumi:"clientCertificateSendChain"`
	ClientId                   *string `pulumi:"clientId"`
	ClientSecret               *string `pulumi:"clientSecret"`
	TenantId                   *string `pulumi:"tenantId"`
}


func (val *ServicePrincipalDefinitionResponse) Defaults() *ServicePrincipalDefinitionResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ClientCertificateSendChain) {
		clientCertificateSendChain_ := false
		tmp.ClientCertificateSendChain = &clientCertificateSendChain_
	}
	return &tmp
}

type ServicePrincipalDefinitionResponseOutput struct{ *pulumi.OutputState }

func (ServicePrincipalDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePrincipalDefinitionResponse)(nil)).Elem()
}

func (o ServicePrincipalDefinitionResponseOutput) ToServicePrincipalDefinitionResponseOutput() ServicePrincipalDefinitionResponseOutput {
	return o
}

func (o ServicePrincipalDefinitionResponseOutput) ToServicePrincipalDefinitionResponseOutputWithContext(ctx context.Context) ServicePrincipalDefinitionResponseOutput {
	return o
}

func (o ServicePrincipalDefinitionResponseOutput) ClientCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServicePrincipalDefinitionResponse) *string { return v.ClientCertificate }).(pulumi.StringPtrOutput)
}

func (o ServicePrincipalDefinitionResponseOutput) ClientCertificatePassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServicePrincipalDefinitionResponse) *string { return v.ClientCertificatePassword }).(pulumi.StringPtrOutput)
}

func (o ServicePrincipalDefinitionResponseOutput) ClientCertificateSendChain() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ServicePrincipalDefinitionResponse) *bool { return v.ClientCertificateSendChain }).(pulumi.BoolPtrOutput)
}

func (o ServicePrincipalDefinitionResponseOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServicePrincipalDefinitionResponse) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o ServicePrincipalDefinitionResponseOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServicePrincipalDefinitionResponse) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

func (o ServicePrincipalDefinitionResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServicePrincipalDefinitionResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type ServicePrincipalDefinitionResponsePtrOutput struct{ *pulumi.OutputState }

func (ServicePrincipalDefinitionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePrincipalDefinitionResponse)(nil)).Elem()
}

func (o ServicePrincipalDefinitionResponsePtrOutput) ToServicePrincipalDefinitionResponsePtrOutput() ServicePrincipalDefinitionResponsePtrOutput {
	return o
}

func (o ServicePrincipalDefinitionResponsePtrOutput) ToServicePrincipalDefinitionResponsePtrOutputWithContext(ctx context.Context) ServicePrincipalDefinitionResponsePtrOutput {
	return o
}

func (o ServicePrincipalDefinitionResponsePtrOutput) Elem() ServicePrincipalDefinitionResponseOutput {
	return o.ApplyT(func(v *ServicePrincipalDefinitionResponse) ServicePrincipalDefinitionResponse {
		if v != nil {
			return *v
		}
		var ret ServicePrincipalDefinitionResponse
		return ret
	}).(ServicePrincipalDefinitionResponseOutput)
}

func (o ServicePrincipalDefinitionResponsePtrOutput) ClientCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicePrincipalDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientCertificate
	}).(pulumi.StringPtrOutput)
}

func (o ServicePrincipalDefinitionResponsePtrOutput) ClientCertificatePassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicePrincipalDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientCertificatePassword
	}).(pulumi.StringPtrOutput)
}

func (o ServicePrincipalDefinitionResponsePtrOutput) ClientCertificateSendChain() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServicePrincipalDefinitionResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ClientCertificateSendChain
	}).(pulumi.BoolPtrOutput)
}

func (o ServicePrincipalDefinitionResponsePtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicePrincipalDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o ServicePrincipalDefinitionResponsePtrOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicePrincipalDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecret
	}).(pulumi.StringPtrOutput)
}

func (o ServicePrincipalDefinitionResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicePrincipalDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AzureBlobDefinitionOutput{})
	pulumi.RegisterOutputType(AzureBlobDefinitionPtrOutput{})
	pulumi.RegisterOutputType(AzureBlobDefinitionResponseOutput{})
	pulumi.RegisterOutputType(AzureBlobDefinitionResponsePtrOutput{})
	pulumi.RegisterOutputType(BucketDefinitionOutput{})
	pulumi.RegisterOutputType(BucketDefinitionPtrOutput{})
	pulumi.RegisterOutputType(BucketDefinitionResponseOutput{})
	pulumi.RegisterOutputType(BucketDefinitionResponsePtrOutput{})
	pulumi.RegisterOutputType(ComplianceStatusResponseOutput{})
	pulumi.RegisterOutputType(ErrorAdditionalInfoResponseOutput{})
	pulumi.RegisterOutputType(ErrorAdditionalInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(ErrorDetailResponseOutput{})
	pulumi.RegisterOutputType(ErrorDetailResponseArrayOutput{})
	pulumi.RegisterOutputType(ExtensionAksAssignedIdentityOutput{})
	pulumi.RegisterOutputType(ExtensionAksAssignedIdentityPtrOutput{})
	pulumi.RegisterOutputType(ExtensionResponseAksAssignedIdentityOutput{})
	pulumi.RegisterOutputType(ExtensionResponseAksAssignedIdentityPtrOutput{})
	pulumi.RegisterOutputType(ExtensionStatusOutput{})
	pulumi.RegisterOutputType(ExtensionStatusArrayOutput{})
	pulumi.RegisterOutputType(ExtensionStatusResponseOutput{})
	pulumi.RegisterOutputType(ExtensionStatusResponseArrayOutput{})
	pulumi.RegisterOutputType(GitRepositoryDefinitionOutput{})
	pulumi.RegisterOutputType(GitRepositoryDefinitionPtrOutput{})
	pulumi.RegisterOutputType(GitRepositoryDefinitionResponseOutput{})
	pulumi.RegisterOutputType(GitRepositoryDefinitionResponsePtrOutput{})
	pulumi.RegisterOutputType(HelmOperatorPropertiesOutput{})
	pulumi.RegisterOutputType(HelmOperatorPropertiesPtrOutput{})
	pulumi.RegisterOutputType(HelmOperatorPropertiesResponseOutput{})
	pulumi.RegisterOutputType(HelmOperatorPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(HelmReleasePropertiesDefinitionResponseOutput{})
	pulumi.RegisterOutputType(HelmReleasePropertiesDefinitionResponsePtrOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(KustomizationDefinitionOutput{})
	pulumi.RegisterOutputType(KustomizationDefinitionMapOutput{})
	pulumi.RegisterOutputType(KustomizationDefinitionResponseOutput{})
	pulumi.RegisterOutputType(KustomizationDefinitionResponseMapOutput{})
	pulumi.RegisterOutputType(ManagedIdentityDefinitionOutput{})
	pulumi.RegisterOutputType(ManagedIdentityDefinitionPtrOutput{})
	pulumi.RegisterOutputType(ManagedIdentityDefinitionResponseOutput{})
	pulumi.RegisterOutputType(ManagedIdentityDefinitionResponsePtrOutput{})
	pulumi.RegisterOutputType(ObjectReferenceDefinitionResponseOutput{})
	pulumi.RegisterOutputType(ObjectReferenceDefinitionResponsePtrOutput{})
	pulumi.RegisterOutputType(ObjectStatusConditionDefinitionResponseOutput{})
	pulumi.RegisterOutputType(ObjectStatusConditionDefinitionResponseArrayOutput{})
	pulumi.RegisterOutputType(ObjectStatusDefinitionResponseOutput{})
	pulumi.RegisterOutputType(ObjectStatusDefinitionResponseArrayOutput{})
	pulumi.RegisterOutputType(PlanOutput{})
	pulumi.RegisterOutputType(PlanPtrOutput{})
	pulumi.RegisterOutputType(PlanResponseOutput{})
	pulumi.RegisterOutputType(PlanResponsePtrOutput{})
	pulumi.RegisterOutputType(RepositoryRefDefinitionOutput{})
	pulumi.RegisterOutputType(RepositoryRefDefinitionPtrOutput{})
	pulumi.RegisterOutputType(RepositoryRefDefinitionResponseOutput{})
	pulumi.RegisterOutputType(RepositoryRefDefinitionResponsePtrOutput{})
	pulumi.RegisterOutputType(ScopeOutput{})
	pulumi.RegisterOutputType(ScopePtrOutput{})
	pulumi.RegisterOutputType(ScopeClusterOutput{})
	pulumi.RegisterOutputType(ScopeClusterPtrOutput{})
	pulumi.RegisterOutputType(ScopeClusterResponseOutput{})
	pulumi.RegisterOutputType(ScopeClusterResponsePtrOutput{})
	pulumi.RegisterOutputType(ScopeNamespaceOutput{})
	pulumi.RegisterOutputType(ScopeNamespacePtrOutput{})
	pulumi.RegisterOutputType(ScopeNamespaceResponseOutput{})
	pulumi.RegisterOutputType(ScopeNamespaceResponsePtrOutput{})
	pulumi.RegisterOutputType(ScopeResponseOutput{})
	pulumi.RegisterOutputType(ScopeResponsePtrOutput{})
	pulumi.RegisterOutputType(ServicePrincipalDefinitionOutput{})
	pulumi.RegisterOutputType(ServicePrincipalDefinitionPtrOutput{})
	pulumi.RegisterOutputType(ServicePrincipalDefinitionResponseOutput{})
	pulumi.RegisterOutputType(ServicePrincipalDefinitionResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
