package main

import (
	"fmt"
	"os"
)

func main() {

	exibeIntroducao()
	exebirMenu()

	comando := leComando()

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do programa")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando")
		os.Exit(-1)
	}

}

func exibeIntroducao() {
	nome := "Bruno"
	versao := 1.1
	fmt.Println("Ola, sr.", nome)
	fmt.Println("Está programa está na versão", versao)
}

func exebirMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}
