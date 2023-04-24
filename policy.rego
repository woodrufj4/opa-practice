package policy

default allow = false

allow = true {
    # Always allow admins
    input.user.roles[_] = "admin"
}

allow = true {
    # Regardless of role, allow access to the
    # public endpoints.
    startswith(input.request.path, "/public")
    input.request.method == "GET"
}