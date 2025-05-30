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
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['EnvironmentArgs', 'Environment']

@pulumi.input_type
class EnvironmentArgs:
    def __init__(__self__, *,
                 kind: pulumi.Input[Union[builtins.str, 'EnvironmentKind']],
                 resource_group_name: pulumi.Input[builtins.str],
                 service_name: pulumi.Input[builtins.str],
                 title: pulumi.Input[builtins.str],
                 workspace_name: pulumi.Input[builtins.str],
                 custom_properties: Optional[Any] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 environment_name: Optional[pulumi.Input[builtins.str]] = None,
                 onboarding: Optional[pulumi.Input['OnboardingArgs']] = None,
                 server: Optional[pulumi.Input['EnvironmentServerArgs']] = None):
        """
        The set of arguments for constructing a Environment resource.
        :param pulumi.Input[Union[builtins.str, 'EnvironmentKind']] kind: Environment kind.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] service_name: The name of Azure API Center service.
        :param pulumi.Input[builtins.str] title: Environment title.
        :param pulumi.Input[builtins.str] workspace_name: The name of the workspace.
        :param Any custom_properties: The custom metadata defined for API catalog entities.
        :param pulumi.Input[builtins.str] description: The environment description.
        :param pulumi.Input[builtins.str] environment_name: The name of the environment.
        :param pulumi.Input['OnboardingArgs'] onboarding: Environment onboarding information
        :param pulumi.Input['EnvironmentServerArgs'] server: Server information of the environment.
        """
        pulumi.set(__self__, "kind", kind)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "service_name", service_name)
        pulumi.set(__self__, "title", title)
        pulumi.set(__self__, "workspace_name", workspace_name)
        if custom_properties is not None:
            pulumi.set(__self__, "custom_properties", custom_properties)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if environment_name is not None:
            pulumi.set(__self__, "environment_name", environment_name)
        if onboarding is not None:
            pulumi.set(__self__, "onboarding", onboarding)
        if server is not None:
            pulumi.set(__self__, "server", server)

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Input[Union[builtins.str, 'EnvironmentKind']]:
        """
        Environment kind.
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: pulumi.Input[Union[builtins.str, 'EnvironmentKind']]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the resource group. The name is case insensitive.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of Azure API Center service.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter
    def title(self) -> pulumi.Input[builtins.str]:
        """
        Environment title.
        """
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "title", value)

    @property
    @pulumi.getter(name="workspaceName")
    def workspace_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the workspace.
        """
        return pulumi.get(self, "workspace_name")

    @workspace_name.setter
    def workspace_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "workspace_name", value)

    @property
    @pulumi.getter(name="customProperties")
    def custom_properties(self) -> Optional[Any]:
        """
        The custom metadata defined for API catalog entities.
        """
        return pulumi.get(self, "custom_properties")

    @custom_properties.setter
    def custom_properties(self, value: Optional[Any]):
        pulumi.set(self, "custom_properties", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The environment description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="environmentName")
    def environment_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the environment.
        """
        return pulumi.get(self, "environment_name")

    @environment_name.setter
    def environment_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "environment_name", value)

    @property
    @pulumi.getter
    def onboarding(self) -> Optional[pulumi.Input['OnboardingArgs']]:
        """
        Environment onboarding information
        """
        return pulumi.get(self, "onboarding")

    @onboarding.setter
    def onboarding(self, value: Optional[pulumi.Input['OnboardingArgs']]):
        pulumi.set(self, "onboarding", value)

    @property
    @pulumi.getter
    def server(self) -> Optional[pulumi.Input['EnvironmentServerArgs']]:
        """
        Server information of the environment.
        """
        return pulumi.get(self, "server")

    @server.setter
    def server(self, value: Optional[pulumi.Input['EnvironmentServerArgs']]):
        pulumi.set(self, "server", value)


@pulumi.type_token("azure-native:apicenter:Environment")
class Environment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 custom_properties: Optional[Any] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 environment_name: Optional[pulumi.Input[builtins.str]] = None,
                 kind: Optional[pulumi.Input[Union[builtins.str, 'EnvironmentKind']]] = None,
                 onboarding: Optional[pulumi.Input[Union['OnboardingArgs', 'OnboardingArgsDict']]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 server: Optional[pulumi.Input[Union['EnvironmentServerArgs', 'EnvironmentServerArgsDict']]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 title: Optional[pulumi.Input[builtins.str]] = None,
                 workspace_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Environment entity.

        Uses Azure REST API version 2024-03-15-preview. In version 2.x of the Azure Native provider, it used API version 2024-03-01.

        Other available API versions: 2024-03-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apicenter [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param Any custom_properties: The custom metadata defined for API catalog entities.
        :param pulumi.Input[builtins.str] description: The environment description.
        :param pulumi.Input[builtins.str] environment_name: The name of the environment.
        :param pulumi.Input[Union[builtins.str, 'EnvironmentKind']] kind: Environment kind.
        :param pulumi.Input[Union['OnboardingArgs', 'OnboardingArgsDict']] onboarding: Environment onboarding information
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[Union['EnvironmentServerArgs', 'EnvironmentServerArgsDict']] server: Server information of the environment.
        :param pulumi.Input[builtins.str] service_name: The name of Azure API Center service.
        :param pulumi.Input[builtins.str] title: Environment title.
        :param pulumi.Input[builtins.str] workspace_name: The name of the workspace.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EnvironmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Environment entity.

        Uses Azure REST API version 2024-03-15-preview. In version 2.x of the Azure Native provider, it used API version 2024-03-01.

        Other available API versions: 2024-03-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apicenter [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param EnvironmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EnvironmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 custom_properties: Optional[Any] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 environment_name: Optional[pulumi.Input[builtins.str]] = None,
                 kind: Optional[pulumi.Input[Union[builtins.str, 'EnvironmentKind']]] = None,
                 onboarding: Optional[pulumi.Input[Union['OnboardingArgs', 'OnboardingArgsDict']]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 server: Optional[pulumi.Input[Union['EnvironmentServerArgs', 'EnvironmentServerArgsDict']]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 title: Optional[pulumi.Input[builtins.str]] = None,
                 workspace_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EnvironmentArgs.__new__(EnvironmentArgs)

            __props__.__dict__["custom_properties"] = custom_properties
            __props__.__dict__["description"] = description
            __props__.__dict__["environment_name"] = environment_name
            if kind is None and not opts.urn:
                raise TypeError("Missing required property 'kind'")
            __props__.__dict__["kind"] = kind
            __props__.__dict__["onboarding"] = onboarding
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["server"] = server
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
            if title is None and not opts.urn:
                raise TypeError("Missing required property 'title'")
            __props__.__dict__["title"] = title
            if workspace_name is None and not opts.urn:
                raise TypeError("Missing required property 'workspace_name'")
            __props__.__dict__["workspace_name"] = workspace_name
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:apicenter/v20240301:Environment"), pulumi.Alias(type_="azure-native:apicenter/v20240315preview:Environment"), pulumi.Alias(type_="azure-native:apicenter/v20240601preview:Environment")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Environment, __self__).__init__(
            'azure-native:apicenter:Environment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Environment':
        """
        Get an existing Environment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = EnvironmentArgs.__new__(EnvironmentArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["custom_properties"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["onboarding"] = None
        __props__.__dict__["server"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["title"] = None
        __props__.__dict__["type"] = None
        return Environment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="customProperties")
    def custom_properties(self) -> pulumi.Output[Optional[Any]]:
        """
        The custom metadata defined for API catalog entities.
        """
        return pulumi.get(self, "custom_properties")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The environment description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[builtins.str]:
        """
        Environment kind.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def onboarding(self) -> pulumi.Output[Optional['outputs.OnboardingResponse']]:
        """
        Environment onboarding information
        """
        return pulumi.get(self, "onboarding")

    @property
    @pulumi.getter
    def server(self) -> pulumi.Output[Optional['outputs.EnvironmentServerResponse']]:
        """
        Server information of the environment.
        """
        return pulumi.get(self, "server")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> pulumi.Output['outputs.SystemDataResponse']:
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def title(self) -> pulumi.Output[builtins.str]:
        """
        Environment title.
        """
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

