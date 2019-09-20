/**
* Book:     The Pragmatic Programmer
* Chapter:  2 - A Pragmatic Approach
* Exercise: 5
* Page:     63
*
* P 2   # select pen 2
* D     # pen down
* W 2   # draw west 2 units
* N 1   # draw north 1 unit
* E 2   # draw east 2 units
* S 1   # draw south 1 unit
* U     # pen up
*
* Implement the code that parses this language. It should be designed so that
* it is simple to add new commands.
*
* Andre Zunino <neyzunino@gmail.com>
* 19 September 2019
*/

package main

import (
    "bufio"
    "fmt"
    "strings"
    "text/scanner"
)

var commandMap = map[string]command {
    "P": command {"P", "Select a pen", true},
    "D": command {"D", "Put the pen down", false},
    "U": command {"U", "Lift the pen", false},
    "W": command {"W", "Move the pen west", true},
    "N": command {"N", "Move the pen north", true},
    "E": command {"E", "Move the pen east", true},
    "S": command {"S", "Move the pen south", true},
}

const CommentDelim = "#"

type command struct {
    keyword string
    desc string
    hasArg bool
}

type ParsedCommand struct {
    keyword string
    arg string
}

func Parse(input string) ([]ParsedCommand, error) {
    strReader := strings.NewReader(input)
    bufReader := bufio.NewReader(strReader)
    b := bufio.NewScanner(bufReader)
    parsedCommands := []ParsedCommand {}
    for b.Scan() {  // for each line of input
        line := strings.TrimSpace(b.Text())
        if len(line) == 0 {
            continue
        }
        parsedCmd, err := parseLine(line)
        if err != nil {
            return parsedCommands, err
        }
        if parsedCmd != nil {
            parsedCommands = append(parsedCommands, *parsedCmd)
        }
    }
    return parsedCommands, nil
}

func parseLine(line string) (*ParsedCommand, error) {
    var s scanner.Scanner
    s.Init(strings.NewReader(line))
    tok := s.Scan()
    if (s.TokenText() == CommentDelim) {
        return nil, nil
    }
    cmd, ok := commandMap[s.TokenText()]
    if !ok {
        return nil, fmt.Errorf("Command '%s' not recognized", s.TokenText())
    }
    arg := ""
    tok = s.Scan()
    if cmd.hasArg {
        if tok == scanner.EOF || s.TokenText() == CommentDelim {
            return nil, fmt.Errorf("Missing required argument for command %s", cmd.keyword)
        }
        arg = s.TokenText()
        tok = s.Scan()
        if tok != scanner.EOF && s.TokenText() != CommentDelim {
            return nil, fmt.Errorf("Command '%s' takes a single argument", cmd.keyword)
        }
    } else {
        if tok != scanner.EOF && s.TokenText() != CommentDelim {
            return nil, fmt.Errorf("Command '%s' does not expect arguments", cmd.keyword)
        }
    }
    return &ParsedCommand{cmd.keyword, arg}, nil
}

