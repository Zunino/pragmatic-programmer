parser: parser.o parser_main.o
	gcc -o $@ $^

%.o: %.c
	gcc -c -o $@ $<

parser.c: time.y parser.h
	bison -o $@ time.y

test: parser_test parser.h
	@./parser_test

parser_test: parser_test.o parser.o
	gcc -o $@ $^

clean:
	rm -f parser parser_test *.o parser.c

