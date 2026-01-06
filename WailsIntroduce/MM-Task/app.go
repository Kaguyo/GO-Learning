package main

import (
	"bufio"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) ProvideWorkers(numWorkers string) string {
	num, err := strconv.Atoi(numWorkers)
	if err != nil {
		log.Panic(err)
	}

	createTestFile("test.txt")

	// Modo assíncrono com N workers
	fmt.Printf("\n=== Modo Assíncrono com %d workers ===\n", num)
	timeAsync := a.ReadFileAsync("test.txt", num, true)
	fmt.Printf("Tempo de execução (async): %.4f segundos\n", timeAsync)

	// Modo síncrono
	fmt.Println("\n=== Modo Síncrono ===")
	timeSync := a.ReadFileAsync("test.txt", 0, false)
	fmt.Printf("Tempo de execução (sync): %.4f segundos\n", timeSync)

	speedup := timeSync / timeAsync
	fmt.Printf("\nSpeedup: %.2fx mais rápido\n", speedup)

	return fmt.Sprintf("Async: %.4fs | Sync: %.4fs | Speedup: %.2fx", timeAsync, timeSync, speedup)
}

func (a *App) ReadFileAsync(filePath string, numWorker int, asyncMode bool) float32 {
	startTime := time.Now()

	lines, err := a.readFileLines(filePath)
	if err != nil {
		fmt.Printf("Erro ao ler arquivo: %v\n", err)
		return float32(time.Since(startTime).Seconds())
	}

	fmt.Printf("Arquivo carregado: %d linhas\n", len(lines))

	if asyncMode {
		a.processLinesAsync(lines, numWorker)
	} else {
		a.processLinesSync(lines)
	}

	elapsed := time.Since(startTime)
	return float32(elapsed.Seconds())
}

func (a *App) readFileLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func (a *App) processLinesSync(lines []string) {
	for i, line := range lines {
		a.processLine(i, line)
	}
}

func (a *App) processLinesAsync(lines []string, numWorker int) {
	if numWorker <= 0 {
		numWorker = 1
	}

	jobs := make(chan struct {
		index int
		line  string
	}, len(lines))

	wg := sync.WaitGroup{}

	// Inicia os workers
	for w := 0; w < numWorker; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range jobs {
				a.processLine(job.index, job.line)
			}
		}(w)
	}

	// Envia as linhas para o canal de jobs
	for i, line := range lines {
		jobs <- struct {
			index int
			line  string
		}{index: i, line: line}
	}
	close(jobs)

	wg.Wait()
}

func (a *App) processLine(index int, line string) {
	// Processamento em string para demandar operações de CPU
	result := strings.ToUpper(line)

	// Operações pesadas de hash
	for i := 0; i < 500; i++ {
		hash := sha256.Sum256([]byte(result))
		result = hex.EncodeToString(hash[:])
	}

	sum := 0.0
	for i := 0; i < 50000; i++ {
		sum += float64(i) * 1.5
	}
}

func createTestFile(filename string) {
	file, _ := os.Create(filename)
	defer file.Close()

	// Criar 1000 linhas
	for i := 1; i <= 1000; i++ {
		fmt.Fprintf(file, "Linha %d do arquivo de teste com conteúdo para processamento intensivo\n", i)
	}
}
