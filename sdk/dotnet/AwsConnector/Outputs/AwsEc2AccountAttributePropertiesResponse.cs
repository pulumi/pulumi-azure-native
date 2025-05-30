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
    /// Definition of awsEc2AccountAttribute
    /// </summary>
    [OutputType]
    public sealed class AwsEc2AccountAttributePropertiesResponse
    {
        /// <summary>
        /// &lt;p&gt;The name of the account attribute.&lt;/p&gt;
        /// </summary>
        public readonly string? AttributeName;
        /// <summary>
        /// &lt;p&gt;The values for the account attribute.&lt;/p&gt;
        /// </summary>
        public readonly ImmutableArray<Outputs.AccountAttributeValueResponse> AttributeValues;

        [OutputConstructor]
        private AwsEc2AccountAttributePropertiesResponse(
            string? attributeName,

            ImmutableArray<Outputs.AccountAttributeValueResponse> attributeValues)
        {
            AttributeName = attributeName;
            AttributeValues = attributeValues;
        }
    }
}
