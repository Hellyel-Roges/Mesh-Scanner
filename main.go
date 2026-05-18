package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

// Paleta de cores ANSI para formatação do terminal
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Cyan   = "\033[36m"
	Bold   = "\033[1m"
)

// escanearPorta tenta estabelecer uma conexão TCP e capturar o banner do serviço
func escanearPorta(ip string, porta int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()

	endereco := net.JoinHostPort(ip, strconv.Itoa(porta))
	conexao, err := net.DialTimeout("tcp", endereco, 1*time.Second)
	if err != nil {
		return
	}
	defer conexao.Close()

	// Timeout para leitura de resposta (Banner Grabbing)
	conexao.SetReadDeadline(time.Now().Add(1 * time.Second))
	leitor := bufio.NewReader(conexao)
	banner, err := leitor.ReadString('\n')

	// Bloqueia o Mutex para garantir que a saída no terminal seja Thread-safe
	mu.Lock()
	defer mu.Unlock()

	fmt.Printf("%s[+] Porta %-4d ABERTA%s ", Green+Bold, porta, Reset)

	if err == nil {
		fmt.Printf("➜ %sServiço: %s%s", Cyan, banner, Reset)
	} else {
		fmt.Printf("➜ %s(Silencioso)%s\n", Yellow, Reset)
	}
}

func main() {
	// Configuração dos argumentos de linha de comando (CLI)
	alvo := flag.String("alvo", "127.0.0.1", "Endereço IP ou Domínio para escanear")
	portaInicio := flag.Int("inicio", 1, "Porta inicial")
	portaFim := flag.Int("fim", 1024, "Porta final")
	flag.Parse()

	fmt.Printf("\n%s==================================================%s\n", Cyan+Bold, Reset)
	fmt.Printf("%s     SENTINEL MESH-SCANNER v2.0 %s\n", Cyan+Bold, Reset)
	fmt.Printf("%s==================================================%s\n", Cyan+Bold, Reset)
	fmt.Printf(" %sAlvo:%s %s\n", Bold, Reset, *alvo)
	fmt.Printf(" %sRange:%s %d até %d\n", Bold, Reset, *portaInicio, *portaFim)
	fmt.Printf("%s--------------------------------------------------%s\n\n", Cyan+Bold, Reset)

	inicio := time.Now()
	var wg sync.WaitGroup
	var mu sync.Mutex

	// Escaneamento concorrente das portas
	for porta := *portaInicio; porta <= *portaFim; porta++ {
		wg.Add(1)
		go escanearPorta(*alvo, porta, &wg, &mu)
	}

	wg.Wait()
	duracao := time.Since(inicio)

	fmt.Printf("\n%s==================================================%s\n", Cyan+Bold, Reset)
	fmt.Printf("%sScan finalizado em: %v%s\n", Green+Bold, duracao, Reset)
}
