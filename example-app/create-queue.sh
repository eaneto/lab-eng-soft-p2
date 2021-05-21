#!bin/sh

awslocal sqs create-queue \
    --queue-name notification-queue
