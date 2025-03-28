# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from .. import _utilities
import typing
# Export this package's modules as members:
from ._enums import *
from .attached_network_by_dev_center import *
from .catalog import *
from .curation_profile import *
from .dev_box_definition import *
from .dev_center import *
from .encryption_set import *
from .environment_type import *
from .gallery import *
from .get_attached_network_by_dev_center import *
from .get_catalog import *
from .get_catalog_dev_box_definition_error_details import *
from .get_catalog_sync_error_details import *
from .get_curation_profile import *
from .get_customization_task_error_details import *
from .get_dev_box_definition import *
from .get_dev_center import *
from .get_encryption_set import *
from .get_environment_definition_error_details import *
from .get_environment_type import *
from .get_gallery import *
from .get_network_connection import *
from .get_plan import *
from .get_plan_member import *
from .get_pool import *
from .get_project import *
from .get_project_catalog import *
from .get_project_catalog_environment_definition_error_details import *
from .get_project_catalog_image_definition_build_details import *
from .get_project_catalog_sync_error_details import *
from .get_project_environment_type import *
from .get_project_inherited_settings import *
from .get_project_policy import *
from .get_schedule import *
from .list_skus_by_project import *
from .network_connection import *
from .plan import *
from .plan_member import *
from .pool import *
from .project import *
from .project_catalog import *
from .project_environment_type import *
from .project_policy import *
from .schedule import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_azure_native.devcenter.v20221111preview as __v20221111preview
    v20221111preview = __v20221111preview
    import pulumi_azure_native.devcenter.v20230401 as __v20230401
    v20230401 = __v20230401
    import pulumi_azure_native.devcenter.v20230801preview as __v20230801preview
    v20230801preview = __v20230801preview
    import pulumi_azure_native.devcenter.v20231001preview as __v20231001preview
    v20231001preview = __v20231001preview
    import pulumi_azure_native.devcenter.v20240201 as __v20240201
    v20240201 = __v20240201
    import pulumi_azure_native.devcenter.v20240501preview as __v20240501preview
    v20240501preview = __v20240501preview
    import pulumi_azure_native.devcenter.v20240601preview as __v20240601preview
    v20240601preview = __v20240601preview
    import pulumi_azure_native.devcenter.v20240701preview as __v20240701preview
    v20240701preview = __v20240701preview
    import pulumi_azure_native.devcenter.v20240801preview as __v20240801preview
    v20240801preview = __v20240801preview
    import pulumi_azure_native.devcenter.v20241001preview as __v20241001preview
    v20241001preview = __v20241001preview
    import pulumi_azure_native.devcenter.v20250201 as __v20250201
    v20250201 = __v20250201
else:
    v20221111preview = _utilities.lazy_import('pulumi_azure_native.devcenter.v20221111preview')
    v20230401 = _utilities.lazy_import('pulumi_azure_native.devcenter.v20230401')
    v20230801preview = _utilities.lazy_import('pulumi_azure_native.devcenter.v20230801preview')
    v20231001preview = _utilities.lazy_import('pulumi_azure_native.devcenter.v20231001preview')
    v20240201 = _utilities.lazy_import('pulumi_azure_native.devcenter.v20240201')
    v20240501preview = _utilities.lazy_import('pulumi_azure_native.devcenter.v20240501preview')
    v20240601preview = _utilities.lazy_import('pulumi_azure_native.devcenter.v20240601preview')
    v20240701preview = _utilities.lazy_import('pulumi_azure_native.devcenter.v20240701preview')
    v20240801preview = _utilities.lazy_import('pulumi_azure_native.devcenter.v20240801preview')
    v20241001preview = _utilities.lazy_import('pulumi_azure_native.devcenter.v20241001preview')
    v20250201 = _utilities.lazy_import('pulumi_azure_native.devcenter.v20250201')

