# Logs in Go (LIG)

Функции записи логов


	// файл c записями логов
	file *os.File
	
	// логи текущей сессии
	session []string

    // Create создает файл в который будут записываться логи
    func Create(path string)
    
    // Quit для закрытия файла с логами
    func Quit() 
    
    // Info для уведомляющих сообщений
    func Info(s string)
    
    // Warning для предупреждающих сообщений
    func Warning(s string)
    
    // Error для указание явной нарушение работы
    func Error(s string, e error)
    
    // Crash для остановки приложения в случаи критичной ошибки
    func Crash(s string, e error)
   
## Пример использования

```go
    // Инициализация Logger
    err := lig.Create("logs.txt")
    if err != nil {
        log.Fatal("[CRASH] Logger isn't init ->", err)
    }
    defer lig.Quit()

    lig.Warning("Warning!")
    lig.Info("Hello world!")

    lig.Error("Oh!!", err)
    lig.Crash("Oh!!", err)
```