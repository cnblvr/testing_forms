```bash
docker build -t testing_forms:v1 .
docker create -p 8080:8080 --name testing_forms testing_forms:v1
docker start testing_forms -a
```
t
