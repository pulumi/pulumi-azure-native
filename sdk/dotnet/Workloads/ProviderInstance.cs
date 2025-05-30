// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Workloads
{
    /// <summary>
    /// A provider instance associated with SAP monitor.
    /// 
    /// Uses Azure REST API version 2024-02-01-preview. In version 2.x of the Azure Native provider, it used API version 2023-04-01.
    /// 
    /// Other available API versions: 2023-04-01, 2023-10-01-preview, 2023-12-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native workloads [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:workloads:ProviderInstance")]
    public partial class ProviderInstance : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Defines the provider instance errors.
        /// </summary>
        [Output("errors")]
        public Output<Outputs.ErrorDetailResponse> Errors { get; private set; } = null!;

        /// <summary>
        /// Resource health details
        /// </summary>
        [Output("health")]
        public Output<Outputs.HealthResponse> Health { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Defines the provider specific properties.
        /// </summary>
        [Output("providerSettings")]
        public Output<object?> ProviderSettings { get; private set; } = null!;

        /// <summary>
        /// State of provisioning of the provider instance
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

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
        /// Create a ProviderInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProviderInstance(string name, ProviderInstanceArgs args, CustomResourceOptions? options = null)
            : base("azure-native:workloads:ProviderInstance", name, args ?? new ProviderInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProviderInstance(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:workloads:ProviderInstance", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:workloads/v20211201preview:ProviderInstance" },
                    new global::Pulumi.Alias { Type = "azure-native:workloads/v20221101preview:ProviderInstance" },
                    new global::Pulumi.Alias { Type = "azure-native:workloads/v20230401:ProviderInstance" },
                    new global::Pulumi.Alias { Type = "azure-native:workloads/v20231001preview:ProviderInstance" },
                    new global::Pulumi.Alias { Type = "azure-native:workloads/v20231201preview:ProviderInstance" },
                    new global::Pulumi.Alias { Type = "azure-native:workloads/v20240201preview:ProviderInstance" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ProviderInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProviderInstance Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ProviderInstance(name, id, options);
        }
    }

    public sealed class ProviderInstanceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the SAP monitor resource.
        /// </summary>
        [Input("monitorName", required: true)]
        public Input<string> MonitorName { get; set; } = null!;

        /// <summary>
        /// Name of the provider instance.
        /// </summary>
        [Input("providerInstanceName")]
        public Input<string>? ProviderInstanceName { get; set; }

        /// <summary>
        /// Defines the provider specific properties.
        /// </summary>
        [Input("providerSettings")]
        public object? ProviderSettings { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public ProviderInstanceArgs()
        {
        }
        public static new ProviderInstanceArgs Empty => new ProviderInstanceArgs();
    }
}
