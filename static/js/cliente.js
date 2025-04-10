$(document).ready(function () {
    $('#createClienteBtn').click(function () {
        const name = $('#nome').val().trim();
        const data_nascimento = $('#data_nascimento').val().trim();
        const idade = $('#idade').val().trim();
        const cidade = $('#cidade').val().trim();
        const estado = $('#estado').val().trim();
        const pais = $('#pais').val().trim();
        const ocupacao = $('#ocupacao').val().trim();

        if (!name || !data_nascimento || !idade || !cidade || !estado || !pais || !ocupacao) {
            $('#error-msg').text("Todos os campos devem ser preenchidos.");
            return;
        }

        // Se passou na validação, prossegue com a requisição
        $.ajax({
            url: '/add-cliente',
            method: 'POST',
            contentType: 'application/json', // <- ESSA LINHA É IMPORTANTE
            data: JSON.stringify({
                nome: name,
                data_nascimento: data_nascimento,
                idade: idade,
                cidade: cidade,
                estado: estado,
                pais: pais,
                ocupacao: ocupacao
            }),
            success: function (response) {
                $('#msg').text(response);
                $('#error-msg').text("");
            },
            error: function (xhr) { 
                $('#error-msg').text("Erro ao adicionar cliente.");
            }
        });
        
        
    });
});