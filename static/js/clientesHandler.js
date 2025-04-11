$(document).ready(function () {
    let selectedClientId = null;
    let currentAction = null;

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
                            <td>
                                <div class="action-group">
                                    <button class="action-btn edit-btn" onclick="openModal('editar', '${cliente.nome}')">Editar</button>
                                    <button class="action-btn delete-btn" onclick="openModal('deletar', '${cliente.nome}')">Deletar</button>
                                </div>
                            </td>
                    </tr>
                `);
            });
        }
    });

    $('#cancelBtn').click(() => {
        $('#actionModal').hide();
        selectedClientId = null;
        currentAction = null;
    });

    $('#confirmBtn').click(() => {
        if (currentAction === 'editar') {
            alert('Editar cliente com ID: ' + selectedClientId);
            // Redirecionar ou abrir form de edição
        } else if (currentAction === 'deletar') {
            alert('Cliente deletado: ID ' + selectedClientId);
            // Chamar API de deletar, depois remover da tabela se quiser
        }
        $('#actionModal').hide();
    });
});

function openModal(action, nome) {
    selectedClientId = nome; // agora armazena o nome
    currentAction = action;
    $('#modalText').text(`Tem certeza que deseja ${action} o cliente ${nome}?`);
    $('#actionModal').fadeIn();
}
