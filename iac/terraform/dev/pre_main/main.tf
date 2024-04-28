locals {
  environment = "dev"
  api_service_name = "template-service-api"
}

provider "aws" {
    region = "ap-northeast-1"
}

terraform {
  required_version = ">= 1.0.0, < 2.0.0"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0"
    }
  }

  backend "s3" {
    # S3バケットとDynamoDBのテーブルは事前にマネージメントコンソールから手動で作成してください。
    bucket         = "dev-terraform-state-gangan-ecr-bucket"
    key            = "template-pj/terraform.tfstate"
    region         = "ap-northeast-1"
    dynamodb_table = "dev-terraform-ecr-lock"
    encrypt        = true
  }
}

module "ecr" {
  source = "./../../modules/ecr"
  environment = local.environment
  api_service_name = local.api_service_name
}

module "iam" {
  source = "./../../modules/iam_pre"
  environment = local.environment
}