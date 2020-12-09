# Build
1. Run the build.sh script to build both the Docker image and binaries for local tests.

2. to run the container
```
docker run -d -p 8080:8080 adisasson-api:latest
```

# Usage
You can now access from your browser to http://127.0.0.1:8080

Access http://127.0.0.1:8080/requests by GET method.

Access http://127.0.0.1:8080/drop by DELETE method.

Access http://127.0.0.1:8080/update?count=541 by POST method.
* You can pass any integer value in the "count" parameter.