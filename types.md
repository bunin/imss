# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [data/imss.proto](#data/imss.proto)
    - [Error](#imss.Error)
    - [Image](#imss.Image)
    - [ImageUpload](#imss.ImageUpload)
    - [Session](#imss.Session)
    - [UploadJob](#imss.UploadJob)
  
    - [UploadStatus](#imss.UploadStatus)
  
  
  

- [Scalar Value Types](#scalar-value-types)



<a name="data/imss.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## data/imss.proto



<a name="imss.Error"></a>

### Error



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [uint64](#uint64) |  | Error code |
| message | [string](#string) |  | Error message |






<a name="imss.Image"></a>

### Image



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | Image ID |
| sessionId | [string](#string) |  | Corresponding session ID |
| localPath | [string](#string) |  | Path to image file on the workstation |
| cloudId | [string](#string) |  | Remote path/ID of uploaded image |
| createdAt | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Creation time |
| size | [uint64](#uint64) |  | File size in bytes |






<a name="imss.ImageUpload"></a>

### ImageUpload



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | Image upload ID |
| jobId | [string](#string) |  | Related upload job ID |
| imageId | [string](#string) |  | ID of the image being uploaded |
| progress | [uint64](#uint64) |  | Uploaded bytes |
| status | [UploadStatus](#imss.UploadStatus) |  | Upload status |
| error | [Error](#imss.Error) |  | Error message, if any |






<a name="imss.Session"></a>

### Session



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | Session ID |
| completed | [bool](#bool) |  | Completion marker, only one active session at a time is allowed |
| createdAt | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Creation time |
| finishedAt | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Completion time |
| name | [string](#string) |  | Session name/title |
| images | [Image](#imss.Image) | repeated | List of related images |






<a name="imss.UploadJob"></a>

### UploadJob



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | Job ID |
| imageIds | [string](#string) | repeated | Related image IDs |
| status | [UploadStatus](#imss.UploadStatus) |  | Upload status |
| createdAt | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Start time |
| finishedAt | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | End time |
| progress | [uint64](#uint64) |  | Uploaded bytes |
| size | [uint64](#uint64) |  | Total amount of bytes to upload |
| recipient | [string](#string) |  | User&#39;s email or other contact |





 


<a name="imss.UploadStatus"></a>

### UploadStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN | 0 |  |
| IN_PROGRESS | 1 |  |
| DONE | 2 |  |
| ERROR | 3 |  |


 

 

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

