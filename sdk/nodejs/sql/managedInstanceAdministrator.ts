// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * An Azure SQL managed instance administrator.
 *
 * Uses Azure REST API version 2023-08-01. In version 2.x of the Azure Native provider, it used API version 2021-11-01.
 *
 * Other available API versions: 2017-03-01-preview, 2020-02-02-preview, 2020-08-01-preview, 2020-11-01-preview, 2021-02-01-preview, 2021-05-01-preview, 2021-08-01-preview, 2021-11-01, 2021-11-01-preview, 2022-02-01-preview, 2022-05-01-preview, 2022-08-01-preview, 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview, 2023-08-01-preview, 2024-05-01-preview, 2024-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export class ManagedInstanceAdministrator extends pulumi.CustomResource {
    /**
     * Get an existing ManagedInstanceAdministrator resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ManagedInstanceAdministrator {
        return new ManagedInstanceAdministrator(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:sql:ManagedInstanceAdministrator';

    /**
     * Returns true if the given object is an instance of ManagedInstanceAdministrator.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ManagedInstanceAdministrator {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ManagedInstanceAdministrator.__pulumiType;
    }

    /**
     * Type of the managed instance administrator.
     */
    public readonly administratorType!: pulumi.Output<string>;
    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * Login name of the managed instance administrator.
     */
    public readonly login!: pulumi.Output<string>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * SID (object ID) of the managed instance administrator.
     */
    public readonly sid!: pulumi.Output<string>;
    /**
     * Tenant ID of the managed instance administrator.
     */
    public readonly tenantId!: pulumi.Output<string | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ManagedInstanceAdministrator resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ManagedInstanceAdministratorArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.administratorType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'administratorType'");
            }
            if ((!args || args.login === undefined) && !opts.urn) {
                throw new Error("Missing required property 'login'");
            }
            if ((!args || args.managedInstanceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'managedInstanceName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.sid === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sid'");
            }
            resourceInputs["administratorName"] = args ? args.administratorName : undefined;
            resourceInputs["administratorType"] = args ? args.administratorType : undefined;
            resourceInputs["login"] = args ? args.login : undefined;
            resourceInputs["managedInstanceName"] = args ? args.managedInstanceName : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["sid"] = args ? args.sid : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["administratorType"] = undefined /*out*/;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["login"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["sid"] = undefined /*out*/;
            resourceInputs["tenantId"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:sql/v20170301preview:ManagedInstanceAdministrator" }, { type: "azure-native:sql/v20200202preview:ManagedInstanceAdministrator" }, { type: "azure-native:sql/v20200801preview:ManagedInstanceAdministrator" }, { type: "azure-native:sql/v20201101preview:ManagedInstanceAdministrator" }, { type: "azure-native:sql/v20210201preview:ManagedInstanceAdministrator" }, { type: "azure-native:sql/v20210501preview:ManagedInstanceAdministrator" }, { type: "azure-native:sql/v20210801preview:ManagedInstanceAdministrator" }, { type: "azure-native:sql/v20211101:ManagedInstanceAdministrator" }, { type: "azure-native:sql/v20211101preview:ManagedInstanceAdministrator" }, { type: "azure-native:sql/v20220201preview:ManagedInstanceAdministrator" }, { type: "azure-native:sql/v20220501preview:ManagedInstanceAdministrator" }, { type: "azure-native:sql/v20220801preview:ManagedInstanceAdministrator" }, { type: "azure-native:sql/v20221101preview:ManagedInstanceAdministrator" }, { type: "azure-native:sql/v20230201preview:ManagedInstanceAdministrator" }, { type: "azure-native:sql/v20230501preview:ManagedInstanceAdministrator" }, { type: "azure-native:sql/v20230801:ManagedInstanceAdministrator" }, { type: "azure-native:sql/v20230801preview:ManagedInstanceAdministrator" }, { type: "azure-native:sql/v20240501preview:ManagedInstanceAdministrator" }, { type: "azure-native:sql/v20241101preview:ManagedInstanceAdministrator" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(ManagedInstanceAdministrator.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ManagedInstanceAdministrator resource.
 */
export interface ManagedInstanceAdministratorArgs {
    administratorName?: pulumi.Input<string>;
    /**
     * Type of the managed instance administrator.
     */
    administratorType: pulumi.Input<string | enums.sql.ManagedInstanceAdministratorType>;
    /**
     * Login name of the managed instance administrator.
     */
    login: pulumi.Input<string>;
    /**
     * The name of the managed instance.
     */
    managedInstanceName: pulumi.Input<string>;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * SID (object ID) of the managed instance administrator.
     */
    sid: pulumi.Input<string>;
    /**
     * Tenant ID of the managed instance administrator.
     */
    tenantId?: pulumi.Input<string>;
}
