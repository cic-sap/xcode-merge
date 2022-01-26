%{
package jsonparser

type pair struct {
  key string
  val interface{}
}

func setResult(l yyLexer, v map[string]interface{}) {
  l.(*lex).result = v
}
%}

%union{
  obj map[string]interface{}
  list []interface{}
  pair pair
  val interface{}
}

%token LexError
%token <val> String Number Literal

%type <obj> object members
%type <pair> pair
%type <val> array
%type <list> elements
%type <val> value

%right UMINUS      /*  supplies  precedence  for  unary  minus  */

%start object

%%
object: '{' members '}'
  {
    $$ = $2
    setResult(yylex, $$)
  }
  | '{' '}'
  {
  $$ = map[string]interface{}{}
  setResult(yylex, $$)
  }

members:
  pair
  {
    $$ = map[string]interface{}{
      $1.key: $1.val,
    }
  }
|   members   pair
  {
    $1[$2.key] = $2.val
    $$ = $1
  }


pair: String '=' value ';'
  {
    $$ = pair{key: $1.(string), val: $3}
  }
| Literal '=' value ';'
 {
  $$ = pair{key: $1.(string), val: $3}
 }

array: '(' elements ')'
  {
    $$ = $2
  }
  | '(' ')'
  {
  $$ = []interface{}{}
  }

elements:
 value ','
  {
    $$ = []interface{}{$1}
  }
| elements   value ','
  {
    $$ = append($1, $2)
  }

value:
  String
| Number
| Literal
| object
  {
    $$ = $1
  }
| array
