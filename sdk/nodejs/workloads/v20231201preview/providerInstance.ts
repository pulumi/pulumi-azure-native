// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * A provider instance associated with SAP monitor.
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
    public static readonly __pulumiType = 'azure-native:workloads/v20231201preview:ProviderInstance';

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
     * Defines the provider instance errors.
     */
    public /*out*/ readonly errors!: pulumi.Output<outputs.workloads.v20231201preview.ErrorDetailResponse>;
    /**
     * Resource health details
     */
    public /*out*/ readonly health!: pulumi.Output<outputs.workloads.v20231201preview.HealthResponse>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Defines the provider specific properties.
     */
    public readonly providerSettings!: pulumi.Output<outputs.workloads.v20231201preview.Db2ProviderInstancePropertiesResponse | outputs.workloads.v20231201preview.HanaDbProviderInstancePropertiesResponse | outputs.workloads.v20231201preview.MsSqlServerProviderInstancePropertiesResponse | outputs.workloads.v20231201preview.PrometheusHaClusterProviderInstancePropertiesResponse | outputs.workloads.v20231201preview.PrometheusOsProviderInstancePropertiesResponse | outputs.workloads.v20231201preview.SapNetWeaverProviderInstancePropertiesResponse | undefined>;
    /**
     * State of provisioning of the provider instance
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.workloads.v20231201preview.SystemDataResponse>;
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
            resourceInputs["errors"] = undefined /*out*/;
            resourceInputs["health"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["errors"] = undefined /*out*/;
            resourceInputs["health"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["providerSettings"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:workloads/v20211201preview:ProviderInstance" }, { type: "azure-native:workloads/v20221101preview:ProviderInstance" }, { type: "azure-native:workloads/v20230401:ProviderInstance" }, { type: "azure-native:workloads/v20231001preview:ProviderInstance" }, { type: "azure-native:workloads/v20240201preview:ProviderInstance" }, { type: "azure-native:workloads:ProviderInstance" }] };
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
    providerSettings?: pulumi.Input<inputs.workloads.v20231201preview.Db2ProviderInstancePropertiesArgs | inputs.workloads.v20231201preview.HanaDbProviderInstancePropertiesArgs | inputs.workloads.v20231201preview.MsSqlServerProviderInstancePropertiesArgs | inputs.workloads.v20231201preview.PrometheusHaClusterProviderInstancePropertiesArgs | inputs.workloads.v20231201preview.PrometheusOsProviderInstancePropertiesArgs | inputs.workloads.v20231201preview.SapNetWeaverProviderInstancePropertiesArgs>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
}
