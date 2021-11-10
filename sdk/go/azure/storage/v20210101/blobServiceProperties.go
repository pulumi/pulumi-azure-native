


package v20210101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BlobServiceProperties struct {
	pulumi.CustomResourceState

	AutomaticSnapshotPolicyEnabled pulumi.BoolPtrOutput                          `pulumi:"automaticSnapshotPolicyEnabled"`
	ChangeFeed                     ChangeFeedResponsePtrOutput                   `pulumi:"changeFeed"`
	ContainerDeleteRetentionPolicy DeleteRetentionPolicyResponsePtrOutput        `pulumi:"containerDeleteRetentionPolicy"`
	Cors                           CorsRulesResponsePtrOutput                    `pulumi:"cors"`
	DefaultServiceVersion          pulumi.StringPtrOutput                        `pulumi:"defaultServiceVersion"`
	DeleteRetentionPolicy          DeleteRetentionPolicyResponsePtrOutput        `pulumi:"deleteRetentionPolicy"`
	IsVersioningEnabled            pulumi.BoolPtrOutput                          `pulumi:"isVersioningEnabled"`
	LastAccessTimeTrackingPolicy   LastAccessTimeTrackingPolicyResponsePtrOutput `pulumi:"lastAccessTimeTrackingPolicy"`
	Name                           pulumi.StringOutput                           `pulumi:"name"`
	RestorePolicy                  RestorePolicyPropertiesResponsePtrOutput      `pulumi:"restorePolicy"`
	Sku                            SkuResponseOutput                             `pulumi:"sku"`
	Type                           pulumi.StringOutput                           `pulumi:"type"`
}


func NewBlobServiceProperties(ctx *pulumi.Context,
	name string, args *BlobServicePropertiesArgs, opts ...pulumi.ResourceOption) (*BlobServiceProperties, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storage:BlobServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20180701:BlobServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20181101:BlobServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20190401:BlobServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20190601:BlobServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20200801preview:BlobServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210201:BlobServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210401:BlobServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210601:BlobServiceProperties"),
		},
	})
	opts = append(opts, aliases)
	var resource BlobServiceProperties
	err := ctx.RegisterResource("azure-native:storage/v20210101:BlobServiceProperties", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBlobServiceProperties(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BlobServicePropertiesState, opts ...pulumi.ResourceOption) (*BlobServiceProperties, error) {
	var resource BlobServiceProperties
	err := ctx.ReadResource("azure-native:storage/v20210101:BlobServiceProperties", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type blobServicePropertiesState struct {
}

type BlobServicePropertiesState struct {
}

func (BlobServicePropertiesState) ElementType() reflect.Type {
	return reflect.TypeOf((*blobServicePropertiesState)(nil)).Elem()
}

type blobServicePropertiesArgs struct {
	AccountName                    string                        `pulumi:"accountName"`
	AutomaticSnapshotPolicyEnabled *bool                         `pulumi:"automaticSnapshotPolicyEnabled"`
	BlobServicesName               *string                       `pulumi:"blobServicesName"`
	ChangeFeed                     *ChangeFeed                   `pulumi:"changeFeed"`
	ContainerDeleteRetentionPolicy *DeleteRetentionPolicy        `pulumi:"containerDeleteRetentionPolicy"`
	Cors                           *CorsRules                    `pulumi:"cors"`
	DefaultServiceVersion          *string                       `pulumi:"defaultServiceVersion"`
	DeleteRetentionPolicy          *DeleteRetentionPolicy        `pulumi:"deleteRetentionPolicy"`
	IsVersioningEnabled            *bool                         `pulumi:"isVersioningEnabled"`
	LastAccessTimeTrackingPolicy   *LastAccessTimeTrackingPolicy `pulumi:"lastAccessTimeTrackingPolicy"`
	ResourceGroupName              string                        `pulumi:"resourceGroupName"`
	RestorePolicy                  *RestorePolicyProperties      `pulumi:"restorePolicy"`
}


type BlobServicePropertiesArgs struct {
	AccountName                    pulumi.StringInput
	AutomaticSnapshotPolicyEnabled pulumi.BoolPtrInput
	BlobServicesName               pulumi.StringPtrInput
	ChangeFeed                     ChangeFeedPtrInput
	ContainerDeleteRetentionPolicy DeleteRetentionPolicyPtrInput
	Cors                           CorsRulesPtrInput
	DefaultServiceVersion          pulumi.StringPtrInput
	DeleteRetentionPolicy          DeleteRetentionPolicyPtrInput
	IsVersioningEnabled            pulumi.BoolPtrInput
	LastAccessTimeTrackingPolicy   LastAccessTimeTrackingPolicyPtrInput
	ResourceGroupName              pulumi.StringInput
	RestorePolicy                  RestorePolicyPropertiesPtrInput
}

func (BlobServicePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*blobServicePropertiesArgs)(nil)).Elem()
}

type BlobServicePropertiesInput interface {
	pulumi.Input

	ToBlobServicePropertiesOutput() BlobServicePropertiesOutput
	ToBlobServicePropertiesOutputWithContext(ctx context.Context) BlobServicePropertiesOutput
}

func (*BlobServiceProperties) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobServiceProperties)(nil))
}

func (i *BlobServiceProperties) ToBlobServicePropertiesOutput() BlobServicePropertiesOutput {
	return i.ToBlobServicePropertiesOutputWithContext(context.Background())
}

func (i *BlobServiceProperties) ToBlobServicePropertiesOutputWithContext(ctx context.Context) BlobServicePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobServicePropertiesOutput)
}

type BlobServicePropertiesOutput struct{ *pulumi.OutputState }

func (BlobServicePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobServiceProperties)(nil))
}

func (o BlobServicePropertiesOutput) ToBlobServicePropertiesOutput() BlobServicePropertiesOutput {
	return o
}

func (o BlobServicePropertiesOutput) ToBlobServicePropertiesOutputWithContext(ctx context.Context) BlobServicePropertiesOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BlobServicePropertiesOutput{})
}
