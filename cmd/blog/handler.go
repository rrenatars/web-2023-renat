package main

import (
	"net/http"
	"text/template"
	"log"
)

func index(w http.ResponseWriter, r *http.Request) { // Функция для отдачи страницы
	ts, err := template.ParseFiles("pages/index.html") // Главная страница блога
	if err != nil {
		http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
		log.Println(err.Error()) // Используем стандартный логгер для вывода ошибки в консоль
		return // Не забываем завершить выполнение ф-ии
	}
 
   // Подготовим данные для шаблона
    data := struct { // Создаем структуру данных для шаблона. По смыслу структура - аналог Record в паскале
	Title string // Заголовок страницы
	Subtitle string // Подзаголовок страницы
	}{
		Title: "BLog", // Заполняем заголовок
		Subtitle: "Best blog", // Заполняем подзаголовок
	}
 
 
	err = ts.Execute(w, data) // Запускаем шаблонизатор для вывода шаблона в тело ответа
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
 }
 