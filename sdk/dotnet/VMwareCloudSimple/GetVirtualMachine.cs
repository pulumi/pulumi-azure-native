// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.VMwareCloudSimple
{
    public static class GetVirtualMachine
    {
        /// <summary>
        /// Get virtual machine
        /// 
        /// Uses Azure REST API version 2019-04-01.
        /// </summary>
        public static Task<GetVirtualMachineResult> InvokeAsync(GetVirtualMachineArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVirtualMachineResult>("azure-native:vmwarecloudsimple:getVirtualMachine", args ?? new GetVirtualMachineArgs(), options.WithDefaults());

        /// <summary>
        /// Get virtual machine
        /// 
        /// Uses Azure REST API version 2019-04-01.
        /// </summary>
        public static Output<GetVirtualMachineResult> Invoke(GetVirtualMachineInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVirtualMachineResult>("azure-native:vmwarecloudsimple:getVirtualMachine", args ?? new GetVirtualMachineInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get virtual machine
        /// 
        /// Uses Azure REST API version 2019-04-01.
        /// </summary>
        public static Output<GetVirtualMachineResult> Invoke(GetVirtualMachineInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetVirtualMachineResult>("azure-native:vmwarecloudsimple:getVirtualMachine", args ?? new GetVirtualMachineInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVirtualMachineArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// virtual machine name
        /// </summary>
        [Input("virtualMachineName", required: true)]
        public string VirtualMachineName { get; set; } = null!;

        public GetVirtualMachineArgs()
        {
        }
        public static new GetVirtualMachineArgs Empty => new GetVirtualMachineArgs();
    }

    public sealed class GetVirtualMachineInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// virtual machine name
        /// </summary>
        [Input("virtualMachineName", required: true)]
        public Input<string> VirtualMachineName { get; set; } = null!;

        public GetVirtualMachineInvokeArgs()
        {
        }
        public static new GetVirtualMachineInvokeArgs Empty => new GetVirtualMachineInvokeArgs();
    }


    [OutputType]
    public sealed class GetVirtualMachineResult
    {
        /// <summary>
        /// The amount of memory
        /// </summary>
        public readonly int AmountOfRam;
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// The list of Virtual Disks' Controllers
        /// </summary>
        public readonly ImmutableArray<Outputs.VirtualDiskControllerResponse> Controllers;
        /// <summary>
        /// Virtual machine properties
        /// </summary>
        public readonly Outputs.GuestOSCustomizationResponse? Customization;
        /// <summary>
        /// The list of Virtual Disks
        /// </summary>
        public readonly ImmutableArray<Outputs.VirtualDiskResponse> Disks;
        /// <summary>
        /// The DNS name of Virtual Machine in VCenter
        /// </summary>
        public readonly string Dnsname;
        /// <summary>
        /// Expose Guest OS or not
        /// </summary>
        public readonly bool? ExposeToGuestVM;
        /// <summary>
        /// The path to virtual machine folder in VCenter
        /// </summary>
        public readonly string Folder;
        /// <summary>
        /// The name of Guest OS
        /// </summary>
        public readonly string GuestOS;
        /// <summary>
        /// The Guest OS type
        /// </summary>
        public readonly string GuestOSType;
        /// <summary>
        /// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/virtualMachines/{virtualMachineName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Azure region
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// {virtualMachineName}
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The list of Virtual NICs
        /// </summary>
        public readonly ImmutableArray<Outputs.VirtualNicResponse> Nics;
        /// <summary>
        /// The number of CPU cores
        /// </summary>
        public readonly int NumberOfCores;
        /// <summary>
        /// Password for login. Deprecated - use customization property
        /// </summary>
        public readonly string? Password;
        /// <summary>
        /// Private Cloud Id
        /// </summary>
        public readonly string PrivateCloudId;
        /// <summary>
        /// The provisioning status of the resource
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The public ip of Virtual Machine
        /// </summary>
        public readonly string PublicIP;
        /// <summary>
        /// Virtual Machines Resource Pool
        /// </summary>
        public readonly Outputs.ResourcePoolResponse? ResourcePool;
        /// <summary>
        /// The status of Virtual machine
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The list of tags
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Virtual Machine Template Id
        /// </summary>
        public readonly string? TemplateId;
        /// <summary>
        /// {resourceProviderNamespace}/{resourceType}
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Username for login. Deprecated - use customization property
        /// </summary>
        public readonly string? Username;
        /// <summary>
        /// The list of Virtual VSphere Networks
        /// </summary>
        public readonly ImmutableArray<string> VSphereNetworks;
        /// <summary>
        /// The internal id of Virtual Machine in VCenter
        /// </summary>
        public readonly string VmId;
        /// <summary>
        /// VMware tools version
        /// </summary>
        public readonly string Vmwaretools;

        [OutputConstructor]
        private GetVirtualMachineResult(
            int amountOfRam,

            string azureApiVersion,

            ImmutableArray<Outputs.VirtualDiskControllerResponse> controllers,

            Outputs.GuestOSCustomizationResponse? customization,

            ImmutableArray<Outputs.VirtualDiskResponse> disks,

            string dnsname,

            bool? exposeToGuestVM,

            string folder,

            string guestOS,

            string guestOSType,

            string id,

            string location,

            string name,

            ImmutableArray<Outputs.VirtualNicResponse> nics,

            int numberOfCores,

            string? password,

            string privateCloudId,

            string provisioningState,

            string publicIP,

            Outputs.ResourcePoolResponse? resourcePool,

            string status,

            ImmutableDictionary<string, string>? tags,

            string? templateId,

            string type,

            string? username,

            ImmutableArray<string> vSphereNetworks,

            string vmId,

            string vmwaretools)
        {
            AmountOfRam = amountOfRam;
            AzureApiVersion = azureApiVersion;
            Controllers = controllers;
            Customization = customization;
            Disks = disks;
            Dnsname = dnsname;
            ExposeToGuestVM = exposeToGuestVM;
            Folder = folder;
            GuestOS = guestOS;
            GuestOSType = guestOSType;
            Id = id;
            Location = location;
            Name = name;
            Nics = nics;
            NumberOfCores = numberOfCores;
            Password = password;
            PrivateCloudId = privateCloudId;
            ProvisioningState = provisioningState;
            PublicIP = publicIP;
            ResourcePool = resourcePool;
            Status = status;
            Tags = tags;
            TemplateId = templateId;
            Type = type;
            Username = username;
            VSphereNetworks = vSphereNetworks;
            VmId = vmId;
            Vmwaretools = vmwaretools;
        }
    }
}
