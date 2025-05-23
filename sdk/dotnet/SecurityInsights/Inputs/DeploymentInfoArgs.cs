// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecurityInsights.Inputs
{

    /// <summary>
    /// Information regarding a deployment.
    /// </summary>
    public sealed class DeploymentInfoArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Deployment information.
        /// </summary>
        [Input("deployment")]
        public Input<Inputs.DeploymentArgs>? Deployment { get; set; }

        /// <summary>
        /// Status while fetching the last deployment.
        /// </summary>
        [Input("deploymentFetchStatus")]
        public InputUnion<string, Pulumi.AzureNative.SecurityInsights.DeploymentFetchStatus>? DeploymentFetchStatus { get; set; }

        /// <summary>
        /// Additional details about the deployment that can be shown to the user.
        /// </summary>
        [Input("message")]
        public Input<string>? Message { get; set; }

        public DeploymentInfoArgs()
        {
        }
        public static new DeploymentInfoArgs Empty => new DeploymentInfoArgs();
    }
}
