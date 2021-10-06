


package v20200201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRegistrationAssignment(ctx *pulumi.Context, args *LookupRegistrationAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupRegistrationAssignmentResult, error) {
	var rv LookupRegistrationAssignmentResult
	err := ctx.Invoke("azure-native:managedservices/v20200201preview:getRegistrationAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRegistrationAssignmentArgs struct {
	ExpandRegistrationDefinition *bool  `pulumi:"expandRegistrationDefinition"`
	RegistrationAssignmentId     string `pulumi:"registrationAssignmentId"`
	Scope                        string `pulumi:"scope"`
}


type LookupRegistrationAssignmentResult struct {
	Id         string                                   `pulumi:"id"`
	Name       string                                   `pulumi:"name"`
	Properties RegistrationAssignmentPropertiesResponse `pulumi:"properties"`
	Type       string                                   `pulumi:"type"`
}
