module "website" {
  source = "./modules/static_website"

  bucket_name            = "${var.company_name}-website-${var.stage}"
  domain                 = var.company_domain_name
}