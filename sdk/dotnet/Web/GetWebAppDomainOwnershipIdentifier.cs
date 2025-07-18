// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web
{
    public static class GetWebAppDomainOwnershipIdentifier
    {
        /// <summary>
        /// Description for Get domain ownership identifier for web app.
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2016-08-01, 2018-02-01, 2018-11-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetWebAppDomainOwnershipIdentifierResult> InvokeAsync(GetWebAppDomainOwnershipIdentifierArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWebAppDomainOwnershipIdentifierResult>("azure-native:web:getWebAppDomainOwnershipIdentifier", args ?? new GetWebAppDomainOwnershipIdentifierArgs(), options.WithDefaults());

        /// <summary>
        /// Description for Get domain ownership identifier for web app.
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2016-08-01, 2018-02-01, 2018-11-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetWebAppDomainOwnershipIdentifierResult> Invoke(GetWebAppDomainOwnershipIdentifierInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWebAppDomainOwnershipIdentifierResult>("azure-native:web:getWebAppDomainOwnershipIdentifier", args ?? new GetWebAppDomainOwnershipIdentifierInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Description for Get domain ownership identifier for web app.
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2016-08-01, 2018-02-01, 2018-11-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetWebAppDomainOwnershipIdentifierResult> Invoke(GetWebAppDomainOwnershipIdentifierInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetWebAppDomainOwnershipIdentifierResult>("azure-native:web:getWebAppDomainOwnershipIdentifier", args ?? new GetWebAppDomainOwnershipIdentifierInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWebAppDomainOwnershipIdentifierArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of domain ownership identifier.
        /// </summary>
        [Input("domainOwnershipIdentifierName", required: true)]
        public string DomainOwnershipIdentifierName { get; set; } = null!;

        /// <summary>
        /// Name of the app.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetWebAppDomainOwnershipIdentifierArgs()
        {
        }
        public static new GetWebAppDomainOwnershipIdentifierArgs Empty => new GetWebAppDomainOwnershipIdentifierArgs();
    }

    public sealed class GetWebAppDomainOwnershipIdentifierInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of domain ownership identifier.
        /// </summary>
        [Input("domainOwnershipIdentifierName", required: true)]
        public Input<string> DomainOwnershipIdentifierName { get; set; } = null!;

        /// <summary>
        /// Name of the app.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetWebAppDomainOwnershipIdentifierInvokeArgs()
        {
        }
        public static new GetWebAppDomainOwnershipIdentifierInvokeArgs Empty => new GetWebAppDomainOwnershipIdentifierInvokeArgs();
    }


    [OutputType]
    public sealed class GetWebAppDomainOwnershipIdentifierResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
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
        /// Resource type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// String representation of the identity.
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private GetWebAppDomainOwnershipIdentifierResult(
            string azureApiVersion,

            string id,

            string? kind,

            string name,

            string type,

            string? value)
        {
            AzureApiVersion = azureApiVersion;
            Id = id;
            Kind = kind;
            Name = name;
            Type = type;
            Value = value;
        }
    }
}
