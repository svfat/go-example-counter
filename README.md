```
git clone git@github.com:svfat/go-example-counter.git
cd go-example-counter
docker build -t test .
docker run -p8080:8080 -n test test
ab -n 10000 -c 64 http://localhost:8080/
```
