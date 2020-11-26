$("#submit-button").click(function() {
    let login = $("#login-field").val();
    let password = $("#password-field").val();

    fetch("/api/users/auth", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            login: login,
            password: password
        })
    }).then(response => {
        if (response.status === 200) {
            window.location.replace("/user");
        } else {
            response.text().then(text => {
                window.location.replace("/error?status=" +
                    response.status + "&message=" + text);
            });
        }
    }).catch(err => {
        window.location.replace("/error?status=" +
            err.status + "&message=" + err.message);
    });
});