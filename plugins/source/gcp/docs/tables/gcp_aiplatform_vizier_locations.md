# Table: gcp_aiplatform_vizier_locations

https://cloud.google.com/api-gateway/docs/reference/rest/v1/projects.locations#Location

The composite primary key for this table is (**project_id**, **name**).

## Relations

The following tables depend on gcp_aiplatform_vizier_locations:
  - [gcp_aiplatform_studies](gcp_aiplatform_studies.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|location_id|String|
|display_name|String|
|labels|JSON|
|metadata|JSON|