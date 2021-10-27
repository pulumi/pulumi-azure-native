


package v20150504preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RecordSet struct {
	pulumi.CustomResourceState

	AAAARecords AaaaRecordResponseArrayOutput `pulumi:"aAAARecords"`
	ARecords    ARecordResponseArrayOutput    `pulumi:"aRecords"`
	CNAMERecord CnameRecordResponsePtrOutput  `pulumi:"cNAMERecord"`
	Etag        pulumi.StringPtrOutput        `pulumi:"etag"`
	Fqdn        pulumi.StringOutput           `pulumi:"fqdn"`
	MXRecords   MxRecordResponseArrayOutput   `pulumi:"mXRecords"`
	NSRecords   NsRecordResponseArrayOutput   `pulumi:"nSRecords"`
	Name        pulumi.StringOutput           `pulumi:"name"`
	PTRRecords  PtrRecordResponseArrayOutput  `pulumi:"pTRRecords"`
	SOARecord   SoaRecordResponsePtrOutput    `pulumi:"sOARecord"`
	SRVRecords  SrvRecordResponseArrayOutput  `pulumi:"sRVRecords"`
	TXTRecords  TxtRecordResponseArrayOutput  `pulumi:"tXTRecords"`
	Ttl         pulumi.Float64PtrOutput       `pulumi:"ttl"`
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
			Type: pulumi.String("azure-nextgen:network/v20150504preview:RecordSet"),
		},
		{
			Type: pulumi.String("azure-native:network:RecordSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network:RecordSet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160401:RecordSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20160401:RecordSet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:RecordSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20170901:RecordSet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:RecordSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20171001:RecordSet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180301preview:RecordSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180301preview:RecordSet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180501:RecordSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180501:RecordSet"),
		},
	})
	opts = append(opts, aliases)
	var resource RecordSet
	err := ctx.RegisterResource("azure-native:network/v20150504preview:RecordSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRecordSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RecordSetState, opts ...pulumi.ResourceOption) (*RecordSet, error) {
	var resource RecordSet
	err := ctx.ReadResource("azure-native:network/v20150504preview:RecordSet", name, id, state, &resource, opts...)
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
	AAAARecords           []AaaaRecord `pulumi:"aAAARecords"`
	ARecords              []ARecord    `pulumi:"aRecords"`
	CNAMERecord           *CnameRecord `pulumi:"cNAMERecord"`
	Etag                  *string      `pulumi:"etag"`
	MXRecords             []MxRecord   `pulumi:"mXRecords"`
	NSRecords             []NsRecord   `pulumi:"nSRecords"`
	PTRRecords            []PtrRecord  `pulumi:"pTRRecords"`
	RecordType            string       `pulumi:"recordType"`
	RelativeRecordSetName *string      `pulumi:"relativeRecordSetName"`
	ResourceGroupName     string       `pulumi:"resourceGroupName"`
	SOARecord             *SoaRecord   `pulumi:"sOARecord"`
	SRVRecords            []SrvRecord  `pulumi:"sRVRecords"`
	TXTRecords            []TxtRecord  `pulumi:"tXTRecords"`
	Ttl                   *float64     `pulumi:"ttl"`
	ZoneName              string       `pulumi:"zoneName"`
}


type RecordSetArgs struct {
	AAAARecords           AaaaRecordArrayInput
	ARecords              ARecordArrayInput
	CNAMERecord           CnameRecordPtrInput
	Etag                  pulumi.StringPtrInput
	MXRecords             MxRecordArrayInput
	NSRecords             NsRecordArrayInput
	PTRRecords            PtrRecordArrayInput
	RecordType            pulumi.StringInput
	RelativeRecordSetName pulumi.StringPtrInput
	ResourceGroupName     pulumi.StringInput
	SOARecord             SoaRecordPtrInput
	SRVRecords            SrvRecordArrayInput
	TXTRecords            TxtRecordArrayInput
	Ttl                   pulumi.Float64PtrInput
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
	pulumi.RegisterInputType(reflect.TypeOf((*RecordSetInput)(nil)).Elem(), &RecordSet{})
	pulumi.RegisterOutputType(RecordSetOutput{})
}
