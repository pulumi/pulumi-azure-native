// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Solutions.Inputs
{

    /// <summary>
    /// Managed application deployment policy.
    /// </summary>
    public sealed class ApplicationDeploymentPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The managed application deployment mode.
        /// </summary>
        [Input("deploymentMode", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Solutions.DeploymentMode> DeploymentMode { get; set; } = null!;

        public ApplicationDeploymentPolicyArgs()
        {
        }
        public static new ApplicationDeploymentPolicyArgs Empty => new ApplicationDeploymentPolicyArgs();
    }
}
