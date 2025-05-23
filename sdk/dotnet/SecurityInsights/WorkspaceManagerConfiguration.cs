// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityInsights
{
    /// <summary>
    /// The workspace manager configuration
    /// 
    /// Uses Azure REST API version 2025-01-01-preview. In version 2.x of the Azure Native provider, it used API version 2023-06-01-preview.
    /// 
    /// Other available API versions: 2023-04-01-preview, 2023-05-01-preview, 2023-06-01-preview, 2023-07-01-preview, 2023-08-01-preview, 2023-09-01-preview, 2023-10-01-preview, 2023-12-01-preview, 2024-01-01-preview, 2024-04-01-preview, 2024-10-01-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native securityinsights [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:securityinsights:WorkspaceManagerConfiguration")]
    public partial class WorkspaceManagerConfiguration : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Resource Etag.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The current mode of the workspace manager configuration
        /// </summary>
        [Output("mode")]
        public Output<string> Mode { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a WorkspaceManagerConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WorkspaceManagerConfiguration(string name, WorkspaceManagerConfigurationArgs args, CustomResourceOptions? options = null)
            : base("azure-native:securityinsights:WorkspaceManagerConfiguration", name, args ?? new WorkspaceManagerConfigurationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WorkspaceManagerConfiguration(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:securityinsights:WorkspaceManagerConfiguration", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230401preview:WorkspaceManagerConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230501preview:WorkspaceManagerConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230601preview:WorkspaceManagerConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230701preview:WorkspaceManagerConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230801preview:WorkspaceManagerConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230901preview:WorkspaceManagerConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20231001preview:WorkspaceManagerConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20231201preview:WorkspaceManagerConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20240101preview:WorkspaceManagerConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20240401preview:WorkspaceManagerConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20241001preview:WorkspaceManagerConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20250101preview:WorkspaceManagerConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20250401preview:WorkspaceManagerConfiguration" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing WorkspaceManagerConfiguration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WorkspaceManagerConfiguration Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new WorkspaceManagerConfiguration(name, id, options);
        }
    }

    public sealed class WorkspaceManagerConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The current mode of the workspace manager configuration
        /// </summary>
        [Input("mode", required: true)]
        public InputUnion<string, Pulumi.AzureNative.SecurityInsights.Mode> Mode { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the workspace manager configuration
        /// </summary>
        [Input("workspaceManagerConfigurationName")]
        public Input<string>? WorkspaceManagerConfigurationName { get; set; }

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public WorkspaceManagerConfigurationArgs()
        {
        }
        public static new WorkspaceManagerConfigurationArgs Empty => new WorkspaceManagerConfigurationArgs();
    }
}
