$(document).ready(function () {
    let selectedClientId = null;

    // Carregar clientes da API
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
                                <button class="action-btn edit-btn" onclick="openModal('editar', ${cliente.id})">Editar</button>
                                <button class="action-btn delete-btn" onclick="openModal('deletar', ${cliente.id})">Deletar</button>
                            </div>
                        </td>
                    </tr>
                `);
            });
        },
        error: function () {
            alert('Erro ao carregar clientes.');
        }
    });

    // Função para abrir os modais
    window.openModal = function (action, id) {
        selectedClientId = id;

        if (action === 'editar') {
            // Buscar cliente pelo ID e preencher os campos
            $.ajax({
                url: `/api/clientes/${id}`,
                method: 'GET',
                success: function (cliente) {
                    $('#editName').val(cliente.nome);
                    $('#editNascimento').val(cliente.data_nascimento);
                    $('#editIdade').val(cliente.idade);
                    $('#editCidade').val(cliente.cidade);
                    $('#editEstado').val(cliente.estado);
                    $('#editPais').val(cliente.pais);
                    $('#editOcupacao').val(cliente.ocupacao);
                    $('#editModal').fadeIn();
                },
                error: function () {
                    alert("Erro ao buscar cliente para edição.");
                }
            });
        } else if (action === 'deletar') {
            $('#deleteModal p').text(`Tem certeza que deseja excluir este cliente?`);
            $('#deleteModal').fadeIn();
        }
    };

    // Fechar os modais ao clicar no X ou cancelar
    $('#closeModal, #cancelDelete, #closeDeleteModal').click(() => {
        $('.modal').fadeOut();
        selectedClientId = null;
    });

    // Confirmar exclusão
    $('#confirmDelete').click(() => {
        if (selectedClientId) {
            $.ajax({
                url: `/api/clientes/${selectedClientId}`,
                method: 'DELETE',
                success: function () {
                    alert('Cliente deletado com sucesso.');
                    location.reload();
                },
                error: function () {
                    alert('Erro ao deletar cliente.');
                }
            });
        }
        $('#deleteModal').fadeOut();
    });

    // Salvar alterações do modal de edição
    $('#saveChanges').click(() => {
        const clientData = {
            nome: $('#editName').val(),
            data_nascimento: $('#editNascimento').val(),
            idade: parseInt($('#editIdade').val()),
            cidade: $('#editCidade').val(),
            estado: $('#editEstado').val(),
            pais: $('#editPais').val(),
            ocupacao: $('#editOcupacao').val()
        };

        $.ajax({
            url: `/api/clientes/${selectedClientId}`,
            method: 'PUT',
            contentType: 'application/json',
            data: JSON.stringify(clientData),
            success: function () {
                alert('Cliente atualizado com sucesso.');
                location.reload();
            },
            error: function () {
                alert('Erro ao atualizar cliente.');
            }
        });

        $('#editModal').fadeOut();
    });

    // Fechar modais ao clicar fora do conteúdo
    $(window).click(function (event) {
        if ($(event.target).is('#editModal')) {
            $('#editModal').fadeOut();
        }
        if ($(event.target).is('#deleteModal')) {
            $('#deleteModal').fadeOut();
        }
    });
});
