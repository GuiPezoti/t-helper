# T-Helper - Cobra + BubbleTea (Separação Perfeita)

## 🎯 Arquitetura

**Separação clara de responsabilidades:**
- 🐍 **Cobra**: Gerencia **TODOS** os comandos e sua lógica
- 🧋 **BubbleTea**: Gerencia **APENAS** a UI interativa

## 📁 Estrutura do Projeto

```
T-HELPER/
├── cmd/                    # 🐍 COBRA - Lógica de Comandos
│   ├── root.go            # Comando raiz
│   ├── add.go             # Comando: add
│   ├── list.go            # Comando: list
│   ├── clear.go           # Comando: clear
│   └── exit.go            # Comando: exit
├── internal/
│   └── app/               # 🧋 BUBBLETEA - UI apenas
│       ├── model.go       # Estado da UI
│       ├── update.go      # Lógica de input/eventos
│       └── view.go        # Renderização
├── main.go                
├── go.mod
└── go.sum
```

## 🔄 Como Funciona

```
Usuario digita "add task"
         ↓
   BubbleTea captura input
         ↓
   Passa para Cobra executar
         ↓
   Cobra processa o comando
         ↓
   Retorna output como string
         ↓
   BubbleTea exibe na tela
```

## 🚀 Uso

### Instalar Cobra

```bash
go get -u github.com/spf13/cobra@latest
```

### Compilar e Executar

```bash
go build -o t-helper
./t-helper
```

### Comandos Disponíveis

```
T-helper > help                    # Lista comandos (Cobra nativo)
T-helper > add Buy groceries       # Adiciona item
T-helper > add "Fix bug" -p high   # Com flags
T-helper > list                    # Lista itens
T-helper > list --all              # Com flags
T-helper > clear                   # Limpa tela
T-helper > exit                    # Sai (ou quit)
```

## 🆕 Adicionar Novo Comando

### 1. Criar arquivo em `cmd/`

```go
// cmd/delete.go
package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
    Use:   "delete [id]",
    Short: "Delete an item",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        id := args[0]
        
        // Sua lógica aqui
        fmt.Printf("✓ Item %s deleted\n", id)
    },
}

func init() {
    rootCmd.AddCommand(deleteCmd)
    
    // Adicione flags se necessário
    deleteCmd.Flags().BoolP("force", "f", false, "Force delete")
}
```

### 2. Pronto!

O comando funciona automaticamente no modo interativo. Você **NÃO** precisa tocar em nada no `internal/app/`.

## 📝 Responsabilidades

### 🐍 Cobra (`cmd/`)
- ✅ Definir comandos
- ✅ Validar argumentos
- ✅ Processar flags
- ✅ Implementar lógica de negócio
- ✅ Retornar output formatado

### 🧋 BubbleTea (`internal/app/`)
- ✅ Capturar teclas
- ✅ Gerenciar input do usuário
- ✅ Chamar comandos Cobra
- ✅ Exibir output
- ✅ Renderizar UI bonita

## 🎨 Comandos Especiais

Alguns comandos precisam interagir com a UI:

### `exit` / `quit`
```go
// cmd/exit.go
fmt.Println("__EXIT__")  // BubbleTea reconhece e sai
```

### `clear`
```go
// cmd/clear.go
fmt.Println("__CLEAR_SCREEN__")  // BubbleTea limpa o histórico
```

## ✨ Vantagens desta Arquitetura

| Aspecto | Vantagem |
|---------|----------|
| **Organização** | Comandos no Cobra, UI no BubbleTea |
| **Manutenção** | Cada parte é independente |
| **Testabilidade** | Comandos podem ser testados isoladamente |
| **Escalabilidade** | Fácil adicionar comandos sem tocar na UI |
| **Help** | Cobra gera automaticamente |
| **Flags** | Cobra gerencia validation |
| **Reutilização** | Comandos podem ser usados fora da UI |

## 🔧 Exemplo Completo

### Comando com Flags

```go
// cmd/add.go
var addCmd = &cobra.Command{
    Use:   "add [item]",
    Short: "Add a new item",
    Args:  cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        item := strings.Join(args, " ")
        priority, _ := cmd.Flags().GetString("priority")
        
        fmt.Printf("✓ Item added: %s\n", item)
        if priority != "" {
            fmt.Printf("  Priority: %s\n", priority)
        }
    },
}

func init() {
    rootCmd.AddCommand(addCmd)
    addCmd.Flags().StringP("priority", "p", "normal", "Priority level")
}
```

### Uso no Terminal

```
T-helper > add "Review PR" -p high

$ add Review PR -p high
✓ Item added: Review PR
  Priority: high

T-helper > 
```

## 📦 Dependências

```go
require (
    github.com/charmbracelet/bubbletea v0.25.0
    github.com/charmbracelet/lipgloss v0.9.1
    github.com/spf13/cobra v1.8.0
)
```

## 🎯 Próximos Passos

1. **Persistência**: Adicionar storage nos comandos
2. **Estado compartilhado**: Service layer entre Cobra e BubbleTea
3. **Mais comandos**: `update`, `complete`, `search`
4. **Validações**: Adicionar no Cobra
5. **Testes**: Testar comandos isoladamente

## 💡 Por que Esta Arquitetura?

### ✅ Você mantém:
- Comandos organizados (Cobra)
- UI bonita (BubbleTea)
- Help automático (Cobra)
- Flags validation (Cobra)

### ✅ Você ganha:
- Separação clara
- Fácil de testar
- Fácil de manter
- Fácil de escalar

### ✅ Você evita:
- Misturar lógica de negócio com UI
- Duplicação de código
- Comandos hardcoded
