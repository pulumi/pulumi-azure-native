# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities

__all__ = ['JobPrivateEndpointArgs', 'JobPrivateEndpoint']

@pulumi.input_type
class JobPrivateEndpointArgs:
    def __init__(__self__, *,
                 job_agent_name: pulumi.Input[builtins.str],
                 resource_group_name: pulumi.Input[builtins.str],
                 server_name: pulumi.Input[builtins.str],
                 target_server_azure_resource_id: pulumi.Input[builtins.str],
                 private_endpoint_name: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a JobPrivateEndpoint resource.
        :param pulumi.Input[builtins.str] job_agent_name: The name of the job agent.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        :param pulumi.Input[builtins.str] server_name: The name of the server.
        :param pulumi.Input[builtins.str] target_server_azure_resource_id: ARM resource id of the server the private endpoint will target.
        :param pulumi.Input[builtins.str] private_endpoint_name: The name of the private endpoint.
        """
        pulumi.set(__self__, "job_agent_name", job_agent_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "server_name", server_name)
        pulumi.set(__self__, "target_server_azure_resource_id", target_server_azure_resource_id)
        if private_endpoint_name is not None:
            pulumi.set(__self__, "private_endpoint_name", private_endpoint_name)

    @property
    @pulumi.getter(name="jobAgentName")
    def job_agent_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the job agent.
        """
        return pulumi.get(self, "job_agent_name")

    @job_agent_name.setter
    def job_agent_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "job_agent_name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="serverName")
    def server_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the server.
        """
        return pulumi.get(self, "server_name")

    @server_name.setter
    def server_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "server_name", value)

    @property
    @pulumi.getter(name="targetServerAzureResourceId")
    def target_server_azure_resource_id(self) -> pulumi.Input[builtins.str]:
        """
        ARM resource id of the server the private endpoint will target.
        """
        return pulumi.get(self, "target_server_azure_resource_id")

    @target_server_azure_resource_id.setter
    def target_server_azure_resource_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "target_server_azure_resource_id", value)

    @property
    @pulumi.getter(name="privateEndpointName")
    def private_endpoint_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the private endpoint.
        """
        return pulumi.get(self, "private_endpoint_name")

    @private_endpoint_name.setter
    def private_endpoint_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "private_endpoint_name", value)


@pulumi.type_token("azure-native:sql:JobPrivateEndpoint")
class JobPrivateEndpoint(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 job_agent_name: Optional[pulumi.Input[builtins.str]] = None,
                 private_endpoint_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 server_name: Optional[pulumi.Input[builtins.str]] = None,
                 target_server_azure_resource_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        A job agent private endpoint.

        Uses Azure REST API version 2023-08-01. In version 2.x of the Azure Native provider, it used API version 2023-05-01-preview.

        Other available API versions: 2023-05-01-preview, 2023-08-01-preview, 2024-05-01-preview, 2024-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] job_agent_name: The name of the job agent.
        :param pulumi.Input[builtins.str] private_endpoint_name: The name of the private endpoint.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        :param pulumi.Input[builtins.str] server_name: The name of the server.
        :param pulumi.Input[builtins.str] target_server_azure_resource_id: ARM resource id of the server the private endpoint will target.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: JobPrivateEndpointArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A job agent private endpoint.

        Uses Azure REST API version 2023-08-01. In version 2.x of the Azure Native provider, it used API version 2023-05-01-preview.

        Other available API versions: 2023-05-01-preview, 2023-08-01-preview, 2024-05-01-preview, 2024-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param JobPrivateEndpointArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(JobPrivateEndpointArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 job_agent_name: Optional[pulumi.Input[builtins.str]] = None,
                 private_endpoint_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 server_name: Optional[pulumi.Input[builtins.str]] = None,
                 target_server_azure_resource_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = JobPrivateEndpointArgs.__new__(JobPrivateEndpointArgs)

            if job_agent_name is None and not opts.urn:
                raise TypeError("Missing required property 'job_agent_name'")
            __props__.__dict__["job_agent_name"] = job_agent_name
            __props__.__dict__["private_endpoint_name"] = private_endpoint_name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            if server_name is None and not opts.urn:
                raise TypeError("Missing required property 'server_name'")
            __props__.__dict__["server_name"] = server_name
            if target_server_azure_resource_id is None and not opts.urn:
                raise TypeError("Missing required property 'target_server_azure_resource_id'")
            __props__.__dict__["target_server_azure_resource_id"] = target_server_azure_resource_id
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["private_endpoint_id"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:sql/v20230501preview:JobPrivateEndpoint"), pulumi.Alias(type_="azure-native:sql/v20230801:JobPrivateEndpoint"), pulumi.Alias(type_="azure-native:sql/v20230801preview:JobPrivateEndpoint"), pulumi.Alias(type_="azure-native:sql/v20240501preview:JobPrivateEndpoint"), pulumi.Alias(type_="azure-native:sql/v20241101preview:JobPrivateEndpoint")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(JobPrivateEndpoint, __self__).__init__(
            'azure-native:sql:JobPrivateEndpoint',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'JobPrivateEndpoint':
        """
        Get an existing JobPrivateEndpoint resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = JobPrivateEndpointArgs.__new__(JobPrivateEndpointArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["private_endpoint_id"] = None
        __props__.__dict__["target_server_azure_resource_id"] = None
        __props__.__dict__["type"] = None
        return JobPrivateEndpoint(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="privateEndpointId")
    def private_endpoint_id(self) -> pulumi.Output[builtins.str]:
        """
        Private endpoint id of the private endpoint.
        """
        return pulumi.get(self, "private_endpoint_id")

    @property
    @pulumi.getter(name="targetServerAzureResourceId")
    def target_server_azure_resource_id(self) -> pulumi.Output[builtins.str]:
        """
        ARM resource id of the server the private endpoint will target.
        """
        return pulumi.get(self, "target_server_azure_resource_id")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

