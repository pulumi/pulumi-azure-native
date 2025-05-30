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
    /// Definition of ProjectBuildBatchConfig
    /// </summary>
    public sealed class ProjectBuildBatchConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;p&gt;Specifies how build status reports are sent to the source provider for the batch build. This property is only used when the source provider for your project is Bitbucket, GitHub, or GitHub Enterprise, and your project is configured to report build statuses to the source provider.&lt;/p&gt; &lt;dl&gt; &lt;dt&gt;REPORT_AGGREGATED_BATCH&lt;/dt&gt; &lt;dd&gt; &lt;p&gt;(Default) Aggregate all of the build statuses into a single status report.&lt;/p&gt; &lt;/dd&gt; &lt;dt&gt;REPORT_INDIVIDUAL_BUILDS&lt;/dt&gt; &lt;dd&gt; &lt;p&gt;Send a separate status report for each individual build.&lt;/p&gt; &lt;/dd&gt; &lt;/dl&gt;
        /// </summary>
        [Input("batchReportMode")]
        public Input<Inputs.BatchReportModeTypeEnumValueArgs>? BatchReportMode { get; set; }

        /// <summary>
        /// &lt;p&gt;Specifies if the build artifacts for the batch build should be combined into a single artifact location.&lt;/p&gt;
        /// </summary>
        [Input("combineArtifacts")]
        public Input<bool>? CombineArtifacts { get; set; }

        /// <summary>
        /// &lt;p&gt;A &lt;code&gt;BatchRestrictions&lt;/code&gt; object that specifies the restrictions for the batch build.&lt;/p&gt;
        /// </summary>
        [Input("restrictions")]
        public Input<Inputs.BatchRestrictionsArgs>? Restrictions { get; set; }

        /// <summary>
        /// &lt;p&gt;Specifies the service role ARN for the batch build project.&lt;/p&gt;
        /// </summary>
        [Input("serviceRole")]
        public Input<string>? ServiceRole { get; set; }

        /// <summary>
        /// &lt;p&gt;Specifies the maximum amount of time, in minutes, that the batch build must be completed in.&lt;/p&gt;
        /// </summary>
        [Input("timeoutInMins")]
        public Input<int>? TimeoutInMins { get; set; }

        public ProjectBuildBatchConfigArgs()
        {
        }
        public static new ProjectBuildBatchConfigArgs Empty => new ProjectBuildBatchConfigArgs();
    }
}
