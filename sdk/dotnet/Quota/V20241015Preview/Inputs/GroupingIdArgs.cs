// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Quota.V20241015Preview.Inputs
{

    /// <summary>
    /// The grouping Id for the group quota. It can be Billing Id or ServiceTreeId if applicable. 
    /// </summary>
    public sealed class GroupingIdArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// GroupingId type. It is a required property. More types of groupIds can be supported in future.
        /// </summary>
        [Input("groupingIdType")]
        public InputUnion<string, Pulumi.AzureNative.Quota.V20241015Preview.GroupingIdType>? GroupingIdType { get; set; }

        /// <summary>
        /// GroupId value based on the groupingType selected - Billing Id or ServiceTreeId.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public GroupingIdArgs()
        {
        }
        public static new GroupingIdArgs Empty => new GroupingIdArgs();
    }
}
