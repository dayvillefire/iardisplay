// IAR Display logic

var incidentCache = {};

$(document).ready(function () {
    // Initial load
    populateIncidents();
    populateOnDuty();
    populateResponding();

    // Set repeating
    setInterval(populateResponding, 10500);
    setInterval(populateOnDuty, 10250)
    setInterval(populateIncidents, 10000);
});

function populateIncidents() {
    $.getJSON("/api/iar/incidents", {}, function (response) {
        //var body = '';
        $('#incidents').html('');
        x = response.incidents.length;
        if (x > 5) {
            x = 5;
        }
        for (var i = 0; i < x; i++) {
            var id = response.incidents[i].Id;
            var body = '<tr><td>';
            if (i == 0) {
                body += "<b>";
            }
            body += response.detail[id + 0].ArrivedOn.replace('T', ' ');
            if (i == 0) {
                body += "</b>";
            }
            body += "</td><td>";
            if (i == 0) {
                body += "<b>";
            }
            body += response.detail[id + 0].Body;
            if (i == 0) {
                body += "</b>";
            }
            body += '</td></tr>';
            $('#incidents').append(body);
        }
    });
    //body += '<tr><td>' + response[i].Body + '</td></tr>';
}

function populateOnDuty() {
    $.getJSON("/api/iar/schedule", {}, function (response) {
        var body = '';
        x = response.length;
        for (i = 0; i < x; i++) {
            body += '<tr>' +
                '<td>' + response[i].MemberName + '</td>' +
                '<td>' + response[i].MemberCat + '</td>' +
                '<td>' + response[i].InStationOrHome + '</td>' +
                '<td>' + response[i].MemberStation + '</td>' +
                '<td>' + response[i].UntilAt + '</td>' +
                '</tr>';
        }
        if (body == '') {
            body = '<tr><td colspan="5"><div align="center">No one scheduled</div></td></tr>';
        }
        $('#schedule').html(body);
        //console.log(body);
    });
}

function populateResponding() {
    $.getJSON("/api/iar/responding", {}, function (response) {
        var body = '';
        console.log("populateResponding(): " + JSON.stringify(response));
        x = response.length;
        for (i = 0; i < x; i++) {
            console.log(JSON.stringify(response[i]));
            body += '<tr>' +
                '<td>' + response[i].MemberFName + '</td>' +
                '<td>' + response[i].MemberCat + '</td>' +
                '<td>' + response[i].MemberStation + '</td>' +
                '<td>' + response[i].CallingTime + '</td>' +
                '<td>' + response[i].ETA + '</td>' +
                '</tr>';
        }
        if (body == '') {
            body = '<tr><td colspan="5"><div align="center">No one responding</div></td></tr>';
        }
        $('#responding').html(body);
        //console.log(body);
    });
}