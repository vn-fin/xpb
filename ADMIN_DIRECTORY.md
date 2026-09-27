# Administrative email directory RPC

`PermissionService.GetUserByEmail(GetUserByEmailRequest)` is an additive RPC for
trusted administrative workflows such as XNOBrain workspace deletion.

Request fields: `token` (administrator access token without a Bearer prefix),
`email` (exact target address). Response fields: `user_id`, `email`,
`email_verified`. The token belongs to the caller, not the target user. No caller
role, user ID, tenant, workspace, or URL is accepted as authorization input.

The service must verify the token and current issuer platform-admin role before
looking up the target. Exact lookup rejects ambiguous addresses. An unverified
address can be returned with `email_verified=false`; consumers must apply their
own verified-email requirement. Tokens and target email must not be logged.

Canonical status codes:

| Status | Meaning |
| --- | --- |
| Unauthenticated | Missing/invalid access token |
| PermissionDenied | Caller is not a platform admin |
| InvalidArgument | Malformed email/request |
| NotFound | No exact address match |
| FailedPrecondition | Ambiguous/mismatched identity |
| Unavailable / DeadlineExceeded | Dependency failure / bounded timeout |
| Unimplemented | Older auth deployment; privileged consumers fail closed |

All existing fields and RPC paths remain unchanged. New clients must handle old
servers; publishing this module does not itself deploy the implementation.
Update xno-auth and Control module pins only after a real reviewed module revision
is published. Local integration uses a Go workspace rather than guessed versions.

Regenerate the changed contract using the repository's generation tools:

```sh
protoc --go_out=. --go-grpc_out=. permission.proto
uv run --frozen python -m grpc_tools.protoc \
  -I . --python_out=xpb --grpc_python_out=xpb permission.proto
```

Generation used the installed protoc 34.0, protoc-gen-go 1.25.0,
protoc-gen-go-grpc 1.1.0 and the locked Python environment. Regeneration also
brings the existing Python stub's GetPictureFromToken RPC into alignment with
the pre-existing permission.proto source; no existing RPC is removed.
