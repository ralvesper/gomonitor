package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const quantidadeMonitoramentos = 3
const intervaloMonitoramento = 5 * time.Second

func main() {

	exibeIntroducao()

	for {

		exibeMenu()

		comando := obtemComando()

		switch comando {
		case 0:
			{
				fmt.Println("Sando do Programa...")
				os.Exit(0)
			}
		case 1:
			{
				iniciarMonitoramento()
			}
		case 2:
			{
				exibirLogs()
			}
		default:
			{
				fmt.Println("Comando inválido!")
				os.Exit(-1)
			}
		}

	}
}

func exibeIntroducao() {
	versao := 1.2
	fmt.Println("Bem vindo ao Programa de Monitoramento")
	fmt.Println("Versão:", versao)
	//fmt.Println("O tipo da variavel versao é", reflect.TypeOf(versao))
	//fmt.Println("O tipo da variavel idade é", reflect.TypeOf(idade))
}

func exibeMenu() {
	fmt.Println("---------------------------")
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
	fmt.Println("---------------------------")
}

func obtemComando() int {
	fmt.Print("Digite o comando: ")
	var comando int
	//fmt.Scanf("%d", &comando)
	fmt.Scan(&comando)
	//time.Sleep(1 * time.Second)
	return comando
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	sites := leSitesDoArquivo()

	for i := 0; i < quantidadeMonitoramentos; i++ {
		fmt.Printf("\nTeste %d/%d\n", i+1, quantidadeMonitoramentos)
		for _, site := range sites {
			testaSite(site)
		}
		time.Sleep(intervaloMonitoramento)
	}

}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
		panic(err)
	}

	fmt.Println("Testando o site:", site, "...")
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com suecesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status:", resp.Status)
		registraLog(site, false)
	}
}

func exibirLogs() {
	fmt.Println("Exibindo Logs...")
	imprimeLogs()
}

func leSitesDoArquivo() []string {

	var sites []string
	arquivo, err := os.Open("sites.txt")
	//arquivo, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		//fmt.Println("Ocorreu um erro:", err)
		panic(err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()

}

func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}
