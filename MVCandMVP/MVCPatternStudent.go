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
    fmt.Println("Model: Set data:", name, id)
    m.st.Name = name
    m.st.ID = id
}

func (m *Model) GetData() Student {
    fmt.Println("Model: Get data.")
    return m.st
}

type View struct {
    model *Model
}

func NewView(model *Model) *View {
    return &View{model: model}
}

func (v *View) Update() {
    st := v.model.GetData()
    fmt.Println("View: Student Info :", st.Name, st.ID)
}

type Controller struct {
    model *Model
    view  *View
}

func NewController() *Controller {
    model := NewModel()
    view := NewView(model)
    return &Controller{model: model, view: view}
}

func (c *Controller) SetData(name string, id int) {
    fmt.Println("Controller: Receive data from client.")
    c.model.SetData(name, id)
}

func (c *Controller) UpdateView() {
    fmt.Println("Controller: Receive update view from client.")
    c.view.Update()
}

func main() {
    controller := NewController()
    controller.UpdateView()

    controller.SetData("jack", 2)
    controller.UpdateView()
}

/*
Controller: Receive update view from client.
Model: Get data.
View: Student Info : Harry 1
Controller: Receive data from client.
Model: Set data: jack 2
Controller: Receive update view from client.
Model: Get data.
View: Student Info : jack 2
*/
