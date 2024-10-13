# Funcionalidade de Fechamento Automático de Leilão

## Objetivo
Adicionar uma nova funcionalidade ao projeto já existente para o leilão fechar automaticamente a partir de um tempo definido.

## Início

### Clonar o Repositório
Clone o seguinte repositório para iniciar o projeto: [Clique aqui para acessar o repositório](#).

### Estado Atual do Projeto
Toda a rotina de criação do leilão e lances já está desenvolvida, entretanto, o projeto clonado necessita de melhoria: adicionar a rotina de fechamento automático a partir de um tempo.

## Descrição da Tarefa

Você deverá se concentrar no processo de criação de leilão (auction). A validação do leilão (auction) estar fechado ou aberto na rotina de novos lances (bid) já está implementada.

### O que Você Deve Desenvolver:
- Uma função que irá calcular o tempo do leilão, baseado em parâmetros previamente definidos em variáveis de ambiente.
- Uma nova Go routine que validará a existência de um leilão (auction) vencido (que o tempo já se esgotou) e deverá realizar o update, fechando o leilão (auction).
- Um teste para validar se o fechamento está acontecendo de forma automatizada.

## Dicas:
- Concentre-se no arquivo `internal/infra/database/auction/create_auction.go`, você deverá implementar a solução nesse arquivo.
- Lembre-se que estamos trabalhando com concorrência, implemente uma solução que lide com isso corretamente:
- Verifique como o cálculo de intervalo para checar se o leilão (auction) ainda é válido está sendo realizado na rotina de criação de bid.
- Para mais informações de como funciona uma goroutine, [clique aqui](#) e acesse nosso módulo de Multithreading no curso Go Expert.

## Entrega:

- O código-fonte completo da implementação.
- Documentação explicando como rodar o projeto em ambiente de desenvolvimento.
- Utilize Docker/Docker Compose para podermos realizar os testes da sua aplicação.

## Passos para executar o desafio

1. **Clone o Repositório:**

   ```bash
   git clone https://github.com/joao-fsilva/pos-go-expert.git
   cd pos-go-expert/2-labs/desafio-2/

2. **Configurar o ambiente:**
  - dentro de pos-go-expert/2-labs/desafio-2:
      ```bash
      docker-compose build
      docker-compose up -d
      docker exec -it desafio-2_app_1 bash

3. **Rodar o teste:**
  - dentro do container, em /app:
      ```bash
       go test ./...