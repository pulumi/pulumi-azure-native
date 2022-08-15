


package v20161201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CustomDomain struct {
	Name             string `pulumi:"name"`
	UseSubDomainName *bool  `pulumi:"useSubDomainName"`
}





type CustomDomainInput interface {
	pulumi.Input

	ToCustomDomainOutput() CustomDomainOutput
	ToCustomDomainOutputWithContext(context.Context) CustomDomainOutput
}

type CustomDomainArgs struct {
	Name             pulumi.StringInput  `pulumi:"name"`
	UseSubDomainName pulumi.BoolPtrInput `pulumi:"useSubDomainName"`
}

func (CustomDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomDomain)(nil)).Elem()
}

func (i CustomDomainArgs) ToCustomDomainOutput() CustomDomainOutput {
	return i.ToCustomDomainOutputWithContext(context.Background())
}

func (i CustomDomainArgs) ToCustomDomainOutputWithContext(ctx context.Context) CustomDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainOutput)
}

func (i CustomDomainArgs) ToCustomDomainPtrOutput() CustomDomainPtrOutput {
	return i.ToCustomDomainPtrOutputWithContext(context.Background())
}

func (i CustomDomainArgs) ToCustomDomainPtrOutputWithContext(ctx context.Context) CustomDomainPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainOutput).ToCustomDomainPtrOutputWithContext(ctx)
}









type CustomDomainPtrInput interface {
	pulumi.Input

	ToCustomDomainPtrOutput() CustomDomainPtrOutput
	ToCustomDomainPtrOutputWithContext(context.Context) CustomDomainPtrOutput
}

type customDomainPtrType CustomDomainArgs

func CustomDomainPtr(v *CustomDomainArgs) CustomDomainPtrInput {
	return (*customDomainPtrType)(v)
}

func (*customDomainPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDomain)(nil)).Elem()
}

func (i *customDomainPtrType) ToCustomDomainPtrOutput() CustomDomainPtrOutput {
	return i.ToCustomDomainPtrOutputWithContext(context.Background())
}

func (i *customDomainPtrType) ToCustomDomainPtrOutputWithContext(ctx context.Context) CustomDomainPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainPtrOutput)
}

type CustomDomainOutput struct{ *pulumi.OutputState }

func (CustomDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomDomain)(nil)).Elem()
}

func (o CustomDomainOutput) ToCustomDomainOutput() CustomDomainOutput {
	return o
}

func (o CustomDomainOutput) ToCustomDomainOutputWithContext(ctx context.Context) CustomDomainOutput {
	return o
}

func (o CustomDomainOutput) ToCustomDomainPtrOutput() CustomDomainPtrOutput {
	return o.ToCustomDomainPtrOutputWithContext(context.Background())
}

func (o CustomDomainOutput) ToCustomDomainPtrOutputWithContext(ctx context.Context) CustomDomainPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomDomain) *CustomDomain {
		return &v
	}).(CustomDomainPtrOutput)
}

func (o CustomDomainOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v CustomDomain) string { return v.Name }).(pulumi.StringOutput)
}

func (o CustomDomainOutput) UseSubDomainName() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CustomDomain) *bool { return v.UseSubDomainName }).(pulumi.BoolPtrOutput)
}

type CustomDomainPtrOutput struct{ *pulumi.OutputState }

func (CustomDomainPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDomain)(nil)).Elem()
}

func (o CustomDomainPtrOutput) ToCustomDomainPtrOutput() CustomDomainPtrOutput {
	return o
}

func (o CustomDomainPtrOutput) ToCustomDomainPtrOutputWithContext(ctx context.Context) CustomDomainPtrOutput {
	return o
}

func (o CustomDomainPtrOutput) Elem() CustomDomainOutput {
	return o.ApplyT(func(v *CustomDomain) CustomDomain {
		if v != nil {
			return *v
		}
		var ret CustomDomain
		return ret
	}).(CustomDomainOutput)
}

func (o CustomDomainPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomDomain) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o CustomDomainPtrOutput) UseSubDomainName() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CustomDomain) *bool {
		if v == nil {
			return nil
		}
		return v.UseSubDomainName
	}).(pulumi.BoolPtrOutput)
}

type CustomDomainResponse struct {
	Name             string `pulumi:"name"`
	UseSubDomainName *bool  `pulumi:"useSubDomainName"`
}

type CustomDomainResponseOutput struct{ *pulumi.OutputState }

func (CustomDomainResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomDomainResponse)(nil)).Elem()
}

func (o CustomDomainResponseOutput) ToCustomDomainResponseOutput() CustomDomainResponseOutput {
	return o
}

func (o CustomDomainResponseOutput) ToCustomDomainResponseOutputWithContext(ctx context.Context) CustomDomainResponseOutput {
	return o
}

func (o CustomDomainResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v CustomDomainResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o CustomDomainResponseOutput) UseSubDomainName() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CustomDomainResponse) *bool { return v.UseSubDomainName }).(pulumi.BoolPtrOutput)
}

type Encryption struct {
	KeySource string              `pulumi:"keySource"`
	Services  *EncryptionServices `pulumi:"services"`
}





type EncryptionInput interface {
	pulumi.Input

	ToEncryptionOutput() EncryptionOutput
	ToEncryptionOutputWithContext(context.Context) EncryptionOutput
}

type EncryptionArgs struct {
	KeySource pulumi.StringInput         `pulumi:"keySource"`
	Services  EncryptionServicesPtrInput `pulumi:"services"`
}

func (EncryptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Encryption)(nil)).Elem()
}

func (i EncryptionArgs) ToEncryptionOutput() EncryptionOutput {
	return i.ToEncryptionOutputWithContext(context.Background())
}

func (i EncryptionArgs) ToEncryptionOutputWithContext(ctx context.Context) EncryptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionOutput)
}

func (i EncryptionArgs) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return i.ToEncryptionPtrOutputWithContext(context.Background())
}

func (i EncryptionArgs) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionOutput).ToEncryptionPtrOutputWithContext(ctx)
}









type EncryptionPtrInput interface {
	pulumi.Input

	ToEncryptionPtrOutput() EncryptionPtrOutput
	ToEncryptionPtrOutputWithContext(context.Context) EncryptionPtrOutput
}

type encryptionPtrType EncryptionArgs

func EncryptionPtr(v *EncryptionArgs) EncryptionPtrInput {
	return (*encryptionPtrType)(v)
}

func (*encryptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Encryption)(nil)).Elem()
}

func (i *encryptionPtrType) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return i.ToEncryptionPtrOutputWithContext(context.Background())
}

func (i *encryptionPtrType) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPtrOutput)
}

type EncryptionOutput struct{ *pulumi.OutputState }

func (EncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Encryption)(nil)).Elem()
}

func (o EncryptionOutput) ToEncryptionOutput() EncryptionOutput {
	return o
}

func (o EncryptionOutput) ToEncryptionOutputWithContext(ctx context.Context) EncryptionOutput {
	return o
}

func (o EncryptionOutput) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return o.ToEncryptionPtrOutputWithContext(context.Background())
}

func (o EncryptionOutput) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Encryption) *Encryption {
		return &v
	}).(EncryptionPtrOutput)
}

func (o EncryptionOutput) KeySource() pulumi.StringOutput {
	return o.ApplyT(func(v Encryption) string { return v.KeySource }).(pulumi.StringOutput)
}

func (o EncryptionOutput) Services() EncryptionServicesPtrOutput {
	return o.ApplyT(func(v Encryption) *EncryptionServices { return v.Services }).(EncryptionServicesPtrOutput)
}

type EncryptionPtrOutput struct{ *pulumi.OutputState }

func (EncryptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Encryption)(nil)).Elem()
}

func (o EncryptionPtrOutput) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return o
}

func (o EncryptionPtrOutput) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return o
}

func (o EncryptionPtrOutput) Elem() EncryptionOutput {
	return o.ApplyT(func(v *Encryption) Encryption {
		if v != nil {
			return *v
		}
		var ret Encryption
		return ret
	}).(EncryptionOutput)
}

func (o EncryptionPtrOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Encryption) *string {
		if v == nil {
			return nil
		}
		return &v.KeySource
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionPtrOutput) Services() EncryptionServicesPtrOutput {
	return o.ApplyT(func(v *Encryption) *EncryptionServices {
		if v == nil {
			return nil
		}
		return v.Services
	}).(EncryptionServicesPtrOutput)
}

type EncryptionResponse struct {
	KeySource string                      `pulumi:"keySource"`
	Services  *EncryptionServicesResponse `pulumi:"services"`
}

type EncryptionResponseOutput struct{ *pulumi.OutputState }

func (EncryptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionResponse)(nil)).Elem()
}

func (o EncryptionResponseOutput) ToEncryptionResponseOutput() EncryptionResponseOutput {
	return o
}

func (o EncryptionResponseOutput) ToEncryptionResponseOutputWithContext(ctx context.Context) EncryptionResponseOutput {
	return o
}

func (o EncryptionResponseOutput) KeySource() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionResponse) string { return v.KeySource }).(pulumi.StringOutput)
}

func (o EncryptionResponseOutput) Services() EncryptionServicesResponsePtrOutput {
	return o.ApplyT(func(v EncryptionResponse) *EncryptionServicesResponse { return v.Services }).(EncryptionServicesResponsePtrOutput)
}

type EncryptionService struct {
	Enabled *bool `pulumi:"enabled"`
}





type EncryptionServiceInput interface {
	pulumi.Input

	ToEncryptionServiceOutput() EncryptionServiceOutput
	ToEncryptionServiceOutputWithContext(context.Context) EncryptionServiceOutput
}

type EncryptionServiceArgs struct {
	Enabled pulumi.BoolPtrInput `pulumi:"enabled"`
}

func (EncryptionServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionService)(nil)).Elem()
}

func (i EncryptionServiceArgs) ToEncryptionServiceOutput() EncryptionServiceOutput {
	return i.ToEncryptionServiceOutputWithContext(context.Background())
}

func (i EncryptionServiceArgs) ToEncryptionServiceOutputWithContext(ctx context.Context) EncryptionServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionServiceOutput)
}

func (i EncryptionServiceArgs) ToEncryptionServicePtrOutput() EncryptionServicePtrOutput {
	return i.ToEncryptionServicePtrOutputWithContext(context.Background())
}

func (i EncryptionServiceArgs) ToEncryptionServicePtrOutputWithContext(ctx context.Context) EncryptionServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionServiceOutput).ToEncryptionServicePtrOutputWithContext(ctx)
}









type EncryptionServicePtrInput interface {
	pulumi.Input

	ToEncryptionServicePtrOutput() EncryptionServicePtrOutput
	ToEncryptionServicePtrOutputWithContext(context.Context) EncryptionServicePtrOutput
}

type encryptionServicePtrType EncryptionServiceArgs

func EncryptionServicePtr(v *EncryptionServiceArgs) EncryptionServicePtrInput {
	return (*encryptionServicePtrType)(v)
}

func (*encryptionServicePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionService)(nil)).Elem()
}

func (i *encryptionServicePtrType) ToEncryptionServicePtrOutput() EncryptionServicePtrOutput {
	return i.ToEncryptionServicePtrOutputWithContext(context.Background())
}

func (i *encryptionServicePtrType) ToEncryptionServicePtrOutputWithContext(ctx context.Context) EncryptionServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionServicePtrOutput)
}

type EncryptionServiceOutput struct{ *pulumi.OutputState }

func (EncryptionServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionService)(nil)).Elem()
}

func (o EncryptionServiceOutput) ToEncryptionServiceOutput() EncryptionServiceOutput {
	return o
}

func (o EncryptionServiceOutput) ToEncryptionServiceOutputWithContext(ctx context.Context) EncryptionServiceOutput {
	return o
}

func (o EncryptionServiceOutput) ToEncryptionServicePtrOutput() EncryptionServicePtrOutput {
	return o.ToEncryptionServicePtrOutputWithContext(context.Background())
}

func (o EncryptionServiceOutput) ToEncryptionServicePtrOutputWithContext(ctx context.Context) EncryptionServicePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionService) *EncryptionService {
		return &v
	}).(EncryptionServicePtrOutput)
}

func (o EncryptionServiceOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EncryptionService) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

type EncryptionServicePtrOutput struct{ *pulumi.OutputState }

func (EncryptionServicePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionService)(nil)).Elem()
}

func (o EncryptionServicePtrOutput) ToEncryptionServicePtrOutput() EncryptionServicePtrOutput {
	return o
}

func (o EncryptionServicePtrOutput) ToEncryptionServicePtrOutputWithContext(ctx context.Context) EncryptionServicePtrOutput {
	return o
}

func (o EncryptionServicePtrOutput) Elem() EncryptionServiceOutput {
	return o.ApplyT(func(v *EncryptionService) EncryptionService {
		if v != nil {
			return *v
		}
		var ret EncryptionService
		return ret
	}).(EncryptionServiceOutput)
}

func (o EncryptionServicePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EncryptionService) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

type EncryptionServiceResponse struct {
	Enabled         *bool  `pulumi:"enabled"`
	LastEnabledTime string `pulumi:"lastEnabledTime"`
}

type EncryptionServiceResponseOutput struct{ *pulumi.OutputState }

func (EncryptionServiceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionServiceResponse)(nil)).Elem()
}

func (o EncryptionServiceResponseOutput) ToEncryptionServiceResponseOutput() EncryptionServiceResponseOutput {
	return o
}

func (o EncryptionServiceResponseOutput) ToEncryptionServiceResponseOutputWithContext(ctx context.Context) EncryptionServiceResponseOutput {
	return o
}

func (o EncryptionServiceResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EncryptionServiceResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o EncryptionServiceResponseOutput) LastEnabledTime() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionServiceResponse) string { return v.LastEnabledTime }).(pulumi.StringOutput)
}

type EncryptionServiceResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionServiceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionServiceResponse)(nil)).Elem()
}

func (o EncryptionServiceResponsePtrOutput) ToEncryptionServiceResponsePtrOutput() EncryptionServiceResponsePtrOutput {
	return o
}

func (o EncryptionServiceResponsePtrOutput) ToEncryptionServiceResponsePtrOutputWithContext(ctx context.Context) EncryptionServiceResponsePtrOutput {
	return o
}

func (o EncryptionServiceResponsePtrOutput) Elem() EncryptionServiceResponseOutput {
	return o.ApplyT(func(v *EncryptionServiceResponse) EncryptionServiceResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionServiceResponse
		return ret
	}).(EncryptionServiceResponseOutput)
}

func (o EncryptionServiceResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EncryptionServiceResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o EncryptionServiceResponsePtrOutput) LastEnabledTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionServiceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastEnabledTime
	}).(pulumi.StringPtrOutput)
}

type EncryptionServices struct {
	Blob *EncryptionService `pulumi:"blob"`
	File *EncryptionService `pulumi:"file"`
}





type EncryptionServicesInput interface {
	pulumi.Input

	ToEncryptionServicesOutput() EncryptionServicesOutput
	ToEncryptionServicesOutputWithContext(context.Context) EncryptionServicesOutput
}

type EncryptionServicesArgs struct {
	Blob EncryptionServicePtrInput `pulumi:"blob"`
	File EncryptionServicePtrInput `pulumi:"file"`
}

func (EncryptionServicesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionServices)(nil)).Elem()
}

func (i EncryptionServicesArgs) ToEncryptionServicesOutput() EncryptionServicesOutput {
	return i.ToEncryptionServicesOutputWithContext(context.Background())
}

func (i EncryptionServicesArgs) ToEncryptionServicesOutputWithContext(ctx context.Context) EncryptionServicesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionServicesOutput)
}

func (i EncryptionServicesArgs) ToEncryptionServicesPtrOutput() EncryptionServicesPtrOutput {
	return i.ToEncryptionServicesPtrOutputWithContext(context.Background())
}

func (i EncryptionServicesArgs) ToEncryptionServicesPtrOutputWithContext(ctx context.Context) EncryptionServicesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionServicesOutput).ToEncryptionServicesPtrOutputWithContext(ctx)
}









type EncryptionServicesPtrInput interface {
	pulumi.Input

	ToEncryptionServicesPtrOutput() EncryptionServicesPtrOutput
	ToEncryptionServicesPtrOutputWithContext(context.Context) EncryptionServicesPtrOutput
}

type encryptionServicesPtrType EncryptionServicesArgs

func EncryptionServicesPtr(v *EncryptionServicesArgs) EncryptionServicesPtrInput {
	return (*encryptionServicesPtrType)(v)
}

func (*encryptionServicesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionServices)(nil)).Elem()
}

func (i *encryptionServicesPtrType) ToEncryptionServicesPtrOutput() EncryptionServicesPtrOutput {
	return i.ToEncryptionServicesPtrOutputWithContext(context.Background())
}

func (i *encryptionServicesPtrType) ToEncryptionServicesPtrOutputWithContext(ctx context.Context) EncryptionServicesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionServicesPtrOutput)
}

type EncryptionServicesOutput struct{ *pulumi.OutputState }

func (EncryptionServicesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionServices)(nil)).Elem()
}

func (o EncryptionServicesOutput) ToEncryptionServicesOutput() EncryptionServicesOutput {
	return o
}

func (o EncryptionServicesOutput) ToEncryptionServicesOutputWithContext(ctx context.Context) EncryptionServicesOutput {
	return o
}

func (o EncryptionServicesOutput) ToEncryptionServicesPtrOutput() EncryptionServicesPtrOutput {
	return o.ToEncryptionServicesPtrOutputWithContext(context.Background())
}

func (o EncryptionServicesOutput) ToEncryptionServicesPtrOutputWithContext(ctx context.Context) EncryptionServicesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionServices) *EncryptionServices {
		return &v
	}).(EncryptionServicesPtrOutput)
}

func (o EncryptionServicesOutput) Blob() EncryptionServicePtrOutput {
	return o.ApplyT(func(v EncryptionServices) *EncryptionService { return v.Blob }).(EncryptionServicePtrOutput)
}

func (o EncryptionServicesOutput) File() EncryptionServicePtrOutput {
	return o.ApplyT(func(v EncryptionServices) *EncryptionService { return v.File }).(EncryptionServicePtrOutput)
}

type EncryptionServicesPtrOutput struct{ *pulumi.OutputState }

func (EncryptionServicesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionServices)(nil)).Elem()
}

func (o EncryptionServicesPtrOutput) ToEncryptionServicesPtrOutput() EncryptionServicesPtrOutput {
	return o
}

func (o EncryptionServicesPtrOutput) ToEncryptionServicesPtrOutputWithContext(ctx context.Context) EncryptionServicesPtrOutput {
	return o
}

func (o EncryptionServicesPtrOutput) Elem() EncryptionServicesOutput {
	return o.ApplyT(func(v *EncryptionServices) EncryptionServices {
		if v != nil {
			return *v
		}
		var ret EncryptionServices
		return ret
	}).(EncryptionServicesOutput)
}

func (o EncryptionServicesPtrOutput) Blob() EncryptionServicePtrOutput {
	return o.ApplyT(func(v *EncryptionServices) *EncryptionService {
		if v == nil {
			return nil
		}
		return v.Blob
	}).(EncryptionServicePtrOutput)
}

func (o EncryptionServicesPtrOutput) File() EncryptionServicePtrOutput {
	return o.ApplyT(func(v *EncryptionServices) *EncryptionService {
		if v == nil {
			return nil
		}
		return v.File
	}).(EncryptionServicePtrOutput)
}

type EncryptionServicesResponse struct {
	Blob  *EncryptionServiceResponse `pulumi:"blob"`
	File  *EncryptionServiceResponse `pulumi:"file"`
	Queue EncryptionServiceResponse  `pulumi:"queue"`
	Table EncryptionServiceResponse  `pulumi:"table"`
}

type EncryptionServicesResponseOutput struct{ *pulumi.OutputState }

func (EncryptionServicesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionServicesResponse)(nil)).Elem()
}

func (o EncryptionServicesResponseOutput) ToEncryptionServicesResponseOutput() EncryptionServicesResponseOutput {
	return o
}

func (o EncryptionServicesResponseOutput) ToEncryptionServicesResponseOutputWithContext(ctx context.Context) EncryptionServicesResponseOutput {
	return o
}

func (o EncryptionServicesResponseOutput) Blob() EncryptionServiceResponsePtrOutput {
	return o.ApplyT(func(v EncryptionServicesResponse) *EncryptionServiceResponse { return v.Blob }).(EncryptionServiceResponsePtrOutput)
}

func (o EncryptionServicesResponseOutput) File() EncryptionServiceResponsePtrOutput {
	return o.ApplyT(func(v EncryptionServicesResponse) *EncryptionServiceResponse { return v.File }).(EncryptionServiceResponsePtrOutput)
}

func (o EncryptionServicesResponseOutput) Queue() EncryptionServiceResponseOutput {
	return o.ApplyT(func(v EncryptionServicesResponse) EncryptionServiceResponse { return v.Queue }).(EncryptionServiceResponseOutput)
}

func (o EncryptionServicesResponseOutput) Table() EncryptionServiceResponseOutput {
	return o.ApplyT(func(v EncryptionServicesResponse) EncryptionServiceResponse { return v.Table }).(EncryptionServiceResponseOutput)
}

type EncryptionServicesResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionServicesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionServicesResponse)(nil)).Elem()
}

func (o EncryptionServicesResponsePtrOutput) ToEncryptionServicesResponsePtrOutput() EncryptionServicesResponsePtrOutput {
	return o
}

func (o EncryptionServicesResponsePtrOutput) ToEncryptionServicesResponsePtrOutputWithContext(ctx context.Context) EncryptionServicesResponsePtrOutput {
	return o
}

func (o EncryptionServicesResponsePtrOutput) Elem() EncryptionServicesResponseOutput {
	return o.ApplyT(func(v *EncryptionServicesResponse) EncryptionServicesResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionServicesResponse
		return ret
	}).(EncryptionServicesResponseOutput)
}

func (o EncryptionServicesResponsePtrOutput) Blob() EncryptionServiceResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionServicesResponse) *EncryptionServiceResponse {
		if v == nil {
			return nil
		}
		return v.Blob
	}).(EncryptionServiceResponsePtrOutput)
}

func (o EncryptionServicesResponsePtrOutput) File() EncryptionServiceResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionServicesResponse) *EncryptionServiceResponse {
		if v == nil {
			return nil
		}
		return v.File
	}).(EncryptionServiceResponsePtrOutput)
}

func (o EncryptionServicesResponsePtrOutput) Queue() EncryptionServiceResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionServicesResponse) *EncryptionServiceResponse {
		if v == nil {
			return nil
		}
		return &v.Queue
	}).(EncryptionServiceResponsePtrOutput)
}

func (o EncryptionServicesResponsePtrOutput) Table() EncryptionServiceResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionServicesResponse) *EncryptionServiceResponse {
		if v == nil {
			return nil
		}
		return &v.Table
	}).(EncryptionServiceResponsePtrOutput)
}

type EndpointsResponse struct {
	Blob  string `pulumi:"blob"`
	File  string `pulumi:"file"`
	Queue string `pulumi:"queue"`
	Table string `pulumi:"table"`
}

type EndpointsResponseOutput struct{ *pulumi.OutputState }

func (EndpointsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointsResponse)(nil)).Elem()
}

func (o EndpointsResponseOutput) ToEndpointsResponseOutput() EndpointsResponseOutput {
	return o
}

func (o EndpointsResponseOutput) ToEndpointsResponseOutputWithContext(ctx context.Context) EndpointsResponseOutput {
	return o
}

func (o EndpointsResponseOutput) Blob() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointsResponse) string { return v.Blob }).(pulumi.StringOutput)
}

func (o EndpointsResponseOutput) File() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointsResponse) string { return v.File }).(pulumi.StringOutput)
}

func (o EndpointsResponseOutput) Queue() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointsResponse) string { return v.Queue }).(pulumi.StringOutput)
}

func (o EndpointsResponseOutput) Table() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointsResponse) string { return v.Table }).(pulumi.StringOutput)
}

type Sku struct {
	Name SkuName `pulumi:"name"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Name SkuNameInput `pulumi:"name"`
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

func (o SkuOutput) Name() SkuNameOutput {
	return o.ApplyT(func(v Sku) SkuName { return v.Name }).(SkuNameOutput)
}

type SkuResponse struct {
	Name string `pulumi:"name"`
	Tier string `pulumi:"tier"`
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

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Tier }).(pulumi.StringOutput)
}

type StorageAccountKeyResponse struct {
	KeyName     string `pulumi:"keyName"`
	Permissions string `pulumi:"permissions"`
	Value       string `pulumi:"value"`
}

type StorageAccountKeyResponseOutput struct{ *pulumi.OutputState }

func (StorageAccountKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountKeyResponse)(nil)).Elem()
}

func (o StorageAccountKeyResponseOutput) ToStorageAccountKeyResponseOutput() StorageAccountKeyResponseOutput {
	return o
}

func (o StorageAccountKeyResponseOutput) ToStorageAccountKeyResponseOutputWithContext(ctx context.Context) StorageAccountKeyResponseOutput {
	return o
}

func (o StorageAccountKeyResponseOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountKeyResponse) string { return v.KeyName }).(pulumi.StringOutput)
}

func (o StorageAccountKeyResponseOutput) Permissions() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountKeyResponse) string { return v.Permissions }).(pulumi.StringOutput)
}

func (o StorageAccountKeyResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountKeyResponse) string { return v.Value }).(pulumi.StringOutput)
}

type StorageAccountKeyResponseArrayOutput struct{ *pulumi.OutputState }

func (StorageAccountKeyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StorageAccountKeyResponse)(nil)).Elem()
}

func (o StorageAccountKeyResponseArrayOutput) ToStorageAccountKeyResponseArrayOutput() StorageAccountKeyResponseArrayOutput {
	return o
}

func (o StorageAccountKeyResponseArrayOutput) ToStorageAccountKeyResponseArrayOutputWithContext(ctx context.Context) StorageAccountKeyResponseArrayOutput {
	return o
}

func (o StorageAccountKeyResponseArrayOutput) Index(i pulumi.IntInput) StorageAccountKeyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StorageAccountKeyResponse {
		return vs[0].([]StorageAccountKeyResponse)[vs[1].(int)]
	}).(StorageAccountKeyResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(CustomDomainOutput{})
	pulumi.RegisterOutputType(CustomDomainPtrOutput{})
	pulumi.RegisterOutputType(CustomDomainResponseOutput{})
	pulumi.RegisterOutputType(EncryptionOutput{})
	pulumi.RegisterOutputType(EncryptionPtrOutput{})
	pulumi.RegisterOutputType(EncryptionResponseOutput{})
	pulumi.RegisterOutputType(EncryptionServiceOutput{})
	pulumi.RegisterOutputType(EncryptionServicePtrOutput{})
	pulumi.RegisterOutputType(EncryptionServiceResponseOutput{})
	pulumi.RegisterOutputType(EncryptionServiceResponsePtrOutput{})
	pulumi.RegisterOutputType(EncryptionServicesOutput{})
	pulumi.RegisterOutputType(EncryptionServicesPtrOutput{})
	pulumi.RegisterOutputType(EncryptionServicesResponseOutput{})
	pulumi.RegisterOutputType(EncryptionServicesResponsePtrOutput{})
	pulumi.RegisterOutputType(EndpointsResponseOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(StorageAccountKeyResponseOutput{})
	pulumi.RegisterOutputType(StorageAccountKeyResponseArrayOutput{})
}
