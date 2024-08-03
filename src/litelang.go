package datamorf

import (
	parser "datamorf/build"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/mohae/deepcopy"
)

type DataMorfVisitor struct {
	parser.DataMorfVisitor
	variables map[string]interface{}
	functions map[string]interface{}
	logs      []string
}

func NewLiteLangVisitor(variables map[string]any) *DataMorfVisitor {
	logs := make([]string, 0)
	return &DataMorfVisitor{variables: variables, logs: logs}
}

func ExecuteCode(
	code string,
	variables map[string]interface{},
) (map[string]interface{}, []string) {

	stream := antlr.NewInputStream(code)
	lexer := parser.NewDataMorfLexer(stream)
	cs := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewDataMorfParser(cs)
	tree := p.Program()
	llvisitor := NewLiteLangVisitor(variables)

	llvisitor.Visit(tree)
	return llvisitor.variables, llvisitor.logs
}

func ExecuteWithReturn(
	code string,
	variables map[string]interface{},
) interface{} {
	test := antlr.NewInputStream(code)
	lexer := parser.NewDataMorfLexer(test)
	cs := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewDataMorfParser(cs)
	tree := p.Program()
	llvisitor := NewLiteLangVisitor(variables)
	return llvisitor.Visit(tree)
}

func (v *DataMorfVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *DataMorfVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	for _, child := range ctx.GetChildren() {
		switch x := child.(type) {
		case *parser.StatementsContext:
			return v.VisitStatements(x)
		}
	}
	return nil
}

func (v *DataMorfVisitor) VisitStatements(ctx *parser.StatementsContext) interface{} {
	var last any
	for _, child := range ctx.GetChildren() {
		switch x := child.(type) {
		case *parser.StatementContext:
			last = v.VisitStatement(x)
		}
	}
	return last
}

func (v *DataMorfVisitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	for _, child := range ctx.GetChildren() {
		switch x := child.(type) {
		case *parser.VariableStatementContext:
			return v.VisitVariableStatement(x)
		case *parser.AssignStatementContext:
			return v.VisitAssignStatement(x)
		case *parser.IfStatementContext:
			return v.VisitIfStatement(x)
		case *parser.ForStatementContext:
			return v.VisitForStatement(x)
		case *parser.UnitStatementContext:
			return v.VisitUnitStatement(x)
		case *parser.ValueStatementContext:
			return v.VisitValueStatement(x)
		}
	}
	return nil
}

func (v *DataMorfVisitor) VisitVariableStatement(ctx *parser.VariableStatementContext) interface{} {
	var value interface{}
	value = nil
	var name string
	value = nil

	if ctx.Value() != nil {
		value = ctx.Value().Accept(v)
	}

	name = ctx.Identifier().GetText()
	v.variables[name] = value
	return nil
}

func (v *DataMorfVisitor) VisitValue(ctx *parser.ValueContext) interface{} {
	if ctx.ObjectLiteral() != nil {
		return ctx.ObjectLiteral().Accept(v)
	}
	if ctx.ArrayLiteral() != nil {
		return ctx.ArrayLiteral().Accept(v)
	}
	if ctx.FunctionCall() != nil {
		return ctx.FunctionCall().Accept(v)
	}
	if ctx.Unit() != nil {
		return ctx.Unit().Accept(v)
	}

	return nil
}

func (v *DataMorfVisitor) VisitObjectLiteral(ctx *parser.ObjectLiteralContext) interface{} {
	objectItems := ctx.AllObjectItem()
	objectMap := make(map[string]interface{})

	for i := range objectItems {
		objectItem := objectItems[i]
		var currentObj map[string]interface{}
		if objectItem.KeyValue() != nil {
			currentObj = objectItem.KeyValue().Accept(v).(map[string]interface{})
			for key, value := range currentObj {
				objectMap[key] = value
			}
		} else if objectItem.Spread() != nil {
			currentObj = objectItem.Spread().Accept(v).(map[string]interface{})
			for key, value := range currentObj {
				objectMap[key] = value
			}
		}
	}

	return objectMap
}

func (v *DataMorfVisitor) VisitSpread(ctx *parser.SpreadContext) interface{} {
	object := ctx.AccessRhs().Accept(v)
	return object
}

func (v *DataMorfVisitor) VisitKeyValue(ctx *parser.KeyValueContext) interface{} {
	object := map[string]interface{}{}
	key := ctx.Key().Accept(v).(string)
	value := ctx.Value().Accept(v)
	object[key] = value
	return object
}

func (v *DataMorfVisitor) VisitArrayLiteral(ctx *parser.ArrayLiteralContext) interface{} {
	values := make([]interface{}, 0)
	for _, valueTree := range ctx.AllValue() {
		value := valueTree.Accept(v)
		values = append(values, value)
	}
	return values
}

func (v *DataMorfVisitor) VisitKey(ctx *parser.KeyContext) interface{} {
	if ctx.StringLiteral() != nil {
		return processStringLiteral(ctx.StringLiteral().GetText())
	}
	panic("No key Found")
}

func (v *DataMorfVisitor) VisitUnitStatement(ctx *parser.UnitStatementContext) interface{} {
	return ctx.Unit().Accept(v)
}

func (v *DataMorfVisitor) VisitUnit(ctx *parser.UnitContext) interface{} {
	if ctx.Constant() != nil {
		return ctx.Constant().Accept(v)
	}

	if ctx.Plus() != nil {
		left := ctx.GetLeft().Accept(v)
		right := ctx.GetRight().Accept(v)
		if isNumber(left) && isNumber(right) {
			return left.(float64) + right.(float64)
		} else if isString(left) && isString(right) {
			return left.(string) + right.(string)
		} else {
			if isString(left) {
				rightBytes, _ := json.Marshal(right)
				return left.(string) + string(rightBytes)
			} else if isString(right) {
				leftBytes, _ := json.Marshal(left)
				return string(leftBytes) + right.(string)
			} else {
				// todo : throw exception here instead
				return nil
			}
		}
	}

	if ctx.Minus() != nil {
		left := ctx.GetLeft().Accept(v)
		right := ctx.GetRight().Accept(v)
		if isNumber(left) && isNumber(right) {
			return left.(float64) - right.(float64)
		} else {
			return nil
		}
	}

	if ctx.MoreThan() != nil {
		left := ctx.GetLeft().Accept(v)
		right := ctx.GetRight().Accept(v)
		if isNumber(left) && isNumber(right) {
			return left.(float64) > right.(float64)
		} else {
			return nil
		}
	}

	if ctx.AccessRhs() != nil {
		return ctx.AccessRhs().Accept(v)
	}
	return nil
}

func (v *DataMorfVisitor) VisitConstant(ctx *parser.ConstantContext) interface{} {
	if ctx.StringLiteral() != nil {
		return processStringLiteral(ctx.StringLiteral().GetText())
	}

	if ctx.DecimalLiteral() != nil {
		num, err := strconv.ParseFloat(ctx.DecimalLiteral().GetText(), 64)
		if err != nil {
			panic("Unable to convert decimal literal: " + ctx.DecimalLiteral().GetText())
		}
		return num
	}

	return nil
}

func (v *DataMorfVisitor) VisitValueStatement(ctx *parser.ValueStatementContext) interface{} {
	if ctx.Value() != nil {
		return ctx.Value().Accept(v)
	}
	return nil
}

func (v *DataMorfVisitor) VisitAccessLhs(ctx *parser.AccessLhsContext) interface{} {
	var currentObject interface{}

	if ctx.AllAccessorLhs() == nil || len(ctx.AllAccessorLhs()) == 0 {
		return v.variables
	}

	if ctx.Identifier() != nil {
		currentObject = v.variables[ctx.Identifier().GetText()]
	}

	for i := 0; i < len(ctx.AllAccessorLhs())-1; i++ {
		childCtx := ctx.AllAccessorLhs()[i]

		if childCtx.Identifier() != nil {
			property := childCtx.Identifier().GetText()
			currentObject = currentObject.(map[string]interface{})[property]
		}

		if childCtx.AccessProperty() != nil {
			if childCtx.AccessProperty().DecimalLiteral() != nil {
				index, _ := strconv.Atoi(childCtx.AccessProperty().DecimalLiteral().GetText())
				currentObject = currentObject.([]interface{})[index]
			}

			if childCtx.AccessProperty().Identifier() != nil {
				property := childCtx.AccessProperty().Identifier().GetText()
				currentObject = currentObject.(map[string]interface{})[property]
			}

			if childCtx.AccessProperty().StringLiteral() != nil {
				property := processStringLiteral(childCtx.AccessProperty().StringLiteral().GetText())
				currentObject = currentObject.(map[string]interface{})[property]
			}
		}
	}

	return currentObject
}

func (v *DataMorfVisitor) VisitFunctionCall(ctx *parser.FunctionCallContext) interface{} {
	functionName := ctx.Identifier().GetText()
	paramVals := ctx.Params().Accept(v).([]interface{})
	if functionName == "copy" {
		return deepcopy.Copy(paramVals[0])
	} else if functionName == "log" {
		var logVal = paramVals[0]
		kind := reflect.TypeOf(paramVals[0]).Kind()
		if kind == reflect.Float64 {
			logVal = strconv.FormatFloat(logVal.(float64), 'f', -1, 64)
		} else if kind == reflect.Map {
			logVal, _ = json.Marshal(logVal)
			logVal = string(logVal.([]byte))
		}
		v.logs = append(v.logs, logVal.(string))
		return nil
	}
	return v.functions[functionName].(*parser.FunctionStatementContext).Accept(v)
}

func (v *DataMorfVisitor) VisitParams(ctx *parser.ParamsContext) interface{} {
	vals := make([]interface{}, 0)
	for i := range ctx.AllUnit() {
		val := ctx.AllUnit()[i].Accept(v)
		vals = append(vals, val)
	}
	return vals
}

func (v *DataMorfVisitor) VisitIfStatement(ctx *parser.IfStatementContext) interface{} {
	conditions := ctx.AllUnit()
	blocks := ctx.AllBlock()
	for i, condition := range conditions {
		conditionEvaluation := condition.Accept(v)
		conditionAsBoolean := conditionEvaluation.(bool)
		if conditionAsBoolean {
			blocks[i].Accept(v)
			return nil
		}
	}

	if len(conditions) < len(blocks) {
		blocks[len(blocks)-1].Accept(v)
		return nil
	}

	return nil
}

func (v *DataMorfVisitor) VisitForStatement(ctx *parser.ForStatementContext) interface{} {
	if ctx.IteratorForParam() != nil {
		iterator := ctx.IteratorForParam().Identifier().GetText()
		parent := ctx.IteratorForParam().AccessRhs().Accept(v).([]interface{})
		for _, x := range parent {
			v.variables[iterator] = x
			ctx.Block().Accept(v)
		}
	}
	return nil
}

func (v *DataMorfVisitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	ctx.Statements().Accept(v)
	return nil
}

func (v *DataMorfVisitor) VisitAccessRhs(ctx *parser.AccessRhsContext) interface{} {
	var currentObject interface{}
	var parentObject interface{}
	parentObject = v.variables
	var parentProperty interface{}

	if ctx.Identifier() != nil {
		parentProperty = ctx.Identifier().GetText()
		currentObject = v.variables[ctx.Identifier().GetText()]
	} else if ctx.FunctionCall() != nil {
		currentObject = ctx.FunctionCall().Accept(v)
	}

	for _, childCtx := range ctx.AllAccessorRhs() {
		if childCtx.Identifier() != nil {
			property := childCtx.Identifier().GetText()
			parentProperty = property
			parentObject = currentObject
			currentObject = currentObject.(map[string]interface{})[property]
		}

		if childCtx.AccessProperty() != nil {
			if childCtx.AccessProperty().DecimalLiteral() != nil {
				index, _ := strconv.Atoi(childCtx.AccessProperty().DecimalLiteral().GetText())
				parentObject = currentObject
				currentObject = currentObject.([]interface{})[index]
			}

			if childCtx.AccessProperty().Identifier() != nil {
				property := childCtx.AccessProperty().Identifier().GetText()
				parentObject = currentObject
				currentObject = currentObject.(map[string]interface{})[property]
			}

			if childCtx.AccessProperty().StringLiteral() != nil {
				property := processStringLiteral(childCtx.AccessProperty().StringLiteral().GetText())
				parentObject = currentObject
				currentObject = currentObject.(map[string]interface{})[property]
			}
		}

		if childCtx.FunctionCall() != nil {
			fnName := childCtx.FunctionCall().Identifier().GetText()
			if fnName == "push" {
				dummyArray := currentObject.([]interface{})
				element := childCtx.FunctionCall().Params().Accept(v).([]interface{})[0]
				currentObject = append(dummyArray, element)
				switch parentPropertyCast := parentProperty.(type) {
				case int:
					parentObject.([]interface{})[parentPropertyCast] = currentObject
				case string:
					parentObject.(map[string]interface{})[parentPropertyCast] = currentObject
				default:
					fmt.Println("Unknown type: " + reflect.TypeOf(parentPropertyCast).Name())
				}
			}
			if fnName == "length" {
				dummyArray := currentObject.([]interface{})
				currentObject = float64(len(dummyArray))
			}
		}
	}

	return currentObject
}

func (v *DataMorfVisitor) VisitAssignStatement(ctx *parser.AssignStatementContext) interface{} {
	value := ctx.Value().Accept(v)
	current := ctx.AccessLhs().Accept(v)
	property := getLastAccessorLhs(ctx.AccessLhs().(*parser.AccessLhsContext))

	switch property := property.(type) {
	case int:
		current.([]interface{})[property] = value
	case string:
		current.(map[string]interface{})[property] = value
	}

	return nil
}

func getLastAccessorLhs(ctx *parser.AccessLhsContext) interface{} {
	if ctx.AllAccessorLhs() == nil || len(ctx.AllAccessorLhs()) == 0 {
		return ctx.Identifier().GetText()
	}

	last := ctx.AllAccessorLhs()[len(ctx.AllAccessorLhs())-1]

	if last.Identifier() != nil {
		return last.Identifier().GetText()
	}

	if last.AccessProperty() != nil {
		if last.AccessProperty().DecimalLiteral() != nil {
			index, _ := strconv.Atoi(last.AccessProperty().DecimalLiteral().GetText())
			return index
		}

		if last.AccessProperty().Identifier() != nil {
			return last.AccessProperty().Identifier().GetText()
		}

		if last.AccessProperty().StringLiteral() != nil {
			return processStringLiteral(last.AccessProperty().StringLiteral().GetText())
		}
	}

	return nil
}

func isNumber(input interface{}) bool {
	_, isInt := input.(int)
	_, isFloat := input.(float64)
	return isInt || isFloat
}

func isString(input interface{}) bool {
	_, isString := input.(string)
	return isString
}

func processStringLiteral(text string) string {
	return text[1 : len(text)-1]
}
