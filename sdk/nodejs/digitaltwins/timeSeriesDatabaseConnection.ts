// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Describes a time series database connection resource.
 *
 * Uses Azure REST API version 2023-01-31. In version 2.x of the Azure Native provider, it used API version 2023-01-31.
 */
export class TimeSeriesDatabaseConnection extends pulumi.CustomResource {
    /**
     * Get an existing TimeSeriesDatabaseConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): TimeSeriesDatabaseConnection {
        return new TimeSeriesDatabaseConnection(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:digitaltwins:TimeSeriesDatabaseConnection';

    /**
     * Returns true if the given object is an instance of TimeSeriesDatabaseConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TimeSeriesDatabaseConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TimeSeriesDatabaseConnection.__pulumiType;
    }

    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * Extension resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Properties of a specific time series database connection.
     */
    public readonly properties!: pulumi.Output<outputs.digitaltwins.AzureDataExplorerConnectionPropertiesResponse>;
    /**
     * Metadata pertaining to creation and last modification of the resource.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.digitaltwins.SystemDataResponse>;
    /**
     * The resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a TimeSeriesDatabaseConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TimeSeriesDatabaseConnectionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.resourceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceName'");
            }
            resourceInputs["properties"] = args ? (args.properties ? pulumi.output(args.properties).apply(inputs.digitaltwins.azureDataExplorerConnectionPropertiesArgsProvideDefaults) : undefined) : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["resourceName"] = args ? args.resourceName : undefined;
            resourceInputs["timeSeriesDatabaseConnectionName"] = args ? args.timeSeriesDatabaseConnectionName : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["properties"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:digitaltwins/v20210630preview:TimeSeriesDatabaseConnection" }, { type: "azure-native:digitaltwins/v20220531:TimeSeriesDatabaseConnection" }, { type: "azure-native:digitaltwins/v20221031:TimeSeriesDatabaseConnection" }, { type: "azure-native:digitaltwins/v20230131:TimeSeriesDatabaseConnection" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(TimeSeriesDatabaseConnection.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a TimeSeriesDatabaseConnection resource.
 */
export interface TimeSeriesDatabaseConnectionArgs {
    /**
     * Properties of a specific time series database connection.
     */
    properties?: pulumi.Input<inputs.digitaltwins.AzureDataExplorerConnectionPropertiesArgs>;
    /**
     * The name of the resource group that contains the DigitalTwinsInstance.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the DigitalTwinsInstance.
     */
    resourceName: pulumi.Input<string>;
    /**
     * Name of time series database connection.
     */
    timeSeriesDatabaseConnectionName?: pulumi.Input<string>;
}
