// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerInstance.Inputs
{

    /// <summary>
    /// Extension sidecars to be added to the deployment.
    /// </summary>
    public sealed class DeploymentExtensionSpecArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of extension to be added.
        /// </summary>
        [Input("extensionType", required: true)]
        public Input<string> ExtensionType { get; set; } = null!;

        /// <summary>
        /// Name of the extension.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Protected settings for the extension.
        /// </summary>
        [Input("protectedSettings")]
        public Input<object>? ProtectedSettings { get; set; }

        /// <summary>
        /// Settings for the extension.
        /// </summary>
        [Input("settings")]
        public Input<object>? Settings { get; set; }

        /// <summary>
        /// Version of the extension being used.
        /// </summary>
        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        public DeploymentExtensionSpecArgs()
        {
        }
        public static new DeploymentExtensionSpecArgs Empty => new DeploymentExtensionSpecArgs();
    }
}
