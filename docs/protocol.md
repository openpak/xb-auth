# Protocol: xb-auth

Flows (all in `internal/kdc`, structures in `xb-common/krb`):

- MACS (realm `MACS.XBOX.COM`): console proves its machine key (client-version
  MAC and encrypted timestamp) and receives its machine account `SN.<serial>`
  sealed under that key. No ticket is issued.
- AS: users and machine accounts sign in with pre-authentication. A user may
  add a machine TGT as compound identity; the reply is sealed with the
  compound key. Machine TGTs never work at the TGS.
- TGS: validates the TGT, authenticator and replay, re-checks the user with
  xb-account on every request, and issues a gateway ticket (sealed with the
  shared SG keyring) plus PA_XBOX_SERVICE_ADDRESS naming the public gateway.

Admission always comes from `xbaccountapi.Resolver` (ResolveGamertag,
ResolveUser, ResolveDevice) over mutual TLS. Unknown, banned, unlinked,
revoked or unreachable principals are refused; nothing is invented.
Devices additionally need proof material (`kdc.DeviceKeys`), fetched per MACS
request from xb-account (`GET /internal/v1/devices/{id}/key`, xb-auth peer
only, released only for an admitted device): without it MACS answers
"principal unknown".

Machine id = the 12-digit serial as a number, 16 lowercase hex digits
(see docs/go-service-contracts.md item 4).
