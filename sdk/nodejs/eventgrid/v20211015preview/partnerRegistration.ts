// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Information about a partner registration.
 */
export class PartnerRegistration extends pulumi.CustomResource {
    /**
     * Get an existing PartnerRegistration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PartnerRegistration {
        return new PartnerRegistration(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:eventgrid/v20211015preview:PartnerRegistration';

    /**
     * Returns true if the given object is an instance of PartnerRegistration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PartnerRegistration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PartnerRegistration.__pulumiType;
    }

    /**
     * List of Azure subscription Ids that are authorized to create a partner namespace
     * associated with this partner registration. This is an optional property. Creating
     * partner namespaces is always permitted under the same Azure subscription as the one used
     * for creating the partner registration.
     */
    public readonly authorizedAzureSubscriptionIds!: pulumi.Output<string[] | undefined>;
    /**
     * The extension of the customer service URI of the publisher.
     */
    public readonly customerServiceUri!: pulumi.Output<string | undefined>;
    /**
     * Location of the resource.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * URI of the logo.
     */
    public readonly logoUri!: pulumi.Output<string | undefined>;
    /**
     * Long description for the custom scenarios and integration to be displayed in the portal if needed.
     * Length of this description should not exceed 2048 characters.
     */
    public readonly longDescription!: pulumi.Output<string | undefined>;
    /**
     * Name of the resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The extension of the customer service number of the publisher. Only digits are allowed and number of digits should not exceed 10.
     */
    public readonly partnerCustomerServiceExtension!: pulumi.Output<string | undefined>;
    /**
     * The customer service number of the publisher. The expected phone format should start with a '+' sign 
     * followed by the country code. The remaining digits are then followed. Only digits and spaces are allowed and its
     * length cannot exceed 16 digits including country code. Examples of valid phone numbers are: +1 515 123 4567 and
     * +966 7 5115 2471. Examples of invalid phone numbers are: +1 (515) 123-4567, 1 515 123 4567 and +966 121 5115 24 7 551 1234 43
     */
    public readonly partnerCustomerServiceNumber!: pulumi.Output<string | undefined>;
    /**
     * Official name of the partner name. For example: "Contoso".
     */
    public readonly partnerName!: pulumi.Output<string | undefined>;
    /**
     * The immutableId of the corresponding partner registration.
     */
    public readonly partnerRegistrationImmutableId!: pulumi.Output<string | undefined>;
    /**
     * Short description of the partner resource type. The length of this description should not exceed 256 characters.
     */
    public readonly partnerResourceTypeDescription!: pulumi.Output<string | undefined>;
    /**
     * Display name of the partner resource type.
     */
    public readonly partnerResourceTypeDisplayName!: pulumi.Output<string | undefined>;
    /**
     * Name of the partner resource type.
     */
    public readonly partnerResourceTypeName!: pulumi.Output<string | undefined>;
    /**
     * Provisioning state of the partner registration.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * URI of the partner website that can be used by Azure customers to setup Event Grid
     * integration on an event source.
     */
    public readonly setupUri!: pulumi.Output<string | undefined>;
    /**
     * The system metadata relating to Partner Registration resource.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.eventgrid.v20211015preview.SystemDataResponse>;
    /**
     * Tags of the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Visibility state of the partner registration.
     */
    public readonly visibilityState!: pulumi.Output<string | undefined>;

    /**
     * Create a PartnerRegistration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PartnerRegistrationArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            resourceInputs["authorizedAzureSubscriptionIds"] = args ? args.authorizedAzureSubscriptionIds : undefined;
            resourceInputs["customerServiceUri"] = args ? args.customerServiceUri : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["logoUri"] = args ? args.logoUri : undefined;
            resourceInputs["longDescription"] = args ? args.longDescription : undefined;
            resourceInputs["partnerCustomerServiceExtension"] = args ? args.partnerCustomerServiceExtension : undefined;
            resourceInputs["partnerCustomerServiceNumber"] = args ? args.partnerCustomerServiceNumber : undefined;
            resourceInputs["partnerName"] = args ? args.partnerName : undefined;
            resourceInputs["partnerRegistrationImmutableId"] = args ? args.partnerRegistrationImmutableId : undefined;
            resourceInputs["partnerRegistrationName"] = args ? args.partnerRegistrationName : undefined;
            resourceInputs["partnerResourceTypeDescription"] = args ? args.partnerResourceTypeDescription : undefined;
            resourceInputs["partnerResourceTypeDisplayName"] = args ? args.partnerResourceTypeDisplayName : undefined;
            resourceInputs["partnerResourceTypeName"] = args ? args.partnerResourceTypeName : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["setupUri"] = args ? args.setupUri : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["visibilityState"] = args ? args.visibilityState : undefined;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["authorizedAzureSubscriptionIds"] = undefined /*out*/;
            resourceInputs["customerServiceUri"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["logoUri"] = undefined /*out*/;
            resourceInputs["longDescription"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["partnerCustomerServiceExtension"] = undefined /*out*/;
            resourceInputs["partnerCustomerServiceNumber"] = undefined /*out*/;
            resourceInputs["partnerName"] = undefined /*out*/;
            resourceInputs["partnerRegistrationImmutableId"] = undefined /*out*/;
            resourceInputs["partnerResourceTypeDescription"] = undefined /*out*/;
            resourceInputs["partnerResourceTypeDisplayName"] = undefined /*out*/;
            resourceInputs["partnerResourceTypeName"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["setupUri"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["visibilityState"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:eventgrid/v20200401preview:PartnerRegistration" }, { type: "azure-native:eventgrid/v20201015preview:PartnerRegistration" }, { type: "azure-native:eventgrid/v20210601preview:PartnerRegistration" }, { type: "azure-native:eventgrid/v20220615:PartnerRegistration" }, { type: "azure-native:eventgrid/v20230601preview:PartnerRegistration" }, { type: "azure-native:eventgrid/v20231215preview:PartnerRegistration" }, { type: "azure-native:eventgrid/v20240601preview:PartnerRegistration" }, { type: "azure-native:eventgrid/v20241215preview:PartnerRegistration" }, { type: "azure-native:eventgrid/v20250215:PartnerRegistration" }, { type: "azure-native:eventgrid:PartnerRegistration" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(PartnerRegistration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a PartnerRegistration resource.
 */
export interface PartnerRegistrationArgs {
    /**
     * List of Azure subscription Ids that are authorized to create a partner namespace
     * associated with this partner registration. This is an optional property. Creating
     * partner namespaces is always permitted under the same Azure subscription as the one used
     * for creating the partner registration.
     */
    authorizedAzureSubscriptionIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The extension of the customer service URI of the publisher.
     */
    customerServiceUri?: pulumi.Input<string>;
    /**
     * Location of the resource.
     */
    location?: pulumi.Input<string>;
    /**
     * URI of the logo.
     */
    logoUri?: pulumi.Input<string>;
    /**
     * Long description for the custom scenarios and integration to be displayed in the portal if needed.
     * Length of this description should not exceed 2048 characters.
     */
    longDescription?: pulumi.Input<string>;
    /**
     * The extension of the customer service number of the publisher. Only digits are allowed and number of digits should not exceed 10.
     */
    partnerCustomerServiceExtension?: pulumi.Input<string>;
    /**
     * The customer service number of the publisher. The expected phone format should start with a '+' sign 
     * followed by the country code. The remaining digits are then followed. Only digits and spaces are allowed and its
     * length cannot exceed 16 digits including country code. Examples of valid phone numbers are: +1 515 123 4567 and
     * +966 7 5115 2471. Examples of invalid phone numbers are: +1 (515) 123-4567, 1 515 123 4567 and +966 121 5115 24 7 551 1234 43
     */
    partnerCustomerServiceNumber?: pulumi.Input<string>;
    /**
     * Official name of the partner name. For example: "Contoso".
     */
    partnerName?: pulumi.Input<string>;
    /**
     * The immutableId of the corresponding partner registration.
     */
    partnerRegistrationImmutableId?: pulumi.Input<string>;
    /**
     * Name of the partner registration.
     */
    partnerRegistrationName?: pulumi.Input<string>;
    /**
     * Short description of the partner resource type. The length of this description should not exceed 256 characters.
     */
    partnerResourceTypeDescription?: pulumi.Input<string>;
    /**
     * Display name of the partner resource type.
     */
    partnerResourceTypeDisplayName?: pulumi.Input<string>;
    /**
     * Name of the partner resource type.
     */
    partnerResourceTypeName?: pulumi.Input<string>;
    /**
     * The name of the resource group within the user's subscription.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * URI of the partner website that can be used by Azure customers to setup Event Grid
     * integration on an event source.
     */
    setupUri?: pulumi.Input<string>;
    /**
     * Tags of the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Visibility state of the partner registration.
     */
    visibilityState?: pulumi.Input<string | enums.eventgrid.v20211015preview.PartnerRegistrationVisibilityState>;
}
