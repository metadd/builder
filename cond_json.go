package builder

import (
	"encoding/json"
	"errors"
	"reflect"
)

func noAttrErr(attr string) error {
	return errors.New("Do not has " + attr + " attribute")
}

func attrTypeErr(attr string) error {
	return errors.New("Attrbute " + attr + " type error")
}

func (and condAnd) MarshalJSON() ([]byte, error) {
	tmp := struct {
		AND []Cond //不能直接用condAnd， 会导致 MarshalJSON 递归调用
	}{AND: and}

	return json.Marshal(tmp)
}

func LoadAnd(in interface{}) (Cond, error) {
	ain, ok := in.([]interface{})
	if !ok {
		return nil, ErrJsonAndShouldArray
	}
	if len(ain) == 0 {
		return nil, ErrJsonAndArrayLength
	}
	conds := []Cond{}
	for _, k := range ain {
		if m, ok := k.(map[string]interface{}); ok {
			c, err := LoadCond(m)
			if err != nil {
				return nil, err
			}
			conds = append(conds, c)
		}
	}

	return And(conds...), nil
}

func (between Between) MarshalJSON() ([]byte, error) {
	tmp := struct {
		BETWEEN interface{}
	}{BETWEEN: struct {
		Col     string
		LessVal interface{}
		MoreVal interface{}
	}{between.Col, between.LessVal, between.MoreVal}}

	return json.Marshal(tmp)
}

func checkKey(m map[string]interface{}, key string) bool {
	_, ok := m[key]
	return ok
}

func LoadBetween(in interface{}) (Cond, error) {
	m, ok := in.(map[string]interface{})
	if !ok {
		return nil, ErrJsonBetweenShouldMap
	}
	col, ok := m["Col"]
	if !ok {
		return nil, ErrJsonBetweenShouldHasCol
	}

	colStr, ok := col.(string)
	if !ok {
		return nil, ErrJsonBetweenColShouldString
	}

	lessVal, ok := m["LessVal"]
	if !ok {
		return nil, ErrJsonBetweenShouldHasLessVal
	}

	moreVal, ok := m["MoreVal"]
	if !ok {
		return nil, ErrJsonBetweenShouldHasMoreVal
	}

	return Between{colStr, lessVal, moreVal}, nil
}

func (lt Lt) MarshalJSON() ([]byte, error) {
	tmp := struct {
		LT map[string]interface{}
	}{LT: lt}

	return json.Marshal(tmp)
}

func LoadLt(in interface{}) (Cond, error) {
	m, ok := in.(map[string]interface{})
	if !ok {
		return nil, ErrJsonCompareShouldMap
	}
	return Lt(m), nil
}

func (lte Lte) MarshalJSON() ([]byte, error) {
	tmp := struct {
		LTE map[string]interface{}
	}{LTE: lte}

	return json.Marshal(tmp)
}

func LoadLte(in interface{}) (Cond, error) {
	m, ok := in.(map[string]interface{})
	if !ok {
		return nil, ErrJsonCompareShouldMap
	}
	return Lte(m), nil
}

func (gt Gt) MarshalJSON() ([]byte, error) {
	tmp := struct {
		GT map[string]interface{}
	}{GT: gt}

	return json.Marshal(tmp)
}

func LoadGt(in interface{}) (Cond, error) {
	m, ok := in.(map[string]interface{})
	if !ok {
		return nil, ErrJsonCompareShouldMap
	}
	return Gt(m), nil
}

func (gte Gte) MarshalJSON() ([]byte, error) {
	tmp := struct {
		GTE map[string]interface{}
	}{GTE: gte}

	return json.Marshal(tmp)
}

func LoadGte(in interface{}) (Cond, error) {
	m, ok := in.(map[string]interface{})
	if !ok {
		return nil, ErrJsonCompareShouldMap
	}
	return Gte(m), nil
}

func (eq Eq) MarshalJSON() ([]byte, error) {
	tmp := struct {
		EQ map[string]interface{}
	}{EQ: eq}
	return json.Marshal(tmp)
}

func LoadEq(in interface{}) (Cond, error) {
	m, ok := in.(map[string]interface{})
	if !ok {
		return nil, ErrJsonCompareShouldMap
	}
	return Eq(m), nil
}

func (ex expr) MarshalJSON() ([]byte, error) {
	tmp := struct {
		EXPR interface{}
	}{EXPR: struct { //因为expr的成员是非public的。所以构造一个临时的结构体来构建json
		Sql  string
		Args []interface{}
	}{Sql: ex.sql, Args: ex.args}}
	return json.Marshal(tmp)
}

func LoadExpr(in interface{}) (Cond, error) {
	m, ok := in.(map[string]interface{})
	if !ok {
		return nil, ErrJsonExprShouldMap
	}

	sql, ok := m["Sql"]
	if !ok {
		return nil, ErrJsonExprShouldHasSql
	}

	sqlStr, ok := sql.(string)
	if !ok {
		return nil, ErrJsonExprSqlShouldString
	}

	args, ok := m["Args"]
	if !ok {
		return nil, ErrJsonExprShouldHasArgs
	}

	argsV, ok := args.([]interface{})
	if !ok {
		return nil, ErrJsonExprArgsShouldArray
	}
	return Expr(sqlStr, argsV...), nil
}

func (cif condIf) MarshalJSON() ([]byte, error) {
	tmp := struct {
		IF interface{}
	}{IF: struct { //因为condIf的成员是非public的。所以构造一个临时的结构体来构建json
		Condition bool
		CondTrue  interface{}
		CondFalse interface{}
	}{Condition: cif.condition, CondTrue: cif.condTrue, CondFalse: cif.condFalse}}
	return json.Marshal(tmp)
}

func LoadIf(in interface{}) (Cond, error) {

	m, ok := in.(map[string]interface{})
	if !ok {
		return nil, errors.New("IF should be a map")
	}

	c, ok := m["Condition"]
	if !ok {
		return nil, noAttrErr("Condition")
	}

	condition, ok := c.(bool)
	if !ok {
		return nil, attrTypeErr("Condition")
	}

	cTrue, ok := m["CondTrue"]
	if !ok {
		return nil, noAttrErr("CondTrue")
	}
	cTrueMap, ok := cTrue.(map[string]interface{})
	if !ok {
		return nil, attrTypeErr("CondTrue")
	}
	condTrue, err := LoadCond(cTrueMap)
	if err != nil {
		return nil, err
	}

	cFalse, ok := m["CondFalse"]
	if !ok {
		return nil, noAttrErr("CondFalse")
	}
	cFalseMap, ok := cFalse.(map[string]interface{})
	if !ok {
		return nil, attrTypeErr("CondFalse")
	}
	condFalse, err := LoadCond(cFalseMap)
	if err != nil {
		return nil, err
	}
	return If(condition, condTrue, condFalse), nil

}

func (cin condIn) MarshalJSON() ([]byte, error) {
	tmp := struct {
		IN interface{}
	}{IN: struct { //因为condIf的成员是非public的。所以构造一个临时的结构体来构建json
		Col  string
		Vals []interface{}
	}{Col: cin.col, Vals: cin.vals}}
	return json.Marshal(tmp)
}

func LoadIn(in interface{}) (Cond, error) {
	m, ok := in.(map[string]interface{})
	if !ok {
		return nil, ErrJsonInShouldMap
	}

	c, ok := m["Col"]
	if !ok {
		return nil, ErrJsonINShouldHasCol
	}
	Col, ok := c.(string)
	if !ok {
		return nil, ErrJsonINColShouldString
	}

	v, ok := m["Vals"]
	if !ok {
		return nil, ErrJsonINShouldHasVals
	}
	//Vals, ok := v.([]interface{}) //TODO: 这里用[]interface{}{1,2,3}没问题，用[]int{1,2,3}会报错。应该用另外的办法判断？？
	vType := reflect.TypeOf(v)
	if vType.Kind() != reflect.Array && vType.Kind() != reflect.Slice {
		return nil, ErrJsonINColShouldArray
	}

	vv := reflect.ValueOf(v)
	Vals := make([]interface{}, vv.Len())
	rv := reflect.ValueOf(Vals)
	for i := 0; i < vv.Len(); i++ {
		ve := reflect.ValueOf(v).Index(i)
		rv.Index(i).Set(ve)
	}
	return In(Col, Vals...), nil
}

func (like Like) MarshalJSON() ([]byte, error) {
	tmp := struct {
		LIKE [2]string
	}{LIKE: like}
	return json.Marshal(tmp)
}

func LoadLike(in interface{}) (Cond, error) {
	inVal := reflect.ValueOf(in)
	inType := inVal.Type()

	if inType.Kind() != reflect.Array && inType.Kind() != reflect.Slice {
		return nil, ErrJsonLikeShouldArray
	}

	if inVal.Len() != 2 {
		return nil, ErrJsonLikeShouldLen2
	}

	v0 := inVal.Index(0)
	if v0.Kind() == reflect.Interface {
		v0 = v0.Elem()
	}

	if v0.Kind() != reflect.String {
		return nil, ErrJsonLikeShouldArrayString
	}

	v1 := inVal.Index(1)
	if v1.Kind() == reflect.Interface {
		v1 = v1.Elem()
	}
	if v1.Kind() != reflect.String {
		return nil, ErrJsonLikeShouldArrayString
	}

	var tmp [2]string
	tmpV := reflect.ValueOf(&tmp)
	tmpV.Elem().Index(0).Set(v0)
	tmpV.Elem().Index(1).Set(v1)

	return Like{tmp[0], tmp[1]}, nil
}

func (neq Neq) MarshalJSON() ([]byte, error) {
	tmp := struct {
		NEQ map[string]interface{}
	}{NEQ: neq}
	return json.Marshal(tmp)
}

func LoadNeq(in interface{}) (Cond, error) {
	m, ok := in.(map[string]interface{})
	if !ok {
		return nil, ErrJsonCompareShouldMap
	}
	return Neq(m), nil
}

func (not Not) MarshalJSON() ([]byte, error) {
	tmp := struct {
		NOT [1]Cond
	}{NOT: not}
	return json.Marshal(tmp)
}

func LoadNot(in interface{}) (Cond, error) {
	v, ok := in.([]interface{})
	if !ok {
		return nil, ErrJsonNotShouldInterfaceArray
	}

	if len(v) != 1 {
		return nil, ErrJsonNotArrayShouldLen1
	}

	inMap, ok := v[0].(map[string]interface{})
	if !ok {
		return nil, ErrJsonNotCondShouldMap
	}

	cond, err := LoadCond(inMap)
	if err != nil {
		return nil, err
	}

	return Not{cond}, nil
}

func (cnin condNotIn) MarshalJSON() ([]byte, error) {
	tmp := struct {
		NOTIN interface{}
	}{NOTIN: struct { //因为condIf的成员是非public的。所以构造一个临时的结构体来构建json
		Col  string
		Vals []interface{}
	}{Col: cnin.col, Vals: cnin.vals}}
	return json.Marshal(tmp)
}

func LoadNotIn(in interface{}) (Cond, error) {
	m, ok := in.(map[string]interface{})
	if !ok {
		return nil, errors.New("IF should be a map")
	}

	c, ok := m["Col"]
	if !ok {
		return nil, noAttrErr("Col")
	}
	Col, ok := c.(string)
	if !ok {
		return nil, attrTypeErr("Col")
	}

	v, ok := m["Vals"]
	if !ok {
		return nil, noAttrErr("Vals")
	}
	Vals, ok := v.([]interface{})
	if !ok {
		return nil, attrTypeErr("Vals")
	}
	return NotIn(Col, Vals...), nil
}

func (isnull IsNull) MarshalJSON() ([]byte, error) {
	tmp := struct {
		ISNULL [1]string
	}{ISNULL: isnull}
	return json.Marshal(tmp)
}

func LoadIsNull(in interface{}) (Cond, error) {
	v, ok := in.([]interface{})
	if !ok {
		return nil, errors.New("NOT should []interface{}")
	}

	if len(v) != 1 {
		return nil, errors.New("NOT cond should [1]interface{}")
	}

	vstr, ok := v[0].(string)
	if !ok {
		return nil, errors.New("IsNull should be [1]string")
	}
	return IsNull{vstr}, nil
}

func (or condOr) MarshalJSON() ([]byte, error) {
	tmp := struct {
		OR []Cond
	}{OR: or}

	return json.Marshal(tmp)
}

func LoadOr(in interface{}) (Cond, error) {
	ain, ok := in.([]interface{})
	if !ok {
		return nil, errors.New("And cond should array")
	}
	if len(ain) == 0 {
		return nil, errors.New("And cond should array size > 0") // need return error???
	}
	conds := []Cond{}
	for _, k := range ain {
		if m, ok := k.(map[string]interface{}); ok {
			c, err := LoadCond(m)
			if err != nil {
				return nil, err
			}
			conds = append(conds, c)
		}
	}

	return Or(conds...), nil
}

//将map[string]interface{}转换成Cond
func LoadCond(m map[string]interface{}) (Cond, error) {
	//retCond := NewCond()

	if len(m) != 1 {
		return nil, errors.New("LoadCond Root should has one node")
	}
	for k, v := range m {
		switch k {
		case "AND":
			return LoadAnd(v)
		case "BETWEEN":
			return LoadBetween(v)
		case "LT":
			return LoadLt(v)
		case "LTE":
			return LoadLte(v)
		case "GT":
			return LoadGt(v)
		case "GTE":
			return LoadGte(v)
		case "EQ":
			return LoadEq(v)
		case "EXPR":
			return LoadExpr(v)
		case "IF":
			return LoadIf(v)
		case "IN":
			return LoadIn(v)
		case "LIKE":
			return LoadLike(v)
		case "NEQ":
			return LoadNeq(v)
		case "NOT":
			return LoadNot(v)
		case "NOTIN":
			return LoadNotIn(v)
		case "ISNULL":
			return LoadIsNull(v)
		case "OR":
			return LoadOr(v)
		default:
			return nil, ErrJsonUnknownCond
		}
	}
	return nil, ErrJsonUnknownCond
}
