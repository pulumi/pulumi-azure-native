// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20240401Preview.Inputs
{

    public sealed class EndpointDeploymentResourcePropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Model used for the endpoint deployment.
        /// </summary>
        [Input("model", required: true)]
        public Input<Inputs.EndpointDeploymentModelArgs> Model { get; set; } = null!;

        /// <summary>
        /// The name of RAI policy.
        /// </summary>
        [Input("raiPolicyName")]
        public Input<string>? RaiPolicyName { get; set; }

        /// <summary>
        /// Deployment model version upgrade option.
        /// </summary>
        [Input("versionUpgradeOption")]
        public InputUnion<string, Pulumi.AzureNative.MachineLearningServices.V20240401Preview.DeploymentModelVersionUpgradeOption>? VersionUpgradeOption { get; set; }

        public EndpointDeploymentResourcePropertiesArgs()
        {
        }
        public static new EndpointDeploymentResourcePropertiesArgs Empty => new EndpointDeploymentResourcePropertiesArgs();
    }
}
