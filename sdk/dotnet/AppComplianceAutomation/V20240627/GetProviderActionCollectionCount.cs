// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AppComplianceAutomation.V20240627
{
    public static class GetProviderActionCollectionCount
    {
        /// <summary>
        /// Get the count of reports.
        /// </summary>
        public static Task<GetProviderActionCollectionCountResult> InvokeAsync(GetProviderActionCollectionCountArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetProviderActionCollectionCountResult>("azure-native:appcomplianceautomation/v20240627:getProviderActionCollectionCount", args ?? new GetProviderActionCollectionCountArgs(), options.WithDefaults());

        /// <summary>
        /// Get the count of reports.
        /// </summary>
        public static Output<GetProviderActionCollectionCountResult> Invoke(GetProviderActionCollectionCountInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetProviderActionCollectionCountResult>("azure-native:appcomplianceautomation/v20240627:getProviderActionCollectionCount", args ?? new GetProviderActionCollectionCountInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get the count of reports.
        /// </summary>
        public static Output<GetProviderActionCollectionCountResult> Invoke(GetProviderActionCollectionCountInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetProviderActionCollectionCountResult>("azure-native:appcomplianceautomation/v20240627:getProviderActionCollectionCount", args ?? new GetProviderActionCollectionCountInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProviderActionCollectionCountArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The resource type.
        /// </summary>
        [Input("type")]
        public string? Type { get; set; }

        public GetProviderActionCollectionCountArgs()
        {
        }
        public static new GetProviderActionCollectionCountArgs Empty => new GetProviderActionCollectionCountArgs();
    }

    public sealed class GetProviderActionCollectionCountInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The resource type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public GetProviderActionCollectionCountInvokeArgs()
        {
        }
        public static new GetProviderActionCollectionCountInvokeArgs Empty => new GetProviderActionCollectionCountInvokeArgs();
    }


    [OutputType]
    public sealed class GetProviderActionCollectionCountResult
    {
        /// <summary>
        /// The count of the specified resource.
        /// </summary>
        public readonly int? Count;

        [OutputConstructor]
        private GetProviderActionCollectionCountResult(int? count)
        {
            Count = count;
        }
    }
}
