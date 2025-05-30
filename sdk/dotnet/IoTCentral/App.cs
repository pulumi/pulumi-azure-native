// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.IoTCentral
{
    /// <summary>
    /// The IoT Central application.
    /// 
    /// Uses Azure REST API version 2021-11-01-preview. In version 2.x of the Azure Native provider, it used API version 2021-06-01.
    /// 
    /// Other available API versions: 2021-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native iotcentral [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:iotcentral:App")]
    public partial class App : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the application.
        /// </summary>
        [Output("applicationId")]
        public Output<string> ApplicationId { get; private set; } = null!;

        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// The display name of the application.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The managed identities for the IoT Central application.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.SystemAssignedServiceIdentityResponse?> Identity { get; private set; } = null!;

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
        /// Network Rule Set Properties of this IoT Central application.
        /// </summary>
        [Output("networkRuleSets")]
        public Output<Outputs.NetworkRuleSetsResponse?> NetworkRuleSets { get; private set; } = null!;

        /// <summary>
        /// Private endpoint connections created on this IoT Central application.
        /// </summary>
        [Output("privateEndpointConnections")]
        public Output<ImmutableArray<Outputs.PrivateEndpointConnectionResponse>> PrivateEndpointConnections { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the application.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Whether requests from the public network are allowed.
        /// </summary>
        [Output("publicNetworkAccess")]
        public Output<string?> PublicNetworkAccess { get; private set; } = null!;

        /// <summary>
        /// A valid instance SKU.
        /// </summary>
        [Output("sku")]
        public Output<Outputs.AppSkuInfoResponse> Sku { get; private set; } = null!;

        /// <summary>
        /// The current state of the application.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The subdomain of the application.
        /// </summary>
        [Output("subdomain")]
        public Output<string?> Subdomain { get; private set; } = null!;

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
        /// The ID of the application template, which is a blueprint that defines the characteristics and behaviors of an application. Optional; if not specified, defaults to a blank blueprint and allows the application to be defined from scratch.
        /// </summary>
        [Output("template")]
        public Output<string?> Template { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a App resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public App(string name, AppArgs args, CustomResourceOptions? options = null)
            : base("azure-native:iotcentral:App", name, args ?? new AppArgs(), MakeResourceOptions(options, ""))
        {
        }

        private App(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:iotcentral:App", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:iotcentral/v20180901:App" },
                    new global::Pulumi.Alias { Type = "azure-native:iotcentral/v20210601:App" },
                    new global::Pulumi.Alias { Type = "azure-native:iotcentral/v20211101preview:App" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing App resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static App Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new App(name, id, options);
        }
    }

    public sealed class AppArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The display name of the application.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The managed identities for the IoT Central application.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.SystemAssignedServiceIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Network Rule Set Properties of this IoT Central application.
        /// </summary>
        [Input("networkRuleSets")]
        public Input<Inputs.NetworkRuleSetsArgs>? NetworkRuleSets { get; set; }

        /// <summary>
        /// Whether requests from the public network are allowed.
        /// </summary>
        [Input("publicNetworkAccess")]
        public InputUnion<string, Pulumi.AzureNative.IoTCentral.PublicNetworkAccess>? PublicNetworkAccess { get; set; }

        /// <summary>
        /// The name of the resource group that contains the IoT Central application.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The ARM resource name of the IoT Central application.
        /// </summary>
        [Input("resourceName")]
        public Input<string>? ResourceName { get; set; }

        /// <summary>
        /// A valid instance SKU.
        /// </summary>
        [Input("sku", required: true)]
        public Input<Inputs.AppSkuInfoArgs> Sku { get; set; } = null!;

        /// <summary>
        /// The subdomain of the application.
        /// </summary>
        [Input("subdomain")]
        public Input<string>? Subdomain { get; set; }

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
        /// The ID of the application template, which is a blueprint that defines the characteristics and behaviors of an application. Optional; if not specified, defaults to a blank blueprint and allows the application to be defined from scratch.
        /// </summary>
        [Input("template")]
        public Input<string>? Template { get; set; }

        public AppArgs()
        {
        }
        public static new AppArgs Empty => new AppArgs();
    }
}
