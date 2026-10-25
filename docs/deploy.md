# Deploy: xb-auth

## Configuration (environment; secrets are files)

| Variable | Meaning |
|---|---|
| `XB_AUTH_UDP_ADDR` | default `:88` (needs CAP_NET_BIND_SERVICE or the unprivileged-port sysctl) |
| `XB_AUTH_HEALTH_ADDR` | default `127.0.0.1:8088`; `/livez`, `/readyz` (listener only, not sign-in health) |
| `XB_GATEWAY_PUBLIC_IPV4` | advertised gateway; default HOST_IP; private/loopback/CGNAT refused |
| `XB_TGS_KEYRING_FILE` | TGT keys, 16 bytes each, local to xb-auth |
| `XB_SG_KEYRING_FILE` | gateway ticket keys, 16 bytes each, shared with xb-gateway |
| `XB_USER_KEY_SECRETS_FILE` | 32-byte secrets deriving user and machine-account keys |
| `XB_ACCOUNT_URL`, `XB_ACCOUNT_CA_FILE`, `XB_TLS_CERT_FILE`, `XB_TLS_KEY_FILE` | xb-account (https, TLS 1.3, client cert) |
| `XB_TGT_LIFETIME`, `XB_TICKET_LIFETIME`, `XB_CLOCK_SKEW`, `XB_REPLAY_MAX` | defaults 24h, 24h, 5m, 65536 |

Keyring files: `{"active":2,"keys":{"1":"<hex>","2":"<hex>"}}`. The id is the
key version carried in tickets; add a key and make it active, remove an old id
once its tickets have expired. Group/world-writable files, zero keys and the
published Xombie prototype keys are refused.

## Repository

- Module `github.com/openpak/xb-auth`, image `ghcr.io/openpak/xb-auth` (static, scratch, uid 10001).
- Shared packages: `github.com/openpak/xb-common` (private; see its README for access). Service map: private notes.
- Build and test: `GOWORK=off go build ./... && go vet ./... && go test -race ./...` (with `GOPRIVATE=github.com/openpak/*` and `gh auth setup-git`).
- Image: `podman build --secret id=netrc,src=$HOME/.netrc -t ghcr.io/openpak/xb-auth:dev .`
- Release: push a `vX.Y.Z` tag; `.github/workflows/release.yml` (tags and manual dispatch only) tests and publishes the multi-arch image. CI reads xb-common with the job's `GITHUB_TOKEN` through a netrc build secret; nothing secret is in git or in the image.
- `deploy/` holds the env example and a Podman quadlet example (digests in it are examples; production pins its own).
