# File Lock

![GitHub](https://img.shields.io/github/license/NalbertLeal/file-lock)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/NalbertLeal/file-lock)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/NalbertLeal/file-lock)

## Descrição

_File Lock_ é uma ferramenta em Go desenvolvida para criptografar arquivos usando o algoritmo _Advanced Encryption Standard_ (AES). Não utilize esta ferramenta como sua principal solução de criptografia; ela é destinada apenas para criptografar rapidamente um arquivo que não deve ser aberto por ninguém.

## Como Usar

### Pré-requisitos

Certifique-se de ter o Go instalado em sua máquina. Para instruções sobre a instalação do Go, visite [https://golang.org/doc/install](https://golang.org/doc/install).

### Instalação

```bash
go get -u github.com/NalbertLeal/file-lock
```

### Uso Básico

Criptografe um arquivo chamado _file\_name.txt_ usando AES com a chave _key123_:

```bash
file-lock file_name.txt key123
```

## Contribuição

Se você deseja contribuir com melhorias, correções de bugs ou novos recursos, sinta-se à vontade para abrir uma issue ou enviar um pull request. Suas contribuições são altamente apreciadas!

## Licença

Este projeto é licenciado sob a [GNU LESSER GENERAL PUBLIC LICENSE](https://choosealicense.com/licenses/lgpl-2.1/).