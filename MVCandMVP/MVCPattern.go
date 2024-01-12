package main

import "fmt"

type Model struct {
    data string
}

func NewModel() *Model {
    return &Model{data: "Hello, World!"}
}

func (m *Model) SetData(data string) {
    fmt.Println("Model: Set data:", data)
    m.data = data
}

func (m *Model) GetData() string {
    fmt.Println("Model: Get data:", m.data)
    return m.data
}

type View struct {
    model *Model
}

func NewView(model *Model) *View {
    return &View{model: model}
}

func (v *View) Update() {
    data := v.model.GetData()
    fmt.Println("View: Updating the view with data:", data)
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

func (c *Controller) SetData(data string) {
    fmt.Println("Controller: Receive data from client.")
    c.model.SetData(data)
}

func (c *Controller) UpdateView() {
    fmt.Println("Controller: Receive update view from client.")
    c.view.Update()
}

func main() {
    controller := NewController()
    controller.UpdateView()

    controller.SetData("Hello, Students!")
    controller.UpdateView()
}

/*
Controller: Receive update view from client.
Model: Get data: Hello, World!
View: Updating the view with data: Hello, World!
Controller: Receive data from client.
Model: Set data: Hello, Students!
Controller: Receive update view from client.
Model: Get data: Hello, Students!
View: Updating the view with data: Hello, Students!
*/