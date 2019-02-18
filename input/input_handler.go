package input

import "strings"
import "github.com/reisraff/goinput/input/nodes"
import "github.com/reisraff/goinput/input/interfaces"
import "reflect"

type InputResult struct {
    root interfaces.NodeInterface
    output interface{}
    errors []string
}

func (self * InputResult) Configure(root interfaces.NodeInterface) {
    self.root = root
}

func (self * InputResult) Add(key string, _type interface{}, options map[string]interface{}) interfaces.NodeInterface {
    node, err := self.root.Add(key, _type, options)

    if err != nil {
        self.errors = append(self.errors, err.Error())
    }

    return node
}

func (self * InputResult) GetData(index string) interface{} {
    var result interface{}
    result = self.output.(map[string]interface{})[index]

    if reflect.TypeOf(result).String() == "reflect.Value" {
        result = result.(reflect.Value).Interface()
    }

    return result
}

func (self * InputResult) IsValid() bool {
    return len(self.errors) == 0
}

func (self * InputResult) GetErrorsAsString() string {
    return strings.Join(self.errors, ", ")
}

type Define func(InputResult)

type InputHandlerInterface interface {
    Configure(interfaces.TypeHandlerInterface)
    Bind(map[string]interface{})
}

type InputHandler struct {
    typeHandler interfaces.TypeHandlerInterface
}

func (self * InputHandler) Configure(typeHandler interfaces.TypeHandlerInterface) {
    self.typeHandler = typeHandler
}

func (self * InputHandler) Bind(input map[string]interface{}, definer Define) InputResult {
    rootNode := nodes.CreateBaseNode()
    rootNode.SetTypeHandler(self.typeHandler)

    result := InputResult{}
    result.Configure(rootNode)

    definer(result)

    result.output = result.root.GetValue("root", result.root.Walk(input, "root"))
    result.errors = self.typeHandler.GetErrors()

    return result
}

