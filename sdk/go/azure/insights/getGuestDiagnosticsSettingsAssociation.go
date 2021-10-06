


package insights

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGuestDiagnosticsSettingsAssociation(ctx *pulumi.Context, args *LookupGuestDiagnosticsSettingsAssociationArgs, opts ...pulumi.InvokeOption) (*LookupGuestDiagnosticsSettingsAssociationResult, error) {
	var rv LookupGuestDiagnosticsSettingsAssociationResult
	err := ctx.Invoke("azure-native:insights:getGuestDiagnosticsSettingsAssociation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGuestDiagnosticsSettingsAssociationArgs struct {
	AssociationName string `pulumi:"associationName"`
	ResourceUri     string `pulumi:"resourceUri"`
}


type LookupGuestDiagnosticsSettingsAssociationResult struct {
	GuestDiagnosticSettingsName string            `pulumi:"guestDiagnosticSettingsName"`
	Id                          string            `pulumi:"id"`
	Location                    string            `pulumi:"location"`
	Name                        string            `pulumi:"name"`
	Tags                        map[string]string `pulumi:"tags"`
	Type                        string            `pulumi:"type"`
}
