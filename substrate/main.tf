# Configure the Google Cloud provider
provider "google" {
  credentials = file("<PATH_TO_YOUR_GCP_SERVICE_ACCOUNT_KEY>")
  project     = "<YOUR_GCP_PROJECT_ID>"
  region      = "us-central1"
}

# Create a Google Compute Engine instance
resource "google_compute_instance" "substrate_node_instance" {
  name         = "substrate-node-instance"
  machine_type = "n1-standard-2"
  zone         = "us-central1-a"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-10"
    }
  }

  network_interface {
    network = "default"
  }

  metadata_startup_script = <<-EOF
    #!/bin/bash
    apt-get update -y
    apt-get install -y git

    # Clone Substrate repository
    git clone https://github.com/paritytech/substrate.git /opt/substrate

    # Install required dependencies
    /opt/substrate/scripts/init.sh

    # Build Substrate
    cd /opt/substrate
    cargo build --release

    # Start Substrate node
    /opt/substrate/target/release/substrate --dev
  EOF
}

# Output the external IP address of the Substrate node instance
output "external_ip" {
  value = google_compute_instance.substrate_node_instance.network_interface.0.access_config.0.assigned_nat_ip
}
