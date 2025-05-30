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
    /// Definition of RoutingRule
    /// </summary>
    public sealed class RoutingRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Container for redirect information. You can redirect requests to another host, to another page, or with another protocol. In the event of an error, you can specify a different error code to return. Specifies how requests are redirected. In the event of an error, you can specify a different error code to return.
        /// </summary>
        [Input("redirectRule")]
        public Input<Inputs.RedirectRuleArgs>? RedirectRule { get; set; }

        /// <summary>
        /// A container for describing a condition that must be met for the specified redirect to apply. For example, 1. If request is for pages in the ``/docs`` folder, redirect to the ``/documents`` folder. 2. If request results in HTTP error 4xx, redirect request to another host where you might process the error. A container for describing a condition that must be met for the specified redirect to apply. For example, 1. If request is for pages in the ``/docs`` folder, redirect to the ``/documents`` folder. 2. If request results in HTTP error 4xx, redirect request to another host where you might process the error.
        /// </summary>
        [Input("routingRuleCondition")]
        public Input<Inputs.RoutingRuleConditionArgs>? RoutingRuleCondition { get; set; }

        public RoutingRuleArgs()
        {
        }
        public static new RoutingRuleArgs Empty => new RoutingRuleArgs();
    }
}
