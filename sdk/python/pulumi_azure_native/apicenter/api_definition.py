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

__all__ = ['ApiDefinitionArgs', 'ApiDefinition']

@pulumi.input_type
class ApiDefinitionArgs:
    def __init__(__self__, *,
                 api_name: pulumi.Input[builtins.str],
                 resource_group_name: pulumi.Input[builtins.str],
                 service_name: pulumi.Input[builtins.str],
                 title: pulumi.Input[builtins.str],
                 version_name: pulumi.Input[builtins.str],
                 workspace_name: pulumi.Input[builtins.str],
                 definition_name: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a ApiDefinition resource.
        :param pulumi.Input[builtins.str] api_name: The name of the API.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] service_name: The name of Azure API Center service.
        :param pulumi.Input[builtins.str] title: API definition title.
        :param pulumi.Input[builtins.str] version_name: The name of the API version.
        :param pulumi.Input[builtins.str] workspace_name: The name of the workspace.
        :param pulumi.Input[builtins.str] definition_name: The name of the API definition.
        :param pulumi.Input[builtins.str] description: API definition description.
        """
        pulumi.set(__self__, "api_name", api_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "service_name", service_name)
        pulumi.set(__self__, "title", title)
        pulumi.set(__self__, "version_name", version_name)
        pulumi.set(__self__, "workspace_name", workspace_name)
        if definition_name is not None:
            pulumi.set(__self__, "definition_name", definition_name)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter(name="apiName")
    def api_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the API.
        """
        return pulumi.get(self, "api_name")

    @api_name.setter
    def api_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "api_name", value)

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
        API definition title.
        """
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "title", value)

    @property
    @pulumi.getter(name="versionName")
    def version_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the API version.
        """
        return pulumi.get(self, "version_name")

    @version_name.setter
    def version_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "version_name", value)

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
    @pulumi.getter(name="definitionName")
    def definition_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the API definition.
        """
        return pulumi.get(self, "definition_name")

    @definition_name.setter
    def definition_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "definition_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        API definition description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)


@pulumi.type_token("azure-native:apicenter:ApiDefinition")
class ApiDefinition(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_name: Optional[pulumi.Input[builtins.str]] = None,
                 definition_name: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 title: Optional[pulumi.Input[builtins.str]] = None,
                 version_name: Optional[pulumi.Input[builtins.str]] = None,
                 workspace_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        API definition entity.

        Uses Azure REST API version 2024-03-15-preview. In version 2.x of the Azure Native provider, it used API version 2024-03-01.

        Other available API versions: 2024-03-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apicenter [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] api_name: The name of the API.
        :param pulumi.Input[builtins.str] definition_name: The name of the API definition.
        :param pulumi.Input[builtins.str] description: API definition description.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] service_name: The name of Azure API Center service.
        :param pulumi.Input[builtins.str] title: API definition title.
        :param pulumi.Input[builtins.str] version_name: The name of the API version.
        :param pulumi.Input[builtins.str] workspace_name: The name of the workspace.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApiDefinitionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        API definition entity.

        Uses Azure REST API version 2024-03-15-preview. In version 2.x of the Azure Native provider, it used API version 2024-03-01.

        Other available API versions: 2024-03-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apicenter [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param ApiDefinitionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApiDefinitionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_name: Optional[pulumi.Input[builtins.str]] = None,
                 definition_name: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 title: Optional[pulumi.Input[builtins.str]] = None,
                 version_name: Optional[pulumi.Input[builtins.str]] = None,
                 workspace_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApiDefinitionArgs.__new__(ApiDefinitionArgs)

            if api_name is None and not opts.urn:
                raise TypeError("Missing required property 'api_name'")
            __props__.__dict__["api_name"] = api_name
            __props__.__dict__["definition_name"] = definition_name
            __props__.__dict__["description"] = description
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
            if title is None and not opts.urn:
                raise TypeError("Missing required property 'title'")
            __props__.__dict__["title"] = title
            if version_name is None and not opts.urn:
                raise TypeError("Missing required property 'version_name'")
            __props__.__dict__["version_name"] = version_name
            if workspace_name is None and not opts.urn:
                raise TypeError("Missing required property 'workspace_name'")
            __props__.__dict__["workspace_name"] = workspace_name
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["specification"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:apicenter/v20240301:ApiDefinition"), pulumi.Alias(type_="azure-native:apicenter/v20240315preview:ApiDefinition"), pulumi.Alias(type_="azure-native:apicenter/v20240601preview:ApiDefinition")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ApiDefinition, __self__).__init__(
            'azure-native:apicenter:ApiDefinition',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ApiDefinition':
        """
        Get an existing ApiDefinition resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ApiDefinitionArgs.__new__(ApiDefinitionArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["specification"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["title"] = None
        __props__.__dict__["type"] = None
        return ApiDefinition(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        API definition description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def specification(self) -> pulumi.Output['outputs.ApiDefinitionPropertiesSpecificationResponse']:
        """
        API specification details.
        """
        return pulumi.get(self, "specification")

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
        API definition title.
        """
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

