# EstudosGO 🚀

API REST simples em Go para gerenciamento de usuários, utilizando Gin Framework e armazenamento em JSON.

## 📋 Sobre o Projeto

Este é um projeto de estudos de Go (Golang) que implementa uma API REST básica para consulta de usuários. O projeto utiliza uma arquitetura limpa com separação em camadas (controller, repository, model).

## 🛠️ Tecnologias Utilizadas

- **Go 1.25.1**
- **Gin Framework** - Framework web rápido e minimalista
- **JSON** - Armazenamento de dados em arquivo

## 📁 Estrutura do Projeto

```
EstudosGO/
├── cmd/
│   └── main.go              # Ponto de entrada da aplicação
├── internal/
│   ├── controller/
│   │   └── user.go          # Handlers HTTP
│   ├── model/
│   │   └── user.go          # Definição das estruturas de dados
│   └── repository/
│       └── user.go          # Camada de acesso aos dados
├── users.json               # Banco de dados (arquivo JSON)
├── Dockerfile               # Configuração Docker
├── go.mod                   # Dependências do projeto
└── go.sum                   # Checksums das dependências
```

## 🚀 Como Executar

### Pré-requisitos

- Go 1.25.1 ou superior instalado
- Git

### Instalação

1. Clone o repositório:
```bash
git clone https://github.com/MarkHiarley/EstudosGO.git
cd EstudosGO
```

2. Instale as dependências:
```bash
go mod download
```

3. Execute a aplicação:
```bash
go run cmd/main.go
```

A API estará disponível em `http://localhost:8080`

## 📡 Endpoints

### Listar Usuários

```http
GET /users
```

**Resposta de Sucesso (200 OK):**
```json
{
  "data": [
    {
      "id": 1,
      "name": "Alice Souza",
      "email": "alice@example.com",
      "age": 25
    },
    {
      "id": 2,
      "name": "Bruno Lima",
      "email": "bruno@example.com",
      "age": 30
    }
  ]
}
```

**Resposta de Erro (500 Internal Server Error):**
```json
{
  "error": "descrição do erro"
}
```

## 🐳 Docker

### Construir a imagem:
```bash
docker build -t estudosgo .
```

### Executar o container:
```bash
docker run -p 8080:8080 estudosgo
```

## 📚 Aprendizados

Este projeto foi desenvolvido para praticar:
- Estruturação de projetos Go
- Uso do Gin Framework
- Arquitetura em camadas (MVC)
- Manipulação de JSON em Go
- Tratamento de erros
- APIs RESTful

## 🤝 Contribuindo

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou enviar pull requests.

## 📝 Licença

Este projeto é de código aberto e está disponível para fins educacionais.

## 👤 Autor

**Mark Hiarley**
- GitHub: [@MarkHiarley](https://github.com/MarkHiarley)

---

⭐ Se este projeto te ajudou nos estudos, considere dar uma estrela!
