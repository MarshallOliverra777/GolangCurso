package prefixo

import "strconv"

//Capital representa o numero do prefixo da capital de um estado/provincia
var Capital = 11

var teste = "teste"

//TesteComPrefixo é apenas um teste
var TesteComPrefixo = teste + " " + strconv.Itoa(Capital)
