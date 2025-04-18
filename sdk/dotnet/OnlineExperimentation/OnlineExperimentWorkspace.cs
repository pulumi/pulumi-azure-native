// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.OnlineExperimentation
{
    /// <summary>
    /// An online experiment workspace resource.
    /// 
    /// Uses Azure REST API version 2025-05-31-preview.
    /// </summary>
    [AzureNativeResourceType("azure-native:onlineexperimentation:OnlineExperimentWorkspace")]
    public partial class OnlineExperimentWorkspace : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// The managed service identities assigned to this resource.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.ManagedServiceIdentityResponse?> Identity { get; private set; } = null!;

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The resource-specific properties for this resource.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.OnlineExperimentWorkspacePropertiesResponse> Properties { get; private set; } = null!;

        /// <summary>
        /// The SKU (Stock Keeping Unit) assigned to this resource.
        /// </summary>
        [Output("sku")]
        public Output<Outputs.OnlineExperimentationWorkspaceSkuResponse?> Sku { get; private set; } = null!;

        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a OnlineExperimentWorkspace resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OnlineExperimentWorkspace(string name, OnlineExperimentWorkspaceArgs args, CustomResourceOptions? options = null)
            : base("azure-native:onlineexperimentation:OnlineExperimentWorkspace", name, args ?? new OnlineExperimentWorkspaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OnlineExperimentWorkspace(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:onlineexperimentation:OnlineExperimentWorkspace", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:onlineexperimentation/v20250531preview:OnlineExperimentWorkspace" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing OnlineExperimentWorkspace resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OnlineExperimentWorkspace Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new OnlineExperimentWorkspace(name, id, options);
        }
    }

    public sealed class OnlineExperimentWorkspaceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The managed service identities assigned to this resource.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.ManagedServiceIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The resource-specific properties for this resource.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.OnlineExperimentWorkspacePropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The SKU (Stock Keeping Unit) assigned to this resource.
        /// </summary>
        [Input("sku")]
        public Input<Inputs.OnlineExperimentationWorkspaceSkuArgs>? Sku { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The name of the OnlineExperimentWorkspace
        /// </summary>
        [Input("workspaceName")]
        public Input<string>? WorkspaceName { get; set; }

        public OnlineExperimentWorkspaceArgs()
        {
        }
        public static new OnlineExperimentWorkspaceArgs Empty => new OnlineExperimentWorkspaceArgs();
    }
}
