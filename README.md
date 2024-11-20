# Shenlong Project

Welcome to the Shenlong Project! This README will guide you through the setup and usage of the project.

## Table of Contents
- [Introduction](#introduction)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)
- [flags](#flags)

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

- Create Cron

```bash
shenlong create-cron --name teste --image busybox --comand echo,hello  

```

- Get 

```bash
shenlong get --job --name teste --namespace default

```

- Server

```bash
shenlong server --port 3001

```


## Contributing


## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.


## flags

| Comando         | Flag           | Tipo         | Valor Padrão      | Descrição                                          | Obrigatório |
|------------------|----------------|--------------|-------------------|--------------------------------------------------|-------------|
| **create-job**   | `--image`      | `String`     | `busybox:latest`  | Imagem para executar o job no Kubernetes         | Não         |
|                  | `--cpu`        | `String`     | `""`              | Quantidade de CPU para o job                     | Não         |
|                  | `--memory`     | `String`     | `""`              | Quantidade de memória para o job                 | Não         |
|                  | `--namespace`  | `String`     | `default`         | Namespace onde o job será criado                 | Não         |
|                  | `--name`       | `String`     | `""`              | Nome do job                                      | Sim         |
|                  | `--command`    | `StringSlice`| `[]`              | Comando para executar no job                     | Sim         |
|                  | `--kubeconfig` | `String`     | `""`              | Caminho para o arquivo kubeconfig                | Não         |
|                  | `--ttl`        | `Int32`      | `100`             | Tempo de vida do job após a execução             | Não         |
| **server**       | `--run`        | `String`     | `""`              | Rodar o servidor                                 | Não         |
| **get**          | `--name`       | `String`     | `""`              | Nome do job                                      | Sim         |
|                  | `--namespace`  | `String`     | `default`         | Namespace do job                                 | Não         |
|                  | `--kubeconfig` | `String`     | `""`              | Caminho para o arquivo kubeconfig                | Não         |
|                  | `--job`        | `String`     | `""`              | Job a ser obtido                                 | Não         |
|                  | `--cron`       | `String`     | `""`              | Cron a ser obtido                                | Não         |
| **create-cron**  | `--image`      | `String`     | `busybox:latest`  | Imagem para executar o cron no Kubernetes        | Não         |
|                  | `--cpu`        | `String`     | `""`              | Quantidade de CPU para o cron                    | Não         |
|                  | `--memory`     | `String`     | `""`              | Quantidade de memória para o cron                | Não         |
|                  | `--namespace`  | `String`     | `default`         | Namespace onde o cron será criado                | Não         |
|                  | `--name`       | `String`     | `""`              | Nome do cron                                     | Sim         |
|                  | `--command`    | `StringSlice`| `[]`              | Comando para executar no cron                    | Sim         |
|                  | `--kubeconfig` | `String`     | `""`              | Caminho para o arquivo kubeconfig                | Não         |
|                  | `--ttl`        | `Int32`      | `100`             | Tempo de vida do cron após a execução            | Não         |
|                  | `--schedule`   | `String`     | `*/5 * * * *`     | Agenda para executar o cron                      | Sim         |
