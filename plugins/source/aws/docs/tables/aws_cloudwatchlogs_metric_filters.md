# Table: aws_cloudwatchlogs_metric_filters

https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_MetricFilter.html

The composite primary key for this table is (**log_group_arn**, **filter_name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|log_group_arn (PK)|String|
|creation_time|Int|
|filter_name (PK)|String|
|filter_pattern|String|
|log_group_name|String|
|metric_transformations|JSON|