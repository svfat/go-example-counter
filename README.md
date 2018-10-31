```
git clone git@github.com:svfat/go-example-counter.git
cd go-example-counter
docker build -t test .
docker run -p8080:8080 -n test test
```
