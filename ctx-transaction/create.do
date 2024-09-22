let {
    base = "http://localhost:8080";
}

do {
    url = "$base/orders";
    method = "POST";
    headers = {
        "Content-Type": "application/json"
    };
    body = `{
        "customer_name": "luis",
        "description": "example order",
        "order_lines": [
            {
                "name": "example",
                "quantity": 12
            }
        ]
    }`;
}