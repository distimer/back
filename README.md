# DISTIMER Backend

디스타이머 프로젝트 백엔드 리포지토리

[![wakatime](https://wakatime.com/badge/user/aeebb3a2-8786-4794-9ad8-bd3812263c99/project/f98bd272-3c29-4214-83a9-b9a583500c5c.svg)](https://wakatime.com/badge/user/aeebb3a2-8786-4794-9ad8-bd3812263c99/project/f98bd272-3c29-4214-83a9-b9a583500c5c) [![Docker](https://github.com/distimer/back/actions/workflows/docker-publish.yml/badge.svg)](https://github.com/distimer/back/actions/workflows/docker-publish.yml)

## Usage

```yml
version: "3"

services:
  distimer-back:
    image: ghcr.io/distimer/distimer-back:main
    restart: always
    expose:
      - 3000
    environment:
      DOPPLER_TOKEN: ${DOPPLER_TOKEN}
```
