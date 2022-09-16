package models

type Empresa struct {
	Id               int    `json:"id"`
	Nome             string `json:"nome"`
	Cnpj             string `json:"cnpj"`
	Razao_social     string `json:"razao_social"`
	Endereco_empresa string `json:"endereco_empresa"`
	Ramo_de_atuacao  string `json:"ramo_de_atuacao"`
}

var Empresas []Empresa
