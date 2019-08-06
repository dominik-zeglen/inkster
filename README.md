# Inkster - Headless content management solution
==

Inkster is a open-source headless cms.

### Getting started
==

The most comfortable way to set up Inkster is by using [Docker](https://www.docker.com/) and [docker-compose](https://docs.docker.com/compose/). Simply run
```shell
$ docker-compose build
$ docker-compose run --rm web /app/migrate
$ docker-compose run --rm web /app/main add-user admin@example.com admin
$ docker-compose up
```

and then go to [localhost:8000/graphql](http://localhost:8000/graphql/) and play with the API.