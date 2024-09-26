let {
    base = "http://localhost:8080";
    reqId = uuid();
}

do {
    method = "GET";
    url = "$base/api/pokemon/:id";
    params = {
        "id": 1
    };
    headers = {
        "Content-Type": "application/json",
        "request-id": reqId
    };
}