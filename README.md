# User API #

example user RESTAPI project

### How to run ###
```
cd my-project 
docker build -t user-api .
docker run --rm -p 8080:8080 user-api
```

### Project Structure ###
With RESTAPI project use 

- app
  - container.go : Initialize database or external service and Start application
  - handler.go : routing http route
- config
  - env.go : define env config
  - file.go : define file config
  - reader.go : initial and read config
- user
  - handler.go : Handle http request and response
  - service.go : Layer for facade design pattern for handle business requirement
  - repository.go : Layer for main database we can change database and query with not impact service layer
- config.yaml : config file
- Dockerfile : docker file for build image
- main.go : main golang file for load config and call start app
- package.json : control version with using lib standard verion https://github.com/conventional-changelog/standard-version