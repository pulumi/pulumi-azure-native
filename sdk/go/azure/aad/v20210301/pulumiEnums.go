


package v20210301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ExternalAccess string

const (
	ExternalAccessEnabled  = ExternalAccess("Enabled")
	ExternalAccessDisabled = ExternalAccess("Disabled")
)

func (ExternalAccess) ElementType() reflect.Type {
	return reflect.TypeOf((*ExternalAccess)(nil)).Elem()
}

func (e ExternalAccess) ToExternalAccessOutput() ExternalAccessOutput {
	return pulumi.ToOutput(e).(ExternalAccessOutput)
}

func (e ExternalAccess) ToExternalAccessOutputWithContext(ctx context.Context) ExternalAccessOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ExternalAccessOutput)
}

func (e ExternalAccess) ToExternalAccessPtrOutput() ExternalAccessPtrOutput {
	return e.ToExternalAccessPtrOutputWithContext(context.Background())
}

func (e ExternalAccess) ToExternalAccessPtrOutputWithContext(ctx context.Context) ExternalAccessPtrOutput {
	return ExternalAccess(e).ToExternalAccessOutputWithContext(ctx).ToExternalAccessPtrOutputWithContext(ctx)
}

func (e ExternalAccess) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExternalAccess) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExternalAccess) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ExternalAccess) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ExternalAccessOutput struct{ *pulumi.OutputState }

func (ExternalAccessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExternalAccess)(nil)).Elem()
}

func (o ExternalAccessOutput) ToExternalAccessOutput() ExternalAccessOutput {
	return o
}

func (o ExternalAccessOutput) ToExternalAccessOutputWithContext(ctx context.Context) ExternalAccessOutput {
	return o
}

func (o ExternalAccessOutput) ToExternalAccessPtrOutput() ExternalAccessPtrOutput {
	return o.ToExternalAccessPtrOutputWithContext(context.Background())
}

func (o ExternalAccessOutput) ToExternalAccessPtrOutputWithContext(ctx context.Context) ExternalAccessPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExternalAccess) *ExternalAccess {
		return &v
	}).(ExternalAccessPtrOutput)
}

func (o ExternalAccessOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ExternalAccessOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExternalAccess) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ExternalAccessOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExternalAccessOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExternalAccess) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ExternalAccessPtrOutput struct{ *pulumi.OutputState }

func (ExternalAccessPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExternalAccess)(nil)).Elem()
}

func (o ExternalAccessPtrOutput) ToExternalAccessPtrOutput() ExternalAccessPtrOutput {
	return o
}

func (o ExternalAccessPtrOutput) ToExternalAccessPtrOutputWithContext(ctx context.Context) ExternalAccessPtrOutput {
	return o
}

func (o ExternalAccessPtrOutput) Elem() ExternalAccessOutput {
	return o.ApplyT(func(v *ExternalAccess) ExternalAccess {
		if v != nil {
			return *v
		}
		var ret ExternalAccess
		return ret
	}).(ExternalAccessOutput)
}

func (o ExternalAccessPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExternalAccessPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ExternalAccess) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ExternalAccessInput interface {
	pulumi.Input

	ToExternalAccessOutput() ExternalAccessOutput
	ToExternalAccessOutputWithContext(context.Context) ExternalAccessOutput
}

var externalAccessPtrType = reflect.TypeOf((**ExternalAccess)(nil)).Elem()

type ExternalAccessPtrInput interface {
	pulumi.Input

	ToExternalAccessPtrOutput() ExternalAccessPtrOutput
	ToExternalAccessPtrOutputWithContext(context.Context) ExternalAccessPtrOutput
}

type externalAccessPtr string

func ExternalAccessPtr(v string) ExternalAccessPtrInput {
	return (*externalAccessPtr)(&v)
}

func (*externalAccessPtr) ElementType() reflect.Type {
	return externalAccessPtrType
}

func (in *externalAccessPtr) ToExternalAccessPtrOutput() ExternalAccessPtrOutput {
	return pulumi.ToOutput(in).(ExternalAccessPtrOutput)
}

func (in *externalAccessPtr) ToExternalAccessPtrOutputWithContext(ctx context.Context) ExternalAccessPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ExternalAccessPtrOutput)
}

type FilteredSync string

const (
	FilteredSyncEnabled  = FilteredSync("Enabled")
	FilteredSyncDisabled = FilteredSync("Disabled")
)

func (FilteredSync) ElementType() reflect.Type {
	return reflect.TypeOf((*FilteredSync)(nil)).Elem()
}

func (e FilteredSync) ToFilteredSyncOutput() FilteredSyncOutput {
	return pulumi.ToOutput(e).(FilteredSyncOutput)
}

func (e FilteredSync) ToFilteredSyncOutputWithContext(ctx context.Context) FilteredSyncOutput {
	return pulumi.ToOutputWithContext(ctx, e).(FilteredSyncOutput)
}

func (e FilteredSync) ToFilteredSyncPtrOutput() FilteredSyncPtrOutput {
	return e.ToFilteredSyncPtrOutputWithContext(context.Background())
}

func (e FilteredSync) ToFilteredSyncPtrOutputWithContext(ctx context.Context) FilteredSyncPtrOutput {
	return FilteredSync(e).ToFilteredSyncOutputWithContext(ctx).ToFilteredSyncPtrOutputWithContext(ctx)
}

func (e FilteredSync) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e FilteredSync) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e FilteredSync) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e FilteredSync) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type FilteredSyncOutput struct{ *pulumi.OutputState }

func (FilteredSyncOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FilteredSync)(nil)).Elem()
}

func (o FilteredSyncOutput) ToFilteredSyncOutput() FilteredSyncOutput {
	return o
}

func (o FilteredSyncOutput) ToFilteredSyncOutputWithContext(ctx context.Context) FilteredSyncOutput {
	return o
}

func (o FilteredSyncOutput) ToFilteredSyncPtrOutput() FilteredSyncPtrOutput {
	return o.ToFilteredSyncPtrOutputWithContext(context.Background())
}

func (o FilteredSyncOutput) ToFilteredSyncPtrOutputWithContext(ctx context.Context) FilteredSyncPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FilteredSync) *FilteredSync {
		return &v
	}).(FilteredSyncPtrOutput)
}

func (o FilteredSyncOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o FilteredSyncOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FilteredSync) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o FilteredSyncOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FilteredSyncOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FilteredSync) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type FilteredSyncPtrOutput struct{ *pulumi.OutputState }

func (FilteredSyncPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FilteredSync)(nil)).Elem()
}

func (o FilteredSyncPtrOutput) ToFilteredSyncPtrOutput() FilteredSyncPtrOutput {
	return o
}

func (o FilteredSyncPtrOutput) ToFilteredSyncPtrOutputWithContext(ctx context.Context) FilteredSyncPtrOutput {
	return o
}

func (o FilteredSyncPtrOutput) Elem() FilteredSyncOutput {
	return o.ApplyT(func(v *FilteredSync) FilteredSync {
		if v != nil {
			return *v
		}
		var ret FilteredSync
		return ret
	}).(FilteredSyncOutput)
}

func (o FilteredSyncPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FilteredSyncPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *FilteredSync) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type FilteredSyncInput interface {
	pulumi.Input

	ToFilteredSyncOutput() FilteredSyncOutput
	ToFilteredSyncOutputWithContext(context.Context) FilteredSyncOutput
}

var filteredSyncPtrType = reflect.TypeOf((**FilteredSync)(nil)).Elem()

type FilteredSyncPtrInput interface {
	pulumi.Input

	ToFilteredSyncPtrOutput() FilteredSyncPtrOutput
	ToFilteredSyncPtrOutputWithContext(context.Context) FilteredSyncPtrOutput
}

type filteredSyncPtr string

func FilteredSyncPtr(v string) FilteredSyncPtrInput {
	return (*filteredSyncPtr)(&v)
}

func (*filteredSyncPtr) ElementType() reflect.Type {
	return filteredSyncPtrType
}

func (in *filteredSyncPtr) ToFilteredSyncPtrOutput() FilteredSyncPtrOutput {
	return pulumi.ToOutput(in).(FilteredSyncPtrOutput)
}

func (in *filteredSyncPtr) ToFilteredSyncPtrOutputWithContext(ctx context.Context) FilteredSyncPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(FilteredSyncPtrOutput)
}

type KerberosArmoring string

const (
	KerberosArmoringEnabled  = KerberosArmoring("Enabled")
	KerberosArmoringDisabled = KerberosArmoring("Disabled")
)

func (KerberosArmoring) ElementType() reflect.Type {
	return reflect.TypeOf((*KerberosArmoring)(nil)).Elem()
}

func (e KerberosArmoring) ToKerberosArmoringOutput() KerberosArmoringOutput {
	return pulumi.ToOutput(e).(KerberosArmoringOutput)
}

func (e KerberosArmoring) ToKerberosArmoringOutputWithContext(ctx context.Context) KerberosArmoringOutput {
	return pulumi.ToOutputWithContext(ctx, e).(KerberosArmoringOutput)
}

func (e KerberosArmoring) ToKerberosArmoringPtrOutput() KerberosArmoringPtrOutput {
	return e.ToKerberosArmoringPtrOutputWithContext(context.Background())
}

func (e KerberosArmoring) ToKerberosArmoringPtrOutputWithContext(ctx context.Context) KerberosArmoringPtrOutput {
	return KerberosArmoring(e).ToKerberosArmoringOutputWithContext(ctx).ToKerberosArmoringPtrOutputWithContext(ctx)
}

func (e KerberosArmoring) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e KerberosArmoring) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e KerberosArmoring) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e KerberosArmoring) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type KerberosArmoringOutput struct{ *pulumi.OutputState }

func (KerberosArmoringOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KerberosArmoring)(nil)).Elem()
}

func (o KerberosArmoringOutput) ToKerberosArmoringOutput() KerberosArmoringOutput {
	return o
}

func (o KerberosArmoringOutput) ToKerberosArmoringOutputWithContext(ctx context.Context) KerberosArmoringOutput {
	return o
}

func (o KerberosArmoringOutput) ToKerberosArmoringPtrOutput() KerberosArmoringPtrOutput {
	return o.ToKerberosArmoringPtrOutputWithContext(context.Background())
}

func (o KerberosArmoringOutput) ToKerberosArmoringPtrOutputWithContext(ctx context.Context) KerberosArmoringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KerberosArmoring) *KerberosArmoring {
		return &v
	}).(KerberosArmoringPtrOutput)
}

func (o KerberosArmoringOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o KerberosArmoringOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e KerberosArmoring) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o KerberosArmoringOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KerberosArmoringOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e KerberosArmoring) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type KerberosArmoringPtrOutput struct{ *pulumi.OutputState }

func (KerberosArmoringPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KerberosArmoring)(nil)).Elem()
}

func (o KerberosArmoringPtrOutput) ToKerberosArmoringPtrOutput() KerberosArmoringPtrOutput {
	return o
}

func (o KerberosArmoringPtrOutput) ToKerberosArmoringPtrOutputWithContext(ctx context.Context) KerberosArmoringPtrOutput {
	return o
}

func (o KerberosArmoringPtrOutput) Elem() KerberosArmoringOutput {
	return o.ApplyT(func(v *KerberosArmoring) KerberosArmoring {
		if v != nil {
			return *v
		}
		var ret KerberosArmoring
		return ret
	}).(KerberosArmoringOutput)
}

func (o KerberosArmoringPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KerberosArmoringPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *KerberosArmoring) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type KerberosArmoringInput interface {
	pulumi.Input

	ToKerberosArmoringOutput() KerberosArmoringOutput
	ToKerberosArmoringOutputWithContext(context.Context) KerberosArmoringOutput
}

var kerberosArmoringPtrType = reflect.TypeOf((**KerberosArmoring)(nil)).Elem()

type KerberosArmoringPtrInput interface {
	pulumi.Input

	ToKerberosArmoringPtrOutput() KerberosArmoringPtrOutput
	ToKerberosArmoringPtrOutputWithContext(context.Context) KerberosArmoringPtrOutput
}

type kerberosArmoringPtr string

func KerberosArmoringPtr(v string) KerberosArmoringPtrInput {
	return (*kerberosArmoringPtr)(&v)
}

func (*kerberosArmoringPtr) ElementType() reflect.Type {
	return kerberosArmoringPtrType
}

func (in *kerberosArmoringPtr) ToKerberosArmoringPtrOutput() KerberosArmoringPtrOutput {
	return pulumi.ToOutput(in).(KerberosArmoringPtrOutput)
}

func (in *kerberosArmoringPtr) ToKerberosArmoringPtrOutputWithContext(ctx context.Context) KerberosArmoringPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(KerberosArmoringPtrOutput)
}

type KerberosRc4Encryption string

const (
	KerberosRc4EncryptionEnabled  = KerberosRc4Encryption("Enabled")
	KerberosRc4EncryptionDisabled = KerberosRc4Encryption("Disabled")
)

func (KerberosRc4Encryption) ElementType() reflect.Type {
	return reflect.TypeOf((*KerberosRc4Encryption)(nil)).Elem()
}

func (e KerberosRc4Encryption) ToKerberosRc4EncryptionOutput() KerberosRc4EncryptionOutput {
	return pulumi.ToOutput(e).(KerberosRc4EncryptionOutput)
}

func (e KerberosRc4Encryption) ToKerberosRc4EncryptionOutputWithContext(ctx context.Context) KerberosRc4EncryptionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(KerberosRc4EncryptionOutput)
}

func (e KerberosRc4Encryption) ToKerberosRc4EncryptionPtrOutput() KerberosRc4EncryptionPtrOutput {
	return e.ToKerberosRc4EncryptionPtrOutputWithContext(context.Background())
}

func (e KerberosRc4Encryption) ToKerberosRc4EncryptionPtrOutputWithContext(ctx context.Context) KerberosRc4EncryptionPtrOutput {
	return KerberosRc4Encryption(e).ToKerberosRc4EncryptionOutputWithContext(ctx).ToKerberosRc4EncryptionPtrOutputWithContext(ctx)
}

func (e KerberosRc4Encryption) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e KerberosRc4Encryption) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e KerberosRc4Encryption) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e KerberosRc4Encryption) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type KerberosRc4EncryptionOutput struct{ *pulumi.OutputState }

func (KerberosRc4EncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KerberosRc4Encryption)(nil)).Elem()
}

func (o KerberosRc4EncryptionOutput) ToKerberosRc4EncryptionOutput() KerberosRc4EncryptionOutput {
	return o
}

func (o KerberosRc4EncryptionOutput) ToKerberosRc4EncryptionOutputWithContext(ctx context.Context) KerberosRc4EncryptionOutput {
	return o
}

func (o KerberosRc4EncryptionOutput) ToKerberosRc4EncryptionPtrOutput() KerberosRc4EncryptionPtrOutput {
	return o.ToKerberosRc4EncryptionPtrOutputWithContext(context.Background())
}

func (o KerberosRc4EncryptionOutput) ToKerberosRc4EncryptionPtrOutputWithContext(ctx context.Context) KerberosRc4EncryptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KerberosRc4Encryption) *KerberosRc4Encryption {
		return &v
	}).(KerberosRc4EncryptionPtrOutput)
}

func (o KerberosRc4EncryptionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o KerberosRc4EncryptionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e KerberosRc4Encryption) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o KerberosRc4EncryptionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KerberosRc4EncryptionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e KerberosRc4Encryption) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type KerberosRc4EncryptionPtrOutput struct{ *pulumi.OutputState }

func (KerberosRc4EncryptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KerberosRc4Encryption)(nil)).Elem()
}

func (o KerberosRc4EncryptionPtrOutput) ToKerberosRc4EncryptionPtrOutput() KerberosRc4EncryptionPtrOutput {
	return o
}

func (o KerberosRc4EncryptionPtrOutput) ToKerberosRc4EncryptionPtrOutputWithContext(ctx context.Context) KerberosRc4EncryptionPtrOutput {
	return o
}

func (o KerberosRc4EncryptionPtrOutput) Elem() KerberosRc4EncryptionOutput {
	return o.ApplyT(func(v *KerberosRc4Encryption) KerberosRc4Encryption {
		if v != nil {
			return *v
		}
		var ret KerberosRc4Encryption
		return ret
	}).(KerberosRc4EncryptionOutput)
}

func (o KerberosRc4EncryptionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KerberosRc4EncryptionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *KerberosRc4Encryption) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type KerberosRc4EncryptionInput interface {
	pulumi.Input

	ToKerberosRc4EncryptionOutput() KerberosRc4EncryptionOutput
	ToKerberosRc4EncryptionOutputWithContext(context.Context) KerberosRc4EncryptionOutput
}

var kerberosRc4EncryptionPtrType = reflect.TypeOf((**KerberosRc4Encryption)(nil)).Elem()

type KerberosRc4EncryptionPtrInput interface {
	pulumi.Input

	ToKerberosRc4EncryptionPtrOutput() KerberosRc4EncryptionPtrOutput
	ToKerberosRc4EncryptionPtrOutputWithContext(context.Context) KerberosRc4EncryptionPtrOutput
}

type kerberosRc4EncryptionPtr string

func KerberosRc4EncryptionPtr(v string) KerberosRc4EncryptionPtrInput {
	return (*kerberosRc4EncryptionPtr)(&v)
}

func (*kerberosRc4EncryptionPtr) ElementType() reflect.Type {
	return kerberosRc4EncryptionPtrType
}

func (in *kerberosRc4EncryptionPtr) ToKerberosRc4EncryptionPtrOutput() KerberosRc4EncryptionPtrOutput {
	return pulumi.ToOutput(in).(KerberosRc4EncryptionPtrOutput)
}

func (in *kerberosRc4EncryptionPtr) ToKerberosRc4EncryptionPtrOutputWithContext(ctx context.Context) KerberosRc4EncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(KerberosRc4EncryptionPtrOutput)
}

type Ldaps string

const (
	LdapsEnabled  = Ldaps("Enabled")
	LdapsDisabled = Ldaps("Disabled")
)

func (Ldaps) ElementType() reflect.Type {
	return reflect.TypeOf((*Ldaps)(nil)).Elem()
}

func (e Ldaps) ToLdapsOutput() LdapsOutput {
	return pulumi.ToOutput(e).(LdapsOutput)
}

func (e Ldaps) ToLdapsOutputWithContext(ctx context.Context) LdapsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(LdapsOutput)
}

func (e Ldaps) ToLdapsPtrOutput() LdapsPtrOutput {
	return e.ToLdapsPtrOutputWithContext(context.Background())
}

func (e Ldaps) ToLdapsPtrOutputWithContext(ctx context.Context) LdapsPtrOutput {
	return Ldaps(e).ToLdapsOutputWithContext(ctx).ToLdapsPtrOutputWithContext(ctx)
}

func (e Ldaps) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Ldaps) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Ldaps) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Ldaps) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type LdapsOutput struct{ *pulumi.OutputState }

func (LdapsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Ldaps)(nil)).Elem()
}

func (o LdapsOutput) ToLdapsOutput() LdapsOutput {
	return o
}

func (o LdapsOutput) ToLdapsOutputWithContext(ctx context.Context) LdapsOutput {
	return o
}

func (o LdapsOutput) ToLdapsPtrOutput() LdapsPtrOutput {
	return o.ToLdapsPtrOutputWithContext(context.Background())
}

func (o LdapsOutput) ToLdapsPtrOutputWithContext(ctx context.Context) LdapsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Ldaps) *Ldaps {
		return &v
	}).(LdapsPtrOutput)
}

func (o LdapsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o LdapsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Ldaps) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o LdapsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LdapsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Ldaps) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type LdapsPtrOutput struct{ *pulumi.OutputState }

func (LdapsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ldaps)(nil)).Elem()
}

func (o LdapsPtrOutput) ToLdapsPtrOutput() LdapsPtrOutput {
	return o
}

func (o LdapsPtrOutput) ToLdapsPtrOutputWithContext(ctx context.Context) LdapsPtrOutput {
	return o
}

func (o LdapsPtrOutput) Elem() LdapsOutput {
	return o.ApplyT(func(v *Ldaps) Ldaps {
		if v != nil {
			return *v
		}
		var ret Ldaps
		return ret
	}).(LdapsOutput)
}

func (o LdapsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LdapsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Ldaps) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type LdapsInput interface {
	pulumi.Input

	ToLdapsOutput() LdapsOutput
	ToLdapsOutputWithContext(context.Context) LdapsOutput
}

var ldapsPtrType = reflect.TypeOf((**Ldaps)(nil)).Elem()

type LdapsPtrInput interface {
	pulumi.Input

	ToLdapsPtrOutput() LdapsPtrOutput
	ToLdapsPtrOutputWithContext(context.Context) LdapsPtrOutput
}

type ldapsPtr string

func LdapsPtr(v string) LdapsPtrInput {
	return (*ldapsPtr)(&v)
}

func (*ldapsPtr) ElementType() reflect.Type {
	return ldapsPtrType
}

func (in *ldapsPtr) ToLdapsPtrOutput() LdapsPtrOutput {
	return pulumi.ToOutput(in).(LdapsPtrOutput)
}

func (in *ldapsPtr) ToLdapsPtrOutputWithContext(ctx context.Context) LdapsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(LdapsPtrOutput)
}

type NotifyDcAdmins string

const (
	NotifyDcAdminsEnabled  = NotifyDcAdmins("Enabled")
	NotifyDcAdminsDisabled = NotifyDcAdmins("Disabled")
)

func (NotifyDcAdmins) ElementType() reflect.Type {
	return reflect.TypeOf((*NotifyDcAdmins)(nil)).Elem()
}

func (e NotifyDcAdmins) ToNotifyDcAdminsOutput() NotifyDcAdminsOutput {
	return pulumi.ToOutput(e).(NotifyDcAdminsOutput)
}

func (e NotifyDcAdmins) ToNotifyDcAdminsOutputWithContext(ctx context.Context) NotifyDcAdminsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(NotifyDcAdminsOutput)
}

func (e NotifyDcAdmins) ToNotifyDcAdminsPtrOutput() NotifyDcAdminsPtrOutput {
	return e.ToNotifyDcAdminsPtrOutputWithContext(context.Background())
}

func (e NotifyDcAdmins) ToNotifyDcAdminsPtrOutputWithContext(ctx context.Context) NotifyDcAdminsPtrOutput {
	return NotifyDcAdmins(e).ToNotifyDcAdminsOutputWithContext(ctx).ToNotifyDcAdminsPtrOutputWithContext(ctx)
}

func (e NotifyDcAdmins) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e NotifyDcAdmins) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e NotifyDcAdmins) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e NotifyDcAdmins) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type NotifyDcAdminsOutput struct{ *pulumi.OutputState }

func (NotifyDcAdminsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotifyDcAdmins)(nil)).Elem()
}

func (o NotifyDcAdminsOutput) ToNotifyDcAdminsOutput() NotifyDcAdminsOutput {
	return o
}

func (o NotifyDcAdminsOutput) ToNotifyDcAdminsOutputWithContext(ctx context.Context) NotifyDcAdminsOutput {
	return o
}

func (o NotifyDcAdminsOutput) ToNotifyDcAdminsPtrOutput() NotifyDcAdminsPtrOutput {
	return o.ToNotifyDcAdminsPtrOutputWithContext(context.Background())
}

func (o NotifyDcAdminsOutput) ToNotifyDcAdminsPtrOutputWithContext(ctx context.Context) NotifyDcAdminsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NotifyDcAdmins) *NotifyDcAdmins {
		return &v
	}).(NotifyDcAdminsPtrOutput)
}

func (o NotifyDcAdminsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o NotifyDcAdminsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NotifyDcAdmins) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o NotifyDcAdminsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NotifyDcAdminsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NotifyDcAdmins) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type NotifyDcAdminsPtrOutput struct{ *pulumi.OutputState }

func (NotifyDcAdminsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotifyDcAdmins)(nil)).Elem()
}

func (o NotifyDcAdminsPtrOutput) ToNotifyDcAdminsPtrOutput() NotifyDcAdminsPtrOutput {
	return o
}

func (o NotifyDcAdminsPtrOutput) ToNotifyDcAdminsPtrOutputWithContext(ctx context.Context) NotifyDcAdminsPtrOutput {
	return o
}

func (o NotifyDcAdminsPtrOutput) Elem() NotifyDcAdminsOutput {
	return o.ApplyT(func(v *NotifyDcAdmins) NotifyDcAdmins {
		if v != nil {
			return *v
		}
		var ret NotifyDcAdmins
		return ret
	}).(NotifyDcAdminsOutput)
}

func (o NotifyDcAdminsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NotifyDcAdminsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *NotifyDcAdmins) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type NotifyDcAdminsInput interface {
	pulumi.Input

	ToNotifyDcAdminsOutput() NotifyDcAdminsOutput
	ToNotifyDcAdminsOutputWithContext(context.Context) NotifyDcAdminsOutput
}

var notifyDcAdminsPtrType = reflect.TypeOf((**NotifyDcAdmins)(nil)).Elem()

type NotifyDcAdminsPtrInput interface {
	pulumi.Input

	ToNotifyDcAdminsPtrOutput() NotifyDcAdminsPtrOutput
	ToNotifyDcAdminsPtrOutputWithContext(context.Context) NotifyDcAdminsPtrOutput
}

type notifyDcAdminsPtr string

func NotifyDcAdminsPtr(v string) NotifyDcAdminsPtrInput {
	return (*notifyDcAdminsPtr)(&v)
}

func (*notifyDcAdminsPtr) ElementType() reflect.Type {
	return notifyDcAdminsPtrType
}

func (in *notifyDcAdminsPtr) ToNotifyDcAdminsPtrOutput() NotifyDcAdminsPtrOutput {
	return pulumi.ToOutput(in).(NotifyDcAdminsPtrOutput)
}

func (in *notifyDcAdminsPtr) ToNotifyDcAdminsPtrOutputWithContext(ctx context.Context) NotifyDcAdminsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(NotifyDcAdminsPtrOutput)
}

type NotifyGlobalAdmins string

const (
	NotifyGlobalAdminsEnabled  = NotifyGlobalAdmins("Enabled")
	NotifyGlobalAdminsDisabled = NotifyGlobalAdmins("Disabled")
)

func (NotifyGlobalAdmins) ElementType() reflect.Type {
	return reflect.TypeOf((*NotifyGlobalAdmins)(nil)).Elem()
}

func (e NotifyGlobalAdmins) ToNotifyGlobalAdminsOutput() NotifyGlobalAdminsOutput {
	return pulumi.ToOutput(e).(NotifyGlobalAdminsOutput)
}

func (e NotifyGlobalAdmins) ToNotifyGlobalAdminsOutputWithContext(ctx context.Context) NotifyGlobalAdminsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(NotifyGlobalAdminsOutput)
}

func (e NotifyGlobalAdmins) ToNotifyGlobalAdminsPtrOutput() NotifyGlobalAdminsPtrOutput {
	return e.ToNotifyGlobalAdminsPtrOutputWithContext(context.Background())
}

func (e NotifyGlobalAdmins) ToNotifyGlobalAdminsPtrOutputWithContext(ctx context.Context) NotifyGlobalAdminsPtrOutput {
	return NotifyGlobalAdmins(e).ToNotifyGlobalAdminsOutputWithContext(ctx).ToNotifyGlobalAdminsPtrOutputWithContext(ctx)
}

func (e NotifyGlobalAdmins) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e NotifyGlobalAdmins) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e NotifyGlobalAdmins) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e NotifyGlobalAdmins) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type NotifyGlobalAdminsOutput struct{ *pulumi.OutputState }

func (NotifyGlobalAdminsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotifyGlobalAdmins)(nil)).Elem()
}

func (o NotifyGlobalAdminsOutput) ToNotifyGlobalAdminsOutput() NotifyGlobalAdminsOutput {
	return o
}

func (o NotifyGlobalAdminsOutput) ToNotifyGlobalAdminsOutputWithContext(ctx context.Context) NotifyGlobalAdminsOutput {
	return o
}

func (o NotifyGlobalAdminsOutput) ToNotifyGlobalAdminsPtrOutput() NotifyGlobalAdminsPtrOutput {
	return o.ToNotifyGlobalAdminsPtrOutputWithContext(context.Background())
}

func (o NotifyGlobalAdminsOutput) ToNotifyGlobalAdminsPtrOutputWithContext(ctx context.Context) NotifyGlobalAdminsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NotifyGlobalAdmins) *NotifyGlobalAdmins {
		return &v
	}).(NotifyGlobalAdminsPtrOutput)
}

func (o NotifyGlobalAdminsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o NotifyGlobalAdminsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NotifyGlobalAdmins) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o NotifyGlobalAdminsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NotifyGlobalAdminsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NotifyGlobalAdmins) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type NotifyGlobalAdminsPtrOutput struct{ *pulumi.OutputState }

func (NotifyGlobalAdminsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotifyGlobalAdmins)(nil)).Elem()
}

func (o NotifyGlobalAdminsPtrOutput) ToNotifyGlobalAdminsPtrOutput() NotifyGlobalAdminsPtrOutput {
	return o
}

func (o NotifyGlobalAdminsPtrOutput) ToNotifyGlobalAdminsPtrOutputWithContext(ctx context.Context) NotifyGlobalAdminsPtrOutput {
	return o
}

func (o NotifyGlobalAdminsPtrOutput) Elem() NotifyGlobalAdminsOutput {
	return o.ApplyT(func(v *NotifyGlobalAdmins) NotifyGlobalAdmins {
		if v != nil {
			return *v
		}
		var ret NotifyGlobalAdmins
		return ret
	}).(NotifyGlobalAdminsOutput)
}

func (o NotifyGlobalAdminsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NotifyGlobalAdminsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *NotifyGlobalAdmins) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type NotifyGlobalAdminsInput interface {
	pulumi.Input

	ToNotifyGlobalAdminsOutput() NotifyGlobalAdminsOutput
	ToNotifyGlobalAdminsOutputWithContext(context.Context) NotifyGlobalAdminsOutput
}

var notifyGlobalAdminsPtrType = reflect.TypeOf((**NotifyGlobalAdmins)(nil)).Elem()

type NotifyGlobalAdminsPtrInput interface {
	pulumi.Input

	ToNotifyGlobalAdminsPtrOutput() NotifyGlobalAdminsPtrOutput
	ToNotifyGlobalAdminsPtrOutputWithContext(context.Context) NotifyGlobalAdminsPtrOutput
}

type notifyGlobalAdminsPtr string

func NotifyGlobalAdminsPtr(v string) NotifyGlobalAdminsPtrInput {
	return (*notifyGlobalAdminsPtr)(&v)
}

func (*notifyGlobalAdminsPtr) ElementType() reflect.Type {
	return notifyGlobalAdminsPtrType
}

func (in *notifyGlobalAdminsPtr) ToNotifyGlobalAdminsPtrOutput() NotifyGlobalAdminsPtrOutput {
	return pulumi.ToOutput(in).(NotifyGlobalAdminsPtrOutput)
}

func (in *notifyGlobalAdminsPtr) ToNotifyGlobalAdminsPtrOutputWithContext(ctx context.Context) NotifyGlobalAdminsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(NotifyGlobalAdminsPtrOutput)
}

type NtlmV1 string

const (
	NtlmV1Enabled  = NtlmV1("Enabled")
	NtlmV1Disabled = NtlmV1("Disabled")
)

func (NtlmV1) ElementType() reflect.Type {
	return reflect.TypeOf((*NtlmV1)(nil)).Elem()
}

func (e NtlmV1) ToNtlmV1Output() NtlmV1Output {
	return pulumi.ToOutput(e).(NtlmV1Output)
}

func (e NtlmV1) ToNtlmV1OutputWithContext(ctx context.Context) NtlmV1Output {
	return pulumi.ToOutputWithContext(ctx, e).(NtlmV1Output)
}

func (e NtlmV1) ToNtlmV1PtrOutput() NtlmV1PtrOutput {
	return e.ToNtlmV1PtrOutputWithContext(context.Background())
}

func (e NtlmV1) ToNtlmV1PtrOutputWithContext(ctx context.Context) NtlmV1PtrOutput {
	return NtlmV1(e).ToNtlmV1OutputWithContext(ctx).ToNtlmV1PtrOutputWithContext(ctx)
}

func (e NtlmV1) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e NtlmV1) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e NtlmV1) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e NtlmV1) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type NtlmV1Output struct{ *pulumi.OutputState }

func (NtlmV1Output) ElementType() reflect.Type {
	return reflect.TypeOf((*NtlmV1)(nil)).Elem()
}

func (o NtlmV1Output) ToNtlmV1Output() NtlmV1Output {
	return o
}

func (o NtlmV1Output) ToNtlmV1OutputWithContext(ctx context.Context) NtlmV1Output {
	return o
}

func (o NtlmV1Output) ToNtlmV1PtrOutput() NtlmV1PtrOutput {
	return o.ToNtlmV1PtrOutputWithContext(context.Background())
}

func (o NtlmV1Output) ToNtlmV1PtrOutputWithContext(ctx context.Context) NtlmV1PtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NtlmV1) *NtlmV1 {
		return &v
	}).(NtlmV1PtrOutput)
}

func (o NtlmV1Output) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o NtlmV1Output) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NtlmV1) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o NtlmV1Output) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NtlmV1Output) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NtlmV1) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type NtlmV1PtrOutput struct{ *pulumi.OutputState }

func (NtlmV1PtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NtlmV1)(nil)).Elem()
}

func (o NtlmV1PtrOutput) ToNtlmV1PtrOutput() NtlmV1PtrOutput {
	return o
}

func (o NtlmV1PtrOutput) ToNtlmV1PtrOutputWithContext(ctx context.Context) NtlmV1PtrOutput {
	return o
}

func (o NtlmV1PtrOutput) Elem() NtlmV1Output {
	return o.ApplyT(func(v *NtlmV1) NtlmV1 {
		if v != nil {
			return *v
		}
		var ret NtlmV1
		return ret
	}).(NtlmV1Output)
}

func (o NtlmV1PtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NtlmV1PtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *NtlmV1) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type NtlmV1Input interface {
	pulumi.Input

	ToNtlmV1Output() NtlmV1Output
	ToNtlmV1OutputWithContext(context.Context) NtlmV1Output
}

var ntlmV1PtrType = reflect.TypeOf((**NtlmV1)(nil)).Elem()

type NtlmV1PtrInput interface {
	pulumi.Input

	ToNtlmV1PtrOutput() NtlmV1PtrOutput
	ToNtlmV1PtrOutputWithContext(context.Context) NtlmV1PtrOutput
}

type ntlmV1Ptr string

func NtlmV1Ptr(v string) NtlmV1PtrInput {
	return (*ntlmV1Ptr)(&v)
}

func (*ntlmV1Ptr) ElementType() reflect.Type {
	return ntlmV1PtrType
}

func (in *ntlmV1Ptr) ToNtlmV1PtrOutput() NtlmV1PtrOutput {
	return pulumi.ToOutput(in).(NtlmV1PtrOutput)
}

func (in *ntlmV1Ptr) ToNtlmV1PtrOutputWithContext(ctx context.Context) NtlmV1PtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(NtlmV1PtrOutput)
}

type SyncKerberosPasswords string

const (
	SyncKerberosPasswordsEnabled  = SyncKerberosPasswords("Enabled")
	SyncKerberosPasswordsDisabled = SyncKerberosPasswords("Disabled")
)

func (SyncKerberosPasswords) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncKerberosPasswords)(nil)).Elem()
}

func (e SyncKerberosPasswords) ToSyncKerberosPasswordsOutput() SyncKerberosPasswordsOutput {
	return pulumi.ToOutput(e).(SyncKerberosPasswordsOutput)
}

func (e SyncKerberosPasswords) ToSyncKerberosPasswordsOutputWithContext(ctx context.Context) SyncKerberosPasswordsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SyncKerberosPasswordsOutput)
}

func (e SyncKerberosPasswords) ToSyncKerberosPasswordsPtrOutput() SyncKerberosPasswordsPtrOutput {
	return e.ToSyncKerberosPasswordsPtrOutputWithContext(context.Background())
}

func (e SyncKerberosPasswords) ToSyncKerberosPasswordsPtrOutputWithContext(ctx context.Context) SyncKerberosPasswordsPtrOutput {
	return SyncKerberosPasswords(e).ToSyncKerberosPasswordsOutputWithContext(ctx).ToSyncKerberosPasswordsPtrOutputWithContext(ctx)
}

func (e SyncKerberosPasswords) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SyncKerberosPasswords) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SyncKerberosPasswords) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SyncKerberosPasswords) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SyncKerberosPasswordsOutput struct{ *pulumi.OutputState }

func (SyncKerberosPasswordsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncKerberosPasswords)(nil)).Elem()
}

func (o SyncKerberosPasswordsOutput) ToSyncKerberosPasswordsOutput() SyncKerberosPasswordsOutput {
	return o
}

func (o SyncKerberosPasswordsOutput) ToSyncKerberosPasswordsOutputWithContext(ctx context.Context) SyncKerberosPasswordsOutput {
	return o
}

func (o SyncKerberosPasswordsOutput) ToSyncKerberosPasswordsPtrOutput() SyncKerberosPasswordsPtrOutput {
	return o.ToSyncKerberosPasswordsPtrOutputWithContext(context.Background())
}

func (o SyncKerberosPasswordsOutput) ToSyncKerberosPasswordsPtrOutputWithContext(ctx context.Context) SyncKerberosPasswordsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SyncKerberosPasswords) *SyncKerberosPasswords {
		return &v
	}).(SyncKerberosPasswordsPtrOutput)
}

func (o SyncKerberosPasswordsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SyncKerberosPasswordsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SyncKerberosPasswords) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SyncKerberosPasswordsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SyncKerberosPasswordsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SyncKerberosPasswords) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SyncKerberosPasswordsPtrOutput struct{ *pulumi.OutputState }

func (SyncKerberosPasswordsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncKerberosPasswords)(nil)).Elem()
}

func (o SyncKerberosPasswordsPtrOutput) ToSyncKerberosPasswordsPtrOutput() SyncKerberosPasswordsPtrOutput {
	return o
}

func (o SyncKerberosPasswordsPtrOutput) ToSyncKerberosPasswordsPtrOutputWithContext(ctx context.Context) SyncKerberosPasswordsPtrOutput {
	return o
}

func (o SyncKerberosPasswordsPtrOutput) Elem() SyncKerberosPasswordsOutput {
	return o.ApplyT(func(v *SyncKerberosPasswords) SyncKerberosPasswords {
		if v != nil {
			return *v
		}
		var ret SyncKerberosPasswords
		return ret
	}).(SyncKerberosPasswordsOutput)
}

func (o SyncKerberosPasswordsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SyncKerberosPasswordsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SyncKerberosPasswords) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SyncKerberosPasswordsInput interface {
	pulumi.Input

	ToSyncKerberosPasswordsOutput() SyncKerberosPasswordsOutput
	ToSyncKerberosPasswordsOutputWithContext(context.Context) SyncKerberosPasswordsOutput
}

var syncKerberosPasswordsPtrType = reflect.TypeOf((**SyncKerberosPasswords)(nil)).Elem()

type SyncKerberosPasswordsPtrInput interface {
	pulumi.Input

	ToSyncKerberosPasswordsPtrOutput() SyncKerberosPasswordsPtrOutput
	ToSyncKerberosPasswordsPtrOutputWithContext(context.Context) SyncKerberosPasswordsPtrOutput
}

type syncKerberosPasswordsPtr string

func SyncKerberosPasswordsPtr(v string) SyncKerberosPasswordsPtrInput {
	return (*syncKerberosPasswordsPtr)(&v)
}

func (*syncKerberosPasswordsPtr) ElementType() reflect.Type {
	return syncKerberosPasswordsPtrType
}

func (in *syncKerberosPasswordsPtr) ToSyncKerberosPasswordsPtrOutput() SyncKerberosPasswordsPtrOutput {
	return pulumi.ToOutput(in).(SyncKerberosPasswordsPtrOutput)
}

func (in *syncKerberosPasswordsPtr) ToSyncKerberosPasswordsPtrOutputWithContext(ctx context.Context) SyncKerberosPasswordsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SyncKerberosPasswordsPtrOutput)
}

type SyncNtlmPasswords string

const (
	SyncNtlmPasswordsEnabled  = SyncNtlmPasswords("Enabled")
	SyncNtlmPasswordsDisabled = SyncNtlmPasswords("Disabled")
)

func (SyncNtlmPasswords) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncNtlmPasswords)(nil)).Elem()
}

func (e SyncNtlmPasswords) ToSyncNtlmPasswordsOutput() SyncNtlmPasswordsOutput {
	return pulumi.ToOutput(e).(SyncNtlmPasswordsOutput)
}

func (e SyncNtlmPasswords) ToSyncNtlmPasswordsOutputWithContext(ctx context.Context) SyncNtlmPasswordsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SyncNtlmPasswordsOutput)
}

func (e SyncNtlmPasswords) ToSyncNtlmPasswordsPtrOutput() SyncNtlmPasswordsPtrOutput {
	return e.ToSyncNtlmPasswordsPtrOutputWithContext(context.Background())
}

func (e SyncNtlmPasswords) ToSyncNtlmPasswordsPtrOutputWithContext(ctx context.Context) SyncNtlmPasswordsPtrOutput {
	return SyncNtlmPasswords(e).ToSyncNtlmPasswordsOutputWithContext(ctx).ToSyncNtlmPasswordsPtrOutputWithContext(ctx)
}

func (e SyncNtlmPasswords) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SyncNtlmPasswords) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SyncNtlmPasswords) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SyncNtlmPasswords) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SyncNtlmPasswordsOutput struct{ *pulumi.OutputState }

func (SyncNtlmPasswordsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncNtlmPasswords)(nil)).Elem()
}

func (o SyncNtlmPasswordsOutput) ToSyncNtlmPasswordsOutput() SyncNtlmPasswordsOutput {
	return o
}

func (o SyncNtlmPasswordsOutput) ToSyncNtlmPasswordsOutputWithContext(ctx context.Context) SyncNtlmPasswordsOutput {
	return o
}

func (o SyncNtlmPasswordsOutput) ToSyncNtlmPasswordsPtrOutput() SyncNtlmPasswordsPtrOutput {
	return o.ToSyncNtlmPasswordsPtrOutputWithContext(context.Background())
}

func (o SyncNtlmPasswordsOutput) ToSyncNtlmPasswordsPtrOutputWithContext(ctx context.Context) SyncNtlmPasswordsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SyncNtlmPasswords) *SyncNtlmPasswords {
		return &v
	}).(SyncNtlmPasswordsPtrOutput)
}

func (o SyncNtlmPasswordsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SyncNtlmPasswordsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SyncNtlmPasswords) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SyncNtlmPasswordsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SyncNtlmPasswordsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SyncNtlmPasswords) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SyncNtlmPasswordsPtrOutput struct{ *pulumi.OutputState }

func (SyncNtlmPasswordsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncNtlmPasswords)(nil)).Elem()
}

func (o SyncNtlmPasswordsPtrOutput) ToSyncNtlmPasswordsPtrOutput() SyncNtlmPasswordsPtrOutput {
	return o
}

func (o SyncNtlmPasswordsPtrOutput) ToSyncNtlmPasswordsPtrOutputWithContext(ctx context.Context) SyncNtlmPasswordsPtrOutput {
	return o
}

func (o SyncNtlmPasswordsPtrOutput) Elem() SyncNtlmPasswordsOutput {
	return o.ApplyT(func(v *SyncNtlmPasswords) SyncNtlmPasswords {
		if v != nil {
			return *v
		}
		var ret SyncNtlmPasswords
		return ret
	}).(SyncNtlmPasswordsOutput)
}

func (o SyncNtlmPasswordsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SyncNtlmPasswordsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SyncNtlmPasswords) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SyncNtlmPasswordsInput interface {
	pulumi.Input

	ToSyncNtlmPasswordsOutput() SyncNtlmPasswordsOutput
	ToSyncNtlmPasswordsOutputWithContext(context.Context) SyncNtlmPasswordsOutput
}

var syncNtlmPasswordsPtrType = reflect.TypeOf((**SyncNtlmPasswords)(nil)).Elem()

type SyncNtlmPasswordsPtrInput interface {
	pulumi.Input

	ToSyncNtlmPasswordsPtrOutput() SyncNtlmPasswordsPtrOutput
	ToSyncNtlmPasswordsPtrOutputWithContext(context.Context) SyncNtlmPasswordsPtrOutput
}

type syncNtlmPasswordsPtr string

func SyncNtlmPasswordsPtr(v string) SyncNtlmPasswordsPtrInput {
	return (*syncNtlmPasswordsPtr)(&v)
}

func (*syncNtlmPasswordsPtr) ElementType() reflect.Type {
	return syncNtlmPasswordsPtrType
}

func (in *syncNtlmPasswordsPtr) ToSyncNtlmPasswordsPtrOutput() SyncNtlmPasswordsPtrOutput {
	return pulumi.ToOutput(in).(SyncNtlmPasswordsPtrOutput)
}

func (in *syncNtlmPasswordsPtr) ToSyncNtlmPasswordsPtrOutputWithContext(ctx context.Context) SyncNtlmPasswordsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SyncNtlmPasswordsPtrOutput)
}

type SyncOnPremPasswords string

const (
	SyncOnPremPasswordsEnabled  = SyncOnPremPasswords("Enabled")
	SyncOnPremPasswordsDisabled = SyncOnPremPasswords("Disabled")
)

func (SyncOnPremPasswords) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncOnPremPasswords)(nil)).Elem()
}

func (e SyncOnPremPasswords) ToSyncOnPremPasswordsOutput() SyncOnPremPasswordsOutput {
	return pulumi.ToOutput(e).(SyncOnPremPasswordsOutput)
}

func (e SyncOnPremPasswords) ToSyncOnPremPasswordsOutputWithContext(ctx context.Context) SyncOnPremPasswordsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SyncOnPremPasswordsOutput)
}

func (e SyncOnPremPasswords) ToSyncOnPremPasswordsPtrOutput() SyncOnPremPasswordsPtrOutput {
	return e.ToSyncOnPremPasswordsPtrOutputWithContext(context.Background())
}

func (e SyncOnPremPasswords) ToSyncOnPremPasswordsPtrOutputWithContext(ctx context.Context) SyncOnPremPasswordsPtrOutput {
	return SyncOnPremPasswords(e).ToSyncOnPremPasswordsOutputWithContext(ctx).ToSyncOnPremPasswordsPtrOutputWithContext(ctx)
}

func (e SyncOnPremPasswords) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SyncOnPremPasswords) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SyncOnPremPasswords) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SyncOnPremPasswords) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SyncOnPremPasswordsOutput struct{ *pulumi.OutputState }

func (SyncOnPremPasswordsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncOnPremPasswords)(nil)).Elem()
}

func (o SyncOnPremPasswordsOutput) ToSyncOnPremPasswordsOutput() SyncOnPremPasswordsOutput {
	return o
}

func (o SyncOnPremPasswordsOutput) ToSyncOnPremPasswordsOutputWithContext(ctx context.Context) SyncOnPremPasswordsOutput {
	return o
}

func (o SyncOnPremPasswordsOutput) ToSyncOnPremPasswordsPtrOutput() SyncOnPremPasswordsPtrOutput {
	return o.ToSyncOnPremPasswordsPtrOutputWithContext(context.Background())
}

func (o SyncOnPremPasswordsOutput) ToSyncOnPremPasswordsPtrOutputWithContext(ctx context.Context) SyncOnPremPasswordsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SyncOnPremPasswords) *SyncOnPremPasswords {
		return &v
	}).(SyncOnPremPasswordsPtrOutput)
}

func (o SyncOnPremPasswordsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SyncOnPremPasswordsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SyncOnPremPasswords) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SyncOnPremPasswordsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SyncOnPremPasswordsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SyncOnPremPasswords) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SyncOnPremPasswordsPtrOutput struct{ *pulumi.OutputState }

func (SyncOnPremPasswordsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncOnPremPasswords)(nil)).Elem()
}

func (o SyncOnPremPasswordsPtrOutput) ToSyncOnPremPasswordsPtrOutput() SyncOnPremPasswordsPtrOutput {
	return o
}

func (o SyncOnPremPasswordsPtrOutput) ToSyncOnPremPasswordsPtrOutputWithContext(ctx context.Context) SyncOnPremPasswordsPtrOutput {
	return o
}

func (o SyncOnPremPasswordsPtrOutput) Elem() SyncOnPremPasswordsOutput {
	return o.ApplyT(func(v *SyncOnPremPasswords) SyncOnPremPasswords {
		if v != nil {
			return *v
		}
		var ret SyncOnPremPasswords
		return ret
	}).(SyncOnPremPasswordsOutput)
}

func (o SyncOnPremPasswordsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SyncOnPremPasswordsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SyncOnPremPasswords) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SyncOnPremPasswordsInput interface {
	pulumi.Input

	ToSyncOnPremPasswordsOutput() SyncOnPremPasswordsOutput
	ToSyncOnPremPasswordsOutputWithContext(context.Context) SyncOnPremPasswordsOutput
}

var syncOnPremPasswordsPtrType = reflect.TypeOf((**SyncOnPremPasswords)(nil)).Elem()

type SyncOnPremPasswordsPtrInput interface {
	pulumi.Input

	ToSyncOnPremPasswordsPtrOutput() SyncOnPremPasswordsPtrOutput
	ToSyncOnPremPasswordsPtrOutputWithContext(context.Context) SyncOnPremPasswordsPtrOutput
}

type syncOnPremPasswordsPtr string

func SyncOnPremPasswordsPtr(v string) SyncOnPremPasswordsPtrInput {
	return (*syncOnPremPasswordsPtr)(&v)
}

func (*syncOnPremPasswordsPtr) ElementType() reflect.Type {
	return syncOnPremPasswordsPtrType
}

func (in *syncOnPremPasswordsPtr) ToSyncOnPremPasswordsPtrOutput() SyncOnPremPasswordsPtrOutput {
	return pulumi.ToOutput(in).(SyncOnPremPasswordsPtrOutput)
}

func (in *syncOnPremPasswordsPtr) ToSyncOnPremPasswordsPtrOutputWithContext(ctx context.Context) SyncOnPremPasswordsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SyncOnPremPasswordsPtrOutput)
}

type TlsV1 string

const (
	TlsV1Enabled  = TlsV1("Enabled")
	TlsV1Disabled = TlsV1("Disabled")
)

func (TlsV1) ElementType() reflect.Type {
	return reflect.TypeOf((*TlsV1)(nil)).Elem()
}

func (e TlsV1) ToTlsV1Output() TlsV1Output {
	return pulumi.ToOutput(e).(TlsV1Output)
}

func (e TlsV1) ToTlsV1OutputWithContext(ctx context.Context) TlsV1Output {
	return pulumi.ToOutputWithContext(ctx, e).(TlsV1Output)
}

func (e TlsV1) ToTlsV1PtrOutput() TlsV1PtrOutput {
	return e.ToTlsV1PtrOutputWithContext(context.Background())
}

func (e TlsV1) ToTlsV1PtrOutputWithContext(ctx context.Context) TlsV1PtrOutput {
	return TlsV1(e).ToTlsV1OutputWithContext(ctx).ToTlsV1PtrOutputWithContext(ctx)
}

func (e TlsV1) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TlsV1) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TlsV1) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TlsV1) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TlsV1Output struct{ *pulumi.OutputState }

func (TlsV1Output) ElementType() reflect.Type {
	return reflect.TypeOf((*TlsV1)(nil)).Elem()
}

func (o TlsV1Output) ToTlsV1Output() TlsV1Output {
	return o
}

func (o TlsV1Output) ToTlsV1OutputWithContext(ctx context.Context) TlsV1Output {
	return o
}

func (o TlsV1Output) ToTlsV1PtrOutput() TlsV1PtrOutput {
	return o.ToTlsV1PtrOutputWithContext(context.Background())
}

func (o TlsV1Output) ToTlsV1PtrOutputWithContext(ctx context.Context) TlsV1PtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TlsV1) *TlsV1 {
		return &v
	}).(TlsV1PtrOutput)
}

func (o TlsV1Output) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TlsV1Output) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TlsV1) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TlsV1Output) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TlsV1Output) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TlsV1) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TlsV1PtrOutput struct{ *pulumi.OutputState }

func (TlsV1PtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TlsV1)(nil)).Elem()
}

func (o TlsV1PtrOutput) ToTlsV1PtrOutput() TlsV1PtrOutput {
	return o
}

func (o TlsV1PtrOutput) ToTlsV1PtrOutputWithContext(ctx context.Context) TlsV1PtrOutput {
	return o
}

func (o TlsV1PtrOutput) Elem() TlsV1Output {
	return o.ApplyT(func(v *TlsV1) TlsV1 {
		if v != nil {
			return *v
		}
		var ret TlsV1
		return ret
	}).(TlsV1Output)
}

func (o TlsV1PtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TlsV1PtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TlsV1) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type TlsV1Input interface {
	pulumi.Input

	ToTlsV1Output() TlsV1Output
	ToTlsV1OutputWithContext(context.Context) TlsV1Output
}

var tlsV1PtrType = reflect.TypeOf((**TlsV1)(nil)).Elem()

type TlsV1PtrInput interface {
	pulumi.Input

	ToTlsV1PtrOutput() TlsV1PtrOutput
	ToTlsV1PtrOutputWithContext(context.Context) TlsV1PtrOutput
}

type tlsV1Ptr string

func TlsV1Ptr(v string) TlsV1PtrInput {
	return (*tlsV1Ptr)(&v)
}

func (*tlsV1Ptr) ElementType() reflect.Type {
	return tlsV1PtrType
}

func (in *tlsV1Ptr) ToTlsV1PtrOutput() TlsV1PtrOutput {
	return pulumi.ToOutput(in).(TlsV1PtrOutput)
}

func (in *tlsV1Ptr) ToTlsV1PtrOutputWithContext(ctx context.Context) TlsV1PtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TlsV1PtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ExternalAccessOutput{})
	pulumi.RegisterOutputType(ExternalAccessPtrOutput{})
	pulumi.RegisterOutputType(FilteredSyncOutput{})
	pulumi.RegisterOutputType(FilteredSyncPtrOutput{})
	pulumi.RegisterOutputType(KerberosArmoringOutput{})
	pulumi.RegisterOutputType(KerberosArmoringPtrOutput{})
	pulumi.RegisterOutputType(KerberosRc4EncryptionOutput{})
	pulumi.RegisterOutputType(KerberosRc4EncryptionPtrOutput{})
	pulumi.RegisterOutputType(LdapsOutput{})
	pulumi.RegisterOutputType(LdapsPtrOutput{})
	pulumi.RegisterOutputType(NotifyDcAdminsOutput{})
	pulumi.RegisterOutputType(NotifyDcAdminsPtrOutput{})
	pulumi.RegisterOutputType(NotifyGlobalAdminsOutput{})
	pulumi.RegisterOutputType(NotifyGlobalAdminsPtrOutput{})
	pulumi.RegisterOutputType(NtlmV1Output{})
	pulumi.RegisterOutputType(NtlmV1PtrOutput{})
	pulumi.RegisterOutputType(SyncKerberosPasswordsOutput{})
	pulumi.RegisterOutputType(SyncKerberosPasswordsPtrOutput{})
	pulumi.RegisterOutputType(SyncNtlmPasswordsOutput{})
	pulumi.RegisterOutputType(SyncNtlmPasswordsPtrOutput{})
	pulumi.RegisterOutputType(SyncOnPremPasswordsOutput{})
	pulumi.RegisterOutputType(SyncOnPremPasswordsPtrOutput{})
	pulumi.RegisterOutputType(TlsV1Output{})
	pulumi.RegisterOutputType(TlsV1PtrOutput{})
}
