#include <stdio.h>

#include "parser.h"

void def_callback(enum status_t status, const char* const message) {
    FILE* stream = status == OK ? stdout : stderr;
    fprintf(stream, "%s\n", message);
}

int main(int argc, char* argv[]) {
    if (argc > 1) {
        parse(argv[1], &def_callback);
    }
    return 0;
}

