# Summarize-Transactions

Projeto dedicado para praticar golang

Ideia do projeto:
- obter CSV do extrato bancário através do e-mail
- gravar essas informações a fim de obter de-para de categoria para cada transação
- consolidar gastos por categorias

Tooling:
- conexão IMAP para buscar e-mail com CSV do extrato bancário
- banco SQL para armazenar transações a serem processadas

TODOs:
- [X] gravar no banco as transações e as tags incluindo de-para (transação 1:N tags)
- [X] estruturar quais termos cada tag deve dar match (alimentar manual por enquanto)
- [X] filtrar por mes-ano as categorias e transações
- [X] tela de categorias por mes-ano
- [ ] refazer transactions para retornar pela API por e-mail da pessoa (front também - fazer endpoint com id por e-mail da pessoa pra ter link acess'vel/compartilhável)
- [ ] retornar consolidado de lista das transações por categoria e mes/ano
- [ ] fazer o motor de dar match entre as transações & tags (já aproveitar para refazer gravando por e-mail da pessoa)
- [ ] deployar motor no cloud run & cloud schedule
- [ ] receber via e-mail da pessoa o CSV de gastos e gravar no db (retornar e-mail para pessoa falando que gravou com sucesso - retornar URL)
- [ ] validar se id do usuário é existente
- [ ] refatorar de uma forma estruturada
- [ ] cobrir com testes
- [ ] refatorar com boas praticas
- [ ] evoluir a obtenção para identificar quais transações já foram processadas para não inserir duplicatas
- [ ] download de CSV das transações que precisam incluir tags
- [ ] upload de CSV de de-para da transação & tag
- [ ] motor para auto-identificar quais termos as tags podem dar match para próximas transações