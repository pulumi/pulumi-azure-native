


package managedservices

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRegistrationDefinition(ctx *pulumi.Context, args *LookupRegistrationDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupRegistrationDefinitionResult, error) {
	var rv LookupRegistrationDefinitionResult
	err := ctx.Invoke("azure-native:managedservices:getRegistrationDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRegistrationDefinitionArgs struct {
	RegistrationDefinitionId string `pulumi:"registrationDefinitionId"`
	Scope                    string `pulumi:"scope"`
}


type LookupRegistrationDefinitionResult struct {
	Id         string                                   `pulumi:"id"`
	Name       string                                   `pulumi:"name"`
	Plan       *PlanResponse                            `pulumi:"plan"`
	Properties RegistrationDefinitionPropertiesResponse `pulumi:"properties"`
	Type       string                                   `pulumi:"type"`
}
