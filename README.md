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

## Um dos modulos para fazer o download dos packages
Utilizando o comando "go mod tidy" ele realiza uma limpeza em seu go.mod instalando packages que por acaso 
voce não importou e está usando, ou desistalando packages inutilizados e atualizando o seu go.sum
 