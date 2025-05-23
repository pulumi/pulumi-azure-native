// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ScVmm
{
    /// <summary>
    /// The VirtualMachines resource definition.
    /// 
    /// Uses Azure REST API version 2023-04-01-preview. In version 2.x of the Azure Native provider, it used API version 2022-05-21-preview.
    /// 
    /// Other available API versions: 2022-05-21-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native scvmm [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:scvmm:VirtualMachine")]
    public partial class VirtualMachine : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Availability Sets in vm.
        /// </summary>
        [Output("availabilitySets")]
        public Output<ImmutableArray<Outputs.VirtualMachinePropertiesResponseAvailabilitySets>> AvailabilitySets { get; private set; } = null!;

        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Type of checkpoint supported for the vm.
        /// </summary>
        [Output("checkpointType")]
        public Output<string?> CheckpointType { get; private set; } = null!;

        /// <summary>
        /// Checkpoints in the vm.
        /// </summary>
        [Output("checkpoints")]
        public Output<ImmutableArray<Outputs.CheckpointResponse>> Checkpoints { get; private set; } = null!;

        /// <summary>
        /// ARM Id of the cloud resource to use for deploying the vm.
        /// </summary>
        [Output("cloudId")]
        public Output<string?> CloudId { get; private set; } = null!;

        /// <summary>
        /// The extended location.
        /// </summary>
        [Output("extendedLocation")]
        public Output<Outputs.ExtendedLocationResponse> ExtendedLocation { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the generation for the vm.
        /// </summary>
        [Output("generation")]
        public Output<int?> Generation { get; private set; } = null!;

        /// <summary>
        /// Guest agent status properties.
        /// </summary>
        [Output("guestAgentProfile")]
        public Output<Outputs.GuestAgentProfileResponse?> GuestAgentProfile { get; private set; } = null!;

        /// <summary>
        /// Hardware properties.
        /// </summary>
        [Output("hardwareProfile")]
        public Output<Outputs.HardwareProfileResponse?> HardwareProfile { get; private set; } = null!;

        /// <summary>
        /// The identity of the resource.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.IdentityResponse?> Identity { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the inventory Item ID for the resource.
        /// </summary>
        [Output("inventoryItemId")]
        public Output<string?> InventoryItemId { get; private set; } = null!;

        /// <summary>
        /// Last restored checkpoint in the vm.
        /// </summary>
        [Output("lastRestoredVMCheckpoint")]
        public Output<Outputs.CheckpointResponse> LastRestoredVMCheckpoint { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the location.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Resource Name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Network properties.
        /// </summary>
        [Output("networkProfile")]
        public Output<Outputs.NetworkProfileResponse?> NetworkProfile { get; private set; } = null!;

        /// <summary>
        /// OS properties.
        /// </summary>
        [Output("osProfile")]
        public Output<Outputs.OsProfileResponse?> OsProfile { get; private set; } = null!;

        /// <summary>
        /// Gets the power state of the virtual machine.
        /// </summary>
        [Output("powerState")]
        public Output<string> PowerState { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the provisioning state.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Storage properties.
        /// </summary>
        [Output("storageProfile")]
        public Output<Outputs.StorageProfileResponse?> StorageProfile { get; private set; } = null!;

        /// <summary>
        /// The system data.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// Resource tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// ARM Id of the template resource to use for deploying the vm.
        /// </summary>
        [Output("templateId")]
        public Output<string?> TemplateId { get; private set; } = null!;

        /// <summary>
        /// Resource Type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Unique ID of the virtual machine.
        /// </summary>
        [Output("uuid")]
        public Output<string?> Uuid { get; private set; } = null!;

        /// <summary>
        /// VMName is the name of VM on the SCVMM server.
        /// </summary>
        [Output("vmName")]
        public Output<string?> VmName { get; private set; } = null!;

        /// <summary>
        /// ARM Id of the vmmServer resource in which this resource resides.
        /// </summary>
        [Output("vmmServerId")]
        public Output<string?> VmmServerId { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualMachine resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualMachine(string name, VirtualMachineArgs args, CustomResourceOptions? options = null)
            : base("azure-native:scvmm:VirtualMachine", name, args ?? new VirtualMachineArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualMachine(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:scvmm:VirtualMachine", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:scvmm/v20200605preview:VirtualMachine" },
                    new global::Pulumi.Alias { Type = "azure-native:scvmm/v20220521preview:VirtualMachine" },
                    new global::Pulumi.Alias { Type = "azure-native:scvmm/v20230401preview:VirtualMachine" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VirtualMachine resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualMachine Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VirtualMachine(name, id, options);
        }
    }

    public sealed class VirtualMachineArgs : global::Pulumi.ResourceArgs
    {
        [Input("availabilitySets")]
        private InputList<Inputs.VirtualMachinePropertiesAvailabilitySetsArgs>? _availabilitySets;

        /// <summary>
        /// Availability Sets in vm.
        /// </summary>
        public InputList<Inputs.VirtualMachinePropertiesAvailabilitySetsArgs> AvailabilitySets
        {
            get => _availabilitySets ?? (_availabilitySets = new InputList<Inputs.VirtualMachinePropertiesAvailabilitySetsArgs>());
            set => _availabilitySets = value;
        }

        /// <summary>
        /// Type of checkpoint supported for the vm.
        /// </summary>
        [Input("checkpointType")]
        public Input<string>? CheckpointType { get; set; }

        [Input("checkpoints")]
        private InputList<Inputs.CheckpointArgs>? _checkpoints;

        /// <summary>
        /// Checkpoints in the vm.
        /// </summary>
        public InputList<Inputs.CheckpointArgs> Checkpoints
        {
            get => _checkpoints ?? (_checkpoints = new InputList<Inputs.CheckpointArgs>());
            set => _checkpoints = value;
        }

        /// <summary>
        /// ARM Id of the cloud resource to use for deploying the vm.
        /// </summary>
        [Input("cloudId")]
        public Input<string>? CloudId { get; set; }

        /// <summary>
        /// The extended location.
        /// </summary>
        [Input("extendedLocation", required: true)]
        public Input<Inputs.ExtendedLocationArgs> ExtendedLocation { get; set; } = null!;

        /// <summary>
        /// Gets or sets the generation for the vm.
        /// </summary>
        [Input("generation")]
        public Input<int>? Generation { get; set; }

        /// <summary>
        /// Guest agent status properties.
        /// </summary>
        [Input("guestAgentProfile")]
        public Input<Inputs.GuestAgentProfileArgs>? GuestAgentProfile { get; set; }

        /// <summary>
        /// Hardware properties.
        /// </summary>
        [Input("hardwareProfile")]
        public Input<Inputs.HardwareProfileArgs>? HardwareProfile { get; set; }

        /// <summary>
        /// The identity of the resource.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.IdentityArgs>? Identity { get; set; }

        /// <summary>
        /// Gets or sets the inventory Item ID for the resource.
        /// </summary>
        [Input("inventoryItemId")]
        public Input<string>? InventoryItemId { get; set; }

        /// <summary>
        /// Gets or sets the location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Network properties.
        /// </summary>
        [Input("networkProfile")]
        public Input<Inputs.NetworkProfileArgs>? NetworkProfile { get; set; }

        /// <summary>
        /// OS properties.
        /// </summary>
        [Input("osProfile")]
        public Input<Inputs.OsProfileArgs>? OsProfile { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Storage properties.
        /// </summary>
        [Input("storageProfile")]
        public Input<Inputs.StorageProfileArgs>? StorageProfile { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// ARM Id of the template resource to use for deploying the vm.
        /// </summary>
        [Input("templateId")]
        public Input<string>? TemplateId { get; set; }

        /// <summary>
        /// Unique ID of the virtual machine.
        /// </summary>
        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        /// <summary>
        /// Name of the VirtualMachine.
        /// </summary>
        [Input("virtualMachineName")]
        public Input<string>? VirtualMachineName { get; set; }

        /// <summary>
        /// VMName is the name of VM on the SCVMM server.
        /// </summary>
        [Input("vmName")]
        public Input<string>? VmName { get; set; }

        /// <summary>
        /// ARM Id of the vmmServer resource in which this resource resides.
        /// </summary>
        [Input("vmmServerId")]
        public Input<string>? VmmServerId { get; set; }

        public VirtualMachineArgs()
        {
        }
        public static new VirtualMachineArgs Empty => new VirtualMachineArgs();
    }
}
