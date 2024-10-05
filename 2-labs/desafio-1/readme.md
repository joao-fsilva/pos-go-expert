## Objetivo
Desenvolver um sistema em Go que receba um CEP, identifique a cidade e retorne o clima atual (temperatura em graus Celsius, Fahrenheit e Kelvin). Este sistema deverá ser publicado no Google Cloud Run.

## Requisitos

- O sistema deve receber um CEP válido de 8 dígitos.
- O sistema deve realizar a pesquisa do CEP e encontrar o nome da localização; a partir disso, deverá retornar as temperaturas formatadas em:
    - Celsius
    - Fahrenheit
    - Kelvin.

- O sistema deve responder adequadamente nos seguintes cenários:
    - **Em caso de sucesso:**
        - Código HTTP: `200`
        - Response Body:
          ```json
          { "temp_C": 28.5, "temp_F": 28.5, "temp_K": 28.5 }
          ```

    - **Em caso de falha (CEP inválido):**
        - Código HTTP: `422`
        - Mensagem: `invalid zipcode`

    - **Em caso de falha (CEP não encontrado):**
        - Código HTTP: `404`
        - Mensagem: `can not find zipcode`

## Deploy
Deverá ser realizado o deploy no Google Cloud Run.

## Dicas
- Utilize a API [viaCEP](https://viacep.com.br/) (ou similar) para encontrar a localização que deseja consultar a temperatura.
- Utilize a API [WeatherAPI](https://www.weatherapi.com/) (ou similar) para consultar as temperaturas desejadas.
- Para realizar a conversão de Celsius para Fahrenheit, utilize a seguinte fórmula:
    - \( F = C * 1,8 + 32 \)
- Para realizar a conversão de Celsius para Kelvin, utilize a seguinte fórmula:
    - \( K = C * 273 \)

  Sendo:
    - \( F \) = Fahrenheit
    - \( C \) = Celsius
    - \( K \) = Kelvin

## Entrega
- O código-fonte completo da implementação.
- Testes automatizados demonstrando o funciona

## Passos para executar o desafio

1. **Clone o Repositório:**

   ```bash
   git clone https://github.com/joao-fsilva/pos-go-expert.git
   cd pos-go-expert/2-labs/desafio-1/

2. **Configurar o .env dentro de cmd/weather-zip-code:**
    - dentro de pos-go-expert/2-labs/desafio-1/cmd/weather-zip-code:

    ```bash
    cp .env.example .env

3. **Configurar o ambiente:**
   - dentro de pos-go-expert/2-labs/desafio-1:
       ```bash
       docker build -t weather-zip-code .
       docker run -d -p 8080:8080 --name weather-zip-code -v $(pwd):/app weather-zip-code
       docker exec -it weather-zip-code bash

4. **Rodar os testes:**
    - dentro do container, em /app:
        ```bash
         go test ./...

5. **Subir a aplicação**
    - dentro do container, em /app/cmd/weather-zip-code:
        ```bash
        go run main.go
      
6. **Executar sucesso:**
    - Acessar: http://localhost:8080/weather?zipcode=07011020

7. **Executar CEP inválido:**
    - Acessar: http://localhost:8080/weather?zipcode=070110200

8. **Executar CEP inexistente:**
    - Acessar: http://localhost:8080/weather?zipcode=00000000


