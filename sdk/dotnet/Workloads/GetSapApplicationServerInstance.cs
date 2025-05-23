// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Workloads
{
    public static class GetSapApplicationServerInstance
    {
        /// <summary>
        /// Gets the SAP Application Server Instance corresponding to the Virtual Instance for SAP solutions resource.
        /// 
        /// Uses Azure REST API version 2024-09-01.
        /// </summary>
        public static Task<GetSapApplicationServerInstanceResult> InvokeAsync(GetSapApplicationServerInstanceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSapApplicationServerInstanceResult>("azure-native:workloads:getSapApplicationServerInstance", args ?? new GetSapApplicationServerInstanceArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the SAP Application Server Instance corresponding to the Virtual Instance for SAP solutions resource.
        /// 
        /// Uses Azure REST API version 2024-09-01.
        /// </summary>
        public static Output<GetSapApplicationServerInstanceResult> Invoke(GetSapApplicationServerInstanceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSapApplicationServerInstanceResult>("azure-native:workloads:getSapApplicationServerInstance", args ?? new GetSapApplicationServerInstanceInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the SAP Application Server Instance corresponding to the Virtual Instance for SAP solutions resource.
        /// 
        /// Uses Azure REST API version 2024-09-01.
        /// </summary>
        public static Output<GetSapApplicationServerInstanceResult> Invoke(GetSapApplicationServerInstanceInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSapApplicationServerInstanceResult>("azure-native:workloads:getSapApplicationServerInstance", args ?? new GetSapApplicationServerInstanceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSapApplicationServerInstanceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of SAP Application Server instance resource.
        /// </summary>
        [Input("applicationInstanceName", required: true)]
        public string ApplicationInstanceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the Virtual Instances for SAP solutions resource
        /// </summary>
        [Input("sapVirtualInstanceName", required: true)]
        public string SapVirtualInstanceName { get; set; } = null!;

        public GetSapApplicationServerInstanceArgs()
        {
        }
        public static new GetSapApplicationServerInstanceArgs Empty => new GetSapApplicationServerInstanceArgs();
    }

    public sealed class GetSapApplicationServerInstanceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of SAP Application Server instance resource.
        /// </summary>
        [Input("applicationInstanceName", required: true)]
        public Input<string> ApplicationInstanceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the Virtual Instances for SAP solutions resource
        /// </summary>
        [Input("sapVirtualInstanceName", required: true)]
        public Input<string> SapVirtualInstanceName { get; set; } = null!;

        public GetSapApplicationServerInstanceInvokeArgs()
        {
        }
        public static new GetSapApplicationServerInstanceInvokeArgs Empty => new GetSapApplicationServerInstanceInvokeArgs();
    }


    [OutputType]
    public sealed class GetSapApplicationServerInstanceResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Application server instance dispatcher status.
        /// </summary>
        public readonly string DispatcherStatus;
        /// <summary>
        /// Defines the Application Instance errors.
        /// </summary>
        public readonly Outputs.SAPVirtualInstanceErrorResponse Errors;
        /// <summary>
        /// Application server instance gateway Port.
        /// </summary>
        public readonly double GatewayPort;
        /// <summary>
        /// Defines the health of SAP Instances.
        /// </summary>
        public readonly string Health;
        /// <summary>
        /// Application server instance SAP hostname.
        /// </summary>
        public readonly string Hostname;
        /// <summary>
        /// Application server instance ICM HTTP Port.
        /// </summary>
        public readonly double IcmHttpPort;
        /// <summary>
        /// Application server instance ICM HTTPS Port.
        /// </summary>
        public readonly double IcmHttpsPort;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Application server Instance Number.
        /// </summary>
        public readonly string InstanceNo;
        /// <summary>
        /// Application server instance SAP IP Address.
        /// </summary>
        public readonly string IpAddress;
        /// <summary>
        /// Application server instance SAP Kernel Patch level.
        /// </summary>
        public readonly string KernelPatch;
        /// <summary>
        /// Application server instance SAP Kernel Version.
        /// </summary>
        public readonly string KernelVersion;
        /// <summary>
        /// The Load Balancer details such as LoadBalancer ID attached to Application Server Virtual Machines
        /// </summary>
        public readonly Outputs.LoadBalancerDetailsResponse LoadBalancerDetails;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Defines the provisioning states.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Defines the SAP Instance status.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Application server Subnet.
        /// </summary>
        public readonly string Subnet;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The list of virtual machines.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationServerVmDetailsResponse> VmDetails;

        [OutputConstructor]
        private GetSapApplicationServerInstanceResult(
            string azureApiVersion,

            string dispatcherStatus,

            Outputs.SAPVirtualInstanceErrorResponse errors,

            double gatewayPort,

            string health,

            string hostname,

            double icmHttpPort,

            double icmHttpsPort,

            string id,

            string instanceNo,

            string ipAddress,

            string kernelPatch,

            string kernelVersion,

            Outputs.LoadBalancerDetailsResponse loadBalancerDetails,

            string location,

            string name,

            string provisioningState,

            string status,

            string subnet,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type,

            ImmutableArray<Outputs.ApplicationServerVmDetailsResponse> vmDetails)
        {
            AzureApiVersion = azureApiVersion;
            DispatcherStatus = dispatcherStatus;
            Errors = errors;
            GatewayPort = gatewayPort;
            Health = health;
            Hostname = hostname;
            IcmHttpPort = icmHttpPort;
            IcmHttpsPort = icmHttpsPort;
            Id = id;
            InstanceNo = instanceNo;
            IpAddress = ipAddress;
            KernelPatch = kernelPatch;
            KernelVersion = kernelVersion;
            LoadBalancerDetails = loadBalancerDetails;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            Status = status;
            Subnet = subnet;
            SystemData = systemData;
            Tags = tags;
            Type = type;
            VmDetails = vmDetails;
        }
    }
}
