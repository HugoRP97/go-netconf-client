> :heavy_exclamation_mark: *Red Hat does not provide commercial support for the content of these repos*

```bash
#############################################################################
DISCLAIMER: THESE ARE UNSUPPORTED COMMUNITY TOOLS.

THE REFERENCES ARE PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
#############################################################################
```

# NETCONF

[![GoDoc](https://godoc.org/github.com/openshift-telco/go-netconf?status.svg)](https://godoc.org/github.com/openshift-telco/go-netconf)
[![Report Card](https://goreportcard.com/badge/github.com/openshift-telco/go-netconf)](https://goreportcard.com/report/github.com/openshift-telco/go-netconf)
[![Build Status](https://travis-ci.org/openshift-telco/go-netconf.png)](https://travis-ci.org/openshift-telco/go-netconf)
[![codecov](https://codecov.io/gh/openshift-telco/go-netconf/branch/main/graph/badge.svg)](https://codecov.io/gh/openshift-telco/go-netconf)

This library is a simple NETCONF client :
- [RFC6241](http://tools.ietf.org/html/rfc6241): **Network Configuration Protocol (NETCONF)** 
    - Support for the following RPC: `lock`, `unlock`, `edit-config`, `comit`, `validate`,`get`, `get-config`
    - Support for custom RPC
- [RFC6242](http://tools.ietf.org/html/rfc6242): **Using the NETCONF Protocol over Secure Shell (SSH)**
    - Support for username/password
    - Support for pub key
- [RFC5277](https://datatracker.ietf.org/doc/html/rfc5277): **NETCONF Event Notifications**
    - Support for `create-subscription`
    - No support for notification filtering
- Partially [RFC8641](https://datatracker.ietf.org/doc/html/rfc8641) and [RFC8639](https://datatracker.ietf.org/doc/html/rfc8639): **Subscription to YANG Notifications for Datastore Updates**
    - Support for `establish-subscription`
    - No support for `delete-subscription` -- TODO

#### Links
This client is an adaptation of the code taken from:
- https://github.com/andaru/netconf
- https://github.com/Juniper/go-netconf