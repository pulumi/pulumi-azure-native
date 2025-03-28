// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.LabServices
{
    public static class GetGlobalUserOperationBatchStatus
    {
        /// <summary>
        /// Get batch operation status
        /// 
        /// Uses Azure REST API version 2018-10-15.
        /// </summary>
        public static Task<GetGlobalUserOperationBatchStatusResult> InvokeAsync(GetGlobalUserOperationBatchStatusArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGlobalUserOperationBatchStatusResult>("azure-native:labservices:getGlobalUserOperationBatchStatus", args ?? new GetGlobalUserOperationBatchStatusArgs(), options.WithDefaults());

        /// <summary>
        /// Get batch operation status
        /// 
        /// Uses Azure REST API version 2018-10-15.
        /// </summary>
        public static Output<GetGlobalUserOperationBatchStatusResult> Invoke(GetGlobalUserOperationBatchStatusInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGlobalUserOperationBatchStatusResult>("azure-native:labservices:getGlobalUserOperationBatchStatus", args ?? new GetGlobalUserOperationBatchStatusInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get batch operation status
        /// 
        /// Uses Azure REST API version 2018-10-15.
        /// </summary>
        public static Output<GetGlobalUserOperationBatchStatusResult> Invoke(GetGlobalUserOperationBatchStatusInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetGlobalUserOperationBatchStatusResult>("azure-native:labservices:getGlobalUserOperationBatchStatus", args ?? new GetGlobalUserOperationBatchStatusInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGlobalUserOperationBatchStatusArgs : global::Pulumi.InvokeArgs
    {
        [Input("urls", required: true)]
        private List<string>? _urls;

        /// <summary>
        /// The operation url of long running operation
        /// </summary>
        public List<string> Urls
        {
            get => _urls ?? (_urls = new List<string>());
            set => _urls = value;
        }

        /// <summary>
        /// The name of the user.
        /// </summary>
        [Input("userName", required: true)]
        public string UserName { get; set; } = null!;

        public GetGlobalUserOperationBatchStatusArgs()
        {
        }
        public static new GetGlobalUserOperationBatchStatusArgs Empty => new GetGlobalUserOperationBatchStatusArgs();
    }

    public sealed class GetGlobalUserOperationBatchStatusInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("urls", required: true)]
        private InputList<string>? _urls;

        /// <summary>
        /// The operation url of long running operation
        /// </summary>
        public InputList<string> Urls
        {
            get => _urls ?? (_urls = new InputList<string>());
            set => _urls = value;
        }

        /// <summary>
        /// The name of the user.
        /// </summary>
        [Input("userName", required: true)]
        public Input<string> UserName { get; set; } = null!;

        public GetGlobalUserOperationBatchStatusInvokeArgs()
        {
        }
        public static new GetGlobalUserOperationBatchStatusInvokeArgs Empty => new GetGlobalUserOperationBatchStatusInvokeArgs();
    }


    [OutputType]
    public sealed class GetGlobalUserOperationBatchStatusResult
    {
        /// <summary>
        /// Gets a collection of items that contain the operation url and status.
        /// </summary>
        public readonly ImmutableArray<Outputs.OperationBatchStatusResponseItemResponse> Items;

        [OutputConstructor]
        private GetGlobalUserOperationBatchStatusResult(ImmutableArray<Outputs.OperationBatchStatusResponseItemResponse> items)
        {
            Items = items;
        }
    }
}
