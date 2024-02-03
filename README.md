# Gokey: Distributed key-value store

**Usage:**

* **Application Prerequisites:**
  - Make sure go is installed version >= 1.21
  - Clone the repository

* **Application Startup:**
  - Execute the following commands to initiate the nodes:
    ```bash
    go run cmd/httpd/httpd.go -node-id A -address localhost:8000 -http-port 9000 -bootstrap
    ```
    ```bash
    go run cmd/httpd/httpd.go -node-id B -address localhost:8001 -http-port 9001
    ```
    ```bash
    go run cmd/httpd/httpd.go -node-id C -address localhost:8002 -http-port 9002

  After successfully launching the nodes, ensure that the leader node (i.e., `node A`, as it was designated with the bootstrap option) adds the other two nodes to the cluster.

* **Adding Nodes API:**
  - Add node B to the cluster:
    ```bash
    curl --location 'localhost:9000/add-replica' \
    --header 'Content-Type: application/json' \
    --data '{
        "node_id": "B",
        "address": "localhost:8001"
    }'
    ```
  - Add node C to the cluster:
    ```bash
    curl --location 'localhost:9000/add-replica' \
    --header 'Content-Type: application/json' \
    --data '{
        "node_id": "C",
        "address": "localhost:8002"
    }'
    ```

- **Examples:**
  1. **SET:**
       - **Request:**
           ```bash
           curl --location 'localhost:9000/key-store' \
           --header 'Content-Type: application/json' \
           --data '{
               "command": "SET",
               "key": "FOO",
               "value": "BAR"
           }'
           ```
      - **Response:**
          ```json
          {"created":true}
          ```
  2. **DELETE:**
        - **Request:**
            ```bash
            curl --location 'localhost:9000/key-store' \
            --header 'Content-Type: application/json' \
            --data '{
                "command": "DELETE",
                "key": "FOO"
            }'
            ```
  3. **GET_OR_CREATE:**
       - **Request:**
         ```bash
         curl --location 'localhost:9000/key-store' \
         --header 'Content-Type: application/json' \
         --data '{
             "command": "GET_OR_CREATE",
             "key": "FOO",
             "value": "FOO"
         }'
         ```
      - **Response:**
        ```json
        {"value":"BAR"}
        ```
