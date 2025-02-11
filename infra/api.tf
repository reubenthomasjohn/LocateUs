module "api" {
  source = "./modules/ec2"

  public_key            = var.public_key
  private_key           = var.private_key
  key_name              = var.key_name
}