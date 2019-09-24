#ifndef AZ_PARSER_H
#define AZ_PARSER_H

enum status_t {
    OK,
    ERROR
};

const char* const str(enum status_t);

typedef void (*callback_fn)(enum status_t status, const char* const message);

void parse(const char*, callback_fn);

#endif

