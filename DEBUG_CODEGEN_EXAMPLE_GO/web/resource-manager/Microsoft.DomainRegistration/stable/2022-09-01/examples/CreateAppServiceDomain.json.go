package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/domainregistration/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := domainregistration.NewDomain(ctx, "domain", &domainregistration.DomainArgs{
			AuthCode:  pulumi.String("exampleAuthCode"),
			AutoRenew: pulumi.Bool(true),
			Consent: &domainregistration.DomainPurchaseConsentArgs{
				AgreedAt: pulumi.String("2021-09-10T19:30:53Z"),
				AgreedBy: pulumi.String("192.0.2.1"),
				AgreementKeys: pulumi.StringArray{
					pulumi.String("agreementKey1"),
				},
			},
			ContactAdmin: &domainregistration.ContactArgs{
				AddressMailing: &domainregistration.AddressArgs{
					Address1:   pulumi.String("3400 State St"),
					City:       pulumi.String("Chicago"),
					Country:    pulumi.String("United States"),
					PostalCode: pulumi.String("67098"),
					State:      pulumi.String("IL"),
				},
				Email:        pulumi.String("admin@email.com"),
				Fax:          pulumi.String("1-245-534-2242"),
				JobTitle:     pulumi.String("Admin"),
				NameFirst:    pulumi.String("John"),
				NameLast:     pulumi.String("Doe"),
				NameMiddle:   pulumi.String(""),
				Organization: pulumi.String("Microsoft Inc."),
				Phone:        pulumi.String("1-245-534-2242"),
			},
			ContactBilling: &domainregistration.ContactArgs{
				AddressMailing: &domainregistration.AddressArgs{
					Address1:   pulumi.String("3400 State St"),
					City:       pulumi.String("Chicago"),
					Country:    pulumi.String("United States"),
					PostalCode: pulumi.String("67098"),
					State:      pulumi.String("IL"),
				},
				Email:        pulumi.String("billing@email.com"),
				Fax:          pulumi.String("1-245-534-2242"),
				JobTitle:     pulumi.String("Billing"),
				NameFirst:    pulumi.String("John"),
				NameLast:     pulumi.String("Doe"),
				NameMiddle:   pulumi.String(""),
				Organization: pulumi.String("Microsoft Inc."),
				Phone:        pulumi.String("1-245-534-2242"),
			},
			ContactRegistrant: &domainregistration.ContactArgs{
				AddressMailing: &domainregistration.AddressArgs{
					Address1:   pulumi.String("3400 State St"),
					City:       pulumi.String("Chicago"),
					Country:    pulumi.String("United States"),
					PostalCode: pulumi.String("67098"),
					State:      pulumi.String("IL"),
				},
				Email:        pulumi.String("registrant@email.com"),
				Fax:          pulumi.String("1-245-534-2242"),
				JobTitle:     pulumi.String("Registrant"),
				NameFirst:    pulumi.String("John"),
				NameLast:     pulumi.String("Doe"),
				NameMiddle:   pulumi.String(""),
				Organization: pulumi.String("Microsoft Inc."),
				Phone:        pulumi.String("1-245-534-2242"),
			},
			ContactTech: &domainregistration.ContactArgs{
				AddressMailing: &domainregistration.AddressArgs{
					Address1:   pulumi.String("3400 State St"),
					City:       pulumi.String("Chicago"),
					Country:    pulumi.String("United States"),
					PostalCode: pulumi.String("67098"),
					State:      pulumi.String("IL"),
				},
				Email:        pulumi.String("tech@email.com"),
				Fax:          pulumi.String("1-245-534-2242"),
				JobTitle:     pulumi.String("Tech"),
				NameFirst:    pulumi.String("John"),
				NameLast:     pulumi.String("Doe"),
				NameMiddle:   pulumi.String(""),
				Organization: pulumi.String("Microsoft Inc."),
				Phone:        pulumi.String("1-245-534-2242"),
			},
			DnsType:           domainregistration.DnsTypeDefaultDomainRegistrarDns,
			DomainName:        pulumi.String("example.com"),
			Location:          pulumi.String("global"),
			Privacy:           pulumi.Bool(false),
			ResourceGroupName: pulumi.String("testrg123"),
			Tags:              nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
