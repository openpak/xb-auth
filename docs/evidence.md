# Evidence: xb-auth

Tests: `GOWORK=off go test -race ./...`
(use a scratch TMPDIR). Everything is unit/in-process; no console or emulator
has signed in against it.

Open items: `KDC.UserAccountBlob` builds the account blob for a provisioned
user, but no network path hands it to a console yet; MACS still delivers only
the machine account (Xombie's MACS delivered the user account instead). Which
of the two a retail console expects is unverified on hardware.
