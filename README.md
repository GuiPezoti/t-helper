# T-Helper com Bubbletea

## 📦 Instalação

```bash
# 1. Instale as dependências
go get github.com/charmbracelet/bubbletea
go get github.com/charmbracelet/lipgloss

# 2. Rode a aplicação
go run .
```

## 🎯 Estrutura do Projeto

```
t-helper/
├── main.go                  # Entrada da aplicação
├── go.mod
└── internal/
    └── app/
        ├── model.go         # Estado da aplicação (dados)
        ├── update.go        # Lógica (como o estado muda)
        ├── view.go          # Visual (como renderizar)
        └── commands.go      # Comandos disponíveis
```

## 🧠 Como funciona (Arquitetura Elm)

### 1. **Model** (Estado)
Define "o que a aplicação sabe":
- O que o usuário digitou
- Histórico de comandos
- Lista de comandos disponíveis

### 2. **Update** (Lógica)
Define "como as coisas mudam":
- Usuário pressiona tecla → atualiza input
- Usuário pressiona Enter → executa comando
- Comando executa → atualiza output

### 3. **View** (Visual)
Define "como mostrar na tela":
- Renderiza o título
- Mostra o histórico
- Mostra o prompt + input

## 🔄 Fluxo de Execução

```
┌─────────────────────────────────────┐
│  1. Usuário pressiona 'h'           │
└──────────────┬──────────────────────┘
               │
               ▼
┌─────────────────────────────────────┐
│  2. Update() recebe KeyMsg('h')     │
│     Adiciona 'h' ao m.input         │
│     Retorna novo Model              │
└──────────────┬──────────────────────┘
               │
               ▼
┌─────────────────────────────────────┐
│  3. View() renderiza novo Model     │
│     Mostra: "T-helper > h_"         │
└─────────────────────────────────────┘
```

## 📝 Comandos Disponíveis

- `help` - Mostra todos os comandos
- `exit` - Sai da aplicação
- `clear` - Limpa a tela

## 🎨 Adicionando Novos Comandos

Edite `internal/app/commands.go`:

```go
func GetCommands() []Command {
    return []Command{
        // ... comandos existentes ...
        {
            Name:        "meucomando",
            Description: "Descrição do comando",
            Execute:     executeMeuComando,
        },
    }
}

func executeMeuComando(m *Model, args []string) tea.Cmd {
    m.output = append(m.output, "Executando meu comando!")
    return nil
}
```

## 🐛 Dicas de Debug

1. **Ver o que está no Model:**
   - Adicione `fmt.Printf("Model: %+v\n", m)` no Update()

2. **Ver mensagens recebidas:**
   - Adicione `fmt.Printf("Msg: %+v\n", msg)` no início do Update()

3. **Rodar sem AltScreen (para ver prints):**
   - Em `main.go`, remova `tea.WithAltScreen()`

## 🚀 Próximos Passos

1. Adicionar histórico com setas ↑↓
2. Adicionar autocompletar com Tab
3. Adicionar comando `ia` interativo
4. Adicionar cores customizadas