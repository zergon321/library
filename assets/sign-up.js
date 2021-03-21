$("#submit-button").click(function() {
    let name = $("#name-field").val();
    let surname = $("#surname-field").val();
    let nickname = $("#nickname-field").val();
    let patronim = $("#patronim-field").val();
    let email = $("#email-field").val();
    let password = $("#password-field").val();
    let confirmPassword = $("#confirm-password-field").val();

    fetch("/ajax/user/sign-up", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            name: name,
            surname: surname,
            patronim: patronim,
            nickname: nickname,
            email: email,
            password: password,
            confirm_password: confirmPassword
        })
    }).then(response => {
        if (response.status === 202) {
            response.json().then(respBody => {
                window.location.replace("/signed-up?personal_number=" +
                    respBody.personal_number + "&nickname=" + nickname);
            });
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

    return false;
});