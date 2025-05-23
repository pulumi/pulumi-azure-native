// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector.Outputs
{

    /// <summary>
    /// Definition of RecordingModeOverride
    /// </summary>
    [OutputType]
    public sealed class RecordingModeOverrideResponse
    {
        /// <summary>
        /// &lt;p&gt;A description that you provide for the override.&lt;/p&gt;
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// &lt;p&gt;The recording frequency that will be applied to all the resource types specified in the override.&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt;Continuous recording allows you to record configuration changes continuously whenever a change occurs.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt;Daily recording allows you to receive a configuration item (CI) representing the most recent state of your resources over the last 24-hour period, only if it’s different from the previous CI recorded. &lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; &lt;note&gt; &lt;p&gt;Firewall Manager depends on continuous recording to monitor your resources. If you are using Firewall Manager, it is recommended that you set the recording frequency to Continuous.&lt;/p&gt; &lt;/note&gt;
        /// </summary>
        public readonly Outputs.RecordingFrequencyEnumValueResponse? RecordingFrequency;
        /// <summary>
        /// &lt;p&gt;A comma-separated list that specifies which resource types Config includes in the override.&lt;/p&gt; &lt;important&gt; &lt;p&gt;Daily recording is not supported for the following resource types:&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;AWS::Config::ResourceCompliance&lt;/code&gt; &lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;AWS::Config::ConformancePackCompliance&lt;/code&gt; &lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;AWS::Config::ConfigurationRecorder&lt;/code&gt; &lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; &lt;/important&gt;
        /// </summary>
        public readonly ImmutableArray<string> ResourceTypes;

        [OutputConstructor]
        private RecordingModeOverrideResponse(
            string? description,

            Outputs.RecordingFrequencyEnumValueResponse? recordingFrequency,

            ImmutableArray<string> resourceTypes)
        {
            Description = description;
            RecordingFrequency = recordingFrequency;
            ResourceTypes = resourceTypes;
        }
    }
}
