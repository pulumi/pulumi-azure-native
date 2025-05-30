// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DevTestLab.Inputs
{

    /// <summary>
    /// Properties for creating a virtual machine.
    /// </summary>
    public sealed class LabVirtualMachineCreationParameterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether another user can take ownership of the virtual machine
        /// </summary>
        [Input("allowClaim")]
        public Input<bool>? AllowClaim { get; set; }

        [Input("artifacts")]
        private InputList<Inputs.ArtifactInstallPropertiesArgs>? _artifacts;

        /// <summary>
        /// The artifacts to be installed on the virtual machine.
        /// </summary>
        public InputList<Inputs.ArtifactInstallPropertiesArgs> Artifacts
        {
            get => _artifacts ?? (_artifacts = new InputList<Inputs.ArtifactInstallPropertiesArgs>());
            set => _artifacts = value;
        }

        /// <summary>
        /// The number of virtual machine instances to create.
        /// </summary>
        [Input("bulkCreationParameters")]
        public Input<Inputs.BulkCreationParametersArgs>? BulkCreationParameters { get; set; }

        /// <summary>
        /// The creation date of the virtual machine.
        /// </summary>
        [Input("createdDate")]
        public Input<string>? CreatedDate { get; set; }

        /// <summary>
        /// The custom image identifier of the virtual machine.
        /// </summary>
        [Input("customImageId")]
        public Input<string>? CustomImageId { get; set; }

        [Input("dataDiskParameters")]
        private InputList<Inputs.DataDiskPropertiesArgs>? _dataDiskParameters;

        /// <summary>
        /// New or existing data disks to attach to the virtual machine after creation
        /// </summary>
        public InputList<Inputs.DataDiskPropertiesArgs> DataDiskParameters
        {
            get => _dataDiskParameters ?? (_dataDiskParameters = new InputList<Inputs.DataDiskPropertiesArgs>());
            set => _dataDiskParameters = value;
        }

        /// <summary>
        /// Indicates whether the virtual machine is to be created without a public IP address.
        /// </summary>
        [Input("disallowPublicIpAddress")]
        public Input<bool>? DisallowPublicIpAddress { get; set; }

        /// <summary>
        /// The resource ID of the environment that contains this virtual machine, if any.
        /// </summary>
        [Input("environmentId")]
        public Input<string>? EnvironmentId { get; set; }

        /// <summary>
        /// The expiration date for VM.
        /// </summary>
        [Input("expirationDate")]
        public Input<string>? ExpirationDate { get; set; }

        /// <summary>
        /// The Microsoft Azure Marketplace image reference of the virtual machine.
        /// </summary>
        [Input("galleryImageReference")]
        public Input<Inputs.GalleryImageReferenceArgs>? GalleryImageReference { get; set; }

        /// <summary>
        /// Indicates whether this virtual machine uses an SSH key for authentication.
        /// </summary>
        [Input("isAuthenticationWithSshKey")]
        public Input<bool>? IsAuthenticationWithSshKey { get; set; }

        /// <summary>
        /// The lab subnet name of the virtual machine.
        /// </summary>
        [Input("labSubnetName")]
        public Input<string>? LabSubnetName { get; set; }

        /// <summary>
        /// The lab virtual network identifier of the virtual machine.
        /// </summary>
        [Input("labVirtualNetworkId")]
        public Input<string>? LabVirtualNetworkId { get; set; }

        /// <summary>
        /// The location of the new virtual machine or environment
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the virtual machine or environment
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The network interface properties.
        /// </summary>
        [Input("networkInterface")]
        public Input<Inputs.NetworkInterfacePropertiesArgs>? NetworkInterface { get; set; }

        /// <summary>
        /// The notes of the virtual machine.
        /// </summary>
        [Input("notes")]
        public Input<string>? Notes { get; set; }

        /// <summary>
        /// The object identifier of the owner of the virtual machine.
        /// </summary>
        [Input("ownerObjectId")]
        public Input<string>? OwnerObjectId { get; set; }

        /// <summary>
        /// The user principal name of the virtual machine owner.
        /// </summary>
        [Input("ownerUserPrincipalName")]
        public Input<string>? OwnerUserPrincipalName { get; set; }

        /// <summary>
        /// The password of the virtual machine administrator.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The id of the plan associated with the virtual machine image
        /// </summary>
        [Input("planId")]
        public Input<string>? PlanId { get; set; }

        [Input("scheduleParameters")]
        private InputList<Inputs.ScheduleCreationParameterArgs>? _scheduleParameters;

        /// <summary>
        /// Virtual Machine schedules to be created
        /// </summary>
        public InputList<Inputs.ScheduleCreationParameterArgs> ScheduleParameters
        {
            get => _scheduleParameters ?? (_scheduleParameters = new InputList<Inputs.ScheduleCreationParameterArgs>());
            set => _scheduleParameters = value;
        }

        /// <summary>
        /// The size of the virtual machine.
        /// </summary>
        [Input("size")]
        public Input<string>? Size { get; set; }

        /// <summary>
        /// The SSH key of the virtual machine administrator.
        /// </summary>
        [Input("sshKey")]
        public Input<string>? SshKey { get; set; }

        /// <summary>
        /// Storage type to use for virtual machine (i.e. Standard, Premium).
        /// </summary>
        [Input("storageType")]
        public Input<string>? StorageType { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The tags of the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The user name of the virtual machine.
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        public LabVirtualMachineCreationParameterArgs()
        {
            AllowClaim = false;
            DisallowPublicIpAddress = false;
            OwnerObjectId = "dynamicValue";
            StorageType = "labStorageType";
        }
        public static new LabVirtualMachineCreationParameterArgs Empty => new LabVirtualMachineCreationParameterArgs();
    }
}
