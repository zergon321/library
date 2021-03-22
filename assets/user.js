$("#rent-button").click(function() {
    let bookID = $("#book-select").val();

    fetch("/ajax/user/rent-book?book_id=" + bookID, {
        method: "GET",
    }).then(response => {
        if (response.status === 200) {
            window.location.reload();
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

$("#order-button").click(function() {
    window.location.replace("/order");
});