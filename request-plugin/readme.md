# Tyk Request Plugin

This repository contains an example of how to write a request plugin for Tyk using Go.

## Prerequisites

Before you begin, make sure you have the following installed:

- Go (version 1.23 or higher)
- Tyk Gateway (version 5.0.10 or higher)

## Getting Started

To get started with writing a request plugin for Tyk, follow these steps:

1. Clone this repository

2. Navigate to the cloned repository:

    ```bash
    cd request-plugin
    ```

3. Implement your request plugin logic in the `plugin.go` file.

4. Build the plugin:

    ```bash
    go build -buildmode=plugin -o request_plugin.so main.go
    ```

    If you use docker follow this step 
    [TYK - Build Plugin Docker - Cross-compiling for different architectures and operating systems](https://tyk.io/docs/product-stack/tyk-gateway/advanced-configurations/plugins/golang/go-plugin-compiler/#cross-compiling-for-different-architectures-and-operating-systems)

    ```
    docker run --rm -v `pwd`:/plugin-source \
           --platform=linux/amd64 \
           tykio/tyk-plugin-compiler:v5.0.10 plugin.so
    ```
    

5. Copy the generated `plugin.so` file to the Tyk plugins directory:

    ```bash
    cp plugin.so /path/to/tyk/plugins/
    ```

6. Configure Tyk to use the request plugin by adding the following to your Tyk Gateway configuration file:

    ```json
    "custom_middleware": {
        "pre": [],
        "post_key_auth": [],
        "auth_check": {},
        "post": [
            {
            "name": "AddCustomHeader",
            "path": "<path>/plugin.so"
            }
        ],
        "driver": "goplugin"
    }
    ```

7. Restart the Tyk Gateway.

## Usage

Once the Tyk Gateway is restarted with the request plugin enabled, it will execute your custom logic for each incoming request.
