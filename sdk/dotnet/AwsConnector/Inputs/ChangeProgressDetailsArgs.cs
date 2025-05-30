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
    /// Definition of ChangeProgressDetails
    /// </summary>
    public sealed class ChangeProgressDetailsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;p&gt;The ID of the configuration change.&lt;/p&gt;
        /// </summary>
        [Input("changeId")]
        public Input<string>? ChangeId { get; set; }

        /// <summary>
        /// &lt;p&gt;The current status of the configuration change.&lt;/p&gt;
        /// </summary>
        [Input("configChangeStatus")]
        public Input<Inputs.ConfigChangeStatusEnumValueArgs>? ConfigChangeStatus { get; set; }

        /// <summary>
        /// &lt;p&gt;The IAM principal who initiated the configuration change.&lt;/p&gt;
        /// </summary>
        [Input("initiatedBy")]
        public Input<Inputs.InitiatedByEnumValueArgs>? InitiatedBy { get; set; }

        /// <summary>
        /// &lt;p&gt;The last time that the configuration change was updated.&lt;/p&gt;
        /// </summary>
        [Input("lastUpdatedTime")]
        public Input<string>? LastUpdatedTime { get; set; }

        /// <summary>
        /// &lt;p&gt;A message corresponding to the status of the configuration change.&lt;/p&gt;
        /// </summary>
        [Input("message")]
        public Input<string>? Message { get; set; }

        /// <summary>
        /// &lt;p&gt;The time that the configuration change was initiated, in Universal Coordinated Time (UTC).&lt;/p&gt;
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        public ChangeProgressDetailsArgs()
        {
        }
        public static new ChangeProgressDetailsArgs Empty => new ChangeProgressDetailsArgs();
    }
}
