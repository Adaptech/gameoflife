FROM mkboudreau/go-present

WORKDIR /app
COPY *.article .
COPY ./curls ./curls

CMD /run/present-run -orighost=localhost -play=false
