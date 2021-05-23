terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }
}

# Configure the AWS Provider
provider "aws" {
  region = "us-east-1"
}

resource "aws_sqs_queue" "notification_queue" {
  name = "notification-queue"
}

resource "aws_sns_topic" "user_updates" {
  name = "notification-topic"
}

