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
    /// Definition of ActionDefinition
    /// </summary>
    public sealed class ActionDefinitionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Property publishMetricAction
        /// </summary>
        [Input("publishMetricAction")]
        public Input<Inputs.PublishMetricActionArgs>? PublishMetricAction { get; set; }

        public ActionDefinitionArgs()
        {
        }
        public static new ActionDefinitionArgs Empty => new ActionDefinitionArgs();
    }
}
