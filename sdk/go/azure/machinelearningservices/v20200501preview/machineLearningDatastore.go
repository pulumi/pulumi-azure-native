


package v20200501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MachineLearningDatastore struct {
	pulumi.CustomResourceState

	Identity   IdentityResponsePtrOutput `pulumi:"identity"`
	Location   pulumi.StringPtrOutput    `pulumi:"location"`
	Name       pulumi.StringOutput       `pulumi:"name"`
	Properties DatastoreResponseOutput   `pulumi:"properties"`
	Sku        SkuResponsePtrOutput      `pulumi:"sku"`
	Tags       pulumi.StringMapOutput    `pulumi:"tags"`
	Type       pulumi.StringOutput       `pulumi:"type"`
}


func NewMachineLearningDatastore(ctx *pulumi.Context,
	name string, args *MachineLearningDatastoreArgs, opts ...pulumi.ResourceOption) (*MachineLearningDatastore, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataStoreType == nil {
		return nil, errors.New("invalid value for required argument 'DataStoreType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	if args.EnforceSSL == nil {
		args.EnforceSSL = pulumi.BoolPtr(true)
	}
	if args.IncludeSecret == nil {
		args.IncludeSecret = pulumi.BoolPtr(true)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:MachineLearningDatastore"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210301preview:MachineLearningDatastore"),
		},
	})
	opts = append(opts, aliases)
	var resource MachineLearningDatastore
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20200501preview:MachineLearningDatastore", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMachineLearningDatastore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MachineLearningDatastoreState, opts ...pulumi.ResourceOption) (*MachineLearningDatastore, error) {
	var resource MachineLearningDatastore
	err := ctx.ReadResource("azure-native:machinelearningservices/v20200501preview:MachineLearningDatastore", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type machineLearningDatastoreState struct {
}

type MachineLearningDatastoreState struct {
}

func (MachineLearningDatastoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*machineLearningDatastoreState)(nil)).Elem()
}

type machineLearningDatastoreArgs struct {
	AccountKey                      *string `pulumi:"accountKey"`
	AccountName                     *string `pulumi:"accountName"`
	AdlsResourceGroup               *string `pulumi:"adlsResourceGroup"`
	AdlsSubscriptionId              *string `pulumi:"adlsSubscriptionId"`
	AuthorityUrl                    *string `pulumi:"authorityUrl"`
	ClientId                        *string `pulumi:"clientId"`
	ClientSecret                    *string `pulumi:"clientSecret"`
	ContainerName                   *string `pulumi:"containerName"`
	DataStoreType                   string  `pulumi:"dataStoreType"`
	DatabaseName                    *string `pulumi:"databaseName"`
	DatastoreName                   *string `pulumi:"datastoreName"`
	Description                     *string `pulumi:"description"`
	Endpoint                        *string `pulumi:"endpoint"`
	EnforceSSL                      *bool   `pulumi:"enforceSSL"`
	FileSystem                      *string `pulumi:"fileSystem"`
	IncludeSecret                   *bool   `pulumi:"includeSecret"`
	Name                            *string `pulumi:"name"`
	Password                        *string `pulumi:"password"`
	Port                            *string `pulumi:"port"`
	Protocol                        *string `pulumi:"protocol"`
	ResourceGroupName               string  `pulumi:"resourceGroupName"`
	ResourceUrl                     *string `pulumi:"resourceUrl"`
	SasToken                        *string `pulumi:"sasToken"`
	ServerName                      *string `pulumi:"serverName"`
	ShareName                       *string `pulumi:"shareName"`
	SkipValidation                  *bool   `pulumi:"skipValidation"`
	StorageAccountResourceGroup     *string `pulumi:"storageAccountResourceGroup"`
	StorageAccountSubscriptionId    *string `pulumi:"storageAccountSubscriptionId"`
	StoreName                       *string `pulumi:"storeName"`
	TenantId                        *string `pulumi:"tenantId"`
	UserId                          *string `pulumi:"userId"`
	UserName                        *string `pulumi:"userName"`
	WorkspaceName                   string  `pulumi:"workspaceName"`
	WorkspaceSystemAssignedIdentity *bool   `pulumi:"workspaceSystemAssignedIdentity"`
}


type MachineLearningDatastoreArgs struct {
	AccountKey                      pulumi.StringPtrInput
	AccountName                     pulumi.StringPtrInput
	AdlsResourceGroup               pulumi.StringPtrInput
	AdlsSubscriptionId              pulumi.StringPtrInput
	AuthorityUrl                    pulumi.StringPtrInput
	ClientId                        pulumi.StringPtrInput
	ClientSecret                    pulumi.StringPtrInput
	ContainerName                   pulumi.StringPtrInput
	DataStoreType                   pulumi.StringInput
	DatabaseName                    pulumi.StringPtrInput
	DatastoreName                   pulumi.StringPtrInput
	Description                     pulumi.StringPtrInput
	Endpoint                        pulumi.StringPtrInput
	EnforceSSL                      pulumi.BoolPtrInput
	FileSystem                      pulumi.StringPtrInput
	IncludeSecret                   pulumi.BoolPtrInput
	Name                            pulumi.StringPtrInput
	Password                        pulumi.StringPtrInput
	Port                            pulumi.StringPtrInput
	Protocol                        pulumi.StringPtrInput
	ResourceGroupName               pulumi.StringInput
	ResourceUrl                     pulumi.StringPtrInput
	SasToken                        pulumi.StringPtrInput
	ServerName                      pulumi.StringPtrInput
	ShareName                       pulumi.StringPtrInput
	SkipValidation                  pulumi.BoolPtrInput
	StorageAccountResourceGroup     pulumi.StringPtrInput
	StorageAccountSubscriptionId    pulumi.StringPtrInput
	StoreName                       pulumi.StringPtrInput
	TenantId                        pulumi.StringPtrInput
	UserId                          pulumi.StringPtrInput
	UserName                        pulumi.StringPtrInput
	WorkspaceName                   pulumi.StringInput
	WorkspaceSystemAssignedIdentity pulumi.BoolPtrInput
}

func (MachineLearningDatastoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*machineLearningDatastoreArgs)(nil)).Elem()
}

type MachineLearningDatastoreInput interface {
	pulumi.Input

	ToMachineLearningDatastoreOutput() MachineLearningDatastoreOutput
	ToMachineLearningDatastoreOutputWithContext(ctx context.Context) MachineLearningDatastoreOutput
}

func (*MachineLearningDatastore) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineLearningDatastore)(nil))
}

func (i *MachineLearningDatastore) ToMachineLearningDatastoreOutput() MachineLearningDatastoreOutput {
	return i.ToMachineLearningDatastoreOutputWithContext(context.Background())
}

func (i *MachineLearningDatastore) ToMachineLearningDatastoreOutputWithContext(ctx context.Context) MachineLearningDatastoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineLearningDatastoreOutput)
}

type MachineLearningDatastoreOutput struct{ *pulumi.OutputState }

func (MachineLearningDatastoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineLearningDatastore)(nil))
}

func (o MachineLearningDatastoreOutput) ToMachineLearningDatastoreOutput() MachineLearningDatastoreOutput {
	return o
}

func (o MachineLearningDatastoreOutput) ToMachineLearningDatastoreOutputWithContext(ctx context.Context) MachineLearningDatastoreOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MachineLearningDatastoreOutput{})
}
