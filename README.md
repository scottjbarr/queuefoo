# Queue Foo

A little SQS example behind a fairly generic Queue interface.


## Config

Config is provided by environment variables.

See the [example config](conf/dev.env.example)


## Run it

See the [Makefile](Makefile)

A raw example

    docker run --env-file ./conf/docker.env queue-foo-receive


## Hit SQS via standard cli tools

    URL="https://sqs.ap-southeast-2.amazonaws.com/123456/queuefoo"
    source conf/dev.env && aws sqs receive-message --queue-url "$URL"


## Licence

The MIT License (MIT)

Copyright (c) 2017 Scott Barr

See [LICENCE.md](LICENCE.md)
