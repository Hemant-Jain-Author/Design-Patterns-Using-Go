package main

import "fmt"

type Student struct {
    Name string
    ID   int
}

type Model struct {
    st Student
}

func NewModel() *Model {
    return &Model{st: Student{Name: "Harry", ID: 1}}
}

func (m *Model) SetData(name string, id int) {
    fmt.Println("Model: Set data :", name, id)
    m.st.Name = name
    m.st.ID = id
}

func (m *Model) GetData() Student {
    fmt.Println("Model: Get data.")
    return m.st
}

type View struct{}

func (v *View) Update(name string, id int) {
    fmt.Println("View: Student Info :", name, id)
}

type Presenter struct {
    model *Model
    view  *View
}

func NewPresenter() *Presenter {
    return &Presenter{
        model: NewModel(),
        view:  &View{},
    }
}

func (p *Presenter) SetData(name string, id int) {
    fmt.Println("Presenter: Receive data from client.")
    p.model.SetData(name, id)
}

func (p *Presenter) UpdateView() {
    fmt.Println("Presenter: Receive update from client.")
    data := p.model.GetData()
    p.view.Update(data.Name, data.ID)
}

func main() {
    presenter := NewPresenter()
    presenter.UpdateView()

    presenter.SetData("jack", 2)
    presenter.UpdateView()
}

/*
Presenter: Receive update from client.
Model: Get data.
View: Student Info : Harry 1
Presenter: Receive data from client.
Model: Set data : jack 2
Presenter: Receive update from client.
Model: Get data.
View: Student Info : jack 2
*/