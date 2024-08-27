# Monitoramento de Servidores

Este projeto é uma aplicação em Go que verifica a disponibilidade de servidores a partir de uma lista de URLs fornecidas em um arquivo CSV. Caso algum servidor esteja fora do ar ou não responda corretamente, a aplicação registra essas informações em outro arquivo CSV.

## Requisitos

- Go 1.16 ou superior.

## Instalação

1. Clone o repositório ou copie o código para o seu ambiente local.
2. Certifique-se de ter o Go instalado em seu ambiente.

## Estrutura do Projeto

- **main.go**: Arquivo principal que contém toda a lógica de verificação e registro de servidores.
- **<server_list>**: Arquivo CSV que contém a lista de servidores a serem monitorados. Cada linha do arquivo deve seguir o seguinte formato:

    ```
    server_name,server_url
    ```
  
  Exemplo:
  
    ```
    Google,https://www.google.com
    GitHub,https://www.github.com
    ```
  
- **<downtime_list>**: Arquivo CSV onde serão registrados os detalhes dos servidores que estiverem fora do ar ou não responderem corretamente. Cada linha do arquivo será escrita no seguinte formato:

    ```
    server_name,server_url,data_falha,tempo_execucao,status_http
    ```

## Como Usar

1. Prepare o arquivo `<server_list>` com a lista de servidores que deseja monitorar.
2. Execute o comando abaixo substituindo `<server_list>` e `<downtime_list>` pelos respectivos arquivos:

    ```bash
    go run main.go <server_list> <downtime_list>
    ```

3. O programa irá verificar cada servidor listado e exibirá no terminal o status HTTP e o tempo de resposta para cada servidor.
4. Se algum servidor estiver fora do ar, os detalhes serão registrados no arquivo `<downtime_list>`.

## Exemplo de Execução

```bash
go run main.go servers.csv downtime.csv
```

**Saída esperada no terminal:**

```
Status [200] Tempo de carga: [0.123456] URL: [https://www.google.com]
Status [503] Tempo de carga: [0.234567] URL: [https://www.example.com]
```

## Funções Principais

- **verificarParametros**: Verifica se os parâmetros necessários foram passados na linha de comando.
- **openFiles**: Abre os arquivos CSV de entrada e saída.
- **criarListaServidores**: Lê o arquivo CSV de servidores e cria uma lista de estruturas `Server`.
- **checkServer**: Faz uma requisição HTTP para cada servidor e verifica se ele está respondendo corretamente.
- **generateDownTime**: Escreve os servidores que estão fora do ar no arquivo de registro.

## Licença

Este projeto é distribuído sob a licença MIT. Consulte o arquivo `LICENSE` para obter mais informações.