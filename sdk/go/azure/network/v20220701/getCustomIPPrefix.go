


package v20220701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCustomIPPrefix(ctx *pulumi.Context, args *LookupCustomIPPrefixArgs, opts ...pulumi.InvokeOption) (*LookupCustomIPPrefixResult, error) {
	var rv LookupCustomIPPrefixResult
	err := ctx.Invoke("azure-native:network/v20220701:getCustomIPPrefix", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCustomIPPrefixArgs struct {
	CustomIpPrefixName string  `pulumi:"customIpPrefixName"`
	Expand             *string `pulumi:"expand"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
}


type LookupCustomIPPrefixResult struct {
	Asn                   *string                   `pulumi:"asn"`
	AuthorizationMessage  *string                   `pulumi:"authorizationMessage"`
	ChildCustomIpPrefixes []SubResourceResponse     `pulumi:"childCustomIpPrefixes"`
	Cidr                  *string                   `pulumi:"cidr"`
	CommissionedState     *string                   `pulumi:"commissionedState"`
	CustomIpPrefixParent  *SubResourceResponse      `pulumi:"customIpPrefixParent"`
	Etag                  string                    `pulumi:"etag"`
	ExpressRouteAdvertise *bool                     `pulumi:"expressRouteAdvertise"`
	ExtendedLocation      *ExtendedLocationResponse `pulumi:"extendedLocation"`
	FailedReason          string                    `pulumi:"failedReason"`
	Geo                   *string                   `pulumi:"geo"`
	Id                    *string                   `pulumi:"id"`
	Location              *string                   `pulumi:"location"`
	Name                  string                    `pulumi:"name"`
	NoInternetAdvertise   *bool                     `pulumi:"noInternetAdvertise"`
	PrefixType            *string                   `pulumi:"prefixType"`
	ProvisioningState     string                    `pulumi:"provisioningState"`
	PublicIpPrefixes      []SubResourceResponse     `pulumi:"publicIpPrefixes"`
	ResourceGuid          string                    `pulumi:"resourceGuid"`
	SignedMessage         *string                   `pulumi:"signedMessage"`
	Tags                  map[string]string         `pulumi:"tags"`
	Type                  string                    `pulumi:"type"`
	Zones                 []string                  `pulumi:"zones"`
}

func LookupCustomIPPrefixOutput(ctx *pulumi.Context, args LookupCustomIPPrefixOutputArgs, opts ...pulumi.InvokeOption) LookupCustomIPPrefixResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCustomIPPrefixResult, error) {
			args := v.(LookupCustomIPPrefixArgs)
			r, err := LookupCustomIPPrefix(ctx, &args, opts...)
			var s LookupCustomIPPrefixResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCustomIPPrefixResultOutput)
}

type LookupCustomIPPrefixOutputArgs struct {
	CustomIpPrefixName pulumi.StringInput    `pulumi:"customIpPrefixName"`
	Expand             pulumi.StringPtrInput `pulumi:"expand"`
	ResourceGroupName  pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupCustomIPPrefixOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomIPPrefixArgs)(nil)).Elem()
}


type LookupCustomIPPrefixResultOutput struct{ *pulumi.OutputState }

func (LookupCustomIPPrefixResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomIPPrefixResult)(nil)).Elem()
}

func (o LookupCustomIPPrefixResultOutput) ToLookupCustomIPPrefixResultOutput() LookupCustomIPPrefixResultOutput {
	return o
}

func (o LookupCustomIPPrefixResultOutput) ToLookupCustomIPPrefixResultOutputWithContext(ctx context.Context) LookupCustomIPPrefixResultOutput {
	return o
}

func (o LookupCustomIPPrefixResultOutput) Asn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) *string { return v.Asn }).(pulumi.StringPtrOutput)
}

func (o LookupCustomIPPrefixResultOutput) AuthorizationMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) *string { return v.AuthorizationMessage }).(pulumi.StringPtrOutput)
}

func (o LookupCustomIPPrefixResultOutput) ChildCustomIpPrefixes() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) []SubResourceResponse { return v.ChildCustomIpPrefixes }).(SubResourceResponseArrayOutput)
}

func (o LookupCustomIPPrefixResultOutput) Cidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) *string { return v.Cidr }).(pulumi.StringPtrOutput)
}

func (o LookupCustomIPPrefixResultOutput) CommissionedState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) *string { return v.CommissionedState }).(pulumi.StringPtrOutput)
}

func (o LookupCustomIPPrefixResultOutput) CustomIpPrefixParent() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) *SubResourceResponse { return v.CustomIpPrefixParent }).(SubResourceResponsePtrOutput)
}

func (o LookupCustomIPPrefixResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupCustomIPPrefixResultOutput) ExpressRouteAdvertise() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) *bool { return v.ExpressRouteAdvertise }).(pulumi.BoolPtrOutput)
}

func (o LookupCustomIPPrefixResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o LookupCustomIPPrefixResultOutput) FailedReason() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) string { return v.FailedReason }).(pulumi.StringOutput)
}

func (o LookupCustomIPPrefixResultOutput) Geo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) *string { return v.Geo }).(pulumi.StringPtrOutput)
}

func (o LookupCustomIPPrefixResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupCustomIPPrefixResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupCustomIPPrefixResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCustomIPPrefixResultOutput) NoInternetAdvertise() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) *bool { return v.NoInternetAdvertise }).(pulumi.BoolPtrOutput)
}

func (o LookupCustomIPPrefixResultOutput) PrefixType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) *string { return v.PrefixType }).(pulumi.StringPtrOutput)
}

func (o LookupCustomIPPrefixResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupCustomIPPrefixResultOutput) PublicIpPrefixes() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) []SubResourceResponse { return v.PublicIpPrefixes }).(SubResourceResponseArrayOutput)
}

func (o LookupCustomIPPrefixResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o LookupCustomIPPrefixResultOutput) SignedMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) *string { return v.SignedMessage }).(pulumi.StringPtrOutput)
}

func (o LookupCustomIPPrefixResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupCustomIPPrefixResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupCustomIPPrefixResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCustomIPPrefixResultOutput{})
}
