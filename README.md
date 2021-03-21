# denv: Your Docker Environment

Use your working directory within a Docker image.

## Example

### Run

    $ cd /path
    $ denv run -i busybox echo Hello World!

Output:

    + docker run --rm -it -v /path:/denv/workdir -w /denv/workdir busybox echo Hello World!
    Hello World!

### Shell

    $ cd /path
    $ denv shell -i busybox

Output:

    + docker run --rm -it -v /path:/denv/workdir -w /denv/workdir busybox /bin/sh
    # ...
    # exit

### Config file

The config file `.denv.yml` is searched in:

- working directory `.`
- any parent of working directory
- home folder `~/.config/denv`

`.denv.yml` or `../.../.denv.yml` or `~/.config/denv/.denv.yml`:

    ---
    image: python:3.8

Skip image option `-i`:

    $ denv run echo 123
    $ denv shell
    $ denv

Explicit path to config file:

    $ denv -c my.denv.yml shell


### Named configuration

`.denv.yml`:

    ---
    image: python:3.8
    ---
    name: 3.9
    image: python:3.9

Use named configuration:

    $ denv run -n 3.9 echo 123
    $ denv shell -n 3.9


### Execute configuration

`.denv.yml`

    ---
    image: busybox
    exec:
      - echo build
    ---
    name: test
    image: python:3.8
    exec:
      - echo test

Execute:

    $ denv exec
    $ denv exec -n test

Execute all configurations:

    $ denv exec --all


## Development

Run:

    $ go run . <args>

Build:

    $ go build -o build/denv .
    $ build/denv <args>




# Notes

## build

    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o build/denv.exe ./denv

## Open questions

    - [ ] reuse container?
