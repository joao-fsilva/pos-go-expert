# Rate Limiter em Go

## Objetivo

Desenvolver um rate limiter em Go que possa ser configurado para limitar o número máximo de requisições por segundo com base em um endereço IP específico ou em um token de acesso.

## Descrição

O objetivo deste desafio é criar um rate limiter em Go que possa ser utilizado para controlar o tráfego de requisições para um serviço web. O rate limiter deve ser capaz de limitar o número de requisições com base em dois critérios:

- **Endereço IP**: O rate limiter deve restringir o número de requisições recebidas de um único endereço IP dentro de um intervalo de tempo definido.
- **Token de Acesso**: O rate limiter deve também poderá limitar as requisições baseadas em um token de acesso único, permitindo diferentes limites de tempo de expiração para diferentes tokens. O Token deve ser informado no header no seguinte formato:
  - API_KEY: token 
- As configurações de limite do token de acesso devem se sobrepor as do IP. Ex: Se o limite por IP é de 10 req/s e a de um determinado token é de 100 req/s, o rate limiter deve utilizar as informações do token.

## Requisitos

- O rate limiter deve poder trabalhar como um middleware que é injetado ao servidor web.
- O rate limiter deve permitir a configuração do número máximo de requisições permitidas por segundo.
- O rate limiter deve ter a opção de escolher o tempo de bloqueio do IP ou do Token caso a quantidade de requisições tenha sido excedida.
- As configurações de limite devem ser realizadas via variáveis de ambiente ou em um arquivo `.env` na pasta raiz.
- Deve ser possível configurar o rate limiter tanto para limitação por IP quanto por token de acesso.
- O sistema deve responder adequadamente quando o limite é excedido:
  - Código HTTP: 429
  - Mensagem: "you have reached the maximum number of requests or actions allowed within a certain time frame"
- Todas as informações de "limiter” devem ser armazenadas e consultadas de um banco de dados Redis. Você pode utilizar docker-compose para subir o Redis.
- Crie uma "strategy" que permita trocar facilmente o Redis por outro mecanismo de persistência.
- A lógica do limiter deve estar separada do middleware.

## Exemplos

### Limitação por IP

Suponha que o rate limiter esteja configurado para permitir no máximo 5 requisições por segundo por IP. Se o IP `192.168.1.1` enviar 6 requisições em um segundo, a sexta requisição deve ser bloqueada.

### Limitação por Token

Se um token `abc123` tiver um limite configurado de 10 requisições por segundo e enviar 11 requisições nesse intervalo, a décima primeira deve ser bloqueada.

Nos dois casos acima, as próximas requisições poderão ser realizadas somente quando o tempo total de expiração ocorrer. Ex: Se o tempo de expiração é de 5 minutos, determinado IP poderá realizar novas requisições somente após os 5 minutos.

## Dicas

- Teste seu rate limiter sob diferentes condições de carga para garantir que ele funcione conforme esperado em situações de alto tráfego.

## Entrega

- O código-fonte completo da implementação.
- Documentação explicando como o rate limiter funciona e como ele pode ser configurado.
- Testes automatizados demonstrando a eficácia e a robustez do rate limiter.
- Utilize docker/docker-compose para que possamos realizar os testes de sua aplicação.
- O servidor web deve responder na porta 8080.


## Como o rate limiter funciona

- Recebe a requisição (cmd/ratelimiter/main.go)
- Encaminha a requsição para o middleware (internal/middleware/ratelimiter_middleware.go)
  - Carrega através das variáveis de ambiente o rate, que é o limite máximo de requisições por segundo e o block duration, que é o tempo em minutos que o usuário ficará sem acessar no caso de bloqueio.
  - Caso seja passado um token, altera o rate e o block duration para os valores passados no token.
    - Exemplo de token: apikey_rate_5_blockduration_1
- O middleware chama o caso de uso, passando todos os parâmetros já setados (internal/usecase/ratelimiter.go)
- No caso de uso, o fluxo é o seguinte:
  - Se a entity (internal/entity/limiter.go) não existe no redis (internal/infra/limiter_repository_redis.go), cria gerando 1 acesso.
    - Na entity, o método IncrementAccessCount lida com as regras de incremento.
  - Se existe e expirou, ou seja, não teve acesso dentro de 1 segundo, o count de acessos será resetado.
  - Se existe, não expirou e ultrapassou o limite, então realiza o bloqueio pelo tempo parametrizado.
  - Se existe, não expirou e não ultrapassou o limite, realiza o incremento na quantidade de acessos.

- Sobre este ponto: Crie uma “strategy” que permita trocar facilmente o Redis por outro mecanismo de persistência.
  - Foi utilizada a arquitetura clean architecture, sendo possível implementar a interface limiter_repository (internal/entity/limiter_repository.go) com outro mecanismo de persistência e realizar a troca no main.go

## Passos para executar o desafio

1. **Clone o Repositório:**

   ```bash
   git clone https://github.com/joao-fsilva/pos-go-expert.git
   cd pos-go-expert/3-desafios-tecnicos/desafio-1/

2. **Configurar o .env (opcional, já tem valores default):**
    - Dentro de pos-go-expert/3-desafios-tecnicos/desafio-1/cmd/.env
    - Caso altere, os testes abaixo terão um comportamento diferente.

3. **Configurar o ambiente:**
    - dentro de pos-go-expert/3-desafios-tecnicos/desafio-1:
        ```bash
        docker-compose build
        docker-compose up app-prod

4. **Teste por IP (utilizar a ferramenta ApacheBench)**
    - Acessar:  ab -n 1001 -c 1 -k http://localhost:8080/
    - Esperado: Ter sucesso em 1000 acessos, dar erro no acesso 1001 e bloquear o IP por 1 minuto.

5. **Teste por token (utilizar a ferramenta ApacheBench)**
    - Acessar: ab -n 10 -c 1 -k -H "API_KEY: apikey_rate_5_blockduration_1"  http://localhost:8080/
    - Esperado: Ter sucesso em 5 acessos, dar erro em 5 acessos e bloquear o token por 1 minuto. 

6. **Rodar os testes automatizados:**
    - dentro de pos-go-expert/3-desafios-tecnicos/desafio-1:
        ```bash
        docker exec -it app-prod bash
        go test ./...

