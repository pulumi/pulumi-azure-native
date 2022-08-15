


package v20150801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListConnectionConsentLinks(ctx *pulumi.Context, args *ListConnectionConsentLinksArgs, opts ...pulumi.InvokeOption) (*ListConnectionConsentLinksResult, error) {
	var rv ListConnectionConsentLinksResult
	err := ctx.Invoke("azure-native:web/v20150801preview:listConnectionConsentLinks", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListConnectionConsentLinksArgs struct {
	ConnectionName    string                      `pulumi:"connectionName"`
	Id                *string                     `pulumi:"id"`
	Kind              *string                     `pulumi:"kind"`
	Location          *string                     `pulumi:"location"`
	Name              *string                     `pulumi:"name"`
	Parameters        []ConsentLinkInputParameter `pulumi:"parameters"`
	ResourceGroupName string                      `pulumi:"resourceGroupName"`
	Tags              map[string]string           `pulumi:"tags"`
	Type              *string                     `pulumi:"type"`
}


type ListConnectionConsentLinksResult struct {
	Value []ConsentLinkResponse `pulumi:"value"`
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
	ConnectionName    pulumi.StringInput                  `pulumi:"connectionName"`
	Id                pulumi.StringPtrInput               `pulumi:"id"`
	Kind              pulumi.StringPtrInput               `pulumi:"kind"`
	Location          pulumi.StringPtrInput               `pulumi:"location"`
	Name              pulumi.StringPtrInput               `pulumi:"name"`
	Parameters        ConsentLinkInputParameterArrayInput `pulumi:"parameters"`
	ResourceGroupName pulumi.StringInput                  `pulumi:"resourceGroupName"`
	Tags              pulumi.StringMapInput               `pulumi:"tags"`
	Type              pulumi.StringPtrInput               `pulumi:"type"`
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

func (o ListConnectionConsentLinksResultOutput) Value() ConsentLinkResponseArrayOutput {
	return o.ApplyT(func(v ListConnectionConsentLinksResult) []ConsentLinkResponse { return v.Value }).(ConsentLinkResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListConnectionConsentLinksResultOutput{})
}
