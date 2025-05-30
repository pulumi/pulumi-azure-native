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
    /// Definition of LabelNameConditionModelProperties
    /// </summary>
    public sealed class LabelNameConditionModelPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The label name that a log record must contain in order to meet the condition. This must be a fully qualified label name. Fully qualified labels have a prefix, optional namespaces, and label name. The prefix identifies the rule group or web ACL context of the rule that added the label.
        /// </summary>
        [Input("labelName")]
        public Input<string>? LabelName { get; set; }

        public LabelNameConditionModelPropertiesArgs()
        {
        }
        public static new LabelNameConditionModelPropertiesArgs Empty => new LabelNameConditionModelPropertiesArgs();
    }
}
