// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * A export resource.
 */
export class Export extends pulumi.CustomResource {
    /**
     * Get an existing Export resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Export {
        return new Export(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:costmanagement/v20191001:Export';

    /**
     * Returns true if the given object is an instance of Export.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Export {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Export.__pulumiType;
    }

    /**
     * Has definition for the export.
     */
    public readonly definition!: pulumi.Output<outputs.costmanagement.v20191001.QueryDefinitionResponse>;
    /**
     * Has delivery information for the export.
     */
    public readonly deliveryInfo!: pulumi.Output<outputs.costmanagement.v20191001.ExportDeliveryInfoResponse>;
    /**
     * The format of the export being delivered.
     */
    public readonly format!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Has schedule information for the export.
     */
    public readonly schedule!: pulumi.Output<outputs.costmanagement.v20191001.ExportScheduleResponse | undefined>;
    /**
     * Resource tags.
     */
    public /*out*/ readonly tags!: pulumi.Output<{[key: string]: string}>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Export resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ExportArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.definition === undefined) && !opts.urn) {
                throw new Error("Missing required property 'definition'");
            }
            if ((!args || args.deliveryInfo === undefined) && !opts.urn) {
                throw new Error("Missing required property 'deliveryInfo'");
            }
            if ((!args || args.scope === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scope'");
            }
            resourceInputs["definition"] = args ? args.definition : undefined;
            resourceInputs["deliveryInfo"] = args ? args.deliveryInfo : undefined;
            resourceInputs["exportName"] = args ? args.exportName : undefined;
            resourceInputs["format"] = args ? args.format : undefined;
            resourceInputs["schedule"] = args ? args.schedule : undefined;
            resourceInputs["scope"] = args ? args.scope : undefined;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["definition"] = undefined /*out*/;
            resourceInputs["deliveryInfo"] = undefined /*out*/;
            resourceInputs["format"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["schedule"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:costmanagement/v20190101:Export" }, { type: "azure-native:costmanagement/v20190901:Export" }, { type: "azure-native:costmanagement/v20191101:Export" }, { type: "azure-native:costmanagement/v20200601:Export" }, { type: "azure-native:costmanagement/v20201201preview:Export" }, { type: "azure-native:costmanagement/v20210101:Export" }, { type: "azure-native:costmanagement/v20211001:Export" }, { type: "azure-native:costmanagement/v20221001:Export" }, { type: "azure-native:costmanagement/v20230301:Export" }, { type: "azure-native:costmanagement/v20230401preview:Export" }, { type: "azure-native:costmanagement/v20230701preview:Export" }, { type: "azure-native:costmanagement/v20230801:Export" }, { type: "azure-native:costmanagement/v20230901:Export" }, { type: "azure-native:costmanagement/v20231101:Export" }, { type: "azure-native:costmanagement/v20240801:Export" }, { type: "azure-native:costmanagement/v20241001preview:Export" }, { type: "azure-native:costmanagement:Export" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(Export.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Export resource.
 */
export interface ExportArgs {
    /**
     * Has definition for the export.
     */
    definition: pulumi.Input<inputs.costmanagement.v20191001.QueryDefinitionArgs>;
    /**
     * Has delivery information for the export.
     */
    deliveryInfo: pulumi.Input<inputs.costmanagement.v20191001.ExportDeliveryInfoArgs>;
    /**
     * Export Name.
     */
    exportName?: pulumi.Input<string>;
    /**
     * The format of the export being delivered.
     */
    format?: pulumi.Input<string | enums.costmanagement.v20191001.FormatType>;
    /**
     * Has schedule information for the export.
     */
    schedule?: pulumi.Input<inputs.costmanagement.v20191001.ExportScheduleArgs>;
    /**
     * The scope associated with query and export operations. This includes '/subscriptions/{subscriptionId}/' for subscription scope, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope and '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, '/providers/Microsoft.Management/managementGroups/{managementGroupId} for Management Group scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for billingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/invoiceSections/{invoiceSectionId}' for invoiceSection scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' specific for partners, 'providers/Microsoft.CostManagement/ExternalSubscriptions/{externalSubscriptionId}' for linked account and 'providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountId}' for consolidated account
     */
    scope: pulumi.Input<string>;
}
