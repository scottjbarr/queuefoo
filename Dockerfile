FROM golang:1.8.1

COPY dist/queue-foo-receive /queue-foo-receive

CMD "/queue-foo-receive"
