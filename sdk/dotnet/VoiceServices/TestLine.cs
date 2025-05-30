// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.VoiceServices
{
    /// <summary>
    /// A TestLine resource
    /// 
    /// Uses Azure REST API version 2023-09-01. In version 2.x of the Azure Native provider, it used API version 2023-04-03.
    /// 
    /// Other available API versions: 2022-12-01-preview, 2023-01-31, 2023-04-03. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native voiceservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:voiceservices:TestLine")]
    public partial class TestLine : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

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
        /// The phone number
        /// </summary>
        [Output("phoneNumber")]
        public Output<string> PhoneNumber { get; private set; } = null!;

        /// <summary>
        /// Resource provisioning state.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Purpose of this test line, e.g. automated or manual testing
        /// </summary>
        [Output("purpose")]
        public Output<string> Purpose { get; private set; } = null!;

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
        /// Create a TestLine resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TestLine(string name, TestLineArgs args, CustomResourceOptions? options = null)
            : base("azure-native:voiceservices:TestLine", name, args ?? new TestLineArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TestLine(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:voiceservices:TestLine", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:voiceservices/v20221201preview:TestLine" },
                    new global::Pulumi.Alias { Type = "azure-native:voiceservices/v20230131:TestLine" },
                    new global::Pulumi.Alias { Type = "azure-native:voiceservices/v20230403:TestLine" },
                    new global::Pulumi.Alias { Type = "azure-native:voiceservices/v20230901:TestLine" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing TestLine resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TestLine Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new TestLine(name, id, options);
        }
    }

    public sealed class TestLineArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Unique identifier for this deployment
        /// </summary>
        [Input("communicationsGatewayName", required: true)]
        public Input<string> CommunicationsGatewayName { get; set; } = null!;

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The phone number
        /// </summary>
        [Input("phoneNumber", required: true)]
        public Input<string> PhoneNumber { get; set; } = null!;

        /// <summary>
        /// Purpose of this test line, e.g. automated or manual testing
        /// </summary>
        [Input("purpose", required: true)]
        public InputUnion<string, Pulumi.AzureNative.VoiceServices.TestLinePurpose> Purpose { get; set; } = null!;

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

        /// <summary>
        /// Unique identifier for this test line
        /// </summary>
        [Input("testLineName")]
        public Input<string>? TestLineName { get; set; }

        public TestLineArgs()
        {
        }
        public static new TestLineArgs Empty => new TestLineArgs();
    }
}
