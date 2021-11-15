


package v20201020

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





type MyWorkbookManagedIdentityResponseInput interface {
	pulumi.Input

	ToMyWorkbookManagedIdentityResponseOutput() MyWorkbookManagedIdentityResponseOutput
	ToMyWorkbookManagedIdentityResponseOutputWithContext(context.Context) MyWorkbookManagedIdentityResponseOutput
}

type MyWorkbookManagedIdentityResponseArgs struct {
	Type                   pulumi.StringPtrInput                            `pulumi:"type"`
	UserAssignedIdentities MyWorkbookUserAssignedIdentitiesResponsePtrInput `pulumi:"userAssignedIdentities"`
}

func (MyWorkbookManagedIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MyWorkbookManagedIdentityResponse)(nil)).Elem()
}

func (i MyWorkbookManagedIdentityResponseArgs) ToMyWorkbookManagedIdentityResponseOutput() MyWorkbookManagedIdentityResponseOutput {
	return i.ToMyWorkbookManagedIdentityResponseOutputWithContext(context.Background())
}

func (i MyWorkbookManagedIdentityResponseArgs) ToMyWorkbookManagedIdentityResponseOutputWithContext(ctx context.Context) MyWorkbookManagedIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MyWorkbookManagedIdentityResponseOutput)
}

func (i MyWorkbookManagedIdentityResponseArgs) ToMyWorkbookManagedIdentityResponsePtrOutput() MyWorkbookManagedIdentityResponsePtrOutput {
	return i.ToMyWorkbookManagedIdentityResponsePtrOutputWithContext(context.Background())
}

func (i MyWorkbookManagedIdentityResponseArgs) ToMyWorkbookManagedIdentityResponsePtrOutputWithContext(ctx context.Context) MyWorkbookManagedIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MyWorkbookManagedIdentityResponseOutput).ToMyWorkbookManagedIdentityResponsePtrOutputWithContext(ctx)
}









type MyWorkbookManagedIdentityResponsePtrInput interface {
	pulumi.Input

	ToMyWorkbookManagedIdentityResponsePtrOutput() MyWorkbookManagedIdentityResponsePtrOutput
	ToMyWorkbookManagedIdentityResponsePtrOutputWithContext(context.Context) MyWorkbookManagedIdentityResponsePtrOutput
}

type myWorkbookManagedIdentityResponsePtrType MyWorkbookManagedIdentityResponseArgs

func MyWorkbookManagedIdentityResponsePtr(v *MyWorkbookManagedIdentityResponseArgs) MyWorkbookManagedIdentityResponsePtrInput {
	return (*myWorkbookManagedIdentityResponsePtrType)(v)
}

func (*myWorkbookManagedIdentityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MyWorkbookManagedIdentityResponse)(nil)).Elem()
}

func (i *myWorkbookManagedIdentityResponsePtrType) ToMyWorkbookManagedIdentityResponsePtrOutput() MyWorkbookManagedIdentityResponsePtrOutput {
	return i.ToMyWorkbookManagedIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *myWorkbookManagedIdentityResponsePtrType) ToMyWorkbookManagedIdentityResponsePtrOutputWithContext(ctx context.Context) MyWorkbookManagedIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MyWorkbookManagedIdentityResponsePtrOutput)
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

func (o MyWorkbookManagedIdentityResponseOutput) ToMyWorkbookManagedIdentityResponsePtrOutput() MyWorkbookManagedIdentityResponsePtrOutput {
	return o.ToMyWorkbookManagedIdentityResponsePtrOutputWithContext(context.Background())
}

func (o MyWorkbookManagedIdentityResponseOutput) ToMyWorkbookManagedIdentityResponsePtrOutputWithContext(ctx context.Context) MyWorkbookManagedIdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MyWorkbookManagedIdentityResponse) *MyWorkbookManagedIdentityResponse {
		return &v
	}).(MyWorkbookManagedIdentityResponsePtrOutput)
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





type MyWorkbookUserAssignedIdentitiesResponseInput interface {
	pulumi.Input

	ToMyWorkbookUserAssignedIdentitiesResponseOutput() MyWorkbookUserAssignedIdentitiesResponseOutput
	ToMyWorkbookUserAssignedIdentitiesResponseOutputWithContext(context.Context) MyWorkbookUserAssignedIdentitiesResponseOutput
}

type MyWorkbookUserAssignedIdentitiesResponseArgs struct {
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
	TenantId    pulumi.StringInput `pulumi:"tenantId"`
}

func (MyWorkbookUserAssignedIdentitiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MyWorkbookUserAssignedIdentitiesResponse)(nil)).Elem()
}

func (i MyWorkbookUserAssignedIdentitiesResponseArgs) ToMyWorkbookUserAssignedIdentitiesResponseOutput() MyWorkbookUserAssignedIdentitiesResponseOutput {
	return i.ToMyWorkbookUserAssignedIdentitiesResponseOutputWithContext(context.Background())
}

func (i MyWorkbookUserAssignedIdentitiesResponseArgs) ToMyWorkbookUserAssignedIdentitiesResponseOutputWithContext(ctx context.Context) MyWorkbookUserAssignedIdentitiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MyWorkbookUserAssignedIdentitiesResponseOutput)
}

func (i MyWorkbookUserAssignedIdentitiesResponseArgs) ToMyWorkbookUserAssignedIdentitiesResponsePtrOutput() MyWorkbookUserAssignedIdentitiesResponsePtrOutput {
	return i.ToMyWorkbookUserAssignedIdentitiesResponsePtrOutputWithContext(context.Background())
}

func (i MyWorkbookUserAssignedIdentitiesResponseArgs) ToMyWorkbookUserAssignedIdentitiesResponsePtrOutputWithContext(ctx context.Context) MyWorkbookUserAssignedIdentitiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MyWorkbookUserAssignedIdentitiesResponseOutput).ToMyWorkbookUserAssignedIdentitiesResponsePtrOutputWithContext(ctx)
}









type MyWorkbookUserAssignedIdentitiesResponsePtrInput interface {
	pulumi.Input

	ToMyWorkbookUserAssignedIdentitiesResponsePtrOutput() MyWorkbookUserAssignedIdentitiesResponsePtrOutput
	ToMyWorkbookUserAssignedIdentitiesResponsePtrOutputWithContext(context.Context) MyWorkbookUserAssignedIdentitiesResponsePtrOutput
}

type myWorkbookUserAssignedIdentitiesResponsePtrType MyWorkbookUserAssignedIdentitiesResponseArgs

func MyWorkbookUserAssignedIdentitiesResponsePtr(v *MyWorkbookUserAssignedIdentitiesResponseArgs) MyWorkbookUserAssignedIdentitiesResponsePtrInput {
	return (*myWorkbookUserAssignedIdentitiesResponsePtrType)(v)
}

func (*myWorkbookUserAssignedIdentitiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MyWorkbookUserAssignedIdentitiesResponse)(nil)).Elem()
}

func (i *myWorkbookUserAssignedIdentitiesResponsePtrType) ToMyWorkbookUserAssignedIdentitiesResponsePtrOutput() MyWorkbookUserAssignedIdentitiesResponsePtrOutput {
	return i.ToMyWorkbookUserAssignedIdentitiesResponsePtrOutputWithContext(context.Background())
}

func (i *myWorkbookUserAssignedIdentitiesResponsePtrType) ToMyWorkbookUserAssignedIdentitiesResponsePtrOutputWithContext(ctx context.Context) MyWorkbookUserAssignedIdentitiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MyWorkbookUserAssignedIdentitiesResponsePtrOutput)
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

func (o MyWorkbookUserAssignedIdentitiesResponseOutput) ToMyWorkbookUserAssignedIdentitiesResponsePtrOutput() MyWorkbookUserAssignedIdentitiesResponsePtrOutput {
	return o.ToMyWorkbookUserAssignedIdentitiesResponsePtrOutputWithContext(context.Background())
}

func (o MyWorkbookUserAssignedIdentitiesResponseOutput) ToMyWorkbookUserAssignedIdentitiesResponsePtrOutputWithContext(ctx context.Context) MyWorkbookUserAssignedIdentitiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MyWorkbookUserAssignedIdentitiesResponse) *MyWorkbookUserAssignedIdentitiesResponse {
		return &v
	}).(MyWorkbookUserAssignedIdentitiesResponsePtrOutput)
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





type WorkbookManagedIdentityResponseInput interface {
	pulumi.Input

	ToWorkbookManagedIdentityResponseOutput() WorkbookManagedIdentityResponseOutput
	ToWorkbookManagedIdentityResponseOutputWithContext(context.Context) WorkbookManagedIdentityResponseOutput
}

type WorkbookManagedIdentityResponseArgs struct {
	Type                   pulumi.StringPtrInput                          `pulumi:"type"`
	UserAssignedIdentities WorkbookUserAssignedIdentitiesResponsePtrInput `pulumi:"userAssignedIdentities"`
}

func (WorkbookManagedIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookManagedIdentityResponse)(nil)).Elem()
}

func (i WorkbookManagedIdentityResponseArgs) ToWorkbookManagedIdentityResponseOutput() WorkbookManagedIdentityResponseOutput {
	return i.ToWorkbookManagedIdentityResponseOutputWithContext(context.Background())
}

func (i WorkbookManagedIdentityResponseArgs) ToWorkbookManagedIdentityResponseOutputWithContext(ctx context.Context) WorkbookManagedIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookManagedIdentityResponseOutput)
}

func (i WorkbookManagedIdentityResponseArgs) ToWorkbookManagedIdentityResponsePtrOutput() WorkbookManagedIdentityResponsePtrOutput {
	return i.ToWorkbookManagedIdentityResponsePtrOutputWithContext(context.Background())
}

func (i WorkbookManagedIdentityResponseArgs) ToWorkbookManagedIdentityResponsePtrOutputWithContext(ctx context.Context) WorkbookManagedIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookManagedIdentityResponseOutput).ToWorkbookManagedIdentityResponsePtrOutputWithContext(ctx)
}









type WorkbookManagedIdentityResponsePtrInput interface {
	pulumi.Input

	ToWorkbookManagedIdentityResponsePtrOutput() WorkbookManagedIdentityResponsePtrOutput
	ToWorkbookManagedIdentityResponsePtrOutputWithContext(context.Context) WorkbookManagedIdentityResponsePtrOutput
}

type workbookManagedIdentityResponsePtrType WorkbookManagedIdentityResponseArgs

func WorkbookManagedIdentityResponsePtr(v *WorkbookManagedIdentityResponseArgs) WorkbookManagedIdentityResponsePtrInput {
	return (*workbookManagedIdentityResponsePtrType)(v)
}

func (*workbookManagedIdentityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkbookManagedIdentityResponse)(nil)).Elem()
}

func (i *workbookManagedIdentityResponsePtrType) ToWorkbookManagedIdentityResponsePtrOutput() WorkbookManagedIdentityResponsePtrOutput {
	return i.ToWorkbookManagedIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *workbookManagedIdentityResponsePtrType) ToWorkbookManagedIdentityResponsePtrOutputWithContext(ctx context.Context) WorkbookManagedIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookManagedIdentityResponsePtrOutput)
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

func (o WorkbookManagedIdentityResponseOutput) ToWorkbookManagedIdentityResponsePtrOutput() WorkbookManagedIdentityResponsePtrOutput {
	return o.ToWorkbookManagedIdentityResponsePtrOutputWithContext(context.Background())
}

func (o WorkbookManagedIdentityResponseOutput) ToWorkbookManagedIdentityResponsePtrOutputWithContext(ctx context.Context) WorkbookManagedIdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkbookManagedIdentityResponse) *WorkbookManagedIdentityResponse {
		return &v
	}).(WorkbookManagedIdentityResponsePtrOutput)
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





type WorkbookUserAssignedIdentitiesResponseInput interface {
	pulumi.Input

	ToWorkbookUserAssignedIdentitiesResponseOutput() WorkbookUserAssignedIdentitiesResponseOutput
	ToWorkbookUserAssignedIdentitiesResponseOutputWithContext(context.Context) WorkbookUserAssignedIdentitiesResponseOutput
}

type WorkbookUserAssignedIdentitiesResponseArgs struct {
	ClientId    pulumi.StringInput `pulumi:"clientId"`
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
	TenantId    pulumi.StringInput `pulumi:"tenantId"`
}

func (WorkbookUserAssignedIdentitiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookUserAssignedIdentitiesResponse)(nil)).Elem()
}

func (i WorkbookUserAssignedIdentitiesResponseArgs) ToWorkbookUserAssignedIdentitiesResponseOutput() WorkbookUserAssignedIdentitiesResponseOutput {
	return i.ToWorkbookUserAssignedIdentitiesResponseOutputWithContext(context.Background())
}

func (i WorkbookUserAssignedIdentitiesResponseArgs) ToWorkbookUserAssignedIdentitiesResponseOutputWithContext(ctx context.Context) WorkbookUserAssignedIdentitiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookUserAssignedIdentitiesResponseOutput)
}

func (i WorkbookUserAssignedIdentitiesResponseArgs) ToWorkbookUserAssignedIdentitiesResponsePtrOutput() WorkbookUserAssignedIdentitiesResponsePtrOutput {
	return i.ToWorkbookUserAssignedIdentitiesResponsePtrOutputWithContext(context.Background())
}

func (i WorkbookUserAssignedIdentitiesResponseArgs) ToWorkbookUserAssignedIdentitiesResponsePtrOutputWithContext(ctx context.Context) WorkbookUserAssignedIdentitiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookUserAssignedIdentitiesResponseOutput).ToWorkbookUserAssignedIdentitiesResponsePtrOutputWithContext(ctx)
}









type WorkbookUserAssignedIdentitiesResponsePtrInput interface {
	pulumi.Input

	ToWorkbookUserAssignedIdentitiesResponsePtrOutput() WorkbookUserAssignedIdentitiesResponsePtrOutput
	ToWorkbookUserAssignedIdentitiesResponsePtrOutputWithContext(context.Context) WorkbookUserAssignedIdentitiesResponsePtrOutput
}

type workbookUserAssignedIdentitiesResponsePtrType WorkbookUserAssignedIdentitiesResponseArgs

func WorkbookUserAssignedIdentitiesResponsePtr(v *WorkbookUserAssignedIdentitiesResponseArgs) WorkbookUserAssignedIdentitiesResponsePtrInput {
	return (*workbookUserAssignedIdentitiesResponsePtrType)(v)
}

func (*workbookUserAssignedIdentitiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkbookUserAssignedIdentitiesResponse)(nil)).Elem()
}

func (i *workbookUserAssignedIdentitiesResponsePtrType) ToWorkbookUserAssignedIdentitiesResponsePtrOutput() WorkbookUserAssignedIdentitiesResponsePtrOutput {
	return i.ToWorkbookUserAssignedIdentitiesResponsePtrOutputWithContext(context.Background())
}

func (i *workbookUserAssignedIdentitiesResponsePtrType) ToWorkbookUserAssignedIdentitiesResponsePtrOutputWithContext(ctx context.Context) WorkbookUserAssignedIdentitiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookUserAssignedIdentitiesResponsePtrOutput)
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

func (o WorkbookUserAssignedIdentitiesResponseOutput) ToWorkbookUserAssignedIdentitiesResponsePtrOutput() WorkbookUserAssignedIdentitiesResponsePtrOutput {
	return o.ToWorkbookUserAssignedIdentitiesResponsePtrOutputWithContext(context.Background())
}

func (o WorkbookUserAssignedIdentitiesResponseOutput) ToWorkbookUserAssignedIdentitiesResponsePtrOutputWithContext(ctx context.Context) WorkbookUserAssignedIdentitiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkbookUserAssignedIdentitiesResponse) *WorkbookUserAssignedIdentitiesResponse {
		return &v
	}).(WorkbookUserAssignedIdentitiesResponsePtrOutput)
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
	pulumi.RegisterOutputType(WorkbookManagedIdentityOutput{})
	pulumi.RegisterOutputType(WorkbookManagedIdentityPtrOutput{})
	pulumi.RegisterOutputType(WorkbookManagedIdentityResponseOutput{})
	pulumi.RegisterOutputType(WorkbookManagedIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(WorkbookUserAssignedIdentitiesResponseOutput{})
	pulumi.RegisterOutputType(WorkbookUserAssignedIdentitiesResponsePtrOutput{})
}
