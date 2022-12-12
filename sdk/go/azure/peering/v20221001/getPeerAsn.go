


package v20221001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPeerAsn(ctx *pulumi.Context, args *LookupPeerAsnArgs, opts ...pulumi.InvokeOption) (*LookupPeerAsnResult, error) {
	var rv LookupPeerAsnResult
	err := ctx.Invoke("azure-native:peering/v20221001:getPeerAsn", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPeerAsnArgs struct {
	PeerAsnName string `pulumi:"peerAsnName"`
}


type LookupPeerAsnResult struct {
	ErrorMessage      string                  `pulumi:"errorMessage"`
	Id                string                  `pulumi:"id"`
	Name              string                  `pulumi:"name"`
	PeerAsn           *int                    `pulumi:"peerAsn"`
	PeerContactDetail []ContactDetailResponse `pulumi:"peerContactDetail"`
	PeerName          *string                 `pulumi:"peerName"`
	Type              string                  `pulumi:"type"`
	ValidationState   string                  `pulumi:"validationState"`
}

func LookupPeerAsnOutput(ctx *pulumi.Context, args LookupPeerAsnOutputArgs, opts ...pulumi.InvokeOption) LookupPeerAsnResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPeerAsnResult, error) {
			args := v.(LookupPeerAsnArgs)
			r, err := LookupPeerAsn(ctx, &args, opts...)
			var s LookupPeerAsnResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPeerAsnResultOutput)
}

type LookupPeerAsnOutputArgs struct {
	PeerAsnName pulumi.StringInput `pulumi:"peerAsnName"`
}

func (LookupPeerAsnOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPeerAsnArgs)(nil)).Elem()
}


type LookupPeerAsnResultOutput struct{ *pulumi.OutputState }

func (LookupPeerAsnResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPeerAsnResult)(nil)).Elem()
}

func (o LookupPeerAsnResultOutput) ToLookupPeerAsnResultOutput() LookupPeerAsnResultOutput {
	return o
}

func (o LookupPeerAsnResultOutput) ToLookupPeerAsnResultOutputWithContext(ctx context.Context) LookupPeerAsnResultOutput {
	return o
}

func (o LookupPeerAsnResultOutput) ErrorMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPeerAsnResult) string { return v.ErrorMessage }).(pulumi.StringOutput)
}

func (o LookupPeerAsnResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPeerAsnResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPeerAsnResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPeerAsnResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPeerAsnResultOutput) PeerAsn() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupPeerAsnResult) *int { return v.PeerAsn }).(pulumi.IntPtrOutput)
}

func (o LookupPeerAsnResultOutput) PeerContactDetail() ContactDetailResponseArrayOutput {
	return o.ApplyT(func(v LookupPeerAsnResult) []ContactDetailResponse { return v.PeerContactDetail }).(ContactDetailResponseArrayOutput)
}

func (o LookupPeerAsnResultOutput) PeerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPeerAsnResult) *string { return v.PeerName }).(pulumi.StringPtrOutput)
}

func (o LookupPeerAsnResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPeerAsnResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupPeerAsnResultOutput) ValidationState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPeerAsnResult) string { return v.ValidationState }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPeerAsnResultOutput{})
}
