import pulumi
import pulumi_random as random
from pulumi_azurerm import resources, storage

random_string = random.RandomString("random_string", length=12, special=False, upper=False)

resource_group = resources.latest.ResourceGroup('resource_group', resource_group_name='azurerm-py', location='westus')

storage_account = storage.latest.StorageAccount('sa',
    account_name='pulumi143pysa',
    resource_group_name=resource_group.name,
    location='westus2',
    sku={
        'name': 'Standard_LRS',
        'tier': 'Standard',
    },
    kind='StorageV2')
