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
    /// Definition of InstanceMetadataOptionsResponse
    /// </summary>
    [OutputType]
    public sealed class InstanceMetadataOptionsResponseResponse
    {
        /// <summary>
        /// &lt;p&gt;Indicates whether the HTTP metadata endpoint on your instances is enabled or disabled.&lt;/p&gt; &lt;p&gt;If the value is &lt;code&gt;disabled&lt;/code&gt;, you cannot access your instance metadata.&lt;/p&gt;
        /// </summary>
        public readonly Outputs.InstanceMetadataEndpointStateEnumValueResponse? HttpEndpoint;
        /// <summary>
        /// &lt;p&gt;Indicates whether the IPv6 endpoint for the instance metadata service is enabled or disabled.&lt;/p&gt; &lt;p&gt;Default: &lt;code&gt;disabled&lt;/code&gt; &lt;/p&gt;
        /// </summary>
        public readonly Outputs.InstanceMetadataProtocolStateEnumValueResponse? HttpProtocolIpv6;
        /// <summary>
        /// &lt;p&gt;The maximum number of hops that the metadata token can travel.&lt;/p&gt; &lt;p&gt;Possible values: Integers from &lt;code&gt;1&lt;/code&gt; to &lt;code&gt;64&lt;/code&gt; &lt;/p&gt;
        /// </summary>
        public readonly int? HttpPutResponseHopLimit;
        /// <summary>
        /// &lt;p&gt;Indicates whether IMDSv2 is required.&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;optional&lt;/code&gt; - IMDSv2 is optional, which means that you can use either IMDSv2 or IMDSv1.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;required&lt;/code&gt; - IMDSv2 is required, which means that IMDSv1 is disabled, and you must use IMDSv2.&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt;
        /// </summary>
        public readonly Outputs.HttpTokensStateEnumValueResponse? HttpTokens;
        /// <summary>
        /// &lt;p&gt;Indicates whether access to instance tags from the instance metadata is enabled or disabled. For more information, see &lt;a href='https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html#work-with-tags-in-IMDS'&gt;Work with instance tags using the instance metadata&lt;/a&gt;.&lt;/p&gt;
        /// </summary>
        public readonly Outputs.InstanceMetadataTagsStateEnumValueResponse? InstanceMetadataTags;
        /// <summary>
        /// &lt;p&gt;The state of the metadata option changes.&lt;/p&gt; &lt;p&gt; &lt;code&gt;pending&lt;/code&gt; - The metadata options are being updated and the instance is not ready to process metadata traffic with the new selection.&lt;/p&gt; &lt;p&gt; &lt;code&gt;applied&lt;/code&gt; - The metadata options have been successfully applied on the instance.&lt;/p&gt;
        /// </summary>
        public readonly Outputs.InstanceMetadataOptionsStateEnumValueResponse? State;

        [OutputConstructor]
        private InstanceMetadataOptionsResponseResponse(
            Outputs.InstanceMetadataEndpointStateEnumValueResponse? httpEndpoint,

            Outputs.InstanceMetadataProtocolStateEnumValueResponse? httpProtocolIpv6,

            int? httpPutResponseHopLimit,

            Outputs.HttpTokensStateEnumValueResponse? httpTokens,

            Outputs.InstanceMetadataTagsStateEnumValueResponse? instanceMetadataTags,

            Outputs.InstanceMetadataOptionsStateEnumValueResponse? state)
        {
            HttpEndpoint = httpEndpoint;
            HttpProtocolIpv6 = httpProtocolIpv6;
            HttpPutResponseHopLimit = httpPutResponseHopLimit;
            HttpTokens = httpTokens;
            InstanceMetadataTags = instanceMetadataTags;
            State = state;
        }
    }
}
