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
from ._enums import *

__all__ = ['DataMaskingPolicyArgs', 'DataMaskingPolicy']

@pulumi.input_type
class DataMaskingPolicyArgs:
    def __init__(__self__, *,
                 data_masking_state: pulumi.Input['DataMaskingState'],
                 database_name: pulumi.Input[builtins.str],
                 resource_group_name: pulumi.Input[builtins.str],
                 server_name: pulumi.Input[builtins.str],
                 data_masking_policy_name: Optional[pulumi.Input[builtins.str]] = None,
                 exempt_principals: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a DataMaskingPolicy resource.
        :param pulumi.Input['DataMaskingState'] data_masking_state: The state of the data masking policy.
        :param pulumi.Input[builtins.str] database_name: The name of the database.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        :param pulumi.Input[builtins.str] server_name: The name of the server.
        :param pulumi.Input[builtins.str] data_masking_policy_name: The name of the database for which the data masking policy applies.
        :param pulumi.Input[builtins.str] exempt_principals: The list of the exempt principals. Specifies the semicolon-separated list of database users for which the data masking policy does not apply. The specified users receive data results without masking for all of the database queries.
        """
        pulumi.set(__self__, "data_masking_state", data_masking_state)
        pulumi.set(__self__, "database_name", database_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "server_name", server_name)
        if data_masking_policy_name is not None:
            pulumi.set(__self__, "data_masking_policy_name", data_masking_policy_name)
        if exempt_principals is not None:
            pulumi.set(__self__, "exempt_principals", exempt_principals)

    @property
    @pulumi.getter(name="dataMaskingState")
    def data_masking_state(self) -> pulumi.Input['DataMaskingState']:
        """
        The state of the data masking policy.
        """
        return pulumi.get(self, "data_masking_state")

    @data_masking_state.setter
    def data_masking_state(self, value: pulumi.Input['DataMaskingState']):
        pulumi.set(self, "data_masking_state", value)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the database.
        """
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "database_name", value)

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
    @pulumi.getter(name="dataMaskingPolicyName")
    def data_masking_policy_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the database for which the data masking policy applies.
        """
        return pulumi.get(self, "data_masking_policy_name")

    @data_masking_policy_name.setter
    def data_masking_policy_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "data_masking_policy_name", value)

    @property
    @pulumi.getter(name="exemptPrincipals")
    def exempt_principals(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The list of the exempt principals. Specifies the semicolon-separated list of database users for which the data masking policy does not apply. The specified users receive data results without masking for all of the database queries.
        """
        return pulumi.get(self, "exempt_principals")

    @exempt_principals.setter
    def exempt_principals(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "exempt_principals", value)


@pulumi.type_token("azure-native:sql:DataMaskingPolicy")
class DataMaskingPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data_masking_policy_name: Optional[pulumi.Input[builtins.str]] = None,
                 data_masking_state: Optional[pulumi.Input['DataMaskingState']] = None,
                 database_name: Optional[pulumi.Input[builtins.str]] = None,
                 exempt_principals: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 server_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        A database data masking policy.

        Uses Azure REST API version 2023-08-01. In version 2.x of the Azure Native provider, it used API version 2021-11-01.

        Other available API versions: 2014-04-01, 2021-11-01, 2022-02-01-preview, 2022-05-01-preview, 2022-08-01-preview, 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview, 2023-08-01-preview, 2024-05-01-preview, 2024-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] data_masking_policy_name: The name of the database for which the data masking policy applies.
        :param pulumi.Input['DataMaskingState'] data_masking_state: The state of the data masking policy.
        :param pulumi.Input[builtins.str] database_name: The name of the database.
        :param pulumi.Input[builtins.str] exempt_principals: The list of the exempt principals. Specifies the semicolon-separated list of database users for which the data masking policy does not apply. The specified users receive data results without masking for all of the database queries.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        :param pulumi.Input[builtins.str] server_name: The name of the server.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DataMaskingPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A database data masking policy.

        Uses Azure REST API version 2023-08-01. In version 2.x of the Azure Native provider, it used API version 2021-11-01.

        Other available API versions: 2014-04-01, 2021-11-01, 2022-02-01-preview, 2022-05-01-preview, 2022-08-01-preview, 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview, 2023-08-01-preview, 2024-05-01-preview, 2024-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param DataMaskingPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DataMaskingPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data_masking_policy_name: Optional[pulumi.Input[builtins.str]] = None,
                 data_masking_state: Optional[pulumi.Input['DataMaskingState']] = None,
                 database_name: Optional[pulumi.Input[builtins.str]] = None,
                 exempt_principals: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 server_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DataMaskingPolicyArgs.__new__(DataMaskingPolicyArgs)

            __props__.__dict__["data_masking_policy_name"] = data_masking_policy_name
            if data_masking_state is None and not opts.urn:
                raise TypeError("Missing required property 'data_masking_state'")
            __props__.__dict__["data_masking_state"] = data_masking_state
            if database_name is None and not opts.urn:
                raise TypeError("Missing required property 'database_name'")
            __props__.__dict__["database_name"] = database_name
            __props__.__dict__["exempt_principals"] = exempt_principals
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            if server_name is None and not opts.urn:
                raise TypeError("Missing required property 'server_name'")
            __props__.__dict__["server_name"] = server_name
            __props__.__dict__["application_principals"] = None
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["kind"] = None
            __props__.__dict__["location"] = None
            __props__.__dict__["masking_level"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:sql/v20140401:DataMaskingPolicy"), pulumi.Alias(type_="azure-native:sql/v20211101:DataMaskingPolicy"), pulumi.Alias(type_="azure-native:sql/v20220201preview:DataMaskingPolicy"), pulumi.Alias(type_="azure-native:sql/v20220501preview:DataMaskingPolicy"), pulumi.Alias(type_="azure-native:sql/v20220801preview:DataMaskingPolicy"), pulumi.Alias(type_="azure-native:sql/v20221101preview:DataMaskingPolicy"), pulumi.Alias(type_="azure-native:sql/v20230201preview:DataMaskingPolicy"), pulumi.Alias(type_="azure-native:sql/v20230501preview:DataMaskingPolicy"), pulumi.Alias(type_="azure-native:sql/v20230801:DataMaskingPolicy"), pulumi.Alias(type_="azure-native:sql/v20230801preview:DataMaskingPolicy"), pulumi.Alias(type_="azure-native:sql/v20240501preview:DataMaskingPolicy"), pulumi.Alias(type_="azure-native:sql/v20241101preview:DataMaskingPolicy")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(DataMaskingPolicy, __self__).__init__(
            'azure-native:sql:DataMaskingPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DataMaskingPolicy':
        """
        Get an existing DataMaskingPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DataMaskingPolicyArgs.__new__(DataMaskingPolicyArgs)

        __props__.__dict__["application_principals"] = None
        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["data_masking_state"] = None
        __props__.__dict__["exempt_principals"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["masking_level"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["type"] = None
        return DataMaskingPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationPrincipals")
    def application_principals(self) -> pulumi.Output[builtins.str]:
        """
        The list of the application principals. This is a legacy parameter and is no longer used.
        """
        return pulumi.get(self, "application_principals")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="dataMaskingState")
    def data_masking_state(self) -> pulumi.Output[builtins.str]:
        """
        The state of the data masking policy.
        """
        return pulumi.get(self, "data_masking_state")

    @property
    @pulumi.getter(name="exemptPrincipals")
    def exempt_principals(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The list of the exempt principals. Specifies the semicolon-separated list of database users for which the data masking policy does not apply. The specified users receive data results without masking for all of the database queries.
        """
        return pulumi.get(self, "exempt_principals")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[builtins.str]:
        """
        The kind of Data Masking Policy. Metadata, used for Azure portal.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[builtins.str]:
        """
        The location of the data masking policy.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="maskingLevel")
    def masking_level(self) -> pulumi.Output[builtins.str]:
        """
        The masking level. This is a legacy parameter and is no longer used.
        """
        return pulumi.get(self, "masking_level")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

