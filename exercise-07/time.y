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

#define YYDEBUG 1

const char* input;
char msg_buffer[80];

bool check_12_hour_value(int);

int yylex();
int yyerror(const char* const err);
%}

%token DIGIT
%token SEP
%token AM
%token PM

%%
time:       hours am_pm {
                if (!check_12_hour_value($1)) YYABORT;
                printf("Hour + AM/PM\n");
            }
            | hours SEP minutes {
                printf("Hour + colon + minutes\n");
            }
            | hours SEP minutes am_pm {
                if (!check_12_hour_value($1)) YYABORT;
                printf("hours: %d | minutes: %d | ampm: %d\n", $1, $3, $4);
                printf("Hour + colon + minutes + AM/PM\n");
            }
            ;

hours:      DIGIT | two_digit {
                if ($1 > 23) {
                    sprintf(msg_buffer, "%d is not valid for hours", $1);
                    yyerror(msg_buffer);
                    YYABORT;
                }
                $$ = $1 * 60;
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
                if ($$ > 59) {
                    sprintf(msg_buffer, "%d is not valid for minutes", $$);
                    yyerror(msg_buffer);
                    YYABORT;
                }
            }
%%

bool check_12_hour_value(int hours) {
    if (hours > 12) {
        sprintf(msg_buffer, "%d is not valid for hours in 12-clock format", hours);
        yyerror(msg_buffer);
        return false;
    }
    return true;
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
    fprintf(stderr, "Error: %s\n", err);
}

int main(int argc, char* argv[]) {
    if (argc > 1) {
        input = argv[1];
        yyparse();
    }
    return 0;
}

