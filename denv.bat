@ECHO OFF
docker run --rm -v %cd%:/denv -w /denv -it %*
