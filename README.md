# 🛡️ Sentinel Mesh-Scanner

> Uma ferramenta de auditoria de redes (Port Scanner) rápida, concorrente e desenvolvida em Golang. 

O **Mesh-Scanner** é um módulo de reconhecimento de rede focado em alta performance. Desenvolvido como parte dos estudos de cibersegurança e arquitetura de redes, ele utiliza a engine de concorrência nativa do Go para escanear milhares de portas em segundos, além de tentar extrair o *banner* dos serviços ativos.

## ✨ Funcionalidades (Features)

* 🚀 **Escaneamento Concorrente:** Utiliza *Goroutines* para testar múltiplas portas simultaneamente, reduzindo drasticamente o tempo de varredura.
* 🔒 **Thread-Safe I/O:** Implementação de `sync.Mutex` para garantir que a saída no terminal seja limpa e não sofra *race conditions* durante a impressão simultânea de resultados.
* 🕵️ **Banner Grabbing:** Não apenas verifica se a porta está aberta, mas tenta interagir com o serviço TCP para capturar o seu banner (ex: versão do SSH, FTP, Apache).
* ⚙️ **CLI Customizável:** Argumentos de linha de comando (`flags`) fáceis de usar para definir o alvo e o *range* de portas.
* 🎨 **Interface Colorida:** Saída formatada no terminal com cores ANSI para rápida identificação visual de portas abertas e anomalias.

## 🛠️ Tecnologias Utilizadas

* **Linguagem:** Go (Golang)
* **Pacotes Nativos:** `net` (Sockets TCP), `sync` (WaitGroups e Mutex), `flag` (CLI), `bufio` (Leitura de Buffers).

## 🚀 Como instalar e usar

### Pré-requisitos
Certifique-se de ter o [Go instalado](https://go.dev/doc/install) na sua máquina.

### Clonando o repositório
```bash
git clone [https://github.com/SEU_USUARIO/Mesh-Scanner.git](https://github.com/SEU_USUARIO/Mesh-Scanner.git)
cd Mesh-Scanner
```

### Executando a ferramenta
Pode rodar o scanner diretamente utilizando as *flags* disponíveis. 

**Comando Básico (Escanear localhost das portas 1 a 1024):**
```bash
go run main.go
```

**Comando Avançado (Definir alvo e range específico):**
```bash
go run main.go -alvo 192.168.1.100 -inicio 20 -fim 10000
```

### Argumentos da CLI

| Flag | Descrição | Valor Padrão |
| :--- | :--- | :--- |
| `-alvo` | Endereço IP ou Domínio que será escaneado. | `127.0.0.1` |
| `-inicio` | A porta inicial do escaneamento. | `1` |
| `-fim` | A porta final do escaneamento. | `1024` |

## 💻 Exemplo de Saída

```text
==================================================
     SENTINEL MESH-SCANNER v2.0 
==================================================
 Alvo: 127.0.0.1
 Range: 1 até 1024
--------------------------------------------------

[+] Porta 137  ABERTA ➜ (Silencioso)
[+] Porta 139  ABERTA ➜ (Silencioso)
[+] Porta 445  ABERTA ➜ (Silencioso)

==================================================
Scan finalizado em: 1.234s
```

## ⚠️ Aviso Legal (Disclaimer)
Este projeto foi desenvolvido para **fins estritamente educacionais e acadêmicos**. A ferramenta deve ser utilizada **apenas** em redes e sistemas em que você tenha permissão explícita para auditar. O desenvolvedor não se responsabiliza pelo mau uso desta ferramenta.

---
*Desenvolvido com 🍵🧊 e Go.*
