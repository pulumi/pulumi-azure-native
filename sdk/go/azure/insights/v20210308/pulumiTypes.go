


package v20210308

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MyWorkbookManagedIdentity struct {
	Type *string `pulumi:"type"`
}





type MyWorkbookManagedIdentityInput interface {
	pulumi.Input

	ToMyWorkbookManagedIdentityOutput() MyWorkbookManagedIdentityOutput
	ToMyWorkbookManagedIdentityOutputWithContext(context.Context) MyWorkbookManagedIdentityOutput
}

type MyWorkbookManagedIdentityArgs struct {
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (MyWorkbookManagedIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MyWorkbookManagedIdentity)(nil)).Elem()
}

func (i MyWorkbookManagedIdentityArgs) ToMyWorkbookManagedIdentityOutput() MyWorkbookManagedIdentityOutput {
	return i.ToMyWorkbookManagedIdentityOutputWithContext(context.Background())
}

func (i MyWorkbookManagedIdentityArgs) ToMyWorkbookManagedIdentityOutputWithContext(ctx context.Context) MyWorkbookManagedIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MyWorkbookManagedIdentityOutput)
}

func (i MyWorkbookManagedIdentityArgs) ToMyWorkbookManagedIdentityPtrOutput() MyWorkbookManagedIdentityPtrOutput {
	return i.ToMyWorkbookManagedIdentityPtrOutputWithContext(context.Background())
}

func (i MyWorkbookManagedIdentityArgs) ToMyWorkbookManagedIdentityPtrOutputWithContext(ctx context.Context) MyWorkbookManagedIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MyWorkbookManagedIdentityOutput).ToMyWorkbookManagedIdentityPtrOutputWithContext(ctx)
}









type MyWorkbookManagedIdentityPtrInput interface {
	pulumi.Input

	ToMyWorkbookManagedIdentityPtrOutput() MyWorkbookManagedIdentityPtrOutput
	ToMyWorkbookManagedIdentityPtrOutputWithContext(context.Context) MyWorkbookManagedIdentityPtrOutput
}

type myWorkbookManagedIdentityPtrType MyWorkbookManagedIdentityArgs

func MyWorkbookManagedIdentityPtr(v *MyWorkbookManagedIdentityArgs) MyWorkbookManagedIdentityPtrInput {
	return (*myWorkbookManagedIdentityPtrType)(v)
}

func (*myWorkbookManagedIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MyWorkbookManagedIdentity)(nil)).Elem()
}

func (i *myWorkbookManagedIdentityPtrType) ToMyWorkbookManagedIdentityPtrOutput() MyWorkbookManagedIdentityPtrOutput {
	return i.ToMyWorkbookManagedIdentityPtrOutputWithContext(context.Background())
}

func (i *myWorkbookManagedIdentityPtrType) ToMyWorkbookManagedIdentityPtrOutputWithContext(ctx context.Context) MyWorkbookManagedIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MyWorkbookManagedIdentityPtrOutput)
}

type MyWorkbookManagedIdentityOutput struct{ *pulumi.OutputState }

func (MyWorkbookManagedIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MyWorkbookManagedIdentity)(nil)).Elem()
}

func (o MyWorkbookManagedIdentityOutput) ToMyWorkbookManagedIdentityOutput() MyWorkbookManagedIdentityOutput {
	return o
}

func (o MyWorkbookManagedIdentityOutput) ToMyWorkbookManagedIdentityOutputWithContext(ctx context.Context) MyWorkbookManagedIdentityOutput {
	return o
}

func (o MyWorkbookManagedIdentityOutput) ToMyWorkbookManagedIdentityPtrOutput() MyWorkbookManagedIdentityPtrOutput {
	return o.ToMyWorkbookManagedIdentityPtrOutputWithContext(context.Background())
}

func (o MyWorkbookManagedIdentityOutput) ToMyWorkbookManagedIdentityPtrOutputWithContext(ctx context.Context) MyWorkbookManagedIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MyWorkbookManagedIdentity) *MyWorkbookManagedIdentity {
		return &v
	}).(MyWorkbookManagedIdentityPtrOutput)
}

func (o MyWorkbookManagedIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MyWorkbookManagedIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type MyWorkbookManagedIdentityPtrOutput struct{ *pulumi.OutputState }

func (MyWorkbookManagedIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MyWorkbookManagedIdentity)(nil)).Elem()
}

func (o MyWorkbookManagedIdentityPtrOutput) ToMyWorkbookManagedIdentityPtrOutput() MyWorkbookManagedIdentityPtrOutput {
	return o
}

func (o MyWorkbookManagedIdentityPtrOutput) ToMyWorkbookManagedIdentityPtrOutputWithContext(ctx context.Context) MyWorkbookManagedIdentityPtrOutput {
	return o
}

func (o MyWorkbookManagedIdentityPtrOutput) Elem() MyWorkbookManagedIdentityOutput {
	return o.ApplyT(func(v *MyWorkbookManagedIdentity) MyWorkbookManagedIdentity {
		if v != nil {
			return *v
		}
		var ret MyWorkbookManagedIdentity
		return ret
	}).(MyWorkbookManagedIdentityOutput)
}

func (o MyWorkbookManagedIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MyWorkbookManagedIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type MyWorkbookManagedIdentityResponse struct {
	Type                   *string                                   `pulumi:"type"`
	UserAssignedIdentities *MyWorkbookUserAssignedIdentitiesResponse `pulumi:"userAssignedIdentities"`
}

type MyWorkbookManagedIdentityResponseOutput struct{ *pulumi.OutputState }

func (MyWorkbookManagedIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MyWorkbookManagedIdentityResponse)(nil)).Elem()
}

func (o MyWorkbookManagedIdentityResponseOutput) ToMyWorkbookManagedIdentityResponseOutput() MyWorkbookManagedIdentityResponseOutput {
	return o
}

func (o MyWorkbookManagedIdentityResponseOutput) ToMyWorkbookManagedIdentityResponseOutputWithContext(ctx context.Context) MyWorkbookManagedIdentityResponseOutput {
	return o
}

func (o MyWorkbookManagedIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MyWorkbookManagedIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o MyWorkbookManagedIdentityResponseOutput) UserAssignedIdentities() MyWorkbookUserAssignedIdentitiesResponsePtrOutput {
	return o.ApplyT(func(v MyWorkbookManagedIdentityResponse) *MyWorkbookUserAssignedIdentitiesResponse {
		return v.UserAssignedIdentities
	}).(MyWorkbookUserAssignedIdentitiesResponsePtrOutput)
}

type MyWorkbookManagedIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (MyWorkbookManagedIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MyWorkbookManagedIdentityResponse)(nil)).Elem()
}

func (o MyWorkbookManagedIdentityResponsePtrOutput) ToMyWorkbookManagedIdentityResponsePtrOutput() MyWorkbookManagedIdentityResponsePtrOutput {
	return o
}

func (o MyWorkbookManagedIdentityResponsePtrOutput) ToMyWorkbookManagedIdentityResponsePtrOutputWithContext(ctx context.Context) MyWorkbookManagedIdentityResponsePtrOutput {
	return o
}

func (o MyWorkbookManagedIdentityResponsePtrOutput) Elem() MyWorkbookManagedIdentityResponseOutput {
	return o.ApplyT(func(v *MyWorkbookManagedIdentityResponse) MyWorkbookManagedIdentityResponse {
		if v != nil {
			return *v
		}
		var ret MyWorkbookManagedIdentityResponse
		return ret
	}).(MyWorkbookManagedIdentityResponseOutput)
}

func (o MyWorkbookManagedIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MyWorkbookManagedIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o MyWorkbookManagedIdentityResponsePtrOutput) UserAssignedIdentities() MyWorkbookUserAssignedIdentitiesResponsePtrOutput {
	return o.ApplyT(func(v *MyWorkbookManagedIdentityResponse) *MyWorkbookUserAssignedIdentitiesResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(MyWorkbookUserAssignedIdentitiesResponsePtrOutput)
}

type MyWorkbookUserAssignedIdentitiesResponse struct {
	PrincipalId string `pulumi:"principalId"`
	TenantId    string `pulumi:"tenantId"`
}

type MyWorkbookUserAssignedIdentitiesResponseOutput struct{ *pulumi.OutputState }

func (MyWorkbookUserAssignedIdentitiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MyWorkbookUserAssignedIdentitiesResponse)(nil)).Elem()
}

func (o MyWorkbookUserAssignedIdentitiesResponseOutput) ToMyWorkbookUserAssignedIdentitiesResponseOutput() MyWorkbookUserAssignedIdentitiesResponseOutput {
	return o
}

func (o MyWorkbookUserAssignedIdentitiesResponseOutput) ToMyWorkbookUserAssignedIdentitiesResponseOutputWithContext(ctx context.Context) MyWorkbookUserAssignedIdentitiesResponseOutput {
	return o
}

func (o MyWorkbookUserAssignedIdentitiesResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v MyWorkbookUserAssignedIdentitiesResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o MyWorkbookUserAssignedIdentitiesResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v MyWorkbookUserAssignedIdentitiesResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

type MyWorkbookUserAssignedIdentitiesResponsePtrOutput struct{ *pulumi.OutputState }

func (MyWorkbookUserAssignedIdentitiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MyWorkbookUserAssignedIdentitiesResponse)(nil)).Elem()
}

func (o MyWorkbookUserAssignedIdentitiesResponsePtrOutput) ToMyWorkbookUserAssignedIdentitiesResponsePtrOutput() MyWorkbookUserAssignedIdentitiesResponsePtrOutput {
	return o
}

func (o MyWorkbookUserAssignedIdentitiesResponsePtrOutput) ToMyWorkbookUserAssignedIdentitiesResponsePtrOutputWithContext(ctx context.Context) MyWorkbookUserAssignedIdentitiesResponsePtrOutput {
	return o
}

func (o MyWorkbookUserAssignedIdentitiesResponsePtrOutput) Elem() MyWorkbookUserAssignedIdentitiesResponseOutput {
	return o.ApplyT(func(v *MyWorkbookUserAssignedIdentitiesResponse) MyWorkbookUserAssignedIdentitiesResponse {
		if v != nil {
			return *v
		}
		var ret MyWorkbookUserAssignedIdentitiesResponse
		return ret
	}).(MyWorkbookUserAssignedIdentitiesResponseOutput)
}

func (o MyWorkbookUserAssignedIdentitiesResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MyWorkbookUserAssignedIdentitiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o MyWorkbookUserAssignedIdentitiesResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MyWorkbookUserAssignedIdentitiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
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

type WorkbookManagedIdentity struct {
	Type *string `pulumi:"type"`
}





type WorkbookManagedIdentityInput interface {
	pulumi.Input

	ToWorkbookManagedIdentityOutput() WorkbookManagedIdentityOutput
	ToWorkbookManagedIdentityOutputWithContext(context.Context) WorkbookManagedIdentityOutput
}

type WorkbookManagedIdentityArgs struct {
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (WorkbookManagedIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookManagedIdentity)(nil)).Elem()
}

func (i WorkbookManagedIdentityArgs) ToWorkbookManagedIdentityOutput() WorkbookManagedIdentityOutput {
	return i.ToWorkbookManagedIdentityOutputWithContext(context.Background())
}

func (i WorkbookManagedIdentityArgs) ToWorkbookManagedIdentityOutputWithContext(ctx context.Context) WorkbookManagedIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookManagedIdentityOutput)
}

func (i WorkbookManagedIdentityArgs) ToWorkbookManagedIdentityPtrOutput() WorkbookManagedIdentityPtrOutput {
	return i.ToWorkbookManagedIdentityPtrOutputWithContext(context.Background())
}

func (i WorkbookManagedIdentityArgs) ToWorkbookManagedIdentityPtrOutputWithContext(ctx context.Context) WorkbookManagedIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookManagedIdentityOutput).ToWorkbookManagedIdentityPtrOutputWithContext(ctx)
}









type WorkbookManagedIdentityPtrInput interface {
	pulumi.Input

	ToWorkbookManagedIdentityPtrOutput() WorkbookManagedIdentityPtrOutput
	ToWorkbookManagedIdentityPtrOutputWithContext(context.Context) WorkbookManagedIdentityPtrOutput
}

type workbookManagedIdentityPtrType WorkbookManagedIdentityArgs

func WorkbookManagedIdentityPtr(v *WorkbookManagedIdentityArgs) WorkbookManagedIdentityPtrInput {
	return (*workbookManagedIdentityPtrType)(v)
}

func (*workbookManagedIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkbookManagedIdentity)(nil)).Elem()
}

func (i *workbookManagedIdentityPtrType) ToWorkbookManagedIdentityPtrOutput() WorkbookManagedIdentityPtrOutput {
	return i.ToWorkbookManagedIdentityPtrOutputWithContext(context.Background())
}

func (i *workbookManagedIdentityPtrType) ToWorkbookManagedIdentityPtrOutputWithContext(ctx context.Context) WorkbookManagedIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookManagedIdentityPtrOutput)
}

type WorkbookManagedIdentityOutput struct{ *pulumi.OutputState }

func (WorkbookManagedIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookManagedIdentity)(nil)).Elem()
}

func (o WorkbookManagedIdentityOutput) ToWorkbookManagedIdentityOutput() WorkbookManagedIdentityOutput {
	return o
}

func (o WorkbookManagedIdentityOutput) ToWorkbookManagedIdentityOutputWithContext(ctx context.Context) WorkbookManagedIdentityOutput {
	return o
}

func (o WorkbookManagedIdentityOutput) ToWorkbookManagedIdentityPtrOutput() WorkbookManagedIdentityPtrOutput {
	return o.ToWorkbookManagedIdentityPtrOutputWithContext(context.Background())
}

func (o WorkbookManagedIdentityOutput) ToWorkbookManagedIdentityPtrOutputWithContext(ctx context.Context) WorkbookManagedIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkbookManagedIdentity) *WorkbookManagedIdentity {
		return &v
	}).(WorkbookManagedIdentityPtrOutput)
}

func (o WorkbookManagedIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkbookManagedIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type WorkbookManagedIdentityPtrOutput struct{ *pulumi.OutputState }

func (WorkbookManagedIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkbookManagedIdentity)(nil)).Elem()
}

func (o WorkbookManagedIdentityPtrOutput) ToWorkbookManagedIdentityPtrOutput() WorkbookManagedIdentityPtrOutput {
	return o
}

func (o WorkbookManagedIdentityPtrOutput) ToWorkbookManagedIdentityPtrOutputWithContext(ctx context.Context) WorkbookManagedIdentityPtrOutput {
	return o
}

func (o WorkbookManagedIdentityPtrOutput) Elem() WorkbookManagedIdentityOutput {
	return o.ApplyT(func(v *WorkbookManagedIdentity) WorkbookManagedIdentity {
		if v != nil {
			return *v
		}
		var ret WorkbookManagedIdentity
		return ret
	}).(WorkbookManagedIdentityOutput)
}

func (o WorkbookManagedIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkbookManagedIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type WorkbookManagedIdentityResponse struct {
	Type                   *string                                 `pulumi:"type"`
	UserAssignedIdentities *WorkbookUserAssignedIdentitiesResponse `pulumi:"userAssignedIdentities"`
}

type WorkbookManagedIdentityResponseOutput struct{ *pulumi.OutputState }

func (WorkbookManagedIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookManagedIdentityResponse)(nil)).Elem()
}

func (o WorkbookManagedIdentityResponseOutput) ToWorkbookManagedIdentityResponseOutput() WorkbookManagedIdentityResponseOutput {
	return o
}

func (o WorkbookManagedIdentityResponseOutput) ToWorkbookManagedIdentityResponseOutputWithContext(ctx context.Context) WorkbookManagedIdentityResponseOutput {
	return o
}

func (o WorkbookManagedIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkbookManagedIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o WorkbookManagedIdentityResponseOutput) UserAssignedIdentities() WorkbookUserAssignedIdentitiesResponsePtrOutput {
	return o.ApplyT(func(v WorkbookManagedIdentityResponse) *WorkbookUserAssignedIdentitiesResponse {
		return v.UserAssignedIdentities
	}).(WorkbookUserAssignedIdentitiesResponsePtrOutput)
}

type WorkbookManagedIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (WorkbookManagedIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkbookManagedIdentityResponse)(nil)).Elem()
}

func (o WorkbookManagedIdentityResponsePtrOutput) ToWorkbookManagedIdentityResponsePtrOutput() WorkbookManagedIdentityResponsePtrOutput {
	return o
}

func (o WorkbookManagedIdentityResponsePtrOutput) ToWorkbookManagedIdentityResponsePtrOutputWithContext(ctx context.Context) WorkbookManagedIdentityResponsePtrOutput {
	return o
}

func (o WorkbookManagedIdentityResponsePtrOutput) Elem() WorkbookManagedIdentityResponseOutput {
	return o.ApplyT(func(v *WorkbookManagedIdentityResponse) WorkbookManagedIdentityResponse {
		if v != nil {
			return *v
		}
		var ret WorkbookManagedIdentityResponse
		return ret
	}).(WorkbookManagedIdentityResponseOutput)
}

func (o WorkbookManagedIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkbookManagedIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o WorkbookManagedIdentityResponsePtrOutput) UserAssignedIdentities() WorkbookUserAssignedIdentitiesResponsePtrOutput {
	return o.ApplyT(func(v *WorkbookManagedIdentityResponse) *WorkbookUserAssignedIdentitiesResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(WorkbookUserAssignedIdentitiesResponsePtrOutput)
}

type WorkbookUserAssignedIdentitiesResponse struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
	TenantId    string `pulumi:"tenantId"`
}

type WorkbookUserAssignedIdentitiesResponseOutput struct{ *pulumi.OutputState }

func (WorkbookUserAssignedIdentitiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookUserAssignedIdentitiesResponse)(nil)).Elem()
}

func (o WorkbookUserAssignedIdentitiesResponseOutput) ToWorkbookUserAssignedIdentitiesResponseOutput() WorkbookUserAssignedIdentitiesResponseOutput {
	return o
}

func (o WorkbookUserAssignedIdentitiesResponseOutput) ToWorkbookUserAssignedIdentitiesResponseOutputWithContext(ctx context.Context) WorkbookUserAssignedIdentitiesResponseOutput {
	return o
}

func (o WorkbookUserAssignedIdentitiesResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v WorkbookUserAssignedIdentitiesResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o WorkbookUserAssignedIdentitiesResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v WorkbookUserAssignedIdentitiesResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o WorkbookUserAssignedIdentitiesResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v WorkbookUserAssignedIdentitiesResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

type WorkbookUserAssignedIdentitiesResponsePtrOutput struct{ *pulumi.OutputState }

func (WorkbookUserAssignedIdentitiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkbookUserAssignedIdentitiesResponse)(nil)).Elem()
}

func (o WorkbookUserAssignedIdentitiesResponsePtrOutput) ToWorkbookUserAssignedIdentitiesResponsePtrOutput() WorkbookUserAssignedIdentitiesResponsePtrOutput {
	return o
}

func (o WorkbookUserAssignedIdentitiesResponsePtrOutput) ToWorkbookUserAssignedIdentitiesResponsePtrOutputWithContext(ctx context.Context) WorkbookUserAssignedIdentitiesResponsePtrOutput {
	return o
}

func (o WorkbookUserAssignedIdentitiesResponsePtrOutput) Elem() WorkbookUserAssignedIdentitiesResponseOutput {
	return o.ApplyT(func(v *WorkbookUserAssignedIdentitiesResponse) WorkbookUserAssignedIdentitiesResponse {
		if v != nil {
			return *v
		}
		var ret WorkbookUserAssignedIdentitiesResponse
		return ret
	}).(WorkbookUserAssignedIdentitiesResponseOutput)
}

func (o WorkbookUserAssignedIdentitiesResponsePtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkbookUserAssignedIdentitiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o WorkbookUserAssignedIdentitiesResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkbookUserAssignedIdentitiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o WorkbookUserAssignedIdentitiesResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkbookUserAssignedIdentitiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(MyWorkbookManagedIdentityOutput{})
	pulumi.RegisterOutputType(MyWorkbookManagedIdentityPtrOutput{})
	pulumi.RegisterOutputType(MyWorkbookManagedIdentityResponseOutput{})
	pulumi.RegisterOutputType(MyWorkbookManagedIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(MyWorkbookUserAssignedIdentitiesResponseOutput{})
	pulumi.RegisterOutputType(MyWorkbookUserAssignedIdentitiesResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(WorkbookManagedIdentityOutput{})
	pulumi.RegisterOutputType(WorkbookManagedIdentityPtrOutput{})
	pulumi.RegisterOutputType(WorkbookManagedIdentityResponseOutput{})
	pulumi.RegisterOutputType(WorkbookManagedIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(WorkbookUserAssignedIdentitiesResponseOutput{})
	pulumi.RegisterOutputType(WorkbookUserAssignedIdentitiesResponsePtrOutput{})
}
