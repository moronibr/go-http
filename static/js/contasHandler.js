$(document).ready(function () {
    $.ajax({
        url: '/api/contas', // novo endpoint
        method : 'GET',
        success: function (data) {
            data.forEach (function (conta) {
                $('#contas').append(`
                    <tr>
                        <td>${conta.name}</td>
                        <td>${conta.email}</td>
                        <td>${conta.registrationNumber}</td>
                    </tr>
                `);
            });
        }
    });
});