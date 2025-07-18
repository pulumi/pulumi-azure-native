// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web
{
    public static class ListWebAppFunctionKeys
    {
        /// <summary>
        /// Description for Get function keys for a function in a web site, or a deployment slot.
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2018-02-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<ListWebAppFunctionKeysResult> InvokeAsync(ListWebAppFunctionKeysArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListWebAppFunctionKeysResult>("azure-native:web:listWebAppFunctionKeys", args ?? new ListWebAppFunctionKeysArgs(), options.WithDefaults());

        /// <summary>
        /// Description for Get function keys for a function in a web site, or a deployment slot.
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2018-02-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListWebAppFunctionKeysResult> Invoke(ListWebAppFunctionKeysInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListWebAppFunctionKeysResult>("azure-native:web:listWebAppFunctionKeys", args ?? new ListWebAppFunctionKeysInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Description for Get function keys for a function in a web site, or a deployment slot.
        /// 
        /// Uses Azure REST API version 2024-04-01.
        /// 
        /// Other available API versions: 2018-02-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListWebAppFunctionKeysResult> Invoke(ListWebAppFunctionKeysInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListWebAppFunctionKeysResult>("azure-native:web:listWebAppFunctionKeys", args ?? new ListWebAppFunctionKeysInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListWebAppFunctionKeysArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Function name.
        /// </summary>
        [Input("functionName", required: true)]
        public string FunctionName { get; set; } = null!;

        /// <summary>
        /// Site name.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public ListWebAppFunctionKeysArgs()
        {
        }
        public static new ListWebAppFunctionKeysArgs Empty => new ListWebAppFunctionKeysArgs();
    }

    public sealed class ListWebAppFunctionKeysInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Function name.
        /// </summary>
        [Input("functionName", required: true)]
        public Input<string> FunctionName { get; set; } = null!;

        /// <summary>
        /// Site name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public ListWebAppFunctionKeysInvokeArgs()
        {
        }
        public static new ListWebAppFunctionKeysInvokeArgs Empty => new ListWebAppFunctionKeysInvokeArgs();
    }


    [OutputType]
    public sealed class ListWebAppFunctionKeysResult
    {
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
        /// Settings.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Properties;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ListWebAppFunctionKeysResult(
            string id,

            string? kind,

            string name,

            ImmutableDictionary<string, string> properties,

            string type)
        {
            Id = id;
            Kind = kind;
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}
