$(document).ready(function () {
    $('#addAccountBtn').click(function () {
        const name = $('#name').val().trim();
        const email = $('#email').val().trim();
        const password = $('#password').val().trim();
        const registrationNumber = $('#registrationNumber').val().trim();
        const status = $('#status').val();

        if (!name || !email || !password || !registrationNumber) {
            $('#error-msg').text("Todos os campos devem ser preenchidos.");
            return;
        }

        // Se passou na validação, prossegue com a requisição
        $.ajax({
            url: '/add-account',
            method: 'POST',
            data: {
                name,
                email,
                password,
                registrationNumber,
                status
            },
            success: function (response) {
                $('#msg').text(response);
                $('#error-msg').text("");
            },
            error: function (xhr) {
                $('#error-msg').text("Erro ao adicionar conta.");
            }
        });
    });
});
