FROM golang:1.6

COPY ./tyne-quiz-api /usr/bin/tyne-quiz-api

# From here we load our application's code in, therefore the previous docker
# "layer" thats been cached will be used if possible
WORKDIR /usr/bin

CMD ["tyne-quiz-api"]
