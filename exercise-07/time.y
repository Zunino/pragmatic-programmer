/**
* Book:     The Pragmatic Programmer
* Chapter:  2 - A Pragmatic Approach
* Exercise: 7
* Page:     63
*
* Implement a parser for the BNF grammar defined in exercise 6 using yacc,
* bison or a similar parser generator.
*
* Andre Zunino <neyzunino@gmail.com>
* 20 September 2019
*/

%{
#include <stdio.h>
#include <stdbool.h>
#include <ctype.h>

#include "parser.h"

#define YYDEBUG 1

const char* input;
char msg_buffer[80];
callback_fn callback;

bool check_value(int value, int max, const char* const message);
const char* const status_str[] = {"OK", "ERROR"};

int yylex();
int yyerror(const char* const err);
%}

%token DIGIT
%token SEP
%token AM
%token PM

%%
time:       hours am_pm {
                if (!check_value($1, 11, "is an invalid hour value in 12-hour format")) YYABORT;
                callback(OK, "Hour + AM/PM");
            }
            | hours SEP minutes {
                callback(OK, "Hour + colon + minutes");
            }
            | hours SEP minutes am_pm {
                if (!check_value($1, 11, "is an invalid hour value in 12-hour format")) YYABORT;
                callback(OK, "Hour + colon + minutes + AM/PM");
            }
            ;

hours:      DIGIT | two_digit {
                if (!check_value($1, 23, "is an invalid hour value")) YYABORT;
                $$ = $1;
            };

two_digit:  DIGIT DIGIT {
                $$ = $1 * 10 + $2;
            }
            ;

am_pm:      AM {
                $$ = 0;
            }
            | PM {
                $$ = 12 * 60;
            }
            ;

minutes:    DIGIT DIGIT {
                $$ = $1 * 10 + $2;
                if (!check_value($$, 59, "is an invalid minute value")) YYABORT;
            }
%%

bool check_value(int value, int max, const char* const message) {
    if (value > max) {
        sprintf(msg_buffer, "%d %s", value, message);
        yyerror(msg_buffer);
        return false;
    }
    return true;
}

const char* const str(enum status_t status) {
    return status_str[status];
}

int yylex() {
    char c = *input;
    if (!c) {
        return 0;
    }
    ++input;
    if (isdigit(c)) {
        yylval = c - '0';
        return DIGIT;
    }
    if ((c == 'a' || c == 'p') && *input == 'm') {
        ++input;
        return c == 'a' ? AM : PM;
    }
    if (c == ':') {
        return SEP;
    }
    return c;
}

int yyerror(const char* const err) {
    callback(ERROR, err);
}

void parse(const char* text, callback_fn fn) {
    input = text;
    callback = fn;
    yyparse();
}

