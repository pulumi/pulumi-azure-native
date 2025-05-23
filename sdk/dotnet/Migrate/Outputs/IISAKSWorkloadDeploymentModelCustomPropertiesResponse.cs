// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Migrate.Outputs
{

    /// <summary>
    /// IIS workload instance model custom properties.
    /// </summary>
    [OutputType]
    public sealed class IISAKSWorkloadDeploymentModelCustomPropertiesResponse
    {
        /// <summary>
        /// IIS AKS workload deployment.
        /// </summary>
        public readonly Outputs.IISAKSWorkloadDeploymentResponse? IisAksWorkloadDeploymentProperties;
        /// <summary>
        /// Gets or sets the instance type.
        /// Expected value is 'IISAKSWorkloadDeploymentModelCustomProperties'.
        /// </summary>
        public readonly string InstanceType;

        [OutputConstructor]
        private IISAKSWorkloadDeploymentModelCustomPropertiesResponse(
            Outputs.IISAKSWorkloadDeploymentResponse? iisAksWorkloadDeploymentProperties,

            string instanceType)
        {
            IisAksWorkloadDeploymentProperties = iisAksWorkloadDeploymentProperties;
            InstanceType = instanceType;
        }
    }
}
