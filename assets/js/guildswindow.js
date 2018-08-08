var t = $('#t_guilds');
window.opener.guilds.forEach(g => {
    t.append(
        $('<tr>').append([
            $('<td>').text(g.ID),
            $('<td>').text(g.Name),
            $('<td>').text(g.OwnerID),
        ])
    );
})