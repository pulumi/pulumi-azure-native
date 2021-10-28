


package media

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ContentKeyPolicy struct {
	pulumi.CustomResourceState

	Created      pulumi.StringOutput                       `pulumi:"created"`
	Description  pulumi.StringPtrOutput                    `pulumi:"description"`
	LastModified pulumi.StringOutput                       `pulumi:"lastModified"`
	Name         pulumi.StringOutput                       `pulumi:"name"`
	Options      ContentKeyPolicyOptionResponseArrayOutput `pulumi:"options"`
	PolicyId     pulumi.StringOutput                       `pulumi:"policyId"`
	SystemData   SystemDataResponseOutput                  `pulumi:"systemData"`
	Type         pulumi.StringOutput                       `pulumi:"type"`
}


func NewContentKeyPolicy(ctx *pulumi.Context,
	name string, args *ContentKeyPolicyArgs, opts ...pulumi.ResourceOption) (*ContentKeyPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Options == nil {
		return nil, errors.New("invalid value for required argument 'Options'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:media:ContentKeyPolicy"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180330preview:ContentKeyPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:media/v20180330preview:ContentKeyPolicy"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180601preview:ContentKeyPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:media/v20180601preview:ContentKeyPolicy"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180701:ContentKeyPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:media/v20180701:ContentKeyPolicy"),
		},
		{
			Type: pulumi.String("azure-native:media/v20200501:ContentKeyPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:media/v20200501:ContentKeyPolicy"),
		},
		{
			Type: pulumi.String("azure-native:media/v20210601:ContentKeyPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:media/v20210601:ContentKeyPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource ContentKeyPolicy
	err := ctx.RegisterResource("azure-native:media:ContentKeyPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetContentKeyPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContentKeyPolicyState, opts ...pulumi.ResourceOption) (*ContentKeyPolicy, error) {
	var resource ContentKeyPolicy
	err := ctx.ReadResource("azure-native:media:ContentKeyPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type contentKeyPolicyState struct {
}

type ContentKeyPolicyState struct {
}

func (ContentKeyPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*contentKeyPolicyState)(nil)).Elem()
}

type contentKeyPolicyArgs struct {
	AccountName          string                   `pulumi:"accountName"`
	ContentKeyPolicyName *string                  `pulumi:"contentKeyPolicyName"`
	Description          *string                  `pulumi:"description"`
	Options              []ContentKeyPolicyOption `pulumi:"options"`
	ResourceGroupName    string                   `pulumi:"resourceGroupName"`
}


type ContentKeyPolicyArgs struct {
	AccountName          pulumi.StringInput
	ContentKeyPolicyName pulumi.StringPtrInput
	Description          pulumi.StringPtrInput
	Options              ContentKeyPolicyOptionArrayInput
	ResourceGroupName    pulumi.StringInput
}

func (ContentKeyPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*contentKeyPolicyArgs)(nil)).Elem()
}

type ContentKeyPolicyInput interface {
	pulumi.Input

	ToContentKeyPolicyOutput() ContentKeyPolicyOutput
	ToContentKeyPolicyOutputWithContext(ctx context.Context) ContentKeyPolicyOutput
}

func (*ContentKeyPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicy)(nil))
}

func (i *ContentKeyPolicy) ToContentKeyPolicyOutput() ContentKeyPolicyOutput {
	return i.ToContentKeyPolicyOutputWithContext(context.Background())
}

func (i *ContentKeyPolicy) ToContentKeyPolicyOutputWithContext(ctx context.Context) ContentKeyPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyOutput)
}

type ContentKeyPolicyOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicy)(nil))
}

func (o ContentKeyPolicyOutput) ToContentKeyPolicyOutput() ContentKeyPolicyOutput {
	return o
}

func (o ContentKeyPolicyOutput) ToContentKeyPolicyOutputWithContext(ctx context.Context) ContentKeyPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ContentKeyPolicyOutput{})
}
