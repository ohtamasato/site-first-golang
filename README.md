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

To work in the container:

```sh
$ docker exec -it go_app bash
```

Then to quit from the container:

```sh
# exit
```

To delete the containers:

```sh
$ docker stop go_app && docker stop go_db && docker rm go_app && docker rm go_db
```

## Customized behaviour
As commit `#830fcb6` statemented, empty slice should return `[]` instead of `null` in JSON.
Generated controllers by `bee generate scaffold <model>` originally returns `null` though,
it's not normal to popular JSON. Please add the below code to every generated controller:

```go
func (c *UserController) GetAll() {
	// var fields []string....
	
	if err != nil {
		c.Data["json"] = err.Error()
	} else if l == nil {
		c.Data["json"] = make([]models.User, 0) // <--- Add this one
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}
```
