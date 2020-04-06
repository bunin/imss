# IMSS

[PRD](https://github.com/verevskoy/imss-notes/blob/master/SPECIFICATION.md) | [data types](./types.md)

Env parameters (see [.env.example](.env.example)):

|Name|Description|
|---|---|
|`IMSS_DIR`| Directory to watch, required, default is `.` (current directory). |
|`IMSS_KEY`| S3 Access key. Required. |
|`IMSS_SECRET`| S3 Secret key. Required. |
|`IMSS_BUCKET`| S3 bucket name. Required, default is `pics`. |