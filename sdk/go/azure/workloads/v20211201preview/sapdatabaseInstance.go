


package v20211201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SAPDatabaseInstance struct {
	pulumi.CustomResourceState

	DatabaseSid       pulumi.StringOutput                   `pulumi:"databaseSid"`
	DatabaseType      pulumi.StringOutput                   `pulumi:"databaseType"`
	Errors            SAPVirtualInstanceErrorResponseOutput `pulumi:"errors"`
	IpAddress         pulumi.StringOutput                   `pulumi:"ipAddress"`
	Location          pulumi.StringOutput                   `pulumi:"location"`
	Name              pulumi.StringOutput                   `pulumi:"name"`
	ProvisioningState pulumi.StringOutput                   `pulumi:"provisioningState"`
	Status            pulumi.StringOutput                   `pulumi:"status"`
	Subnet            pulumi.StringOutput                   `pulumi:"subnet"`
	SystemData        SystemDataResponseOutput              `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput                `pulumi:"tags"`
	Type              pulumi.StringOutput                   `pulumi:"type"`
	VmDetails         DatabaseVmDetailsResponseArrayOutput  `pulumi:"vmDetails"`
}


func NewSAPDatabaseInstance(ctx *pulumi.Context,
	name string, args *SAPDatabaseInstanceArgs, opts ...pulumi.ResourceOption) (*SAPDatabaseInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SapVirtualInstanceName == nil {
		return nil, errors.New("invalid value for required argument 'SapVirtualInstanceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:workloads:SAPDatabaseInstance"),
		},
	})
	opts = append(opts, aliases)
	var resource SAPDatabaseInstance
	err := ctx.RegisterResource("azure-native:workloads/v20211201preview:SAPDatabaseInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSAPDatabaseInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SAPDatabaseInstanceState, opts ...pulumi.ResourceOption) (*SAPDatabaseInstance, error) {
	var resource SAPDatabaseInstance
	err := ctx.ReadResource("azure-native:workloads/v20211201preview:SAPDatabaseInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type sapdatabaseInstanceState struct {
}

type SAPDatabaseInstanceState struct {
}

func (SAPDatabaseInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*sapdatabaseInstanceState)(nil)).Elem()
}

type sapdatabaseInstanceArgs struct {
	DatabaseInstanceName   *string           `pulumi:"databaseInstanceName"`
	Location               *string           `pulumi:"location"`
	ResourceGroupName      string            `pulumi:"resourceGroupName"`
	SapVirtualInstanceName string            `pulumi:"sapVirtualInstanceName"`
	Tags                   map[string]string `pulumi:"tags"`
}


type SAPDatabaseInstanceArgs struct {
	DatabaseInstanceName   pulumi.StringPtrInput
	Location               pulumi.StringPtrInput
	ResourceGroupName      pulumi.StringInput
	SapVirtualInstanceName pulumi.StringInput
	Tags                   pulumi.StringMapInput
}

func (SAPDatabaseInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sapdatabaseInstanceArgs)(nil)).Elem()
}

type SAPDatabaseInstanceInput interface {
	pulumi.Input

	ToSAPDatabaseInstanceOutput() SAPDatabaseInstanceOutput
	ToSAPDatabaseInstanceOutputWithContext(ctx context.Context) SAPDatabaseInstanceOutput
}

func (*SAPDatabaseInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**SAPDatabaseInstance)(nil)).Elem()
}

func (i *SAPDatabaseInstance) ToSAPDatabaseInstanceOutput() SAPDatabaseInstanceOutput {
	return i.ToSAPDatabaseInstanceOutputWithContext(context.Background())
}

func (i *SAPDatabaseInstance) ToSAPDatabaseInstanceOutputWithContext(ctx context.Context) SAPDatabaseInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SAPDatabaseInstanceOutput)
}

type SAPDatabaseInstanceOutput struct{ *pulumi.OutputState }

func (SAPDatabaseInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SAPDatabaseInstance)(nil)).Elem()
}

func (o SAPDatabaseInstanceOutput) ToSAPDatabaseInstanceOutput() SAPDatabaseInstanceOutput {
	return o
}

func (o SAPDatabaseInstanceOutput) ToSAPDatabaseInstanceOutputWithContext(ctx context.Context) SAPDatabaseInstanceOutput {
	return o
}

func (o SAPDatabaseInstanceOutput) DatabaseSid() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPDatabaseInstance) pulumi.StringOutput { return v.DatabaseSid }).(pulumi.StringOutput)
}

func (o SAPDatabaseInstanceOutput) DatabaseType() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPDatabaseInstance) pulumi.StringOutput { return v.DatabaseType }).(pulumi.StringOutput)
}

func (o SAPDatabaseInstanceOutput) Errors() SAPVirtualInstanceErrorResponseOutput {
	return o.ApplyT(func(v *SAPDatabaseInstance) SAPVirtualInstanceErrorResponseOutput { return v.Errors }).(SAPVirtualInstanceErrorResponseOutput)
}

func (o SAPDatabaseInstanceOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPDatabaseInstance) pulumi.StringOutput { return v.IpAddress }).(pulumi.StringOutput)
}

func (o SAPDatabaseInstanceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPDatabaseInstance) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o SAPDatabaseInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPDatabaseInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SAPDatabaseInstanceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPDatabaseInstance) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o SAPDatabaseInstanceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPDatabaseInstance) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o SAPDatabaseInstanceOutput) Subnet() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPDatabaseInstance) pulumi.StringOutput { return v.Subnet }).(pulumi.StringOutput)
}

func (o SAPDatabaseInstanceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SAPDatabaseInstance) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SAPDatabaseInstanceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SAPDatabaseInstance) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SAPDatabaseInstanceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPDatabaseInstance) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o SAPDatabaseInstanceOutput) VmDetails() DatabaseVmDetailsResponseArrayOutput {
	return o.ApplyT(func(v *SAPDatabaseInstance) DatabaseVmDetailsResponseArrayOutput { return v.VmDetails }).(DatabaseVmDetailsResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(SAPDatabaseInstanceOutput{})
}
