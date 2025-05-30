// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector.Inputs
{

    /// <summary>
    /// Definition of ElasticGpuAssociation
    /// </summary>
    public sealed class ElasticGpuAssociationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;p&gt;The ID of the association.&lt;/p&gt;
        /// </summary>
        [Input("elasticGpuAssociationId")]
        public Input<string>? ElasticGpuAssociationId { get; set; }

        /// <summary>
        /// &lt;p&gt;The state of the association between the instance and the Elastic Graphics accelerator.&lt;/p&gt;
        /// </summary>
        [Input("elasticGpuAssociationState")]
        public Input<string>? ElasticGpuAssociationState { get; set; }

        /// <summary>
        /// &lt;p&gt;The time the Elastic Graphics accelerator was associated with the instance.&lt;/p&gt;
        /// </summary>
        [Input("elasticGpuAssociationTime")]
        public Input<string>? ElasticGpuAssociationTime { get; set; }

        /// <summary>
        /// &lt;p&gt;The ID of the Elastic Graphics accelerator.&lt;/p&gt;
        /// </summary>
        [Input("elasticGpuId")]
        public Input<string>? ElasticGpuId { get; set; }

        public ElasticGpuAssociationArgs()
        {
        }
        public static new ElasticGpuAssociationArgs Empty => new ElasticGpuAssociationArgs();
    }
}
