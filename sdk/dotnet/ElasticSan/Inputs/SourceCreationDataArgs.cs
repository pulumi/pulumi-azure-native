// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ElasticSan.Inputs
{

    /// <summary>
    /// Data source used when creating the volume.
    /// </summary>
    public sealed class SourceCreationDataArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This enumerates the possible sources of a volume creation.
        /// </summary>
        [Input("createSource")]
        public InputUnion<string, Pulumi.AzureNative.ElasticSan.VolumeCreateOption>? CreateSource { get; set; }

        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        [Input("sourceId")]
        public Input<string>? SourceId { get; set; }

        public SourceCreationDataArgs()
        {
        }
        public static new SourceCreationDataArgs Empty => new SourceCreationDataArgs();
    }
}
