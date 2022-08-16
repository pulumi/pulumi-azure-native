


package v20160601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListConnectionConsentLinks(ctx *pulumi.Context, args *ListConnectionConsentLinksArgs, opts ...pulumi.InvokeOption) (*ListConnectionConsentLinksResult, error) {
	var rv ListConnectionConsentLinksResult
	err := ctx.Invoke("azure-native:web/v20160601:listConnectionConsentLinks", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListConnectionConsentLinksArgs struct {
	ConnectionName    string                           `pulumi:"connectionName"`
	Parameters        []ConsentLinkParameterDefinition `pulumi:"parameters"`
	ResourceGroupName string                           `pulumi:"resourceGroupName"`
	SubscriptionId    *string                          `pulumi:"subscriptionId"`
}


type ListConnectionConsentLinksResult struct {
	Value []ConsentLinkDefinitionResponse `pulumi:"value"`
}

func ListConnectionConsentLinksOutput(ctx *pulumi.Context, args ListConnectionConsentLinksOutputArgs, opts ...pulumi.InvokeOption) ListConnectionConsentLinksResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListConnectionConsentLinksResult, error) {
			args := v.(ListConnectionConsentLinksArgs)
			r, err := ListConnectionConsentLinks(ctx, &args, opts...)
			var s ListConnectionConsentLinksResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListConnectionConsentLinksResultOutput)
}

type ListConnectionConsentLinksOutputArgs struct {
	ConnectionName    pulumi.StringInput                       `pulumi:"connectionName"`
	Parameters        ConsentLinkParameterDefinitionArrayInput `pulumi:"parameters"`
	ResourceGroupName pulumi.StringInput                       `pulumi:"resourceGroupName"`
	SubscriptionId    pulumi.StringPtrInput                    `pulumi:"subscriptionId"`
}

func (ListConnectionConsentLinksOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListConnectionConsentLinksArgs)(nil)).Elem()
}


type ListConnectionConsentLinksResultOutput struct{ *pulumi.OutputState }

func (ListConnectionConsentLinksResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListConnectionConsentLinksResult)(nil)).Elem()
}

func (o ListConnectionConsentLinksResultOutput) ToListConnectionConsentLinksResultOutput() ListConnectionConsentLinksResultOutput {
	return o
}

func (o ListConnectionConsentLinksResultOutput) ToListConnectionConsentLinksResultOutputWithContext(ctx context.Context) ListConnectionConsentLinksResultOutput {
	return o
}

func (o ListConnectionConsentLinksResultOutput) Value() ConsentLinkDefinitionResponseArrayOutput {
	return o.ApplyT(func(v ListConnectionConsentLinksResult) []ConsentLinkDefinitionResponse { return v.Value }).(ConsentLinkDefinitionResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListConnectionConsentLinksResultOutput{})
}
