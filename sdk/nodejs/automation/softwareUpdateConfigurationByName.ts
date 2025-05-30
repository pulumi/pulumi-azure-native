// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Software update configuration properties.
 *
 * Uses Azure REST API version 2023-05-15-preview. In version 2.x of the Azure Native provider, it used API version 2019-06-01.
 *
 * Other available API versions: 2017-05-15-preview, 2019-06-01, 2024-10-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native automation [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export class SoftwareUpdateConfigurationByName extends pulumi.CustomResource {
    /**
     * Get an existing SoftwareUpdateConfigurationByName resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SoftwareUpdateConfigurationByName {
        return new SoftwareUpdateConfigurationByName(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:automation:SoftwareUpdateConfigurationByName';

    /**
     * Returns true if the given object is an instance of SoftwareUpdateConfigurationByName.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SoftwareUpdateConfigurationByName {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SoftwareUpdateConfigurationByName.__pulumiType;
    }

    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * CreatedBy property, which only appears in the response.
     */
    public /*out*/ readonly createdBy!: pulumi.Output<string>;
    /**
     * Creation time of the resource, which only appears in the response.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    /**
     * Details of provisioning error
     */
    public readonly error!: pulumi.Output<outputs.automation.ErrorResponseResponse | undefined>;
    /**
     * LastModifiedBy property, which only appears in the response.
     */
    public /*out*/ readonly lastModifiedBy!: pulumi.Output<string>;
    /**
     * Last time resource was modified, which only appears in the response.
     */
    public /*out*/ readonly lastModifiedTime!: pulumi.Output<string>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Provisioning state for the software update configuration, which only appears in the response.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Schedule information for the Software update configuration
     */
    public readonly scheduleInfo!: pulumi.Output<outputs.automation.SUCSchedulePropertiesResponse>;
    /**
     * Tasks information for the Software update configuration.
     */
    public readonly tasks!: pulumi.Output<outputs.automation.SoftwareUpdateConfigurationTasksResponse | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * update specific properties for the Software update configuration
     */
    public readonly updateConfiguration!: pulumi.Output<outputs.automation.UpdateConfigurationResponse>;

    /**
     * Create a SoftwareUpdateConfigurationByName resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SoftwareUpdateConfigurationByNameArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.automationAccountName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'automationAccountName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.scheduleInfo === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scheduleInfo'");
            }
            if ((!args || args.updateConfiguration === undefined) && !opts.urn) {
                throw new Error("Missing required property 'updateConfiguration'");
            }
            resourceInputs["automationAccountName"] = args ? args.automationAccountName : undefined;
            resourceInputs["error"] = args ? args.error : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["scheduleInfo"] = args ? (args.scheduleInfo ? pulumi.output(args.scheduleInfo).apply(inputs.automation.sucschedulePropertiesArgsProvideDefaults) : undefined) : undefined;
            resourceInputs["softwareUpdateConfigurationName"] = args ? args.softwareUpdateConfigurationName : undefined;
            resourceInputs["tasks"] = args ? args.tasks : undefined;
            resourceInputs["updateConfiguration"] = args ? args.updateConfiguration : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["createdBy"] = undefined /*out*/;
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["lastModifiedBy"] = undefined /*out*/;
            resourceInputs["lastModifiedTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["createdBy"] = undefined /*out*/;
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["error"] = undefined /*out*/;
            resourceInputs["lastModifiedBy"] = undefined /*out*/;
            resourceInputs["lastModifiedTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["scheduleInfo"] = undefined /*out*/;
            resourceInputs["tasks"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["updateConfiguration"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:automation/v20170515preview:SoftwareUpdateConfigurationByName" }, { type: "azure-native:automation/v20190601:SoftwareUpdateConfigurationByName" }, { type: "azure-native:automation/v20230515preview:SoftwareUpdateConfigurationByName" }, { type: "azure-native:automation/v20241023:SoftwareUpdateConfigurationByName" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(SoftwareUpdateConfigurationByName.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a SoftwareUpdateConfigurationByName resource.
 */
export interface SoftwareUpdateConfigurationByNameArgs {
    /**
     * The name of the automation account.
     */
    automationAccountName: pulumi.Input<string>;
    /**
     * Details of provisioning error
     */
    error?: pulumi.Input<inputs.automation.ErrorResponseArgs>;
    /**
     * Name of an Azure Resource group.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Schedule information for the Software update configuration
     */
    scheduleInfo: pulumi.Input<inputs.automation.SUCSchedulePropertiesArgs>;
    /**
     * The name of the software update configuration to be created.
     */
    softwareUpdateConfigurationName?: pulumi.Input<string>;
    /**
     * Tasks information for the Software update configuration.
     */
    tasks?: pulumi.Input<inputs.automation.SoftwareUpdateConfigurationTasksArgs>;
    /**
     * update specific properties for the Software update configuration
     */
    updateConfiguration: pulumi.Input<inputs.automation.UpdateConfigurationArgs>;
}
