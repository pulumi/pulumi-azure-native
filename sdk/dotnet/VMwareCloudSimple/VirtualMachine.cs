// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.VMwareCloudSimple
{
    /// <summary>
    /// Virtual machine model
    /// 
    /// Uses Azure REST API version 2019-04-01. In version 2.x of the Azure Native provider, it used API version 2019-04-01.
    /// </summary>
    [AzureNativeResourceType("azure-native:vmwarecloudsimple:VirtualMachine")]
    public partial class VirtualMachine : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The amount of memory
        /// </summary>
        [Output("amountOfRam")]
        public Output<int> AmountOfRam { get; private set; } = null!;

        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// The list of Virtual Disks' Controllers
        /// </summary>
        [Output("controllers")]
        public Output<ImmutableArray<Outputs.VirtualDiskControllerResponse>> Controllers { get; private set; } = null!;

        /// <summary>
        /// Virtual machine properties
        /// </summary>
        [Output("customization")]
        public Output<Outputs.GuestOSCustomizationResponse?> Customization { get; private set; } = null!;

        /// <summary>
        /// The list of Virtual Disks
        /// </summary>
        [Output("disks")]
        public Output<ImmutableArray<Outputs.VirtualDiskResponse>> Disks { get; private set; } = null!;

        /// <summary>
        /// The DNS name of Virtual Machine in VCenter
        /// </summary>
        [Output("dnsname")]
        public Output<string> Dnsname { get; private set; } = null!;

        /// <summary>
        /// Expose Guest OS or not
        /// </summary>
        [Output("exposeToGuestVM")]
        public Output<bool?> ExposeToGuestVM { get; private set; } = null!;

        /// <summary>
        /// The path to virtual machine folder in VCenter
        /// </summary>
        [Output("folder")]
        public Output<string> Folder { get; private set; } = null!;

        /// <summary>
        /// The name of Guest OS
        /// </summary>
        [Output("guestOS")]
        public Output<string> GuestOS { get; private set; } = null!;

        /// <summary>
        /// The Guest OS type
        /// </summary>
        [Output("guestOSType")]
        public Output<string> GuestOSType { get; private set; } = null!;

        /// <summary>
        /// Azure region
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// {virtualMachineName}
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The list of Virtual NICs
        /// </summary>
        [Output("nics")]
        public Output<ImmutableArray<Outputs.VirtualNicResponse>> Nics { get; private set; } = null!;

        /// <summary>
        /// The number of CPU cores
        /// </summary>
        [Output("numberOfCores")]
        public Output<int> NumberOfCores { get; private set; } = null!;

        /// <summary>
        /// Password for login. Deprecated - use customization property
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// Private Cloud Id
        /// </summary>
        [Output("privateCloudId")]
        public Output<string> PrivateCloudId { get; private set; } = null!;

        /// <summary>
        /// The provisioning status of the resource
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The public ip of Virtual Machine
        /// </summary>
        [Output("publicIP")]
        public Output<string> PublicIP { get; private set; } = null!;

        /// <summary>
        /// Virtual Machines Resource Pool
        /// </summary>
        [Output("resourcePool")]
        public Output<Outputs.ResourcePoolResponse?> ResourcePool { get; private set; } = null!;

        /// <summary>
        /// The status of Virtual machine
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The list of tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Virtual Machine Template Id
        /// </summary>
        [Output("templateId")]
        public Output<string?> TemplateId { get; private set; } = null!;

        /// <summary>
        /// {resourceProviderNamespace}/{resourceType}
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Username for login. Deprecated - use customization property
        /// </summary>
        [Output("username")]
        public Output<string?> Username { get; private set; } = null!;

        /// <summary>
        /// The list of Virtual VSphere Networks
        /// </summary>
        [Output("vSphereNetworks")]
        public Output<ImmutableArray<string>> VSphereNetworks { get; private set; } = null!;

        /// <summary>
        /// The internal id of Virtual Machine in VCenter
        /// </summary>
        [Output("vmId")]
        public Output<string> VmId { get; private set; } = null!;

        /// <summary>
        /// VMware tools version
        /// </summary>
        [Output("vmwaretools")]
        public Output<string> Vmwaretools { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualMachine resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualMachine(string name, VirtualMachineArgs args, CustomResourceOptions? options = null)
            : base("azure-native:vmwarecloudsimple:VirtualMachine", name, args ?? new VirtualMachineArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualMachine(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:vmwarecloudsimple:VirtualMachine", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:vmwarecloudsimple/v20190401:VirtualMachine" },
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
        /// <summary>
        /// The amount of memory
        /// </summary>
        [Input("amountOfRam", required: true)]
        public Input<int> AmountOfRam { get; set; } = null!;

        /// <summary>
        /// Virtual machine properties
        /// </summary>
        [Input("customization")]
        public Input<Inputs.GuestOSCustomizationArgs>? Customization { get; set; }

        [Input("disks")]
        private InputList<Inputs.VirtualDiskArgs>? _disks;

        /// <summary>
        /// The list of Virtual Disks
        /// </summary>
        public InputList<Inputs.VirtualDiskArgs> Disks
        {
            get => _disks ?? (_disks = new InputList<Inputs.VirtualDiskArgs>());
            set => _disks = value;
        }

        /// <summary>
        /// Expose Guest OS or not
        /// </summary>
        [Input("exposeToGuestVM")]
        public Input<bool>? ExposeToGuestVM { get; set; }

        /// <summary>
        /// Azure region
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("nics")]
        private InputList<Inputs.VirtualNicArgs>? _nics;

        /// <summary>
        /// The list of Virtual NICs
        /// </summary>
        public InputList<Inputs.VirtualNicArgs> Nics
        {
            get => _nics ?? (_nics = new InputList<Inputs.VirtualNicArgs>());
            set => _nics = value;
        }

        /// <summary>
        /// The number of CPU cores
        /// </summary>
        [Input("numberOfCores", required: true)]
        public Input<int> NumberOfCores { get; set; } = null!;

        /// <summary>
        /// Password for login. Deprecated - use customization property
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// Private Cloud Id
        /// </summary>
        [Input("privateCloudId", required: true)]
        public Input<string> PrivateCloudId { get; set; } = null!;

        /// <summary>
        /// The name of the resource group
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Virtual Machines Resource Pool
        /// </summary>
        [Input("resourcePool")]
        public Input<Inputs.ResourcePoolArgs>? ResourcePool { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The list of tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Virtual Machine Template Id
        /// </summary>
        [Input("templateId")]
        public Input<string>? TemplateId { get; set; }

        /// <summary>
        /// Username for login. Deprecated - use customization property
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        [Input("vSphereNetworks")]
        private InputList<string>? _vSphereNetworks;

        /// <summary>
        /// The list of Virtual VSphere Networks
        /// </summary>
        public InputList<string> VSphereNetworks
        {
            get => _vSphereNetworks ?? (_vSphereNetworks = new InputList<string>());
            set => _vSphereNetworks = value;
        }

        /// <summary>
        /// virtual machine name
        /// </summary>
        [Input("virtualMachineName")]
        public Input<string>? VirtualMachineName { get; set; }

        public VirtualMachineArgs()
        {
        }
        public static new VirtualMachineArgs Empty => new VirtualMachineArgs();
    }
}
