// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Maps
{
    /// <summary>
    /// An Azure resource which represents access to a suite of Maps REST APIs.
    /// 
    /// Uses Azure REST API version 2024-07-01-preview. In version 2.x of the Azure Native provider, it used API version 2021-02-01.
    /// 
    /// Other available API versions: 2020-02-01-preview, 2021-02-01, 2021-07-01-preview, 2021-12-01-preview, 2023-06-01, 2023-08-01-preview, 2023-12-01-preview, 2024-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native maps [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:maps:Account")]
    public partial class Account : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Managed service identity (system assigned and/or user assigned identities)
        /// </summary>
        [Output("identity")]
        public Output<Outputs.ManagedServiceIdentityResponse?> Identity { get; private set; } = null!;

        /// <summary>
        /// Get or Set Kind property.
        /// </summary>
        [Output("kind")]
        public Output<string?> Kind { get; private set; } = null!;

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
        /// The map account properties.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.MapsAccountPropertiesResponse> Properties { get; private set; } = null!;

        /// <summary>
        /// The SKU of this account.
        /// </summary>
        [Output("sku")]
        public Output<Outputs.SkuResponse> Sku { get; private set; } = null!;

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
        /// Create a Account resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Account(string name, AccountArgs args, CustomResourceOptions? options = null)
            : base("azure-native:maps:Account", name, args ?? new AccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Account(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:maps:Account", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:maps/v20170101preview:Account" },
                    new global::Pulumi.Alias { Type = "azure-native:maps/v20180501:Account" },
                    new global::Pulumi.Alias { Type = "azure-native:maps/v20200201preview:Account" },
                    new global::Pulumi.Alias { Type = "azure-native:maps/v20210201:Account" },
                    new global::Pulumi.Alias { Type = "azure-native:maps/v20210701preview:Account" },
                    new global::Pulumi.Alias { Type = "azure-native:maps/v20211201preview:Account" },
                    new global::Pulumi.Alias { Type = "azure-native:maps/v20230601:Account" },
                    new global::Pulumi.Alias { Type = "azure-native:maps/v20230801preview:Account" },
                    new global::Pulumi.Alias { Type = "azure-native:maps/v20231201preview:Account" },
                    new global::Pulumi.Alias { Type = "azure-native:maps/v20240101preview:Account" },
                    new global::Pulumi.Alias { Type = "azure-native:maps/v20240701preview:Account" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Account resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Account Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Account(name, id, options);
        }
    }

    public sealed class AccountArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Maps Account.
        /// </summary>
        [Input("accountName")]
        public Input<string>? AccountName { get; set; }

        /// <summary>
        /// Managed service identity (system assigned and/or user assigned identities)
        /// </summary>
        [Input("identity")]
        public Input<Inputs.ManagedServiceIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// Get or Set Kind property.
        /// </summary>
        [Input("kind")]
        public InputUnion<string, Pulumi.AzureNative.Maps.Kind>? Kind { get; set; }

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The map account properties.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.MapsAccountPropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The SKU of this account.
        /// </summary>
        [Input("sku", required: true)]
        public Input<Inputs.SkuArgs> Sku { get; set; } = null!;

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

        public AccountArgs()
        {
        }
        public static new AccountArgs Empty => new AccountArgs();
    }
}
