// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Scheduled action definition.
 *
 * Uses Azure REST API version 2024-08-01. In version 2.x of the Azure Native provider, it used API version 2023-03-01.
 *
 * Other available API versions: 2022-04-01-preview, 2022-06-01-preview, 2022-10-01, 2023-03-01, 2023-04-01-preview, 2023-07-01-preview, 2023-08-01, 2023-09-01, 2023-11-01, 2024-10-01-preview, 2025-03-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native costmanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export class ScheduledAction extends pulumi.CustomResource {
    /**
     * Get an existing ScheduledAction resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ScheduledAction {
        return new ScheduledAction(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:costmanagement:ScheduledAction';

    /**
     * Returns true if the given object is an instance of ScheduledAction.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ScheduledAction {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ScheduledAction.__pulumiType;
    }

    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * Scheduled action name.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Resource Etag. For update calls, eTag is optional and can be specified to achieve optimistic concurrency. Fetch the resource's eTag by doing a 'GET' call first and then including the latest eTag as part of the request body or 'If-Match' header while performing the update. For create calls, eTag is not required.
     */
    public /*out*/ readonly eTag!: pulumi.Output<string>;
    /**
     * Destination format of the view data. This is optional.
     */
    public readonly fileDestination!: pulumi.Output<outputs.costmanagement.FileDestinationResponse | undefined>;
    /**
     * Kind of the scheduled action.
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Notification properties based on scheduled action kind.
     */
    public readonly notification!: pulumi.Output<outputs.costmanagement.NotificationPropertiesResponse>;
    /**
     * Email address of the point of contact that should get the unsubscribe requests and notification emails.
     */
    public readonly notificationEmail!: pulumi.Output<string | undefined>;
    /**
     * Schedule of the scheduled action.
     */
    public readonly schedule!: pulumi.Output<outputs.costmanagement.SchedulePropertiesResponse>;
    /**
     * For private scheduled action(Create or Update), scope will be empty.<br /> For shared scheduled action(Create or Update By Scope), Cost Management scope can be 'subscriptions/{subscriptionId}' for subscription scope, 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for InvoiceSection scope, '/providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountName}' for ExternalBillingAccount scope, and '/providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}' for ExternalSubscription scope.
     */
    public readonly scope!: pulumi.Output<string | undefined>;
    /**
     * Status of the scheduled action.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Kind of the scheduled action.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.costmanagement.SystemDataResponse>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Cost analysis viewId used for scheduled action. For example, '/providers/Microsoft.CostManagement/views/swaggerExample'
     */
    public readonly viewId!: pulumi.Output<string>;

    /**
     * Create a ScheduledAction resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ScheduledActionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.notification === undefined) && !opts.urn) {
                throw new Error("Missing required property 'notification'");
            }
            if ((!args || args.schedule === undefined) && !opts.urn) {
                throw new Error("Missing required property 'schedule'");
            }
            if ((!args || args.status === undefined) && !opts.urn) {
                throw new Error("Missing required property 'status'");
            }
            if ((!args || args.viewId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'viewId'");
            }
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["fileDestination"] = args ? args.fileDestination : undefined;
            resourceInputs["kind"] = args ? args.kind : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["notification"] = args ? args.notification : undefined;
            resourceInputs["notificationEmail"] = args ? args.notificationEmail : undefined;
            resourceInputs["schedule"] = args ? args.schedule : undefined;
            resourceInputs["scope"] = args ? args.scope : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["viewId"] = args ? args.viewId : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["eTag"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["eTag"] = undefined /*out*/;
            resourceInputs["fileDestination"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["notification"] = undefined /*out*/;
            resourceInputs["notificationEmail"] = undefined /*out*/;
            resourceInputs["schedule"] = undefined /*out*/;
            resourceInputs["scope"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["viewId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:costmanagement/v20220401preview:ScheduledAction" }, { type: "azure-native:costmanagement/v20220601preview:ScheduledAction" }, { type: "azure-native:costmanagement/v20221001:ScheduledAction" }, { type: "azure-native:costmanagement/v20230301:ScheduledAction" }, { type: "azure-native:costmanagement/v20230401preview:ScheduledAction" }, { type: "azure-native:costmanagement/v20230701preview:ScheduledAction" }, { type: "azure-native:costmanagement/v20230801:ScheduledAction" }, { type: "azure-native:costmanagement/v20230901:ScheduledAction" }, { type: "azure-native:costmanagement/v20231101:ScheduledAction" }, { type: "azure-native:costmanagement/v20240801:ScheduledAction" }, { type: "azure-native:costmanagement/v20241001preview:ScheduledAction" }, { type: "azure-native:costmanagement/v20250301:ScheduledAction" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(ScheduledAction.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ScheduledAction resource.
 */
export interface ScheduledActionArgs {
    /**
     * Scheduled action name.
     */
    displayName: pulumi.Input<string>;
    /**
     * Destination format of the view data. This is optional.
     */
    fileDestination?: pulumi.Input<inputs.costmanagement.FileDestinationArgs>;
    /**
     * Kind of the scheduled action.
     */
    kind?: pulumi.Input<string | enums.costmanagement.ScheduledActionKind>;
    /**
     * Scheduled action name.
     */
    name?: pulumi.Input<string>;
    /**
     * Notification properties based on scheduled action kind.
     */
    notification: pulumi.Input<inputs.costmanagement.NotificationPropertiesArgs>;
    /**
     * Email address of the point of contact that should get the unsubscribe requests and notification emails.
     */
    notificationEmail?: pulumi.Input<string>;
    /**
     * Schedule of the scheduled action.
     */
    schedule: pulumi.Input<inputs.costmanagement.SchedulePropertiesArgs>;
    /**
     * For private scheduled action(Create or Update), scope will be empty.<br /> For shared scheduled action(Create or Update By Scope), Cost Management scope can be 'subscriptions/{subscriptionId}' for subscription scope, 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for InvoiceSection scope, '/providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountName}' for ExternalBillingAccount scope, and '/providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}' for ExternalSubscription scope.
     */
    scope?: pulumi.Input<string>;
    /**
     * Status of the scheduled action.
     */
    status: pulumi.Input<string | enums.costmanagement.ScheduledActionStatus>;
    /**
     * Cost analysis viewId used for scheduled action. For example, '/providers/Microsoft.CostManagement/views/swaggerExample'
     */
    viewId: pulumi.Input<string>;
}
