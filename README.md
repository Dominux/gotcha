# gotcha

[<img alt="GitHub Workflow Status (with event)" src="https://img.shields.io/github/actions/workflow/status/Dominux/Pentaract/docker-image.yml?style=plastic&logo=github">](https://github.com/Dominux/gotcha/actions)
[<img alt="Dockerhub latest" src="https://img.shields.io/badge/dockerhub-latest-blue?logo=docker&style=plastic">](https://hub.docker.com/r/thedominux/gotcha)
[<img alt="Docker Image Size (tag)" src="https://img.shields.io/docker/image-size/thedominux/pentaract/latest?style=plastic&logo=docker&color=gold">](https://hub.docker.com/r/thedominux/gotcha/tags?page=1&name=latest)
[<img alt="Any platform" src="https://img.shields.io/badge/platform-any-green?style=plastic&logo=linux&logoColor=white">](https://github.com/Dominux/gotcha)

# Arch

![Gotcha service high-level arch](https://raw.githubusercontent.com/Dominux/gotcha/refs/heads/main/docs/gotcha_service.svg)

## Running this sht

No need to clone the proj, to download a go compiler or smth else. No even need to build the image on ur own - just take [the pre-built one](https://hub.docker.com/r/thedominux/gotcha)! Oh yeah, it's based on [the scratch image](https://hub.docker.com/_/scratch) so that's why it's 146% not bloat - just the runtime itself, so it will take as less disk space as possible (9.35MB total)

### 1. Declaring env vars

There are 2 ways to declare them, but we'll got with copypasting [the original .env.example file](https://github.com/Dominux/gotcha/blob/main/.env.example) into any directory with whatever name ya wish. **Don't forget to set proper values for them!!!**

### 2. Running the container itself

Example of such a command. Set right `--env-file` and `-p` params (inner port is 8000 and isn't meant to be changed anyhow - who cares, it's internal anyway)

```sh
docker run -d --env-file ~/gotcha/.gotcha.env -p 1337:8000 --name gotcha thedominux/gotcha
```
