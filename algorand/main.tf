provider "aws" {
  region = "us-east-1"  # Choose your desired AWS region
}

resource "aws_instance" "algorand_node" {
  ami           = "ami-0c55b159cbfafe1f0"  # Algorand Node AMI for Ubuntu 20.04 (Update this with the latest Algorand AMI)
  instance_type = "t2.micro"                # Choose an appropriate instance type

  key_name = "your-key-pair"  # Replace with your AWS key pair name

  tags = {
    Name = "AlgorandNode"
  }

  user_data = <<-EOF
              #!/bin/bash
              # Algorand Node Installation Script
              sudo apt-get update -y
              sudo apt-get install -y curl
              curl -O https://algorand-releases.s3.us-east-2.amazonaws.com/channel/stable/mainnet/latest/linux-amd64/algorand
              chmod +x algorand
              sudo mv algorand /usr/local/bin/
              algorand init
              algorand start -d /var/lib/algorand
              EOF

  lifecycle {
    create_before_destroy = true
  }
}

output "algorand_node_public_ip" {
  value = aws_instance.algorand_node.public_ip
}
