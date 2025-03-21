// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web.Inputs
{

    /// <summary>
    /// Container App versioned application definition.
    /// Defines the desired state of an immutable revision.
    /// Any changes to this section Will result in a new revision being created
    /// </summary>
    public sealed class TemplateArgs : global::Pulumi.ResourceArgs
    {
        [Input("containers")]
        private InputList<Inputs.ContainerArgs>? _containers;

        /// <summary>
        /// List of container definitions for the Container App.
        /// </summary>
        public InputList<Inputs.ContainerArgs> Containers
        {
            get => _containers ?? (_containers = new InputList<Inputs.ContainerArgs>());
            set => _containers = value;
        }

        /// <summary>
        /// Dapr configuration for the Container App.
        /// </summary>
        [Input("dapr")]
        public Input<Inputs.DaprArgs>? Dapr { get; set; }

        /// <summary>
        /// User friendly suffix that is appended to the revision name
        /// </summary>
        [Input("revisionSuffix")]
        public Input<string>? RevisionSuffix { get; set; }

        /// <summary>
        /// Scaling properties for the Container App.
        /// </summary>
        [Input("scale")]
        public Input<Inputs.ScaleArgs>? Scale { get; set; }

        public TemplateArgs()
        {
        }
        public static new TemplateArgs Empty => new TemplateArgs();
    }
}
