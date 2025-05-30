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
    public sealed class ManagedOnlineEndpointDeploymentResourcePropertiesResponse
    {
        public readonly string? EndpointComputeType;
        /// <summary>
        /// The failure reason if the creation failed.
        /// </summary>
        public readonly string? FailureReason;
        public readonly string? Model;
        /// <summary>
        /// Read-only provision state status property.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Kind of the deployment.
        /// Expected value is 'managedOnlineEndpoint'.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ManagedOnlineEndpointDeploymentResourcePropertiesResponse(
            string? endpointComputeType,

            string? failureReason,

            string? model,

            string provisioningState,

            string type)
        {
            EndpointComputeType = endpointComputeType;
            FailureReason = failureReason;
            Model = model;
            ProvisioningState = provisioningState;
            Type = type;
        }
    }
}
