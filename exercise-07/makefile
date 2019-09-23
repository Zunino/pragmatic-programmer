parser: parser.c
	gcc -o $@ $<

parser.c: time.y
	bison -o $@ $<

clean:
	rm -f parser parser.c

