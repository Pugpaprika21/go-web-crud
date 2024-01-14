const app_url = $("#app_url").val();

const onSubmitFormCreateUser = () => {
    const username = $("#username").val();
    const password = $("#password").val();
    const fd = new FormData($("#form-create-user")[0]);

    if (username == "") {
        Swal.fire({
            title: "Username not Valid.",
            text: "",
            icon: "warning",
            timer: 1000
        });
        return;
    }

    if (password == "") {
        Swal.fire({
            title: "Password not Valid.",
            text: "",
            icon: "warning",
            timer: 1000
        });
        return;
    }

    if (fd.get("profile").name == "") {
        Swal.fire({
            title: "Profile not Valid.",
            text: "",
            icon: "warning",
            timer: 1000
        });
        return;
    }

    fd.append("username", username);
    fd.append("password", password);

    $.ajax({
        type: "POST",
        url: app_url + "user/create",
        data: fd,
        processData: false,
        contentType: false,
        dataType: "json",
        success: function(resp) {
            if (resp.data != null) {
                Swal.fire({
                    title: resp.msg,
                    text: "",
                    icon: "success",
                    timer: 1000,
                }).then((result) => {
                    window.location.href = app_url + "user/";
                });
            }
        },
        error: function(resp) {}
    });
}