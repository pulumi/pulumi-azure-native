// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Outputs
{

    [OutputType]
    public sealed class ServerlessComputeSettingsResponse
    {
        /// <summary>
        /// The resource ID of an existing virtual network subnet in which serverless compute nodes should be deployed
        /// </summary>
        public readonly string? ServerlessComputeCustomSubnet;
        /// <summary>
        /// The flag to signal if serverless compute nodes deployed in custom vNet would have no public IP addresses for a workspace with private endpoint
        /// </summary>
        public readonly bool? ServerlessComputeNoPublicIP;

        [OutputConstructor]
        private ServerlessComputeSettingsResponse(
            string? serverlessComputeCustomSubnet,

            bool? serverlessComputeNoPublicIP)
        {
            ServerlessComputeCustomSubnet = serverlessComputeCustomSubnet;
            ServerlessComputeNoPublicIP = serverlessComputeNoPublicIP;
        }
    }
}
