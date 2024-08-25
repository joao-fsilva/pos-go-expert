# Desafio - Clean Architecture

## Enunciado

Olá devs!
Agora é a hora de botar a mão na massa. Para este desafio, você precisará criar o usecase de listagem das orders.
Esta listagem precisa ser feita com:
- Endpoint REST (GET /order)
- Service ListOrders com GRPC
- Query ListOrders GraphQL
  Não esqueça de criar as migrações necessárias e o arquivo api.http com a request para criar e listar as orders.

Para a criação do banco de dados, utilize o Docker (Dockerfile / docker-compose.yaml), com isso ao rodar o comando docker compose up tudo deverá subir, preparando o banco de dados.
Inclua um README.md com os passos a serem executados no desafio e a porta em que a aplicação deverá responder em cada serviço.

## Passos para Executar o Projeto

1. **Clone o Repositório:**

   ```bash
   git clone https://github.com/joao-fsilva/pos-go-expert.git
   cd pos-go-expert/1-go-expert/3-desafio

2. **Instalar as dependências:**

    ```bash
    go mod tidy

3. **Subir os containers:**
    
    ```bash
   docker-compose up -d
   
4. **Subir a aplicação:**
     ```bash
    cd cmd/ordersystem/
    go run main.go wire_gen.go

5. **Executar o Endpoint REST (GET /order):**
   - A porta utilizada é 8000.
   - Acessar: http://localhost:8000/orders
   - Se prefefir, acessar api/list_orders.http e execute a requisição.

 
6. **Service ListOrders com GRPC:**
   - A porta utilizada é 50051.
   - Executar os comandos abaixo:
     ```bash
      evans -r repl
      call ListOrders

7. **Query ListOrders GraphQL:**
   - A porta utilizada é 8080.
   - Acessar: http://localhost:8080
   - Executar a seguinte query:
     - `query listOrders {
           listOrders {
               id
               Price
               Tax
               FinalPrice
           }
     }`