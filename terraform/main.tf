terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }
}

# Configura provider da AWS na região de virgínia.
provider "aws" {
  region = "us-east-1"
}

# Cria fila SQS de notificação.
resource "aws_sqs_queue" "notification_queue" {
  name = "notification-queue"
}

# Cria tópico SNS de notificação.
resource "aws_sns_topic" "user_updates" {
  name = "notification-topic"
}

