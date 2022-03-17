terraform {
  required_providers {
    download = {
      version = "~> 0.0.1"
      source  = "github.com/bjunker99/download"
    }
  }
}

data "download_file" "test" {
  url = "https://github.com/philips-labs/terraform-aws-github-runner/releases/download/v0.37.0/runners.zip"
  output_file = "runners.zip"  
}
