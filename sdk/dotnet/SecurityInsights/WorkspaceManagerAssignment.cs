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
    /// The workspace manager assignment
    /// 
    /// Uses Azure REST API version 2025-01-01-preview. In version 2.x of the Azure Native provider, it used API version 2023-06-01-preview.
    /// 
    /// Other available API versions: 2023-04-01-preview, 2023-05-01-preview, 2023-06-01-preview, 2023-07-01-preview, 2023-08-01-preview, 2023-09-01-preview, 2023-10-01-preview, 2023-12-01-preview, 2024-01-01-preview, 2024-04-01-preview, 2024-10-01-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native securityinsights [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:securityinsights:WorkspaceManagerAssignment")]
    public partial class WorkspaceManagerAssignment : global::Pulumi.CustomResource
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
        /// List of resources included in this workspace manager assignment
        /// </summary>
        [Output("items")]
        public Output<ImmutableArray<Outputs.AssignmentItemResponse>> Items { get; private set; } = null!;

        /// <summary>
        /// The time the last job associated to this assignment ended at
        /// </summary>
        [Output("lastJobEndTime")]
        public Output<string> LastJobEndTime { get; private set; } = null!;

        /// <summary>
        /// State of the last job associated to this assignment
        /// </summary>
        [Output("lastJobProvisioningState")]
        public Output<string> LastJobProvisioningState { get; private set; } = null!;

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
        /// The resource name of the workspace manager group targeted by the workspace manager assignment
        /// </summary>
        [Output("targetResourceName")]
        public Output<string> TargetResourceName { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a WorkspaceManagerAssignment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WorkspaceManagerAssignment(string name, WorkspaceManagerAssignmentArgs args, CustomResourceOptions? options = null)
            : base("azure-native:securityinsights:WorkspaceManagerAssignment", name, args ?? new WorkspaceManagerAssignmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WorkspaceManagerAssignment(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:securityinsights:WorkspaceManagerAssignment", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230401preview:WorkspaceManagerAssignment" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230501preview:WorkspaceManagerAssignment" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230601preview:WorkspaceManagerAssignment" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230701preview:WorkspaceManagerAssignment" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230801preview:WorkspaceManagerAssignment" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20230901preview:WorkspaceManagerAssignment" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20231001preview:WorkspaceManagerAssignment" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20231201preview:WorkspaceManagerAssignment" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20240101preview:WorkspaceManagerAssignment" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20240401preview:WorkspaceManagerAssignment" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20241001preview:WorkspaceManagerAssignment" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20250101preview:WorkspaceManagerAssignment" },
                    new global::Pulumi.Alias { Type = "azure-native:securityinsights/v20250401preview:WorkspaceManagerAssignment" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing WorkspaceManagerAssignment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WorkspaceManagerAssignment Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new WorkspaceManagerAssignment(name, id, options);
        }
    }

    public sealed class WorkspaceManagerAssignmentArgs : global::Pulumi.ResourceArgs
    {
        [Input("items", required: true)]
        private InputList<Inputs.AssignmentItemArgs>? _items;

        /// <summary>
        /// List of resources included in this workspace manager assignment
        /// </summary>
        public InputList<Inputs.AssignmentItemArgs> Items
        {
            get => _items ?? (_items = new InputList<Inputs.AssignmentItemArgs>());
            set => _items = value;
        }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The resource name of the workspace manager group targeted by the workspace manager assignment
        /// </summary>
        [Input("targetResourceName", required: true)]
        public Input<string> TargetResourceName { get; set; } = null!;

        /// <summary>
        /// The name of the workspace manager assignment
        /// </summary>
        [Input("workspaceManagerAssignmentName")]
        public Input<string>? WorkspaceManagerAssignmentName { get; set; }

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public WorkspaceManagerAssignmentArgs()
        {
        }
        public static new WorkspaceManagerAssignmentArgs Empty => new WorkspaceManagerAssignmentArgs();
    }
}
