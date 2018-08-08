var info = $('#info');
var guilds;

var socket = io();
$('#submit').click(function() {
    socket.emit('tokencheck', $('#inpt_token').val())
});

socket.on('login', function() {
    info.empty();
    info.append($('<p>').text('Logging in...'));
});

socket.on('response', function(data) {
    var dataDom = [
        $('<tr>').append([
            $('<td>').text('Valid'),
            $('<td>').text(data.Valid ? '✔️' : '❌'),
        ])
    ];

    if (data.Valid) {
        var appends = [
            $('<tr>').append([
                $('<td>').text('Account ID'),
                $('<td>').text(data.User.ID),
            ]),
            $('<tr>').append([
                $('<td>').text('Tag'),
                $('<td>').text(`${data.User.Username}#${data.User.Discriminator}`),
            ]),
            $('<tr>').append([
                $('<td>').text('Guilds Number'),
                $('<td>').text(data.NGuilds),
            ]),
            $('<tr>').append([
                $('<td>').text('Guilds'),
                $('<td>').attr('id', 't_c_guilds').append(
                    $('<p>')
                        .attr('id', 'p_collecting_data')
                        .text('Collecting Guild Data')
                ),
            ])
        ];

        appends.forEach(d => dataDom.push(d));
        var c = 1;
        setInterval(function() {
            if (c >= data.NGuilds)
                c = data.NGuilds;
            $('#p_collecting_data').text(`Collecting Guild Data ${c}/${data.NGuilds}`);
            c++;
        }, 40);
    }

    info.empty();
    info.append(
        $('<table>').attr('id', 't_data')
            .attr('style', '/*width:25%;*/ margin-top: 80px;')
            .attr('border', '1')
            .append(dataDom)
    );
})

socket.on('response_guilds', function(gdata) {
    guilds = gdata;
    var t_c_guilds = $('#t_c_guilds');
    t_c_guilds.empty();
    t_c_guilds.append(
        $('<button>')
            .attr('class', 'button')
            .attr('style', 'width: 100%;')
            .text('SHOW')
            .click(() => {
                window.open('/guildswindow.html', 'Guilds', 'width=800,height=500');
            })
    );
})