# CRUD de Usuários em Golang 🚀

Bem-vindo ao projeto **CRUD de Usuários** desenvolvido em **Golang**! Este projeto implementa as operações básicas de um CRUD (*Create, Read, Update, Delete*) utilizando o pacote [**Gorilla/mux**](https://github.com/gorilla/mux) para gerenciar as rotas da API.

## 🌐 Tecnologias Utilizadas
- **Golang**: Linguagem principal do projeto.
- **Gorilla/mux**: Gerenciamento de rotas HTTP.
- **JSON**: Formato para comunicação entre cliente e servidor.

## 🔧 Funcionalidades
O sistema permite:
1. ✍️ **Criar** um novo usuário.
2. 🔎 **Listar** todos os usuários.
3. ⚙️ **Atualizar** os dados de um usuário.
4. ❌ **Excluir** um usuário pelo seu ID.


O servidor estará rodando em `http://localhost:3000`.


## 🕹️ Endpoints da API
Abaixo está a lista de endpoints disponíveis:

### 1. Listar Usuários
- **GET** `/usuarios`
- **Resposta**:
```json
[
  {
    "id": 1,
    "nome": "João"
  }
]
```

### 2. Obter Usuário por ID
- **GET** `/usuarios/{id}`
- **Resposta**:
```json
{
  "id": 1,
  "nome": "João"
}
```

### 3. Criar Usuário
- **POST** `/usuarios`
- **Corpo da Requisição**:
```json
{
  "nome": "Maria"
}
```
- **Resposta**:
```json
{
  "id": 2,
  "nome": "Maria"
}
```

### 4. Atualizar Usuário
- **PUT** `/usuarios/{id}`
- **Corpo da Requisição**:
```json
{
  "nome": "Pedro"
}
```
- **Resposta**:
```json
{
  "id": 1,
  "nome": "Pedro"
}
```

### 5. Excluir Usuário
- **DELETE** `/usuarios/{id}`
- **Resposta**:
```text
Usuário 1 excluído com sucesso
```

## 🔄 Exemplo de Requisições com cURL
Aqui estão alguns exemplos práticos para testar os endpoints:

### Criar Usuário
```bash
curl -X POST http://localhost:3000/usuarios -H "Content-Type: application/json" -d '{"nome": "Maria"}'
```

### Atualizar Usuário
```bash
curl -X PUT http://localhost:3000/usuarios/1 -H "Content-Type: application/json" -d '{"nome": "Pedro"}'
```

### Excluir Usuário
```bash
curl -X DELETE http://localhost:3000/usuarios/1
```

## 🚫 Limitações do Projeto
- Os dados são armazenados em memória (não há banco de dados).
- O ID dos usuários é gerado sequencialmente e não é persistido.

## ✨ Melhorias Futuras
- Integração com banco de dados (SQLite, PostgreSQL, etc.).
- Validação mais robusta para entradas de dados.
- Implementação de autenticação.

## ⚖️ Licença
Este projeto está licenciado sob a [Licença MIT](LICENSE).

---

Feito com ❤️ por Thiago Fernandes(https://github.com/ThiagooFX). Contribuições são bem-vindas! 🚀

