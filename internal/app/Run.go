package handlers

import (
	"context"

	"groupie-tracker/internal/router"
	"groupie-tracker/internal/server"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	port = "7000"
)

func Run() {
	// Подготовка сервера

	router := router.NewRouter()
	srv := &server.Server{}

	// Запуск сервера
	//go func() { ... }() - запускает сервер в отдельной горутине для того, чтобы основная горутина могла продолжить выполнение и ждать сигналы.
	go func() {
		if err := srv.Run(port, router); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error occured while running HTTP server: %s", err.Error())
		}
	}()

	//Graceful Shutdowdn
	// Graceful Shutdown важен для обеспечения стабильной работы сервера и избегания потери данных при его остановке.
	// Вместо простого прерывания работы сервера, Graceful Shutdown позволяет завершить обработку текущих запросов и корректно освободить ресурсы перед завершением работы.

	//os.Sygnal это тип, представляющий сигналы операционной системы.
	signChan := make(chan os.Signal, 1) //создаем канал для отлавливанием сигнала от ОП

	//регистрирует сигналы SIGINT и SIGTERM для обработки.
	//регистрирует сигналы SIGINT (Ctrl+C) и SIGTERM. Когда любой из этих сигналов будет получен, он будет отправлен в канал signChan.
	signal.Notify(signChan, syscall.SIGINT, syscall.SIGTERM)
	log.Println("Server is running. Press Ctrl+C to shut down.")
	//sig := <-signChan - ожидает получения сигнала из канала signChan
	//ожидает, пока сигнал не будет получен. Как только сигнал будет получен, выполнение кода продолжится.
	sig := <-signChan

	log.Println("Recieved terminate signal, graceful shutdown", sig)

	//ctx используется для установки времени, в течение которого сервер должен завершить текущие операции.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	//создает контекст с таймаутом 3 минуты для грациозного завершения сервера.
	//функция cancel будет вызвана после завершения Run, что гарантирует освобождение ресурсов, связанных с контекстом.
	defer cancel()

	if err := srv.Stop(ctx); err != nil {
		// выполняет грациозное завершение сервера с использованием созданного контекста.
		log.Fatalf("Server forced to shutdownЖ %s", err)
	}
	log.Println("Server exiting")
}
