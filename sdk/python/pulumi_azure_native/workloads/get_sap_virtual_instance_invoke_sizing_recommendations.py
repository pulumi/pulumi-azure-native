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

__all__ = [
    'GetSapVirtualInstanceInvokeSizingRecommendationsResult',
    'AwaitableGetSapVirtualInstanceInvokeSizingRecommendationsResult',
    'get_sap_virtual_instance_invoke_sizing_recommendations',
    'get_sap_virtual_instance_invoke_sizing_recommendations_output',
]

@pulumi.output_type
class GetSapVirtualInstanceInvokeSizingRecommendationsResult:
    """
    The SAP sizing recommendation result.
    """
    def __init__(__self__, deployment_type=None):
        if deployment_type and not isinstance(deployment_type, str):
            raise TypeError("Expected argument 'deployment_type' to be a str")
        pulumi.set(__self__, "deployment_type", deployment_type)

    @property
    @pulumi.getter(name="deploymentType")
    def deployment_type(self) -> builtins.str:
        """
        The deployment type. Eg: SingleServer/ThreeTier
        """
        return pulumi.get(self, "deployment_type")


class AwaitableGetSapVirtualInstanceInvokeSizingRecommendationsResult(GetSapVirtualInstanceInvokeSizingRecommendationsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSapVirtualInstanceInvokeSizingRecommendationsResult(
            deployment_type=self.deployment_type)


def get_sap_virtual_instance_invoke_sizing_recommendations(app_location: Optional[builtins.str] = None,
                                                           database_type: Optional[Union[builtins.str, 'SAPDatabaseType']] = None,
                                                           db_memory: Optional[builtins.float] = None,
                                                           db_scale_method: Optional[Union[builtins.str, 'SAPDatabaseScaleMethod']] = None,
                                                           deployment_type: Optional[Union[builtins.str, 'SAPDeploymentType']] = None,
                                                           environment: Optional[Union[builtins.str, 'SAPEnvironmentType']] = None,
                                                           high_availability_type: Optional[Union[builtins.str, 'SAPHighAvailabilityType']] = None,
                                                           location: Optional[builtins.str] = None,
                                                           sap_product: Optional[Union[builtins.str, 'SAPProductType']] = None,
                                                           saps: Optional[builtins.float] = None,
                                                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSapVirtualInstanceInvokeSizingRecommendationsResult:
    """
    Gets the sizing recommendations.

    Uses Azure REST API version 2024-09-01.


    :param builtins.str app_location: The geo-location where the resource is to be created.
    :param Union[builtins.str, 'SAPDatabaseType'] database_type: The database type.
    :param builtins.float db_memory: The database memory configuration.
    :param Union[builtins.str, 'SAPDatabaseScaleMethod'] db_scale_method: The DB scale method.
    :param Union[builtins.str, 'SAPDeploymentType'] deployment_type: The deployment type. Eg: SingleServer/ThreeTier
    :param Union[builtins.str, 'SAPEnvironmentType'] environment: Defines the environment type - Production/Non Production.
    :param Union[builtins.str, 'SAPHighAvailabilityType'] high_availability_type: The high availability type.
    :param builtins.str location: The name of the Azure region.
    :param Union[builtins.str, 'SAPProductType'] sap_product: Defines the SAP Product type.
    :param builtins.float saps: The SAP Application Performance Standard measurement.
    """
    __args__ = dict()
    __args__['appLocation'] = app_location
    __args__['databaseType'] = database_type
    __args__['dbMemory'] = db_memory
    __args__['dbScaleMethod'] = db_scale_method
    __args__['deploymentType'] = deployment_type
    __args__['environment'] = environment
    __args__['highAvailabilityType'] = high_availability_type
    __args__['location'] = location
    __args__['sapProduct'] = sap_product
    __args__['saps'] = saps
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:workloads:getSapVirtualInstanceInvokeSizingRecommendations', __args__, opts=opts, typ=GetSapVirtualInstanceInvokeSizingRecommendationsResult).value

    return AwaitableGetSapVirtualInstanceInvokeSizingRecommendationsResult(
        deployment_type=pulumi.get(__ret__, 'deployment_type'))
def get_sap_virtual_instance_invoke_sizing_recommendations_output(app_location: Optional[pulumi.Input[builtins.str]] = None,
                                                                  database_type: Optional[pulumi.Input[Union[builtins.str, 'SAPDatabaseType']]] = None,
                                                                  db_memory: Optional[pulumi.Input[builtins.float]] = None,
                                                                  db_scale_method: Optional[pulumi.Input[Optional[Union[builtins.str, 'SAPDatabaseScaleMethod']]]] = None,
                                                                  deployment_type: Optional[pulumi.Input[Union[builtins.str, 'SAPDeploymentType']]] = None,
                                                                  environment: Optional[pulumi.Input[Union[builtins.str, 'SAPEnvironmentType']]] = None,
                                                                  high_availability_type: Optional[pulumi.Input[Optional[Union[builtins.str, 'SAPHighAvailabilityType']]]] = None,
                                                                  location: Optional[pulumi.Input[builtins.str]] = None,
                                                                  sap_product: Optional[pulumi.Input[Union[builtins.str, 'SAPProductType']]] = None,
                                                                  saps: Optional[pulumi.Input[builtins.float]] = None,
                                                                  opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetSapVirtualInstanceInvokeSizingRecommendationsResult]:
    """
    Gets the sizing recommendations.

    Uses Azure REST API version 2024-09-01.


    :param builtins.str app_location: The geo-location where the resource is to be created.
    :param Union[builtins.str, 'SAPDatabaseType'] database_type: The database type.
    :param builtins.float db_memory: The database memory configuration.
    :param Union[builtins.str, 'SAPDatabaseScaleMethod'] db_scale_method: The DB scale method.
    :param Union[builtins.str, 'SAPDeploymentType'] deployment_type: The deployment type. Eg: SingleServer/ThreeTier
    :param Union[builtins.str, 'SAPEnvironmentType'] environment: Defines the environment type - Production/Non Production.
    :param Union[builtins.str, 'SAPHighAvailabilityType'] high_availability_type: The high availability type.
    :param builtins.str location: The name of the Azure region.
    :param Union[builtins.str, 'SAPProductType'] sap_product: Defines the SAP Product type.
    :param builtins.float saps: The SAP Application Performance Standard measurement.
    """
    __args__ = dict()
    __args__['appLocation'] = app_location
    __args__['databaseType'] = database_type
    __args__['dbMemory'] = db_memory
    __args__['dbScaleMethod'] = db_scale_method
    __args__['deploymentType'] = deployment_type
    __args__['environment'] = environment
    __args__['highAvailabilityType'] = high_availability_type
    __args__['location'] = location
    __args__['sapProduct'] = sap_product
    __args__['saps'] = saps
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:workloads:getSapVirtualInstanceInvokeSizingRecommendations', __args__, opts=opts, typ=GetSapVirtualInstanceInvokeSizingRecommendationsResult)
    return __ret__.apply(lambda __response__: GetSapVirtualInstanceInvokeSizingRecommendationsResult(
        deployment_type=pulumi.get(__response__, 'deployment_type')))
