


package v20171001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RecordSet struct {
	pulumi.CustomResourceState

	ARecords    ARecordResponseArrayOutput    `pulumi:"aRecords"`
	AaaaRecords AaaaRecordResponseArrayOutput `pulumi:"aaaaRecords"`
	CaaRecords  CaaRecordResponseArrayOutput  `pulumi:"caaRecords"`
	CnameRecord CnameRecordResponsePtrOutput  `pulumi:"cnameRecord"`
	Etag        pulumi.StringPtrOutput        `pulumi:"etag"`
	Fqdn        pulumi.StringOutput           `pulumi:"fqdn"`
	Metadata    pulumi.StringMapOutput        `pulumi:"metadata"`
	MxRecords   MxRecordResponseArrayOutput   `pulumi:"mxRecords"`
	Name        pulumi.StringOutput           `pulumi:"name"`
	NsRecords   NsRecordResponseArrayOutput   `pulumi:"nsRecords"`
	PtrRecords  PtrRecordResponseArrayOutput  `pulumi:"ptrRecords"`
	SoaRecord   SoaRecordResponsePtrOutput    `pulumi:"soaRecord"`
	SrvRecords  SrvRecordResponseArrayOutput  `pulumi:"srvRecords"`
	Ttl         pulumi.Float64PtrOutput       `pulumi:"ttl"`
	TxtRecords  TxtRecordResponseArrayOutput  `pulumi:"txtRecords"`
	Type        pulumi.StringOutput           `pulumi:"type"`
}


func NewRecordSet(ctx *pulumi.Context,
	name string, args *RecordSetArgs, opts ...pulumi.ResourceOption) (*RecordSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RecordType == nil {
		return nil, errors.New("invalid value for required argument 'RecordType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ZoneName == nil {
		return nil, errors.New("invalid value for required argument 'ZoneName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:RecordSet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150504preview:RecordSet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160401:RecordSet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:RecordSet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180301preview:RecordSet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180501:RecordSet"),
		},
	})
	opts = append(opts, aliases)
	var resource RecordSet
	err := ctx.RegisterResource("azure-native:network/v20171001:RecordSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRecordSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RecordSetState, opts ...pulumi.ResourceOption) (*RecordSet, error) {
	var resource RecordSet
	err := ctx.ReadResource("azure-native:network/v20171001:RecordSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type recordSetState struct {
}

type RecordSetState struct {
}

func (RecordSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*recordSetState)(nil)).Elem()
}

type recordSetArgs struct {
	ARecords              []ARecord         `pulumi:"aRecords"`
	AaaaRecords           []AaaaRecord      `pulumi:"aaaaRecords"`
	CaaRecords            []CaaRecord       `pulumi:"caaRecords"`
	CnameRecord           *CnameRecord      `pulumi:"cnameRecord"`
	Etag                  *string           `pulumi:"etag"`
	Metadata              map[string]string `pulumi:"metadata"`
	MxRecords             []MxRecord        `pulumi:"mxRecords"`
	NsRecords             []NsRecord        `pulumi:"nsRecords"`
	PtrRecords            []PtrRecord       `pulumi:"ptrRecords"`
	RecordType            string            `pulumi:"recordType"`
	RelativeRecordSetName *string           `pulumi:"relativeRecordSetName"`
	ResourceGroupName     string            `pulumi:"resourceGroupName"`
	SoaRecord             *SoaRecord        `pulumi:"soaRecord"`
	SrvRecords            []SrvRecord       `pulumi:"srvRecords"`
	Ttl                   *float64          `pulumi:"ttl"`
	TxtRecords            []TxtRecord       `pulumi:"txtRecords"`
	ZoneName              string            `pulumi:"zoneName"`
}


type RecordSetArgs struct {
	ARecords              ARecordArrayInput
	AaaaRecords           AaaaRecordArrayInput
	CaaRecords            CaaRecordArrayInput
	CnameRecord           CnameRecordPtrInput
	Etag                  pulumi.StringPtrInput
	Metadata              pulumi.StringMapInput
	MxRecords             MxRecordArrayInput
	NsRecords             NsRecordArrayInput
	PtrRecords            PtrRecordArrayInput
	RecordType            pulumi.StringInput
	RelativeRecordSetName pulumi.StringPtrInput
	ResourceGroupName     pulumi.StringInput
	SoaRecord             SoaRecordPtrInput
	SrvRecords            SrvRecordArrayInput
	Ttl                   pulumi.Float64PtrInput
	TxtRecords            TxtRecordArrayInput
	ZoneName              pulumi.StringInput
}

func (RecordSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*recordSetArgs)(nil)).Elem()
}

type RecordSetInput interface {
	pulumi.Input

	ToRecordSetOutput() RecordSetOutput
	ToRecordSetOutputWithContext(ctx context.Context) RecordSetOutput
}

func (*RecordSet) ElementType() reflect.Type {
	return reflect.TypeOf((*RecordSet)(nil))
}

func (i *RecordSet) ToRecordSetOutput() RecordSetOutput {
	return i.ToRecordSetOutputWithContext(context.Background())
}

func (i *RecordSet) ToRecordSetOutputWithContext(ctx context.Context) RecordSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecordSetOutput)
}

type RecordSetOutput struct{ *pulumi.OutputState }

func (RecordSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecordSet)(nil))
}

func (o RecordSetOutput) ToRecordSetOutput() RecordSetOutput {
	return o
}

func (o RecordSetOutput) ToRecordSetOutputWithContext(ctx context.Context) RecordSetOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RecordSetOutput{})
}
