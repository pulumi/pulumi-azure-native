


package v20210404preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSpacecraftAvailableContacts(ctx *pulumi.Context, args *ListSpacecraftAvailableContactsArgs, opts ...pulumi.InvokeOption) (*ListSpacecraftAvailableContactsResult, error) {
	var rv ListSpacecraftAvailableContactsResult
	err := ctx.Invoke("azure-native:orbital/v20210404preview:listSpacecraftAvailableContacts", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSpacecraftAvailableContactsArgs struct {
	ContactProfile    ResourceReference `pulumi:"contactProfile"`
	EndTime           string            `pulumi:"endTime"`
	GroundStationName string            `pulumi:"groundStationName"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	SpacecraftName    string            `pulumi:"spacecraftName"`
	StartTime         string            `pulumi:"startTime"`
}


type ListSpacecraftAvailableContactsResult struct {
	NextLink string                      `pulumi:"nextLink"`
	Value    []AvailableContactsResponse `pulumi:"value"`
}

func ListSpacecraftAvailableContactsOutput(ctx *pulumi.Context, args ListSpacecraftAvailableContactsOutputArgs, opts ...pulumi.InvokeOption) ListSpacecraftAvailableContactsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSpacecraftAvailableContactsResult, error) {
			args := v.(ListSpacecraftAvailableContactsArgs)
			r, err := ListSpacecraftAvailableContacts(ctx, &args, opts...)
			var s ListSpacecraftAvailableContactsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSpacecraftAvailableContactsResultOutput)
}

type ListSpacecraftAvailableContactsOutputArgs struct {
	ContactProfile    ResourceReferenceInput `pulumi:"contactProfile"`
	EndTime           pulumi.StringInput     `pulumi:"endTime"`
	GroundStationName pulumi.StringInput     `pulumi:"groundStationName"`
	ResourceGroupName pulumi.StringInput     `pulumi:"resourceGroupName"`
	SpacecraftName    pulumi.StringInput     `pulumi:"spacecraftName"`
	StartTime         pulumi.StringInput     `pulumi:"startTime"`
}

func (ListSpacecraftAvailableContactsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSpacecraftAvailableContactsArgs)(nil)).Elem()
}


type ListSpacecraftAvailableContactsResultOutput struct{ *pulumi.OutputState }

func (ListSpacecraftAvailableContactsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSpacecraftAvailableContactsResult)(nil)).Elem()
}

func (o ListSpacecraftAvailableContactsResultOutput) ToListSpacecraftAvailableContactsResultOutput() ListSpacecraftAvailableContactsResultOutput {
	return o
}

func (o ListSpacecraftAvailableContactsResultOutput) ToListSpacecraftAvailableContactsResultOutputWithContext(ctx context.Context) ListSpacecraftAvailableContactsResultOutput {
	return o
}

func (o ListSpacecraftAvailableContactsResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListSpacecraftAvailableContactsResult) string { return v.NextLink }).(pulumi.StringOutput)
}

func (o ListSpacecraftAvailableContactsResultOutput) Value() AvailableContactsResponseArrayOutput {
	return o.ApplyT(func(v ListSpacecraftAvailableContactsResult) []AvailableContactsResponse { return v.Value }).(AvailableContactsResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSpacecraftAvailableContactsResultOutput{})
}
