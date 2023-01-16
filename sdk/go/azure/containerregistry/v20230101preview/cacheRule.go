


package v20230101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CacheRule struct {
	pulumi.CustomResourceState

	CreationDate            pulumi.StringOutput      `pulumi:"creationDate"`
	CredentialSetResourceId pulumi.StringPtrOutput   `pulumi:"credentialSetResourceId"`
	Name                    pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState       pulumi.StringOutput      `pulumi:"provisioningState"`
	SourceRepository        pulumi.StringPtrOutput   `pulumi:"sourceRepository"`
	SystemData              SystemDataResponseOutput `pulumi:"systemData"`
	TargetRepository        pulumi.StringPtrOutput   `pulumi:"targetRepository"`
	Type                    pulumi.StringOutput      `pulumi:"type"`
}


func NewCacheRule(ctx *pulumi.Context,
	name string, args *CacheRuleArgs, opts ...pulumi.ResourceOption) (*CacheRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource CacheRule
	err := ctx.RegisterResource("azure-native:containerregistry/v20230101preview:CacheRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCacheRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CacheRuleState, opts ...pulumi.ResourceOption) (*CacheRule, error) {
	var resource CacheRule
	err := ctx.ReadResource("azure-native:containerregistry/v20230101preview:CacheRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type cacheRuleState struct {
}

type CacheRuleState struct {
}

func (CacheRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*cacheRuleState)(nil)).Elem()
}

type cacheRuleArgs struct {
	CacheRuleName           *string `pulumi:"cacheRuleName"`
	CredentialSetResourceId *string `pulumi:"credentialSetResourceId"`
	RegistryName            string  `pulumi:"registryName"`
	ResourceGroupName       string  `pulumi:"resourceGroupName"`
	SourceRepository        *string `pulumi:"sourceRepository"`
	TargetRepository        *string `pulumi:"targetRepository"`
}


type CacheRuleArgs struct {
	CacheRuleName           pulumi.StringPtrInput
	CredentialSetResourceId pulumi.StringPtrInput
	RegistryName            pulumi.StringInput
	ResourceGroupName       pulumi.StringInput
	SourceRepository        pulumi.StringPtrInput
	TargetRepository        pulumi.StringPtrInput
}

func (CacheRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cacheRuleArgs)(nil)).Elem()
}

type CacheRuleInput interface {
	pulumi.Input

	ToCacheRuleOutput() CacheRuleOutput
	ToCacheRuleOutputWithContext(ctx context.Context) CacheRuleOutput
}

func (*CacheRule) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheRule)(nil)).Elem()
}

func (i *CacheRule) ToCacheRuleOutput() CacheRuleOutput {
	return i.ToCacheRuleOutputWithContext(context.Background())
}

func (i *CacheRule) ToCacheRuleOutputWithContext(ctx context.Context) CacheRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheRuleOutput)
}

type CacheRuleOutput struct{ *pulumi.OutputState }

func (CacheRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheRule)(nil)).Elem()
}

func (o CacheRuleOutput) ToCacheRuleOutput() CacheRuleOutput {
	return o
}

func (o CacheRuleOutput) ToCacheRuleOutputWithContext(ctx context.Context) CacheRuleOutput {
	return o
}

func (o CacheRuleOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *CacheRule) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

func (o CacheRuleOutput) CredentialSetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheRule) pulumi.StringPtrOutput { return v.CredentialSetResourceId }).(pulumi.StringPtrOutput)
}

func (o CacheRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CacheRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CacheRuleOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *CacheRule) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o CacheRuleOutput) SourceRepository() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheRule) pulumi.StringPtrOutput { return v.SourceRepository }).(pulumi.StringPtrOutput)
}

func (o CacheRuleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *CacheRule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o CacheRuleOutput) TargetRepository() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheRule) pulumi.StringPtrOutput { return v.TargetRepository }).(pulumi.StringPtrOutput)
}

func (o CacheRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CacheRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CacheRuleOutput{})
}
