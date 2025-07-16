# Aplicativo de Lista de Compras

## 1. Contextualização do Problema
Muitas pessoas enfrentam dificuldades ao gerenciar suas listas de compras, seja por esquecerem itens importantes ou por não conseguirem acompanhar a organização de suas necessidades de consumo de forma eficiente. Um aplicativo móvel que permita criar e gerenciar listas de compras personalizadas resolve esse problema ao centralizar e simplificar o processo.

## 2. Objetivo
Desenvolver um aplicativo que facilite a criação, edição e gerenciamento de listas de compras em dispositivos móveis, com foco em simplicidade e usabilidade.

## 3. Requisitos do Sistema

### 3.1. Requisitos Funcionais
- **Criação de listas:** O usuário pode criar múltiplas listas de compras.
- **Adicionar itens à lista:** O usuário pode adicionar itens às listas.
- **Organização por categorias:** Itens podem ser organizados por categorias (ex: alimentos, limpeza).
- **Exibição de histórico:** O usuário pode visualizar o histórico de compras feitas.
- **Dicas de economia:** O usuário pode ver dicas para economizar nas compras.

### 3.2. Requisitos Não Funcionais
- **Disponibilidade:** O aplicativo deve estar disponível 99,5% do tempo.
- **Desempenho:** O aplicativo deve carregar listas de compras em menos de 2 segundos.
- **Usabilidade:** O aplicativo deve ser fácil de usar, com uma interface intuitiva.
- **Compatibilidade:** O aplicativo deve funcionar em dispositivos Android e iOS.

## 4. Estrutura do Projeto (Backend)

O backend é desenvolvido em Go, seguindo princípios de Clean Architecture e DDD. As principais pastas são:

- `cmd/api/`: Ponto de entrada da aplicação (main.go).
- `internal/application/usecase/`: Casos de uso da aplicação (regras de negócio).
- `internal/domain/entity/`: Entidades de domínio.
- `internal/domain/repository/`: Interfaces de repositórios.
- `internal/domain/service/`: Serviços de domínio (ex: autenticação JWT).
- `internal/infrastructure/persistence/`: Implementações de persistência (ex: PostgreSQL).
- `internal/interfaces/api/handler/`: Handlers HTTP (controllers).
- `internal/interfaces/api/middleware/`: Middlewares HTTP (ex: autenticação).
- `internal/interfaces/api/router.go`: Definição das rotas da API.

## 5. Como rodar o backend localmente

1. **Pré-requisitos:**
   - Go 1.20+
   - Docker e Docker Compose (opcional, para banco de dados)

2. **Configuração do banco de dados:**
   - Configure as variáveis de ambiente no arquivo `.env` ou exporte no terminal:
     - `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`
   - Você pode usar o `docker-compose.yml` para subir um PostgreSQL local:
     ```bash
     docker-compose up -d
     ```

3. **Rodando a API:**
   ```bash
   go run cmd/api/main.go
   ```
   A API estará disponível em `http://localhost:8080`.

## 6. Tecnologias Utilizadas
- Go (Golang)
- PostgreSQL
- Docker
- JWT para autenticação
- Bcrypt para hash de senhas

---

Se precisar de mais detalhes ou quiser adicionar instruções para o frontend, me avise! 