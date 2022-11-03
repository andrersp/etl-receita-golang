# E.T.L Receita Federal

> Desenvolvimento de microsserviço para realizar o **Extração**, **Transformação** e **Carregamento** dos dados públicos CNPJ do site da [Receita Federal](https://www.gov.br/receitafederal/pt-br/assuntos/orientacao-tributaria/cadastros/consultas/dados-publicos-cnpj) utilizando **Golang**.

![Badge](https://img.shields.io/static/v1?label=Go&message=1.18&color=green&style=flat&logo=GO)

> Status do Projeto e Documentação: Em desenvolvimento :warning:

# Features
### Extração
- [x] Scrap do site para obter as *urls* com os arquivos que precisamos. 
- [x] Inicar o download de cada arquivo em Goroutine.
- [x] Criar um retry para reiniciar o download caso haja alguma falha na conexão.
- [x] Descompactar o arquivo .csv ao final de cada download e remover o .zip para consumir menos espaço em disco

### Transformação
- [x] Tratar os caracteres com charset inválidos.
- [x] Remover caracteres não numéricos de string com números.
- [x] Unir os arquivos .csv por categoria.

### Carregamento
- [ ] Armazenar no PostgreSQL




### Maintainers:
* André França              rsp.assistencia@gmail.com

### Contributing

1. Faça o _fork_ do projeto (<https://github.com/andrersp/go-etl-receita-federal/fork>)
2. Crie uma _branch_ para sua modificação (`git checkout -b feature/fooBar`)
3. Faça o _commit_ (`git commit -am 'Add some fooBar'`)
4. _Push_ (`git push origin feature/fooBar`)
5. Crie um novo _Pull Request_



