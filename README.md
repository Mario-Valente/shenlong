# Shenlong Project

Welcome to the Shenlong Project! This README will guide you through the setup and usage of the project.

## Table of Contents
- [Introduction](#introduction)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Introduction
The Shenlong Project is designed to run jobs in k8 via api or cli comands

## Installation
To install the Shenlong Project, follow these steps:

1. Clone the repository:
    ```bash
    git clone 
    ```
2. Navigate to the project directory:
    ```bash
    cd shenlong
    ```
3. Install the dependencies:
    ```bash
    make build
    ```

## Usage
To use the Shenlong Project, follow these steps:

## Uso/Exemplos

- Create job 
```bash
shenlong create-job --name teste --image busybox --command echo,hello 

```

- Server

```bash
shenlong server --port 3001

```


## Contributing


## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.