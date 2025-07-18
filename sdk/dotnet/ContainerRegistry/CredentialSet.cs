// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerRegistry
{
    /// <summary>
    /// An object that represents a credential set resource for a container registry.
    /// 
    /// Uses Azure REST API version 2023-01-01-preview. In version 2.x of the Azure Native provider, it used API version 2023-01-01-preview.
    /// 
    /// Other available API versions: 2023-06-01-preview, 2023-07-01, 2023-08-01-preview, 2023-11-01-preview, 2024-11-01-preview, 2025-03-01-preview, 2025-04-01, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native containerregistry [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:containerregistry:CredentialSet")]
    public partial class CredentialSet : global::Pulumi.CustomResource
    {
        /// <summary>
        /// List of authentication credentials stored for an upstream.
        /// Usually consists of a primary and an optional secondary credential.
        /// </summary>
        [Output("authCredentials")]
        public Output<ImmutableArray<Outputs.AuthCredentialResponse>> AuthCredentials { get; private set; } = null!;

        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// The creation date of credential store resource.
        /// </summary>
        [Output("creationDate")]
        public Output<string> CreationDate { get; private set; } = null!;

        /// <summary>
        /// Identities associated with the resource. This is used to access the KeyVault secrets.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.IdentityPropertiesResponse?> Identity { get; private set; } = null!;

        /// <summary>
        /// The credentials are stored for this upstream or login server.
        /// </summary>
        [Output("loginServer")]
        public Output<string?> LoginServer { get; private set; } = null!;

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Provisioning state of the resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Metadata pertaining to creation and last modification of the resource.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// The type of the resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a CredentialSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CredentialSet(string name, CredentialSetArgs args, CustomResourceOptions? options = null)
            : base("azure-native:containerregistry:CredentialSet", name, args ?? new CredentialSetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CredentialSet(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:containerregistry:CredentialSet", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:containerregistry/v20230101preview:CredentialSet" },
                    new global::Pulumi.Alias { Type = "azure-native:containerregistry/v20230601preview:CredentialSet" },
                    new global::Pulumi.Alias { Type = "azure-native:containerregistry/v20230701:CredentialSet" },
                    new global::Pulumi.Alias { Type = "azure-native:containerregistry/v20230801preview:CredentialSet" },
                    new global::Pulumi.Alias { Type = "azure-native:containerregistry/v20231101preview:CredentialSet" },
                    new global::Pulumi.Alias { Type = "azure-native:containerregistry/v20241101preview:CredentialSet" },
                    new global::Pulumi.Alias { Type = "azure-native:containerregistry/v20250301preview:CredentialSet" },
                    new global::Pulumi.Alias { Type = "azure-native:containerregistry/v20250401:CredentialSet" },
                    new global::Pulumi.Alias { Type = "azure-native:containerregistry/v20250501preview:CredentialSet" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CredentialSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CredentialSet Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new CredentialSet(name, id, options);
        }
    }

    public sealed class CredentialSetArgs : global::Pulumi.ResourceArgs
    {
        [Input("authCredentials")]
        private InputList<Inputs.AuthCredentialArgs>? _authCredentials;

        /// <summary>
        /// List of authentication credentials stored for an upstream.
        /// Usually consists of a primary and an optional secondary credential.
        /// </summary>
        public InputList<Inputs.AuthCredentialArgs> AuthCredentials
        {
            get => _authCredentials ?? (_authCredentials = new InputList<Inputs.AuthCredentialArgs>());
            set => _authCredentials = value;
        }

        /// <summary>
        /// The name of the credential set.
        /// </summary>
        [Input("credentialSetName")]
        public Input<string>? CredentialSetName { get; set; }

        /// <summary>
        /// Identities associated with the resource. This is used to access the KeyVault secrets.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.IdentityPropertiesArgs>? Identity { get; set; }

        /// <summary>
        /// The credentials are stored for this upstream or login server.
        /// </summary>
        [Input("loginServer")]
        public Input<string>? LoginServer { get; set; }

        /// <summary>
        /// The name of the container registry.
        /// </summary>
        [Input("registryName", required: true)]
        public Input<string> RegistryName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public CredentialSetArgs()
        {
        }
        public static new CredentialSetArgs Empty => new CredentialSetArgs();
    }
}
