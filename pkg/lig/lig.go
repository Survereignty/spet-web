package lig

import (
	"fmt"
	"log"
	"os"
)

var (
	// файл c записями логов
	file *os.File
	// логи текущей сессии
	session []string
)

// Create создает файл в который будут записываться логи
func Create(path string) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	file = f
	log.SetOutput(f)
	return nil
}

// Quit для закрытия файла с логами
func Quit() {
	err := file.Close()
	if err != nil {
		Error("File "+file.Name()+" couldn't close", err)
	}
}

// Info для уведомляющих сообщений
func Info(s string) {
	session = append(session, "[INFO]"+s)
	fmt.Println("[INFO]", s)
}

// Warning для предупреждающих сообщений
func Warning(s string) {
	session = append(session, "[WARNING]"+s)
	fmt.Println("[WARNING]", s)
}

// Error для указание явной нарушение работы
func Error(s string, e error) {
	session = append(session, "[ERROR]"+s+"->"+e.Error())
	log.Println("[ERROR]", s, "->", e)
	fmt.Println("[ERROR]", s, "->", e)
}

// Crash для остановки приложения в случаи критичной ошибки
func Crash(s string, e error) {
	session = append(session, "[Crash]"+s+"->"+e.Error())
	fmt.Println("[Crash]", s, "->", e)
	log.Fatal("[Crash]", s, "->", e)
}
