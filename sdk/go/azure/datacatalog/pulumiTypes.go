


package datacatalog

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Principals struct {
	ObjectId *string `pulumi:"objectId"`
	Upn      *string `pulumi:"upn"`
}





type PrincipalsInput interface {
	pulumi.Input

	ToPrincipalsOutput() PrincipalsOutput
	ToPrincipalsOutputWithContext(context.Context) PrincipalsOutput
}

type PrincipalsArgs struct {
	ObjectId pulumi.StringPtrInput `pulumi:"objectId"`
	Upn      pulumi.StringPtrInput `pulumi:"upn"`
}

func (PrincipalsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Principals)(nil)).Elem()
}

func (i PrincipalsArgs) ToPrincipalsOutput() PrincipalsOutput {
	return i.ToPrincipalsOutputWithContext(context.Background())
}

func (i PrincipalsArgs) ToPrincipalsOutputWithContext(ctx context.Context) PrincipalsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrincipalsOutput)
}





type PrincipalsArrayInput interface {
	pulumi.Input

	ToPrincipalsArrayOutput() PrincipalsArrayOutput
	ToPrincipalsArrayOutputWithContext(context.Context) PrincipalsArrayOutput
}

type PrincipalsArray []PrincipalsInput

func (PrincipalsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Principals)(nil)).Elem()
}

func (i PrincipalsArray) ToPrincipalsArrayOutput() PrincipalsArrayOutput {
	return i.ToPrincipalsArrayOutputWithContext(context.Background())
}

func (i PrincipalsArray) ToPrincipalsArrayOutputWithContext(ctx context.Context) PrincipalsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrincipalsArrayOutput)
}

type PrincipalsOutput struct{ *pulumi.OutputState }

func (PrincipalsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Principals)(nil)).Elem()
}

func (o PrincipalsOutput) ToPrincipalsOutput() PrincipalsOutput {
	return o
}

func (o PrincipalsOutput) ToPrincipalsOutputWithContext(ctx context.Context) PrincipalsOutput {
	return o
}

func (o PrincipalsOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Principals) *string { return v.ObjectId }).(pulumi.StringPtrOutput)
}

func (o PrincipalsOutput) Upn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Principals) *string { return v.Upn }).(pulumi.StringPtrOutput)
}

type PrincipalsArrayOutput struct{ *pulumi.OutputState }

func (PrincipalsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Principals)(nil)).Elem()
}

func (o PrincipalsArrayOutput) ToPrincipalsArrayOutput() PrincipalsArrayOutput {
	return o
}

func (o PrincipalsArrayOutput) ToPrincipalsArrayOutputWithContext(ctx context.Context) PrincipalsArrayOutput {
	return o
}

func (o PrincipalsArrayOutput) Index(i pulumi.IntInput) PrincipalsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Principals {
		return vs[0].([]Principals)[vs[1].(int)]
	}).(PrincipalsOutput)
}

type PrincipalsResponse struct {
	ObjectId *string `pulumi:"objectId"`
	Upn      *string `pulumi:"upn"`
}





type PrincipalsResponseInput interface {
	pulumi.Input

	ToPrincipalsResponseOutput() PrincipalsResponseOutput
	ToPrincipalsResponseOutputWithContext(context.Context) PrincipalsResponseOutput
}

type PrincipalsResponseArgs struct {
	ObjectId pulumi.StringPtrInput `pulumi:"objectId"`
	Upn      pulumi.StringPtrInput `pulumi:"upn"`
}

func (PrincipalsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrincipalsResponse)(nil)).Elem()
}

func (i PrincipalsResponseArgs) ToPrincipalsResponseOutput() PrincipalsResponseOutput {
	return i.ToPrincipalsResponseOutputWithContext(context.Background())
}

func (i PrincipalsResponseArgs) ToPrincipalsResponseOutputWithContext(ctx context.Context) PrincipalsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrincipalsResponseOutput)
}





type PrincipalsResponseArrayInput interface {
	pulumi.Input

	ToPrincipalsResponseArrayOutput() PrincipalsResponseArrayOutput
	ToPrincipalsResponseArrayOutputWithContext(context.Context) PrincipalsResponseArrayOutput
}

type PrincipalsResponseArray []PrincipalsResponseInput

func (PrincipalsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrincipalsResponse)(nil)).Elem()
}

func (i PrincipalsResponseArray) ToPrincipalsResponseArrayOutput() PrincipalsResponseArrayOutput {
	return i.ToPrincipalsResponseArrayOutputWithContext(context.Background())
}

func (i PrincipalsResponseArray) ToPrincipalsResponseArrayOutputWithContext(ctx context.Context) PrincipalsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrincipalsResponseArrayOutput)
}

type PrincipalsResponseOutput struct{ *pulumi.OutputState }

func (PrincipalsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrincipalsResponse)(nil)).Elem()
}

func (o PrincipalsResponseOutput) ToPrincipalsResponseOutput() PrincipalsResponseOutput {
	return o
}

func (o PrincipalsResponseOutput) ToPrincipalsResponseOutputWithContext(ctx context.Context) PrincipalsResponseOutput {
	return o
}

func (o PrincipalsResponseOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrincipalsResponse) *string { return v.ObjectId }).(pulumi.StringPtrOutput)
}

func (o PrincipalsResponseOutput) Upn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrincipalsResponse) *string { return v.Upn }).(pulumi.StringPtrOutput)
}

type PrincipalsResponseArrayOutput struct{ *pulumi.OutputState }

func (PrincipalsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrincipalsResponse)(nil)).Elem()
}

func (o PrincipalsResponseArrayOutput) ToPrincipalsResponseArrayOutput() PrincipalsResponseArrayOutput {
	return o
}

func (o PrincipalsResponseArrayOutput) ToPrincipalsResponseArrayOutputWithContext(ctx context.Context) PrincipalsResponseArrayOutput {
	return o
}

func (o PrincipalsResponseArrayOutput) Index(i pulumi.IntInput) PrincipalsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrincipalsResponse {
		return vs[0].([]PrincipalsResponse)[vs[1].(int)]
	}).(PrincipalsResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(PrincipalsOutput{})
	pulumi.RegisterOutputType(PrincipalsArrayOutput{})
	pulumi.RegisterOutputType(PrincipalsResponseOutput{})
	pulumi.RegisterOutputType(PrincipalsResponseArrayOutput{})
}
