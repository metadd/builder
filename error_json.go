package builder

import "errors"

var (
	//Unknown err
	ErrJsonUnknownCond = errors.New("Unknown Condition type")
	// Add
	ErrJsonAndShouldArray = errors.New("And Cond should be array")
	ErrJsonAndArrayLength = errors.New("And Cond array length must larger than 0")

	// Between
	ErrJsonBetweenShouldMap        = errors.New("Between Cond should be map")
	ErrJsonBetweenShouldHasCol     = errors.New("Between Cond should has \"Col\" attribute")
	ErrJsonBetweenColShouldString  = errors.New("Between attribute \"Col\" should be String")
	ErrJsonBetweenShouldHasLessVal = errors.New("Between Cond should has \"LessVal\" attribute")
	ErrJsonBetweenShouldHasMoreVal = errors.New("Between Cond should has \"MoreVal\" attribute")

	//Lt, Lte, Gt,  Gte ,Eq, NEQ
	ErrJsonCompareShouldMap = errors.New("Compare(LT,LTE,GT,GTE,EQ,NEQ) should be map")

	//Expr
	ErrJsonExprShouldMap       = errors.New("Expr should be map")
	ErrJsonExprShouldHasSql    = errors.New("Expr should has \"Sql\" attribute")
	ErrJsonExprSqlShouldString = errors.New("Expr should has \"Sql\" should be String")
	ErrJsonExprShouldHasArgs   = errors.New("Expr should has \"Args\" attrbitue")
	ErrJsonExprArgsShouldArray = errors.New("Expr should has \"Args\" should be Array")

	//IF
	ErrJsonIfShouldMap = errors.New("If should be map")

	//IN
	ErrJsonInShouldMap       = errors.New("IN should be map")
	ErrJsonINShouldHasCol    = errors.New("IN should has \"Col\" attribute")
	ErrJsonINColShouldString = errors.New("In \"Col\" attribute should be String")
	ErrJsonINShouldHasVals   = errors.New("IN should has \"Vals\" attribute")
	ErrJsonINColShouldArray  = errors.New("In \"Vals\" attribute should be Array")

	//Like
	ErrJsonLikeShouldArray       = errors.New("LIKE should be Array")
	ErrJsonLikeShouldLen2        = errors.New("LIKE Array should be Len(2)")
	ErrJsonLikeShouldArrayString = errors.New("LIKE should be Array of String")

	//Not
	ErrJsonNotShouldInterfaceArray = errors.New("NOT should []interface{}")
	ErrJsonNotArrayShouldLen1      = errors.New("NOT cond should [1]interface{}")
	ErrJsonNotCondShouldMap        = errors.New("NOT cond should be map[string]interface{}")
)
