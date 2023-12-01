<!-- <p align="center">
  <a href="" rel="noopener">
 <img width=200px height=200px src="https://i.imgur.com/6wj0hh6.jpg" alt="Project logo"></a>
</p> -->

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
- [Related Repositories](#related_repositories)
- [Getting Started](#getting_started)
- [Tests](#tests)
- [Built Using](#built_using)
- [TODO](../TODO.md)
- [Authors](#authors)
- [Acknowledgments](#acknowledgement)
<!-- - [Contributing](../CONTRIBUTING.md) -->
<!-- - [Deployment](#deployment) -->
<!-- - [Usage](#usage) -->

## 🧐 About <a name = "about"></a>

Awaik is a service uptime monitoring tool that periodically sends requests to services, awaits a response, and checks that they are awake (get it... awaik 👈👈)

## 💾 Related Repositories <a name = "related_repositories"></a>

- [awaik](https://github.com/taylow/awaik) - Landing page
- [awaik-frontend](https://github.com/taylow/awaik-frontend) - Frontend repository for connecting to Awaik services
- [awaik-backend](https://github.com/taylow/awaik-backend) - Backend repository responsible for Awaik's functionality
- [awaik-infrastructure](https://github.com/taylow/awaik-infrastructure) - Scripts and configs for Awaik's infrastructure

## 🏁 Getting Started <a name = "getting_started"></a>

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See [deployment](#deployment) for notes on how to deploy the project on a live system.

This project is currently developed on macOS, and so all instructions are written as such. They will be updated overtime, or if you want to update them sooner open a PR!

**COMING SOON**

## 🔧 Running the tests <a name = "tests"></a>

Running tests is as running the following:

```sh
make test-setup
make test
```

Tests will also automatically run on pushes and pull requests.

<!-- ## 🎈 Usage <a name="usage"></a> -->

<!-- ## 🚀 Deployment <a name = "deployment"></a> -->

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
