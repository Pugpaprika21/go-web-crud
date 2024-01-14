const app_url = $("#app_url").val();

const onDelete = (userId) => {
    Swal.fire({
        title: "Are you sure?",
        text: "You won't be able to revert this!",
        icon: "warning",
        showCancelButton: true,
        confirmButtonColor: "#3085d6",
        cancelButtonColor: "#d33",
        confirmButtonText: "Yes, delete it!"
    }).then((result) => {
        if (result.isConfirmed) {
            $.ajax({
                type: "DELETE",
                dataType: "json",
                url: app_url + "user/delete/" + userId,
                success: function(resp) {
                    console.log(resp);
                    if (resp.data != null) {
                        Swal.fire({
                            title: "Deleted!",
                            text: "Your file has been deleted.",
                            icon: "success"
                        })
                        window.location.href = app_url + "user/";
                        return
                    }
                }
            });
        }
    });
}

const onEdit = (userId) => {
    window.location.href = app_url + "user/update/" + userId;
}