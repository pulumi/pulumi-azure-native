// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ManagedNetworkFabric
{
    /// <summary>
    /// The IP Community resource definition.
    /// 
    /// Uses Azure REST API version 2023-06-15. In version 2.x of the Azure Native provider, it used API version 2023-02-01-preview.
    /// 
    /// Other available API versions: 2023-02-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native managednetworkfabric [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:managednetworkfabric:IpCommunity")]
    public partial class IpCommunity : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Administrative state of the resource.
        /// </summary>
        [Output("administrativeState")]
        public Output<string> AdministrativeState { get; private set; } = null!;

        /// <summary>
        /// Switch configuration description.
        /// </summary>
        [Output("annotation")]
        public Output<string?> Annotation { get; private set; } = null!;

        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Configuration state of the resource.
        /// </summary>
        [Output("configurationState")]
        public Output<string> ConfigurationState { get; private set; } = null!;

        /// <summary>
        /// List of IP Community Rules.
        /// </summary>
        [Output("ipCommunityRules")]
        public Output<ImmutableArray<Outputs.IpCommunityRuleResponse>> IpCommunityRules { get; private set; } = null!;

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
        /// Provisioning state of the resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

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
        /// Create a IpCommunity resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IpCommunity(string name, IpCommunityArgs args, CustomResourceOptions? options = null)
            : base("azure-native:managednetworkfabric:IpCommunity", name, args ?? new IpCommunityArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IpCommunity(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:managednetworkfabric:IpCommunity", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:managednetworkfabric/v20230201preview:IpCommunity" },
                    new global::Pulumi.Alias { Type = "azure-native:managednetworkfabric/v20230615:IpCommunity" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing IpCommunity resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IpCommunity Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new IpCommunity(name, id, options);
        }
    }

    public sealed class IpCommunityArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Switch configuration description.
        /// </summary>
        [Input("annotation")]
        public Input<string>? Annotation { get; set; }

        /// <summary>
        /// Name of the IP Community.
        /// </summary>
        [Input("ipCommunityName")]
        public Input<string>? IpCommunityName { get; set; }

        [Input("ipCommunityRules", required: true)]
        private InputList<Inputs.IpCommunityRuleArgs>? _ipCommunityRules;

        /// <summary>
        /// List of IP Community Rules.
        /// </summary>
        public InputList<Inputs.IpCommunityRuleArgs> IpCommunityRules
        {
            get => _ipCommunityRules ?? (_ipCommunityRules = new InputList<Inputs.IpCommunityRuleArgs>());
            set => _ipCommunityRules = value;
        }

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

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

        public IpCommunityArgs()
        {
        }
        public static new IpCommunityArgs Empty => new IpCommunityArgs();
    }
}
