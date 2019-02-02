# fake-server

A fake server for developing and testing client (web, mobile, console) applications.

## How to run?

- `go get github.com/prakashpandey/fake-server`

- `cd fake-server`

- `make run`

By default the server will start at `localhost:8284`

## Apis

- `/superAdmin`: Added dynamic endpoints with the desire response

    - Method: `POST`
    
    - Body: 
        ```
        {
            "path": "/cart",
            "method": "GET",
            "response": "I am cart endpoint"
        }	
        ```

- `/getEndpoints`: Get list of dynamically added endpoints

    - Example response: 
        ```
            ["","/cart","/user"]
        ```

### How to access your dynamically added Api

Go the your endpoint path, e.g if you have added endpoint `/cart`, then you can access your endpoint by going to
`localhost:8284/cart`
