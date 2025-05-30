// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.TestBase
{
    public static class GetCustomerEvent
    {
        /// <summary>
        /// Gets a Test Base CustomerEvent.
        /// 
        /// Uses Azure REST API version 2023-11-01-preview.
        /// 
        /// Other available API versions: 2022-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native testbase [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetCustomerEventResult> InvokeAsync(GetCustomerEventArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCustomerEventResult>("azure-native:testbase:getCustomerEvent", args ?? new GetCustomerEventArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a Test Base CustomerEvent.
        /// 
        /// Uses Azure REST API version 2023-11-01-preview.
        /// 
        /// Other available API versions: 2022-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native testbase [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetCustomerEventResult> Invoke(GetCustomerEventInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCustomerEventResult>("azure-native:testbase:getCustomerEvent", args ?? new GetCustomerEventInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a Test Base CustomerEvent.
        /// 
        /// Uses Azure REST API version 2023-11-01-preview.
        /// 
        /// Other available API versions: 2022-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native testbase [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetCustomerEventResult> Invoke(GetCustomerEventInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetCustomerEventResult>("azure-native:testbase:getCustomerEvent", args ?? new GetCustomerEventInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCustomerEventArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The resource name of the Test Base Customer event.
        /// </summary>
        [Input("customerEventName", required: true)]
        public string CustomerEventName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The resource name of the Test Base Account.
        /// </summary>
        [Input("testBaseAccountName", required: true)]
        public string TestBaseAccountName { get; set; } = null!;

        public GetCustomerEventArgs()
        {
        }
        public static new GetCustomerEventArgs Empty => new GetCustomerEventArgs();
    }

    public sealed class GetCustomerEventInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The resource name of the Test Base Customer event.
        /// </summary>
        [Input("customerEventName", required: true)]
        public Input<string> CustomerEventName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The resource name of the Test Base Account.
        /// </summary>
        [Input("testBaseAccountName", required: true)]
        public Input<string> TestBaseAccountName { get; set; } = null!;

        public GetCustomerEventInvokeArgs()
        {
        }
        public static new GetCustomerEventInvokeArgs Empty => new GetCustomerEventInvokeArgs();
    }


    [OutputType]
    public sealed class GetCustomerEventResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// The name of the event subscribed to.
        /// </summary>
        public readonly string EventName;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The notification event receivers.
        /// </summary>
        public readonly ImmutableArray<Outputs.NotificationEventReceiverResponse> Receivers;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetCustomerEventResult(
            string azureApiVersion,

            string eventName,

            string id,

            string name,

            ImmutableArray<Outputs.NotificationEventReceiverResponse> receivers,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            EventName = eventName;
            Id = id;
            Name = name;
            Receivers = receivers;
            SystemData = systemData;
            Type = type;
        }
    }
}
