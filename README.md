# **Korp Lab**

<p align="justify">
   <img src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/go/go-original.svg" width="50" height="50"/>
   <img src="https://github.com/gin-gonic/logo/blob/master/color.svg" width="50" height="50"/>
   <img src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/docker/docker-original.svg" width="50" height="50"/>
   <img src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/nginx/nginx-original.svg" width="50" height="50"/>
   <img src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/prometheus/prometheus-original.svg" width="50" height="50"/>
   <img src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/grafana/grafana-original.svg" width="50" height="50"/>
   <img src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/ansible/ansible-original.svg" width="50" height="50"/>
   <img src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/vagrant/vagrant-original.svg" width="50" height="50"/>
   <img src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/ubuntu/ubuntu-original.svg" width="50" height="50"/>
</p>

A containerized HTTP service built in Go with Gin, fully provisioned and automated with Ansible on a Vagrant-managed virtual machine. The service's behavior is monitored in real time with a complete observability stack using Prometheus and Grafana.

## Table of Contents

- [Overview](#overview)
- [Architecture](#architecture)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
- [Usage Example](#usage-example)
- [Observability Stack](#observability-stack)
- [Observability Pillars Implemented](#observability-pillars-implemented)
- [Infrastructure Provisioning with Vagrant and Ansible](#infrastructure-provisioning-with-vagrant-and-ansible)
- [Containers and Docker Compose](#containers-and-docker-compose)
- [Proof of Execution](#proof-of-execution)
- [Architectural Decisions](#architectural-decisions)
- [Future Improvements](#future-improvements)

## Overview

Korp Lab is a small HTTP service, http-server-projeto-korp, that exposes a single endpoint returning its name and the current UTC. While the service itself is intentionally simple, the project's real focus is the infrastructure and automation around it: a reverse proxy, a full observability stack, and an entirely automated provisioning pipeline that takes a basic Linux VM and turns it into a fully running environment with a single Ansible command.

The project was built as a technical assessment, with an emphasis on clear architectural decisions, reproducibility, and commitment to infrastructure best practices.


## Architecture

[I'll add a DrawIO image]

All services run on a shared Docker bridge network (korp-network), created by Ansible before any container starts.

The application service is **not** exposed directly to the host. All external traffic reaches it exclusively through the Nginx reverse proxy. Prometheus reaches the application's `/metrics` endpoint directly over the internal Docker network, bypassing Nginx entirely, since that endpoint is meant for machine-to-machine scraping, not external or human consumption.

## Tech Stack

| Layer | Technology |
|---|---|
| Application | Go 1.26, [Gin](https://github.com/gin-gonic/gin) |
| Containerization | Docker, Docker Compose |
| Reverse proxy | Nginx |
| Metrics | Prometheus (`client_golang`) |
| Visualization | Grafana |
| Infrastructure provisioning | Vagrant, VirtualBox |
| Configuration management | Ansible |
| Operating System | Ubuntu 22.04 LTS (Jammy) |
> [!IMPORTANT]
> Be aware that Ubuntu was chosen as the operating system for the virtual machine in this project. This decision was made solely for convenience during development and testing.
> 
> That said, feel free to use a different Linux distribution if you prefer.
> 
> Just keep in mind that if you choose a distribution that is not Debian-based or does not use the APT package manager, you will need to adapt the Ansible playbooks. They currently rely on modules and tasks designed for APT-based environments, so they will not work correctly on distributions that use other package managers (such as dnf, pacman, etc) without the necessary adjustments.


## Project Structure

```yaml
korp-lab
├── ansible
│   ├── inventory.ini
│   ├── inventory.ini.example
│   ├── playbook.yaml
│   ├── requirements.yaml
│   └── roles
│       ├── app
│       │   └── tasks
│       │       └── main.yaml
│       ├── clone-repository
│       │   └── tasks
│       │       └── main.yaml
│       ├── docker-install
│       │   └── tasks
│       │       └── main.yaml
│       ├── docker-network
│       │   └── tasks
│       │       └── main.yaml
│       ├── docker-non-root
│       │   └── tasks
│       │       └── main.yaml
│       ├── grafana
│       │   └── tasks
│       │       └── main.yaml
│       ├── nginx
│       │   └── tasks
│       │       └── main.yaml
│       ├── prometheus
│       │   └── tasks
│       │       └── main.yaml
│       ├── service-validation
│       │   └── tasks
│       │       └── main.yaml
│       └── update-os
│           └── tasks
│               └── main.yaml
├── cmd
│   └── api
│       └── main.go
├── docker-compose.yaml
├── Dockerfile
├── docs
├── go.mod
├── go.sum
├── grafana
│   ├── dashboards
│   │   └── http-server-projeto-korp-dashboard.json
│   └── provisioning
│       ├── dashboards
│       │   └── dashboards.yaml
│       └── datasources
│           └── datasources.yaml
├── LICENSE
├── nginx
│   └── http-server-projeto-korp.conf
├── prometheus
│   └── prometheus.yaml
├── README.md
└── Vagrantfile
```

## Getting Started

This repository does not cover the installation process for the technologies used in the project, such as Vagrant, VirtualBox, Ansible, Go, Docker, and Docker Compose.

The installation process for these tools can vary significantly depending on your operating system (Linux, macOS, or Windows), Linux distribution, package manager, and even the installation method you choose.

For this reason, the recommended approach is always to follow the official documentation for each technology, using the instructions specific to your environment.

Keep in mind that there are multiple ways to install the same tool. Depending on your preferences or environment, you may choose to use a package manager, compressed archives, native installers, AppImage, or any other officially supported installation method.

Once all the required dependencies are installed and you’ve confirmed they’re working correctly in your environment, you’ll be ready to run this project.

### Prerequisites

- [Vagrant](https://www.vagrantup.com/)
- [VirtualBox](https://www.virtualbox.org/)
- [Ansible](https://docs.ansible.com/) (on the control machine, not the VM)

### 1. Clone this repository

```bash
git clone https://github.com/thomaspalma1/korp-lab.git
cd korp-lab
```

(Talvez trocar por um curl usando um install sh ou makefile?)

### 2. Bring up the virtual machine

```bash
vagrant up
```

This creates an Ubuntu 22.04 VM with a fixed private network IP (`192.168.56.10`). No software provisioning happens at this stage, since Vagrant is only responsible for creating the machine itself.
> [!NOTE]  
> This is a private IP address chosen solely for convenience during the project's development.
> 
> If you want to change it, feel free to use another IP address within your private network. However, remember to update every location where this value is used.
> 
> When changing the IP address, you will need to make the same update in the following files:
> 
> - Vagrantfile, where the virtual machine's IP address is defined.
> - The Ansible `inventory.ini` file, inside the `korp_lab` group, where the managed machine's address is configured.
> 
> Make sure both files are using the same address to avoid communication issues between Vagrant and Ansible.

### 3. Prepare the Ansible inventory

```bash
cp ansible/inventory.ini.example ansible/inventory.ini
```

### 4. Install required Ansible collections

Inside the Ansible directory, run the following command:

```bash
ansible-galaxy collection install -r requirements.yaml
```

### 5. Provision the entire environment

Still inside the same directory, run this command:

```bash
ansible-playbook -i inventory.ini playbook.yaml
```

This single command installs Docker, configures a non-root user, creates the Docker network, clones this repository into the VM, configures Nginx, Prometheus, and Grafana, builds the application image, brings up the full stack, and finally validates the service with an HTTP request, printing the response directly to your console.



## Usage Example

Once provisioning completes, the service is reachable through the Nginx reverse proxy:

```bash
curl http://192.168.56.10/projeto-korp
```

```json
{
  "nome": "Projeto Korp",
  "horario": "2026-08-02T00:00:18Z"
}
```

Direct access to the application container is intentionally blocked. The service never exposes port 8080 to the host:

```bash
curl http://192.168.56.10:8080/projeto-korp
# curl: (7) Failed to connect to 192.168.56.10 port 8080: Connection refused
```

## Observability Stack

| Tool | Role |
|---|---|
| **Prometheus** | Scrapes and stores metrics exposed by the Go service |
| **Grafana** | Visualizes metrics through an auto-provisioned dashboard |

Access, once the environment is up:

| Service | URL |
|---|---|
| Application (via Nginx) | http://192.168.56.10 |
| Prometheus UI | http://192.168.56.10:9090 |
| Grafana UI | http://192.168.56.10:3000 (`admin` / `admin`) |

## Observability Pillars Implemented

### Availability
Exposed through Prometheus's native `up` metric, generated automatically for every scrape target. No custom instrumentation is required for this pillar.

### Request volume
A custom counter, `http_requests_total`, incremented on every call to `/projeto-korp`, exposed in Prometheus exposition format via `client_golang`.

### Dashboard
A Grafana dashboard (`Projeto Korp - Overview`) with two panels: service availability (stat panel) and request volume (time series, using `rate(http_requests_total[1m])`).

## Infrastructure Provisioning with Vagrant and Ansible

- **Vagrant** is scoped narrowly: it only creates the VM (box, hostname, network, resources). It performs no software provisioning, since that responsibility belongs entirely to Ansible, avoiding duplicated responsibility between the two tools.
- **Ansible** handles everything inside the guest OS, organized into single-responsibility roles:

| Role | Responsibility |
|---|---|
| `update-os` | Updates the apt cache and upgrades packages |
| `docker-install` | Installs Docker following the official documentation for Ubuntu |
| `docker-non-root` | Adds the deployment user to the `docker` group |
| `docker-network` | Creates the `korp-network` bridge network |
| `clone-repository` | Clones this repository into the VM |
| `nginx` | Deploys the reverse proxy configuration |
| `prometheus` | Deploys the scrape configuration |
| `grafana` | Deploys datasource and dashboard provisioning files |
| `service-deploy` | Builds the application image and brings up the full stack |
| `service-validation` | Validates the service via HTTP and prints the response to the console |

## Containers and Docker Compose

Four containers, all connected to the same externally-managed bridge network:

| Container | Exposed to host? | Purpose |
|---|---|---|
| `http-server-projeto-korp` | No | The Go application |
| `nginx` | Yes (`80:80`) | Reverse proxy, the only public entry point |
| `prometheus` | Yes (`9090:9090`) | Metrics collection and querying |
| `grafana` | Yes (`3000:3000`) | Metrics visualization |

## Proof of Execution

The following captures follow the natural order of validation. The environment is provisioned, the service responds, Prometheus confirms it is collecting metrics, and Grafana confirms it is visualizing them.

### 1. Ansible playbook run, ending with the smoke test response

Shows the final task of the playbook printing the service's JSON response directly to the console, satisfying the technical challenge's explicit validation requirement.

**[Screenshot placeholder: terminal output of `ansible-playbook`, showing the `smoke-test` role's final task]**

### 2. Prometheus, Targets page

Confirms the scrape configuration is active and the service is being monitored continuously, not just queried once manually.

**[Screenshot placeholder: Prometheus Targets page, showing `http-server-projeto-korp` as UP]**

### 3. Grafana, Data Sources page

Confirms the Prometheus data source was connected automatically through provisioning files, not configured manually through the UI.

**[Screenshot placeholder: Grafana Data Sources page, showing Prometheus connected]**

### 4. Grafana Dashboard

The final deliverable of the monitoring requirement: both required metrics visualized side by side.

**[Screenshot placeholder: Grafana dashboard "Projeto Korp - Overview", showing both panels]**

## Architectural Decisions

A few decisions were made where the technical challenge document left room for interpretation. Each is documented here with its rationale.

### The network is created by Ansible, not by Docker Compose

Declaring the network inside `docker-compose.yaml` would work, but it introduces a real operational risk: running `docker compose down` would remove a network that could, in a broader setup, be shared by other services. Creating it explicitly via Ansible (`community.docker.docker_network`) and referencing it as `external: true` in Compose keeps its lifecycle independent from the application stack, which is a common pattern in real infrastructure.

### Prometheus scrapes the service directly, not through Nginx

The `/metrics` endpoint is meant for machine-to-machine consumption, not for external or human access. Routing it through Nginx would add a layer with no functional benefit, since Nginx currently has no route configured for it. Prometheus reaches the service directly over the internal Docker network, which the "no port exposed to host" requirement does not prohibit, since it only restricts host-level access.

### Code delivery via `git clone`, not local file copy

Rather than copying files from the control machine into the VM, the `clone-repository` role clones the project's public GitHub repository directly. This was chosen deliberately to mirror a more realistic deployment flow, and to guarantee that the provisioned environment always reflects exactly what is published in the repository, eliminating any risk of divergence between local, uncommitted changes and what actually gets deployed.

### Nginx, Prometheus, and Grafana configuration files live in `/opt` and are deployed to `/etc`

Following the Linux Filesystem Hierarchy Standard, the cloned repository (source code) lives under `/opt/korp-lab`, while the configuration files actually consumed by each container are copied to `/etc/korp-lab/<service>` by dedicated Ansible roles. This keeps the git checkout and the deployed configuration as two clearly separated concerns, even though both ultimately originate from the same single source of truth: the repository.

### Nginx configuration is deployed via a dedicated Ansible role; Prometheus and Grafana follow the same pattern for consistency

The technical challenge document explicitly allows the Grafana dashboard to be configured manually, treating file-based provisioning as a bonus. Nginx configuration, however, has no such exception in the text. To keep the automation consistent and avoid two different delivery strategies for conceptually similar configuration files, all three services (Nginx, Prometheus, Grafana) are configured through dedicated, single-responsibility Ansible roles, even though strictly only Nginx required it.

### The `vagrant` user is added to the `docker` group

This is not strictly required for the Ansible-driven provisioning to work, since all Ansible tasks already run with elevated privileges via `become: true`. However, it follows Docker's official post-installation recommendation and makes manual inspection of the environment (`vagrant ssh` followed by `docker ps`, without `sudo`) considerably more convenient during a live demonstration.