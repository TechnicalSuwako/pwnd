NAME != cat main.go | grep "var sofname" | awk '{print $$4}' | sed "s/\"//g"
VERSION != cat main.go | grep "var version" | awk '{print $$4}' | sed "s/\"//g"
PREFIX = /usr/local

CC = CGO_ENABLED=0 go build
RELEASE = -ldflags="-s -w" -buildvcs=false

all:
	${CC} ${RELEASE} -o ${NAME}

release:
	mkdir -p release/bin/${VERSION}/openbsd/amd64
	env GOOS=openbsd GOARCH=amd64 ${CC} ${RELEASE} -o\
		release/bin/${VERSION}/openbsd/amd64/${NAME}

clean:
	rm -f ${NAME}

dist:
	mkdir -p ${NAME}-${VERSION} release/src
	cp -R LICENSE.txt Makefile README.md CHANGELOG.md\
		main.go ${NAME}.rc src go.mod go.sum ${NAME}-${VERSION}
	tar zcfv release/src/${NAME}-${VERSION}.tar.gz ${NAME}-${VERSION}
	rm -rf ${NAME}-${VERSION}

install:
	mkdir -p ${DESTDIR}${PREFIX}/bin ${DESTDIR}/etc/rc.d
	cp -f ${NAME} ${DESTDIR}${PREFIX}/bin
	chmod 755 ${DESTDIR}${PREFIX}/bin/${NAME}
	cp -f ${NAME}.rc ${DESTDIR}/etc/rc.d/${NAME}
	chmod +x ${DESTDIR}/etc/rc.d/${NAME}

uninstall:
	rm -f ${DESTDIR}${PREFIX}/bin/${NAME}

.PHONY: all release clean dist install uninstall
