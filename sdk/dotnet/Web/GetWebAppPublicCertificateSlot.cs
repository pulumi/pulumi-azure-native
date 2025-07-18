// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web
{
    public static class GetWebAppPublicCertificateSlot
    {
        /// <summary>
        /// Description for Get the named public certificate for an app (or deployment slot, if specified).
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2016-08-01, 2018-02-01, 2018-11-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetWebAppPublicCertificateSlotResult> InvokeAsync(GetWebAppPublicCertificateSlotArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWebAppPublicCertificateSlotResult>("azure-native:web:getWebAppPublicCertificateSlot", args ?? new GetWebAppPublicCertificateSlotArgs(), options.WithDefaults());

        /// <summary>
        /// Description for Get the named public certificate for an app (or deployment slot, if specified).
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2016-08-01, 2018-02-01, 2018-11-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetWebAppPublicCertificateSlotResult> Invoke(GetWebAppPublicCertificateSlotInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWebAppPublicCertificateSlotResult>("azure-native:web:getWebAppPublicCertificateSlot", args ?? new GetWebAppPublicCertificateSlotInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Description for Get the named public certificate for an app (or deployment slot, if specified).
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2016-08-01, 2018-02-01, 2018-11-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetWebAppPublicCertificateSlotResult> Invoke(GetWebAppPublicCertificateSlotInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetWebAppPublicCertificateSlotResult>("azure-native:web:getWebAppPublicCertificateSlot", args ?? new GetWebAppPublicCertificateSlotInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWebAppPublicCertificateSlotArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the app.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Public certificate name.
        /// </summary>
        [Input("publicCertificateName", required: true)]
        public string PublicCertificateName { get; set; } = null!;

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of the deployment slot. If a slot is not specified, the API the named binding for the production slot.
        /// </summary>
        [Input("slot", required: true)]
        public string Slot { get; set; } = null!;

        public GetWebAppPublicCertificateSlotArgs()
        {
        }
        public static new GetWebAppPublicCertificateSlotArgs Empty => new GetWebAppPublicCertificateSlotArgs();
    }

    public sealed class GetWebAppPublicCertificateSlotInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the app.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Public certificate name.
        /// </summary>
        [Input("publicCertificateName", required: true)]
        public Input<string> PublicCertificateName { get; set; } = null!;

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of the deployment slot. If a slot is not specified, the API the named binding for the production slot.
        /// </summary>
        [Input("slot", required: true)]
        public Input<string> Slot { get; set; } = null!;

        public GetWebAppPublicCertificateSlotInvokeArgs()
        {
        }
        public static new GetWebAppPublicCertificateSlotInvokeArgs Empty => new GetWebAppPublicCertificateSlotInvokeArgs();
    }


    [OutputType]
    public sealed class GetWebAppPublicCertificateSlotResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Public Certificate byte array
        /// </summary>
        public readonly string? Blob;
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
        /// Public Certificate Location
        /// </summary>
        public readonly string? PublicCertificateLocation;
        /// <summary>
        /// Certificate Thumbprint
        /// </summary>
        public readonly string Thumbprint;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetWebAppPublicCertificateSlotResult(
            string azureApiVersion,

            string? blob,

            string id,

            string? kind,

            string name,

            string? publicCertificateLocation,

            string thumbprint,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Blob = blob;
            Id = id;
            Kind = kind;
            Name = name;
            PublicCertificateLocation = publicCertificateLocation;
            Thumbprint = thumbprint;
            Type = type;
        }
    }
}
