# SaaS
services as a service

This service will keep a running list of servers for various services
servers ping a machine with udp, pinging the machine with tcp will demand
a machine for a service. SaaS machines are services themselves. SaaS
machines should share their service lists.

* service sends udp saying it's alive to a network
* SaaS machine picks it up and adds it to a list of servers in a list of services
* TCP message to a server returns a machine that can perform a service
* SaaS machines come online and send out a udp message asking for SaaS machines
* TCP message between SaaS machines retrieves all the known services of that machine
