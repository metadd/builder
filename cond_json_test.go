package builder

import (
	"encoding/json"
	"errors"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

func condToMap(cond Cond) map[string]interface{} {
	s1, err := json.Marshal(cond)
	So(err, ShouldEqual, nil)

	var m1 map[string]interface{}
	err = json.Unmarshal(s1, &m1)
	So(err, ShouldEqual, nil)

	return m1
}

func jsonTestCond(cond Cond) {
	s1, err := json.Marshal(cond)

	So(err, ShouldEqual, nil)

	var m1 map[string]interface{}
	err = json.Unmarshal(s1, &m1)
	So(err, ShouldEqual, nil)
	c1, err := LoadCond(m1)
	So(err, ShouldEqual, nil)
	cs1, err := json.Marshal(c1)
	fmt.Println("s1  = ", string(s1))
	fmt.Println("cs1 = ", string(cs1))
	So(err, ShouldEqual, nil)
	So(string(s1), ShouldEqual, string(cs1))

}

func getInVal(in interface{}) interface{} {
	iType := reflect.TypeOf(in)
	if iType.Kind() != reflect.Ptr {
		return in
	}

	switch iType.Kind() {
	case reflect.Bool:
		return reflect.ValueOf(in).Bool()
	case reflect.String:
		return reflect.ValueOf(in).String()
	}

	return nil
}

//反射学习
func reflectLoad(inVal reflect.Value, vVal reflect.Value) error {
	vType := vVal.Type()
	vKind := vType.Kind()

	if vKind == reflect.Interface || vKind == reflect.Ptr {
		vVal = vVal.Elem()
	}

	if inVal.Type().Kind() == reflect.Interface || inVal.Type().Kind() == reflect.Ptr {
		inVal = inVal.Elem()
	}
	inKind := inVal.Kind()
	switch vKind {
	case reflect.Bool:
		if inKind != reflect.Bool {
			return errors.New("in should be bool")
		}
		vVal.Set(inVal)
	case reflect.Int:
		if inKind != reflect.Int {
			return errors.New("in need to be Int")
		}
		vVal.Set(inVal)
	case reflect.Int8:
		if inKind != reflect.Int8 {
			return errors.New("in need to be Unt8")
		}
		vVal.Set(inVal)
	case reflect.Int16:
		if inKind != reflect.Int16 {
			return errors.New("in need to be Int16")
		}
		vVal.Set(inVal)
	case reflect.Int32:
		var tmp int32
		if inKind == reflect.Int {
			tmp = int32(inVal.Interface().(int))
		} else if inKind == reflect.Int8 {
			tmp = int32(inVal.Interface().(int8))
		} else if inKind == reflect.Int16 {
			tmp = int32(inVal.Interface().(int16))
		} else if inKind == reflect.Int32 {
			tmp = inVal.Interface().(int32)
		} else {
			return errors.New("in need to be Int/Int8/Int16/INt32")
		}
		vVal.Set(reflect.ValueOf(tmp))
	case reflect.Int64:
		var tmp int64
		if inKind == reflect.Int {
			tmp = int64(inVal.Interface().(int))
		} else if inKind == reflect.Int8 {
			tmp = int64(inVal.Interface().(int8))
		} else if inKind == reflect.Int16 {
			tmp = int64(inVal.Interface().(int16))
		} else if inKind == reflect.Int32 {
			tmp = int64(inVal.Interface().(int32))
		} else if inKind == reflect.Int64 {
			tmp = inVal.Interface().(int64)
		} else {
			return errors.New("in need to be Int/Int8/Int16/INt32/Int64")
		}

		vVal.Set(reflect.ValueOf(tmp))
	case reflect.Uint:
		if inVal.Kind() != reflect.Uint {
			return errors.New("in need to be Uint")
		}
		vVal.Set(inVal)
	case reflect.Uint8:
		if inVal.Kind() != reflect.Uint8 {
			return errors.New("in need to be Uint8")
		}
		vVal.Set(inVal)
	case reflect.Uint16:
		if inVal.Kind() != reflect.Uint16 {
			return errors.New("in need to be Uint16")
		}
		vVal.Set(inVal)
	case reflect.Uint32:
		if inVal.Kind() != reflect.Uint32 {
			return errors.New("in need to be Uint32")
		}
		vVal.Set(inVal)
	case reflect.Uint64:
		if inVal.Kind() != reflect.Uint64 {
			return errors.New("in need to be Uint64")
		}
		vVal.Set(inVal)
		vVal.SetInt(123)
	case reflect.Uintptr:
	case reflect.Float32:
		if inVal.Kind() != reflect.Float32 {
			return errors.New("in need to be Float32")
		}
		vVal.Set(inVal)
	case reflect.Float64:
		if inVal.Kind() != reflect.Float64 {
			return errors.New("in need to be Float64")
		}
		vVal.Set(inVal)
	case reflect.Complex64:
		return errors.New("Do not support Complex64")
	case reflect.Complex128:
		return errors.New("Do not support Complex128")
	case reflect.Slice:
		//直接扩展val Slice到能够容纳in？？
		inLen := inVal.Len()
		if inLen > vVal.Cap() {
			vVal.Set(reflect.MakeSlice(vType, inLen, inLen))
		}
		fallthrough
	case reflect.Array:
		//判断inVal是否是Array或者slice
		if inVal.Kind() != reflect.Array && inVal.Kind() != reflect.Slice {
			return errors.New("in value should be Array or Slice")
		}
		//判断in和val元素类型相同
		if vVal.Type().Elem() != inVal.Type().Elem() &&
			vVal.Type().Elem().Kind() != reflect.Interface {
			fmt.Println("vVal type =", vVal.Type().Elem())
			fmt.Println("inVal type =", inVal.Type().Elem())
			return errors.New("in value type not same to val type")
		}

		//判断in的size是否和val的size是否一致. 如果val是Array，则in必须一样，如果val是Slice，只需要Cap大于等于in
		if vVal.Type().Kind() == reflect.Array && inVal.Len() != vVal.Cap() {
			return errors.New("in value len should little or equals to val len")
		}
		//将in的内容赋值到val
		for i := 0; i < inVal.Len(); i++ {
			fmt.Println("in i=", i, inVal.Index(i))
			vVal.Index(i).Set(inVal.Index(i))
		}
		fmt.Println("inVal.Kind = ", inVal.Kind())
	case reflect.Chan:
		return errors.New("Do not support Chan")
	case reflect.Func:
		return errors.New("Do not support Func")
	case reflect.Interface:
		return errors.New("Do not support Interface")
	case reflect.Map:
		return errors.New("Do not support Map")
	case reflect.Ptr:
		return errors.New("Do not support Ptr")
	case reflect.String:
		//vVal := reflect.ValueOf(val).Elem()
		if inVal.Kind() != reflect.String {
			return errors.New("in should to be string")
		}
		vVal.Set(inVal)
	case reflect.Struct:
		if inVal.Kind() != reflect.Map {
			return errors.New("in should be map")
		}

		for i := 0; i < vType.NumField(); i++ {
			vsf := vType.Field(i)
			//判断in是否有对应的字段
			insfv := inVal.MapIndex(reflect.ValueOf(vsf.Name))
			if insfv.IsZero() { //没有包含对应的字段
				return errors.New("in map should contain key:" + vsf.Name)
			}
			err := reflectLoad(insfv, vVal.Field(i))
			if err != nil {
				return err
			}
			/*
				//判断类型是否可一样， TODO:其实某些情况下不一样也可以赋值，先忽略
				if vsf.Type.Kind() != insfv.Elem().Kind() {
					return errors.New("in filed " + vsf.Name + " should be type " + vsf.Type.Kind().String())
				}
				//将对应的值设置到field上
				//是否会有问题， 递归的调用是否更好？试试结构体套结构体的例子
				vsfValue := vVal.Field(i)
				vsfValue.Set(insfv.Elem())

			*/
		}
	case reflect.UnsafePointer:
		return errors.New("Do not support UnsafePointer")
	default:
		return errors.New("Unkown type")
	}

	return nil
}

func LoadVal(in interface{}, val interface{}) error {
	//或者in的reflect.Value
	inVal := reflect.ValueOf(in)
	if inVal.Kind() == reflect.Ptr {
		inVal = inVal.Elem()
	}

	vType := reflect.TypeOf(val)

	if vType.Kind() != reflect.Ptr {
		return errors.New("val should pointer type")
	}

	fmt.Println("vType = ", vType)
	fmt.Println("vType.kind = ", vType.Kind())

	//判断val指针指向的的类型
	vType = vType.Elem()
	fmt.Println("vType = ", vType)
	fmt.Println("vType.kind = ", vType.Kind())
	vVal := reflect.ValueOf(val).Elem()

	return reflectLoad(inVal, vVal)
}

func TestLoadVal(t *testing.T) {
	Convey("Test Load bool", t, func() {
		Convey("Test Load bool val", func() {
			a := false
			b := true
			err := LoadVal(a, &b)
			So(err, ShouldEqual, nil)
			So(b, ShouldEqual, a)
		})
		Convey("Test Load bool Ptr", func() {
			a := false
			b := true
			err := LoadVal(&a, &b)
			So(err, ShouldEqual, nil)
			So(a, ShouldEqual, false)
			So(b, ShouldEqual, a)
		})
		Convey("Test Load bool Type error", func() {
			a := 1
			b := true
			err := LoadVal(a, &b)
			So(err, ShouldNotEqual, nil)
		})
	})

	Convey("Test Load string", t, func() {
		Convey("Test Load string val", func() {
			a := "xxxx"
			b := ""
			err := LoadVal(a, &b)
			So(err, ShouldEqual, nil)
			So(a, ShouldEqual, b)
		})

		Convey("Test Load string Ptr", func() {
			a := "xxxx"
			b := ""
			err := LoadVal(&a, &b)
			So(err, ShouldEqual, nil)
			So(a, ShouldEqual, "xxxx")
			So(a, ShouldEqual, b)
		})

		Convey("Test Load string Type error", func() {
			a := 1
			b := ""
			err := LoadVal(&a, &b)
			So(err, ShouldNotEqual, nil)
		})
	})

	Convey("Test Load Array", t, func() {
		Convey("Test Load xxx", func() {
			a := []int{12, 3, 3}
			b := [3]int{}
			err := LoadVal(a, &b)
			So(err, ShouldEqual, nil)
			for i := range a {
				So(a[i], ShouldEqual, b[i])
			}
		})

		Convey("Test Load size error", func() {
			a := []int{12, 3}
			b := [3]int{}
			err := LoadVal(a, &b)
			So(err, ShouldNotEqual, nil)
			fmt.Println("err =", err)
		})

		Convey("Test Load base val interface{}  ok", func() {
			a := []string{"12", "3", "4"}
			b := [3]interface{}{}
			err := LoadVal(a, &b)
			So(err, ShouldEqual, nil)
			for i := range a {
				So(a[i], ShouldEqual, b[i])
			}
			fmt.Println("a = ", a)
			fmt.Println("b = ", b)
		})

		Convey("Test Load base Type  error", func() {
			a := []string{"12", "3", "4"}
			b := [3]int{}
			err := LoadVal(a, &b)
			So(err, ShouldNotEqual, nil)
			fmt.Println("err =", err)
		})
	})

	Convey("Test Load Slice", t, func() {
		Convey("Test Load Slice auto extend Slice size", func() {
			a := [3]int{12, 3, 3}
			b := []int{}
			err := LoadVal(a, &b)
			So(err, ShouldEqual, nil)
			for i := range a {
				So(a[i], ShouldEqual, b[i])
			}
		})
	})

	Convey("Tesst Load Struct", t, func() {
		Convey("Test Load Struct success", func() {

			a := map[string]interface{}{"A": 1, "B": "xxx"}
			b := struct {
				A int
				B string
			}{}
			err := LoadVal(a, &b)
			So(err, ShouldEqual, nil)
			aStr, err := json.Marshal(a)
			So(err, ShouldEqual, nil)
			bStr, err := json.Marshal(b)
			So(err, ShouldEqual, nil)
			So(string(aStr), ShouldEqual, string(bStr))
		})

		Convey("Test Load Struct embedded", func() {

			a := map[string]interface{}{"A": 1, "B": map[string]interface{}{"C": "1234", "D": 111}}
			b := struct {
				A int
				B struct {
					C string
					D int64
				}
			}{}
			err := LoadVal(a, &b)
			So(err, ShouldEqual, nil)
			aStr, err := json.Marshal(a)
			So(err, ShouldEqual, nil)
			bStr, err := json.Marshal(b)
			So(err, ShouldEqual, nil)
			So(string(aStr), ShouldEqual, string(bStr))
		})
	})
}

func TestJSON_BASE(t *testing.T) {
	Convey("Test JSON BASE", t, func() {
		var cond Cond

		Convey("BETWEEN", func() {
			cond = Between{"xxx", 1, 4}
			jsonTestCond(cond)
		})

		Convey("IF", func() {
			cond = If(true, Eq{"aaa": 1, "bbb": "123"}, Gte{"ccc": 123})
			jsonTestCond(cond)
		})

		Convey("IN", func() {
			cond = In("xxx", 1, 2, 3, 4)
			jsonTestCond(cond)
		})

		Convey("Like", func() {
			cond = Like{"1", "2"}
			jsonTestCond(cond)
		})

		Convey("NEQ", func() {
			cond = Neq{"aaa": 123, "bbb": "123"}
			jsonTestCond(cond)
		})

		Convey("NOT", func() {
			cond = Not{Eq{"aaa": 123}}
			jsonTestCond(cond)
		})

		Convey("NOTIN", func() {
			cond = NotIn("aaa", 3, 1, 2)
			jsonTestCond(cond)
		})

		Convey("ISNULL", func() {
			cond = IsNull{"xxxx"}
			jsonTestCond(cond)
		})

		Convey("OR", func() {
			cond = Or(Eq{"aaa": 123}, Neq{"bbb": "123"})
			jsonTestCond(cond)
		})
	})
}

func TestJSON_condIf(t *testing.T) {
	Convey("IF abnormal", t, func() {
		m := make(map[string]interface{})
		m["IF"] = make(map[string]interface{})
		_, err := LoadCond(m)
		fmt.Println("err = ", err)
	})
}

func TeestJSON_condbetween(t *testing.T) {
	Convey("BETWEEN abnormal", t, func() {
		m := make(map[string]interface{})

		Convey("BETWEEN abnormal root is not map", func() {
			m["BETWEEN"] = "xxx"
			_, err := LoadCond(m)
			So(err, ShouldEqual, ErrJsonBetweenShouldMap)
		})

		Convey("BETWEEN abnormal without Col attribute", func() {
			sub := make(map[string]interface{})
			m["BETWEEN"] = sub
			_, err := LoadCond(m)
			So(err, ShouldEqual, ErrJsonBetweenShouldHasCol)

			Convey("BETWEEN abnormal Col should be string", func() {
				sub["Col"] = 123
				_, err := LoadCond(m)
				So(err, ShouldEqual, ErrJsonBetweenColShouldString)

				Convey("BETWEEN abnormal without LessVal attribute", func() {
					sub["Col"] = "abc"
					_, err := LoadCond(m)
					So(err, ShouldEqual, ErrJsonBetweenShouldHasLessVal)

					Convey("BETWEEN abnormal without MoreVal attribute", func() {
						sub["LessVal"] = 10
						_, err := LoadCond(m)
						So(err, ShouldEqual, ErrJsonBetweenShouldHasMoreVal)
						Convey("BETWEEN ok", func() {
							sub["MoreVal"] = 10
							_, err := LoadCond(m)
							So(err, ShouldEqual, nil)
						})
					})
				})
			})
		})
	})
}

func TestJSON_condexpr(t *testing.T) {
	Convey("Expr abnormal", t, func() {
		m := make(map[string]interface{})

		Convey("Expr  root is not map", func() {
			m["EXPR"] = "xxx"
			_, err := LoadCond(m)
			So(err, ShouldEqual, ErrJsonExprShouldMap)
		})

		Convey("Expr abnormal without Sql attribute", func() {
			sub := make(map[string]interface{})
			m["EXPR"] = sub
			_, err := LoadCond(m)
			So(err, ShouldEqual, ErrJsonExprShouldHasSql)

			Convey("Expr abnormal Sql should be string", func() {
				sub["Sql"] = 123
				_, err := LoadCond(m)
				So(err, ShouldEqual, ErrJsonExprSqlShouldString)
			})

			Convey("Expr abnormal without Args attribute", func() {
				sub["Sql"] = "xxx"
				_, err := LoadCond(m)
				So(err, ShouldEqual, ErrJsonExprShouldHasArgs)

				Convey("Expr abnormal  Args should be array", func() {
					sub["Args"] = "xxx"
					_, err := LoadCond(m)
					So(err, ShouldEqual, ErrJsonExprArgsShouldArray)

					Convey("Expr abnormal  Args should be array", func() {
						sub["Args"] = []interface{}{1, "123", 3}
						_, err := LoadCond(m)
						So(err, ShouldEqual, nil)
					})
				})
			})
		})
	})

	Convey("EXPR xxyy", t, func() {
		var cond = expr{sql: "xxx = ? & yyy = ?", args: []interface{}{123, "123"}}
		jsonTestCond(cond)
	})
}

func TestJSON_condand(t *testing.T) {
	Convey("AND", t, func() {
		m := make(map[string]interface{})

		Convey("AND abnormal root is not slice", func() {
			m["AND"] = "abc"
			_, err := LoadCond(m)
			So(err, ShouldEqual, ErrJsonAndShouldArray)
		})

		Convey("AND abnormal root slice len 0", func() {
			m["AND"] = []interface{}{}
			_, err := LoadCond(m)
			So(err, ShouldEqual, ErrJsonAndArrayLength)
		})

		Convey("AND ok", func() {
			var cond = And(Eq{"a": 1}, Eq{"b": 1})
			jsonTestCond(cond)
		})
	})
}

func TestJSON_compare(t *testing.T) {
	Convey("Compare abnormal", t, func() {
		m := make(map[string]interface{})
		Convey("Compare LT is not map", func() {
			m["LT"] = "xxx"
			_, err := LoadCond(m)
			So(err, ShouldEqual, ErrJsonCompareShouldMap)
		})

		Convey("Compare LTE is not map", func() {
			m["LTE"] = "xxx"
			_, err := LoadCond(m)
			So(err, ShouldEqual, ErrJsonCompareShouldMap)
		})

		Convey("Compare GT is not map", func() {
			m["GT"] = "xxx"
			_, err := LoadCond(m)
			So(err, ShouldEqual, ErrJsonCompareShouldMap)
		})

		Convey("Compare GTE is not map", func() {
			m["GTE"] = "xxx"
			_, err := LoadCond(m)
			So(err, ShouldEqual, ErrJsonCompareShouldMap)
		})

		Convey("Compare EQ is not map", func() {
			m["EQ"] = "xxx"
			_, err := LoadCond(m)
			So(err, ShouldEqual, ErrJsonCompareShouldMap)
		})

		Convey("LT", func() {
			var cond = Lt{"aaa": 123, "bbb": "123"}
			jsonTestCond(cond)
		})

		Convey("LTE", func() {
			var cond = Lte{"aaa": 123, "bbb": "123"}
			jsonTestCond(cond)
		})

		Convey("GT", func() {
			var cond = Gt{"aaa": 123, "bbb": "123"}
			jsonTestCond(cond)
		})

		Convey("GTE", func() {
			var cond = Gte{"aaa": 123, "bbb": "123"}
			jsonTestCond(cond)
		})

		Convey("EQ", func() {
			var cond = Eq{"aaa": 123, "bbb": "123"}
			jsonTestCond(cond)
		})

		Convey("NEQ", func() {
			var cond = Eq{"aaa": 123, "bbb": "123"}
			jsonTestCond(cond)
		})
	})
}

func TestJSON_condin(t *testing.T) {
	Convey("In abnormal", t, func() {
		m := make(map[string]interface{})
		Convey("IN abnormal root is not map", func() {
			m["IN"] = "xxx"
			_, err := LoadCond(m)
			So(err, ShouldEqual, ErrJsonInShouldMap)

			sub := make(map[string]interface{})
			m["IN"] = sub
			Convey("IN abnormal without Col", func() {
				_, err := LoadCond(m)
				So(err, ShouldEqual, ErrJsonINShouldHasCol)
			})

			Convey("IN abnormal  Col should be string", func() {
				sub["Col"] = 123
				_, err := LoadCond(m)
				So(err, ShouldEqual, ErrJsonINColShouldString)
			})

			Convey("IN abnormal  without Vals", func() {
				sub["Col"] = "aaa"
				_, err := LoadCond(m)
				So(err, ShouldEqual, ErrJsonINShouldHasVals)
			})

			Convey("IN abnormal  vals should be Array", func() {
				sub["Col"] = "aaa"
				sub["Vals"] = "aaa"
				_, err := LoadCond(m)
				So(err, ShouldEqual, ErrJsonINColShouldArray)
			})

			Convey("IN OK int", func() {
				sub["Col"] = "aaa"
				sub["Vals"] = []int{1, 2, 3}
				_, err := LoadCond(m)
				So(err, ShouldEqual, nil)
			})

			Convey("IN OK string", func() {
				sub["Col"] = "aaa"
				sub["Vals"] = []string{"1", "2", "3"}
				_, err := LoadCond(m)
				So(err, ShouldEqual, nil)
			})
		})
	})
}

func TestJSON_condlike(t *testing.T) {
	Convey("Test Like", t, func() {
		m := make(map[string]interface{})

		Convey("Like should be array", func() {
			m["LIKE"] = "xxx"
			_, err := LoadCond(m)
			So(err, ShouldEqual, ErrJsonLikeShouldArray)
		})

		Convey("Like should be array stirng 1", func() {
			m["LIKE"] = []int{1, 2}
			_, err := LoadCond(m)
			So(err, ShouldEqual, ErrJsonLikeShouldArrayString)
		})

		Convey("Like should be array string 2", func() {
			m["LIKE"] = []interface{}{1, "2"}
			_, err := LoadCond(m)
			So(err, ShouldEqual, ErrJsonLikeShouldArrayString)
		})

		Convey("Like should be array ok 1", func() {
			m["LIKE"] = []interface{}{"1", "2"}
			_, err := LoadCond(m)
			So(err, ShouldEqual, nil)
		})

		Convey("Like should be array ok 2", func() {
			m["LIKE"] = []string{"1", "2"}
			_, err := LoadCond(m)
			So(err, ShouldEqual, nil)
		})
	})

}

func TestJSON_condnot(t *testing.T) {
	Convey("Test NOT", t, func() {
		m := make(map[string]interface{})
		Convey("NOT should be Array", func() {
			m["NOT"] = "xxx"
			_, err := LoadCond(m)
			So(err, ShouldEqual, ErrJsonNotShouldInterfaceArray)
		})

		Convey("NOT should be []interface", func() {
			m["NOT"] = []interface{}{12, "12"}
			_, err := LoadCond(m)
			So(err, ShouldEqual, ErrJsonNotArrayShouldLen1)
		})

		Convey("NOT should be Cond map", func() {
			m["NOT"] = []interface{}{12}
			_, err := LoadCond(m)
			So(err, ShouldEqual, ErrJsonNotCondShouldMap)
		})

		Convey("NOT ok", func() {
			lt := make(map[string]interface{})
			lt["LT"] = map[string]interface{}{
				"aaa": 1,
			}
			m["NOT"] = []interface{}{lt}
			_, err := LoadCond(m)
			So(err, ShouldEqual, nil)
		})
	})
}
