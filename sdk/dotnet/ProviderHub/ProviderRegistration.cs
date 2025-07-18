// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ProviderHub
{
    /// <summary>
    /// Uses Azure REST API version 2024-09-01. In version 2.x of the Azure Native provider, it used API version 2021-09-01-preview.
    /// 
    /// Other available API versions: 2021-09-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native providerhub [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:providerhub:ProviderRegistration")]
    public partial class ProviderRegistration : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Provider registration kind. This Metadata is also used by portal/tooling/etc to render different UX experiences for resources of the same type.
        /// </summary>
        [Output("kind")]
        public Output<string?> Kind { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("properties")]
        public Output<Outputs.ProviderRegistrationPropertiesResponse> Properties { get; private set; } = null!;

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
        /// Create a ProviderRegistration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProviderRegistration(string name, ProviderRegistrationArgs? args = null, CustomResourceOptions? options = null)
            : base("azure-native:providerhub:ProviderRegistration", name, args ?? new ProviderRegistrationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProviderRegistration(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:providerhub:ProviderRegistration", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:providerhub/v20201120:ProviderRegistration" },
                    new global::Pulumi.Alias { Type = "azure-native:providerhub/v20210501preview:ProviderRegistration" },
                    new global::Pulumi.Alias { Type = "azure-native:providerhub/v20210601preview:ProviderRegistration" },
                    new global::Pulumi.Alias { Type = "azure-native:providerhub/v20210901preview:ProviderRegistration" },
                    new global::Pulumi.Alias { Type = "azure-native:providerhub/v20240901:ProviderRegistration" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ProviderRegistration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProviderRegistration Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ProviderRegistration(name, id, options);
        }
    }

    public sealed class ProviderRegistrationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Provider registration kind. This Metadata is also used by portal/tooling/etc to render different UX experiences for resources of the same type.
        /// </summary>
        [Input("kind")]
        public InputUnion<string, Pulumi.AzureNative.ProviderHub.ProviderRegistrationKind>? Kind { get; set; }

        [Input("properties")]
        public Input<Inputs.ProviderRegistrationPropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// The name of the resource provider hosted within ProviderHub.
        /// </summary>
        [Input("providerNamespace")]
        public Input<string>? ProviderNamespace { get; set; }

        public ProviderRegistrationArgs()
        {
            Kind = "Managed";
        }
        public static new ProviderRegistrationArgs Empty => new ProviderRegistrationArgs();
    }
}
