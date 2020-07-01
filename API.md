# API methods

Table of Contents
=================

* [GET /auth](#get-auth)
* [GET|HEAD /files/:localPath](#gethead-fileslocalpath)
* [GET /api/auth/check](#get-apiauthcheck)
* [GET /api/session](#get-apisession)
* [GET /api/session/:id](#get-apisessionid)
* [PATCH /api/session/:id](#patch-apisessionid)
* [POST /api/session](#post-apisession)
* [POST /api/upload](#post-apiupload)

## GET /auth

Redirect target from Google's OAuth request, reads the code and saves the token

## GET|HEAD /files/:localPath

Returns saved [image](./types.md#imss.Image)'s contents by its `localPath`

## GET /api/auth/check

Checks if the Google OAuth token is saved and valid, returns an empty object on success (`{}`) or `{"url": "https://google...."}` where the user should be redirected to in order to get a new token.

## GET /api/session

Returns an array of up to 20 latest [sessions](./types.md#imss.Session) sorted from newer to older.
Or an error with `500` status.

Response example:

```json
[
    {
        "id": "brqr7ipjpb3jqs7hnbcg",
        "isActive": true,
        "createdAt": {
            "seconds": 1593160651,
            "nanos": 499095300
        },
        "images": [
            {
                "id": "brqr7ipjpb3jqs7hnbcg:brs5fj1jpb3jmu66lc30",
                "sessionId": "brqr7ipjpb3jqs7hnbcg",
                "localPath": "C:%5C%20tmp%5CIMG_20200628_154144.jpg",
                "createdAt": {
                    "seconds": 1593333708,
                    "nanos": 600069300
                },
                "size": 4461125
            },
            {
                "id": "brqr7ipjpb3jqs7hnbcg:brs5ge1jpb3jmu66lc40",
                "sessionId": "brqr7ipjpb3jqs7hnbcg",
                "localPath": "C:%5C%20tmp%5CIMG_20200628_154332.jpg",
                "createdAt": {
                    "seconds": 1593333816,
                    "nanos": 669570500
                },
                "size": 5196783
            },
            {
                "id": "brqr7ipjpb3jqs7hnbcg:brs5vrpjpb3gne1rbm0g",
                "sessionId": "brqr7ipjpb3jqs7hnbcg",
                "localPath": "C:%5Ctmp%5CIMG_20200628_161627.jpg",
                "createdAt": {
                    "seconds": 1593335791,
                    "nanos": 825199500
                },
                "size": 4573970
            },
            {
                "id": "brqr7ipjpb3jqs7hnbcg:brs5vshjpb3gne1rbm1g",
                "sessionId": "brqr7ipjpb3jqs7hnbcg",
                "localPath": "C:%5Ctmp%5CVID_20200628_161633.mp4",
                "createdAt": {
                    "seconds": 1593335794,
                    "nanos": 677720800
                }
            },
            {
                "id": "brqr7ipjpb3jqs7hnbcg:brs5vtpjpb3gne1rbm2g",
                "sessionId": "brqr7ipjpb3jqs7hnbcg",
                "localPath": "C:%5Ctmp%5CVID_20200628_161637.mp4",
                "createdAt": {
                    "seconds": 1593335799,
                    "nanos": 874471000
                },
                "size": 6935762
            },
            {
                "id": "brqr7ipjpb3jqs7hnbcg:brs75bhjpb3gbh0tfrbg",
                "sessionId": "brqr7ipjpb3jqs7hnbcg",
                "localPath": "C:%5Ctmp%5CIMG_20200628_173622.jpg",
                "createdAt": {
                    "seconds": 1593340590,
                    "nanos": 291296300
                },
                "size": 4222656
            }
        ]
    }
]
```

## GET /api/session/:id

Returns a single [session](./types.md#imss.Session) by its ID.

Or `404` error if session was not found.

Or `500` error in case of any other error.

Response example:

```json
{
    "id": "brqr7ipjpb3jqs7hnbcg",
    "isActive": true,
    "createdAt": {
        "seconds": 1593160651,
        "nanos": 499095300
    },
    "images": [
        {
            "id": "brqr7ipjpb3jqs7hnbcg:brs5fj1jpb3jmu66lc30",
            "sessionId": "brqr7ipjpb3jqs7hnbcg",
            "localPath": "C:%5C%20tmp%5CIMG_20200628_154144.jpg",
            "createdAt": {
                "seconds": 1593333708,
                "nanos": 600069300
            },
            "size": 4461125
        },
        {
            "id": "brqr7ipjpb3jqs7hnbcg:brs5ge1jpb3jmu66lc40",
            "sessionId": "brqr7ipjpb3jqs7hnbcg",
            "localPath": "C:%5C%20tmp%5CIMG_20200628_154332.jpg",
            "createdAt": {
                "seconds": 1593333816,
                "nanos": 669570500
            },
            "size": 5196783
        },
        {
            "id": "brqr7ipjpb3jqs7hnbcg:brs5vrpjpb3gne1rbm0g",
            "sessionId": "brqr7ipjpb3jqs7hnbcg",
            "localPath": "C:%5Ctmp%5CIMG_20200628_161627.jpg",
            "createdAt": {
                "seconds": 1593335791,
                "nanos": 825199500
            },
            "size": 4573970
        },
        {
            "id": "brqr7ipjpb3jqs7hnbcg:brs5vshjpb3gne1rbm1g",
            "sessionId": "brqr7ipjpb3jqs7hnbcg",
            "localPath": "C:%5Ctmp%5CVID_20200628_161633.mp4",
            "createdAt": {
                "seconds": 1593335794,
                "nanos": 677720800
            }
        },
        {
            "id": "brqr7ipjpb3jqs7hnbcg:brs5vtpjpb3gne1rbm2g",
            "sessionId": "brqr7ipjpb3jqs7hnbcg",
            "localPath": "C:%5Ctmp%5CVID_20200628_161637.mp4",
            "createdAt": {
                "seconds": 1593335799,
                "nanos": 874471000
            },
            "size": 6935762
        },
        {
            "id": "brqr7ipjpb3jqs7hnbcg:brs75bhjpb3gbh0tfrbg",
            "sessionId": "brqr7ipjpb3jqs7hnbcg",
            "localPath": "C:%5Ctmp%5CIMG_20200628_173622.jpg",
            "createdAt": {
                "seconds": 1593340590,
                "nanos": 291296300
            },
            "size": 4222656
        }
    ]
}
```

## PATCH /api/session/:id

Updates a [session](./types.md#imss.Session) by its ID.

Response codes:

`200` - session was successfully updated. Response body will contain an updated [session](./types.md#imss.Session) object (only `isActive` value can be changed).

`302` - request contains `{"isActive": true}` but there is another active session. `Location` header will contain the current active session URL (eg. `/api/session/brqr7ipjpb3jqs7hnbcg`)

`400` - request body could not be unmarshaled into a [session](./types.md#imss.Session) object.

`404` - session not found.

`500` - any other error.

## POST /api/session

Creates a new active session. Request body should be a [session](./types.md#imss.Session) object.

Response codes:

`201` - session was successfully created. `Location` header will contain a new session URL (eg. `/api/session/brqr7ipjpb3jqs7hnbcg`), and the body will be a new [session](./types.md#imss.Session) object.

`302` - there is another active session. `Location` header will contain the current active session URL (eg. `/api/session/brqr7ipjpb3jqs7hnbcg`)

`500` - any other error.

## POST /api/upload

Creates and starts an [upload job](./types.md#imss.UploadJob).

Request example:

```json
{
    "images": [
        {"id": "brqr7ipjpb3jqs7hnbcg:brs75bhjpb3gbh0tfrbg"},
        {"id": "brqr7ipjpb3jqs7hnbcg:brs5vtpjpb3gne1rbm2g"}
    ],
    "recipient": "user@example.com"
}
```

Response codes:

`201` - job was successfully created.

`400` - invalid request.

`500` - any other error.

Response example:
```json
{
    "id": "brudva9jpb3ivb35po1g",
    "status": 1,
    "createdAt": {
        "seconds": 1593630633,
        "nanos": 920027500
    },
    "size": 11158418,
    "recipient": "user@example.com",
    "images": [
        {
            "id": "brudva9jpb3ivb35po1g:brudva9jpb3ivb35po20",
            "jobId": "brudva9jpb3ivb35po1g",
            "imageId": "brqr7ipjpb3jqs7hnbcg:brs75bhjpb3gbh0tfrbg"
        },
        {
            "id": "brudva9jpb3ivb35po1g:brudva9jpb3ivb35po2g",
            "jobId": "brudva9jpb3ivb35po1g",
            "imageId": "brqr7ipjpb3jqs7hnbcg:brs5vtpjpb3gne1rbm2g"
        }
    ]
}
```

