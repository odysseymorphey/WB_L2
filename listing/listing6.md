Что выведет программа? Объяснить вывод программы. Рассказать про внутреннее устройство слайсов и что происходит при передаче их в качестве аргументов функции.

```go
package main
 
import (
    "fmt"
)
 
func main() {
    var s = []string{"1", "2", "3"}
    modifySlice(s)
    fmt.Println(s)
}
 
func modifySlice(i []string) {
    i[0] = "3"
    i = append(i, "4")
    i[1] = "5"
    i = append(i, "6")
}
```
Ответ:

```txt
Вывод: 
[3 2 3]

Слайс это контейнер который передается по указателю, соответственно в функции moidfySlice() содержится тот же указатель, что и в main() и поэтому когда мы меняем в функции значение по индексу, то оно изменится и в главной функции. Затем происходит добавление нового элемента в слайс через append() и так как базовый массив начального слайса имел максимальную вместимость 3, то после append() создасться новый массив с вместимостью 4 и вернется указатель на него, и это будет уже другой слайс, отличный от того, что мы инициализировали в главной функции, поэтому дальнейшие изменения уже никак не отразятся на слайсе в main()
```