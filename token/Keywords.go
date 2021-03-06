package token

// Keywords is a map to help lookup simple keyword types
var Keywords = map[string]Type{
	`aint`:       Aint,
	`as`:         As,
	`and`:        And,
	`build`:      Build,
	`down`:       Down,
	`else`:       Else,
	`false`:      False,
	`wrong`:      False,
	`no`:         False,
	`lies`:       False,
	`greater`:    Greater,
	`higher`:     Greater,
	`bigger`:     Greater,
	`stronger`:   Greater,
	`great`:      GreaterEqual,
	`high`:       GreaterEqual,
	`big`:        GreaterEqual,
	`strong`:     GreaterEqual,
	`if`:         If,
	`is`:         Is,
	`was`:        Is,
	`were`:       Is,
	`knock`:      Knock,
	`less`:       Less,
	`lower`:      Less,
	`smaller`:    Less,
	`weaker`:     Less,
	`low`:        LessEqual,
	`little`:     LessEqual,
	`small`:      LessEqual,
	`weak`:       LessEqual,
	`minus`:      Minus,
	`without`:    Minus,
	`mysterious`: Mysterious,
	`not`:        Not,
	`null`:       Null,
	`nothing`:    Null,
	`nowhere`:    Null,
	`nobody`:     Null,
	`empty`:      Null,
	`gone`:       Null,
	`over`:       Over,
	`plus`:       Plus,
	`with`:       Plus,
	`it`:         Pronoun,
	`he`:         Pronoun,
	`she`:        Pronoun,
	`him`:        Pronoun,
	`her`:        Pronoun,
	`they`:       Pronoun,
	`them`:       Pronoun,
	`ze`:         Pronoun,
	`hir`:        Pronoun,
	`zie`:        Pronoun,
	`zir`:        Pronoun,
	`xe`:         Pronoun,
	`xem`:        Pronoun,
	`ve`:         Pronoun,
	`ver`:        Pronoun,
	`say`:        Say,
	`shout`:      Say,
	`whisper`:    Say,
	`scream`:     Say,
	`says`:       Says,
	`takes`:      Takes,
	`taking`:     Taking,
	`than`:       Than,
	`times`:      Times,
	`of`:         Times,
	`true`:       True,
	`right`:      True,
	`yes`:        True,
	`ok`:         True,
	`while`:      While,
	`until`:      While,
	`up`:         Up,
	`put`:        Put,
	`into`:       Into,
}
