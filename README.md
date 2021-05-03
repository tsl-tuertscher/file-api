![example workflow](https://github.com/tsl-tuertscher/file-api/actions/workflows/main.yml/badge.svg)
# file-api

## Description

File api for Docker container

## Usage

### Get tile
```
curl localhost:8080/tiles/base/180/9/137.hgt?key=23
```

### Post data
```
curl localhost:8080/tiles/base -X POST -H "key: asdf" -H "Content-Type: application/json" --data "{\"url\": \"https://onedrive.live.com/"}"
```

## Config
```
{
  "workingDirectory": "workdir",
  "key":[
    "pwieub3289h"
  ],
  "offset":"work"
}

```
