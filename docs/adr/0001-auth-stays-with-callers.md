# Token handling stays with callers

The client does not ship a hand-written constructor that fetches and refreshes OAuth tokens, even though every caller has to wire auth itself. Its only consumer is terraform-provider-camunda, so such a module would have one adapter: a hypothetical seam. The capability already exists one line away: `clientcredentials.Config.Client(ctx)` from `golang.org/x/oauth2`, passed as `Configuration.HTTPClient`, fetches, caches and refreshes tokens. A constructor here would wrap that call plus two defaults (token URL `https://login.cloud.camunda.io/oauth/token`, audience = API host), and would add the module's first dependency. Callers fix token handling in their own code instead.

Revisit if a second consumer appears that repeats the same wiring.
