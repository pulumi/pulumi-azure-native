// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Databricks.Outputs
{

    /// <summary>
    /// Custom Parameters used for Cluster Creation.
    /// </summary>
    [OutputType]
    public sealed class WorkspaceCustomParametersResponse
    {
        /// <summary>
        /// The ID of a Azure Machine Learning workspace to link with Databricks workspace
        /// </summary>
        public readonly Outputs.WorkspaceCustomStringParameterResponse? AmlWorkspaceId;
        /// <summary>
        /// The name of the Private Subnet within the Virtual Network
        /// </summary>
        public readonly Outputs.WorkspaceCustomStringParameterResponse? CustomPrivateSubnetName;
        /// <summary>
        /// The name of a Public Subnet within the Virtual Network
        /// </summary>
        public readonly Outputs.WorkspaceCustomStringParameterResponse? CustomPublicSubnetName;
        /// <summary>
        /// The ID of a Virtual Network where this Databricks Cluster should be created
        /// </summary>
        public readonly Outputs.WorkspaceCustomStringParameterResponse? CustomVirtualNetworkId;
        /// <summary>
        /// Boolean indicating whether the public IP should be disabled. Default value is true
        /// </summary>
        public readonly Outputs.WorkspaceNoPublicIPBooleanParameterResponse? EnableNoPublicIp;
        /// <summary>
        /// Contains the encryption details for Customer-Managed Key (CMK) enabled workspace.
        /// </summary>
        public readonly Outputs.WorkspaceEncryptionParameterResponse? Encryption;
        /// <summary>
        /// Name of the outbound Load Balancer Backend Pool for Secure Cluster Connectivity (No Public IP).
        /// </summary>
        public readonly Outputs.WorkspaceCustomStringParameterResponse? LoadBalancerBackendPoolName;
        /// <summary>
        /// Resource URI of Outbound Load balancer for Secure Cluster Connectivity (No Public IP) workspace.
        /// </summary>
        public readonly Outputs.WorkspaceCustomStringParameterResponse? LoadBalancerId;
        /// <summary>
        /// Name of the NAT gateway for Secure Cluster Connectivity (No Public IP) workspace subnets.
        /// </summary>
        public readonly Outputs.WorkspaceCustomStringParameterResponse? NatGatewayName;
        /// <summary>
        /// Prepare the workspace for encryption. Enables the Managed Identity for managed storage account.
        /// </summary>
        public readonly Outputs.WorkspaceCustomBooleanParameterResponse? PrepareEncryption;
        /// <summary>
        /// Name of the Public IP for No Public IP workspace with managed vNet.
        /// </summary>
        public readonly Outputs.WorkspaceCustomStringParameterResponse? PublicIpName;
        /// <summary>
        /// A boolean indicating whether or not the DBFS root file system will be enabled with secondary layer of encryption with platform managed keys for data at rest.
        /// </summary>
        public readonly Outputs.WorkspaceCustomBooleanParameterResponse? RequireInfrastructureEncryption;
        /// <summary>
        /// Tags applied to resources under Managed resource group. These can be updated by updating tags at workspace level.
        /// </summary>
        public readonly Outputs.WorkspaceCustomObjectParameterResponse ResourceTags;
        /// <summary>
        /// Default DBFS storage account name.
        /// </summary>
        public readonly Outputs.WorkspaceCustomStringParameterResponse? StorageAccountName;
        /// <summary>
        /// Storage account SKU name, ex: Standard_GRS, Standard_LRS. Refer https://aka.ms/storageskus for valid inputs.
        /// </summary>
        public readonly Outputs.WorkspaceCustomStringParameterResponse? StorageAccountSkuName;
        /// <summary>
        /// Address prefix for Managed virtual network. Default value for this input is 10.139.
        /// </summary>
        public readonly Outputs.WorkspaceCustomStringParameterResponse? VnetAddressPrefix;

        [OutputConstructor]
        private WorkspaceCustomParametersResponse(
            Outputs.WorkspaceCustomStringParameterResponse? amlWorkspaceId,

            Outputs.WorkspaceCustomStringParameterResponse? customPrivateSubnetName,

            Outputs.WorkspaceCustomStringParameterResponse? customPublicSubnetName,

            Outputs.WorkspaceCustomStringParameterResponse? customVirtualNetworkId,

            Outputs.WorkspaceNoPublicIPBooleanParameterResponse? enableNoPublicIp,

            Outputs.WorkspaceEncryptionParameterResponse? encryption,

            Outputs.WorkspaceCustomStringParameterResponse? loadBalancerBackendPoolName,

            Outputs.WorkspaceCustomStringParameterResponse? loadBalancerId,

            Outputs.WorkspaceCustomStringParameterResponse? natGatewayName,

            Outputs.WorkspaceCustomBooleanParameterResponse? prepareEncryption,

            Outputs.WorkspaceCustomStringParameterResponse? publicIpName,

            Outputs.WorkspaceCustomBooleanParameterResponse? requireInfrastructureEncryption,

            Outputs.WorkspaceCustomObjectParameterResponse resourceTags,

            Outputs.WorkspaceCustomStringParameterResponse? storageAccountName,

            Outputs.WorkspaceCustomStringParameterResponse? storageAccountSkuName,

            Outputs.WorkspaceCustomStringParameterResponse? vnetAddressPrefix)
        {
            AmlWorkspaceId = amlWorkspaceId;
            CustomPrivateSubnetName = customPrivateSubnetName;
            CustomPublicSubnetName = customPublicSubnetName;
            CustomVirtualNetworkId = customVirtualNetworkId;
            EnableNoPublicIp = enableNoPublicIp;
            Encryption = encryption;
            LoadBalancerBackendPoolName = loadBalancerBackendPoolName;
            LoadBalancerId = loadBalancerId;
            NatGatewayName = natGatewayName;
            PrepareEncryption = prepareEncryption;
            PublicIpName = publicIpName;
            RequireInfrastructureEncryption = requireInfrastructureEncryption;
            ResourceTags = resourceTags;
            StorageAccountName = storageAccountName;
            StorageAccountSkuName = storageAccountSkuName;
            VnetAddressPrefix = vnetAddressPrefix;
        }
    }
}
