# denv: Your Docker Environment

Use your working directory within a Docker image.

## Example

Run:

    $ denv busybox echo Hello World!

Output:

    + docker run --rm -it -v /path:/denv/workdir -w /denv/workdir busybox echo Hello World!
    Hello World!

# Ideas

https://github.com/drone/drone-cli/blob/master/drone/main.go
https://github.com/urfave/cli/blob/master/docs/v2/manual.md#full-api-example

## todo

    - [x] setup go project
    - [ ] `run` command with `-i` option
    - [ ] `run` command with `- n` option, local `.denv.yml` + `image`/`name`/`args`
    - [ ] `shell` command

## build

    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o build/denv.exe ./denv

## Docker approach

    docker build -t fipsi .
    docker run --rm -v %cd%:/denv -w /denv fipsi pipenv --help

## Syntax

    denv <opts> <cmd> <name> <args>

## run a single command

    denv -i python:3.8 run /bin/sh
    denv -i python:3.8 run pipenv install --dev

## open shell

    denv -i python:3.8 run /bin/sh
    denv -i python:3.8 shell

## options

image name:

    denv -i python shell

name from `.denv.yml`:

    denv -n python shell

name is 'default' or '' from `.denv.yml`:

    denv shell
    denv

## config file

Priorities:

    - `./.denv.yml`
    - `~/.denv.yml`

Use specific file:

    denv -f .denv.custom.yml ...

## execute script

run commands of `name: default`:

    denv exec

run commands of `name: test`:

    denv -n test exec

run commands of multiple configurations:

    denv -n test -n build exec

run specific script:

    denv exec scripts/execute.sh

https://github.com/drone/drone-yaml/blob/master/yaml/compiler/script_posix.go

## Open questions

    - [ ] reuse container?
