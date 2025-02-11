## Steps to setup:

### Locally run terraform to test configuration:

1. Create an IAM User with Administrator Access. Note down the SECRET_ACCESS and ACCESS_KEY
2. Create an S3 bucket to hold the `.tfstate` file.
3. Create a secret key to ssh into an EC2 machine in the EC2 tab. Once the .pem file is downloaded:
   a. cat path/to/file - this is the private ssh key.
   b. ssh-keygen -y -f path/to/pemfile - this is the public ssh key.
4. Run `aws configure`. Paste in the ACCESS_KEY and SECRET_ACCESS key for the create IAM User.
5. Run `terraform init`

TODO:

1. Create terraform to apply S3 to host S3
2. Create cloudfront, and route53 to connect alias.
3. github action to deploy to s3, and invalidate cache when theres a new deployment.

##

`sudo snap install docker`
`sudo snap start docker`
`sudo apt install make`

```sh
sudo groupadd docker
sudo usermod -aG docker $USER
su -s ${USER}
```
