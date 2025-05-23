// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * A provider instance associated with SAP monitor.
 *
 * Uses Azure REST API version 2024-02-01-preview. In version 2.x of the Azure Native provider, it used API version 2023-04-01.
 *
 * Other available API versions: 2023-04-01, 2023-10-01-preview, 2023-12-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native workloads [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export class ProviderInstance extends pulumi.CustomResource {
    /**
     * Get an existing ProviderInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ProviderInstance {
        return new ProviderInstance(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:workloads:ProviderInstance';

    /**
     * Returns true if the given object is an instance of ProviderInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProviderInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProviderInstance.__pulumiType;
    }

    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * Defines the provider instance errors.
     */
    public /*out*/ readonly errors!: pulumi.Output<outputs.workloads.ErrorDetailResponse>;
    /**
     * Resource health details
     */
    public /*out*/ readonly health!: pulumi.Output<outputs.workloads.HealthResponse>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Defines the provider specific properties.
     */
    public readonly providerSettings!: pulumi.Output<outputs.workloads.Db2ProviderInstancePropertiesResponse | outputs.workloads.HanaDbProviderInstancePropertiesResponse | outputs.workloads.MsSqlServerProviderInstancePropertiesResponse | outputs.workloads.OracleProviderInstancePropertiesResponse | outputs.workloads.PrometheusHaClusterProviderInstancePropertiesResponse | outputs.workloads.PrometheusOsProviderInstancePropertiesResponse | outputs.workloads.SapNetWeaverProviderInstancePropertiesResponse | undefined>;
    /**
     * State of provisioning of the provider instance
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.workloads.SystemDataResponse>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ProviderInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProviderInstanceArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.monitorName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'monitorName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            resourceInputs["monitorName"] = args ? args.monitorName : undefined;
            resourceInputs["providerInstanceName"] = args ? args.providerInstanceName : undefined;
            resourceInputs["providerSettings"] = args ? args.providerSettings : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["errors"] = undefined /*out*/;
            resourceInputs["health"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["errors"] = undefined /*out*/;
            resourceInputs["health"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["providerSettings"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:workloads/v20211201preview:ProviderInstance" }, { type: "azure-native:workloads/v20221101preview:ProviderInstance" }, { type: "azure-native:workloads/v20230401:ProviderInstance" }, { type: "azure-native:workloads/v20231001preview:ProviderInstance" }, { type: "azure-native:workloads/v20231201preview:ProviderInstance" }, { type: "azure-native:workloads/v20240201preview:ProviderInstance" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(ProviderInstance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ProviderInstance resource.
 */
export interface ProviderInstanceArgs {
    /**
     * Name of the SAP monitor resource.
     */
    monitorName: pulumi.Input<string>;
    /**
     * Name of the provider instance.
     */
    providerInstanceName?: pulumi.Input<string>;
    /**
     * Defines the provider specific properties.
     */
    providerSettings?: pulumi.Input<inputs.workloads.Db2ProviderInstancePropertiesArgs | inputs.workloads.HanaDbProviderInstancePropertiesArgs | inputs.workloads.MsSqlServerProviderInstancePropertiesArgs | inputs.workloads.OracleProviderInstancePropertiesArgs | inputs.workloads.PrometheusHaClusterProviderInstancePropertiesArgs | inputs.workloads.PrometheusOsProviderInstancePropertiesArgs | inputs.workloads.SapNetWeaverProviderInstancePropertiesArgs>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
