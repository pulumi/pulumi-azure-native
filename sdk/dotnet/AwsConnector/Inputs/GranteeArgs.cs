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
    /// Definition of Grantee
    /// </summary>
    public sealed class GranteeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;p&gt;Screen name of the grantee.&lt;/p&gt;
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// &lt;p&gt;Email address of the grantee.&lt;/p&gt; &lt;note&gt; &lt;p&gt;Using email addresses to specify a grantee is only supported in the following Amazon Web Services Regions: &lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt;US East (N. Virginia)&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt;US West (N. California)&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; US West (Oregon)&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; Asia Pacific (Singapore)&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt;Asia Pacific (Sydney)&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt;Asia Pacific (Tokyo)&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt;Europe (Ireland)&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt;South America (São Paulo)&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; &lt;p&gt;For a list of all the Amazon S3 supported Regions and endpoints, see &lt;a href='https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region'&gt;Regions and Endpoints&lt;/a&gt; in the Amazon Web Services General Reference.&lt;/p&gt; &lt;/note&gt;
        /// </summary>
        [Input("emailAddress")]
        public Input<string>? EmailAddress { get; set; }

        /// <summary>
        /// &lt;p&gt;The canonical user ID of the grantee.&lt;/p&gt;
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// &lt;p&gt;Type of grantee&lt;/p&gt;
        /// </summary>
        [Input("type")]
        public Input<Inputs.TypeEnumValueArgs>? Type { get; set; }

        /// <summary>
        /// &lt;p&gt;URI of the grantee group.&lt;/p&gt;
        /// </summary>
        [Input("uri")]
        public Input<string>? Uri { get; set; }

        public GranteeArgs()
        {
        }
        public static new GranteeArgs Empty => new GranteeArgs();
    }
}
