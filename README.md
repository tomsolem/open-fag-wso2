# OpenFGA Development Environment

This project provides a fully configured, containerized development environment for building applications with [OpenFGA](https://openfga.dev/). It uses VS Code Dev Containers to streamline setup, so you can start coding in a consistent and reproducible workspace.

The environment comes pre-installed with Go, the `fga` CLI, `psql` client, and other essential development tools, all connected to a running OpenFGA instance with a PostgreSQL backend.

## Features

- **One-click Setup**: Launch a complete environment with a single VS Code command.
- **Integrated Services**: Includes OpenFGA server and PostgreSQL database, networked together.
- **Pre-installed Tooling**:
  - Go (latest version)
  - Delve (Go debugger)
  - OpenFGA CLI (`fga`)
  - PostgreSQL Client (`psql`)
  - Zsh with a clean default setup
- **Reproducible**: Ensures all developers work with the same tool versions and configuration.
- **Cross-platform**: Works on any machine that can run Docker and VS Code.

## Prerequisites

Before you begin, ensure you have the following installed:

1.  [Docker Desktop](https://www.docker.com/products/docker-desktop/)
2.  [Visual Studio Code](https://code.visualstudio.com/)
3.  [VS Code Dev Containers extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)

## Quick Start

1.  **Clone the repository:**
    ```bash
    git clone <your-repository-url>
    cd <your-repository-name>
    ```

2.  **Open in VS Code:**
    ```bash
    code .
    ```

3.  **Reopen in Container:**
    -   After opening the folder, VS Code will detect the `.devcontainer` configuration and show a notification in the bottom-right corner.
    -   Click **"Reopen in Container"**.
    -   VS Code will build the Docker images and set up the environment. This may take a few minutes on the first run.

Once the container is built, you will have a VS Code window connected to your new development environment, with the terminal using Zsh. Short note: the go debugger may not start stating that the delve is not installed. You need to "reload" the vs code window to make it work.

## Using the Environment

The environment is designed to be ready to use out-of-the-box. All services are started and configured automatically.

### Verify Services

- **Check Docker Containers**: Open a new terminal in VS Code and run `docker ps` to see the running services (`openfga`, `postgres`, `dev-container`).
- **Check Go version**:
  ```bash
  go version
  ```
- **Check FGA CLI**:
  ```bash
  fga version
  ```

### Interacting with OpenFGA

The `fga` CLI is pre-configured to communicate with the `openfga` service.

- **Create a new store**:
See the /open-fga folder. There is a script called configure-model.sh. This will import the store and model to the postgres db. It allso adds env.vars for store and model id.
  ```bash
  source /workspace/open-fga/configure-model.sh
  ```

### Connecting to PostgreSQL

The `psql` client is configured via environment variables to connect directly to the `postgres` service.

- **Open a `psql` session**:
  ```bash
  psql
  ```
- **List databases**:
  ```sql
  \l
  ```
- **Exit `psql`**:
  ```sql
  \q
  ```

## Development Workflow

1.  **Define your model**: Create a `.fga` file to define your authorization model's types and relations.
2.  **Write Go code**: Develop your application logic in Go, using the OpenFGA Go SDK to interact with the `openfga` service. To debug the code start the launcher. The code write out a curl request with a JWT token that can be used for testing.
3.  **Test**: Write and run tests for your application and authorization logic. The integrated Go tooling and Delve debugger are available for use.
4.  **Commit**: Commit your changes from within the dev container. Git is pre-installed and configured.

## Services Overview

This environment is managed by the `docker/docker-compose.yaml` file and consists of three main services:

-   `postgres`: A PostgreSQL 17 database instance used as the data store for OpenFGA.
-   `openfga`: The OpenFGA server. It is configured to use the `postgres` service.
-   `dev-container`: Your main development workspace. It contains the Go toolchain, CLIs, and your source code mounted from

## TODO

- Setup WSO2 docker image. See here for details: [wso2 docker compose](https://github.com/wso2/docker-apim/blob/master/docker-compose/apim-with-mi/README.md).
- Create a API swagger for the GoLang backend and publis this to WSO2.
- Use custom transformation in WSO2 to do the authorization using openFGA.
- Create a better store and model to show more complex senarios.