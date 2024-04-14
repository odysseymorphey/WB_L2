package main

import "fmt"

type MultimediaConverter struct{}

func (m *MultimediaConverter) ConvertVideo(filename, format string) {
	fmt.Printf("Конвертируем видео файл %s в формат %s\n", filename, format)
}

func (m *MultimediaConverter) ConvertAudio(filename, format string) {
	fmt.Printf("Конвертируем аудио файл %s в формат %s\n", filename, format)
}

func main() {
	converter := MultimediaConverter{}
	converter.ConvertVideo("video.mp4", "avi")
	converter.ConvertAudio("audio.wav", "mp3")
}

// Паттерн фасад- это структурный паттерн проектирования,
// который предоставляет простой интерфейс к сложной системе классов,
// библиотеке или фреймворку. Фасад позволяет скрыть сложность системы
// и предоставляет упрощенный интерфейс для взаимодействия с ней.
