// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Inputs
{

    /// <summary>
    /// Resource requirements for each container instance within an online deployment.
    /// </summary>
    public sealed class ContainerResourceRequirementsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Container resource limit info:
        /// </summary>
        [Input("containerResourceLimits")]
        public Input<Inputs.ContainerResourceSettingsArgs>? ContainerResourceLimits { get; set; }

        /// <summary>
        /// Container resource request info:
        /// </summary>
        [Input("containerResourceRequests")]
        public Input<Inputs.ContainerResourceSettingsArgs>? ContainerResourceRequests { get; set; }

        public ContainerResourceRequirementsArgs()
        {
        }
        public static new ContainerResourceRequirementsArgs Empty => new ContainerResourceRequirementsArgs();
    }
}
