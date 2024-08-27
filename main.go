package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"time"
)

func verificarParametros() {
	if len(os.Args) != 3 {
		fmt.Println("Uso: go run main.go <server_list> <downtime_list>")
		os.Exit(1)
	}
}

func criarListaServidores(serverList *os.File) []Server {

	csvReader := csv.NewReader(serverList)
	data, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var servidores []Server
	for i, line := range data {
		if i > 0 {
			servidor := Server{
				serverName: line[0],
				serverURL:  line[1],
			}
			servidores = append(servidores, servidor)
		}
	}
	return servidores
}

type Server struct {
	serverName    string
	serverURL     string
	tempoExecucao float64
	status        int
	dataFalha     string
}

func checkServer(servidores []Server) []Server {
	var downServers []Server

	for _, servidor := range servidores {
		agora := time.Now()
		get, err := http.Get(servidor.serverURL)

		if err != nil {
			fmt.Println("Server %s is down [%s]", servidor.serverName, err.Error())
			servidor.status = 0
			servidor.dataFalha = agora.Format("02/01/2006 15:04:05")
			downServers = append(downServers, servidor)
			continue
		}
		servidor.status = get.StatusCode
		if servidor.status != 200 {
			servidor.dataFalha = agora.Format("02/01/2006 15:04:05")
			downServers = append(downServers, servidor)
		}
		servidor.tempoExecucao = time.Since(agora).Seconds()
		fmt.Printf("Status [%d] Tempo de carga: [%f] URL: [%s]\n", servidor.status, servidor.tempoExecucao, servidor.serverURL)
	}
	return downServers
}

func openFiles(serverListFile string, downTimeFile string) (*os.File, *os.File) {
	serverList, err := os.OpenFile(serverListFile, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	downTimeList, err := os.OpenFile(downTimeFile, os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return serverList, downTimeList
}

func generateDownTime(downTimeList *os.File, downServers []Server) {
	csvWriter := csv.NewWriter(downTimeList)
	for _, servidor := range downServers {
		line := []string{servidor.serverName, servidor.serverURL, servidor.dataFalha, fmt.Sprintf("%f", servidor.tempoExecucao), fmt.Sprintf("%d", servidor.status)}
		csvWriter.Write(line)
	}
	csvWriter.Flush()
}

func main() {

	verificarParametros()

	serverList, downTimeList := openFiles(os.Args[1], os.Args[2])

	defer serverList.Close()
	defer downTimeList.Close()
	servidores := criarListaServidores(serverList)

	downServers := checkServer(servidores)

	generateDownTime(downTimeList, downServers)

}
