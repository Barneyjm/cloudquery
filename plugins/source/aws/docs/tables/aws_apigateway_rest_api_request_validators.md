# Table: aws_apigateway_rest_api_request_validators

https://docs.aws.amazon.com/apigateway/latest/api/API_RequestValidator.html

The composite primary key for this table is (**account_id**, **arn**).

## Relations

This table depends on [aws_apigateway_rest_apis](aws_apigateway_rest_apis.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region|String|
|rest_api_arn|String|
|arn (PK)|String|
|id|String|
|name|String|
|validate_request_body|Bool|
|validate_request_parameters|Bool|