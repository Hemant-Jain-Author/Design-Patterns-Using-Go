package main

import "fmt"

type Model struct {
    data string
}

func NewModel() *Model {
    return &Model{data: "Hello, World!"}
}

func (m *Model) SetData(data string) {
    fmt.Println("Model: Set data :", data)
    m.data = data
}

func (m *Model) GetData() string {
    fmt.Println("Model: Get data.")
    return m.data
}

type View struct{}

func (v *View) Update(data string) {
    fmt.Println("View: Updating the view with data: ", data)
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

func (p *Presenter) SetData(data string) {
    fmt.Println("Presenter: Receive data from client.")
    p.model.SetData(data)
}

func (p *Presenter) UpdateView() {
    fmt.Println("Presenter: Receive update view from client.")
    data := p.model.GetData()
    p.view.Update(data)
}

func main() {
    fmt.Println("Client: Pass trigger to Presenter.")
    presenter := NewPresenter()
    presenter.UpdateView()

    presenter.SetData("Hello, Students!")
    presenter.UpdateView()
}

/*
Client: Pass trigger to Presenter.
Presenter: Receive update view from client.
Model: Get data.
View: Updating the view with data:  Hello, World!
Presenter: Receive data from client.
Model: Set data : Hello, Students!
Presenter: Receive update view from client.
Model: Get data.
View: Updating the view with data:  Hello, Students!
*/