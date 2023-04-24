package policy_test

import data.policy.allow

test_deny_by_default {
    not allow
}

test_allow_if_admin {
    allow with input as {
        "user": {
            "roles": ["admin"]
        }
    }
}

test_allow_for_public_paths {

    allow with input as {
        "request": {
            "method": "GET",
            "path": "/public"
        }
    }

    allow with input as {
        "request": {
            "method": "GET",
            "path": "/public/pictures"
        }
    }

}