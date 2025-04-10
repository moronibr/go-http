$(document).ready(function () {
    $("#loginBtn").click(function () {
      const name = $("#name").val();
      const password = $("#password").val();
  
      $.post("/login", { name, password })
        .done(function () {
          // Redirecionar após login com sucesso
          window.location.href = "/index";
        })
        .fail(function (xhr) {
          $("#error-msg").text(xhr.responseText);
        });
    });
  });
  