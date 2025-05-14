package main

import (
	"fmt"
	"net"
	"sort"
	"strconv"
	"sync"
	"time"
)

func worker(host string, ports <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for p := range ports {
		address := fmt.Sprintf("%s:%d", host, p)
		conn, err := net.DialTimeout("tcp", address, 1*time.Second)
		if err == nil {
			conn.Close()
			results <- p
		} else {
			results <- 0
		}
	}
}

func main() {
	const workerCount = 100
	startTime := time.Now()

	// Запит на введення хоста та кількості портів
	fmt.Print("🖥️ Введіть хост для сканування: ")
	var host string
	fmt.Scanln(&host)

	fmt.Print("🔢 Введіть кількість портів для сканування (наприклад, 1024): ")
	var maxPorts string
	fmt.Scanln(&maxPorts)

	numPorts, err := strconv.Atoi(maxPorts)
	if err != nil || numPorts <= 0 {
		fmt.Println("❌ Некоректне число портів:", err)
		return
	}

	// Повідомлення про початок сканування
	fmt.Println("🚀 Сканування портів...")

	// Канали для комунікації між горутинами
	ports := make(chan int, workerCount)
	results := make(chan int, numPorts)

	var wg sync.WaitGroup
	var openPorts []int

	// Стартуємо воркери
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(host, ports, results, &wg)
	}

	// Додаємо задачі на сканування портів
	go func() {
		for i := 1; i <= numPorts; i++ {
			ports <- i
		}
		close(ports) // Закриваємо канал після додавання всіх портів
	}()

	// Обробка результатів після завершення всіх горутин
	go func() {
		for port := range results {
			if port != 0 {
				openPorts = append(openPorts, port)
			}
		}
	}()

	// Чекаємо завершення всіх горутин
	wg.Wait()

	// Закриваємо канал results після завершення роботи з ним
	close(results)

	// Обчислюємо час виконання
	elapsed := time.Since(startTime)

	// Виводимо відкриті порти
	sort.Ints(openPorts)
	fmt.Println("\n📋 Відкриті порти:")
	if len(openPorts) > 0 {
		for _, port := range openPorts {
			fmt.Printf("  ✅  %d\n", port)
		}
	} else {
		fmt.Println("❌ Не знайдено відкритих портів.")
	}

	// Виведення часу та завершення сканування
	fmt.Printf("\n⏳ Час: %s\n", elapsed.Round(time.Second))
	fmt.Println("\n✅ Сканування завершено.")
}
