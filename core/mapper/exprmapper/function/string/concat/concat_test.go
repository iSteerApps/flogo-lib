package concat

import (
	"fmt"
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/core/mapper/exprmapper/expression"
	"github.com/stretchr/testify/assert"
)

var s = &Concat{}

func TestStaticFunc_Concat(t *testing.T) {
	final, err := s.Eval("TIBCO", "FLOGO", "IOT")
	assert.Nil(t, err)
	fmt.Println(final)
	assert.Equal(t, final, "TIBCOFLOGOIOT")
}

func TestExpressionDoubleQuotes(t *testing.T) {
	fun, err := expression.ParseExpression(`string.concat('TIBCO',' Flo"go')`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.Equal(t, `TIBCO Flo"go`, v)
}

func TestExpressionSingleQuote(t *testing.T) {
	fun, err := expression.ParseExpression(`string.concat("TIBCO"," Flo'o\o{go")`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.Equal(t, "TIBCO Flo'o\\o{go", v)
}

func TestExpressionCombine(t *testing.T) {
	fun, err := expression.ParseExpression(`string.concat('Hello', " 'World'")`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.Equal(t, `Hello 'World'`, v)
}

func TestExpressionCombine2(t *testing.T) {
	fun, err := expression.ParseExpression(`string.concat('Hello', ' "World"')`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.Equal(t, `Hello "World"`, v)
}
func TestExpressionNewLine(t *testing.T) {
	fun, err := expression.ParseExpression(`string.concat(
	"TIBCO",
	" FLOGO"
	)`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.Equal(t, "TIBCO FLOGO", v)
}

func TestExpressionSpace(t *testing.T) {
	fun, err := expression.ParseExpression(`string.concat(    "TIBCO"  ,  " FLOGO")   `)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.Equal(t, "TIBCO FLOGO", v)
}

func TestExpressionSpaceNewLineTab(t *testing.T) {
	fun, err := expression.ParseExpression(`string.concat(    "TIBCO" 
		 ,	" FLOGO"	
		 )`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.Equal(t, "TIBCO FLOGO", v)
}

func TestExpressionDoubleDoubleQuotes(t *testing.T) {
	fun, err := expression.ParseExpression(`string.concat("\"abc\"", "dddd")`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.Equal(t, `"abc"dddd`, v)
}

func TestExpressionSingleSingleQuote(t *testing.T) {
	fun, err := expression.ParseExpression(`string.concat('\'b\'ac\'', "dddd")`)
	assert.Nil(t, err)
	assert.NotNil(t, fun)
	v, err := fun.Eval()
	assert.Nil(t, err)
	assert.Equal(t, `'b'ac'dddd`, v)
}
