package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {

	handleRequest()

}
func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page) //отслежываем другой переход на юрл
	http.ListenAndServe(":8080", nil)            //start server
}

func contacts_page(page http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(page, "SomeText CONTACKT ")
}

func home_page(page http.ResponseWriter, r *http.Request) { // Request - запрос клиента  //ResponseWriter - для отображения инфы лдя клиента
	bob := User{Name: "Bob", Age: 18, Money: -25, avg_grades: 4.2, happines: 5.0, Hobbies: []string{"Footbal", "Tenis", "Dance", "Drive"}}
	bob.setNewName("ALEX")
	//fmt.Fprintf(page, bob.getInfo()) // если я снач запущю эту фигню то экзекют выведет код html

	templ, err := template.ParseFiles("C:\\Users\\Best\\GolandProjects\\awesomeProject\\GOSHAGolangWeb\\templates\\home_page.html") // с помощю этого метода мы подгрузим html
	if err != nil {
		log.Fatal(err)
	}
	templ.Execute(page, bob) // страница на которой будем выводит инфу и доп значения
}

type User struct {
	Name                 string
	Age                  uint16
	Money                int16
	avg_grades, happines float64
	Hobbies              []string
}

// если структура больше то лутше ставить *
func (u *User) getInfo() string {
	return fmt.Sprintf("User Name is %s. He is %d, "+
		"and he has %d money", u.Name, u.Age, u.Money)
}

// мы передаем здес сам обект а не ссылку на него тоесть мы рабоатем с копией поэтому  используй *
func (u *User) setNewName(Name string) {
	u.Name = Name
}
