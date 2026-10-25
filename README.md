# Xbox Auth Server

OpenPak's Xbox Auth Server.

| | |
| --- | --- |
| Network | Original Xbox |
| Role | Original Xbox Kerberos KDC: MACS machine proof, AS and TGS sign-in, gateway tickets. UDP 88 plus a private health listener |
| Depends on | [`xb-account`](https://github.com/openpak/xb-account), [`xb-common`](https://github.com/openpak/xb-common) (Go module) |
| Used by | [`xb-gateway`](https://github.com/openpak/xb-gateway) (shares the SG keyring) |
| Image | `ghcr.io/openpak/xb-auth` |
| Based on | [Xombie `services/kdc`](https://github.com/XombieOnline/xombie) (AGPL-3.0) |
| Licence | AGPL-3.0-only |

## What it serves

| Flow | What |
| --- | --- |
| MACS | Console proves its machine key and receives its machine account; no ticket is issued |
| AS | Pre-authenticated sign-in for users and machine accounts |
| TGS | Validates the TGT and issues a gateway ticket naming the public gateway |

## Run

The service runs as a container from `deploy/openpak-xb-auth.container`.
Configuration is environment variables; secrets are files. Build the image:

```sh
podman build --secret id=netrc,src=$HOME/.netrc -t ghcr.io/openpak/xb-auth:dev .
```

## Docs

- [Protocol](docs/protocol.md)
- [Evidence](docs/evidence.md)
- [Deploy](docs/deploy.md)
- [Notes](docs/notes.md)

## Licence

AGPL-3.0-only. See [LICENSE](LICENSE).
