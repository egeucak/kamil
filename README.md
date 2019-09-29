# Kamil API Gateway

Kamil is a dynamic API gateway to be used with alongside micro service architecture. All you need to do is set the `config.yaml` file. The good thing about Kamil is that you don't need to restart the gateway when you need to add another service.

## Example
A simple config file looks like this;

```yaml
routes:
 - route: /test1
   name: test1
   port: 8080
   request-types:
    - POST
   host: localhost
   
 - route: /test2/.*
   name: test2
   port: 8081
   host: localhost
```

 - `route` is the regular expression that matches the request path. Being able to use regular expressions for route matching gives the user great flexibility.

 - `name` field is mostly for the developer, it does not serve a goal for now.

 - `port` is the port that the target server is listening

 - `request-types` is an optional field for being able to control the requests more precisely.

 - `host` is the IP of the target server.

A POST request sent to the `${KAMIL_IP}/test1` would be sent to the `localhost:8080/test1`. If request type is not POST, an error would be returned.

## Program Arguments
 - `config-check-interval`: Interval time in seconds that the config file would be checked for updates. Default is 2 seconds
  
 - `port`: The port which the gateway will run. Default is 3000.
  
 - `config-file-name`: Path of the config file. Default is `./config.yaml`

## TODO
 - [ ] Add plugin support, and build some plugins for authantication.
 - [ ] Enable adding new services from an endpoint.
 - [ ] Create a User Interface
 - [ ] Add health checks
 - [ ] Write some tests
 - [ ] Create a docker version