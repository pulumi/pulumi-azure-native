// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataLakeAnalytics.Inputs
{

    /// <summary>
    /// The parameters used to create a new compute policy while creating a new Data Lake Analytics account.
    /// </summary>
    public sealed class CreateComputePolicyWithAccountParametersArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum degree of parallelism per job this user can use to submit jobs. This property, the min priority per job property, or both must be passed.
        /// </summary>
        [Input("maxDegreeOfParallelismPerJob")]
        public Input<int>? MaxDegreeOfParallelismPerJob { get; set; }

        /// <summary>
        /// The minimum priority per job this user can use to submit jobs. This property, the max degree of parallelism per job property, or both must be passed.
        /// </summary>
        [Input("minPriorityPerJob")]
        public Input<int>? MinPriorityPerJob { get; set; }

        /// <summary>
        /// The unique name of the compute policy to create.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The AAD object identifier for the entity to create a policy for.
        /// </summary>
        [Input("objectId", required: true)]
        public Input<string> ObjectId { get; set; } = null!;

        /// <summary>
        /// The type of AAD object the object identifier refers to.
        /// </summary>
        [Input("objectType", required: true)]
        public InputUnion<string, Pulumi.AzureNative.DataLakeAnalytics.AADObjectType> ObjectType { get; set; } = null!;

        public CreateComputePolicyWithAccountParametersArgs()
        {
        }
        public static new CreateComputePolicyWithAccountParametersArgs Empty => new CreateComputePolicyWithAccountParametersArgs();
    }
}
