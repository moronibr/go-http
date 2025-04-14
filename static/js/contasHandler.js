$(document).ready(function () {
    $.ajax({
        url: '/api/contas',
        method: 'GET',
        success: function (data) {
            data.forEach(function (conta) {
                $('#contas').append(`
                    <tr>
                        <td>${conta.name}</td>
                        <td>${conta.email}</td>
                        <td>${conta.registrationNumber}</td>
                        <td>
                            <button class="edit-btn" data-id="${conta.id}" data-name="${conta.name}" data-email="${conta.email}" data-reg="${conta.registrationNumber}">Editar</button>
                            <button class="delete-btn" data-id="${conta.id}">Deletar</button>
                        </td>
                    </tr>
                `);
            });

            // Botão Editar
            $('.edit-btn').click(function () {
                const id = $(this).data('id');
                const name = $(this).data('name');
                const email = $(this).data('email');
                const reg = $(this).data('reg');

                $('#editId').val(id);
                $('#editName').val(name);
                $('#editEmail').val(email);
                $('#editReg').val(reg);

                $('#editModal').fadeIn();
            });

            // Botão Deletar - agora usando modal estilizado
            $('.delete-btn').click(function () {
                const id = $(this).data('id');
                $('#deleteId').val(id);
                $('#deleteModal').fadeIn();
            });
        }
    });

    // Fechar modal de edição
    $('#closeModal').click(function () {
        $('#editModal').fadeOut();
    });

    // Salvar alterações
    $('#saveChanges').click(function () {
        const id = $('#editId').val();
        const updatedConta = {
            name: $('#editName').val(),
            email: $('#editEmail').val(),
            registrationNumber: $('#editReg').val()
        };

        $.ajax({
            url: `/api/contas/${id}`,
            method: 'PUT',
            contentType: 'application/json',
            data: JSON.stringify(updatedConta),
            success: function () {
                alert('Conta atualizada com sucesso!');
                $('#editModal').fadeOut();
                location.reload();
            },
            error: function () {
                alert('Erro ao atualizar conta.');
            }
        });
    });

    // Fechar modal de delete
    $('#closeDeleteModal, #cancelDelete').click(function () {
        $('#deleteModal').fadeOut();
    });

    // Confirmar deletar
    $('#confirmDelete').click(function () {
        const id = $('#deleteId').val();

        $.ajax({
            url: `/api/contas/${id}`,
            method: 'DELETE',
            success: function () {
                $('#deleteModal').fadeOut();
                location.reload();
            },
            error: function () {
                alert('Erro ao deletar.');
            }
        });
    });
});
