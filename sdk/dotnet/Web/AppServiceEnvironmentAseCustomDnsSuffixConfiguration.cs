// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web
{
    /// <summary>
    /// Full view of the custom domain suffix configuration for ASEv3.
    /// 
    /// Uses Azure REST API version 2024-04-01. In version 2.x of the Azure Native provider, it used API version 2022-09-01.
    /// 
    /// Other available API versions: 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:web:AppServiceEnvironmentAseCustomDnsSuffixConfiguration")]
    public partial class AppServiceEnvironmentAseCustomDnsSuffixConfiguration : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// The URL referencing the Azure Key Vault certificate secret that should be used as the default SSL/TLS certificate for sites with the custom domain suffix.
        /// </summary>
        [Output("certificateUrl")]
        public Output<string?> CertificateUrl { get; private set; } = null!;

        /// <summary>
        /// The default custom domain suffix to use for all sites deployed on the ASE.
        /// </summary>
        [Output("dnsSuffix")]
        public Output<string?> DnsSuffix { get; private set; } = null!;

        /// <summary>
        /// The user-assigned identity to use for resolving the key vault certificate reference. If not specified, the system-assigned ASE identity will be used if available.
        /// </summary>
        [Output("keyVaultReferenceIdentity")]
        public Output<string?> KeyVaultReferenceIdentity { get; private set; } = null!;

        /// <summary>
        /// Kind of resource.
        /// </summary>
        [Output("kind")]
        public Output<string?> Kind { get; private set; } = null!;

        /// <summary>
        /// Resource Name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("provisioningDetails")]
        public Output<string> ProvisioningDetails { get; private set; } = null!;

        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a AppServiceEnvironmentAseCustomDnsSuffixConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AppServiceEnvironmentAseCustomDnsSuffixConfiguration(string name, AppServiceEnvironmentAseCustomDnsSuffixConfigurationArgs args, CustomResourceOptions? options = null)
            : base("azure-native:web:AppServiceEnvironmentAseCustomDnsSuffixConfiguration", name, args ?? new AppServiceEnvironmentAseCustomDnsSuffixConfigurationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AppServiceEnvironmentAseCustomDnsSuffixConfiguration(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:web:AppServiceEnvironmentAseCustomDnsSuffixConfiguration", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:web/v20220301:AppServiceEnvironmentAseCustomDnsSuffixConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20220901:AppServiceEnvironmentAseCustomDnsSuffixConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20230101:AppServiceEnvironmentAseCustomDnsSuffixConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20231201:AppServiceEnvironmentAseCustomDnsSuffixConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20240401:AppServiceEnvironmentAseCustomDnsSuffixConfiguration" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20241101:AppServiceEnvironmentAseCustomDnsSuffixConfiguration" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AppServiceEnvironmentAseCustomDnsSuffixConfiguration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AppServiceEnvironmentAseCustomDnsSuffixConfiguration Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AppServiceEnvironmentAseCustomDnsSuffixConfiguration(name, id, options);
        }
    }

    public sealed class AppServiceEnvironmentAseCustomDnsSuffixConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The URL referencing the Azure Key Vault certificate secret that should be used as the default SSL/TLS certificate for sites with the custom domain suffix.
        /// </summary>
        [Input("certificateUrl")]
        public Input<string>? CertificateUrl { get; set; }

        /// <summary>
        /// The default custom domain suffix to use for all sites deployed on the ASE.
        /// </summary>
        [Input("dnsSuffix")]
        public Input<string>? DnsSuffix { get; set; }

        /// <summary>
        /// The user-assigned identity to use for resolving the key vault certificate reference. If not specified, the system-assigned ASE identity will be used if available.
        /// </summary>
        [Input("keyVaultReferenceIdentity")]
        public Input<string>? KeyVaultReferenceIdentity { get; set; }

        /// <summary>
        /// Kind of resource.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Name of the App Service Environment.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public AppServiceEnvironmentAseCustomDnsSuffixConfigurationArgs()
        {
        }
        public static new AppServiceEnvironmentAseCustomDnsSuffixConfigurationArgs Empty => new AppServiceEnvironmentAseCustomDnsSuffixConfigurationArgs();
    }
}
