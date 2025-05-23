// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The metric setting details for the role
 *
 * Uses Azure REST API version 2023-07-01. In version 2.x of the Azure Native provider, it used API version 2022-03-01.
 *
 * Other available API versions: 2022-03-01, 2022-04-01-preview, 2022-12-01-preview, 2023-01-01-preview, 2023-12-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native databoxedge [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export class MonitoringConfig extends pulumi.CustomResource {
    /**
     * Get an existing MonitoringConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): MonitoringConfig {
        return new MonitoringConfig(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:databoxedge:MonitoringConfig';

    /**
     * Returns true if the given object is an instance of MonitoringConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MonitoringConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MonitoringConfig.__pulumiType;
    }

    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * The metrics configuration details
     */
    public readonly metricConfigurations!: pulumi.Output<outputs.databoxedge.MetricConfigurationResponse[]>;
    /**
     * The object name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Metadata pertaining to creation and last modification of MonitoringConfiguration
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.databoxedge.SystemDataResponse>;
    /**
     * The hierarchical type of the object.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a MonitoringConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MonitoringConfigArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.deviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'deviceName'");
            }
            if ((!args || args.metricConfigurations === undefined) && !opts.urn) {
                throw new Error("Missing required property 'metricConfigurations'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.roleName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleName'");
            }
            resourceInputs["deviceName"] = args ? args.deviceName : undefined;
            resourceInputs["metricConfigurations"] = args ? args.metricConfigurations : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["roleName"] = args ? args.roleName : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["metricConfigurations"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:databoxedge/v20200901:MonitoringConfig" }, { type: "azure-native:databoxedge/v20200901preview:MonitoringConfig" }, { type: "azure-native:databoxedge/v20201201:MonitoringConfig" }, { type: "azure-native:databoxedge/v20210201:MonitoringConfig" }, { type: "azure-native:databoxedge/v20210201preview:MonitoringConfig" }, { type: "azure-native:databoxedge/v20210601:MonitoringConfig" }, { type: "azure-native:databoxedge/v20210601preview:MonitoringConfig" }, { type: "azure-native:databoxedge/v20220301:MonitoringConfig" }, { type: "azure-native:databoxedge/v20220401preview:MonitoringConfig" }, { type: "azure-native:databoxedge/v20221201preview:MonitoringConfig" }, { type: "azure-native:databoxedge/v20230101preview:MonitoringConfig" }, { type: "azure-native:databoxedge/v20230701:MonitoringConfig" }, { type: "azure-native:databoxedge/v20231201:MonitoringConfig" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(MonitoringConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a MonitoringConfig resource.
 */
export interface MonitoringConfigArgs {
    /**
     * The device name.
     */
    deviceName: pulumi.Input<string>;
    /**
     * The metrics configuration details
     */
    metricConfigurations: pulumi.Input<pulumi.Input<inputs.databoxedge.MetricConfigurationArgs>[]>;
    /**
     * The resource group name.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The role name.
     */
    roleName: pulumi.Input<string>;
}
