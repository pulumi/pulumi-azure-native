// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web
{
    public static class GetWebAppHybridConnectionSlot
    {
        /// <summary>
        /// Description for Retrieves a specific Service Bus Hybrid Connection used by this Web App.
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2016-08-01, 2018-02-01, 2018-11-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetWebAppHybridConnectionSlotResult> InvokeAsync(GetWebAppHybridConnectionSlotArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWebAppHybridConnectionSlotResult>("azure-native:web:getWebAppHybridConnectionSlot", args ?? new GetWebAppHybridConnectionSlotArgs(), options.WithDefaults());

        /// <summary>
        /// Description for Retrieves a specific Service Bus Hybrid Connection used by this Web App.
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2016-08-01, 2018-02-01, 2018-11-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetWebAppHybridConnectionSlotResult> Invoke(GetWebAppHybridConnectionSlotInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWebAppHybridConnectionSlotResult>("azure-native:web:getWebAppHybridConnectionSlot", args ?? new GetWebAppHybridConnectionSlotInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Description for Retrieves a specific Service Bus Hybrid Connection used by this Web App.
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2016-08-01, 2018-02-01, 2018-11-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetWebAppHybridConnectionSlotResult> Invoke(GetWebAppHybridConnectionSlotInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetWebAppHybridConnectionSlotResult>("azure-native:web:getWebAppHybridConnectionSlot", args ?? new GetWebAppHybridConnectionSlotInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWebAppHybridConnectionSlotArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the web app.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The namespace for this hybrid connection.
        /// </summary>
        [Input("namespaceName", required: true)]
        public string NamespaceName { get; set; } = null!;

        /// <summary>
        /// The relay name for this hybrid connection.
        /// </summary>
        [Input("relayName", required: true)]
        public string RelayName { get; set; } = null!;

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the slot for the web app.
        /// </summary>
        [Input("slot", required: true)]
        public string Slot { get; set; } = null!;

        public GetWebAppHybridConnectionSlotArgs()
        {
        }
        public static new GetWebAppHybridConnectionSlotArgs Empty => new GetWebAppHybridConnectionSlotArgs();
    }

    public sealed class GetWebAppHybridConnectionSlotInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the web app.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The namespace for this hybrid connection.
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        /// <summary>
        /// The relay name for this hybrid connection.
        /// </summary>
        [Input("relayName", required: true)]
        public Input<string> RelayName { get; set; } = null!;

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the slot for the web app.
        /// </summary>
        [Input("slot", required: true)]
        public Input<string> Slot { get; set; } = null!;

        public GetWebAppHybridConnectionSlotInvokeArgs()
        {
        }
        public static new GetWebAppHybridConnectionSlotInvokeArgs Empty => new GetWebAppHybridConnectionSlotInvokeArgs();
    }


    [OutputType]
    public sealed class GetWebAppHybridConnectionSlotResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// The hostname of the endpoint.
        /// </summary>
        public readonly string? Hostname;
        /// <summary>
        /// Resource Id.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Kind of resource.
        /// </summary>
        public readonly string? Kind;
        /// <summary>
        /// Resource Name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The port of the endpoint.
        /// </summary>
        public readonly int? Port;
        /// <summary>
        /// The ARM URI to the Service Bus relay.
        /// </summary>
        public readonly string? RelayArmUri;
        /// <summary>
        /// The name of the Service Bus relay.
        /// </summary>
        public readonly string? RelayName;
        /// <summary>
        /// The name of the Service Bus key which has Send permissions. This is used to authenticate to Service Bus.
        /// </summary>
        public readonly string? SendKeyName;
        /// <summary>
        /// The value of the Service Bus key. This is used to authenticate to Service Bus. In ARM this key will not be returned
        /// normally, use the POST /listKeys API instead.
        /// </summary>
        public readonly string? SendKeyValue;
        /// <summary>
        /// The name of the Service Bus namespace.
        /// </summary>
        public readonly string? ServiceBusNamespace;
        /// <summary>
        /// The suffix for the service bus endpoint. By default this is .servicebus.windows.net
        /// </summary>
        public readonly string? ServiceBusSuffix;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetWebAppHybridConnectionSlotResult(
            string azureApiVersion,

            string? hostname,

            string id,

            string? kind,

            string name,

            int? port,

            string? relayArmUri,

            string? relayName,

            string? sendKeyName,

            string? sendKeyValue,

            string? serviceBusNamespace,

            string? serviceBusSuffix,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Hostname = hostname;
            Id = id;
            Kind = kind;
            Name = name;
            Port = port;
            RelayArmUri = relayArmUri;
            RelayName = relayName;
            SendKeyName = sendKeyName;
            SendKeyValue = sendKeyValue;
            ServiceBusNamespace = serviceBusNamespace;
            ServiceBusSuffix = serviceBusSuffix;
            Type = type;
        }
    }
}
