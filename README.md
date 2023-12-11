<p align="center">
  <a href="" rel="noopener">
 <img width=300px src="https://raw.githubusercontent.com/taylow/awaik-backend/main/docs/img/awaikcat.png" alt="Awaik logo - fat cat sleeping"></a>
</p>

<!-- 
```
                       _ _               ____             _                  _ 
     /\               (_) |             |  _ \           | |                | | z
    /  \__      ____ _ _| | __  ______  | |_) | __ _  ___| | _____ _ __   __| |   z  ／l、
   / /\ \ \ /\ / / _` | | |/ / |______| |  _ < / _` |/ __| |/ / _ \ '_ \ / _` |    （ﾟ､ ｡７
  / ____ \ V  V / (_| | |   <           | |_) | (_| | (__|   <  __/ | | | (_| |      l、ﾞ~ヽ
 /_/    \_\_/\_/ \__,_|_|_|\_\          |____/ \__,_|\___|_|\_\___|_| |_|\__,_|      じしf_,)ノ
``` -->

<h3 align="center">Awaik - Don't hesitate; await the awake state!</h3>

<div align="center">

[![Status](https://img.shields.io/badge/status-active-success.svg)]()
[![GitHub Issues](https://img.shields.io/github/issues/taylow/awaik-backend.svg)](https://github.com/taylow/awaik-backend/issues)
[![GitHub Pull Requests](https://img.shields.io/github/issues-pr/taylow/awaik-backend.svg)](https://github.com/taylow/awaik-backend/pulls)
<!-- [![License](https://img.shields.io/badge/license-CC--BY--NC--SA--4.0-blue)](/LICENSE) -->

</div>

---

<p align="center"> This repository holds the backend code for Awaik.
    <br> 
</p>

## 📝 Table of Contents

- [About](#about)
- [Goals](#Goals)
- [Related Repositories](#related_repositories)
- [Getting Started](#getting_started)
- [Tests](#tests)
- [Built Using](#built_using)
- [Deployment](#deployment)
- [Todo](https://github.com/taylow/awaik-backend/blob/main/TODO.md)
- [Authors](#authors)
- [Acknowledgments](#acknowledgement)
- [Contributing](https://github.com/taylow/awaik-backend/blob/main/CONTRIBUTING.md)

## 🧐 About <a name = "about"></a>

Awaik is a service uptime monitoring tool that periodically sends requests to services, awaits a response, and checks that they are awake (get it... awaik 👈👈)

## 🥅 Goals <a name = "goals"></a>

Awaik is a fairly new project, but one that has been in the books for a while. It shares a nice balance between useful, real-world, challenging, yet simple in its goal.

Here is a list of features Awaik intends to support, and their current status!

✅ = implemented
🚧 = in development
⏰ = not yet started

Phase 0 - design and architect the project, event flow, service composition, etc.

- ✅ Microservice Structure
  - Using a similar approach to the go `database/sql` package, services now register themselves so that the entrypoint at `cmd/awaik/main.go` can dynamically load and host services as they grow
  - Services now have a clearly defined structure, utilising concepts from hexagonal architecture, domain-driven design, while keeping the codebase idiomatic.
- 🚧 Event Producers/Consumers
  - [x] Producing - Kafka integration is well under way, with the MonitorCommandService producing events with every write operation.
  - [ ] Consuming - this is currently in development, and should be ready to PR soon!
- 🚧 Persistence
  - Cassandra is a bit of a beast compared to other NoSQL/document-based databases, and the way the data is structured is quite important to its performance! I still have a bit of work to do before I have the right data structure, but this is at the highest priority right now!
- 🚧 DevOps & Community
  - [ ] Containerisation - write dockerfiles to encapsulate build and runtime environments
  - [ ] Repo automation - dependencies, builds, releases, etc.
  - [ ] Repository love - setup the things that are missing
- 😸 Fun stuff
  - Keeping the project fun is crucial to its success, that's why I added a cute little [terminal user interface](#tui-mode) and will be spending as much time researching and learning as necessary!

Phase 1 - with the groundwork done, the actual features can be implemented

- Infrastructure - move from proof-of-concept/dev configs to production-ready infrastructure (see [awaik-infrastructure](https://github.com/taylow/awaik-infrastructure))
- 🚧 Monitors - allow users to monitor services through numerous protocols (HTTP, ICMP, Browser Automation, etc.), across various regions
  - [x] CRUD operations for monitors
  - [ ] HTTP
  - [ ] ICMP
  - [ ] Port
  - [ ] Browser Automation
  - [ ] Multi-region
- ⏰ Heartbeats - provide a URL for services to send a pulse to
  - [ ] Heartbeat endpoint
- ⏰ Status Page - display uptime, downtime, and service health on a customisable status page
  - [ ] Editor (simplified)
  - [ ] Preview (availability reports)
  - [ ] Custom URL
- ⏰ Alerts & Notifications - send alerts through various integrations when services are down
  - [ ] Email
  - [ ] Slack
  - [ ] Discord
  - [ ] Many more

Phase 2 - once the core features are in, development of nice-to-have features can commence

- ⏰ Users & Teams - add a user and team layer
- ⏰ Incidents & Maintenance - allow users to report incidents, tag downtime, and schedule maintenance
- ⏰ Public API - manage all aspects of Awaik via HTTP requests, fully documented with OpenAPI
- ⏰ Secrets - allow users to store secrets for their monitors and heartbeats to allow the use of protected endpoints
- ⏰ Integrations & Automation - allow Awaik to connect to various services, allowing users to define how to react to unhealthy services

Phase 3 - SaaS

- ⏰ Infrastructure as Code
- ⏰ Scaling with Kubernetes
- ⏰ Venture into the reams of hosting Awaik as a SaaS platform for those who do not with to self-host to enjoy!

## 💾 Related Repositories <a name = "related_repositories"></a>

- [awaik](https://github.com/taylow/awaik) - Landing page
- [awaik-frontend](https://github.com/taylow/awaik-frontend) - Frontend repository for connecting to Awaik services
- [awaik-backend](https://github.com/taylow/awaik-backend) - Backend repository responsible for Awaik's functionality
- [awaik-infrastructure](https://github.com/taylow/awaik-infrastructure) - Scripts and configs for Awaik's infrastructure

## 🏁 Getting Started <a name = "getting_started"></a>

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See [deployment](#deployment) for notes on how to deploy the project on a live system.

This project is currently developed on macOS, and so all instructions are written as such. They will be updated overtime, or if you want to update them sooner open a PR!

### tui mode

Awaik backend is made up of many services. As a way to keep development fun and easy, Awaik comes with an optional terminal user interface!

Simply add the `--tui` flag and you will be greeted with the following:

<p align="center">
  <a href="" rel="noopener">
 <img src="https://raw.githubusercontent.com/taylow/awaik-backend/main/docs/img/tui.gif" alt="demo of the terminal user interface in use"></a>
</p>

```sh
make install
awaik --tui
```

### cli headless mode

To run Awaik headless, simply run the binary with either the flag `--all` for all services, or multiple `--service [SERVICE NAME]` flags indicating the various services to host!

```sh
make install
awaik --service MonitorCommandService --service MonitorQueryService
```

## 🔧 Running the tests <a name = "tests"></a>

Running tests is as running the following:

```sh
make test-setup
make test
```

Tests will also automatically run on pushes and pull requests.

## 🚀 Deployment <a name = "deployment"></a>

Deployment details will come once services and decisions are more solidified. Check back here every so often!

## ⛏️ Built Using <a name = "built_using"></a>

- [**Go**](http://golang.org/) - Backend code
- [**Cassandra**](http://cassandra.apache.org/) - Persistent data storage
- [**Kafka**](http://kafka.apache.org/) - Asynchronous message queues
- [**Hashicorp Vault**](https://www.vaultproject.io/) - Secret storage
- [**Envoy Proxy**](https://envoyproxy.io/) and [**Istio**](https://istio.io/) - Reverse proxy and service mesh
- [**Docker**](https://docker.com/) and [**Kubernetes**](http://kubernetes.io/) - Containerisation and orchestration of various services
- [**AWS**](http://aws.amazon.com/) - Infrastructure
- [**Terraform/OpenTofu**](https://opentofu.org/) - Infrastructure as Code
- [**React**](https://facebook.github.io/react/) - Frontend and internal web dashboards

## ✍️ Authors <a name = "authors"></a>

- [@taylow](https://github.com/taylow) - Founder

See also the list of [contributors](https://github.com/taylow/awaik-backend/contributors) who participated in this project.

## 🎉 Acknowledgements <a name = "acknowledgement"></a>

- Stack inspiration came from [Monzo](https://monzo.com/), as this is what they listed in a job opening, and I wanted to build something that used the same stack
- Hat tip to [Christian Selig](https://github.com/christianselig) for their [apollo-backend](https://github.com/christianselig/apollo-backend) repository; a fantastic example of a production-ready server written in Go
- Thanks to my friend [Connor O'Brien](https://connorobrienbusibddf.myportfolio.com/) for helping me come up with a silly name for it, after all, you can't truly start a project until you know the domain exists and the SEO isn't already saturated, right? right??
- Another thanks to my talented friend [Dylan Thomas](#) for the adorable cat logo!
