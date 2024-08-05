# Simple Golang web server image with different output for uris

Build the image:
```
# build image and push:

$ docker buildx build --push -t grahamh/hello-suse:1.0 .

```
  
Start a local container instance to test:
```
$ docker run -d --rm -p 8080:8080 grahamh/hello-suse:1.0

```

Test:
Note: includes hostname so can test replica sets hit different instances
```
$ curl localhost:8080

 _          _ _
| |__   ___| | | ___
| '_ \ / _ \ | |/ _ \
| | | |  __/ | | (_) |
|_| |_|\___|_|_|\___/

[01-17-2024 14:46:16.49] Container hostname: b5c935d605cc
```

```
$ curl localhost:8080/suse

  ____  _   _ ____  _____
 / ___|| | | / ___|| ____|
 \___ \| | | \___ \|  _|
  ___) | |_| |___) | |___
 |____/ \___/|____/|_____|

[08-05-2024 09:17:38.69] Container hostname: 0cc85b5b976a
```
