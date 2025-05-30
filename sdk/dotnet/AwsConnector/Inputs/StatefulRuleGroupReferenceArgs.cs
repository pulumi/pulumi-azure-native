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
    /// Definition of StatefulRuleGroupReference
    /// </summary>
    public sealed class StatefulRuleGroupReferenceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Property override
        /// </summary>
        [Input("override")]
        public Input<Inputs.StatefulRuleGroupOverrideArgs>? Override { get; set; }

        /// <summary>
        /// Property priority
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// A resource ARN.
        /// </summary>
        [Input("resourceArn")]
        public Input<string>? ResourceArn { get; set; }

        public StatefulRuleGroupReferenceArgs()
        {
        }
        public static new StatefulRuleGroupReferenceArgs Empty => new StatefulRuleGroupReferenceArgs();
    }
}
