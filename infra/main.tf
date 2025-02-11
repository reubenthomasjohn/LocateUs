terraform {
    required_providers{
        aws = {
            source = "hashicorp/aws"
            version = "~> 5.0"
        }
    }
    backend "s3" {
        bucket = "ccc-heatmapping"  
        key    = "aws/ec2-api/terraform.tfstate"
        region = "ap-south-1"
    }
}

provider "aws" {
    region = var.region
}
