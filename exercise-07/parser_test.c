#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <assert.h>

#include "parser.h"

enum status_t last_status = OK;
const char* last_message = "";

const char* const GREEN = "\x1b[32m";
const char* const RED = "\x1b[31m";
const char* const RESET = "\x1b[0m";

void parser_callback(enum status_t status, const char* const message) {
    last_status = status;
    last_message = message;
}

void run_test(const char* const context, const char* input) {
    parse(input, &parser_callback);
    if (last_status == OK) {
        printf("%s-- %s%s\n\tParse '%s' => OK\n", GREEN, context, RESET, input);
    } else {
        printf("%s-- %s%s\n\tParse '%s' => FAIL\n", RED, context, RESET, input);
    }
}

void assert_str(const char* const expected, const char* const actual) {
    if (strcmp(expected, actual) != 0) {
        fprintf(stderr, "%s\tExpected: '%s'\n\tGot:      '%s'%s\n", RED, expected, actual, RESET);
        abort();
    }
}

void assert_int(int expected, int actual) {
    if (actual != expected) {
        fprintf(stderr, "\tExpected: '%d'\n\tGot:      '%d'\n", expected, actual);
        abort();
    }
}

void assert_status(enum status_t expected, enum status_t actual) {
    if (actual != expected) {
        fprintf(stderr, "\tExpected: '%s'\n\tGot:      '%s'\n", str(expected), str(actual));
        abort();
    }
}

void test_hour_without_ampm_should_fail() {
    run_test(__func__, "10");
    assert_status(ERROR, last_status);
}

void test_invalid_hour_without_ampm_should_fail() {
    run_test(__func__, "50");
    assert_status(ERROR, last_status);
    assert_str("50 is an invalid hour value", last_message);
}

void test_hour_ampm_should_succeed() {
    run_test(__func__, "3am");
    assert_status(OK, last_status);
    assert_str("Hour + AM/PM", last_message);
}

void test_hour_minutes_without_ampm_should_succeed() {
    run_test(__func__, "15:20");
    assert_status(OK, last_status);
    assert_str("Hour + colon + minutes", last_message);
}

void test_invalid_hour_minutes_without_ampm_should_fail() {
    run_test(__func__, "25:20");
    assert_status(ERROR, last_status);
    assert_str("25 is an invalid hour value", last_message);
}

void test_hour_minutes_with_ampm_should_succeed() {
    run_test(__func__, "5:40pm");
    assert_status(OK, last_status);
    assert_str("Hour + colon + minutes + AM/PM", last_message);
}

void test_invalid_hour_minutes_with_ampm_should_fail() {
    run_test(__func__, "15:22pm");
    assert_status(ERROR, last_status);
    assert_str("15 is an invalid hour value in 12-hour format", last_message);
}

void test_invalid_minutes_without_ampm_should_fail() {
    run_test(__func__, "15:72pm");
    assert_status(ERROR, last_status);
    assert_str("72 is an invalid minute value", last_message);
}

int main() {
    printf("Running tests...\n");
    test_hour_without_ampm_should_fail();
    test_invalid_hour_without_ampm_should_fail();
    test_hour_ampm_should_succeed();
    test_hour_minutes_without_ampm_should_succeed();
    test_invalid_hour_minutes_without_ampm_should_fail();
    test_hour_minutes_with_ampm_should_succeed();
    test_invalid_hour_minutes_with_ampm_should_fail();
    test_invalid_minutes_without_ampm_should_fail();
}

