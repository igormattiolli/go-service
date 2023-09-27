<img src="docs/assego/blue-logo-stargoe.svg" alt="Logo Stargoe" width="250">

# Checkout API

## Sobre

O checkout-api é o responsável por todas as compras realizadas pelo checkout v3, que está mantido no repositório https://github.com/Stargoe/checkout-web.

Essa aplicação está desacoplada, seguindo boas práticas de Arquitetura Limpa, bem como os princípios do SOLID.

## Estrutura

Estrutura de exemplo com o dominio do produto:

```ssh
.
├── src
│   ├── app.go
│   ├── domain
│   │   ├── product
│   │   │   ├── data
│   │   │   │   └── IProductRepository.go
│   │   │   ├── errors
│   │   │   │   └── ProductNotFound.go
│   │   │   ├── entities
│   │   │   │   └── ProductEntity.go
│   │   │   └── useCases
│   │   └──     └── ProductUseCase.go
│   ├── infra
│   │   └── http
│   │       ├── product
│   │          ├── controllers
│   │          │   └── ProductController.go
│   │          └── routes
│   │            └── Routes.go
│   ├── providers
│   │   ├── database
│   │   │       ├── index.go
│   │   │       ├── repositories
│   │   │           └── ProductRepository.go
```

### Camada de Domínio

Caminho: `src/domain`.

É o diretório de alto nível. Onde se encontra os dominios que compoem o checkout.

#### Entities

Caminho: `src/domain/**/entities`.

As entidades são as formas como modelamos nossos domínios. É importante entender que as entidades que estão no nosso projeto são diferentes das que estão no banco de dados, pois elas possuem apenas as propriedades que necessitamos e não toda a informação.

#### Use Cases

Caminho: `src/domain/**/use-cases`.

Os casos de usos são as classes que utilizam as entidades junto dos providers e repositórios para manipular as regras de negócios. Por exemplo, quando precisamos fazer uma compra, independente do método de pagamento, criamos um caso de uso chamado TransactionUseCase que fará toda a lógica de transação utilizando nosso gateways de pagamento.

#### Data

Caminho: `src/domain/**/data`.

Em nosso diretório data possuímos as nossas interfaces que podemos chamá-las de contratos Esses contratos são usados para desacoplar os nossos códigos de nivel alto (regras de negócios e entidades) do nosso código de nivel baixo (frameworks e base dados) de acordo com o SOLID(Interface Segregation Principle):

### Camada de Providers

Caminho: `src/providers`.

No diretório de providers estão todos os nossos providers, gateways, banco de dados e mocks que usamos, ele tem como principal objetivo separar em um local todas as nossas ligações externas que fazemos em nosso projeto.

### Camada de Infra

Caminho: `src/infra`.

A infra é a nossa camada que vai se comporta como um intermediário, ela é responsável pela conexão entre as nossas regras de negócios e com os nossos providers, nela temos nossos design patterns (controllers, factories, observers e proxies).

#### Factory

Caminho: `src/infra/factory`.

As Factories são responsáveis pela injeção de dependências dentro dos nossos use cases, ou seja, é nela que possuímos a chamada direta do nossos providers e que através do constructor injetamos ele.

#### http

Caminho: `src/infra/http`.

No diretório http possuímos nossa implementação do protocolo http que consiste na divisão entre controllers e routes.
Os controllers são responsáveis por fazer a intermediação entre a informação que vem das requests com a chamada das factories, assim apenas fazendo a passagem de informação, sem ter qualquer tipo de regras de negócio.

O diretório routes é responsável por mapear as rotas que o nosso projeto possui, nela colocamos qual vai ser o endpoint junto com os nossos middlewares que serão executados antes das chamadas dos controllers.

## Como rodar

Rode o seguinte comando:

Para baixar as dependencias

```sh
go mod tidy
```

Para subir a API:

```sh
go run main.go
```
