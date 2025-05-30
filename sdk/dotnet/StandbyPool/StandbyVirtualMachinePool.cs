// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.StandbyPool
{
    /// <summary>
    /// A StandbyVirtualMachinePoolResource.
    /// 
    /// Uses Azure REST API version 2024-03-01. In version 2.x of the Azure Native provider, it used API version 2023-12-01-preview.
    /// 
    /// Other available API versions: 2023-12-01-preview, 2024-03-01-preview, 2025-03-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native standbypool [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:standbypool:StandbyVirtualMachinePool")]
    public partial class StandbyVirtualMachinePool : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the fully qualified resource ID of a virtual machine scale set the pool is attached to.
        /// </summary>
        [Output("attachedVirtualMachineScaleSetId")]
        public Output<string?> AttachedVirtualMachineScaleSetId { get; private set; } = null!;

        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Specifies the elasticity profile of the standby virtual machine pools.
        /// </summary>
        [Output("elasticityProfile")]
        public Output<Outputs.StandbyVirtualMachinePoolElasticityProfileResponse?> ElasticityProfile { get; private set; } = null!;

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
        /// The status of the last operation.
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
        /// Specifies the desired state of virtual machines in the pool.
        /// </summary>
        [Output("virtualMachineState")]
        public Output<string> VirtualMachineState { get; private set; } = null!;


        /// <summary>
        /// Create a StandbyVirtualMachinePool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StandbyVirtualMachinePool(string name, StandbyVirtualMachinePoolArgs args, CustomResourceOptions? options = null)
            : base("azure-native:standbypool:StandbyVirtualMachinePool", name, args ?? new StandbyVirtualMachinePoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StandbyVirtualMachinePool(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:standbypool:StandbyVirtualMachinePool", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:standbypool/v20231201preview:StandbyVirtualMachinePool" },
                    new global::Pulumi.Alias { Type = "azure-native:standbypool/v20240301:StandbyVirtualMachinePool" },
                    new global::Pulumi.Alias { Type = "azure-native:standbypool/v20240301preview:StandbyVirtualMachinePool" },
                    new global::Pulumi.Alias { Type = "azure-native:standbypool/v20250301:StandbyVirtualMachinePool" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing StandbyVirtualMachinePool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StandbyVirtualMachinePool Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new StandbyVirtualMachinePool(name, id, options);
        }
    }

    public sealed class StandbyVirtualMachinePoolArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the fully qualified resource ID of a virtual machine scale set the pool is attached to.
        /// </summary>
        [Input("attachedVirtualMachineScaleSetId")]
        public Input<string>? AttachedVirtualMachineScaleSetId { get; set; }

        /// <summary>
        /// Specifies the elasticity profile of the standby virtual machine pools.
        /// </summary>
        [Input("elasticityProfile")]
        public Input<Inputs.StandbyVirtualMachinePoolElasticityProfileArgs>? ElasticityProfile { get; set; }

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

        /// <summary>
        /// Name of the standby virtual machine pool
        /// </summary>
        [Input("standbyVirtualMachinePoolName")]
        public Input<string>? StandbyVirtualMachinePoolName { get; set; }

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
        /// Specifies the desired state of virtual machines in the pool.
        /// </summary>
        [Input("virtualMachineState", required: true)]
        public InputUnion<string, Pulumi.AzureNative.StandbyPool.VirtualMachineState> VirtualMachineState { get; set; } = null!;

        public StandbyVirtualMachinePoolArgs()
        {
        }
        public static new StandbyVirtualMachinePoolArgs Empty => new StandbyVirtualMachinePoolArgs();
    }
}
