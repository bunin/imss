# IMSS

[PRD](https://github.com/verevskoy/imss-notes/blob/master/SPECIFICATION.md) | [data types](./types.md) | [API](./API.md)

This software requires [Android Debug Bridge (adb)](https://developer.android.com/studio/command-line/adb) installed and accessible from `%PATH%`.
Also, a Google OAuth 2.0 app with access to these scopes `https://www.googleapis.com/auth/photoslibrary`, `https://www.googleapis.com/auth/photoslibrary.sharing` and `http://localhost:8080/auth` as an allowed redirect URL.

Configuration (see [.env.example](.env.example)):

|Name|Description|
|---|---|
|`IMSS_DB`| Local database file, default is `./imss.db`. |
|`IMSS_GOOGLE_ID`| Google OAuth 2.0 client ID ([docs](https://developers.google.com/photos/library/guides/get-started#enable-the-api)). Required. |
|`IMSS_GOOGLE_SECRET`|  Google OAuth 2.0 client Secret  ([docs](https://developers.google.com/photos/library/guides/get-started#enable-the-api)). Required. |
|`IMSS_LOG_LEVEL`| Log level. `debug` by default. Available levels are: `debug`, `info`, `warn`, `error`, `dpanic`, `panic`, `fatal`. |
|`IMSS_LOG`| Where to writes log, `./imss.log` by default. Can be a path to a file, `stdout` or `stderr`. |
|`IMSS_PORT`| Port to listen on. `8080` by default. |
|`LOCAL_DIR`| Where to store files locally, default is `C:\tmp`. |
|`PHONE_DIR`| Directory to scan for a new files on the phone. `/sdcard/DCIM/Camera/` by default. |
|`SMTP_FROM_EMAIL`| |
|`SMTP_FROM_NAME`| |
|`SMTP_HOST`| `smtp.gmail.com` by default. |
|`SMTP_LOGIN`| |
|`SMTP_PASSWORD`| |
|`SMTP_PORT`| `587` by default. |

