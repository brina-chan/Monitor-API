# Monitor de Disponibilidade de APIs

![Go Version](https://img.shields.io/badge/Go-1.18%2B-blue?logo=go)
![License](https://img.shields.io/badge/License-MIT-green)

Esta é uma ferramenta de linha de comando (CLI) desenvolvida em **Go** para monitorar a disponibilidade de múltiplas URLs simultaneamente. O projeto foca em demonstrar o uso prático de concorrência nativa para otimizar operações de rede de alta performance.

##  Funcionalidades

- **Processamento Paralelo:** Verificação simultânea de URLs utilizando *Goroutines* e *Channels*.
- **Sincronização Segura:** Controle do ciclo de vida das execuções com *WaitGroups*.
- **Resiliência de Protocolo:** Normalização automática de endereços (garante o prefixo `https://`).
- **Métricas de Resposta:** Relatório detalhado com Status HTTP e latência em milissegundos.
- **Qualidade de Código:** Suíte de testes unitários integrada para validação de lógica e erros.

##  Tecnologias e Dependências

- **Linguagem:** Go
- **Bibliotecas Padrão:**
  - `net/http`: Comunicação e requisições.
  - `sync`: Sincronização e concorrência.
  - `flag`: Interface de linha de comando.
  - `testing` e `httptest`: Infraestrutura de testes.

##  Estrutura do Projeto

```text
├── monitor.go       # Lógica principal da CLI e concorrência
├── monitor_test.go  # Testes unitários e mocks de servidor
├── go.mod           # Definição do módulo
└── README.md        # Documentação do projeto
```

##  Como Utilizar

### 1. Pré-requisitos
Certifique-se de ter o Go 1.18 ou superior instalado em sua máquina.

### 2. Instalação e Configuração
Clone o repositório e inicialize o módulo:

```bash
git clone [https://github.com/brina-chan/Monitor-API.git](https://github.com/brina-chan/Monitor-API.git)
cd Monitor-API
go mod init monitorapi
go mod tidy
```

### 3. Execução
Para testar a disponibilidade de sites, utilize a flag `-urls` com os endereços separados por vírgula. 

**No Windows (PowerShell), utilize aspas duplas:**
```powershell
go run monitor.go -urls="google.com,github.com,uol.com.br"
```

### 4. Executando Testes
Para validar a integridade das funções e o tratamento de erros:
```bash
go test -v
```

##  Exemplo de Saída
```text
Iniciando verificação de 3 endereços...

🟢 [https://google.com](https://google.com)             | Status: 200 | Tempo: 145ms
🟢 [https://github.com](https://github.com)             | Status: 200 | Tempo: 210ms
❌ [https://site-invalido.com.br](https://site-invalido.com.br)   | ERRO: lookup site-invalido.com.br: no such host
```

## ⚖️ Licença
Este projeto está sob a licença MIT. Consulte o arquivo [LICENSE](LICENSE) para mais detalhes.

---
Desenvolvido por [Sabrina Tavares](https://github.com/brina-chan) 