## Criando modulo GO
Para criar um modulo go utilize "go mod init [nome_do_modulo]" o nome do modulo pode ser qualquer coisa
Como fazemos a criação do modulo de forma atrelada a um repositorio então virou um costume adicionar o
endereço do repo onde o modulo irá ficar.

ex: go mod init github.com/PauloHPMKT/goopportunities

## Arquivo go.mod
É o arquivo central do Go, é onde ele inicializa o projeto descrevendo qual o nome do modulo e qual a 
versão que voce esta usando

## Arquivo main.go
É o arquivo que inicialmente será procurado e carregará inicialmente a aplicação
a definição de package geralmente é o escopo em que estamos mexendo. 
Ex: main.go --> package main

Imagina o package como um subprojeto dentro de seu projeto
Então dentro de um package voce tem acesso a todas as variaveis independentes de elas estarem em arquivos
diferentes, encapsulando tudo dentro do mesmo package de modulo.

A func main também é uma convensão e é utilizada para inicializar a aplicação.

## Ao Instanciar um modulo dentro do Go
Quando voce tem qualquer tipo de package externo ao seu package, por padrão o ultimo nome do package importado
será o nome referenciado dentro do seu codigo.
```sh
  import "github.com/gin-gonic/gin"

  import (
    gin "github.com/gin-gonic/gin"
  )
```
Ou seja, é totalmente possível trocar, sobreescrever o nome do package
gin --> g "github.com/gin-gonic/gin" 
Com isso, sempre que voce não der um nome para o package externo o Go utilizará o que já está por padrão

## Um dos modulos para fazer o download e gerenciamento dos packages
Utilizando o comando "go mod tidy" ele realiza uma limpeza em seu go.mod instalando packages que por acaso 
voce não importou e está usando, ou desistalando packages inutilizados e atualizando o seu go.sum

## Go sum
A linguagem Go sempre faz o download de um arquivo binario (exe) para o GOPATH. O arquivo Go Sum representa
o equivalente a um lock do javaScript, então ele concentra as dependencias diretas ou indiretas.

## Padrões de projeto
Seguindo uma convensão do Go, toda pasta que criamos dentro de um projeto Go assumimos que aquela pasta é um novo 
subpackage, ou seja, ele passa a ser um submodulo novo.
Ex: module github.com/PauloHPMKT/goopportunities, quando criamos um router ele passa a ser o path do modulo
principal.

Diante disso é feito o import do package e não do arquivo.
Ex: 
```sh

## pacote router
package router

import "github.com/gin-gonic/gin"

func InitializeRouter(port string) {}

## utilizando o pacote router
package main

import	"github.com/PauloHPMKT/goopportunities/router"

func main() {
	router.InitializeRouter(port)
}

```
Os Imports e Exports dos packages em Go são por convensão de nomenclatura, ou seja, se voce tem qualquer variavel,
função, constante ou qualquer tipo de entidade na linguagem go ela automaticamente esta sendo exportada do seu 
package. Ou seja, se voce os cria com letra Maiuscula ele pode ser utilizado em qualquer package da aplicação, se 
for feito em letra minuscula ele fica retido ao subpackage.

No contexto interno do package ao criar um arquivo auxiliar àquele pacote voce pode criar funçoes privadas, variaveis e instancias privadas que só são utilizadas dentro daquele package.
Ex: 
|__ router
    |___ router.go
    |___ routes.go -> O que for criado dentro de routes.go estará disponível no package principal

Ex: 
```sh

## router.go. No exemplo da funçao utilizada printRouter, por ser criada com a primeira letra minuscula
## ela é um recurso privado e está acessível apenas dentro do package
func InitializeRouter() {
  printRouter()
}

## routes.go
func printRouter() {}
```

## Qual a ideia do Pointer?
Imagina que tudo que temos em codigo esta em algum lugar em nosso sistema operacional... Como uma variavel por exemplo,
Essa variavel ela tem uma casinha, em um endereço (Linguagens que não são fortemente tipadas geralmente temos um conceito
chamado pass by value, que é voce pegar a casinha e enviar literalmente de um lado para o outro). A ideia do pointer é que 
voce envie o CEP para quem está chamando a casinha, ou seja, voce manda um ponteiro para a casa, então se voce precisa
acessar o conteudo daquela casinha voce acessa o endereço da casa.

Ou seja, ele não é a coisa em si, ele é o endereço (referencia) para a coisa.

O Ganho em performance é que haverá momentos que um recurso passado de um lado para outro, novas instancias serão geradas
e isso pode comprometer o funcionamento e o comportamento da entidade, pois o ponteiro indica o endereço para uma unica
entidade.

Para acessar o endereço utilizamos * Ponteiro, já para enviar o endereço utilizamos &

## Live reload Go
go install github.com/air-verse/air@latest
https://aprendagolang.com.br/como-utilizar-live-reload-para-melhorar-sua-produtividade/

  __    _   ___
 / /\  | | | |_)
/_/--\ |_| |_| \_ v1.52.3, built with Go go1.22.4

 __    _   ___
 / /\  | | | |_)
/_/--\ |_| |_| \_ v1.52.3, built with Go go1.22.4

.air.toml file created to the current directory with the default settings

## O que é uma struct?
É equivalente a um objeto ou uma interface que define os tipos dos campos a serem criados dentro de uma estrutura

## Obs:
--> Quando criamos uma variavel var fora de uma função é uma variavel global, do package
--> Curiosidades sobre o print 
    |___ Print("Eu quero printar esse texto") = Faz que o texto printado seja exatamente o de dentro da linha
    |___ Printf("Eu quero imprimir isso aqui %s", string) = Concatena o print com a string