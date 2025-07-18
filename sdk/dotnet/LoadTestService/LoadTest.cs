// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.LoadTestService
{
    /// <summary>
    /// LoadTest details.
    /// 
    /// Uses Azure REST API version 2023-12-01-preview. In version 2.x of the Azure Native provider, it used API version 2022-12-01.
    /// 
    /// Other available API versions: 2022-12-01, 2024-12-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native loadtestservice [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:loadtestservice:LoadTest")]
    public partial class LoadTest : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Resource data plane URI.
        /// </summary>
        [Output("dataPlaneURI")]
        public Output<string> DataPlaneURI { get; private set; } = null!;

        /// <summary>
        /// Description of the resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// CMK Encryption property.
        /// </summary>
        [Output("encryption")]
        public Output<Outputs.EncryptionPropertiesResponse?> Encryption { get; private set; } = null!;

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
        /// Resource provisioning state.
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
        /// Create a LoadTest resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LoadTest(string name, LoadTestArgs args, CustomResourceOptions? options = null)
            : base("azure-native:loadtestservice:LoadTest", name, args ?? new LoadTestArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LoadTest(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:loadtestservice:LoadTest", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:loadtestservice/v20211201preview:LoadTest" },
                    new global::Pulumi.Alias { Type = "azure-native:loadtestservice/v20221201:LoadTest" },
                    new global::Pulumi.Alias { Type = "azure-native:loadtestservice/v20231201preview:LoadTest" },
                    new global::Pulumi.Alias { Type = "azure-native:loadtestservice/v20241201preview:LoadTest" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing LoadTest resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LoadTest Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new LoadTest(name, id, options);
        }
    }

    public sealed class LoadTestArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// CMK Encryption property.
        /// </summary>
        [Input("encryption")]
        public Input<Inputs.EncryptionPropertiesArgs>? Encryption { get; set; }

        /// <summary>
        /// The managed service identities assigned to this resource.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.ManagedServiceIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// Load Test name
        /// </summary>
        [Input("loadTestName")]
        public Input<string>? LoadTestName { get; set; }

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

        public LoadTestArgs()
        {
        }
        public static new LoadTestArgs Empty => new LoadTestArgs();
    }
}
