// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Uses Azure REST API version 2024-10-01. In version 2.x of the Azure Native provider, it used API version 2023-04-01.
 *
 * Other available API versions: 2020-06-01, 2020-08-01, 2020-09-01-preview, 2021-01-01, 2021-03-01-preview, 2021-04-01, 2021-07-01, 2022-01-01-preview, 2022-02-01-preview, 2022-05-01, 2022-06-01-preview, 2022-10-01, 2022-10-01-preview, 2022-12-01-preview, 2023-02-01-preview, 2023-04-01, 2023-04-01-preview, 2023-06-01-preview, 2023-08-01-preview, 2023-10-01, 2024-01-01-preview, 2024-04-01, 2024-07-01-preview, 2024-10-01-preview, 2025-01-01-preview, 2025-04-01, 2025-04-01-preview, 2025-06-01, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native machinelearningservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export class WorkspaceConnection extends pulumi.CustomResource {
    /**
     * Get an existing WorkspaceConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): WorkspaceConnection {
        return new WorkspaceConnection(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:machinelearningservices:WorkspaceConnection';

    /**
     * Returns true if the given object is an instance of WorkspaceConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WorkspaceConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WorkspaceConnection.__pulumiType;
    }

    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    public readonly properties!: pulumi.Output<outputs.machinelearningservices.AADAuthTypeWorkspaceConnectionPropertiesResponse | outputs.machinelearningservices.AccessKeyAuthTypeWorkspaceConnectionPropertiesResponse | outputs.machinelearningservices.AccountKeyAuthTypeWorkspaceConnectionPropertiesResponse | outputs.machinelearningservices.ApiKeyAuthWorkspaceConnectionPropertiesResponse | outputs.machinelearningservices.CustomKeysWorkspaceConnectionPropertiesResponse | outputs.machinelearningservices.ManagedIdentityAuthTypeWorkspaceConnectionPropertiesResponse | outputs.machinelearningservices.NoneAuthTypeWorkspaceConnectionPropertiesResponse | outputs.machinelearningservices.OAuth2AuthTypeWorkspaceConnectionPropertiesResponse | outputs.machinelearningservices.PATAuthTypeWorkspaceConnectionPropertiesResponse | outputs.machinelearningservices.SASAuthTypeWorkspaceConnectionPropertiesResponse | outputs.machinelearningservices.ServicePrincipalAuthTypeWorkspaceConnectionPropertiesResponse | outputs.machinelearningservices.UsernamePasswordAuthTypeWorkspaceConnectionPropertiesResponse>;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.machinelearningservices.SystemDataResponse>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a WorkspaceConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WorkspaceConnectionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.properties === undefined) && !opts.urn) {
                throw new Error("Missing required property 'properties'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.workspaceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workspaceName'");
            }
            resourceInputs["connectionName"] = args ? args.connectionName : undefined;
            resourceInputs["properties"] = args ? args.properties : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["workspaceName"] = args ? args.workspaceName : undefined;
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
        const aliasOpts = { aliases: [{ type: "azure-native:machinelearningservices/v20200601:WorkspaceConnection" }, { type: "azure-native:machinelearningservices/v20200801:WorkspaceConnection" }, { type: "azure-native:machinelearningservices/v20200901preview:WorkspaceConnection" }, { type: "azure-native:machinelearningservices/v20210101:WorkspaceConnection" }, { type: "azure-native:machinelearningservices/v20210301preview:WorkspaceConnection" }, { type: "azure-native:machinelearningservices/v20210401:WorkspaceConnection" }, { type: "azure-native:machinelearningservices/v20210701:WorkspaceConnection" }, { type: "azure-native:machinelearningservices/v20220101preview:WorkspaceConnection" }, { type: "azure-native:machinelearningservices/v20220201preview:WorkspaceConnection" }, { type: "azure-native:machinelearningservices/v20220501:WorkspaceConnection" }, { type: "azure-native:machinelearningservices/v20220601preview:WorkspaceConnection" }, { type: "azure-native:machinelearningservices/v20221001:WorkspaceConnection" }, { type: "azure-native:machinelearningservices/v20221001preview:WorkspaceConnection" }, { type: "azure-native:machinelearningservices/v20221201preview:WorkspaceConnection" }, { type: "azure-native:machinelearningservices/v20230201preview:WorkspaceConnection" }, { type: "azure-native:machinelearningservices/v20230401:WorkspaceConnection" }, { type: "azure-native:machinelearningservices/v20230401preview:WorkspaceConnection" }, { type: "azure-native:machinelearningservices/v20230601preview:WorkspaceConnection" }, { type: "azure-native:machinelearningservices/v20230801preview:WorkspaceConnection" }, { type: "azure-native:machinelearningservices/v20231001:WorkspaceConnection" }, { type: "azure-native:machinelearningservices/v20240101preview:WorkspaceConnection" }, { type: "azure-native:machinelearningservices/v20240401:WorkspaceConnection" }, { type: "azure-native:machinelearningservices/v20240401preview:WorkspaceConnection" }, { type: "azure-native:machinelearningservices/v20240701preview:WorkspaceConnection" }, { type: "azure-native:machinelearningservices/v20241001:WorkspaceConnection" }, { type: "azure-native:machinelearningservices/v20241001preview:WorkspaceConnection" }, { type: "azure-native:machinelearningservices/v20250101preview:WorkspaceConnection" }, { type: "azure-native:machinelearningservices/v20250401:WorkspaceConnection" }, { type: "azure-native:machinelearningservices/v20250401preview:WorkspaceConnection" }, { type: "azure-native:machinelearningservices/v20250601:WorkspaceConnection" }, { type: "azure-native:machinelearningservices/v20250701preview:WorkspaceConnection" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(WorkspaceConnection.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a WorkspaceConnection resource.
 */
export interface WorkspaceConnectionArgs {
    /**
     * Friendly name of the workspace connection
     */
    connectionName?: pulumi.Input<string>;
    properties: pulumi.Input<inputs.machinelearningservices.AADAuthTypeWorkspaceConnectionPropertiesArgs | inputs.machinelearningservices.AccessKeyAuthTypeWorkspaceConnectionPropertiesArgs | inputs.machinelearningservices.AccountKeyAuthTypeWorkspaceConnectionPropertiesArgs | inputs.machinelearningservices.ApiKeyAuthWorkspaceConnectionPropertiesArgs | inputs.machinelearningservices.CustomKeysWorkspaceConnectionPropertiesArgs | inputs.machinelearningservices.ManagedIdentityAuthTypeWorkspaceConnectionPropertiesArgs | inputs.machinelearningservices.NoneAuthTypeWorkspaceConnectionPropertiesArgs | inputs.machinelearningservices.OAuth2AuthTypeWorkspaceConnectionPropertiesArgs | inputs.machinelearningservices.PATAuthTypeWorkspaceConnectionPropertiesArgs | inputs.machinelearningservices.SASAuthTypeWorkspaceConnectionPropertiesArgs | inputs.machinelearningservices.ServicePrincipalAuthTypeWorkspaceConnectionPropertiesArgs | inputs.machinelearningservices.UsernamePasswordAuthTypeWorkspaceConnectionPropertiesArgs>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Name of Azure Machine Learning workspace.
     */
    workspaceName: pulumi.Input<string>;
}
