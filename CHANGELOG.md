# Changelog

## v14 R2 — experimental security revision

- Restrict GitHub metadata redirects to the authorized HTTPS host.
- Check ancestor directories of installer destinations for pre-existing symbolic links.
- Preserve existing engine paths, configurations and shader caches.
- Add regression tests for redirect and installation safety.

Limitations: this is a partial review, not a comprehensive security audit. It does not establish the safety of external engines, build scripts or all Windows filesystem race conditions. Antivirus detections are unresolved and Windows runtime validation remains pending.
