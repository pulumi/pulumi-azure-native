


package v20180601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BlockchainProtocol string

const (
	BlockchainProtocolNotSpecified = BlockchainProtocol("NotSpecified")
	BlockchainProtocolParity       = BlockchainProtocol("Parity")
	BlockchainProtocolQuorum       = BlockchainProtocol("Quorum")
	BlockchainProtocolCorda        = BlockchainProtocol("Corda")
)

func (BlockchainProtocol) ElementType() reflect.Type {
	return reflect.TypeOf((*BlockchainProtocol)(nil)).Elem()
}

func (e BlockchainProtocol) ToBlockchainProtocolOutput() BlockchainProtocolOutput {
	return pulumi.ToOutput(e).(BlockchainProtocolOutput)
}

func (e BlockchainProtocol) ToBlockchainProtocolOutputWithContext(ctx context.Context) BlockchainProtocolOutput {
	return pulumi.ToOutputWithContext(ctx, e).(BlockchainProtocolOutput)
}

func (e BlockchainProtocol) ToBlockchainProtocolPtrOutput() BlockchainProtocolPtrOutput {
	return e.ToBlockchainProtocolPtrOutputWithContext(context.Background())
}

func (e BlockchainProtocol) ToBlockchainProtocolPtrOutputWithContext(ctx context.Context) BlockchainProtocolPtrOutput {
	return BlockchainProtocol(e).ToBlockchainProtocolOutputWithContext(ctx).ToBlockchainProtocolPtrOutputWithContext(ctx)
}

func (e BlockchainProtocol) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e BlockchainProtocol) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e BlockchainProtocol) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e BlockchainProtocol) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type BlockchainProtocolOutput struct{ *pulumi.OutputState }

func (BlockchainProtocolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlockchainProtocol)(nil)).Elem()
}

func (o BlockchainProtocolOutput) ToBlockchainProtocolOutput() BlockchainProtocolOutput {
	return o
}

func (o BlockchainProtocolOutput) ToBlockchainProtocolOutputWithContext(ctx context.Context) BlockchainProtocolOutput {
	return o
}

func (o BlockchainProtocolOutput) ToBlockchainProtocolPtrOutput() BlockchainProtocolPtrOutput {
	return o.ToBlockchainProtocolPtrOutputWithContext(context.Background())
}

func (o BlockchainProtocolOutput) ToBlockchainProtocolPtrOutputWithContext(ctx context.Context) BlockchainProtocolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BlockchainProtocol) *BlockchainProtocol {
		return &v
	}).(BlockchainProtocolPtrOutput)
}

func (o BlockchainProtocolOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o BlockchainProtocolOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BlockchainProtocol) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o BlockchainProtocolOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BlockchainProtocolOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BlockchainProtocol) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type BlockchainProtocolPtrOutput struct{ *pulumi.OutputState }

func (BlockchainProtocolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlockchainProtocol)(nil)).Elem()
}

func (o BlockchainProtocolPtrOutput) ToBlockchainProtocolPtrOutput() BlockchainProtocolPtrOutput {
	return o
}

func (o BlockchainProtocolPtrOutput) ToBlockchainProtocolPtrOutputWithContext(ctx context.Context) BlockchainProtocolPtrOutput {
	return o
}

func (o BlockchainProtocolPtrOutput) Elem() BlockchainProtocolOutput {
	return o.ApplyT(func(v *BlockchainProtocol) BlockchainProtocol {
		if v != nil {
			return *v
		}
		var ret BlockchainProtocol
		return ret
	}).(BlockchainProtocolOutput)
}

func (o BlockchainProtocolPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BlockchainProtocolPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *BlockchainProtocol) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type BlockchainProtocolInput interface {
	pulumi.Input

	ToBlockchainProtocolOutput() BlockchainProtocolOutput
	ToBlockchainProtocolOutputWithContext(context.Context) BlockchainProtocolOutput
}

var blockchainProtocolPtrType = reflect.TypeOf((**BlockchainProtocol)(nil)).Elem()

type BlockchainProtocolPtrInput interface {
	pulumi.Input

	ToBlockchainProtocolPtrOutput() BlockchainProtocolPtrOutput
	ToBlockchainProtocolPtrOutputWithContext(context.Context) BlockchainProtocolPtrOutput
}

type blockchainProtocolPtr string

func BlockchainProtocolPtr(v string) BlockchainProtocolPtrInput {
	return (*blockchainProtocolPtr)(&v)
}

func (*blockchainProtocolPtr) ElementType() reflect.Type {
	return blockchainProtocolPtrType
}

func (in *blockchainProtocolPtr) ToBlockchainProtocolPtrOutput() BlockchainProtocolPtrOutput {
	return pulumi.ToOutput(in).(BlockchainProtocolPtrOutput)
}

func (in *blockchainProtocolPtr) ToBlockchainProtocolPtrOutputWithContext(ctx context.Context) BlockchainProtocolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(BlockchainProtocolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(BlockchainProtocolOutput{})
	pulumi.RegisterOutputType(BlockchainProtocolPtrOutput{})
}
