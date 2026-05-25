# go-api-lab

API REST desenvolvida em Go (Golang) com foco em aprendizado prático de backend moderno, aplicando conceitos de autenticação JWT, arquitetura em camadas, integração com PostgreSQL e containerização com Docker.

O objetivo deste projeto foi consolidar fundamentos importantes do ecossistema Go através da construção de uma aplicação backend organizada, segura e próxima de cenários reais de desenvolvimento.

---

# 🚀 Tecnologias Utilizadas

- **Go (Golang)**
- **Gin Gonic** — framework HTTP de alta performance
- **PostgreSQL**
- **JWT (JSON Web Token)**
- **bcrypt**
- **Docker**
- **Docker Compose**
- **Git/GitHub**

---

# 🏗 Arquitetura e Organização

O projeto foi estruturado utilizando separação em camadas para melhorar organização, manutenção e responsabilidade de cada parte da aplicação.

```bash
├── cmd/          # Ponto de entrada da aplicação
├── controller/   # Handlers HTTP e rotas
├── db/           # Configuração do banco de dados
├── model/        # Estruturas e entidades
├── repository/   # Acesso a dados e queries SQL
├── usecase/      # Regras de negócio
├── docker-compose.yml
├── dockerfile
├── go.mod
└── go.sum
```

## Estrutura aplicada

- **Controller → Usecase → Repository**
- Separação entre regras de negócio e acesso a dados
- Organização modular do código
- Estrutura inspirada em conceitos de Clean Architecture

---

# 🛠 Funcionalidades

## Usuários

### Cadastro de usuários

```http
POST /users/signup
```

- Criptografia de senha utilizando bcrypt
- Validação de usuário existente

---

### Login e autenticação

```http
POST /users/signin
```

- Geração de token JWT
- Autenticação baseada em Bearer Token

---

## Produtos (Rotas protegidas)

### Listar produtos

```http
GET /products
```

### Buscar produto por ID

```http
GET /product/:productId
```

### Criar produto

```http
POST /product
```

### Atualizar produto

```http
PUT /product/:productId
```

### Remover produto

```http
DELETE /product/:productId
```

---

# 🔐 Segurança

As rotas protegidas utilizam autenticação JWT através de middleware responsável pela validação do token enviado no cabeçalho:

```http
Authorization: Bearer <token>
```

Também foi implementada:

- Criptografia de senhas com bcrypt
- Validação de autenticação
- Tratamento de erros para acessos inválidos

---

# 🐳 Dockerização

A aplicação foi containerizada utilizando Docker e Docker Compose para facilitar configuração e padronização do ambiente.

## Executar o projeto

```bash
docker compose up --build
```

A API ficará disponível em:

```bash
http://localhost:8000
```

---

# ⚙ Variáveis de ambiente

Exemplo do arquivo `.env`:

```env
SECRET_KEY=sua_secret_key

DB_HOST=go_db
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=1234
DB_NAME=postgres
```

---

# 📚 Conceitos praticados no projeto

Durante o desenvolvimento deste projeto, foram aplicados conceitos importantes de backend como:

- Desenvolvimento de APIs REST
- Organização de projetos em Go
- Autenticação JWT
- Middleware de autenticação
- Criptografia de senhas
- Integração com PostgreSQL
- Dockerização de aplicações
- Tratamento de erros
- Separação de responsabilidades
- Arquitetura em camadas
- Estruturação de backend em Go

---

# 🎯 Objetivo do projeto

Projeto desenvolvido com foco em aprofundar conhecimentos em backend com Go, aplicando conceitos de arquitetura, autenticação, organização em camadas e construção de APIs REST.

Durante o desenvolvimento, foram trabalhados aspectos como:

- Estruturação de aplicações backend
- Fluxo de autenticação com JWT
- Integração com PostgreSQL
- Organização de código em camadas
- Tratamento de erros
- Construção de APIs seguras e organizadas

---

# 👨‍💻 Autor

Alexander Leal Pastana

- GitHub: https://github.com/alexander-pastana
- LinkedIn: https://www.linkedin.com/in/alexanderpastana/
