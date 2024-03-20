[x] Os 3 contextos deverão retornar erro nos logs caso o tempo de execução seja insuficiente.

- client.go
  [x] O client.go deverá realizar uma requisição HTTP no server.go solicitando a cotação do dólar.

  [x] O client.go precisará receber do server.go apenas o valor atual do câmbio (campo "bid" do JSON). 
  
  Context:
  [x] timeout máximo de 300ms para receber o resultado do server.go.

  File:
  [x] O client.go terá que salvar a cotação atual em um arquivo "cotacao.txt" no formato: Dólar: {valor}


- server.go
  [x] requisição para https://economia.awesomeapi.com.br/json/last/USD-BRL e retornar no formato JSON o resultado para o cliente.
 
  BD:
  [x] registrar no banco de dados SQLite cada cotação recebida

  Context:
  [x] timeout máximo para chamar a API de cotação do dólar deverá ser de 200ms
  [x] timeout máximo para gravar no BD deverá ser de 10ms.
 
  [x] O endpoint necessário gerado pelo server.go para este desafio será: /cotacao e a porta a ser utilizada pelo servidor HTTP será a 8080.
