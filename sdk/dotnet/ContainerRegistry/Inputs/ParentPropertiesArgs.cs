// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerRegistry.Inputs
{

    /// <summary>
    /// The properties of the connected registry parent.
    /// </summary>
    public sealed class ParentPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource ID of the parent to which the connected registry will be associated.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The sync properties of the connected registry with its parent.
        /// </summary>
        [Input("syncProperties", required: true)]
        public Input<Inputs.SyncPropertiesArgs> SyncProperties { get; set; } = null!;

        public ParentPropertiesArgs()
        {
        }
        public static new ParentPropertiesArgs Empty => new ParentPropertiesArgs();
    }
}
