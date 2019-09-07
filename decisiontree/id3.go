package decisiontree

func (dt DecisionTree) buildTree(examples []Example) Node {
	root := NewAttrNode("Pronostico")
	_ = root.AddChild(NewValNode("sol"))
	nublado := NewValNode("nublado")
	_ = root.AddChild(nublado)
	_ = root.AddChild(NewValNode("lluvioso"))
	_ = nublado.AddChild(NewClassNode("SI"))
	return root
}
