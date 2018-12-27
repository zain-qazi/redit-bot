FROM golang:1.8

WORKDIR /app

#Env var for the repo
ENV SRC_DIR=/go/src/redit-bot

#Add the source code
ADD . $SRC_DIR

#mount volume for user agent file
#VOLUME ~/go /app
#VOLUME ~/go/src /go/src

# #Build it
# RUN 

#Get dependencies
RUN go get -v -d "github.com/turnage/graw"
RUN go get -v -d "github.com/spf13/cobra"

#RUN cd $SRC_DIR; go build -o redit-bot; cp redit-bot /app/

RUN cd $SRC_DIR; go install redit-bot

ENTRYPOINT ["redit-bot"]

#CMD ["redit-bot"]