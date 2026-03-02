# Dependency Update Report

Generated: 2026-03-02T08:17:45Z

## Applied Updates

| Dependency | From | To | Type | CVEs Fixed |
|------------|------|----|------|------------|
| k8s.io/api | v0.32.0 | v0.35.2 | minor | - |
| k8s.io/apimachinery | v0.32.0 | v0.35.2 | minor | - |
| k8s.io/client-go | v0.32.0 | v0.35.2 | minor | - |
| docker.io/library/golang:1.26-alpine | 1.26-alpine | 1.26.0-alpine3.23 | tag | - |
| docker.io/library/alpine:3.23 | 3.23 | 3.23.3 | patch | - |

## Skipped Dependencies

| Dependency | Current | Latest | Reason |
|------------|---------|--------|--------|
| github.com/spf13/cobra | v1.10.2 | v1.10.2 | up to date |
| github.com/davecgh/go-spew | v1.1.2-0.20180830191138-d8f796af33cc | - | up to date |
| github.com/emicklei/go-restful/v3 | v3.11.0 | - | up to date |
| github.com/fxamacker/cbor/v2 | v2.7.0 | - | up to date |
| github.com/go-logr/logr | v1.4.2 | - | up to date |
| github.com/go-openapi/jsonpointer | v0.21.0 | - | up to date |
| github.com/go-openapi/jsonreference | v0.20.2 | - | up to date |
| github.com/go-openapi/swag | v0.23.0 | - | up to date |
| github.com/gogo/protobuf | v1.3.2 | - | up to date |
| github.com/golang/protobuf | v1.5.4 | - | up to date |
| github.com/google/gnostic-models | v0.6.8 | - | up to date |
| github.com/google/go-cmp | v0.6.0 | - | up to date |
| github.com/google/gofuzz | v1.2.0 | - | up to date |
| github.com/google/uuid | v1.6.0 | - | up to date |
| github.com/gorilla/websocket | v1.5.0 | - | up to date |
| github.com/inconshreveable/mousetrap | v1.1.0 | - | up to date |
| github.com/josharian/intern | v1.0.0 | - | up to date |
| github.com/json-iterator/go | v1.1.12 | - | up to date |
| github.com/mailru/easyjson | v0.7.7 | - | up to date |
| github.com/moby/spdystream | v0.5.0 | - | up to date |
| github.com/modern-go/concurrent | v0.0.0-20180306012644-bacd9c7ef1dd | - | up to date |
| github.com/modern-go/reflect2 | v1.0.2 | - | up to date |
| github.com/munnerz/goautoneg | v0.0.0-20191010083416-a7dc8b61c822 | - | up to date |
| github.com/mxk/go-flowrate | v0.0.0-20140419014527-cca7078d478f | - | up to date |
| github.com/pkg/errors | v0.9.1 | - | up to date |
| github.com/spf13/pflag | v1.0.9 | - | up to date |
| github.com/x448/float16 | v0.8.4 | - | up to date |
| golang.org/x/net | v0.30.0 | - | up to date |
| golang.org/x/oauth2 | v0.23.0 | - | up to date |
| golang.org/x/sys | v0.26.0 | - | up to date |
| golang.org/x/term | v0.25.0 | - | up to date |
| golang.org/x/text | v0.19.0 | - | up to date |
| golang.org/x/time | v0.7.0 | - | up to date |
| google.golang.org/protobuf | v1.35.1 | - | up to date |
| gopkg.in/evanphx/json-patch.v4 | v4.12.0 | - | up to date |
| gopkg.in/inf.v0 | v0.9.1 | - | up to date |
| gopkg.in/yaml.v3 | v3.0.1 | - | up to date |
| k8s.io/klog/v2 | v2.130.1 | - | up to date |
| k8s.io/kube-openapi | v0.0.0-20241105132330-32ad38e42d3f | - | up to date |
| k8s.io/utils | v0.0.0-20241104100929-3ea5e8cea738 | - | up to date |
| sigs.k8s.io/json | v0.0.0-20241010143419-9aa6b5e7a4b3 | - | up to date |
| sigs.k8s.io/structured-merge-diff/v4 | v4.4.2 | - | up to date |
| sigs.k8s.io/yaml | v1.4.0 | - | up to date |

## Verification

**Status: FAILED**

verification failed; patch still provided for review.

```
=== go test ./... (/home/kai/repositories/hasteward) ===
go: downloading k8s.io/apimachinery v0.35.2
go: downloading github.com/robfig/cron/v3 v3.0.1
go: downloading k8s.io/api v0.35.2
go: downloading k8s.io/client-go v0.35.2
go: downloading sigs.k8s.io/controller-runtime v0.23.1
go: downloading github.com/prometheus/client_golang v1.23.2
go: downloading github.com/spf13/cobra v1.10.2
go: downloading k8s.io/klog/v2 v2.130.1
go: downloading k8s.io/kube-openapi v0.0.0-20250910181357-589584f1c912
go: downloading sigs.k8s.io/structured-merge-diff/v6 v6.3.2-0.20260122202528-d9cc6641c482
go: downloading k8s.io/utils v0.0.0-20251002143259-bc988d571ff4
go: downloading sigs.k8s.io/randfill v1.0.0
go: downloading github.com/evanphx/json-patch/v5 v5.9.11
go: downloading github.com/go-logr/logr v1.4.3
go: downloading github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822
go: downloading golang.org/x/net v0.47.0
go: downloading github.com/spf13/pflag v1.0.9
go: downloading golang.org/x/term v0.37.0
go: downloading github.com/gorilla/websocket v1.5.4-0.20250319132907-e064f32e3674
go: downloading sigs.k8s.io/json v0.0.0-20250730193827-2d320260d730
go: downloading github.com/beorn7/perks v1.0.1
go: downloading github.com/cespare/xxhash/v2 v2.3.0
go: downloading github.com/prometheus/client_model v0.6.2
go: downloading github.com/prometheus/common v0.66.1
go: downloading github.com/prometheus/procfs v0.16.1
go: downloading google.golang.org/protobuf v1.36.8
go: downloading github.com/json-iterator/go v1.1.12
go: downloading go.yaml.in/yaml/v2 v2.4.3
go: downloading gopkg.in/inf.v0 v0.9.1
go: downloading github.com/fsnotify/fsnotify v1.9.0
go: downloading github.com/google/gnostic-models v0.7.0
go: downloading golang.org/x/time v0.9.0
go: downloading github.com/fxamacker/cbor/v2 v2.9.0
go: downloading golang.org/x/oauth2 v0.30.0
go: downloading golang.org/x/sys v0.38.0
go: downloading sigs.k8s.io/yaml v1.6.0
go: downloading github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd
go: downloading github.com/modern-go/reflect2 v1.0.3-0.20250322232337-35a7c28c31ee
go: downloading gomodules.xyz/jsonpatch/v2 v2.4.0
go: downloading k8s.io/apiextensions-apiserver v0.35.0
go: downloading go.yaml.in/yaml/v3 v3.0.4
go: downloading github.com/go-openapi/jsonreference v0.20.2
go: downloading github.com/go-openapi/swag v0.23.0
go: downloading github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc
go: downloading github.com/x448/float16 v0.8.4
go: downloading github.com/moby/spdystream v0.5.0
go: downloading golang.org/x/text v0.31.0
go: downloading github.com/google/btree v1.1.3
go: downloading github.com/google/uuid v1.6.0
go: downloading golang.org/x/sync v0.18.0
go: downloading github.com/go-openapi/jsonpointer v0.21.0
go: downloading gopkg.in/evanphx/json-patch.v4 v4.13.0
go: downloading github.com/mailru/easyjson v0.7.7
go: downloading gopkg.in/yaml.v3 v3.0.1
go: downloading github.com/mxk/go-flowrate v0.0.0-20140419014527-cca7078d478f
go: downloading github.com/pmezard/go-difflib v1.0.0
go: downloading github.com/emicklei/go-restful/v3 v3.12.2
go: downloading github.com/josharian/intern v1.0.0
?   	gitlab.prplanit.com/precisionplanit/hasteward	[no test files]
?   	gitlab.prplanit.com/precisionplanit/hasteward/api/v1alpha1	[no test files]
?   	gitlab.prplanit.com/precisionplanit/hasteward/common	[no test files]
?   	gitlab.prplanit.com/precisionplanit/hasteward/controller	[no test files]
?   	gitlab.prplanit.com/precisionplanit/hasteward/engine	[no test files]
?   	gitlab.prplanit.com/precisionplanit/hasteward/engine/cnpg	[no test files]
?   	gitlab.prplanit.com/precisionplanit/hasteward/engine/galera	[no test files]
?   	gitlab.prplanit.com/precisionplanit/hasteward/k8s	[no test files]
?   	gitlab.prplanit.com/precisionplanit/hasteward/metrics	[no test files]
?   	gitlab.prplanit.com/precisionplanit/hasteward/output	[no test files]
?   	gitlab.prplanit.com/precisionplanit/hasteward/restic	[no test files]

=== govulncheck ./... (/home/kai/repositories/hasteward) ===
go: downloading golang.org/x/vuln v1.1.4
go: downloading golang.org/x/telemetry v0.0.0-20240522233618-39ace7a40ae7
go: downloading golang.org/x/mod v0.22.0
go: downloading golang.org/x/tools v0.29.0
go: downloading golang.org/x/sync v0.10.0
=== Symbol Results ===

Vulnerability #1: GO-2026-4341
    Memory exhaustion in query parameter parsing in net/url
  More info: https://pkg.go.dev/vuln/GO-2026-4341
  Standard library
    Found in: net/url@go1.25
    Fixed in: net/url@go1.25.6
    Example traces found:
      #1: k8s/client.go:61:16: k8s.Init calls sync.Once.Do, which eventually calls url.ParseQuery
      #2: k8s/client.go:61:16: k8s.Init calls sync.Once.Do, which eventually calls url.URL.Query

Vulnerability #2: GO-2026-4340
    Handshake messages may be processed at the incorrect encryption level in
    crypto/tls
  More info: https://pkg.go.dev/vuln/GO-2026-4340
  Standard library
    Found in: crypto/tls@go1.25
    Fixed in: crypto/tls@go1.25.6
    Example traces found:
      #1: k8s/client.go:61:16: k8s.Init calls sync.Once.Do, which eventually calls tls.Conn.HandshakeContext
      #2: engine/cnpg/heal.go:363:23: cnpg.Engine.logHealPodOutput calls io.ReadAll, which eventually calls tls.Conn.Read
      #3: k8s/exec.go:111:14: k8s.ExecCommandWithEnv calls fmt.Fprintf, which calls tls.Conn.Write
      #4: engine/cnpg/triage.go:179:32: cnpg.Engine.triageCollect calls rest.Request.DoRaw, which eventually calls tls.Dialer.DialContext

Vulnerability #3: GO-2026-4337
    Unexpected session resumption in crypto/tls
  More info: https://pkg.go.dev/vuln/GO-2026-4337
  Standard library
    Found in: crypto/tls@go1.25
    Fixed in: crypto/tls@go1.25.7
    Example traces found:
      #1: k8s/client.go:61:16: k8s.Init calls sync.Once.Do, which eventually calls tls.Conn.HandshakeContext
      #2: engine/cnpg/heal.go:363:23: cnpg.Engine.logHealPodOutput calls io.ReadAll, which eventually calls tls.Conn.Read
      #3: k8s/exec.go:111:14: k8s.ExecCommandWithEnv calls fmt.Fprintf, which calls tls.Conn.Write
      #4: engine/cnpg/triage.go:179:32: cnpg.Engine.triageCollect calls rest.Request.DoRaw, which eventually calls tls.Dialer.DialContext

Vulnerability #4: GO-2025-4175
    Improper application of excluded DNS name constraints when verifying
    wildcard names in crypto/x509
  More info: https://pkg.go.dev/vuln/GO-2025-4175
  Standard library
    Found in: crypto/x509@go1.25
    Fixed in: crypto/x509@go1.25.5
    Example traces found:
      #1: k8s/exec.go:111:14: k8s.ExecCommandWithEnv calls fmt.Fprintf, which eventually calls x509.Certificate.Verify

Vulnerability #5: GO-2025-4155
    Excessive resource consumption when printing error string for host
    certificate validation in crypto/x509
  More info: https://pkg.go.dev/vuln/GO-2025-4155
  Standard library
    Found in: crypto/x509@go1.25
    Fixed in: crypto/x509@go1.25.5
    Example traces found:
      #1: k8s/exec.go:111:14: k8s.ExecCommandWithEnv calls fmt.Fprintf, which eventually calls x509.Certificate.Verify
      #2: k8s/exec.go:111:14: k8s.ExecCommandWithEnv calls fmt.Fprintf, which eventually calls x509.Certificate.VerifyHostname

Vulnerability #6: GO-2025-4013
    Panic when validating certificates with DSA public keys in crypto/x509
  More info: https://pkg.go.dev/vuln/GO-2025-4013
  Standard library
    Found in: crypto/x509@go1.25
    Fixed in: crypto/x509@go1.25.2
    Example traces found:
      #1: k8s/exec.go:111:14: k8s.ExecCommandWithEnv calls fmt.Fprintf, which eventually calls x509.Certificate.Verify

Vulnerability #7: GO-2025-4012
    Lack of limit when parsing cookies can cause memory exhaustion in net/http
  More info: https://pkg.go.dev/vuln/GO-2025-4012
  Standard library
    Found in: net/http@go1.25
    Fixed in: net/http@go1.25.2
    Example traces found:
      #1: engine/cnpg/triage.go:179:32: cnpg.Engine.triageCollect calls rest.Request.DoRaw, which eventually calls http.Client.Do

Vulnerability #8: GO-2025-4011
    Parsing DER payload can cause memory exhaustion in encoding/asn1
  More info: https://pkg.go.dev/vuln/GO-2025-4011
  Standard library
    Found in: encoding/asn1@go1.25
    Fixed in: encoding/asn1@go1.25.2
    Example traces found:
      #1: k8s/client.go:61:16: k8s.Init calls sync.Once.Do, which eventually calls asn1.Unmarshal

Vulnerability #9: GO-2025-4010
    Insufficient validation of bracketed IPv6 hostnames in net/url
  More info: https://pkg.go.dev/vuln/GO-2025-4010
  Standard library
    Found in: net/url@go1.25
    Fixed in: net/url@go1.25.2
    Example traces found:
      #1: k8s/exec.go:94:31: k8s.ExecStream calls remotecommand.spdyStreamExecutor.StreamWithContext, which eventually calls url.Parse
      #2: k8s/client.go:77:49: k8s.Init calls clientcmd.DeferredLoadingClientConfig.ClientConfig, which eventually calls url.ParseRequestURI
      #3: engine/cnpg/triage.go:179:32: cnpg.Engine.triageCollect calls rest.Request.DoRaw, which eventually calls url.URL.Parse

Vulnerability #10: GO-2025-4009
    Quadratic complexity when parsing some invalid inputs in encoding/pem
  More info: https://pkg.go.dev/vuln/GO-2025-4009
  Standard library
    Found in: encoding/pem@go1.25
    Fixed in: encoding/pem@go1.25.2
    Example traces found:
      #1: k8s/exec.go:89:44: k8s.ExecStream calls remotecommand.NewSPDYExecutor, which eventually calls pem.Decode

Vulnerability #11: GO-2025-4008
    ALPN negotiation error contains attacker controlled information in
    crypto/tls
  More info: https://pkg.go.dev/vuln/GO-2025-4008
  Standard library
    Found in: crypto/tls@go1.25
    Fixed in: crypto/tls@go1.25.2
    Example traces found:
      #1: k8s/client.go:61:16: k8s.Init calls sync.Once.Do, which eventually calls tls.Conn.HandshakeContext
      #2: engine/cnpg/heal.go:363:23: cnpg.Engine.logHealPodOutput calls io.ReadAll, which eventually calls tls.Conn.Read
      #3: k8s/exec.go:111:14: k8s.ExecCommandWithEnv calls fmt.Fprintf, which calls tls.Conn.Write
      #4: engine/cnpg/triage.go:179:32: cnpg.Engine.triageCollect calls rest.Request.DoRaw, which eventually calls tls.Dialer.DialContext

Vulnerability #12: GO-2025-4007
    Quadratic complexity when checking name constraints in crypto/x509
  More info: https://pkg.go.dev/vuln/GO-2025-4007
  Standard library
    Found in: crypto/x509@go1.25
    Fixed in: crypto/x509@go1.25.3
    Example traces found:
      #1: k8s/client.go:61:16: k8s.Init calls sync.Once.Do, which eventually calls x509.CertPool.AppendCertsFromPEM
      #2: k8s/exec.go:111:14: k8s.ExecCommandWithEnv calls fmt.Fprintf, which eventually calls x509.Certificate.Verify
      #3: k8s/client.go:61:16: k8s.Init calls sync.Once.Do, which eventually calls x509.CreateCertificate
      #4: k8s/client.go:61:16: k8s.Init calls sync.Once.Do, which eventually calls x509.MarshalPKCS1PrivateKey
      #5: k8s/client.go:61:16: k8s.Init calls sync.Once.Do, which eventually calls x509.ParseCertificate
      #6: k8s/exec.go:89:44: k8s.ExecStream calls remotecommand.NewSPDYExecutor, which eventually calls x509.ParseECPrivateKey
      #7: k8s/exec.go:89:44: k8s.ExecStream calls remotecommand.NewSPDYExecutor, which eventually calls x509.ParsePKCS1PrivateKey
      #8: k8s/exec.go:89:44: k8s.ExecStream calls remotecommand.NewSPDYExecutor, which eventually calls x509.ParsePKCS8PrivateKey

Your code is affected by 12 vulnerabilities from the Go standard library.
This scan also found 2 vulnerabilities in packages you import and 3
vulnerabilities in modules you require, but your code doesn't appear to call
these vulnerabilities.
Use '-show verbose' for more details.
exit status 3

```

## How to apply

```bash
git apply deps.patch
```

## Verify locally

```bash
go test ./... && govulncheck ./...
```
