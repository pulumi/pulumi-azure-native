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

__all__ = ['DatabasePrincipalAssignmentArgs', 'DatabasePrincipalAssignment']

@pulumi.input_type
class DatabasePrincipalAssignmentArgs:
    def __init__(__self__, *,
                 cluster_name: pulumi.Input[builtins.str],
                 database_name: pulumi.Input[builtins.str],
                 principal_id: pulumi.Input[builtins.str],
                 principal_type: pulumi.Input[Union[builtins.str, 'PrincipalType']],
                 resource_group_name: pulumi.Input[builtins.str],
                 role: pulumi.Input[Union[builtins.str, 'DatabasePrincipalRole']],
                 principal_assignment_name: Optional[pulumi.Input[builtins.str]] = None,
                 tenant_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a DatabasePrincipalAssignment resource.
        :param pulumi.Input[builtins.str] cluster_name: The name of the Kusto cluster.
        :param pulumi.Input[builtins.str] database_name: The name of the database in the Kusto cluster.
        :param pulumi.Input[builtins.str] principal_id: The principal ID assigned to the database principal. It can be a user email, application ID, or security group name.
        :param pulumi.Input[Union[builtins.str, 'PrincipalType']] principal_type: Principal type.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[Union[builtins.str, 'DatabasePrincipalRole']] role: Database principal role.
        :param pulumi.Input[builtins.str] principal_assignment_name: The name of the Kusto principalAssignment.
        :param pulumi.Input[builtins.str] tenant_id: The tenant id of the principal
        """
        pulumi.set(__self__, "cluster_name", cluster_name)
        pulumi.set(__self__, "database_name", database_name)
        pulumi.set(__self__, "principal_id", principal_id)
        pulumi.set(__self__, "principal_type", principal_type)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "role", role)
        if principal_assignment_name is not None:
            pulumi.set(__self__, "principal_assignment_name", principal_assignment_name)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the Kusto cluster.
        """
        return pulumi.get(self, "cluster_name")

    @cluster_name.setter
    def cluster_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "cluster_name", value)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the database in the Kusto cluster.
        """
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "database_name", value)

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> pulumi.Input[builtins.str]:
        """
        The principal ID assigned to the database principal. It can be a user email, application ID, or security group name.
        """
        return pulumi.get(self, "principal_id")

    @principal_id.setter
    def principal_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "principal_id", value)

    @property
    @pulumi.getter(name="principalType")
    def principal_type(self) -> pulumi.Input[Union[builtins.str, 'PrincipalType']]:
        """
        Principal type.
        """
        return pulumi.get(self, "principal_type")

    @principal_type.setter
    def principal_type(self, value: pulumi.Input[Union[builtins.str, 'PrincipalType']]):
        pulumi.set(self, "principal_type", value)

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
    @pulumi.getter
    def role(self) -> pulumi.Input[Union[builtins.str, 'DatabasePrincipalRole']]:
        """
        Database principal role.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: pulumi.Input[Union[builtins.str, 'DatabasePrincipalRole']]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter(name="principalAssignmentName")
    def principal_assignment_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the Kusto principalAssignment.
        """
        return pulumi.get(self, "principal_assignment_name")

    @principal_assignment_name.setter
    def principal_assignment_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "principal_assignment_name", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The tenant id of the principal
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "tenant_id", value)


@pulumi.type_token("azure-native:kusto:DatabasePrincipalAssignment")
class DatabasePrincipalAssignment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_name: Optional[pulumi.Input[builtins.str]] = None,
                 database_name: Optional[pulumi.Input[builtins.str]] = None,
                 principal_assignment_name: Optional[pulumi.Input[builtins.str]] = None,
                 principal_id: Optional[pulumi.Input[builtins.str]] = None,
                 principal_type: Optional[pulumi.Input[Union[builtins.str, 'PrincipalType']]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 role: Optional[pulumi.Input[Union[builtins.str, 'DatabasePrincipalRole']]] = None,
                 tenant_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Class representing a database principal assignment.

        Uses Azure REST API version 2024-04-13. In version 2.x of the Azure Native provider, it used API version 2022-12-29.

        Other available API versions: 2019-11-09, 2020-02-15, 2020-06-14, 2020-09-18, 2021-01-01, 2021-08-27, 2022-02-01, 2022-07-07, 2022-11-11, 2022-12-29, 2023-05-02, 2023-08-15. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native kusto [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] cluster_name: The name of the Kusto cluster.
        :param pulumi.Input[builtins.str] database_name: The name of the database in the Kusto cluster.
        :param pulumi.Input[builtins.str] principal_assignment_name: The name of the Kusto principalAssignment.
        :param pulumi.Input[builtins.str] principal_id: The principal ID assigned to the database principal. It can be a user email, application ID, or security group name.
        :param pulumi.Input[Union[builtins.str, 'PrincipalType']] principal_type: Principal type.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[Union[builtins.str, 'DatabasePrincipalRole']] role: Database principal role.
        :param pulumi.Input[builtins.str] tenant_id: The tenant id of the principal
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DatabasePrincipalAssignmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Class representing a database principal assignment.

        Uses Azure REST API version 2024-04-13. In version 2.x of the Azure Native provider, it used API version 2022-12-29.

        Other available API versions: 2019-11-09, 2020-02-15, 2020-06-14, 2020-09-18, 2021-01-01, 2021-08-27, 2022-02-01, 2022-07-07, 2022-11-11, 2022-12-29, 2023-05-02, 2023-08-15. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native kusto [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param DatabasePrincipalAssignmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DatabasePrincipalAssignmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_name: Optional[pulumi.Input[builtins.str]] = None,
                 database_name: Optional[pulumi.Input[builtins.str]] = None,
                 principal_assignment_name: Optional[pulumi.Input[builtins.str]] = None,
                 principal_id: Optional[pulumi.Input[builtins.str]] = None,
                 principal_type: Optional[pulumi.Input[Union[builtins.str, 'PrincipalType']]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 role: Optional[pulumi.Input[Union[builtins.str, 'DatabasePrincipalRole']]] = None,
                 tenant_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DatabasePrincipalAssignmentArgs.__new__(DatabasePrincipalAssignmentArgs)

            if cluster_name is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_name'")
            __props__.__dict__["cluster_name"] = cluster_name
            if database_name is None and not opts.urn:
                raise TypeError("Missing required property 'database_name'")
            __props__.__dict__["database_name"] = database_name
            __props__.__dict__["principal_assignment_name"] = principal_assignment_name
            if principal_id is None and not opts.urn:
                raise TypeError("Missing required property 'principal_id'")
            __props__.__dict__["principal_id"] = principal_id
            if principal_type is None and not opts.urn:
                raise TypeError("Missing required property 'principal_type'")
            __props__.__dict__["principal_type"] = principal_type
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            if role is None and not opts.urn:
                raise TypeError("Missing required property 'role'")
            __props__.__dict__["role"] = role
            __props__.__dict__["tenant_id"] = tenant_id
            __props__.__dict__["aad_object_id"] = None
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["principal_name"] = None
            __props__.__dict__["provisioning_state"] = None
            __props__.__dict__["tenant_name"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:kusto/v20191109:DatabasePrincipalAssignment"), pulumi.Alias(type_="azure-native:kusto/v20200215:DatabasePrincipalAssignment"), pulumi.Alias(type_="azure-native:kusto/v20200614:DatabasePrincipalAssignment"), pulumi.Alias(type_="azure-native:kusto/v20200918:DatabasePrincipalAssignment"), pulumi.Alias(type_="azure-native:kusto/v20210101:DatabasePrincipalAssignment"), pulumi.Alias(type_="azure-native:kusto/v20210827:DatabasePrincipalAssignment"), pulumi.Alias(type_="azure-native:kusto/v20220201:DatabasePrincipalAssignment"), pulumi.Alias(type_="azure-native:kusto/v20220707:DatabasePrincipalAssignment"), pulumi.Alias(type_="azure-native:kusto/v20221111:DatabasePrincipalAssignment"), pulumi.Alias(type_="azure-native:kusto/v20221229:DatabasePrincipalAssignment"), pulumi.Alias(type_="azure-native:kusto/v20230502:DatabasePrincipalAssignment"), pulumi.Alias(type_="azure-native:kusto/v20230815:DatabasePrincipalAssignment"), pulumi.Alias(type_="azure-native:kusto/v20240413:DatabasePrincipalAssignment")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(DatabasePrincipalAssignment, __self__).__init__(
            'azure-native:kusto:DatabasePrincipalAssignment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DatabasePrincipalAssignment':
        """
        Get an existing DatabasePrincipalAssignment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DatabasePrincipalAssignmentArgs.__new__(DatabasePrincipalAssignmentArgs)

        __props__.__dict__["aad_object_id"] = None
        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["principal_id"] = None
        __props__.__dict__["principal_name"] = None
        __props__.__dict__["principal_type"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["role"] = None
        __props__.__dict__["tenant_id"] = None
        __props__.__dict__["tenant_name"] = None
        __props__.__dict__["type"] = None
        return DatabasePrincipalAssignment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="aadObjectId")
    def aad_object_id(self) -> pulumi.Output[builtins.str]:
        """
        The service principal object id in AAD (Azure active directory)
        """
        return pulumi.get(self, "aad_object_id")

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
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> pulumi.Output[builtins.str]:
        """
        The principal ID assigned to the database principal. It can be a user email, application ID, or security group name.
        """
        return pulumi.get(self, "principal_id")

    @property
    @pulumi.getter(name="principalName")
    def principal_name(self) -> pulumi.Output[builtins.str]:
        """
        The principal name
        """
        return pulumi.get(self, "principal_name")

    @property
    @pulumi.getter(name="principalType")
    def principal_type(self) -> pulumi.Output[builtins.str]:
        """
        Principal type.
        """
        return pulumi.get(self, "principal_type")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[builtins.str]:
        """
        The provisioned state of the resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[builtins.str]:
        """
        Database principal role.
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The tenant id of the principal
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter(name="tenantName")
    def tenant_name(self) -> pulumi.Output[builtins.str]:
        """
        The tenant name of the principal
        """
        return pulumi.get(self, "tenant_name")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

