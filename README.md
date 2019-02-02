# fake-server

A fake server for developing and testing client (web, mobile, console) applications.

## How to run?

- `git clone github.com/prakashpandey/fake-server`

- `cd fake-server`

- `make run`

By default the server will start at `localhost:8284`

## Apis

- `/superAdmin`: Add dynamic endpoint with the desire response

    - Method: `POST`
    
    - Body: 
        ```
        {
            "endpoint": "/cart", // 'endpoint' is the path of your endpoint
            "method": "GET", // 'method' the http method this endpoint accept
            "response": "I am cart endpoint" // 'response' is the desire response you want when you will access this endpoint
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
