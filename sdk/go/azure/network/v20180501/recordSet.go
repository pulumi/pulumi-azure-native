


package v20180501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RecordSet struct {
	pulumi.CustomResourceState

	ARecords          ARecordResponseArrayOutput    `pulumi:"aRecords"`
	AaaaRecords       AaaaRecordResponseArrayOutput `pulumi:"aaaaRecords"`
	CaaRecords        CaaRecordResponseArrayOutput  `pulumi:"caaRecords"`
	CnameRecord       CnameRecordResponsePtrOutput  `pulumi:"cnameRecord"`
	Etag              pulumi.StringPtrOutput        `pulumi:"etag"`
	Fqdn              pulumi.StringOutput           `pulumi:"fqdn"`
	Metadata          pulumi.StringMapOutput        `pulumi:"metadata"`
	MxRecords         MxRecordResponseArrayOutput   `pulumi:"mxRecords"`
	Name              pulumi.StringOutput           `pulumi:"name"`
	NsRecords         NsRecordResponseArrayOutput   `pulumi:"nsRecords"`
	ProvisioningState pulumi.StringOutput           `pulumi:"provisioningState"`
	PtrRecords        PtrRecordResponseArrayOutput  `pulumi:"ptrRecords"`
	SoaRecord         SoaRecordResponsePtrOutput    `pulumi:"soaRecord"`
	SrvRecords        SrvRecordResponseArrayOutput  `pulumi:"srvRecords"`
	TargetResource    SubResourceResponsePtrOutput  `pulumi:"targetResource"`
	Ttl               pulumi.Float64PtrOutput       `pulumi:"ttl"`
	TxtRecords        TxtRecordResponseArrayOutput  `pulumi:"txtRecords"`
	Type              pulumi.StringOutput           `pulumi:"type"`
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
			Type: pulumi.String("azure-native:network/v20171001:RecordSet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180301preview:RecordSet"),
		},
	})
	opts = append(opts, aliases)
	var resource RecordSet
	err := ctx.RegisterResource("azure-native:network/v20180501:RecordSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRecordSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RecordSetState, opts ...pulumi.ResourceOption) (*RecordSet, error) {
	var resource RecordSet
	err := ctx.ReadResource("azure-native:network/v20180501:RecordSet", name, id, state, &resource, opts...)
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
	Metadata              map[string]string `pulumi:"metadata"`
	MxRecords             []MxRecord        `pulumi:"mxRecords"`
	NsRecords             []NsRecord        `pulumi:"nsRecords"`
	PtrRecords            []PtrRecord       `pulumi:"ptrRecords"`
	RecordType            string            `pulumi:"recordType"`
	RelativeRecordSetName *string           `pulumi:"relativeRecordSetName"`
	ResourceGroupName     string            `pulumi:"resourceGroupName"`
	SoaRecord             *SoaRecord        `pulumi:"soaRecord"`
	SrvRecords            []SrvRecord       `pulumi:"srvRecords"`
	TargetResource        *SubResource      `pulumi:"targetResource"`
	Ttl                   *float64          `pulumi:"ttl"`
	TxtRecords            []TxtRecord       `pulumi:"txtRecords"`
	ZoneName              string            `pulumi:"zoneName"`
}


type RecordSetArgs struct {
	ARecords              ARecordArrayInput
	AaaaRecords           AaaaRecordArrayInput
	CaaRecords            CaaRecordArrayInput
	CnameRecord           CnameRecordPtrInput
	Metadata              pulumi.StringMapInput
	MxRecords             MxRecordArrayInput
	NsRecords             NsRecordArrayInput
	PtrRecords            PtrRecordArrayInput
	RecordType            pulumi.StringInput
	RelativeRecordSetName pulumi.StringPtrInput
	ResourceGroupName     pulumi.StringInput
	SoaRecord             SoaRecordPtrInput
	SrvRecords            SrvRecordArrayInput
	TargetResource        SubResourcePtrInput
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
	return reflect.TypeOf((**RecordSet)(nil)).Elem()
}

func (i *RecordSet) ToRecordSetOutput() RecordSetOutput {
	return i.ToRecordSetOutputWithContext(context.Background())
}

func (i *RecordSet) ToRecordSetOutputWithContext(ctx context.Context) RecordSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecordSetOutput)
}

type RecordSetOutput struct{ *pulumi.OutputState }

func (RecordSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RecordSet)(nil)).Elem()
}

func (o RecordSetOutput) ToRecordSetOutput() RecordSetOutput {
	return o
}

func (o RecordSetOutput) ToRecordSetOutputWithContext(ctx context.Context) RecordSetOutput {
	return o
}

func (o RecordSetOutput) ARecords() ARecordResponseArrayOutput {
	return o.ApplyT(func(v *RecordSet) ARecordResponseArrayOutput { return v.ARecords }).(ARecordResponseArrayOutput)
}

func (o RecordSetOutput) AaaaRecords() AaaaRecordResponseArrayOutput {
	return o.ApplyT(func(v *RecordSet) AaaaRecordResponseArrayOutput { return v.AaaaRecords }).(AaaaRecordResponseArrayOutput)
}

func (o RecordSetOutput) CaaRecords() CaaRecordResponseArrayOutput {
	return o.ApplyT(func(v *RecordSet) CaaRecordResponseArrayOutput { return v.CaaRecords }).(CaaRecordResponseArrayOutput)
}

func (o RecordSetOutput) CnameRecord() CnameRecordResponsePtrOutput {
	return o.ApplyT(func(v *RecordSet) CnameRecordResponsePtrOutput { return v.CnameRecord }).(CnameRecordResponsePtrOutput)
}

func (o RecordSetOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RecordSet) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o RecordSetOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v *RecordSet) pulumi.StringOutput { return v.Fqdn }).(pulumi.StringOutput)
}

func (o RecordSetOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RecordSet) pulumi.StringMapOutput { return v.Metadata }).(pulumi.StringMapOutput)
}

func (o RecordSetOutput) MxRecords() MxRecordResponseArrayOutput {
	return o.ApplyT(func(v *RecordSet) MxRecordResponseArrayOutput { return v.MxRecords }).(MxRecordResponseArrayOutput)
}

func (o RecordSetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RecordSet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RecordSetOutput) NsRecords() NsRecordResponseArrayOutput {
	return o.ApplyT(func(v *RecordSet) NsRecordResponseArrayOutput { return v.NsRecords }).(NsRecordResponseArrayOutput)
}

func (o RecordSetOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *RecordSet) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o RecordSetOutput) PtrRecords() PtrRecordResponseArrayOutput {
	return o.ApplyT(func(v *RecordSet) PtrRecordResponseArrayOutput { return v.PtrRecords }).(PtrRecordResponseArrayOutput)
}

func (o RecordSetOutput) SoaRecord() SoaRecordResponsePtrOutput {
	return o.ApplyT(func(v *RecordSet) SoaRecordResponsePtrOutput { return v.SoaRecord }).(SoaRecordResponsePtrOutput)
}

func (o RecordSetOutput) SrvRecords() SrvRecordResponseArrayOutput {
	return o.ApplyT(func(v *RecordSet) SrvRecordResponseArrayOutput { return v.SrvRecords }).(SrvRecordResponseArrayOutput)
}

func (o RecordSetOutput) TargetResource() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *RecordSet) SubResourceResponsePtrOutput { return v.TargetResource }).(SubResourceResponsePtrOutput)
}

func (o RecordSetOutput) Ttl() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *RecordSet) pulumi.Float64PtrOutput { return v.Ttl }).(pulumi.Float64PtrOutput)
}

func (o RecordSetOutput) TxtRecords() TxtRecordResponseArrayOutput {
	return o.ApplyT(func(v *RecordSet) TxtRecordResponseArrayOutput { return v.TxtRecords }).(TxtRecordResponseArrayOutput)
}

func (o RecordSetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RecordSet) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RecordSetOutput{})
}
