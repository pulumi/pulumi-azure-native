// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Sql.V20230801
{
    public static class GetJobTargetGroup
    {
        /// <summary>
        /// Gets a target group.
        /// </summary>
        public static Task<GetJobTargetGroupResult> InvokeAsync(GetJobTargetGroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetJobTargetGroupResult>("azure-native:sql/v20230801:getJobTargetGroup", args ?? new GetJobTargetGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a target group.
        /// </summary>
        public static Output<GetJobTargetGroupResult> Invoke(GetJobTargetGroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetJobTargetGroupResult>("azure-native:sql/v20230801:getJobTargetGroup", args ?? new GetJobTargetGroupInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a target group.
        /// </summary>
        public static Output<GetJobTargetGroupResult> Invoke(GetJobTargetGroupInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetJobTargetGroupResult>("azure-native:sql/v20230801:getJobTargetGroup", args ?? new GetJobTargetGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetJobTargetGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the job agent.
        /// </summary>
        [Input("jobAgentName", required: true)]
        public string JobAgentName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the server.
        /// </summary>
        [Input("serverName", required: true)]
        public string ServerName { get; set; } = null!;

        /// <summary>
        /// The name of the target group.
        /// </summary>
        [Input("targetGroupName", required: true)]
        public string TargetGroupName { get; set; } = null!;

        public GetJobTargetGroupArgs()
        {
        }
        public static new GetJobTargetGroupArgs Empty => new GetJobTargetGroupArgs();
    }

    public sealed class GetJobTargetGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the job agent.
        /// </summary>
        [Input("jobAgentName", required: true)]
        public Input<string> JobAgentName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the server.
        /// </summary>
        [Input("serverName", required: true)]
        public Input<string> ServerName { get; set; } = null!;

        /// <summary>
        /// The name of the target group.
        /// </summary>
        [Input("targetGroupName", required: true)]
        public Input<string> TargetGroupName { get; set; } = null!;

        public GetJobTargetGroupInvokeArgs()
        {
        }
        public static new GetJobTargetGroupInvokeArgs Empty => new GetJobTargetGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetJobTargetGroupResult
    {
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Members of the target group.
        /// </summary>
        public readonly ImmutableArray<Outputs.JobTargetResponse> Members;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetJobTargetGroupResult(
            string id,

            ImmutableArray<Outputs.JobTargetResponse> members,

            string name,

            string type)
        {
            Id = id;
            Members = members;
            Name = name;
            Type = type;
        }
    }
}
