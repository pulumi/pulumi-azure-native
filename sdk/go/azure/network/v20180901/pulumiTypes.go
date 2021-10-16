


package v20180901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ARecord struct {
	Ipv4Address *string `pulumi:"ipv4Address"`
}





type ARecordInput interface {
	pulumi.Input

	ToARecordOutput() ARecordOutput
	ToARecordOutputWithContext(context.Context) ARecordOutput
}

type ARecordArgs struct {
	Ipv4Address pulumi.StringPtrInput `pulumi:"ipv4Address"`
}

func (ARecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ARecord)(nil)).Elem()
}

func (i ARecordArgs) ToARecordOutput() ARecordOutput {
	return i.ToARecordOutputWithContext(context.Background())
}

func (i ARecordArgs) ToARecordOutputWithContext(ctx context.Context) ARecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ARecordOutput)
}





type ARecordArrayInput interface {
	pulumi.Input

	ToARecordArrayOutput() ARecordArrayOutput
	ToARecordArrayOutputWithContext(context.Context) ARecordArrayOutput
}

type ARecordArray []ARecordInput

func (ARecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ARecord)(nil)).Elem()
}

func (i ARecordArray) ToARecordArrayOutput() ARecordArrayOutput {
	return i.ToARecordArrayOutputWithContext(context.Background())
}

func (i ARecordArray) ToARecordArrayOutputWithContext(ctx context.Context) ARecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ARecordArrayOutput)
}

type ARecordOutput struct{ *pulumi.OutputState }

func (ARecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ARecord)(nil)).Elem()
}

func (o ARecordOutput) ToARecordOutput() ARecordOutput {
	return o
}

func (o ARecordOutput) ToARecordOutputWithContext(ctx context.Context) ARecordOutput {
	return o
}

func (o ARecordOutput) Ipv4Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ARecord) *string { return v.Ipv4Address }).(pulumi.StringPtrOutput)
}

type ARecordArrayOutput struct{ *pulumi.OutputState }

func (ARecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ARecord)(nil)).Elem()
}

func (o ARecordArrayOutput) ToARecordArrayOutput() ARecordArrayOutput {
	return o
}

func (o ARecordArrayOutput) ToARecordArrayOutputWithContext(ctx context.Context) ARecordArrayOutput {
	return o
}

func (o ARecordArrayOutput) Index(i pulumi.IntInput) ARecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ARecord {
		return vs[0].([]ARecord)[vs[1].(int)]
	}).(ARecordOutput)
}

type ARecordResponse struct {
	Ipv4Address *string `pulumi:"ipv4Address"`
}





type ARecordResponseInput interface {
	pulumi.Input

	ToARecordResponseOutput() ARecordResponseOutput
	ToARecordResponseOutputWithContext(context.Context) ARecordResponseOutput
}

type ARecordResponseArgs struct {
	Ipv4Address pulumi.StringPtrInput `pulumi:"ipv4Address"`
}

func (ARecordResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ARecordResponse)(nil)).Elem()
}

func (i ARecordResponseArgs) ToARecordResponseOutput() ARecordResponseOutput {
	return i.ToARecordResponseOutputWithContext(context.Background())
}

func (i ARecordResponseArgs) ToARecordResponseOutputWithContext(ctx context.Context) ARecordResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ARecordResponseOutput)
}





type ARecordResponseArrayInput interface {
	pulumi.Input

	ToARecordResponseArrayOutput() ARecordResponseArrayOutput
	ToARecordResponseArrayOutputWithContext(context.Context) ARecordResponseArrayOutput
}

type ARecordResponseArray []ARecordResponseInput

func (ARecordResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ARecordResponse)(nil)).Elem()
}

func (i ARecordResponseArray) ToARecordResponseArrayOutput() ARecordResponseArrayOutput {
	return i.ToARecordResponseArrayOutputWithContext(context.Background())
}

func (i ARecordResponseArray) ToARecordResponseArrayOutputWithContext(ctx context.Context) ARecordResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ARecordResponseArrayOutput)
}

type ARecordResponseOutput struct{ *pulumi.OutputState }

func (ARecordResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ARecordResponse)(nil)).Elem()
}

func (o ARecordResponseOutput) ToARecordResponseOutput() ARecordResponseOutput {
	return o
}

func (o ARecordResponseOutput) ToARecordResponseOutputWithContext(ctx context.Context) ARecordResponseOutput {
	return o
}

func (o ARecordResponseOutput) Ipv4Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ARecordResponse) *string { return v.Ipv4Address }).(pulumi.StringPtrOutput)
}

type ARecordResponseArrayOutput struct{ *pulumi.OutputState }

func (ARecordResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ARecordResponse)(nil)).Elem()
}

func (o ARecordResponseArrayOutput) ToARecordResponseArrayOutput() ARecordResponseArrayOutput {
	return o
}

func (o ARecordResponseArrayOutput) ToARecordResponseArrayOutputWithContext(ctx context.Context) ARecordResponseArrayOutput {
	return o
}

func (o ARecordResponseArrayOutput) Index(i pulumi.IntInput) ARecordResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ARecordResponse {
		return vs[0].([]ARecordResponse)[vs[1].(int)]
	}).(ARecordResponseOutput)
}

type AaaaRecord struct {
	Ipv6Address *string `pulumi:"ipv6Address"`
}





type AaaaRecordInput interface {
	pulumi.Input

	ToAaaaRecordOutput() AaaaRecordOutput
	ToAaaaRecordOutputWithContext(context.Context) AaaaRecordOutput
}

type AaaaRecordArgs struct {
	Ipv6Address pulumi.StringPtrInput `pulumi:"ipv6Address"`
}

func (AaaaRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AaaaRecord)(nil)).Elem()
}

func (i AaaaRecordArgs) ToAaaaRecordOutput() AaaaRecordOutput {
	return i.ToAaaaRecordOutputWithContext(context.Background())
}

func (i AaaaRecordArgs) ToAaaaRecordOutputWithContext(ctx context.Context) AaaaRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AaaaRecordOutput)
}





type AaaaRecordArrayInput interface {
	pulumi.Input

	ToAaaaRecordArrayOutput() AaaaRecordArrayOutput
	ToAaaaRecordArrayOutputWithContext(context.Context) AaaaRecordArrayOutput
}

type AaaaRecordArray []AaaaRecordInput

func (AaaaRecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AaaaRecord)(nil)).Elem()
}

func (i AaaaRecordArray) ToAaaaRecordArrayOutput() AaaaRecordArrayOutput {
	return i.ToAaaaRecordArrayOutputWithContext(context.Background())
}

func (i AaaaRecordArray) ToAaaaRecordArrayOutputWithContext(ctx context.Context) AaaaRecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AaaaRecordArrayOutput)
}

type AaaaRecordOutput struct{ *pulumi.OutputState }

func (AaaaRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AaaaRecord)(nil)).Elem()
}

func (o AaaaRecordOutput) ToAaaaRecordOutput() AaaaRecordOutput {
	return o
}

func (o AaaaRecordOutput) ToAaaaRecordOutputWithContext(ctx context.Context) AaaaRecordOutput {
	return o
}

func (o AaaaRecordOutput) Ipv6Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AaaaRecord) *string { return v.Ipv6Address }).(pulumi.StringPtrOutput)
}

type AaaaRecordArrayOutput struct{ *pulumi.OutputState }

func (AaaaRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AaaaRecord)(nil)).Elem()
}

func (o AaaaRecordArrayOutput) ToAaaaRecordArrayOutput() AaaaRecordArrayOutput {
	return o
}

func (o AaaaRecordArrayOutput) ToAaaaRecordArrayOutputWithContext(ctx context.Context) AaaaRecordArrayOutput {
	return o
}

func (o AaaaRecordArrayOutput) Index(i pulumi.IntInput) AaaaRecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AaaaRecord {
		return vs[0].([]AaaaRecord)[vs[1].(int)]
	}).(AaaaRecordOutput)
}

type AaaaRecordResponse struct {
	Ipv6Address *string `pulumi:"ipv6Address"`
}





type AaaaRecordResponseInput interface {
	pulumi.Input

	ToAaaaRecordResponseOutput() AaaaRecordResponseOutput
	ToAaaaRecordResponseOutputWithContext(context.Context) AaaaRecordResponseOutput
}

type AaaaRecordResponseArgs struct {
	Ipv6Address pulumi.StringPtrInput `pulumi:"ipv6Address"`
}

func (AaaaRecordResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AaaaRecordResponse)(nil)).Elem()
}

func (i AaaaRecordResponseArgs) ToAaaaRecordResponseOutput() AaaaRecordResponseOutput {
	return i.ToAaaaRecordResponseOutputWithContext(context.Background())
}

func (i AaaaRecordResponseArgs) ToAaaaRecordResponseOutputWithContext(ctx context.Context) AaaaRecordResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AaaaRecordResponseOutput)
}





type AaaaRecordResponseArrayInput interface {
	pulumi.Input

	ToAaaaRecordResponseArrayOutput() AaaaRecordResponseArrayOutput
	ToAaaaRecordResponseArrayOutputWithContext(context.Context) AaaaRecordResponseArrayOutput
}

type AaaaRecordResponseArray []AaaaRecordResponseInput

func (AaaaRecordResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AaaaRecordResponse)(nil)).Elem()
}

func (i AaaaRecordResponseArray) ToAaaaRecordResponseArrayOutput() AaaaRecordResponseArrayOutput {
	return i.ToAaaaRecordResponseArrayOutputWithContext(context.Background())
}

func (i AaaaRecordResponseArray) ToAaaaRecordResponseArrayOutputWithContext(ctx context.Context) AaaaRecordResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AaaaRecordResponseArrayOutput)
}

type AaaaRecordResponseOutput struct{ *pulumi.OutputState }

func (AaaaRecordResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AaaaRecordResponse)(nil)).Elem()
}

func (o AaaaRecordResponseOutput) ToAaaaRecordResponseOutput() AaaaRecordResponseOutput {
	return o
}

func (o AaaaRecordResponseOutput) ToAaaaRecordResponseOutputWithContext(ctx context.Context) AaaaRecordResponseOutput {
	return o
}

func (o AaaaRecordResponseOutput) Ipv6Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AaaaRecordResponse) *string { return v.Ipv6Address }).(pulumi.StringPtrOutput)
}

type AaaaRecordResponseArrayOutput struct{ *pulumi.OutputState }

func (AaaaRecordResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AaaaRecordResponse)(nil)).Elem()
}

func (o AaaaRecordResponseArrayOutput) ToAaaaRecordResponseArrayOutput() AaaaRecordResponseArrayOutput {
	return o
}

func (o AaaaRecordResponseArrayOutput) ToAaaaRecordResponseArrayOutputWithContext(ctx context.Context) AaaaRecordResponseArrayOutput {
	return o
}

func (o AaaaRecordResponseArrayOutput) Index(i pulumi.IntInput) AaaaRecordResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AaaaRecordResponse {
		return vs[0].([]AaaaRecordResponse)[vs[1].(int)]
	}).(AaaaRecordResponseOutput)
}

type CnameRecord struct {
	Cname *string `pulumi:"cname"`
}





type CnameRecordInput interface {
	pulumi.Input

	ToCnameRecordOutput() CnameRecordOutput
	ToCnameRecordOutputWithContext(context.Context) CnameRecordOutput
}

type CnameRecordArgs struct {
	Cname pulumi.StringPtrInput `pulumi:"cname"`
}

func (CnameRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CnameRecord)(nil)).Elem()
}

func (i CnameRecordArgs) ToCnameRecordOutput() CnameRecordOutput {
	return i.ToCnameRecordOutputWithContext(context.Background())
}

func (i CnameRecordArgs) ToCnameRecordOutputWithContext(ctx context.Context) CnameRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CnameRecordOutput)
}

func (i CnameRecordArgs) ToCnameRecordPtrOutput() CnameRecordPtrOutput {
	return i.ToCnameRecordPtrOutputWithContext(context.Background())
}

func (i CnameRecordArgs) ToCnameRecordPtrOutputWithContext(ctx context.Context) CnameRecordPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CnameRecordOutput).ToCnameRecordPtrOutputWithContext(ctx)
}









type CnameRecordPtrInput interface {
	pulumi.Input

	ToCnameRecordPtrOutput() CnameRecordPtrOutput
	ToCnameRecordPtrOutputWithContext(context.Context) CnameRecordPtrOutput
}

type cnameRecordPtrType CnameRecordArgs

func CnameRecordPtr(v *CnameRecordArgs) CnameRecordPtrInput {
	return (*cnameRecordPtrType)(v)
}

func (*cnameRecordPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CnameRecord)(nil)).Elem()
}

func (i *cnameRecordPtrType) ToCnameRecordPtrOutput() CnameRecordPtrOutput {
	return i.ToCnameRecordPtrOutputWithContext(context.Background())
}

func (i *cnameRecordPtrType) ToCnameRecordPtrOutputWithContext(ctx context.Context) CnameRecordPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CnameRecordPtrOutput)
}

type CnameRecordOutput struct{ *pulumi.OutputState }

func (CnameRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CnameRecord)(nil)).Elem()
}

func (o CnameRecordOutput) ToCnameRecordOutput() CnameRecordOutput {
	return o
}

func (o CnameRecordOutput) ToCnameRecordOutputWithContext(ctx context.Context) CnameRecordOutput {
	return o
}

func (o CnameRecordOutput) ToCnameRecordPtrOutput() CnameRecordPtrOutput {
	return o.ToCnameRecordPtrOutputWithContext(context.Background())
}

func (o CnameRecordOutput) ToCnameRecordPtrOutputWithContext(ctx context.Context) CnameRecordPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CnameRecord) *CnameRecord {
		return &v
	}).(CnameRecordPtrOutput)
}

func (o CnameRecordOutput) Cname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CnameRecord) *string { return v.Cname }).(pulumi.StringPtrOutput)
}

type CnameRecordPtrOutput struct{ *pulumi.OutputState }

func (CnameRecordPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CnameRecord)(nil)).Elem()
}

func (o CnameRecordPtrOutput) ToCnameRecordPtrOutput() CnameRecordPtrOutput {
	return o
}

func (o CnameRecordPtrOutput) ToCnameRecordPtrOutputWithContext(ctx context.Context) CnameRecordPtrOutput {
	return o
}

func (o CnameRecordPtrOutput) Elem() CnameRecordOutput {
	return o.ApplyT(func(v *CnameRecord) CnameRecord {
		if v != nil {
			return *v
		}
		var ret CnameRecord
		return ret
	}).(CnameRecordOutput)
}

func (o CnameRecordPtrOutput) Cname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CnameRecord) *string {
		if v == nil {
			return nil
		}
		return v.Cname
	}).(pulumi.StringPtrOutput)
}

type CnameRecordResponse struct {
	Cname *string `pulumi:"cname"`
}





type CnameRecordResponseInput interface {
	pulumi.Input

	ToCnameRecordResponseOutput() CnameRecordResponseOutput
	ToCnameRecordResponseOutputWithContext(context.Context) CnameRecordResponseOutput
}

type CnameRecordResponseArgs struct {
	Cname pulumi.StringPtrInput `pulumi:"cname"`
}

func (CnameRecordResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CnameRecordResponse)(nil)).Elem()
}

func (i CnameRecordResponseArgs) ToCnameRecordResponseOutput() CnameRecordResponseOutput {
	return i.ToCnameRecordResponseOutputWithContext(context.Background())
}

func (i CnameRecordResponseArgs) ToCnameRecordResponseOutputWithContext(ctx context.Context) CnameRecordResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CnameRecordResponseOutput)
}

func (i CnameRecordResponseArgs) ToCnameRecordResponsePtrOutput() CnameRecordResponsePtrOutput {
	return i.ToCnameRecordResponsePtrOutputWithContext(context.Background())
}

func (i CnameRecordResponseArgs) ToCnameRecordResponsePtrOutputWithContext(ctx context.Context) CnameRecordResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CnameRecordResponseOutput).ToCnameRecordResponsePtrOutputWithContext(ctx)
}









type CnameRecordResponsePtrInput interface {
	pulumi.Input

	ToCnameRecordResponsePtrOutput() CnameRecordResponsePtrOutput
	ToCnameRecordResponsePtrOutputWithContext(context.Context) CnameRecordResponsePtrOutput
}

type cnameRecordResponsePtrType CnameRecordResponseArgs

func CnameRecordResponsePtr(v *CnameRecordResponseArgs) CnameRecordResponsePtrInput {
	return (*cnameRecordResponsePtrType)(v)
}

func (*cnameRecordResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CnameRecordResponse)(nil)).Elem()
}

func (i *cnameRecordResponsePtrType) ToCnameRecordResponsePtrOutput() CnameRecordResponsePtrOutput {
	return i.ToCnameRecordResponsePtrOutputWithContext(context.Background())
}

func (i *cnameRecordResponsePtrType) ToCnameRecordResponsePtrOutputWithContext(ctx context.Context) CnameRecordResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CnameRecordResponsePtrOutput)
}

type CnameRecordResponseOutput struct{ *pulumi.OutputState }

func (CnameRecordResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CnameRecordResponse)(nil)).Elem()
}

func (o CnameRecordResponseOutput) ToCnameRecordResponseOutput() CnameRecordResponseOutput {
	return o
}

func (o CnameRecordResponseOutput) ToCnameRecordResponseOutputWithContext(ctx context.Context) CnameRecordResponseOutput {
	return o
}

func (o CnameRecordResponseOutput) ToCnameRecordResponsePtrOutput() CnameRecordResponsePtrOutput {
	return o.ToCnameRecordResponsePtrOutputWithContext(context.Background())
}

func (o CnameRecordResponseOutput) ToCnameRecordResponsePtrOutputWithContext(ctx context.Context) CnameRecordResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CnameRecordResponse) *CnameRecordResponse {
		return &v
	}).(CnameRecordResponsePtrOutput)
}

func (o CnameRecordResponseOutput) Cname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CnameRecordResponse) *string { return v.Cname }).(pulumi.StringPtrOutput)
}

type CnameRecordResponsePtrOutput struct{ *pulumi.OutputState }

func (CnameRecordResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CnameRecordResponse)(nil)).Elem()
}

func (o CnameRecordResponsePtrOutput) ToCnameRecordResponsePtrOutput() CnameRecordResponsePtrOutput {
	return o
}

func (o CnameRecordResponsePtrOutput) ToCnameRecordResponsePtrOutputWithContext(ctx context.Context) CnameRecordResponsePtrOutput {
	return o
}

func (o CnameRecordResponsePtrOutput) Elem() CnameRecordResponseOutput {
	return o.ApplyT(func(v *CnameRecordResponse) CnameRecordResponse {
		if v != nil {
			return *v
		}
		var ret CnameRecordResponse
		return ret
	}).(CnameRecordResponseOutput)
}

func (o CnameRecordResponsePtrOutput) Cname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CnameRecordResponse) *string {
		if v == nil {
			return nil
		}
		return v.Cname
	}).(pulumi.StringPtrOutput)
}

type MxRecord struct {
	Exchange   *string `pulumi:"exchange"`
	Preference *int    `pulumi:"preference"`
}





type MxRecordInput interface {
	pulumi.Input

	ToMxRecordOutput() MxRecordOutput
	ToMxRecordOutputWithContext(context.Context) MxRecordOutput
}

type MxRecordArgs struct {
	Exchange   pulumi.StringPtrInput `pulumi:"exchange"`
	Preference pulumi.IntPtrInput    `pulumi:"preference"`
}

func (MxRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MxRecord)(nil)).Elem()
}

func (i MxRecordArgs) ToMxRecordOutput() MxRecordOutput {
	return i.ToMxRecordOutputWithContext(context.Background())
}

func (i MxRecordArgs) ToMxRecordOutputWithContext(ctx context.Context) MxRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MxRecordOutput)
}





type MxRecordArrayInput interface {
	pulumi.Input

	ToMxRecordArrayOutput() MxRecordArrayOutput
	ToMxRecordArrayOutputWithContext(context.Context) MxRecordArrayOutput
}

type MxRecordArray []MxRecordInput

func (MxRecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MxRecord)(nil)).Elem()
}

func (i MxRecordArray) ToMxRecordArrayOutput() MxRecordArrayOutput {
	return i.ToMxRecordArrayOutputWithContext(context.Background())
}

func (i MxRecordArray) ToMxRecordArrayOutputWithContext(ctx context.Context) MxRecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MxRecordArrayOutput)
}

type MxRecordOutput struct{ *pulumi.OutputState }

func (MxRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MxRecord)(nil)).Elem()
}

func (o MxRecordOutput) ToMxRecordOutput() MxRecordOutput {
	return o
}

func (o MxRecordOutput) ToMxRecordOutputWithContext(ctx context.Context) MxRecordOutput {
	return o
}

func (o MxRecordOutput) Exchange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MxRecord) *string { return v.Exchange }).(pulumi.StringPtrOutput)
}

func (o MxRecordOutput) Preference() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MxRecord) *int { return v.Preference }).(pulumi.IntPtrOutput)
}

type MxRecordArrayOutput struct{ *pulumi.OutputState }

func (MxRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MxRecord)(nil)).Elem()
}

func (o MxRecordArrayOutput) ToMxRecordArrayOutput() MxRecordArrayOutput {
	return o
}

func (o MxRecordArrayOutput) ToMxRecordArrayOutputWithContext(ctx context.Context) MxRecordArrayOutput {
	return o
}

func (o MxRecordArrayOutput) Index(i pulumi.IntInput) MxRecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MxRecord {
		return vs[0].([]MxRecord)[vs[1].(int)]
	}).(MxRecordOutput)
}

type MxRecordResponse struct {
	Exchange   *string `pulumi:"exchange"`
	Preference *int    `pulumi:"preference"`
}





type MxRecordResponseInput interface {
	pulumi.Input

	ToMxRecordResponseOutput() MxRecordResponseOutput
	ToMxRecordResponseOutputWithContext(context.Context) MxRecordResponseOutput
}

type MxRecordResponseArgs struct {
	Exchange   pulumi.StringPtrInput `pulumi:"exchange"`
	Preference pulumi.IntPtrInput    `pulumi:"preference"`
}

func (MxRecordResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MxRecordResponse)(nil)).Elem()
}

func (i MxRecordResponseArgs) ToMxRecordResponseOutput() MxRecordResponseOutput {
	return i.ToMxRecordResponseOutputWithContext(context.Background())
}

func (i MxRecordResponseArgs) ToMxRecordResponseOutputWithContext(ctx context.Context) MxRecordResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MxRecordResponseOutput)
}





type MxRecordResponseArrayInput interface {
	pulumi.Input

	ToMxRecordResponseArrayOutput() MxRecordResponseArrayOutput
	ToMxRecordResponseArrayOutputWithContext(context.Context) MxRecordResponseArrayOutput
}

type MxRecordResponseArray []MxRecordResponseInput

func (MxRecordResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MxRecordResponse)(nil)).Elem()
}

func (i MxRecordResponseArray) ToMxRecordResponseArrayOutput() MxRecordResponseArrayOutput {
	return i.ToMxRecordResponseArrayOutputWithContext(context.Background())
}

func (i MxRecordResponseArray) ToMxRecordResponseArrayOutputWithContext(ctx context.Context) MxRecordResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MxRecordResponseArrayOutput)
}

type MxRecordResponseOutput struct{ *pulumi.OutputState }

func (MxRecordResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MxRecordResponse)(nil)).Elem()
}

func (o MxRecordResponseOutput) ToMxRecordResponseOutput() MxRecordResponseOutput {
	return o
}

func (o MxRecordResponseOutput) ToMxRecordResponseOutputWithContext(ctx context.Context) MxRecordResponseOutput {
	return o
}

func (o MxRecordResponseOutput) Exchange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MxRecordResponse) *string { return v.Exchange }).(pulumi.StringPtrOutput)
}

func (o MxRecordResponseOutput) Preference() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MxRecordResponse) *int { return v.Preference }).(pulumi.IntPtrOutput)
}

type MxRecordResponseArrayOutput struct{ *pulumi.OutputState }

func (MxRecordResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MxRecordResponse)(nil)).Elem()
}

func (o MxRecordResponseArrayOutput) ToMxRecordResponseArrayOutput() MxRecordResponseArrayOutput {
	return o
}

func (o MxRecordResponseArrayOutput) ToMxRecordResponseArrayOutputWithContext(ctx context.Context) MxRecordResponseArrayOutput {
	return o
}

func (o MxRecordResponseArrayOutput) Index(i pulumi.IntInput) MxRecordResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MxRecordResponse {
		return vs[0].([]MxRecordResponse)[vs[1].(int)]
	}).(MxRecordResponseOutput)
}

type PtrRecord struct {
	Ptrdname *string `pulumi:"ptrdname"`
}





type PtrRecordInput interface {
	pulumi.Input

	ToPtrRecordOutput() PtrRecordOutput
	ToPtrRecordOutputWithContext(context.Context) PtrRecordOutput
}

type PtrRecordArgs struct {
	Ptrdname pulumi.StringPtrInput `pulumi:"ptrdname"`
}

func (PtrRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PtrRecord)(nil)).Elem()
}

func (i PtrRecordArgs) ToPtrRecordOutput() PtrRecordOutput {
	return i.ToPtrRecordOutputWithContext(context.Background())
}

func (i PtrRecordArgs) ToPtrRecordOutputWithContext(ctx context.Context) PtrRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PtrRecordOutput)
}





type PtrRecordArrayInput interface {
	pulumi.Input

	ToPtrRecordArrayOutput() PtrRecordArrayOutput
	ToPtrRecordArrayOutputWithContext(context.Context) PtrRecordArrayOutput
}

type PtrRecordArray []PtrRecordInput

func (PtrRecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PtrRecord)(nil)).Elem()
}

func (i PtrRecordArray) ToPtrRecordArrayOutput() PtrRecordArrayOutput {
	return i.ToPtrRecordArrayOutputWithContext(context.Background())
}

func (i PtrRecordArray) ToPtrRecordArrayOutputWithContext(ctx context.Context) PtrRecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PtrRecordArrayOutput)
}

type PtrRecordOutput struct{ *pulumi.OutputState }

func (PtrRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PtrRecord)(nil)).Elem()
}

func (o PtrRecordOutput) ToPtrRecordOutput() PtrRecordOutput {
	return o
}

func (o PtrRecordOutput) ToPtrRecordOutputWithContext(ctx context.Context) PtrRecordOutput {
	return o
}

func (o PtrRecordOutput) Ptrdname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PtrRecord) *string { return v.Ptrdname }).(pulumi.StringPtrOutput)
}

type PtrRecordArrayOutput struct{ *pulumi.OutputState }

func (PtrRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PtrRecord)(nil)).Elem()
}

func (o PtrRecordArrayOutput) ToPtrRecordArrayOutput() PtrRecordArrayOutput {
	return o
}

func (o PtrRecordArrayOutput) ToPtrRecordArrayOutputWithContext(ctx context.Context) PtrRecordArrayOutput {
	return o
}

func (o PtrRecordArrayOutput) Index(i pulumi.IntInput) PtrRecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PtrRecord {
		return vs[0].([]PtrRecord)[vs[1].(int)]
	}).(PtrRecordOutput)
}

type PtrRecordResponse struct {
	Ptrdname *string `pulumi:"ptrdname"`
}





type PtrRecordResponseInput interface {
	pulumi.Input

	ToPtrRecordResponseOutput() PtrRecordResponseOutput
	ToPtrRecordResponseOutputWithContext(context.Context) PtrRecordResponseOutput
}

type PtrRecordResponseArgs struct {
	Ptrdname pulumi.StringPtrInput `pulumi:"ptrdname"`
}

func (PtrRecordResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PtrRecordResponse)(nil)).Elem()
}

func (i PtrRecordResponseArgs) ToPtrRecordResponseOutput() PtrRecordResponseOutput {
	return i.ToPtrRecordResponseOutputWithContext(context.Background())
}

func (i PtrRecordResponseArgs) ToPtrRecordResponseOutputWithContext(ctx context.Context) PtrRecordResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PtrRecordResponseOutput)
}





type PtrRecordResponseArrayInput interface {
	pulumi.Input

	ToPtrRecordResponseArrayOutput() PtrRecordResponseArrayOutput
	ToPtrRecordResponseArrayOutputWithContext(context.Context) PtrRecordResponseArrayOutput
}

type PtrRecordResponseArray []PtrRecordResponseInput

func (PtrRecordResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PtrRecordResponse)(nil)).Elem()
}

func (i PtrRecordResponseArray) ToPtrRecordResponseArrayOutput() PtrRecordResponseArrayOutput {
	return i.ToPtrRecordResponseArrayOutputWithContext(context.Background())
}

func (i PtrRecordResponseArray) ToPtrRecordResponseArrayOutputWithContext(ctx context.Context) PtrRecordResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PtrRecordResponseArrayOutput)
}

type PtrRecordResponseOutput struct{ *pulumi.OutputState }

func (PtrRecordResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PtrRecordResponse)(nil)).Elem()
}

func (o PtrRecordResponseOutput) ToPtrRecordResponseOutput() PtrRecordResponseOutput {
	return o
}

func (o PtrRecordResponseOutput) ToPtrRecordResponseOutputWithContext(ctx context.Context) PtrRecordResponseOutput {
	return o
}

func (o PtrRecordResponseOutput) Ptrdname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PtrRecordResponse) *string { return v.Ptrdname }).(pulumi.StringPtrOutput)
}

type PtrRecordResponseArrayOutput struct{ *pulumi.OutputState }

func (PtrRecordResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PtrRecordResponse)(nil)).Elem()
}

func (o PtrRecordResponseArrayOutput) ToPtrRecordResponseArrayOutput() PtrRecordResponseArrayOutput {
	return o
}

func (o PtrRecordResponseArrayOutput) ToPtrRecordResponseArrayOutputWithContext(ctx context.Context) PtrRecordResponseArrayOutput {
	return o
}

func (o PtrRecordResponseArrayOutput) Index(i pulumi.IntInput) PtrRecordResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PtrRecordResponse {
		return vs[0].([]PtrRecordResponse)[vs[1].(int)]
	}).(PtrRecordResponseOutput)
}

type SoaRecord struct {
	Email        *string  `pulumi:"email"`
	ExpireTime   *float64 `pulumi:"expireTime"`
	Host         *string  `pulumi:"host"`
	MinimumTtl   *float64 `pulumi:"minimumTtl"`
	RefreshTime  *float64 `pulumi:"refreshTime"`
	RetryTime    *float64 `pulumi:"retryTime"`
	SerialNumber *float64 `pulumi:"serialNumber"`
}





type SoaRecordInput interface {
	pulumi.Input

	ToSoaRecordOutput() SoaRecordOutput
	ToSoaRecordOutputWithContext(context.Context) SoaRecordOutput
}

type SoaRecordArgs struct {
	Email        pulumi.StringPtrInput  `pulumi:"email"`
	ExpireTime   pulumi.Float64PtrInput `pulumi:"expireTime"`
	Host         pulumi.StringPtrInput  `pulumi:"host"`
	MinimumTtl   pulumi.Float64PtrInput `pulumi:"minimumTtl"`
	RefreshTime  pulumi.Float64PtrInput `pulumi:"refreshTime"`
	RetryTime    pulumi.Float64PtrInput `pulumi:"retryTime"`
	SerialNumber pulumi.Float64PtrInput `pulumi:"serialNumber"`
}

func (SoaRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SoaRecord)(nil)).Elem()
}

func (i SoaRecordArgs) ToSoaRecordOutput() SoaRecordOutput {
	return i.ToSoaRecordOutputWithContext(context.Background())
}

func (i SoaRecordArgs) ToSoaRecordOutputWithContext(ctx context.Context) SoaRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SoaRecordOutput)
}

func (i SoaRecordArgs) ToSoaRecordPtrOutput() SoaRecordPtrOutput {
	return i.ToSoaRecordPtrOutputWithContext(context.Background())
}

func (i SoaRecordArgs) ToSoaRecordPtrOutputWithContext(ctx context.Context) SoaRecordPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SoaRecordOutput).ToSoaRecordPtrOutputWithContext(ctx)
}









type SoaRecordPtrInput interface {
	pulumi.Input

	ToSoaRecordPtrOutput() SoaRecordPtrOutput
	ToSoaRecordPtrOutputWithContext(context.Context) SoaRecordPtrOutput
}

type soaRecordPtrType SoaRecordArgs

func SoaRecordPtr(v *SoaRecordArgs) SoaRecordPtrInput {
	return (*soaRecordPtrType)(v)
}

func (*soaRecordPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SoaRecord)(nil)).Elem()
}

func (i *soaRecordPtrType) ToSoaRecordPtrOutput() SoaRecordPtrOutput {
	return i.ToSoaRecordPtrOutputWithContext(context.Background())
}

func (i *soaRecordPtrType) ToSoaRecordPtrOutputWithContext(ctx context.Context) SoaRecordPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SoaRecordPtrOutput)
}

type SoaRecordOutput struct{ *pulumi.OutputState }

func (SoaRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SoaRecord)(nil)).Elem()
}

func (o SoaRecordOutput) ToSoaRecordOutput() SoaRecordOutput {
	return o
}

func (o SoaRecordOutput) ToSoaRecordOutputWithContext(ctx context.Context) SoaRecordOutput {
	return o
}

func (o SoaRecordOutput) ToSoaRecordPtrOutput() SoaRecordPtrOutput {
	return o.ToSoaRecordPtrOutputWithContext(context.Background())
}

func (o SoaRecordOutput) ToSoaRecordPtrOutputWithContext(ctx context.Context) SoaRecordPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SoaRecord) *SoaRecord {
		return &v
	}).(SoaRecordPtrOutput)
}

func (o SoaRecordOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SoaRecord) *string { return v.Email }).(pulumi.StringPtrOutput)
}

func (o SoaRecordOutput) ExpireTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SoaRecord) *float64 { return v.ExpireTime }).(pulumi.Float64PtrOutput)
}

func (o SoaRecordOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SoaRecord) *string { return v.Host }).(pulumi.StringPtrOutput)
}

func (o SoaRecordOutput) MinimumTtl() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SoaRecord) *float64 { return v.MinimumTtl }).(pulumi.Float64PtrOutput)
}

func (o SoaRecordOutput) RefreshTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SoaRecord) *float64 { return v.RefreshTime }).(pulumi.Float64PtrOutput)
}

func (o SoaRecordOutput) RetryTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SoaRecord) *float64 { return v.RetryTime }).(pulumi.Float64PtrOutput)
}

func (o SoaRecordOutput) SerialNumber() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SoaRecord) *float64 { return v.SerialNumber }).(pulumi.Float64PtrOutput)
}

type SoaRecordPtrOutput struct{ *pulumi.OutputState }

func (SoaRecordPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SoaRecord)(nil)).Elem()
}

func (o SoaRecordPtrOutput) ToSoaRecordPtrOutput() SoaRecordPtrOutput {
	return o
}

func (o SoaRecordPtrOutput) ToSoaRecordPtrOutputWithContext(ctx context.Context) SoaRecordPtrOutput {
	return o
}

func (o SoaRecordPtrOutput) Elem() SoaRecordOutput {
	return o.ApplyT(func(v *SoaRecord) SoaRecord {
		if v != nil {
			return *v
		}
		var ret SoaRecord
		return ret
	}).(SoaRecordOutput)
}

func (o SoaRecordPtrOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SoaRecord) *string {
		if v == nil {
			return nil
		}
		return v.Email
	}).(pulumi.StringPtrOutput)
}

func (o SoaRecordPtrOutput) ExpireTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SoaRecord) *float64 {
		if v == nil {
			return nil
		}
		return v.ExpireTime
	}).(pulumi.Float64PtrOutput)
}

func (o SoaRecordPtrOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SoaRecord) *string {
		if v == nil {
			return nil
		}
		return v.Host
	}).(pulumi.StringPtrOutput)
}

func (o SoaRecordPtrOutput) MinimumTtl() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SoaRecord) *float64 {
		if v == nil {
			return nil
		}
		return v.MinimumTtl
	}).(pulumi.Float64PtrOutput)
}

func (o SoaRecordPtrOutput) RefreshTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SoaRecord) *float64 {
		if v == nil {
			return nil
		}
		return v.RefreshTime
	}).(pulumi.Float64PtrOutput)
}

func (o SoaRecordPtrOutput) RetryTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SoaRecord) *float64 {
		if v == nil {
			return nil
		}
		return v.RetryTime
	}).(pulumi.Float64PtrOutput)
}

func (o SoaRecordPtrOutput) SerialNumber() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SoaRecord) *float64 {
		if v == nil {
			return nil
		}
		return v.SerialNumber
	}).(pulumi.Float64PtrOutput)
}

type SoaRecordResponse struct {
	Email        *string  `pulumi:"email"`
	ExpireTime   *float64 `pulumi:"expireTime"`
	Host         *string  `pulumi:"host"`
	MinimumTtl   *float64 `pulumi:"minimumTtl"`
	RefreshTime  *float64 `pulumi:"refreshTime"`
	RetryTime    *float64 `pulumi:"retryTime"`
	SerialNumber *float64 `pulumi:"serialNumber"`
}





type SoaRecordResponseInput interface {
	pulumi.Input

	ToSoaRecordResponseOutput() SoaRecordResponseOutput
	ToSoaRecordResponseOutputWithContext(context.Context) SoaRecordResponseOutput
}

type SoaRecordResponseArgs struct {
	Email        pulumi.StringPtrInput  `pulumi:"email"`
	ExpireTime   pulumi.Float64PtrInput `pulumi:"expireTime"`
	Host         pulumi.StringPtrInput  `pulumi:"host"`
	MinimumTtl   pulumi.Float64PtrInput `pulumi:"minimumTtl"`
	RefreshTime  pulumi.Float64PtrInput `pulumi:"refreshTime"`
	RetryTime    pulumi.Float64PtrInput `pulumi:"retryTime"`
	SerialNumber pulumi.Float64PtrInput `pulumi:"serialNumber"`
}

func (SoaRecordResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SoaRecordResponse)(nil)).Elem()
}

func (i SoaRecordResponseArgs) ToSoaRecordResponseOutput() SoaRecordResponseOutput {
	return i.ToSoaRecordResponseOutputWithContext(context.Background())
}

func (i SoaRecordResponseArgs) ToSoaRecordResponseOutputWithContext(ctx context.Context) SoaRecordResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SoaRecordResponseOutput)
}

func (i SoaRecordResponseArgs) ToSoaRecordResponsePtrOutput() SoaRecordResponsePtrOutput {
	return i.ToSoaRecordResponsePtrOutputWithContext(context.Background())
}

func (i SoaRecordResponseArgs) ToSoaRecordResponsePtrOutputWithContext(ctx context.Context) SoaRecordResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SoaRecordResponseOutput).ToSoaRecordResponsePtrOutputWithContext(ctx)
}









type SoaRecordResponsePtrInput interface {
	pulumi.Input

	ToSoaRecordResponsePtrOutput() SoaRecordResponsePtrOutput
	ToSoaRecordResponsePtrOutputWithContext(context.Context) SoaRecordResponsePtrOutput
}

type soaRecordResponsePtrType SoaRecordResponseArgs

func SoaRecordResponsePtr(v *SoaRecordResponseArgs) SoaRecordResponsePtrInput {
	return (*soaRecordResponsePtrType)(v)
}

func (*soaRecordResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SoaRecordResponse)(nil)).Elem()
}

func (i *soaRecordResponsePtrType) ToSoaRecordResponsePtrOutput() SoaRecordResponsePtrOutput {
	return i.ToSoaRecordResponsePtrOutputWithContext(context.Background())
}

func (i *soaRecordResponsePtrType) ToSoaRecordResponsePtrOutputWithContext(ctx context.Context) SoaRecordResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SoaRecordResponsePtrOutput)
}

type SoaRecordResponseOutput struct{ *pulumi.OutputState }

func (SoaRecordResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SoaRecordResponse)(nil)).Elem()
}

func (o SoaRecordResponseOutput) ToSoaRecordResponseOutput() SoaRecordResponseOutput {
	return o
}

func (o SoaRecordResponseOutput) ToSoaRecordResponseOutputWithContext(ctx context.Context) SoaRecordResponseOutput {
	return o
}

func (o SoaRecordResponseOutput) ToSoaRecordResponsePtrOutput() SoaRecordResponsePtrOutput {
	return o.ToSoaRecordResponsePtrOutputWithContext(context.Background())
}

func (o SoaRecordResponseOutput) ToSoaRecordResponsePtrOutputWithContext(ctx context.Context) SoaRecordResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SoaRecordResponse) *SoaRecordResponse {
		return &v
	}).(SoaRecordResponsePtrOutput)
}

func (o SoaRecordResponseOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SoaRecordResponse) *string { return v.Email }).(pulumi.StringPtrOutput)
}

func (o SoaRecordResponseOutput) ExpireTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SoaRecordResponse) *float64 { return v.ExpireTime }).(pulumi.Float64PtrOutput)
}

func (o SoaRecordResponseOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SoaRecordResponse) *string { return v.Host }).(pulumi.StringPtrOutput)
}

func (o SoaRecordResponseOutput) MinimumTtl() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SoaRecordResponse) *float64 { return v.MinimumTtl }).(pulumi.Float64PtrOutput)
}

func (o SoaRecordResponseOutput) RefreshTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SoaRecordResponse) *float64 { return v.RefreshTime }).(pulumi.Float64PtrOutput)
}

func (o SoaRecordResponseOutput) RetryTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SoaRecordResponse) *float64 { return v.RetryTime }).(pulumi.Float64PtrOutput)
}

func (o SoaRecordResponseOutput) SerialNumber() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SoaRecordResponse) *float64 { return v.SerialNumber }).(pulumi.Float64PtrOutput)
}

type SoaRecordResponsePtrOutput struct{ *pulumi.OutputState }

func (SoaRecordResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SoaRecordResponse)(nil)).Elem()
}

func (o SoaRecordResponsePtrOutput) ToSoaRecordResponsePtrOutput() SoaRecordResponsePtrOutput {
	return o
}

func (o SoaRecordResponsePtrOutput) ToSoaRecordResponsePtrOutputWithContext(ctx context.Context) SoaRecordResponsePtrOutput {
	return o
}

func (o SoaRecordResponsePtrOutput) Elem() SoaRecordResponseOutput {
	return o.ApplyT(func(v *SoaRecordResponse) SoaRecordResponse {
		if v != nil {
			return *v
		}
		var ret SoaRecordResponse
		return ret
	}).(SoaRecordResponseOutput)
}

func (o SoaRecordResponsePtrOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SoaRecordResponse) *string {
		if v == nil {
			return nil
		}
		return v.Email
	}).(pulumi.StringPtrOutput)
}

func (o SoaRecordResponsePtrOutput) ExpireTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SoaRecordResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.ExpireTime
	}).(pulumi.Float64PtrOutput)
}

func (o SoaRecordResponsePtrOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SoaRecordResponse) *string {
		if v == nil {
			return nil
		}
		return v.Host
	}).(pulumi.StringPtrOutput)
}

func (o SoaRecordResponsePtrOutput) MinimumTtl() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SoaRecordResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.MinimumTtl
	}).(pulumi.Float64PtrOutput)
}

func (o SoaRecordResponsePtrOutput) RefreshTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SoaRecordResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.RefreshTime
	}).(pulumi.Float64PtrOutput)
}

func (o SoaRecordResponsePtrOutput) RetryTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SoaRecordResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.RetryTime
	}).(pulumi.Float64PtrOutput)
}

func (o SoaRecordResponsePtrOutput) SerialNumber() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SoaRecordResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.SerialNumber
	}).(pulumi.Float64PtrOutput)
}

type SrvRecord struct {
	Port     *int    `pulumi:"port"`
	Priority *int    `pulumi:"priority"`
	Target   *string `pulumi:"target"`
	Weight   *int    `pulumi:"weight"`
}





type SrvRecordInput interface {
	pulumi.Input

	ToSrvRecordOutput() SrvRecordOutput
	ToSrvRecordOutputWithContext(context.Context) SrvRecordOutput
}

type SrvRecordArgs struct {
	Port     pulumi.IntPtrInput    `pulumi:"port"`
	Priority pulumi.IntPtrInput    `pulumi:"priority"`
	Target   pulumi.StringPtrInput `pulumi:"target"`
	Weight   pulumi.IntPtrInput    `pulumi:"weight"`
}

func (SrvRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SrvRecord)(nil)).Elem()
}

func (i SrvRecordArgs) ToSrvRecordOutput() SrvRecordOutput {
	return i.ToSrvRecordOutputWithContext(context.Background())
}

func (i SrvRecordArgs) ToSrvRecordOutputWithContext(ctx context.Context) SrvRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SrvRecordOutput)
}





type SrvRecordArrayInput interface {
	pulumi.Input

	ToSrvRecordArrayOutput() SrvRecordArrayOutput
	ToSrvRecordArrayOutputWithContext(context.Context) SrvRecordArrayOutput
}

type SrvRecordArray []SrvRecordInput

func (SrvRecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SrvRecord)(nil)).Elem()
}

func (i SrvRecordArray) ToSrvRecordArrayOutput() SrvRecordArrayOutput {
	return i.ToSrvRecordArrayOutputWithContext(context.Background())
}

func (i SrvRecordArray) ToSrvRecordArrayOutputWithContext(ctx context.Context) SrvRecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SrvRecordArrayOutput)
}

type SrvRecordOutput struct{ *pulumi.OutputState }

func (SrvRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SrvRecord)(nil)).Elem()
}

func (o SrvRecordOutput) ToSrvRecordOutput() SrvRecordOutput {
	return o
}

func (o SrvRecordOutput) ToSrvRecordOutputWithContext(ctx context.Context) SrvRecordOutput {
	return o
}

func (o SrvRecordOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SrvRecord) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o SrvRecordOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SrvRecord) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o SrvRecordOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SrvRecord) *string { return v.Target }).(pulumi.StringPtrOutput)
}

func (o SrvRecordOutput) Weight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SrvRecord) *int { return v.Weight }).(pulumi.IntPtrOutput)
}

type SrvRecordArrayOutput struct{ *pulumi.OutputState }

func (SrvRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SrvRecord)(nil)).Elem()
}

func (o SrvRecordArrayOutput) ToSrvRecordArrayOutput() SrvRecordArrayOutput {
	return o
}

func (o SrvRecordArrayOutput) ToSrvRecordArrayOutputWithContext(ctx context.Context) SrvRecordArrayOutput {
	return o
}

func (o SrvRecordArrayOutput) Index(i pulumi.IntInput) SrvRecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SrvRecord {
		return vs[0].([]SrvRecord)[vs[1].(int)]
	}).(SrvRecordOutput)
}

type SrvRecordResponse struct {
	Port     *int    `pulumi:"port"`
	Priority *int    `pulumi:"priority"`
	Target   *string `pulumi:"target"`
	Weight   *int    `pulumi:"weight"`
}





type SrvRecordResponseInput interface {
	pulumi.Input

	ToSrvRecordResponseOutput() SrvRecordResponseOutput
	ToSrvRecordResponseOutputWithContext(context.Context) SrvRecordResponseOutput
}

type SrvRecordResponseArgs struct {
	Port     pulumi.IntPtrInput    `pulumi:"port"`
	Priority pulumi.IntPtrInput    `pulumi:"priority"`
	Target   pulumi.StringPtrInput `pulumi:"target"`
	Weight   pulumi.IntPtrInput    `pulumi:"weight"`
}

func (SrvRecordResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SrvRecordResponse)(nil)).Elem()
}

func (i SrvRecordResponseArgs) ToSrvRecordResponseOutput() SrvRecordResponseOutput {
	return i.ToSrvRecordResponseOutputWithContext(context.Background())
}

func (i SrvRecordResponseArgs) ToSrvRecordResponseOutputWithContext(ctx context.Context) SrvRecordResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SrvRecordResponseOutput)
}





type SrvRecordResponseArrayInput interface {
	pulumi.Input

	ToSrvRecordResponseArrayOutput() SrvRecordResponseArrayOutput
	ToSrvRecordResponseArrayOutputWithContext(context.Context) SrvRecordResponseArrayOutput
}

type SrvRecordResponseArray []SrvRecordResponseInput

func (SrvRecordResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SrvRecordResponse)(nil)).Elem()
}

func (i SrvRecordResponseArray) ToSrvRecordResponseArrayOutput() SrvRecordResponseArrayOutput {
	return i.ToSrvRecordResponseArrayOutputWithContext(context.Background())
}

func (i SrvRecordResponseArray) ToSrvRecordResponseArrayOutputWithContext(ctx context.Context) SrvRecordResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SrvRecordResponseArrayOutput)
}

type SrvRecordResponseOutput struct{ *pulumi.OutputState }

func (SrvRecordResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SrvRecordResponse)(nil)).Elem()
}

func (o SrvRecordResponseOutput) ToSrvRecordResponseOutput() SrvRecordResponseOutput {
	return o
}

func (o SrvRecordResponseOutput) ToSrvRecordResponseOutputWithContext(ctx context.Context) SrvRecordResponseOutput {
	return o
}

func (o SrvRecordResponseOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SrvRecordResponse) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o SrvRecordResponseOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SrvRecordResponse) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o SrvRecordResponseOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SrvRecordResponse) *string { return v.Target }).(pulumi.StringPtrOutput)
}

func (o SrvRecordResponseOutput) Weight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SrvRecordResponse) *int { return v.Weight }).(pulumi.IntPtrOutput)
}

type SrvRecordResponseArrayOutput struct{ *pulumi.OutputState }

func (SrvRecordResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SrvRecordResponse)(nil)).Elem()
}

func (o SrvRecordResponseArrayOutput) ToSrvRecordResponseArrayOutput() SrvRecordResponseArrayOutput {
	return o
}

func (o SrvRecordResponseArrayOutput) ToSrvRecordResponseArrayOutputWithContext(ctx context.Context) SrvRecordResponseArrayOutput {
	return o
}

func (o SrvRecordResponseArrayOutput) Index(i pulumi.IntInput) SrvRecordResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SrvRecordResponse {
		return vs[0].([]SrvRecordResponse)[vs[1].(int)]
	}).(SrvRecordResponseOutput)
}

type SubResource struct {
	Id *string `pulumi:"id"`
}





type SubResourceInput interface {
	pulumi.Input

	ToSubResourceOutput() SubResourceOutput
	ToSubResourceOutputWithContext(context.Context) SubResourceOutput
}

type SubResourceArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (SubResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResource)(nil)).Elem()
}

func (i SubResourceArgs) ToSubResourceOutput() SubResourceOutput {
	return i.ToSubResourceOutputWithContext(context.Background())
}

func (i SubResourceArgs) ToSubResourceOutputWithContext(ctx context.Context) SubResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceOutput)
}

func (i SubResourceArgs) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return i.ToSubResourcePtrOutputWithContext(context.Background())
}

func (i SubResourceArgs) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceOutput).ToSubResourcePtrOutputWithContext(ctx)
}









type SubResourcePtrInput interface {
	pulumi.Input

	ToSubResourcePtrOutput() SubResourcePtrOutput
	ToSubResourcePtrOutputWithContext(context.Context) SubResourcePtrOutput
}

type subResourcePtrType SubResourceArgs

func SubResourcePtr(v *SubResourceArgs) SubResourcePtrInput {
	return (*subResourcePtrType)(v)
}

func (*subResourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SubResource)(nil)).Elem()
}

func (i *subResourcePtrType) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return i.ToSubResourcePtrOutputWithContext(context.Background())
}

func (i *subResourcePtrType) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourcePtrOutput)
}

type SubResourceOutput struct{ *pulumi.OutputState }

func (SubResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResource)(nil)).Elem()
}

func (o SubResourceOutput) ToSubResourceOutput() SubResourceOutput {
	return o
}

func (o SubResourceOutput) ToSubResourceOutputWithContext(ctx context.Context) SubResourceOutput {
	return o
}

func (o SubResourceOutput) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return o.ToSubResourcePtrOutputWithContext(context.Background())
}

func (o SubResourceOutput) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SubResource) *SubResource {
		return &v
	}).(SubResourcePtrOutput)
}

func (o SubResourceOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubResource) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type SubResourcePtrOutput struct{ *pulumi.OutputState }

func (SubResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubResource)(nil)).Elem()
}

func (o SubResourcePtrOutput) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return o
}

func (o SubResourcePtrOutput) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return o
}

func (o SubResourcePtrOutput) Elem() SubResourceOutput {
	return o.ApplyT(func(v *SubResource) SubResource {
		if v != nil {
			return *v
		}
		var ret SubResource
		return ret
	}).(SubResourceOutput)
}

func (o SubResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubResource) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type SubResourceResponse struct {
	Id *string `pulumi:"id"`
}





type SubResourceResponseInput interface {
	pulumi.Input

	ToSubResourceResponseOutput() SubResourceResponseOutput
	ToSubResourceResponseOutputWithContext(context.Context) SubResourceResponseOutput
}

type SubResourceResponseArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (SubResourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResourceResponse)(nil)).Elem()
}

func (i SubResourceResponseArgs) ToSubResourceResponseOutput() SubResourceResponseOutput {
	return i.ToSubResourceResponseOutputWithContext(context.Background())
}

func (i SubResourceResponseArgs) ToSubResourceResponseOutputWithContext(ctx context.Context) SubResourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceResponseOutput)
}

func (i SubResourceResponseArgs) ToSubResourceResponsePtrOutput() SubResourceResponsePtrOutput {
	return i.ToSubResourceResponsePtrOutputWithContext(context.Background())
}

func (i SubResourceResponseArgs) ToSubResourceResponsePtrOutputWithContext(ctx context.Context) SubResourceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceResponseOutput).ToSubResourceResponsePtrOutputWithContext(ctx)
}









type SubResourceResponsePtrInput interface {
	pulumi.Input

	ToSubResourceResponsePtrOutput() SubResourceResponsePtrOutput
	ToSubResourceResponsePtrOutputWithContext(context.Context) SubResourceResponsePtrOutput
}

type subResourceResponsePtrType SubResourceResponseArgs

func SubResourceResponsePtr(v *SubResourceResponseArgs) SubResourceResponsePtrInput {
	return (*subResourceResponsePtrType)(v)
}

func (*subResourceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SubResourceResponse)(nil)).Elem()
}

func (i *subResourceResponsePtrType) ToSubResourceResponsePtrOutput() SubResourceResponsePtrOutput {
	return i.ToSubResourceResponsePtrOutputWithContext(context.Background())
}

func (i *subResourceResponsePtrType) ToSubResourceResponsePtrOutputWithContext(ctx context.Context) SubResourceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceResponsePtrOutput)
}

type SubResourceResponseOutput struct{ *pulumi.OutputState }

func (SubResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResourceResponse)(nil)).Elem()
}

func (o SubResourceResponseOutput) ToSubResourceResponseOutput() SubResourceResponseOutput {
	return o
}

func (o SubResourceResponseOutput) ToSubResourceResponseOutputWithContext(ctx context.Context) SubResourceResponseOutput {
	return o
}

func (o SubResourceResponseOutput) ToSubResourceResponsePtrOutput() SubResourceResponsePtrOutput {
	return o.ToSubResourceResponsePtrOutputWithContext(context.Background())
}

func (o SubResourceResponseOutput) ToSubResourceResponsePtrOutputWithContext(ctx context.Context) SubResourceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SubResourceResponse) *SubResourceResponse {
		return &v
	}).(SubResourceResponsePtrOutput)
}

func (o SubResourceResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubResourceResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type SubResourceResponsePtrOutput struct{ *pulumi.OutputState }

func (SubResourceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubResourceResponse)(nil)).Elem()
}

func (o SubResourceResponsePtrOutput) ToSubResourceResponsePtrOutput() SubResourceResponsePtrOutput {
	return o
}

func (o SubResourceResponsePtrOutput) ToSubResourceResponsePtrOutputWithContext(ctx context.Context) SubResourceResponsePtrOutput {
	return o
}

func (o SubResourceResponsePtrOutput) Elem() SubResourceResponseOutput {
	return o.ApplyT(func(v *SubResourceResponse) SubResourceResponse {
		if v != nil {
			return *v
		}
		var ret SubResourceResponse
		return ret
	}).(SubResourceResponseOutput)
}

func (o SubResourceResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubResourceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type TxtRecord struct {
	Value []string `pulumi:"value"`
}





type TxtRecordInput interface {
	pulumi.Input

	ToTxtRecordOutput() TxtRecordOutput
	ToTxtRecordOutputWithContext(context.Context) TxtRecordOutput
}

type TxtRecordArgs struct {
	Value pulumi.StringArrayInput `pulumi:"value"`
}

func (TxtRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TxtRecord)(nil)).Elem()
}

func (i TxtRecordArgs) ToTxtRecordOutput() TxtRecordOutput {
	return i.ToTxtRecordOutputWithContext(context.Background())
}

func (i TxtRecordArgs) ToTxtRecordOutputWithContext(ctx context.Context) TxtRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TxtRecordOutput)
}





type TxtRecordArrayInput interface {
	pulumi.Input

	ToTxtRecordArrayOutput() TxtRecordArrayOutput
	ToTxtRecordArrayOutputWithContext(context.Context) TxtRecordArrayOutput
}

type TxtRecordArray []TxtRecordInput

func (TxtRecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TxtRecord)(nil)).Elem()
}

func (i TxtRecordArray) ToTxtRecordArrayOutput() TxtRecordArrayOutput {
	return i.ToTxtRecordArrayOutputWithContext(context.Background())
}

func (i TxtRecordArray) ToTxtRecordArrayOutputWithContext(ctx context.Context) TxtRecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TxtRecordArrayOutput)
}

type TxtRecordOutput struct{ *pulumi.OutputState }

func (TxtRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TxtRecord)(nil)).Elem()
}

func (o TxtRecordOutput) ToTxtRecordOutput() TxtRecordOutput {
	return o
}

func (o TxtRecordOutput) ToTxtRecordOutputWithContext(ctx context.Context) TxtRecordOutput {
	return o
}

func (o TxtRecordOutput) Value() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TxtRecord) []string { return v.Value }).(pulumi.StringArrayOutput)
}

type TxtRecordArrayOutput struct{ *pulumi.OutputState }

func (TxtRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TxtRecord)(nil)).Elem()
}

func (o TxtRecordArrayOutput) ToTxtRecordArrayOutput() TxtRecordArrayOutput {
	return o
}

func (o TxtRecordArrayOutput) ToTxtRecordArrayOutputWithContext(ctx context.Context) TxtRecordArrayOutput {
	return o
}

func (o TxtRecordArrayOutput) Index(i pulumi.IntInput) TxtRecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TxtRecord {
		return vs[0].([]TxtRecord)[vs[1].(int)]
	}).(TxtRecordOutput)
}

type TxtRecordResponse struct {
	Value []string `pulumi:"value"`
}





type TxtRecordResponseInput interface {
	pulumi.Input

	ToTxtRecordResponseOutput() TxtRecordResponseOutput
	ToTxtRecordResponseOutputWithContext(context.Context) TxtRecordResponseOutput
}

type TxtRecordResponseArgs struct {
	Value pulumi.StringArrayInput `pulumi:"value"`
}

func (TxtRecordResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TxtRecordResponse)(nil)).Elem()
}

func (i TxtRecordResponseArgs) ToTxtRecordResponseOutput() TxtRecordResponseOutput {
	return i.ToTxtRecordResponseOutputWithContext(context.Background())
}

func (i TxtRecordResponseArgs) ToTxtRecordResponseOutputWithContext(ctx context.Context) TxtRecordResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TxtRecordResponseOutput)
}





type TxtRecordResponseArrayInput interface {
	pulumi.Input

	ToTxtRecordResponseArrayOutput() TxtRecordResponseArrayOutput
	ToTxtRecordResponseArrayOutputWithContext(context.Context) TxtRecordResponseArrayOutput
}

type TxtRecordResponseArray []TxtRecordResponseInput

func (TxtRecordResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TxtRecordResponse)(nil)).Elem()
}

func (i TxtRecordResponseArray) ToTxtRecordResponseArrayOutput() TxtRecordResponseArrayOutput {
	return i.ToTxtRecordResponseArrayOutputWithContext(context.Background())
}

func (i TxtRecordResponseArray) ToTxtRecordResponseArrayOutputWithContext(ctx context.Context) TxtRecordResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TxtRecordResponseArrayOutput)
}

type TxtRecordResponseOutput struct{ *pulumi.OutputState }

func (TxtRecordResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TxtRecordResponse)(nil)).Elem()
}

func (o TxtRecordResponseOutput) ToTxtRecordResponseOutput() TxtRecordResponseOutput {
	return o
}

func (o TxtRecordResponseOutput) ToTxtRecordResponseOutputWithContext(ctx context.Context) TxtRecordResponseOutput {
	return o
}

func (o TxtRecordResponseOutput) Value() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TxtRecordResponse) []string { return v.Value }).(pulumi.StringArrayOutput)
}

type TxtRecordResponseArrayOutput struct{ *pulumi.OutputState }

func (TxtRecordResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TxtRecordResponse)(nil)).Elem()
}

func (o TxtRecordResponseArrayOutput) ToTxtRecordResponseArrayOutput() TxtRecordResponseArrayOutput {
	return o
}

func (o TxtRecordResponseArrayOutput) ToTxtRecordResponseArrayOutputWithContext(ctx context.Context) TxtRecordResponseArrayOutput {
	return o
}

func (o TxtRecordResponseArrayOutput) Index(i pulumi.IntInput) TxtRecordResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TxtRecordResponse {
		return vs[0].([]TxtRecordResponse)[vs[1].(int)]
	}).(TxtRecordResponseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ARecordInput)(nil)).Elem(), ARecordArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ARecordArrayInput)(nil)).Elem(), ARecordArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ARecordResponseInput)(nil)).Elem(), ARecordResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ARecordResponseArrayInput)(nil)).Elem(), ARecordResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AaaaRecordInput)(nil)).Elem(), AaaaRecordArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AaaaRecordArrayInput)(nil)).Elem(), AaaaRecordArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AaaaRecordResponseInput)(nil)).Elem(), AaaaRecordResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AaaaRecordResponseArrayInput)(nil)).Elem(), AaaaRecordResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CnameRecordInput)(nil)).Elem(), CnameRecordArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CnameRecordPtrInput)(nil)).Elem(), CnameRecordArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CnameRecordResponseInput)(nil)).Elem(), CnameRecordResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CnameRecordResponsePtrInput)(nil)).Elem(), CnameRecordResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MxRecordInput)(nil)).Elem(), MxRecordArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MxRecordArrayInput)(nil)).Elem(), MxRecordArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MxRecordResponseInput)(nil)).Elem(), MxRecordResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MxRecordResponseArrayInput)(nil)).Elem(), MxRecordResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PtrRecordInput)(nil)).Elem(), PtrRecordArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PtrRecordArrayInput)(nil)).Elem(), PtrRecordArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PtrRecordResponseInput)(nil)).Elem(), PtrRecordResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PtrRecordResponseArrayInput)(nil)).Elem(), PtrRecordResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SoaRecordInput)(nil)).Elem(), SoaRecordArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SoaRecordPtrInput)(nil)).Elem(), SoaRecordArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SoaRecordResponseInput)(nil)).Elem(), SoaRecordResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SoaRecordResponsePtrInput)(nil)).Elem(), SoaRecordResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SrvRecordInput)(nil)).Elem(), SrvRecordArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SrvRecordArrayInput)(nil)).Elem(), SrvRecordArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SrvRecordResponseInput)(nil)).Elem(), SrvRecordResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SrvRecordResponseArrayInput)(nil)).Elem(), SrvRecordResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubResourceInput)(nil)).Elem(), SubResourceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubResourcePtrInput)(nil)).Elem(), SubResourceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubResourceResponseInput)(nil)).Elem(), SubResourceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubResourceResponsePtrInput)(nil)).Elem(), SubResourceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TxtRecordInput)(nil)).Elem(), TxtRecordArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TxtRecordArrayInput)(nil)).Elem(), TxtRecordArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TxtRecordResponseInput)(nil)).Elem(), TxtRecordResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TxtRecordResponseArrayInput)(nil)).Elem(), TxtRecordResponseArray{})
	pulumi.RegisterOutputType(ARecordOutput{})
	pulumi.RegisterOutputType(ARecordArrayOutput{})
	pulumi.RegisterOutputType(ARecordResponseOutput{})
	pulumi.RegisterOutputType(ARecordResponseArrayOutput{})
	pulumi.RegisterOutputType(AaaaRecordOutput{})
	pulumi.RegisterOutputType(AaaaRecordArrayOutput{})
	pulumi.RegisterOutputType(AaaaRecordResponseOutput{})
	pulumi.RegisterOutputType(AaaaRecordResponseArrayOutput{})
	pulumi.RegisterOutputType(CnameRecordOutput{})
	pulumi.RegisterOutputType(CnameRecordPtrOutput{})
	pulumi.RegisterOutputType(CnameRecordResponseOutput{})
	pulumi.RegisterOutputType(CnameRecordResponsePtrOutput{})
	pulumi.RegisterOutputType(MxRecordOutput{})
	pulumi.RegisterOutputType(MxRecordArrayOutput{})
	pulumi.RegisterOutputType(MxRecordResponseOutput{})
	pulumi.RegisterOutputType(MxRecordResponseArrayOutput{})
	pulumi.RegisterOutputType(PtrRecordOutput{})
	pulumi.RegisterOutputType(PtrRecordArrayOutput{})
	pulumi.RegisterOutputType(PtrRecordResponseOutput{})
	pulumi.RegisterOutputType(PtrRecordResponseArrayOutput{})
	pulumi.RegisterOutputType(SoaRecordOutput{})
	pulumi.RegisterOutputType(SoaRecordPtrOutput{})
	pulumi.RegisterOutputType(SoaRecordResponseOutput{})
	pulumi.RegisterOutputType(SoaRecordResponsePtrOutput{})
	pulumi.RegisterOutputType(SrvRecordOutput{})
	pulumi.RegisterOutputType(SrvRecordArrayOutput{})
	pulumi.RegisterOutputType(SrvRecordResponseOutput{})
	pulumi.RegisterOutputType(SrvRecordResponseArrayOutput{})
	pulumi.RegisterOutputType(SubResourceOutput{})
	pulumi.RegisterOutputType(SubResourcePtrOutput{})
	pulumi.RegisterOutputType(SubResourceResponseOutput{})
	pulumi.RegisterOutputType(SubResourceResponsePtrOutput{})
	pulumi.RegisterOutputType(TxtRecordOutput{})
	pulumi.RegisterOutputType(TxtRecordArrayOutput{})
	pulumi.RegisterOutputType(TxtRecordResponseOutput{})
	pulumi.RegisterOutputType(TxtRecordResponseArrayOutput{})
}
