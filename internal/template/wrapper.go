package template

import (
	"text/template"
	"text/template/parse"
)

type Placeholder struct {
	Name string // field name
	Pos  int    // position in the template
}

type Wrapper struct {
	t *template.Template
}

func NewWrapper(t *template.Template) *Wrapper {
	return &Wrapper{
		t: t,
	}
}

func (t *Wrapper) ExtractPlaceholders() (placeholders []Placeholder) {
	nodes := t.t.Tree.Root.Nodes
	for _, node := range nodes {
		if node.Type() == parse.NodeAction {
			actionNode := node.(*parse.ActionNode)
			pipeline := actionNode.Pipe
			for _, command := range pipeline.Cmds {
				for _, arg := range command.Args {
					if arg.Type() == parse.NodeField {
						fieldNode := arg.(*parse.FieldNode)
						for _, ident := range fieldNode.Ident {
							placeholders = append(placeholders, Placeholder{
								Name: ident,
								Pos:  int(fieldNode.Pos),
							})
						}
					}
				}
			}

		}
	}
	return
}
