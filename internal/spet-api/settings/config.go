package settings

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Структура конфиг файла приложения
type Database struct {
	Name string `json:"Name"`
	URL  string `json:"URL"`
}

type Config struct {
	PathToLogger string   `json:"PathToLogger"`
	Address      string   `json:"Address"`
	Database     Database `json:"Database"`
}

// Создаем или читаем файл с настройками
func NewConfig(c *Config, path string) (*Config, error) {
	// Проверяем файл на существование
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			// Нет, значит создаем
			err := CreateConfig(c, path)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	// Да, читаем данные
	dataJSON, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Парсим файл
	data := &Config{}
	err = json.Unmarshal(dataJSON, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Создание конфиг файла и занесение туда данных по умолчанию
func CreateConfig(c *Config, path string) error {
	// Создаем файл
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil
	}

	// Преобразуем данные в json
	dataJSON, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return err
	}

	// Записываем туда данные
	err = ioutil.WriteFile(path, dataJSON, 0)
	if err != nil {
		return err
	}

	// Закрываем файл
	err = f.Close()
	if err != nil {
		return nil
	}
	return nil
}
