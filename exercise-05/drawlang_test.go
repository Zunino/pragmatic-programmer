package main

import (
    "testing"
)

func TestEmptyParse(t *testing.T) {
    // given
    src := ""

    // when
    parsed, _ := Parse(src)

    // then
    if len(parsed) != 0 {
        t.Errorf("Expected an empty collection of parsed commands\n")
    }
}

func TestSingleCommandWithoutArg(t * testing.T) {
    // given
    src := `
        U
    `
    // when
    parsed, _ := Parse(src)

    // then
    if len(parsed) != 1 {
        t.Errorf("Expected a collection with 1 parsed command\n")
        return
    }
    cmd := parsed[0]
    if cmd.keyword != "U" {
        t.Errorf("Expected keyword to be 'P', but was '%s'\n", cmd.keyword)
    }
    if cmd.arg != "" {
        t.Errorf("Expected argument to be empty, but was '%s'\n", cmd.arg)
    }
}

func TestInvalidSingleCommandWithoutArg(t * testing.T) {
    // given
    src := `
        Z
    `
    // when
    _, err := Parse(src)

    // then
    if err == nil {
        t.Errorf("Expected a parsing error to ocurr\n")
        return
    }
    if err.Error() != "Command 'Z' not recognized" {
        t.Errorf("Parsing error was not the one expected\n");
    }
}

func TestSingleCommandWithArg(t * testing.T) {
    // given
    src := `
        P 1
    `
    // when
    parsed, _ := Parse(src)

    // then
    if len(parsed) != 1 {
        t.Errorf("Expected a collection with 1 parsed command, but got %d\n", len(parsed))
        return
    }
    cmd := parsed[0]
    if cmd.keyword != "P" {
        t.Errorf("Expected keyword to be 'P', but was '%s'\n", cmd.keyword)
    }
    if cmd.arg != "1" {
        t.Errorf("Expected argument to be '1', but was '%s'\n", cmd.arg)
    }
}

func TestUnexpectedArgument(t * testing.T) {
    // given
    src := `
        D 1
    `
    // when
    _, err := Parse(src)

    // then
    if err == nil {
        t.Errorf("Expected a parsing error to ocurr\n");
        return
    }
    if err.Error() != "Command 'D' does not expect arguments" {
        t.Errorf("Parsing error was not the one expected\n");
    }
}

func TestOnlyComments(t* testing.T) {
    // given
    src := `
        # first comment
        # second comment
    `

    // when
    parsed, _ := Parse(src)

    // then
    if len(parsed) != 0 {
        t.Errorf("Expected an empty collection of parsed commands\n")
    }
}

func TestCommentAfterCommand(t* testing.T) {
    // given
    src := `
        S 1 # first comment
        # second comment
    `

    // when
    parsed, _ := Parse(src)

    // then
    if len(parsed) != 1 {
        t.Errorf("Expected a collection with 1 parsed command\n")
        return
    }
    cmd := parsed[0]
    if cmd.keyword != "S" {
        t.Errorf("Expected keyword to be 'S', but was '%s'\n", cmd.keyword)
    }
    if cmd.arg != "1" {
        t.Errorf("Expected argument to be '1', but was '%s'\n", cmd.arg)
    }
}

func TestExtraArgument(t * testing.T) {
    // given
    src := `
        W 10 11
    `
    // when
    _, err := Parse(src)

    // then
    if err == nil {
        t.Errorf("Expected a parsing error to ocurr\n");
        return
    }
    if err.Error() != "Command 'W' takes a single argument" {
        t.Errorf("Parsing error was not the one expected\n");
    }
}

func TestMultipleCommands(t* testing.T) {
    // given
    src := `
        P 1

        D

        N 10
        E 5

        U
    `

    // when
    parsed, _ := Parse(src)

    // then
    if len(parsed) != 5 {
        t.Errorf("Expected a collection with 1 parsed command\n")
        return
    }
    cmd1 := parsed[0]
    if cmd1.keyword != "P" {
        t.Errorf("Expected keyword to be 'P', but was '%s'\n", cmd1.keyword)
    }
    if cmd1.arg != "1" {
        t.Errorf("Expected argument to be '1', but was '%s'\n", cmd1.arg)
    }
    cmd5 := parsed[4]
    if cmd5.keyword != "U" {
        t.Errorf("Expected keyword to be 'U', but was '%s'\n", cmd5.keyword)
    }
    if cmd5.arg != "" {
        t.Errorf("Expected no argument, but was '%s'\n", cmd5.arg)
    }
}

func TestMultipleCommandsWithComments(t* testing.T) {
    // given
    src := `
        P 2   # select pen 2
        D     # pen down
        W 2   # draw west 2 units
        N 1   # draw north 1 unit
        E 2   # draw east 2 units
        S 1   # draw south 1 unit
        U     # pen up
    `

    // when
    parsed, _ := Parse(src)

    // then
    if !checkSize(t, parsed, 7) {
        return
    }
    cmd1 := parsed[0]
    if cmd1.keyword != "P" {
        t.Errorf("Expected keyword to be 'P', but was '%s'\n", cmd1.keyword)
    }
    if cmd1.arg != "2" {
        t.Errorf("Expected argument to be '2', but was '%s'\n", cmd1.arg)
    }
    cmd7 := parsed[6]
    if cmd7.keyword != "U" {
        t.Errorf("Expected keyword to be 'U', but was '%s'\n", cmd7.keyword)
    }
    if cmd7.arg != "" {
        t.Errorf("Expected no argument, but was '%s'\n", cmd7.arg)
    }
}

func checkSize(t *testing.T, coll []ParsedCommand, expected int) bool {
    if len(coll) != expected {
        t.Errorf("Expected a collection with %d parsed commands, but got %d\n", expected, len(coll))
        return false
    }
    return true
}

