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
- [ ] fazer o motor de dar match entre as transações & tags
- [ ] retornar consolidado de lista & total de transações por tag e mes/ano
- [ ] refatorar de uma forma estruturada
- [ ] cobrir com testes
- [ ] refatorar com boas praticas
- [ ] evoluir a obtenção para identificar quais transações já foram processadas para não inserir duplicatas
- [ ] download de CSV das transações que precisam incluir tags
- [ ] upload de CSV de de-para da transação & tag
- [ ] motor para auto-identificar quais termos as tags podem dar match para próximas transações