#not an actual package yet: just main functions. This file produces an executable, main
include $(GOROOT)/src/Make.$(GOARCH)

TARG=ed2kcrawler
GOFILES=\
        main.go\
        loadfile.go\
        httpget.go\
        urlpaser.go\

main: main.${O}
	${LD} -o ${TARG} main.${O}

main.${O}: ${GOFILES}
	${GC} -o main.${O} ${GOFILES}

clean:
	rm ${TARG};rm *.6
