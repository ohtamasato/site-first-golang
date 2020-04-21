# site-first-golang
It's my first experience with Golang.

## Installation for development
You need Docker to run / develop this app:

1. Visit https://hub.docker.com/editions/community/docker-ce-desktop-mac/
2. Get Docker Desktop
3. Install Docker Desktop
4. Check below:

```sh
$ docker version
```

Then you also need to run the below commands:

```sh
$ docker-compose up -d --build
```

Although, to avoid using Docker cache:

```sh
$ docker-compose build --no-cache && docker-compose up -d --no-build
```
