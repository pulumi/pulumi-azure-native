// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * A database security alert policy.
 *
 * Uses Azure REST API version 2023-08-01. In version 2.x of the Azure Native provider, it used API version 2021-11-01.
 *
 * Other available API versions: 2018-06-01-preview, 2020-02-02-preview, 2020-08-01-preview, 2020-11-01-preview, 2021-02-01-preview, 2021-05-01-preview, 2021-08-01-preview, 2021-11-01, 2021-11-01-preview, 2022-02-01-preview, 2022-05-01-preview, 2022-08-01-preview, 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview, 2023-08-01-preview, 2024-05-01-preview, 2024-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export class DatabaseSecurityAlertPolicy extends pulumi.CustomResource {
    /**
     * Get an existing DatabaseSecurityAlertPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DatabaseSecurityAlertPolicy {
        return new DatabaseSecurityAlertPolicy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:sql:DatabaseSecurityAlertPolicy';

    /**
     * Returns true if the given object is an instance of DatabaseSecurityAlertPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatabaseSecurityAlertPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatabaseSecurityAlertPolicy.__pulumiType;
    }

    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * Specifies the UTC creation time of the policy.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    /**
     * Specifies an array of alerts that are disabled. Allowed values are: Sql_Injection, Sql_Injection_Vulnerability, Access_Anomaly, Data_Exfiltration, Unsafe_Action, Brute_Force
     */
    public readonly disabledAlerts!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies that the alert is sent to the account administrators.
     */
    public readonly emailAccountAdmins!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies an array of e-mail addresses to which the alert is sent.
     */
    public readonly emailAddresses!: pulumi.Output<string[] | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Specifies the number of days to keep in the Threat Detection audit logs.
     */
    public readonly retentionDays!: pulumi.Output<number | undefined>;
    /**
     * Specifies the state of the policy, whether it is enabled or disabled or a policy has not been applied yet on the specific database.
     */
    public readonly state!: pulumi.Output<string>;
    /**
     * Specifies the identifier key of the Threat Detection audit storage account.
     */
    public readonly storageAccountAccessKey!: pulumi.Output<string | undefined>;
    /**
     * Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs.
     */
    public readonly storageEndpoint!: pulumi.Output<string | undefined>;
    /**
     * SystemData of SecurityAlertPolicyResource.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.sql.SystemDataResponse>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a DatabaseSecurityAlertPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatabaseSecurityAlertPolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.databaseName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'databaseName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.serverName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serverName'");
            }
            if ((!args || args.state === undefined) && !opts.urn) {
                throw new Error("Missing required property 'state'");
            }
            resourceInputs["databaseName"] = args ? args.databaseName : undefined;
            resourceInputs["disabledAlerts"] = args ? args.disabledAlerts : undefined;
            resourceInputs["emailAccountAdmins"] = args ? args.emailAccountAdmins : undefined;
            resourceInputs["emailAddresses"] = args ? args.emailAddresses : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["retentionDays"] = args ? args.retentionDays : undefined;
            resourceInputs["securityAlertPolicyName"] = args ? args.securityAlertPolicyName : undefined;
            resourceInputs["serverName"] = args ? args.serverName : undefined;
            resourceInputs["state"] = args ? args.state : undefined;
            resourceInputs["storageAccountAccessKey"] = args ? args.storageAccountAccessKey : undefined;
            resourceInputs["storageEndpoint"] = args ? args.storageEndpoint : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["disabledAlerts"] = undefined /*out*/;
            resourceInputs["emailAccountAdmins"] = undefined /*out*/;
            resourceInputs["emailAddresses"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["retentionDays"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["storageAccountAccessKey"] = undefined /*out*/;
            resourceInputs["storageEndpoint"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:sql/v20140401:DatabaseSecurityAlertPolicy" }, { type: "azure-native:sql/v20140401:DatabaseThreatDetectionPolicy" }, { type: "azure-native:sql/v20180601preview:DatabaseSecurityAlertPolicy" }, { type: "azure-native:sql/v20200202preview:DatabaseSecurityAlertPolicy" }, { type: "azure-native:sql/v20200801preview:DatabaseSecurityAlertPolicy" }, { type: "azure-native:sql/v20201101preview:DatabaseSecurityAlertPolicy" }, { type: "azure-native:sql/v20210201preview:DatabaseSecurityAlertPolicy" }, { type: "azure-native:sql/v20210501preview:DatabaseSecurityAlertPolicy" }, { type: "azure-native:sql/v20210801preview:DatabaseSecurityAlertPolicy" }, { type: "azure-native:sql/v20211101:DatabaseSecurityAlertPolicy" }, { type: "azure-native:sql/v20211101preview:DatabaseSecurityAlertPolicy" }, { type: "azure-native:sql/v20220201preview:DatabaseSecurityAlertPolicy" }, { type: "azure-native:sql/v20220501preview:DatabaseSecurityAlertPolicy" }, { type: "azure-native:sql/v20220801preview:DatabaseSecurityAlertPolicy" }, { type: "azure-native:sql/v20221101preview:DatabaseSecurityAlertPolicy" }, { type: "azure-native:sql/v20230201preview:DatabaseSecurityAlertPolicy" }, { type: "azure-native:sql/v20230501preview:DatabaseSecurityAlertPolicy" }, { type: "azure-native:sql/v20230801:DatabaseSecurityAlertPolicy" }, { type: "azure-native:sql/v20230801preview:DatabaseSecurityAlertPolicy" }, { type: "azure-native:sql/v20240501preview:DatabaseSecurityAlertPolicy" }, { type: "azure-native:sql/v20241101preview:DatabaseSecurityAlertPolicy" }, { type: "azure-native:sql:DatabaseThreatDetectionPolicy" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(DatabaseSecurityAlertPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a DatabaseSecurityAlertPolicy resource.
 */
export interface DatabaseSecurityAlertPolicyArgs {
    /**
     * The name of the  database for which the security alert policy is defined.
     */
    databaseName: pulumi.Input<string>;
    /**
     * Specifies an array of alerts that are disabled. Allowed values are: Sql_Injection, Sql_Injection_Vulnerability, Access_Anomaly, Data_Exfiltration, Unsafe_Action, Brute_Force
     */
    disabledAlerts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies that the alert is sent to the account administrators.
     */
    emailAccountAdmins?: pulumi.Input<boolean>;
    /**
     * Specifies an array of e-mail addresses to which the alert is sent.
     */
    emailAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Specifies the number of days to keep in the Threat Detection audit logs.
     */
    retentionDays?: pulumi.Input<number>;
    /**
     * The name of the security alert policy.
     */
    securityAlertPolicyName?: pulumi.Input<string>;
    /**
     * The name of the  server.
     */
    serverName: pulumi.Input<string>;
    /**
     * Specifies the state of the policy, whether it is enabled or disabled or a policy has not been applied yet on the specific database.
     */
    state: pulumi.Input<enums.sql.SecurityAlertsPolicyState>;
    /**
     * Specifies the identifier key of the Threat Detection audit storage account.
     */
    storageAccountAccessKey?: pulumi.Input<string>;
    /**
     * Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs.
     */
    storageEndpoint?: pulumi.Input<string>;
}
