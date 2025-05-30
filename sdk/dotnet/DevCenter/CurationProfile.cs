// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DevCenter
{
    /// <summary>
    /// Represents an curation profile resource.
    /// 
    /// Uses Azure REST API version 2024-10-01-preview. In version 2.x of the Azure Native provider, it used API version 2024-08-01-preview.
    /// 
    /// Other available API versions: 2024-08-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native devcenter [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:devcenter:CurationProfile")]
    public partial class CurationProfile : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Resource policies that are a part of this curation profile.
        /// </summary>
        [Output("resourcePolicies")]
        public Output<ImmutableArray<Outputs.ResourcePolicyResponse>> ResourcePolicies { get; private set; } = null!;

        /// <summary>
        /// Resources that have access to the shared resources that are a part of this curation profile.
        /// </summary>
        [Output("scopes")]
        public Output<ImmutableArray<string>> Scopes { get; private set; } = null!;

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
        /// Create a CurationProfile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CurationProfile(string name, CurationProfileArgs args, CustomResourceOptions? options = null)
            : base("azure-native:devcenter:CurationProfile", name, args ?? new CurationProfileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CurationProfile(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:devcenter:CurationProfile", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:devcenter/v20240801preview:CurationProfile" },
                    new global::Pulumi.Alias { Type = "azure-native:devcenter/v20241001preview:CurationProfile" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CurationProfile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CurationProfile Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new CurationProfile(name, id, options);
        }
    }

    public sealed class CurationProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the curation profile.
        /// </summary>
        [Input("curationProfileName")]
        public Input<string>? CurationProfileName { get; set; }

        /// <summary>
        /// The name of the devcenter.
        /// </summary>
        [Input("devCenterName", required: true)]
        public Input<string> DevCenterName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("resourcePolicies")]
        private InputList<Inputs.ResourcePolicyArgs>? _resourcePolicies;

        /// <summary>
        /// Resource policies that are a part of this curation profile.
        /// </summary>
        public InputList<Inputs.ResourcePolicyArgs> ResourcePolicies
        {
            get => _resourcePolicies ?? (_resourcePolicies = new InputList<Inputs.ResourcePolicyArgs>());
            set => _resourcePolicies = value;
        }

        [Input("scopes")]
        private InputList<string>? _scopes;

        /// <summary>
        /// Resources that have access to the shared resources that are a part of this curation profile.
        /// </summary>
        public InputList<string> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<string>());
            set => _scopes = value;
        }

        public CurationProfileArgs()
        {
        }
        public static new CurationProfileArgs Empty => new CurationProfileArgs();
    }
}
