## projects
{: #project-projects-cli}

Commands for Projects resource.

### `ibmcloud project list`
{: #project-cli-list-command}

List existing Projects. Projects are sorted by id.
Note: If the `--all-pages` option is not set, the command will only retrieve a single page of the collection.

```sh
ibmcloud project list [--start START] [--limit LIMIT] [--complete COMPLETE]
```


#### Command options
{: #project-list-cli-options}

`--start` (string)
:   Page token query parameter that is used to determine what resource to start the page after. If not specified, the logical first page is returned.

    The maximum length is `512` characters. The minimum length is `1` character.

`--limit` (int64)
:   Determine the maximum number of resources to return. The number of resources returned is the same, with exception of the last page.

    The maximum value is `100`. The minimum value is `1`.

`--complete` (bool)
:   The flag to tell if full metadata should be returned.

    The default value is `false`.

`--all-pages` (bool)
:   Invoke multiple requests to display all pages of the collection for list.

#### Example
{: #project-list-examples}

```sh
ibmcloud project list \
    --start=exampleString \
    --limit=10 \
    --complete=false
```
{: pre}

#### Example output
{: #project-list-cli-output}

A Projects list results example

```json
{
  "limit" : 10,
  "total_count" : 25,
  "first" : {
    "href" : "https://projects.test.cloud.ibm.com/v1/projects/cfbf9050-ab8e-ac97-b01b-ab5af830be8a",
    "start" : "start-here-for-next-page-dude"
  },
  "last" : {
    "href" : "https://projects.test.cloud.ibm.com/v1/projects/99999050-1234-ac97-0000-ba5a12fe0945"
  },
  "next" : {
    "href" : "https://projects.test.cloud.ibm.com/v1/projects/12349050-1234-ac97-0000-ba5a12fe9087"
  },
  "previous" : {
    "href" : "https://projects.test.cloud.ibm.com/v1/projects/12349000-6756-abcd-0000-ba5a12fe9087"
  },
  "projects" : [ {
    "description" : "a project example",
    "name" : "iaas-infra-prestage-env",
    "id" : "cfbf9050-ab8e-ac97-b01b-ab5af830be8a",
    "metadata" : {
      "crn" : "crn:v1:staging:public:project:us-south:a/06580c923e40314421d3b6cb40c01c68:cfbf9050-ab8e-ac97-b01b-ab5af830be8a::",
      "location" : "us-south",
      "resource_group" : "Default",
      "state" : "READY",
      "cumulative_needs_attention_view" : [ {
        "event" : "project.instance.update",
        "event_id" : "489f0090-6d7c-4af5-8f20-9106543e4974",
        "config_id" : "069ab83e-5016-4bf2-bd50-cc95cf678293",
        "config_version" : 1
      } ]
    }
  }, {
    "name" : "iaas-infra-stage-env",
    "id" : "1123ed42-4356-efa1-1101-235900fe9087",
    "metadata" : {
      "crn" : "crn:v1:staging:public:project:eu-de:a/06580d923e40314421d3b6cb40c01c68:cfbf9050-ab8e-ac97-b01b-ab5af830be8a::",
      "location" : "eu-gb",
      "resource_group" : "Default",
      "state" : "UPDATING",
      "cumulative_needs_attention_view" : [ ]
    }
  } ]
}
```
{: screen}

### `ibmcloud project get`
{: #project-cli-get-command}

Get a project definition document by id.

```sh
ibmcloud project get --id ID [--exclude-configs EXCLUDE-CONFIGS] [--complete COMPLETE]
```


#### Command options
{: #project-get-cli-options}

`--id` (string)
:   The id of the project, which uniquely identifies it. Required.

    The maximum length is `128` characters. The value must match regular expression `/^[\\-0-9a-z]+$/`.

`--exclude-configs` (bool)
:   Only return with the active configs, no drafts.

    The default value is `false`.

`--complete` (bool)
:   The flag to tell if full metadata should be returned.

    The default value is `false`.

#### Example
{: #project-get-examples}

```sh
ibmcloud project get \
    --id=exampleString \
    --exclude-configs=false \
    --complete=false
```
{: pre}

#### Example output
{: #project-get-cli-output}

Sample response for retrieving a project

```json
{
  "id" : "cfbf9050-ab8e-ac97-b01b-ab5af830be8a",
  "name" : "acme-microservice",
  "description" : "A microservice to deploy on top of ACME infrastructure",
  "configs" : [ {
    "id" : "673d79e4-52bf-4184-b8e9-d3ca3c110f96",
    "name" : "common-variables",
    "type" : "manual",
    "external_resources_account" : "e5ed08b9203bad3e4b6f57f0d1675a88",
    "output" : [ {
      "name" : "tags",
      "value" : [ "project:ghost", "type:infrastructure" ]
    } ]
  }, {
    "id" : "4a1d4ba2-54ba-43a7-975a-d82b5a7612d1",
    "name" : "account-stage",
    "description" : "Stage account configuration. The stage account hosts test environments prestage, performance, stage. This config configures services common to all these environments and regions. It is a terraform_template type of config that points to a github repo hosting the terraform modules that can be deployed via Schematics Workspace.\n",
    "locator_id" : "1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global",
    "type" : "terraform_template",
    "input" : [ {
      "name" : "account_id",
      "value" : "id of the stage account"
    }, {
      "name" : "resource_group",
      "value" : "cross-environments"
    }, {
      "name" : "access_tags",
      "value" : [ "account:stage" ]
    }, {
      "name" : "cis_name",
      "value" : "name of the CIS service to create"
    }, {
      "name" : "cm_name",
      "value" : "name of the Certiticate Manager serice to create"
    }, {
      "name" : "sm_name",
      "value" : "name of the Secrets Manager serice to create"
    } ],
    "output" : [ {
      "name" : "resource_group_id"
    }, {
      "name" : "cis_id"
    }, {
      "name" : "cm_id"
    }, {
      "name" : "sm_id"
    } ]
  }, {
    "id" : "293c3c36-a094-4115-a12b-de0a9ca39be5",
    "name" : "env-stage",
    "description" : "Stage environment configuration, which includes services common to all the environment regions. There must be a blueprint configuring all the services common to the stage regions. It is a terraform_template type of config that points to a github repo hosting the terraform modules that can be deployed via Schematics Workspace.\n",
    "labels" : [ "env:stage", "governance:test", "build:0" ],
    "locator_id" : "1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global",
    "type" : "terraform_template",
    "input" : [ {
      "name" : "account_id",
      "value" : "$configs[].name["account-stage"].input.account_id"
    }, {
      "name" : "resource_group",
      "value" : "stage"
    }, {
      "name" : "access_tags",
      "value" : [ "env:stage" ]
    }, {
      "name" : "logdna_name",
      "value" : "name of the LogDNA stage service instance"
    }, {
      "name" : "sysdig_name",
      "value" : "name of the SysDig stage service instance"
    } ],
    "output" : [ {
      "name" : "resource_group_id"
    }, {
      "name" : "logdna_id"
    }, {
      "name" : "sysdig_id"
    } ]
  }, {
    "id" : "596e8656-9d4b-41a5-8340-b0cbe8bd374a",
    "name" : "region-us-south-stage",
    "description" : "Stage us-south configuration. There must be a blueprint configuring the VPC + ROKS stage us-south. It is a schematics_blueprint type of config that points to a github repo hosting a Schematics Blueprint that can be deployed via Schematics Blueprint.\n",
    "locator_id" : "1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global",
    "type" : "schematics_blueprint",
    "labels" : [ "env:stage", "region:us-south", "governance:test", "build:0" ],
    "input" : [ {
      "name" : "account_id",
      "value" : "$configs[].name["account-stage"].input.account_id"
    }, {
      "name" : "resource_group_id",
      "value" : "$configs[].name["env-stage"].output.resource_group_id"
    }, {
      "name" : "logdna_id",
      "value" : "$configs[].name["env-stage"].output.logdna_id"
    }, {
      "name" : "sysdig_id",
      "value" : "$configs[].name["env-stage"].output.sysdig_id"
    }, {
      "name" : "access_tags",
      "value" : [ "region:us-south" ]
    } ],
    "output" : [ {
      "name" : "vpc_id"
    }, {
      "name" : "roks_cluster_id"
    } ]
  }, {
    "id" : "9c7afed6-17fb-4c56-a13d-440a78f936bd",
    "name" : "region-eu-de-stage",
    "description" : "Stage eu-de configuration. There must be a blueprint configuring the VPC + ROKS stage eu-de. It is a schematics_blueprint type of config that points to a github repo hosting a Schematics Blueprint that can be deployed via Schematics Blueprint.\n",
    "locator_id" : "1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global",
    "type" : "schematics_blueprint",
    "labels" : [ "env:stage", "region:eu-de", "governance:test", "build:0" ],
    "input" : [ {
      "name" : "account_id",
      "value" : "$configs[].name["account-stage"].input.account_id"
    }, {
      "name" : "resource_group_id",
      "value" : "$configs[].name["env-stage"].output.resource_group_id"
    }, {
      "name" : "logdna_id",
      "value" : "$configs[].name["env-stage"].output.logdna_id"
    }, {
      "name" : "sysdig_id",
      "value" : "$configs[].name["env-stage"].output.sysdig_id"
    }, {
      "name" : "access_tags",
      "value" : [ "region:eu-de" ]
    } ],
    "output" : [ {
      "name" : "vpc_id"
    }, {
      "name" : "roks_cluster_id"
    } ]
  } ],
  "metadata" : {
    "crn" : "crn:v1:staging:public:project:us-south:a/06580c923e40314421d3b6cb40c01c68:cfbf9050-ab8e-ac97-b01b-ab5af830be8a::",
    "location" : "us-south",
    "resource_group" : "Default",
    "state" : "READY",
    "cumulative_needs_attention_view" : [ {
      "event" : "project.instance.update",
      "event_id" : "489f0090-6d7c-4af5-8f20-9106543e4974",
      "config_id" : "069ab83e-5016-4bf2-bd50-cc95cf678293",
      "config_version" : 1
    } ],
    "event_notifications_crn" : "crn:v1:staging:public:event-notifications:us-south:a/06580c923e40314421d3b6cb40c01c68:instance-id::"
  }
}
```
{: screen}

### `ibmcloud project delete`
{: #project-cli-delete-command}

Delete a project document. A project can only be deleted after deleting all its artifacts.

```sh
ibmcloud project delete --id ID
```


#### Command options
{: #project-delete-cli-options}

`--id` (string)
:   The id of the project, which uniquely identifies it. Required.

    The maximum length is `128` characters. The value must match regular expression `/^[\\-0-9a-z]+$/`.

#### Example
{: #project-delete-examples}

```sh
ibmcloud project delete \
    --id=exampleString
```
{: pre}

## configs
{: #project-configs-cli}

Commands for Configs resource.

### `ibmcloud project configs`
{: #project-cli-configs-command}

Returns all project configs for a given project.

```sh
ibmcloud project configs --id ID [--version VERSION] [--complete COMPLETE]
```


#### Command options
{: #project-configs-cli-options}

`--id` (string)
:   The id of the project, which uniquely identifies it. Required.

    The maximum length is `128` characters. The value must match regular expression `/^[\\-0-9a-z]+$/`.

`--version` (string)
:   The version of configs to return.

    The default value is `active`. Allowable values are: `active`, `draft`, `mixed`.

`--complete` (bool)
:   The flag to tell if full metadata should be returned.

    The default value is `false`.

#### Example
{: #project-configs-examples}

```sh
ibmcloud project configs \
    --id=exampleString \
    --version=active \
    --complete=false
```
{: pre}

#### Example output
{: #project-configs-cli-output}

Sample response for get project configs

```json
{
  "configs" : [ {
    "id" : "293c3c36-a094-4115-a12b-de0a9ca39be5",
    "name" : "env-stage",
    "description" : "Stage environment configuration, which includes services common to all the environment regions. There must be a blueprint configuring all the services common to the stage regions. It is a terraform_template type of config that points to a github repo hosting the terraform modules that can be deployed via Schematics Workspace.\n",
    "labels" : [ "env:stage", "governance:test", "build:0" ],
    "locator_id" : "1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global",
    "type" : "terraform_template",
    "input" : [ {
      "name" : "account_id",
      "value" : "$configs[].name["account-stage"].input.account_id"
    }, {
      "name" : "resource_group",
      "value" : "stage"
    }, {
      "name" : "access_tags",
      "value" : [ "env:stage" ]
    }, {
      "name" : "logdna_name",
      "value" : "name of the LogDNA stage service instance"
    }, {
      "name" : "sysdig_name",
      "value" : "name of the SysDig stage service instance"
    } ],
    "output" : [ {
      "name" : "resource_group_id"
    }, {
      "name" : "logdna_id"
    }, {
      "name" : "sysdig_id"
    } ]
  }, {
    "id" : "9c7afed6-17fb-4c56-a13d-440a78f936bd",
    "name" : "region-eu-de-stage",
    "description" : "Stage eu-de configuration. There must be a blueprint configuring the VPC + ROKS stage eu-de. It is a schematics_blueprint type of config that points to a github repo hosting a Schematics Blueprint that can be deployed via Schematics Blueprint.\n",
    "locator_id" : "1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global",
    "type" : "schematics_blueprint",
    "labels" : [ "env:stage", "region:eu-de", "governance:test", "build:0" ],
    "input" : [ {
      "name" : "account_id",
      "value" : "$configs[].name["account-stage"].input.account_id"
    }, {
      "name" : "resource_group_id",
      "value" : "$configs[].name["env-stage"].output.resource_group_id"
    }, {
      "name" : "logdna_id",
      "value" : "$configs[].name["env-stage"].output.logdna_id"
    }, {
      "name" : "sysdig_id",
      "value" : "$configs[].name["env-stage"].output.sysdig_id"
    }, {
      "name" : "access_tags",
      "value" : [ "region:eu-de" ]
    } ],
    "output" : [ {
      "name" : "vpc_id"
    }, {
      "name" : "roks_cluster_id"
    } ]
  } ]
}
```
{: screen}

### `ibmcloud project config-operation`
{: #project-cli-config-operation-command}

Returns the specified project config in a given project.

```sh
ibmcloud project config-operation --id ID --config-id CONFIG-ID [--version VERSION] [--complete COMPLETE]
```


#### Command options
{: #project-config-operation-cli-options}

`--id` (string)
:   The id of the project, which uniquely identifies it. Required.

    The maximum length is `128` characters. The value must match regular expression `/^[\\-0-9a-z]+$/`.

`--config-id` (string)
:   The id of the config, which uniquely identifies it. Required.

    The maximum length is `128` characters. The value must match regular expression `/^[\\-0-9a-z]+$/`.

`--version` (string)
:   The version of the config to return.

    The default value is `active`. The maximum length is `6` characters. The minimum length is `1` character. The value must match regular expression `/^(active|draft|\\d+)$/`.

`--complete` (bool)
:   The flag to tell if full metadata should be returned.

    The default value is `false`.

#### Example
{: #project-config-operation-examples}

```sh
ibmcloud project config-operation \
    --id=exampleString \
    --config-id=exampleString \
    --version=active \
    --complete=false
```
{: pre}

#### Example output
{: #project-config-operation-cli-output}

Sample response for a project config

```json
{
  "id" : "293c3c36-a094-4115-a12b-de0a9ca39be5",
  "name" : "env-stage",
  "description" : "Stage environment configuration, which includes services common to all the environment regions. There must be a blueprint configuring all the services common to the stage regions. It is a terraform_template type of config that points to a github repo hosting the terraform modules that can be deployed via Schematics Workspace.\n",
  "labels" : [ "env:stage", "governance:test", "build:0" ],
  "locator_id" : "1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global",
  "type" : "terraform_template",
  "input" : [ {
    "name" : "account_id",
    "value" : "$configs[].name["account-stage"].input.account_id"
  }, {
    "name" : "resource_group",
    "value" : "stage"
  }, {
    "name" : "access_tags",
    "value" : [ "env:stage" ]
  }, {
    "name" : "logdna_name",
    "value" : "name of the LogDNA stage service instance"
  }, {
    "name" : "sysdig_name",
    "value" : "name of the SysDig stage service instance"
  } ],
  "output" : [ {
    "name" : "resource_group_id"
  }, {
    "name" : "logdna_id"
  }, {
    "name" : "sysdig_id"
  } ]
}
```
{: screen}

### `ibmcloud project delete_config`
{: #project-cli-delete_config-command}

Delete a config in a project. Deleting the config will also destroy all the resources deployed by the config.

```sh
ibmcloud project delete_config --id ID --config-id CONFIG-ID
```


#### Command options
{: #project-delete_config-cli-options}

`--id` (string)
:   The id of the project, which uniquely identifies it. Required.

    The maximum length is `128` characters. The value must match regular expression `/^[\\-0-9a-z]+$/`.

`--config-id` (string)
:   The id of the config, which uniquely identifies it. Required.

    The maximum length is `128` characters. The value must match regular expression `/^[\\-0-9a-z]+$/`.

#### Example
{: #project-delete_config-examples}

```sh
ibmcloud project delete_config \
    --id=exampleString \
    --config-id=exampleString
```
{: pre}

#### Example output
{: #project-delete_config-cli-output}

An example of the delete config response

```json
{
  "id" : "293c3c36-a094-4115-a12b-de0a9ca39be5",
  "name" : "env-stage"
}
```
{: screen}

### `ibmcloud project config-diff`
{: #project-cli-config-diff-command}

Returns a diff summary of the specified project config between its current draft and active version of a given project.

```sh
ibmcloud project config-diff --id ID --config-id CONFIG-ID
```


#### Command options
{: #project-config-diff-cli-options}

`--id` (string)
:   The id of the project, which uniquely identifies it. Required.

    The maximum length is `128` characters. The value must match regular expression `/^[\\-0-9a-z]+$/`.

`--config-id` (string)
:   The id of the config, which uniquely identifies it. Required.

    The maximum length is `128` characters. The value must match regular expression `/^[\\-0-9a-z]+$/`.

#### Example
{: #project-config-diff-examples}

```sh
ibmcloud project config-diff \
    --id=exampleString \
    --config-id=exampleString
```
{: pre}

#### Example output
{: #project-config-diff-cli-output}

Sample response for a project config diff summary

```json
{
  "added" : {
    "input" : [ {
      "name" : "account_id",
      "type" : "string",
      "new_value" : "5addba84e3f845be9acc27edee75a145"
    } ]
  },
  "changed" : {
    "input" : [ {
      "name" : "ibm_cloud_api_key",
      "type" : "string",
      "new_value" : "6c573253f969440ebd113fa3bf688809",
      "old_value" : "9f100bfc7ef245119581a4aa4b4dc53a"
    }, {
      "name" : "region",
      "type" : "string",
      "new_value" : "us-south",
      "old_value" : "us-east"
    } ]
  },
  "removed" : {
    "input" : [ {
      "name" : "env",
      "type" : "string",
      "old_value" : "stage"
    } ]
  }
}
```
{: screen}

### `ibmcloud project install`
{: #project-cli-install-command}

Install a project's configuration. It is an asynchronous operation that can be tracked using the project status api.

```sh
ibmcloud project install --id ID --config-id CONFIG-ID
```


#### Command options
{: #project-install-cli-options}

`--id` (string)
:   The id of the project, which uniquely identifies it. Required.

    The maximum length is `128` characters. The value must match regular expression `/^[\\-0-9a-z]+$/`.

`--config-id` (string)
:   The id of the config to install. Required.

    The maximum length is `128` characters. The value must match regular expression `/^[\\-0-9a-z]+$/`.

#### Example
{: #project-install-examples}

```sh
ibmcloud project install \
    --id=exampleString \
    --config-id=exampleString
```
{: pre}

### `ibmcloud project uninstall`
{: #project-cli-uninstall-command}

Uninstall a project's configuration. The operation uninstall all the resources deployed with the given configuration. You can track it by using the project status api.

```sh
ibmcloud project uninstall --id ID --config-id CONFIG-ID
```


#### Command options
{: #project-uninstall-cli-options}

`--id` (string)
:   The id of the project, which uniquely identifies it. Required.

    The maximum length is `128` characters. The value must match regular expression `/^[\\-0-9a-z]+$/`.

`--config-id` (string)
:   The id of the config to uninstall. Required.

    The maximum length is `128` characters. The value must match regular expression `/^[\\-0-9a-z]+$/`.

#### Example
{: #project-uninstall-examples}

```sh
ibmcloud project uninstall \
    --id=exampleString \
    --config-id=exampleString
```
{: pre}

## toolchain
{: #project-toolchain-cli}

Commands for Toolchain resource.

### `ibmcloud project get_schematics_job`
{: #project-cli-get_schematics_job-command}

Fetch and find the latest schematics job corresponds to a plan, install or uninstall action.

```sh
ibmcloud project get_schematics_job --id ID --config-id CONFIG-ID --action ACTION [--since SINCE]
```


#### Command options
{: #project-get_schematics_job-cli-options}

`--id` (string)
:   The id of the project, which uniquely identifies it. Required.

    The maximum length is `128` characters. The value must match regular expression `/^[\\-0-9a-z]+$/`.

`--config-id` (string)
:   The id of the config that triggered the action. Required.

    The maximum length is `128` characters. The value must match regular expression `/^[\\-0-9a-z]+$/`.

`--action` (string)
:   The triggered action. Required.

    Allowable values are: `plan`, `install`, `uninstall`.

`--since` (int64)
:   The timestamp of when the action was triggered.

#### Example
{: #project-get_schematics_job-examples}

```sh
ibmcloud project get_schematics_job \
    --id=exampleString \
    --config-id=exampleString \
    --action=plan \
    --since=38
```
{: pre}

#### Example output
{: #project-get_schematics_job-cli-output}

Sample response for retrieving the job of a project action

```json
{
  "id" : "345b03e7440fce352ffc652050c34dbf"
}
```
{: screen}

### `ibmcloud project get_cost_estimate`
{: #project-cli-get_cost_estimate-command}

Fetch the cost estimate for a given configuraton.

```sh
ibmcloud project get_cost_estimate --id ID --config-id CONFIG-ID [--version VERSION]
```


#### Command options
{: #project-get_cost_estimate-cli-options}

`--id` (string)
:   The id of the project, which uniquely identifies it. Required.

    The maximum length is `128` characters. The value must match regular expression `/^[\\-0-9a-z]+$/`.

`--config-id` (string)
:   The id of the config of the cost estimate to fetch. Required.

    The maximum length is `128` characters. The value must match regular expression `/^[\\-0-9a-z]+$/`.

`--version` (string)
:   The version of the config that the cost estimate will be fetched.

    The default value is `active`. The maximum length is `10` characters. The minimum length is `1` character. The value must match regular expression `/^(active|draft)$/`.

#### Example
{: #project-get_cost_estimate-examples}

```sh
ibmcloud project get_cost_estimate \
    --id=exampleString \
    --config-id=exampleString \
    --version=active
```
{: pre}

## event
{: #project-event-cli}

Commands for Event resource.

### `ibmcloud project post-notification`
{: #project-cli-post-notification-command}

Creates a notification event to be stored on the project definition.

```sh
ibmcloud project post-notification --id ID [--notifications NOTIFICATIONS]
```


#### Command options
{: #project-post-notification-cli-options}

`--id` (string)
:   The id of the project, which uniquely identifies it. Required.

    The maximum length is `128` characters. The value must match regular expression `/^[\\-0-9a-z]+$/`.

`--notifications` ([`NotificationEvent[]`](#cli-notification-event-example-schema))
:   &nbsp;

    The minimum length is `1` item.

#### Example
{: #project-post-notification-examples}

```sh
ibmcloud project post-notification \
    --id=exampleString \
    --notifications='[{"event": "project.create.failed", "target": "234234324-3444-4556-224232432", "source": "id.of.project.service.instance", "action_url": "url.for.project.documentation", "data": {"anyKey": "anyValue"}}]'
```
{: pre}

#### Example output
{: #project-post-notification-cli-output}

A post notifications response example

```json
[ {
  "_id" : "9n121be-3b7d-4bad-9bdd-2b0d7b3dcb6d",
  "event" : "project.create.failed",
  "target" : "234234324-3444-4556-224232432",
  "source" : "id.of.project.service.instance",
  "action_url" : "url.for.project.documentation",
  "data" : {
    "field1" : 1
  },
  "status" : "SUCCESS"
} ]
```
{: screen}

## healthcheck
{: #project-healthcheck-cli}

Commands for Healthcheck resource.

### `ibmcloud project health`
{: #project-cli-health-command}

```sh
ibmcloud project health [--info INFO]
```


#### Command options
{: #project-health-cli-options}

`--info` (bool)
:   Set this parameter if you want to get the version information in the output response.

    The default value is `false`.

#### Example
{: #project-health-examples}

```sh
ibmcloud project health \
    --info=false
```
{: pre}

## integration
{: #project-integration-cli}

Commands for Integration resource.

### `ibmcloud project post-event-notifications-integration`
{: #project-cli-post-event-notifications-integration-command}

connects a project instance to an event notifications instance.

```sh
ibmcloud project post-event-notifications-integration --id ID --instance-crn INSTANCE-CRN [--description DESCRIPTION] [--name NAME] [--enabled ENABLED] [--source SOURCE]
```


#### Command options
{: #project-post-event-notifications-integration-cli-options}

`--id` (string)
:   The id of the project, which uniquely identifies it. Required.

    The maximum length is `128` characters. The value must match regular expression `/^[\\-0-9a-z]+$/`.

`--instance-crn` (string)
:   &nbsp; Required.

    The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^.*$/`.

`--description` (string)
:   description of the instance of event.

    The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.

`--name` (string)
:   name of the instance of event.

    The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.

`--enabled` (bool)
:   status of instance of event.

`--source` (string)
:   source of instance of event.

    The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^.*$/`.

#### Example
{: #project-post-event-notifications-integration-examples}

```sh
ibmcloud project post-event-notifications-integration \
    --id=exampleString \
    --instance-crn='crn of event notifications instance' \
    --description='A sample project source' \
    --name='Project name' \
    --enabled=true \
    --source='CRN of the project instance'
```
{: pre}

#### Example output
{: #project-post-event-notifications-integration-cli-output}

A post integrations event notifications response example

```json
{
  "description" : "A sample project source",
  "enabled" : true,
  "id" : "CRN of the project instance",
  "name" : "Project name",
  "created_at" : "2017-10-10T01:22:38.665Z"
}
```
{: screen}

### `ibmcloud project post-event-notification`
{: #project-cli-post-event-notification-command}

sends notification to event notifications instance.

```sh
ibmcloud project post-event-notification --id ID --new-id NEW-ID --new-source NEW-SOURCE [--new-datacontenttype NEW-DATACONTENTTYPE] [--new-ibmendefaultlong NEW-IBMENDEFAULTLONG] [--new-ibmendefaultshort NEW-IBMENDEFAULTSHORT] [--new-ibmensourceid NEW-IBMENSOURCEID] [--new-specversion NEW-SPECVERSION] [--new-type NEW-TYPE]
```


#### Command options
{: #project-post-event-notification-cli-options}

`--id` (string)
:   The id of the project, which uniquely identifies it. Required.

    The maximum length is `128` characters. The value must match regular expression `/^[\\-0-9a-z]+$/`.

`--new-id` (string)
:   &nbsp; Required.

    The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^.*$/`.

`--new-source` (string)
:   source of instance of event. Required.

    The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.

`--new-datacontenttype` (string)
:   data content type of the instance of event.

    The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.

`--new-ibmendefaultlong` (string)
:   ibm default long message of the instance of event.

    The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.

`--new-ibmendefaultshort` (string)
:   ibm default short message of the instance of event.

    The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.

`--new-ibmensourceid` (string)
:   ibm source id of the instance of event.

    The maximum length is `512` characters. The minimum length is `9` characters. The value must match regular expression `/^.+$/`.

`--new-specversion` (string)
:   spec version of instance of event.

    The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.

`--new-type` (string)
:   type of instance of event.

    The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.

#### Example
{: #project-post-event-notification-examples}

```sh
ibmcloud project post-event-notification \
    --id=exampleString \
    --new-id=5f208fef-6b64-413c-aa07-dfed0b46abc1236 \
    --new-source='crn of project' \
    --new-datacontenttype=application/json \
    --new-ibmendefaultlong='long test notification message' \
    --new-ibmendefaultshort='Test notification' \
    --new-ibmensourceid='crn of project' \
    --new-specversion=1.0 \
    --new-type=com.ibm.cloud.project.project.test_notification
```
{: pre}

#### Example output
{: #project-post-event-notification-cli-output}

A post event notification response example

```json
{
  "datacontenttype" : "application/json",
  "ibmendefaultlong" : "long test notification message",
  "ibmendefaultshort" : "Test notification",
  "ibmensourceid" : "crn of project",
  "id" : "5f208fef-6b64-413c-aa07-dfed0b46abc1236",
  "notification_id" : "234234324-3444-4556-224232432",
  "source" : "crn of project",
  "specversion" : "1.0",
  "type" : "com.ibm.cloud.project.project.test_notification"
}
```
{: screen}

## Schema examples
{: #project-schema-examples}

The following schema examples represent the data that you need to specify for a command option. These examples model the data structure and include placeholder values for the expected value type. When you run a command, replace these values with the values that apply to your environment as appropriate.

### NotificationEvent[]
{: #cli-notification-event-example-schema}

The following example shows the format of the NotificationEvent[] object.

```json

[ {
  "event" : "project.create.failed",
  "target" : "234234324-3444-4556-224232432",
  "source" : "id.of.project.service.instance",
  "action_url" : "url.for.project.documentation",
  "data" : {
    "anyKey" : "anyValue"
  }
} ]
```
{: codeblock}
