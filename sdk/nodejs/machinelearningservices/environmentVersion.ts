// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Azure Resource Manager resource envelope.
 *
 * Uses Azure REST API version 2024-10-01. In version 2.x of the Azure Native provider, it used API version 2023-04-01.
 *
 * Other available API versions: 2022-02-01-preview, 2022-05-01, 2022-06-01-preview, 2022-10-01, 2022-10-01-preview, 2022-12-01-preview, 2023-02-01-preview, 2023-04-01, 2023-04-01-preview, 2023-06-01-preview, 2023-08-01-preview, 2023-10-01, 2024-01-01-preview, 2024-04-01, 2024-07-01-preview, 2024-10-01-preview, 2025-01-01-preview, 2025-04-01, 2025-04-01-preview, 2025-06-01, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native machinelearningservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export class EnvironmentVersion extends pulumi.CustomResource {
    /**
     * Get an existing EnvironmentVersion resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): EnvironmentVersion {
        return new EnvironmentVersion(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:machinelearningservices:EnvironmentVersion';

    /**
     * Returns true if the given object is an instance of EnvironmentVersion.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EnvironmentVersion {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EnvironmentVersion.__pulumiType;
    }

    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * [Required] Additional attributes of the entity.
     */
    public readonly environmentVersionProperties!: pulumi.Output<outputs.machinelearningservices.EnvironmentVersionResponse>;
    /**
     * The name of the resource
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.machinelearningservices.SystemDataResponse>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a EnvironmentVersion resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EnvironmentVersionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.environmentVersionProperties === undefined) && !opts.urn) {
                throw new Error("Missing required property 'environmentVersionProperties'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.workspaceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workspaceName'");
            }
            resourceInputs["environmentVersionProperties"] = args ? (args.environmentVersionProperties ? pulumi.output(args.environmentVersionProperties).apply(inputs.machinelearningservices.environmentVersionArgsProvideDefaults) : undefined) : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["workspaceName"] = args ? args.workspaceName : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["environmentVersionProperties"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:machinelearningservices/v20210301preview:EnvironmentSpecificationVersion" }, { type: "azure-native:machinelearningservices/v20210301preview:EnvironmentVersion" }, { type: "azure-native:machinelearningservices/v20220201preview:EnvironmentVersion" }, { type: "azure-native:machinelearningservices/v20220501:EnvironmentVersion" }, { type: "azure-native:machinelearningservices/v20220601preview:EnvironmentVersion" }, { type: "azure-native:machinelearningservices/v20221001:EnvironmentVersion" }, { type: "azure-native:machinelearningservices/v20221001preview:EnvironmentVersion" }, { type: "azure-native:machinelearningservices/v20221201preview:EnvironmentVersion" }, { type: "azure-native:machinelearningservices/v20230201preview:EnvironmentVersion" }, { type: "azure-native:machinelearningservices/v20230401:EnvironmentVersion" }, { type: "azure-native:machinelearningservices/v20230401preview:EnvironmentVersion" }, { type: "azure-native:machinelearningservices/v20230601preview:EnvironmentVersion" }, { type: "azure-native:machinelearningservices/v20230801preview:EnvironmentVersion" }, { type: "azure-native:machinelearningservices/v20231001:EnvironmentVersion" }, { type: "azure-native:machinelearningservices/v20240101preview:EnvironmentVersion" }, { type: "azure-native:machinelearningservices/v20240401:EnvironmentVersion" }, { type: "azure-native:machinelearningservices/v20240401preview:EnvironmentVersion" }, { type: "azure-native:machinelearningservices/v20240701preview:EnvironmentVersion" }, { type: "azure-native:machinelearningservices/v20241001:EnvironmentVersion" }, { type: "azure-native:machinelearningservices/v20241001preview:EnvironmentVersion" }, { type: "azure-native:machinelearningservices/v20250101preview:EnvironmentVersion" }, { type: "azure-native:machinelearningservices/v20250401:EnvironmentVersion" }, { type: "azure-native:machinelearningservices/v20250401preview:EnvironmentVersion" }, { type: "azure-native:machinelearningservices/v20250601:EnvironmentVersion" }, { type: "azure-native:machinelearningservices/v20250701preview:EnvironmentVersion" }, { type: "azure-native:machinelearningservices:EnvironmentSpecificationVersion" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(EnvironmentVersion.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a EnvironmentVersion resource.
 */
export interface EnvironmentVersionArgs {
    /**
     * [Required] Additional attributes of the entity.
     */
    environmentVersionProperties: pulumi.Input<inputs.machinelearningservices.EnvironmentVersionArgs>;
    /**
     * Name of EnvironmentVersion. This is case-sensitive.
     */
    name: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Version of EnvironmentVersion.
     */
    version?: pulumi.Input<string>;
    /**
     * Name of Azure Machine Learning workspace.
     */
    workspaceName: pulumi.Input<string>;
}
