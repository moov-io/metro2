---
layout: page
title: Docker
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Docker

We publish a [public Docker image `moov/metro2`](https://hub.docker.com/r/moov/metro2/) on Docker Hub with each tagged release of Metro2. No configuration is required to serve on `:8080`. We also have Docker images for [OpenShift](https://quay.io/repository/moov/metro2?tab=tags) published as `quay.io/moov/metro2`.

Moov Metro2 is dependent on Docker being properly installed and running on your machine. Ensure that Docker is running. If your Docker client has issues connecting to the service, review the [Docker getting started guide](https://docs.docker.com/get-started/).

```
docker ps
```
```
CONTAINER ID        IMAGE        COMMAND        CREATED        STATUS        PORTS        NAMES
```

Pull & start the Docker image:
```
docker pull moov/metro2:latest
docker run -p 8080:8080 moov/metro2:latest
```

Upload a file and validate it:
```
curl -X POST --form "file=@./test/testdata/packed_file.json" http://localhost:8080/validator
```
```
valid file
```

Convert a file from JSON to raw Metro format:
```
curl -X POST --form "file=@./test/testdata/packed_file.json" --form "format=metro" http://localhost:8080/convert
```
```
0480HEADER                           555555555508202002051019990510199905101999YOUR BUSINESS NAME HERE                 LINE ONE OF YOUR ADDRESS LINE TWO OF YOUR ADDRESS LINE THERE OF YOUR ADDRESS                    1234567890
...
```

Convert a file from raw Metro format to JSON:
```
curl -X POST --form "file=@./test/testdata/packed_file.dat" --form "format=json" http://localhost:8080/convert
```
```
{"header": {"blockDescriptorWord": 370,"recordDescriptorWord": 366,"recordIdentifier": "HEADER","transUnionProgramIdentifier": "5555555555","activityDate": "2002-08-20T00:00:00Z", ...
```