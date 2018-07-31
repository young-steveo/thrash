package token

// Literals is a map to help lookup literal types
var Literals = map[string]Type{
	`false`:      False,
	`wrong`:      False,
	`no`:         False,
	`lies`:       False,
	`mysterious`: Mysterious,
	`null`:       Null,
	`nothing`:    Null,
	`nowhere`:    Null,
	`nobody`:     Null,
	`empty`:      Null,
	`gone`:       Null,
	`true`:       True,
	`right`:      True,
	`yes`:        True,
	`ok`:         True,
}
