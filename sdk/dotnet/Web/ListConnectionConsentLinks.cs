// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web
{
    public static class ListConnectionConsentLinks
    {
        /// <summary>
        /// Lists the consent links of a connection
        /// 
        /// Uses Azure REST API version 2016-06-01.
        /// 
        /// Other available API versions: 2015-08-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<ListConnectionConsentLinksResult> InvokeAsync(ListConnectionConsentLinksArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListConnectionConsentLinksResult>("azure-native:web:listConnectionConsentLinks", args ?? new ListConnectionConsentLinksArgs(), options.WithDefaults());

        /// <summary>
        /// Lists the consent links of a connection
        /// 
        /// Uses Azure REST API version 2016-06-01.
        /// 
        /// Other available API versions: 2015-08-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListConnectionConsentLinksResult> Invoke(ListConnectionConsentLinksInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListConnectionConsentLinksResult>("azure-native:web:listConnectionConsentLinks", args ?? new ListConnectionConsentLinksInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Lists the consent links of a connection
        /// 
        /// Uses Azure REST API version 2016-06-01.
        /// 
        /// Other available API versions: 2015-08-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<ListConnectionConsentLinksResult> Invoke(ListConnectionConsentLinksInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<ListConnectionConsentLinksResult>("azure-native:web:listConnectionConsentLinks", args ?? new ListConnectionConsentLinksInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListConnectionConsentLinksArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Connection name
        /// </summary>
        [Input("connectionName", required: true)]
        public string ConnectionName { get; set; } = null!;

        [Input("parameters")]
        private List<Inputs.ConsentLinkParameterDefinition>? _parameters;

        /// <summary>
        /// Collection of resources
        /// </summary>
        public List<Inputs.ConsentLinkParameterDefinition> Parameters
        {
            get => _parameters ?? (_parameters = new List<Inputs.ConsentLinkParameterDefinition>());
            set => _parameters = value;
        }

        /// <summary>
        /// The resource group
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Subscription Id
        /// </summary>
        [Input("subscriptionId")]
        public string? SubscriptionId { get; set; }

        public ListConnectionConsentLinksArgs()
        {
        }
        public static new ListConnectionConsentLinksArgs Empty => new ListConnectionConsentLinksArgs();
    }

    public sealed class ListConnectionConsentLinksInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Connection name
        /// </summary>
        [Input("connectionName", required: true)]
        public Input<string> ConnectionName { get; set; } = null!;

        [Input("parameters")]
        private InputList<Inputs.ConsentLinkParameterDefinitionArgs>? _parameters;

        /// <summary>
        /// Collection of resources
        /// </summary>
        public InputList<Inputs.ConsentLinkParameterDefinitionArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.ConsentLinkParameterDefinitionArgs>());
            set => _parameters = value;
        }

        /// <summary>
        /// The resource group
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Subscription Id
        /// </summary>
        [Input("subscriptionId")]
        public Input<string>? SubscriptionId { get; set; }

        public ListConnectionConsentLinksInvokeArgs()
        {
        }
        public static new ListConnectionConsentLinksInvokeArgs Empty => new ListConnectionConsentLinksInvokeArgs();
    }


    [OutputType]
    public sealed class ListConnectionConsentLinksResult
    {
        /// <summary>
        /// Collection of resources
        /// </summary>
        public readonly ImmutableArray<Outputs.ConsentLinkDefinitionResponse> Value;

        [OutputConstructor]
        private ListConnectionConsentLinksResult(ImmutableArray<Outputs.ConsentLinkDefinitionResponse> value)
        {
            Value = value;
        }
    }
}
