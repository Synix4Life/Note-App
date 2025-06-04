function send(action) {
    const title = document.getElementById("title").value;
    const content = document.getElementById("content").value;

    fetch("/" + action, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ title: title, content: content })
    })
        .then(res => res.text())
        .then(data => {
            document.getElementById("message").innerText = data;
        });
}