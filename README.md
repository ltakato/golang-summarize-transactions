# Summarize-Transactions

Projeto dedicado para praticar golang

Ideia do projeto:
- obter CSV do extrato bancário através do e-mail
- gravar essas informações a fim de obter de-para de categoria para cada transação
- consolidar gastos por categorias

Tooling:
- conexão IMAP para buscar e-mail com CSV do extrato bancário
- banco SQL para armazenar transações a serem processadas