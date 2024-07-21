package main

import (
    "bufio"
    "bytes"
    "fmt"
    "os"
)

// Функция mask заменяет содержимое URL с префиксом "http://" на звёздочки
func mask(input string) string {
    strIn := []byte(input) // Переводим входящую строку в строку байтов.
    findHttp := []byte("http://") // Строка-префикс "http://"
    httpLength := len(findHttp) // Длина этого префикса "http://"
    result := make([]byte, 0, len(strIn)) // Проинициализировали новый байтовый срез, аллоцировав под него память, куда будем складывать байты для выходной строки

    for i := 0; i < len(strIn); {
        // Если текущий индекс позволяет проверить наличие префикса "http://" и он найден
        if i <= len(strIn)-httpLength && bytes.Equal(strIn[i:i+httpLength], findHttp) {
            result = append(result, findHttp...) // Добавляем префикс "http://" в результат
            i += httpLength // Перемещаем индекс на длину префикса "http://"
            // Идем по символам ссылки до пробела или конца строки.
            for i < len(strIn) && strIn[i] != ' ' {
                result = append(result, '*') // Заменяем каждый символ ссылки на '*'.
                i++
            }
        } else {
            result = append(result, strIn[i]) // Добавляем текущий символ из strIn в result.
            i++
        }
    }

    return string(result) // Преобразуем байтовый срез result обратно в строку и возвращаем
}

func main() {
    reader := bufio.NewReader(os.Stdin) // Создаем новый reader для чтения ввода с консоли
    fmt.Print("Введите строку: ")
    input, _ := reader.ReadString('\n') // Читаем строку до символа новой строки
    masked := mask(input) // Применяем функцию mask к введенной строке
    fmt.Println(masked)
}
