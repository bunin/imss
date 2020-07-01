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
    - [google.protobuf.Timestamp](#google.protobuf.Timestamp)
  
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
| cloudId | [string](#string) |  | Image ID in the cloud |
| status | [UploadStatus](#imss.UploadStatus) |  | Upload status |
| error | [Error](#imss.Error) |  | Error message, if any |






<a name="imss.Session"></a>

### Session



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | Session ID |
| isActive | [bool](#bool) |  | Completion marker, only one active session at a time is allowed |
| createdAt | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Creation time |
| finishedAt | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Completion time |
| name | [string](#string) |  | Session name/title |
| images | [Image](#imss.Image) | repeated | List of related images (output only) |






<a name="imss.UploadJob"></a>

### UploadJob



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | Job ID |
| status | [UploadStatus](#imss.UploadStatus) |  | Upload status |
| createdAt | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Start time |
| finishedAt | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | End time |
| progress | [uint64](#uint64) |  | Uploaded bytes |
| size | [uint64](#uint64) |  | Total amount of bytes to upload |
| recipient | [string](#string) |  | User&#39;s email or other contact info |
| images | [ImageUpload](#imss.ImageUpload) | repeated | Images upload data |
| CloudLink | [string](#string) |  | Link to share with a recipient |
| error | [Error](#imss.Error) |  | Error message, if any |
| Name | [string](#string) |  | Album name (optional) |


<a name="imss.UploadStatus"></a>

### UploadStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN | 0 |  |
| IN_PROGRESS | 1 |  |
| DONE | 2 |  |
| ERROR | 3 |  |


<a name="google.protobuf.Timestamp"></a>

### Timestamp
A Timestamp represents a point in time independent of any time zone
or calendar, represented as seconds and fractions of seconds at
nanosecond resolution in UTC Epoch time. It is encoded using the
Proleptic Gregorian Calendar which extends the Gregorian calendar
backwards to year one. It is encoded assuming all minutes are 60
seconds long, i.e. leap seconds are &#34;smeared&#34; so that no leap second
table is needed for interpretation. Range is from
0001-01-01T00:00:00Z to 9999-12-31T23:59:59.999999999Z.
By restricting to that range, we ensure that we can convert to
and from  RFC 3339 date strings.
See [https://www.ietf.org/rfc/rfc3339.txt](https://www.ietf.org/rfc/rfc3339.txt).

#### Examples

Example 1: Compute Timestamp from POSIX `time()`.

    Timestamp timestamp;
    timestamp.set_seconds(time(NULL));
    timestamp.set_nanos(0);

Example 2: Compute Timestamp from POSIX `gettimeofday()`.

    struct timeval tv;
    gettimeofday(&amp;tv, NULL);

    Timestamp timestamp;
    timestamp.set_seconds(tv.tv_sec);
    timestamp.set_nanos(tv.tv_usec * 1000);

Example 3: Compute Timestamp from Win32 `GetSystemTimeAsFileTime()`.

    FILETIME ft;
    GetSystemTimeAsFileTime(&amp;ft);
    UINT64 ticks = (((UINT64)ft.dwHighDateTime) &lt;&lt; 32) | ft.dwLowDateTime;

    // A Windows tick is 100 nanoseconds. Windows epoch 1601-01-01T00:00:00Z
    // is 11644473600 seconds before Unix epoch 1970-01-01T00:00:00Z.
    Timestamp timestamp;
    timestamp.set_seconds((INT64) ((ticks / 10000000) - 11644473600LL));
    timestamp.set_nanos((INT32) ((ticks % 10000000) * 100));

Example 4: Compute Timestamp from Java `System.currentTimeMillis()`.

    long millis = System.currentTimeMillis();

    Timestamp timestamp = Timestamp.newBuilder().setSeconds(millis / 1000)
        .setNanos((int) ((millis % 1000) * 1000000)).build();


Example 5: Compute Timestamp from current time in Python.

    timestamp = Timestamp()
    timestamp.GetCurrentTime()

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| seconds | [int64](#int64) |  | Represents seconds of UTC time since Unix epoch 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59Z inclusive. |
| nanos | [int32](#int32) |  | Non-negative fractions of a second at nanosecond resolution. Negative second values with fractions must still have non-negative nanos values that count forward in time. Must be from 0 to 999,999,999 inclusive. |


 

 

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | --- | --- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

