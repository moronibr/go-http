$(document).ready(function (){
    $.ajax({
        url: '/api/clientes',
        method: 'GET',
        success: function (data) {
            data.forEach(function (cliente) {
                $('#clientes').append(`
                    <tr>
                        <td>${cliente.nome}</td>
                        <td>${cliente.data_nascimento}</td>
                        <td>${cliente.idade}</td>
                        <td>${cliente.cidade}</td>
                        <td>${cliente.estado}</td>
                        <td>${cliente.pais}</td>
                        <td>${cliente.ocupacao}</td>
                    </tr>
                `);     
                
                });
        }
    });
})